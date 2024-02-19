package main

import (
	"net/http"
	"github.com/iainquayle/armfinity/components"
	"github.com/a-h/templ"
	"fmt"
)

func main() {
	http.Handle("/", templ.Handler(components.Home()))
	fmt.Println("Server started on port 8080")
	http.ListenAndServe(":8080", nil)
}
