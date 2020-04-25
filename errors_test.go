package clickup

import (
	"net/http"
	"testing"
)

func TestCheckResponse(t *testing.T) {
	errClient := NewClient(nil, str("invalid_token"))
	req, _ := errClient.NewRequest(http.MethodGet, "teams", nil)
	_, err := errClient.Do(req, nil)
	if err == nil {
		t.Error("error expected")
	}
}

func str(s string) *string {
	return &s
}
