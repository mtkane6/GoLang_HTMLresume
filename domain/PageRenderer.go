package domain

import (
	"html/template"
	"net/http"

	repo "../repository"
)

func errorCheck(e error) {
	if e != nil {
		panic(e)
	}
}

// HTMLRender reads the html file and ioWriter writes to localhost:8080
func HTMLRender(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("./repository/resume.html") // returns []byte of file contents
	errorCheck(err)
	Resume := CreateResume()
	tmpl.Execute(w, Resume)
	// fmt.Fprintf(w, "%s", body) // prints html contents to webpage "localhost:8080"
}

// CreateResume builds and returns new ResumeStruct for template
func CreateResume() repo.ResumeStruct {
	var Res repo.ResumeStruct

	prog1 := repo.ProgEnvs{Env: "Go, Python, C++, Java, Not HTML"}
	prog2 := repo.ProgEnvs{Env: "iOS, MacOS"}
	Res.Prog = append(Res.Prog, prog1, prog2)

	Res.Prog = Res.AddProg()
	Res.Mathexp = Res.AddMath()
	Res.Prof = Res.AddProfessional()
	Res.Educ = Res.AddEducation()
	return Res
}
