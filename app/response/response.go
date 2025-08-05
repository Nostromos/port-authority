package response

import "net/http"

type ResponseOptions struct {
	StatusCode int
	Headers http.Header
	Body []byte
}


func BuildResponse(opts ResponseOptions) *http.Response {
	status := opts.StatusCode
	if status == 0 {
		status = 200
	}

	resp := &http.Response{
		StatusCode: status,
		Header: make(http.Header),
		Body: opts.Body,
	}
}

func WithStatus() {

}