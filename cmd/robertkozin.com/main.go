package main

import (
	"net/http"
	"robertkozin.com/services/website"
)

func main() {
	site := website.New()
	http.ListenAndServe("0.0.0.0:8080", site)
}
