package s3

import (
	"gf-vue3-admin/cmd/internal/CplatformClinet"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/gogf/gf/v2/os/gctx"
	"testing"
)

func TestS3(t *testing.T) {
	cl := CplatformClinet.GetClient[*s3.Client](CplatformClinet.WithRegion("us-east-1"), CplatformClinet.WithClientType("s3"))
	err := ListS3Opject(gctx.New(), cl)
	if err != nil {
		t.Error(err)
	}
}
