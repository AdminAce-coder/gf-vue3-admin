package EC2

import (
	"gf-vue3-admin/cmd/internal/CplatformClinet"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/gogf/gf/v2/os/gctx"
	"testing"
)

func TestCreatrEc2(t *testing.T) {
	cl := CplatformClinet.GetClient[*ec2.Client](CplatformClinet.WithRegion("us-east-1"), CplatformClinet.WithClientType("ec2"))
	if cl == nil {
		t.Error("cl is nil")
	}
	err := CreatrEc2(gctx.New(), cl)
	if err != nil {
		t.Error(err)
	}

}
