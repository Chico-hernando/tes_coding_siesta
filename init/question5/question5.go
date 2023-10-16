package question5

type Question5 struct {
	Id                   int    `json:"id"`
	Name                 string `json:"name"`
	Nisn                 string `json:"nisn"`
	Class                string `json:"class"`
	School               string `json:"scholl"`
	TotalSick            int    `json:"total_sick"`
	TotalPermittedLeaves int    `json:"total_permitted_leaves"`
	TotalPresent         int    `json:"total_present"`
}

type Class struct {
	Classes         int       `json:"classes"`
	Students        []Student `json:"students"`
	Sick            []Student `json:"sicks"`
	PermittedLeaves []Student `json:"permitted_leaves"`
}

type Student struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Nisn   string `json:"nisn"`
	Class  string `json:"class"`
	School string `json:"school"`
}
