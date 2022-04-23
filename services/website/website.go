package website

import (
	"fmt"
	"net/http"
)

type Website struct {

}

func New() *Website {
	return &Website{}
}

func (w *Website) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(res, "Hello, World")
}
