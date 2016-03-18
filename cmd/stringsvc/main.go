package main

import (
	"net/http"
	"strings"

	trans "github.com/wavecell/stringsvc/transport/http"
)

type StringService struct{}

func (StringService) Uppercase(str string) (string, error) {
	return strings.ToUpper(str), nil
}

func (StringService) Count(str string) int {
	return len(str)
}

func main() {
	var svc StringService

	trans.HTTPServersForEndpoints(svc)
	http.ListenAndServe(":12345", nil)
}
