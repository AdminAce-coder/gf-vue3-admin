package menu

import (
	"os"
	"runtime"
	"testing"
)

func TestMunus(t *testing.T) {

	//ExecCmd(gctx.New())
	goRootForFilter := runtime.GOROOT()
	goRoot := os.Getenv("GOROOT")
	t.Logf("路径是%s====%s", goRootForFilter, goRoot)
}
