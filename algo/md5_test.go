package algo

import (
	"testing"
)

func TestMD5(t *testing.T) {
	str := "123456"
	res := MD5(str)
	t.Log(res)

	res = Md516bit(str)
	t.Log(res)
}
