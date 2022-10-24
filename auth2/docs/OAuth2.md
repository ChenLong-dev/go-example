# OAuth 2.0

[OAuth 2.0 实战课](https://time.geekbang.org/column/intro/100053901?tab=catalog)

[Web Authorization Protocol (oauth)](https://datatracker.ietf.org/group/oauth/documents/)

## 基础篇
### OAuth 2.0 是什么？
OAuth 2.0 就是一种授权协议
### 为什么用 OAuth 2.0？
OAuth 2.0 这种授权协议，就是保证第三方（软件）只有在获得授权之后，才可以进一步访问授权者的数据。
### OAuth 2.0 是怎样运转的？
![img.png](img.png)
OAuth 2.0 授权的核心就是颁发访问令牌、使用访问令牌
### 为什么需要授权码？
- 资源拥有者 -> 小明
- 第三方软件 -> 小兔软件
- 授权服务 -> 京东商家开放平台的授权服务
- 受保护资源 -> 小明店铺在京东上面的订单

OAuth 诞生之初就是为了解决 Web 浏览器场景下的授权问题

![img_1.png](img_1.png)

### 授权码许可类型的通信过程
#### 间接通信
![img_2.png](img_2.png)
通过浏览器这个“中间人” 来 “搭线”的

#### 直接通信
![img_3.png](img_3.png)

### 授权服务：授权码和访问令牌的颁发流程是怎样的？
#### 授权服务的工作过程
```
Map<String,String> appMap =  new HashMap<String, String>();//模拟第三方软件注册之后的数据库存储

appMap.put("app_id","APPID_RABBIT");
appMap.put("app_secret","APPSECRET_RABBIT");
appMap.put("redirect_uri","http://localhost:8080/AppServlet-ch03");
appMap.put("scope","nickname address pic");
```
![img_4.png](img_4.png)
##### 过程一：颁发授权码 code
- 第一步，验证基本信息。
```
# 授权服务需要对回调地址做基本的校验
if(!appMap.get("redirect_uri").equals(redirectUri)){
    //回调地址不存在
}
```
- 第二步，验证权限范围（第一次）。
```
# 校验申请过来的权限。记住，此刻是第一次权限校验。
String scope = request.getParameter("scope");
if(!checkScope(scope)){
//超出注册的权限范围
}
```
- 第三步，生成授权请求页面。
  
只有用户登录了才可以对第三方软件进行授权，授权服务才能够获得用户信息并最终生成 code 和 app_id（第三方软件的应用标识） + user（资源拥有者标识）之间的对应关系
- 第四步，验证权限范围（第二次）。
```
# 校验用户输入的权限
String[] rscope =request.getParameterValues("rscope");

if(!checkScope(rscope)){
    //超出注册的权限范围
}
```
- 第五步，处理授权请求，生成授权码 code。
```
# 授权服务会校验响应类型 response_type 的值，response_type 有 code 和 token 两种类型的值。
# 代码要验证 response_type 的值是否为 code
String responseType = request.getParameter("response_type");
if("code".equals(responseType)){
  
}
```
在授权服务中，需要将生成的授权码 code 值与 app_id、user 进行关系映射。也就是说，一个授权码 code，表示某一个用户给某一个第三方软件进行授权
```
String code = generateCode(appId,"USERTEST");//模拟登录用户为USERTEST

private String generateCode(String appId,String user) {
  ...
  String code = strb.toString();
  codeMap.put(code,appId+"|"+user+"|"+System.currentTimeMillis());
  return code;
}
```
OAuth 2.0 规范建议授权码 code 值有效期为 10 分钟，并且一个授权码 code 只能被使用一次。不过根据经验呢，在生产环境中 code 的有效期一般不会超过 5 分钟。

同时，授权服务还需要将生成的授权码 code 跟已经授权的权限范围 rscope 进行绑定并存储，以便后续颁发访问令牌时，我们能够通过 code 值取出授权范围并与访问令牌绑定。因为第三方软件最终是通过访问令牌来请求受保护资源的。
```
Map<String,String[]> codeScopeMap =  new HashMap<String, String[]>();

codeScopeMap.put(code,rscope);//授权范围与授权码做绑定
```
- 第六步，重定向至第三方软件。
```
Map<String, String> params = new HashMap<String, String>();
params.put("code",code);

String toAppUrl = URLParamsUtil.appendParams(redirectUri,params);//构造第三方软件的回调地址，并重定向到该地址

response.sendRedirect(toAppUrl);//授权码流程的“第二次”重定向
```

##### 过程二：颁发访问令牌 access_token
- 第一步，验证第三方软件是否存在。
```
# 接收到的 grant_type 的类型为 authorization_code
String grantType = request.getParameter("grant_type");
if("authorization_code".equals(grantType)){
  
}
```
由于颁发访问令牌是通过后端通信完成的，所以这里除了要校验 app_id 外，还要校验 app_secret。
```
if(!appMap.get("app_id").equals(appId)){
    //app_id不存在
    }

    if(!appMap.get("app_secret").equals(appSecret)){
        //app_secret不合法
        }
```
- 第二步，验证授权码 code 值是否合法。
```
# code 值对应的 key 是 app_id 和 user 的组合值
String code = request.getParameter("code");
if(!isExistCode(code)){//验证code值
  //code不存在
  return;
}
codeMap.remove(code);//授权码一旦被使用，须立即作废
```
确认过授权码 code 值有效以后，应该立刻从存储中删除当前的 code 值，以防止第三方软件恶意使用一个失窃的授权码 code 值来请求授权服务。

- 第三步，生成访问令牌 access_token 值

访问令牌 access_token 的值，OAuth 2.0 规范中并没有明确规定，但必须符合三个原则：唯一性、不连续性、不可猜性

将访问令牌 access_token 值存储起来，并将其与第三方软件的应用标识 app_id 和资源拥有者标识 user 进行关系映射。也就是说，一个访问令牌 access_token 表示某一个用户给某一个第三方软件进行授权。

同时，授权服务还需要将授权范围跟访问令牌 access_token 做绑定。最后，还需要为该访问令牌设置一个过期时间 expires_in，比如 1 天。

```
Map<String,String[]> tokenScopeMap =  new HashMap<String, String[]>();

String accessToken = generateAccessToken(appId,"USERTEST");//生成访问令牌access_token的值
tokenScopeMap.put(accessToken,codeScopeMap.get(code));//授权范围与访问令牌绑定

//生成访问令牌的方法
private String generateAccessToken(String appId,String user){
  
  String accessToken = UUID.randomUUID().toString();
  String expires_in = "1";//1天时间过期
  tokenMap.put(accessToken,appId+"|"+user+"|"+System.currentTimeMillis()+"|"+expires_in);

  return accessToken;
}
```
#### 刷新令牌
刷新令牌也是给第三方软件使用的，同样需要遵循先颁发再使用的原则

##### 颁发刷新令牌
颁发刷新令牌和颁发访问令牌是一起实现的，都是在过程二的步骤三生成访问令牌 access_token 中生成的。也就是说，第三方软件得到一个访问令牌的同时，也会得到一个刷新令牌
```
Map<String,String> refreshTokenMap =  new HashMap<String, String>();

String refreshToken = generateRefreshToken(appId,"USERTEST");//生成刷新令牌refresh_token的值

private String generateRefreshToken(String appId,String user){

  String refreshToken = UUID.randomUUID().toString();

  refreshTokenMap.put(refreshToken,appId+"|"+user+"|"+System.currentTimeMillis());
  return refreshToken;
  
} 
```
##### 使用刷新令牌
- 第一步，接收刷新令牌请求，验证基本信息。
```
# 此时请求中的 grant_type 值为 refresh_token
String grantType = request.getParameter("grant_type");
if("refresh_token".equals(grantType)){

}
```
和颁发访问令牌前的验证流程一样，这里我们也需要验证第三方软件是否存在。需要注意的是，这里需要同时验证刷新令牌是否存在，目的就是要保证传过来的刷新令牌的合法性。
```
String refresh_token = request.getParameter("refresh_token");

if(!refreshTokenMap.containsKey(refresh_token)){
    //该refresh_token值不存在
    }
```
另外，我们还需要验证刷新令牌是否属于该第三方软件。授权服务是将颁发的刷新令牌与第三方软件、当时的授权用户绑定在一起的，因此这里需要判断该刷新令牌的归属合法性
```
String appStr = refreshTokenMap.get("refresh_token");
if(!appStr.startsWith(appId+"|"+"USERTEST")){
    //该refresh_token值不是颁发给该第三方软件的
}
```
需要注意，一个刷新令牌被使用以后，授权服务需要将其废弃，并重新颁发一个刷新令牌。

- 第二步，重新生成访问令牌。

生成访问令牌的处理流程，与颁发访问令牌环节的生成流程是一致的。授权服务会将新的访问令牌和新的刷新令牌，一起返回给第三方软件

### 在OAuth 2.0中，如何使用JWT结构化令牌？
#### JWT 结构化令牌
JWT 这种结构化体可以分为 HEADER（头部）、PAYLOAD（数据体）和 SIGNATURE（签名）三部分。经过签名之后的 JWT 的整体结构，是被句点符号分割的三段内容，结构为 header.payload.signature 。

- **HEADER** 表示装载令牌类型和算法等信息，是 JWT 的头部。其中，typ 表示第二部分 PAYLOAD 是 JWT 类型，alg 表示使用 HS256 对称签名的算法。

- **PAYLOAD** 表示是 JWT 的数据体，代表了一组数据。其中，sub（令牌的主体，一般设为资源拥有者的唯一标识）、exp（令牌的过期时间戳）、iat（令牌颁发的时间戳）是 JWT 规范性的声明，代表的是常规性操作。更多的通用声明，你可以参考RFC 7519 开放标准。不过，在一个 JWT 内可以包含一切合法的 JSON 格式的数据，也就是说，PAYLOAD 表示的一组数据允许我们自定义声明。

- **SIGNATURE** 表示对 JWT 信息的签名。那么，它有什么作用呢？我们可能认为，有了 HEADER 和 PAYLOAD 两部分内容后，就可以让令牌携带信息了，似乎就可以在网络中传输了，但是在网络中传输这样的信息体是不安全的，因为你在“裸奔”啊。所以，我们还需要对其进行加密签名处理，而 SIGNATURE 就是对信息的签名结果，当受保护资源接收到第三方软件的签名后需要验证令牌的签名是否合法。

#### 令牌内检
受保护资源来调用授权服务提供的检验令牌的服务，我们把这种校验令牌的方式称为令牌内检。
![img_5.png](img_5.png)

### JWT 是如何被使用的？
授权服务“扔出”一个令牌，受保护资源服务“接住”这个令牌，然后自己开始解析令牌本身所包含的信息就可以了，而不需要再去查询数据库或者请求 RPC 服务。这样也实现了我们上面说的令牌内检。
![img_6.png](img_6.png)

### 为什么要使用 JWT 令牌？
- 第一，JWT 的核心思想，就是用计算代替存储，有些 “时间换空间” 的 “味道”。当然，这种经过计算并结构化封装的方式，也减少了“共享数据库” 因远程调用而带来的网络传输消耗，所以也有可能是节省时间的。
- 第二，也是一个重要特性，是加密。因为 JWT 令牌内部已经包含了重要的信息，所以在整个传输过程中都必须被要求是密文传输的，这样被强制要求了加密也就保障了传输过程中的安全性。这里的加密算法，既可以是对称加密，也可以是非对称加密。
- 第三，使用 JWT 格式的令牌，有助于增强系统的可用性和可伸缩性。这一点要怎么理解呢？我们前面讲到了，这种 JWT 格式的令牌，通过“自编码”的方式包含了身份验证需要的信息，不再需要服务端进行额外的存储，所以每次的请求都是无状态会话。这就符合了我们尽可能遵循无状态架构设计的原则，也就是增强了系统的可用性和伸缩性。

#### JWT 令牌的缺点：没办法在使用过程中修改令牌状态
为了解决这个问题，我们可以把 JWT 令牌存储到远程的分布式内存数据库中吗？显然不能，因为这会违背 JWT 的初衷（将信息通过结构化的方式存入令牌本身）。因此，我们通常会有两种做法：
- 一是，将每次生成 JWT 令牌时的秘钥粒度缩小到用户级别，也就是一个用户一个秘钥。这样，当用户取消授权或者修改密码后，就可以让这个密钥一起修改。一般情况下，这种方案需要配套一个单独的密钥管理服务。
- 二是，在不提供用户主动取消授权的环境里面，如果只考虑到修改密码的情况，那么我们就可以把用户密码作为 JWT 的密钥。当然，这也是用户粒度级别的。这样一来，用户修改密码也就相当于修改了密钥。

### 令牌的生命周期
OAuth 2.0 的令牌生命周期，通常会有三种情况：
- 第一种情况是令牌的自然过期过程，这也是最常见的情况。这个过程是，从授权服务创建一个令牌开始，到第三方软件使用令牌，再到受保护资源服务验证令牌，最后再到令牌失效。同时，这个过程也不排除主动销毁令牌的事情发生，比如令牌被泄露，授权服务可以做主让令牌失效。

- 第二种情况是访问令牌失效之后可以使用刷新令牌请求新的访问令牌来代替失效的访问令牌，以提升用户使用第三方软件的体验。

- 第三种情况是让第三方软件比如小兔，主动发起令牌失效的请求，然后授权服务收到请求之后让令牌立即失效。我们来想一下，什么情况下会需要这种机制，也就是想一下第三方软件这样做的 “动机”，毕竟一般情况下 “我们很难放弃已经拥有的事物”。
![img_7.png](img_7.png)

### 如何安全、快速地接入OAuth 2.0？
#### 构建第三方软件应用
![img_8.png](img_8.png)
- 第一点，注册信息

软件需要先拥有自己的 app_id 和 app_serect 等信息，同时还要填写自己的回调地址 redirect_uri、申请权限等信息。

- 第二点，引导授权

当用户需要使用第三方软件，来操作其在受保护资源上的数据，就需要第三方软件来引导授权

- 第三点，使用访问令牌。

官方规范给出的使用访问令牌请求的方式，有三种
1. Form-Encoded Body Parameter（表单参数）
```
POST /resource HTTP/1.1
Host: server.example.com
Content-Type: application/x-www-form-urlencoded

access_token=b1a64d5c-5e0c-4a70-9711-7af6568a61fb
```
2. URI Query Parameter（URI 查询参数）
```
GET /resource?access_token=b1a64d5c-5e0c-4a70-9711-7af6568a61fb HTTP/1.1
Host: server.example.com
```
3. Authorization Request Header Field（授权请求头部字段）
```
GET /resource HTTP/1.1
Host: server.example.com
Authorization: Bearer b1a64d5c-5e0c-4a70-9711-7af6568a61fb
```

建议你采用表单提交，也就是 POST 的方式来提交令牌，因为表单提交的方式在保证安全传输的同时，还不需要去额外处理 Authorization 头部信息。

- 第四点，使用刷新令牌。

应该将 expires_in 值保存下来并定时检测；如果发现 expires_in 即将过期，则需要利用 refresh_token 去重新请求授权服务，以便获取新的、有效的访问令牌。

刷新令牌是一次性的，使用之后就会失效，但是它的有效期会比访问令牌要长。这个时候我们可能会想到，如果刷新令牌也过期了怎么办？在这种情况下，我们就需要将刷新令牌和访问令牌都放弃，相当于回到了系统的初始状态，只能让用户小明重新授权了。

#### 服务市场中的第三方应用软件
##### 构建受保护资源服务
![img_9.png](img_9.png)
1. 不同的权限对应不同的操作

2. 不同的权限对应不同的数据

3. 不同的用户对应不同的数据

### 除了授权码许可类型，OAuth 2.0还支持什么授权流程？

#### 资源拥有者凭据许可
![img_10.png](img_10.png)
- 步骤 1：当用户访问第三方软件小兔时，会提示输入用户名和密码。索要用户名和密码，就是资源拥有者凭据许可类型的特点。
- 步骤 2：这里的 grant_type 的值为 password，告诉授权服务使用资源拥有者凭据许可凭据的方式去请求访问。
- 步骤 3：授权服务在验证用户名和密码之后，生成 access_token 的值并返回给第三方软件。

#### 客户端凭据许可
![img_11.png](img_11.png)
- 步骤 1：第三方软件小兔通过后端服务向授权服务发送请求，这里 grant_type 的值为 client_credentials，告诉授权服务要使用第三方软件凭据的方式去请求访问。

- 步骤 2：在验证 app_id 和 app_secret 的合法性之后，生成 access_token 的值并返回。

#### 隐式许可
![img_12.png](img_12.png)

- 步骤 1：用户通过浏览器访问第三方软件小兔。此时，第三方软件小兔实际上是嵌入浏览器中执行的应用程序。
- 步骤 2：这个流程和授权码流程类似，只是需要特别注意一点，response_type 的值变成了 token，是要告诉授权服务直接返回 access_token 的值。

#### 如何选择？
在对接 OAuth 2.0 的时候先考虑授权码许可类型，其次再结合现实生产环境来选择：
- 如果小兔软件是官方出品，那么可以直接使用资源拥有者凭据许可；
- 如果小兔软件就是只嵌入到浏览器端的应用且没有服务端，那就只能选择隐式许可；
- 如果小兔软件获取的信息不属于任何一个第三方用户，那可以直接使用客户端凭据许可类型。
![img_13.png](img_13.png)

1. 所有的授权许可类型中，授权码许可类型的安全性是最高的。因此，只要具备使用授权码许可类型的条件，我们一定要首先授权码许可类型。

2. 所有的授权许可类型都是为了解决现实中的实际问题，因此我们还要结合实际的生产环境，在保障安全性的前提下选择最合适的授权许可类型，比如使用客户端凭据许可类型的小兔软件就是一个案例。

### 实践OAuth 2.0时，使用不当可能会导致哪些安全漏洞？

#### CSRF 攻击
![img_14.png](img_14.png)
那如何避免这种攻击呢？方法也很简单，实际上 OAuth 2.0 中也有这样的建议，就是使用 state 参数，它是一个随机值的参数。

#### XSS 攻击
![img_15.png](img_15.png)
受保护资源服务就需要对这类 XSS 漏洞做修复，而具体的修复方法跟其它网站防御 XSS 类似，最简单的方法就是对此类非法信息做转义过滤，比如对包含

#### 水平越权
![img_16.png](img_16.png)
![img_17.png](img_17.png)
在开放平台环境下，由于开放网关和数据接口提供方来自不同的业务部门，防止水平校验的逻辑处理很容易被遗漏：
- 一方面，开放网关的作用是将用户授权之后的访问令牌 access_token 信息转换成真实的用户信息，比如上面提到的商家 ID，然后传递到接口提供方，数据归属判断逻辑只能在接口提供方内部处理；
- 另一方面，数据提供方往往会认为开放出的接口是被“跟自己一个公司的系统所调用的”，容易忽略水平校验的逻辑处理。

#### 授权码失窃
![img_18.png](img_18.png)
这时问题的根源就在于两点：
- 授权服务在进行授权码校验的时候，没有校验 app_id_B；
- 软件 B（也就是极客时间）使用过一次 codeB 的值之后，授权服务没有删除这个 codeB；

#### 重定向 URI 被篡改
![img_19.png](img_19.png)
授权服务提供对第三方软件的回调URL做完整性要求和完整性校验，并增加签名校验

#### 总结
![img_20.png](img_20.png)

###  实战：利用OAuth 2.0实现一个OpenID Connect用户身份认证协议

#### OIDC 是什么？
OIDC= 授权协议 + 身份认证，是 OAuth 2.0 的超集

#### OIDC 和 OAuth 2.0 的角色对应关系
OIDC的三个主要角色：
- EU（End User），代表最终用户。
- RP（Relying Party），代表认证服务的依赖方，就是上面我提到的第三方软件。
- OP（OpenID Provider），代表提供身份认证服务方。
![img_21.png](img_21.png)

#### OIDC 和 OAuth 2.0 的关键区别
![img_22.png](img_22.png)
一个基于授权码流程的 OIDC 协议流程，跟 OAuth 2.0 中的授权码许可的流程几乎完全一致，唯一的区别就是多返回了一个 ID_TOKEN，我们称之为 ID 令牌

#### OIDC 中的 ID 令牌生成和解析方法
##### ID 令牌的内容是什么呢？
ID 令牌是一个 JWT 格式的令牌

5 个 JWT 声明参数也是必须要有的:
- iss，令牌的颁发者，其值就是身份认证服务（OP）的 URL。
- sub，令牌的主题，其值是一个能够代表最终用户（EU）的全局唯一标识符。
- aud，令牌的目标受众，其值是三方软件（RP）的 app_id。
- exp，令牌的到期时间戳，所有的 ID 令牌都会有一个过期时间。
- iat，颁发令牌的时间戳。

##### 用访问令牌获取 ID 令牌之外的信息
OIDC 的流程就是：生成 ID 令牌 -> 创建 UserInfo 端点 -> 解析 ID 令牌 -> 记录登录状态 -> 获取 UserInfo。

#### 单点登录
![img_23.png](img_23.png)


## 进阶篇
### 实现 [go-oauth2/oauth2](https://github.com/go-oauth2/oauth2)
