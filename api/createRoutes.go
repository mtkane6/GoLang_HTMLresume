package api

import (
	"log"
	"net/http"

	"../domain"
)

// BeginServer opens server and calls HtmlRenderer function
func BeginServer() {
	http.HandleFunc("/", domain.HTMLRender)
	http.Handle("/repository/", http.StripPrefix("/repository/", http.FileServer(http.Dir("repository"))))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
