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

// type pageTitle struct {
// 	title string
// }

// HTMLRender reads the html file and ioWriter writes to localhost:8080
func HTMLRender(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("./repository/Resume.html") // returns []byte of file contents
	errorCheck(err)
	resume := CreateTemplate()
	tmpl.Execute(w, resume)
	// fmt.Fprintf(w, "%s", body) // prints html contents to webpage "localhost:8080"
}

// CreateTemplate builds and returns new ResumeStruct for template
func CreateTemplate() *repo.ResumeStruct {
	res := new(repo.ResumeStruct)
	prog1 := repo.ProgEnvs{Env: "Go, Python, C++, Java, Not HTML"}
	prog2 := repo.ProgEnvs{Env: "iOS, MacOS"}
	res.Prog = append(res.Prog, prog1, prog2)

	math1 := repo.MathExp{Math: "Bayesian Statistics"}
	math2 := repo.MathExp{Math: "Calculus I, II"}
	math3 := repo.MathExp{Math: "Linear Algebra"}
	res.Mathexp = append(res.Mathexp, math1, math2, math3)

	&res.Prof = res.AddProfessional()
	&res.Educ = res.AddEducation()
	return res
}
