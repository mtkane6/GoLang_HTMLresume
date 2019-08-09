package repository

// ResumeStruct holds template values for resume
type ResumeStruct struct {
	Prog    []ProgEnvs
	Mathexp []MathExp
	Prof    []ProfExp
	Educ    []Education
}

// ProgEnvs lists programming environments
type ProgEnvs struct {
	Env string
}

// MathExp lists math experience
type MathExp struct {
	Math string
}

// ProfExp lists work experience
type ProfExp struct {
	Company string
	Dates   string
	Duties  []string
}

// Education lists education experience
type Education struct {
	School string
	Topics string
	Dates  string
}

// AddProg appends instances of programming experience to the resume
func (r *ResumeStruct) AddProg() []ProgEnvs {
	var Pro []ProgEnvs
	prog1 := ProgEnvs{Env: "Go, Python, C++, Java, Not HTML"}
	prog2 := ProgEnvs{Env: "iOS, MacOS"}
	Pro = append(Pro, prog1, prog2)
	return Pro
}

// AddMath appends instances of math experience to the resume
func (r *ResumeStruct) AddMath() []MathExp {
	var M []MathExp
	math1 := MathExp{Math: "Bayesian Statistics"}
	math2 := MathExp{Math: "Calculus I, II"}
	math3 := MathExp{Math: "Linear Algebra"}
	M = append(M, math1, math2, math3)
	return M
}

// AddProfessional appends instances of work experience to the resume
func (r *ResumeStruct) AddProfessional() []ProfExp {
	var Pro []ProfExp
	prof1 := ProfExp{
		Company: "Genius - Apple, Inc.",
		Dates:   "April 2017 - Present",
		Duties: []string{
			"Proficient with MacOS and iOS platforms, technical troubleshooting.",
			"Average 50 Genius bar appointments per week; Physical repairs of iPhone and Mac.",
			"Promoted from Technical Specialist to Technical Expert to Genius; Train others the same.",
		},
	}
	prof2 := ProfExp{
		Company: "Software Engineer Intern",
		Duties: []string{
			"Experienced agile development process with scrum meetings and story boards.",
			"Developed web app weather microservice using Go, Docker, Jenkins, and PostGres.",
			"Utilized GitFlow to control project development progress.",
			"Presented to team, as well as University colloquium.",
		},
	}
	prof3 := ProfExp{
		Company: "Biotech Lab Manager - Celgene",
		Dates:   "Feb 2016 - March 2017",
		Duties: []string{
			"Organized, negotiated, purchased contracts for consumables, lab capital equipment, and service agreements.",
			"Coordinate EH&S compliance, corporate facilities liaison, responsible for all DOT/IATA shipping and receiving.",
		},
	}
	Pro = append(Pro, prof1, prof2, prof3)
	return Pro
}

// AddEducation appends instances of education to the resume
func (r *ResumeStruct) AddEducation() []Education {
	var Ed []Education
	educ1 := Education{
		School: "University of Washington",
		Topics: "B.S. Computer Science and Software Engineering",
		Dates:  "2020 Grad.",
	}
	educ2 := Education{
		School: "Bellevue College",
		Topics: "Java, JavaScript, Calculus, Statistics",
		Dates:  "2017 - 2018",
	}
	educ3 := Education{
		School: "Shoreline Community College",
		Topics: "Biotechnology Lab Technician Program",
		Dates:  "2014 - 2015",
	}
	educ4 := Education{
		School: "University of Washington",
		Topics: "B.S. Biology",
		Dates:  "2007 - 2011",
	}
	educ5 := Education{
		School: "Northway Aviation",
		Topics: "Private Pilot Certificate",
		Dates:  "2011",
	}
	Ed = append(Ed, educ1, educ2, educ3, educ4, educ5)
	return Ed
}
