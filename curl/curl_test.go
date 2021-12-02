package curl

import (
	"fmt"
	"net/url"
	"testing"
)

func TestGet(t *testing.T) {
	t.Log("Test api")

	resp, err := Get("https://caizyang.kmdns.net:8001/api/v1/auth/get-access-token", nil)
	if err != nil {
		t.Error(err.Error())
		return
	}
	fmt.Println(resp)

	p := url.Values{}
	p.Set("access_token", "123456")
	Timeout = 2
	resp, err = Post("https://caizyang.kmdns.net:8001/api/v1/user/create", p, nil)
	fmt.Println(resp)
}
