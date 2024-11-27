package register

import "testing"

func TestRegister(t *testing.T) {
	//c1 := &CreateApiGroupReq{
	//	ApiPath: "/user/tessdst",
	//	Register: RouteConfig{
	//		NeedAuth:  true,
	//		GroupName: "tessdsdt",
	//		Enable:    true,
	//	},
	//}
	err := DeleteRouteConfig("/api/v1/apitest")
	if err != nil {
		t.Error(err)
	}
}
