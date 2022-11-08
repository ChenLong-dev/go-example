# 图片转字符画

## 代码实现功能描述

- 看到很多人喜欢在代码开始或者结束位置打印比较好看的图案，比如佛祖保佑无BUG；正好在学习go，于是就想实现一下这个功能。将图片内容或者文字内容转成字符串形式输出，方便我们可以将喜欢的图案转成字符串放到我们写的代码中。

## 实现思路简述

- 遍历获取图片像素的RGB信息，根据RGB的值去选择对应像素点的替换字符。
- 把文字转成字符画，就是先把文字内容写到图片中，然后回到遍历图片像素信息方式实现字符替换像素

## 代码整体实现

### 1、通过文件路径获取一个图片

```go
// GetImage 根据路径获取一个图片
func GetImage(imagePath string) *image.Image {
    // 读取image文件
    imageFile, _ := ioutil.ReadFile(imagePath)
    // bytes.buffer 是一个缓冲byte类型的缓冲器存放着都是byte
    imageBytes := bytes.NewBuffer(imageFile)
    // 解码
    im, _, err := image.Decode(imageBytes)
    if err != nil {
        log.Printf("解码图像失败%#v", err)
        return nil
    }
    return &im
}
```

### 2、根据图片的宽度，等比例缩放图片

> 如果有的图片太宽的话，打印的时候终端屏幕显示宽度太窄，导致显示换行问题

```go
// equalScaleImageFromWidth 根据宽度，等比例缩放图片, 返回新的像素值（等比例处理后的）
// maxWidth 图片最大宽度 如果超过此值就按照宽度 targetWidth 等比例缩放图片
func equalScaleImageFromWidth(img image.Image, maxWidth int, targetWidth int) *image.Image {
    // 原图像界限范围
    bounds := img.Bounds()
    dx := bounds.Dx()
    dy := bounds.Dy()
    log.Printf("原图宽=%d,高=%d", dx, dy)
    m := img
    if dx > maxWidth {
        // 按照宽度缩放
        m = resize.Resize(uint(targetWidth), 0, img, resize.Lanczos3)
    }
    return &m
}
```

### 3、将图片保存到文件中

```go
// SaveNewImage 保存一个新的图片文件
func SaveNewImage(targetPath string, newImg *image.Image) {
    f, err := os.Create(targetPath)
    if err != nil {
        log.Printf("创建文件失败%#v", err)
    }
    defer f.Close()
    // 获取文件后缀判断编码方法
    suffix := path.Ext(targetPath) // strings.Replace(path.Ext(targetPath), ".", "", -1)
    if suffix == ".jpg" || suffix == ".jpge" {
        err = jpeg.Encode(f, *newImg, nil)
    } else if suffix == "png" {
        err = png.Encode(f, *newImg)
    } else if suffix == "gif" {
        err = gif.Encode(f, *newImg, nil)
    } else {
        log.Printf("非可处理图片格式")
    }
    if err != nil {
        log.Printf("保存一个新的图片失败%#v", err)
    }
    log.Printf("保存一个新的图片成功！")
}
```

### 4、实现图片转字符串

> 这里我简单根据像素RGB相加的值判断颜色选择字符（越大越接近白色，选择空格替换）

```go
// ImageToChars 图像转成字符展示
// RGBA(R,G,B,A) 像素值
// R：红色值 G：绿色值 B：蓝色值 A：Alpha透明度
func ImageToChars(img *image.Image, inputChars string) string {
    sourceImage := *img
    bounds := sourceImage.Bounds()
    dx := bounds.Dx()
    dy := bounds.Dy()
    imgChars := "**  "
    if inputChars != "" {
        imgChars = inputChars
    }
    resultString := ""
    intSliceRGB := []int{}
    maxIntRGB := 0
    minIntRGB := 255 * 3
    for i := 0; i < dy; i++ {
        for j := 0; j < dx; j++ {
            colorRgb := sourceImage.At(j, i)
            // 获取 uint32 像素值 >> 8 转换为 255 值
            r, g, b, _ := colorRgb.RGBA()
            // sumRGB 越大越趋近于白色，越小越趋近于黑色
            sumRGB := int(uint8(r>>8)) + int(uint8(g>>8)) + int(uint8(b>>8))
            // 找到最大值和最小值，方便将像素值划分不同区间段
            if maxIntRGB < sumRGB {
                maxIntRGB = sumRGB
            }
            if minIntRGB > sumRGB {
                minIntRGB = sumRGB
            }
            intSliceRGB = append(intSliceRGB, sumRGB)
        }
    }
    for index, val := range intSliceRGB {
        // partLen为区间跨度 我们按照传入字符串元素个数平均分配像素值，将像素值分成几个区间
        //  +1 防止下标溢出
        // 例如 像素值 0~500 用两个字符替换 0~250 251~500 两个区间分别与其对应即可
        partLen := (maxIntRGB-minIntRGB)/(len(imgChars)) + 1
        // 根据像素值取不同的字符替换像素
        str := string(imgChars[(val-minIntRGB)/partLen])
        resultString += str
        // 判断换行
        if (index+1)%dx == 0 {
            resultString += "\n"
        }
    }
    return resultString
}
```

### 5、把字符串写入文件

> 我们也可以把程序打包二进制文件，在我们需要的地方调用，然后打印出我们喜欢的内容即可

```go
// WriteStringToFile 通过 io.WriteString 写入文件
func WriteStringToFile(fileName string, writeInfo string) {
    f, err := os.Create(fileName)
    if err != nil {
        log.Printf("创建文件失败:%s，日志%#v", fileName, err)
        return
    }
    log.Printf("成功创建文件:%s", fileName)
    defer f.Close()
    // 将文件写进去
    if _, err = io.WriteString(f, writeInfo); err != nil {
        log.Printf("WriteStringToFile 写入文件失败:%+v", err)
        return
    }
    log.Printf("WriteStringToFile 写入文件成功")
}
```

### 6、根据文字内容创建图片

> 这里需要我们下载ttf的字体样式，然后根据字体样式把文字渲染到图片上，用到了`github.com/golang/freetype`

```go
// CrateImageFromString 根据字符串创建图片
func CrateImageFromString(fileName string, imgString string, width int, height int, fontSize float64) *image.Image {
    // 读字体数据
    fontBytes, err := ioutil.ReadFile("ht.ttf")
    if err != nil {
        log.Printf("CrateImageFromString 读字体信息失败:%#v", err)
        return nil
    }
    font, err := freetype.ParseFont(fontBytes)
    if err != nil {
        log.Printf("CrateImageFromString 解析字体失败:%#v", err)
        return nil
    }

    // 设置图片范围和底色
    img := image.NewRGBA(image.Rect(0, 0, width, height))
    draw.Draw(img, img.Bounds(), image.White, image.ZP, draw.Src)

    c := freetype.NewContext()
    c.SetDPI(100) // 设置像素分辨率 每英寸点数 越大像素越好
    c.SetFont(font)
    c.SetFontSize(fontSize)
    c.SetClip(img.Bounds())
    c.SetDst(img)
    c.SetSrc(image.Black)

    // 设置字体显示位置
    // >>6 把Int26_6转回int
    pt := freetype.Pt(10, int(c.PointToFixed(fontSize)>>6))
    for _, s := range strings.Split(imgString, "\n") {
        _, err = c.DrawString(s, pt)
        if err != nil {
            log.Printf("CrateImageFromString 字符串画图失败:%#v", err)
            return nil
        }
        pt.Y += c.PointToFixed(fontSize)
    }

    // 创建图片文件
    imgFile, err := os.Create(fileName)
    if err != nil {
        log.Printf("创建文件失败:%s，日志%#v", fileName, err)
        return nil
    }
    defer imgFile.Close()

    // 保存图像到文件
    err = png.Encode(imgFile, img)
    if err != nil {
        log.Fatal(err)
    }
    return GetImage(fileName)
}
```
