package dataprocess

import (
	"testing"

	"github.com/gogf/gf/v2/os/gctx"
)

var ctx = gctx.New()

func TestDataprocee(t *testing.T) {

	resulten := AesEcbEncrypt(ctx, "testtest")
	t.Logf("加密后的数据是,%s", resulten)
	resultout := AesEcbdecrypted(ctx, resulten)
	t.Logf("解密后的数据是,%s", resultout)
}
