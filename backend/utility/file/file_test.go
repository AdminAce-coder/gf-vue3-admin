package file

import "testing"

func TestDeleteDir(t *testing.T) {
	err := DeleteDir("D:\\code\\GF-VUE3-ADMIN-V2\\backend\\internal\\controller\\apitest")
	if err != nil {
		t.Error(err)
	}
}
