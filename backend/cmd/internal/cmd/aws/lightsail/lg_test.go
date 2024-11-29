package lightsail

import (
	"gf-vue3-admin/cmd/internal/CplatformClinet"
	"github.com/aws/aws-sdk-go-v2/service/lightsail"
	"github.com/gogf/gf/v2/os/gctx"
	"testing"
)

func TestListIstance(t *testing.T) {

	cl := CplatformClinet.GetClient[*lightsail.Client](CplatformClinet.WithRegion("ap-northeast-1"), CplatformClinet.WithClientType("lightsail"))
	if cl == nil {
		t.Error("cl is nil")
	}
	err := ListIstance(gctx.New(), cl)
	if err != nil {
		t.Error(err)
	}

}
