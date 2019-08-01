package api

import (
	"log"
	"net/http"

	"../domain"
)

// BeginServer opens server and calls HtmlRenderer function
func BeginServer() {
	http.HandleFunc("/", domain.HTMLRender)
	// http.HandleFunc("/login", login)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
