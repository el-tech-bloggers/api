package api

import (
	"fmt"
	"net/http"
)

func RootHandler(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprintln(w, "Hello World")
}
