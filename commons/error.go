package commons

import "fmt"

type RequestError struct {
	StatusCode      int
	Err             error
	ErrorOccurredIn string
}

func (r *RequestError) Error() string {
	return fmt.Sprintf("status %d: err %v", r.StatusCode, r.Err)
}
