package http

import (
  "net/http"
)

const (
	MethodGet     = "GET"
	MethodHead    = "HEAD"
	MethodPost    = "POST"
	MethodPut     = "PUT"
	MethodPatch   = "PATCH" // RFC 5789
	MethodDelete  = "DELETE"
	MethodConnect = "CONNECT"
	MethodOptions = "OPTIONS"
	MethodTrace   = "TRACE"
)

func MakeHttpRequest(url string, method string) *http.Response{
  req, err := http.NewRequest(method, url, nil)

  if err != nil {
    log.Fatalln(err)
  }

  return resp
}
