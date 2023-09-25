package main

import (
	"context"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	abstractions "github.com/microsoft/kiota-abstractions-go"
	absser "github.com/microsoft/kiota-abstractions-go/serialization"
	jsonserialization "github.com/microsoft/kiota-serialization-json-go"
	graph "github.com/microsoftgraph/msgraph-sdk-go"
	graphgroups "github.com/microsoftgraph/msgraph-sdk-go/groups"
	"github.com/microsoftgraph/msgraph-sdk-go/models"
)

func PrintStr(prefix interface{}, v absser.Parsable) {
	vbyte, err := jsonserialization.Marshal(v)
	if err != nil {
		fmt.Printf("[PrintJson] err:%v\n", err)
		return
	}
	fmt.Printf("[%s] %s\n", prefix, string(vbyte))
}

func PrintStr2(prefix interface{}, vv []models.DirectoryObjectable) {
	for i, v := range vv {
		PrintStr(fmt.Sprintf("%s-%d", prefix, i), v)
	}
}

/* AAD ******************************************************************************************/

type AAD struct {
	graphClient *graph.GraphServiceClient
}

func NewClient(tenantID string, clientID string, clientSecret string) (aad *AAD, err error) {
	aad = &AAD{}
	cred, err := azidentity.NewClientSecretCredential(tenantID, clientID, clientSecret, nil)
	if err != nil {
		fmt.Printf("[GetClient] err:%v\n", err)
		return
	}
	aad.graphClient, err = graph.NewGraphServiceClientWithCredentials(
		cred, []string{"https://graph.microsoft.com/.default"})
	if err != nil {
		fmt.Printf("[GetClient] err:%v\n", err)
		return
	}
	return
}

/* Users *****************************************************************************************/

// UsersGet 列出所有用户: Users().Get()
func (aad *AAD) UsersGet() {
	result, err := aad.graphClient.Users().Get(context.Background(), nil)
	if err != nil {
		fmt.Printf("[UsersGet] err:%v\n", err)
		return
	}
	res := result.GetValue()
	for i, v := range res {
		PrintStr(fmt.Sprintf("UsersGet-%d", i), v)
		fmt.Printf("- v.GetDisplayName():%s\n", *(v.GetDisplayName()))
		fmt.Printf("- v.GetId():%s\n", *(v.GetId()))
	}
}

// UsersByUserIdGet 通过ID获取用户信息：Users().ByUserId().Get()
func (aad *AAD) UsersByUserIdGet(ID string) {
	result, err := aad.graphClient.Users().ByUserId(ID).Get(context.Background(), nil)
	if err != nil {
		fmt.Printf("[UsersByUserIdGet] err:%v\n", err)
		return
	}
	PrintStr("UsersByUserIdGet", result)
	fmt.Printf("- result.GetDisplayName():%s\n", *(result.GetDisplayName()))
	fmt.Printf("- result.GetId():%s\n", *(result.GetId()))
	resmo := result.GetMemberOf()
	PrintStr2("resmo", resmo)

	restmo := result.GetTransitiveMemberOf()
	PrintStr2("restmo", restmo)
}

// UsersByUserIdMemberOfGet 通过ID获取用户是其直接成员的组、目录角色和管理单元：Users().ByUserId().MemberOf().Get()
func (aad *AAD) UsersByUserIdMemberOfGet(ID string) {
	result, err := aad.graphClient.Users().ByUserId(ID).Get(context.Background(), nil)
	if err != nil {
		fmt.Printf("[UsersByUserIdMemberOfGet] err:%v\n", err)
		return
	}
	PrintStr("UsersByUserIdMemberOfGet", result)
	fmt.Printf("- result.GetDisplayName():%s\n", *(result.GetDisplayName()))
	fmt.Printf("- result.GetId():%s\n", *(result.GetId()))
}

// UsersByUserIdTransitiveMemberOfGet 通过ID获取用户所属的组、目录角色和管理单元：Users().ByUserId().TransitiveMemberOfGet().Get()
func (aad *AAD) UsersByUserIdTransitiveMemberOfGet(ID string) {
	result, err := aad.graphClient.Users().ByUserId(ID).Get(context.Background(), nil)
	if err != nil {
		fmt.Printf("[UsersByUserIdTransitiveMemberOfGet] err:%v\n", err)
		return
	}
	PrintStr("UsersByUserIdTransitiveMemberOfGet", result)
	fmt.Printf("- result.GetDisplayName():%s\n", *(result.GetDisplayName()))
	fmt.Printf("- result.GetId():%s\n", *(result.GetId()))
}

/* Groups *****************************************************************************************/

// GroupsGet 获取组列表：Groups().Get()
func (aad *AAD) GroupsGet() {
	result, err := aad.graphClient.Groups().Get(context.Background(), nil)
	if err != nil {
		fmt.Printf("[GroupsGet] err:%v\n", err)
		return
	}
	res := result.GetValue()
	for i, v := range res {
		PrintStr(fmt.Sprintf("GroupsGet-%d", i), v)
		fmt.Printf("- v.GetDisplayName():%s\n", *(v.GetDisplayName()))
		fmt.Printf("- v.GetId():%s\n", *(v.GetId()))
		PrintStr2("- v.GetMemberOf()", v.GetMemberOf())
		PrintStr2("- v.GetMembers()", v.GetMembers())
		PrintStr2("- v.GetTransitiveMemberOf()", v.GetTransitiveMemberOf())
		PrintStr2("- v.GetTransitiveMembers()", v.GetTransitiveMembers())
	}
}

// GroupsGet2 获取组的筛选列表：Groups().Get()
func (aad *AAD) GroupsGet2() {
	headers := abstractions.NewRequestHeaders()
	headers.Add("ConsistencyLevel", "eventual")

	requestCount := true
	requestFilter := "hasMembersWithLicenseErrors eq true"

	requestParameters := &graphgroups.GroupsRequestBuilderGetQueryParameters{
		Count:  &requestCount,
		Filter: &requestFilter,
		Select: []string{"id", "displayName"},
	}
	configuration := &graphgroups.GroupsRequestBuilderGetRequestConfiguration{
		Headers:         headers,
		QueryParameters: requestParameters,
	}
	result, err := aad.graphClient.Groups().Get(context.Background(), configuration)
	if err != nil {
		fmt.Printf("[GroupsGet2] err:%v\n", err)
		return
	}
	res := result.GetValue()
	for i, v := range res {
		PrintStr(fmt.Sprintf("GroupsGet2-%d", i), v)
		fmt.Printf("- v.GetDisplayName():%s\n", *(v.GetDisplayName()))
		fmt.Printf("- v.GetId():%s\n", *(v.GetId()))
		PrintStr2("- v.GetMemberOf()", v.GetMemberOf())
		PrintStr2("- v.GetMembers()", v.GetMembers())
		PrintStr2("- v.GetTransitiveMemberOf()", v.GetTransitiveMemberOf())
		PrintStr2("- v.GetTransitiveMembers()", v.GetTransitiveMembers())
	}
}

// GroupsByGroupIdMemberOfGet 获取组是其直接成员的组和管理单元:Groups().ByGroupId().MemberOf().Get()
func (aad *AAD) GroupsByGroupIdMemberOfGet(ID string) {
	result, err := aad.graphClient.Groups().ByGroupId(ID).MemberOf().Get(context.Background(), nil)
	if err != nil {
		fmt.Printf("[GroupsByGroupIdMemberOfGet] err:%v\n", err)
		return
	}
	res := result.GetValue()
	for i, v := range res {
		PrintStr(fmt.Sprintf("GroupsByGroupIdMemberOfGet-%s-%d", ID, i), v)
		fmt.Printf("- v.GetId():%s\n", *(v.GetId()))
		fmt.Printf("- v.GetOdataType():%s\n", *(v.GetOdataType()))
		value, _ := v.GetBackingStore().Get("displayName")
		fmt.Printf("- v.displayName:%s\n", *value.(*string))
	}
}

// GroupsByGroupIdMembersGet 获取组是其直接成员的组和管理单元:Groups().ByGroupId().Members().Get()
func (aad *AAD) GroupsByGroupIdMembersGet(ID string) {
	result, err := aad.graphClient.Groups().ByGroupId(ID).Members().Get(context.Background(), nil)
	if err != nil {
		fmt.Printf("[GroupsByGroupIdMembersGet] err:%v\n", err)
		return
	}
	res := result.GetValue()
	for i, v := range res {
		PrintStr(fmt.Sprintf("GroupsByGroupIdMembersGet-%s-%d", ID, i), v)
		fmt.Printf("- v.GetId():%s\n", *(v.GetId()))
		fmt.Printf("- v.GetOdataType():%s\n", *(v.GetOdataType()))
		value, _ := v.GetBackingStore().Get("displayName")
		fmt.Printf("- v.displayName:%s\n", *value.(*string))
	}
}

// GroupsByGroupIdTransitiveMemberOfGet 获取组是其直接成员的组和管理单元:Groups().ByGroupId().TransitiveMemberOf().Get()
func (aad *AAD) GroupsByGroupIdTransitiveMemberOfGet(ID string) {
	result, err := aad.graphClient.Groups().ByGroupId(ID).TransitiveMemberOf().Get(context.Background(), nil)
	if err != nil {
		fmt.Printf("[GroupsByGroupIdTransitiveMemberOfGet] err:%v\n", err)
		return
	}
	res := result.GetValue()
	for i, v := range res {
		PrintStr(fmt.Sprintf("GroupsByGroupIdTransitiveMemberOfGet-%s-%d", ID, i), v)
		fmt.Printf("- v.GetId():%s\n", *(v.GetId()))
		fmt.Printf("- v.GetOdataType():%s\n", *(v.GetOdataType()))
		value, _ := v.GetBackingStore().Get("displayName")
		fmt.Printf("- v.displayName:%s\n", *value.(*string))
	}
}

// GroupsByGroupIdTransitiveMembersGet 获取组的可传递成员身份:Groups().ByGroupId().TransitiveMembers().Get()
func (aad *AAD) GroupsByGroupIdTransitiveMembersGet(ID string) {
	result, err := aad.graphClient.Groups().ByGroupId(ID).TransitiveMembers().Get(context.Background(), nil)
	if err != nil {
		fmt.Printf("[GroupsByGroupIdTransitiveMembersGet] err:%v\n", err)
		return
	}
	res := result.GetValue()
	for i, v := range res {
		PrintStr(fmt.Sprintf("GroupsByGroupIdTransitiveMembersGet-%s-%d", ID, i), v)
		fmt.Printf("- v.GetId():%s\n", *(v.GetId()))
		fmt.Printf("- v.GetOdataType():%s\n", *(v.GetOdataType()))
		value, _ := v.GetBackingStore().Get("displayName")
		fmt.Printf("- v.displayName:%s\n", *value.(*string))
	}
}

// GroupsByGroupIdOwnersGet 列出组所有者:Groups().ByGroupId().Owners().Get()
func (aad *AAD) GroupsByGroupIdOwnersGet(ID string) {
	result, err := aad.graphClient.Groups().ByGroupId(ID).Owners().Get(context.Background(), nil)
	if err != nil {
		fmt.Printf("[GroupsByGroupIdOwnersGet] err:%v\n", err)
		return
	}
	res := result.GetValue()
	for i, v := range res {
		PrintStr(fmt.Sprintf("GroupsByGroupIdOwnersGet-%s-%d", ID, i), v)
		fmt.Printf("- v.GetId():%s\n", *(v.GetId()))
		fmt.Printf("- v.GetOdataType():%s\n", *(v.GetOdataType()))
		value, _ := v.GetBackingStore().Get("displayName")
		fmt.Printf("- v.displayName:%s\n", *value.(*string))
	}
}

/* DirectoryObjects *****************************************************************************************/

// DirectoryObjectsGet 获取 directoryObject：DirectoryObjects().Get()
func (aad *AAD) DirectoryObjectsGet() {
	result, err := aad.graphClient.DirectoryObjects().Get(context.Background(), nil)
	if err != nil {
		fmt.Printf("[DirectoryObjectsGet] err:%v\n", err)
		return
	}
	res := result.GetValue()
	for i, v := range res {
		PrintStr(fmt.Sprintf("DirectoryObjectsGet-%d", i), v)
		fmt.Printf("- v.GetId():%s\n", *(v.GetId()))
	}
}

/* main *****************************************************************************************/

func main() {
	// 丁新的应用
	tenantID := "471239f6-f844-478a-bb21-41baeb227175"
	clientID := "e54ac9e1-cd6d-48d9-be52-e3b742d58adc"
	clientSecret := "Kyp8Q~dxGnLNXi~aMCMgPDJAcZMXftMg_1DMCaf5"
	//UserID := "325edea9-0da1-4626-8155-7745e7020673"

	// 宝哥的20230725test应用
	//tenantID := "62e6f7d0-8971-46ac-ab5e-6fcb4f7cb156"
	//clientID := "992dde7a-d4fe-4a1d-8573-fdb0203c06f0"
	//clientSecret := "Add8Q~rPPDAW9ycZcbmVfRjvld5mZGblNprEVbnd"
	//ID := "f46c9fd9-f727-4bf7-8365-880418f0b5e0"

	aad, err := NewClient(tenantID, clientID, clientSecret)
	if err != nil {
		fmt.Printf("err:%v\n", err)
		return
	}

	fmt.Printf("=== new client is successful! ===\n\n")
	//aad.UsersGet()
	fmt.Printf("---------- [UsersGet] ----------\n\n")

	//aad.UsersByUserIdGet(UserID)
	fmt.Printf("---------- [UsersByUserIdGet] ----------\n\n")

	//aad.UsersByUserIdMemberOfGet(UserID)
	fmt.Printf("---------- [UsersByUserIdMemberOfGet] ----------\n\n")

	//aad.UsersByUserIdTransitiveMemberOfGet(UserID)
	fmt.Printf("---------- [UsersByUserIdTransitiveMemberOfGet] ----------\n\n")

	/********************************************************************************************/

	aad.GroupsGet()
	fmt.Printf("---------- [GroupsGet] ----------\n\n")

	//aad.GroupsGet2()
	//fmt.Printf("---------- [GroupsGet2] ----------\n\n")

	//GroupID := "7ff7a529-8fbc-4068-bc9b-38af5a53ae97"
	//aad.GroupsByGroupIdMemberOfGet(GroupID)
	fmt.Printf("---------- [GroupsByGroupIdMemberOfGet] ----------\n\n")

	//aad.GroupsByGroupIdMembersGet(GroupID)
	fmt.Printf("---------- [GroupsByGroupIdMembersGet] ----------\n\n")

	//aad.GroupsByGroupIdTransitiveMemberOfGet(GroupID)
	fmt.Printf("---------- [GroupsByGroupIdTransitiveMemberOfGet] ----------\n\n")

	//aad.GroupsByGroupIdTransitiveMembersGet(GroupID)
	fmt.Printf("---------- [GroupsByGroupIdTransitiveMembersGet] ----------\n\n")

	//aad.GroupsByGroupIdOwnersGet(GroupID)
	fmt.Printf("---------- [GroupsByGroupIdOwnersGet] ----------\n\n")

	aad.DirectoryObjectsGet()
	//fmt.Printf("---------- [DirectoryObjectsGet] ----------\n\n")

}
