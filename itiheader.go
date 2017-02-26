package itiheader

import (
	"net/http"
)

// SchemesMatcher Store the schema to match with the request Path
type SchemesMatcher struct {
	header string
	value  string
}

// New is the constructor to ItiHost
func New(header, value string) *SchemesMatcher {
	return &SchemesMatcher{
		header: header,
		value:  value,
	}
}

// Match returns if the request can be handled by this Route.
func (h *SchemesMatcher) Match(req *http.Request) bool {
	if req.Header.Get(h.header) != h.value {
		return false
	}
	return true
}
