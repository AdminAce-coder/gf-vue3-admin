package apictrl

import (
	"testing"

	"github.com/gogf/gf/v2/os/gctx"
)

func TestDeleApi(t *testing.T) {
	//testapi := DeleteApi{
	//	ApiVersion: "v1",
	//	ApiGroup:   "apitest",
	//	ApiName:    "methodb",
	//}
	//
	//err := testapi.DeleteApi(gctx.New())
	//if err != nil {
	//	t.Error(err)
	//}
	//dl := DeleteApiGroup{
	//	ApiGroupName: "apitest",
	//}
	//err := dl.DeleteGroup(gctx.New())
	//if err != nil {
	//	t.Error(err)
	//}

}
func TestDelCtrlFile(t *testing.T) {
	DelCtrlFile(gctx.New(), "/api/v1/apidemov1", "apitest")
}

func TestFoudApiName(t *testing.T) {
	FoudApiName(gctx.New(), "D:\\code\\GF-VUE3-ADMIN-V2\\backend\\api\\apitest\\v1\\apitest.go", "apidemo")

}
