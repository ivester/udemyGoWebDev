package main

import (
	"net/http"
)

// TODO: Use curl with cool options and stuff

func main() {
	http.ListenAndServe(":3000", nil)
}
