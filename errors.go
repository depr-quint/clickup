package clickup

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type ErrorResponse struct {
	Response *http.Response
	Err      string `json:"err"`
	Code     string `json:"ECODE"`
}

func (r *ErrorResponse) Error() string {
	return fmt.Sprintf("%v %v: %d %v %+v",
		r.Response.Request.Method, r.Response.Request.URL,
		r.Response.StatusCode, r.Err, r.Code)
}

func CheckResponse(resp *http.Response) error {
	if c := resp.StatusCode; 200 <= c && c <= 299 {
		return nil
	}

	errorResponse := &ErrorResponse{Response: resp}
	data, err := ioutil.ReadAll(resp.Body)
	if err == nil && data != nil {
		_ = json.Unmarshal(data, errorResponse)
	}

	// repopulate body
	resp.Body = ioutil.NopCloser(bytes.NewBuffer(data))

	switch {
	default:
		return errorResponse
	}
}
