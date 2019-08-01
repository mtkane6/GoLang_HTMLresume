package domain

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func errorCheck(e error) {
	if e != nil {
		panic(e)
	}
}

// HTMLRender reads the html file and ioWriter writes to localhost:8080
func HTMLRender(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadFile("./repository/Resume.html") // returns []byte of file contents
	errorCheck(err)
	fmt.Fprintf(w, "%s", body) // prints html contents to webpage "localhost:8080"
}
