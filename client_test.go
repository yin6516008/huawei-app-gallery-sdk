package huawei

import (
	"net/http"
	"testing"
)

func TestNewRequest(t *testing.T) {
	client, err := NewHuaweiClientWithEnv()
	if err != nil {
		t.Log(err)
	}
	req, err := client.NewRequest(http.MethodGet, "/api/token", nil, nil)
	if err != nil {
		t.Error(err)
	}
	t.Log(req.URL)
}
