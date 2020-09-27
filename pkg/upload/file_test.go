package upload

import "testing"

func TestGetFileName(t *testing.T) {
	println(GetFileExt("zhangsan.jpg"))
}
