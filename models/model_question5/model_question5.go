package model_question5

import (
	"encoding/json"
	"io/ioutil"
	"path/filepath"
	"tes_coding_siesta/init/question5"

	"golang.org/x/exp/slices"
)

func Question5() (response []question5.Question5, err error) {

	var class question5.Class
	var sick, permittedLeaves, totalSick, totalPermittedLeaves int

	err = fetchData(&class)
	if err != nil {
		return response, err
	}

	for _, value := range class.Students {

		sick, permittedLeaves, totalSick, totalPermittedLeaves = 0, 0, 0, 0

		sick = slices.IndexFunc(class.Sick, func(c question5.Student) bool { return c.Id == value.Id })
		if sick >= 0 {
			for _, sickValue := range class.Sick {
				if sickValue.Id == value.Id {
					totalSick++
				}
			}
		}

		permittedLeaves = slices.IndexFunc(class.PermittedLeaves, func(c question5.Student) bool { return c.Id == value.Id })
		if permittedLeaves >= 0 {
			for _, permittedLeavesValue := range class.PermittedLeaves {
				if permittedLeavesValue.Id == value.Id {
					totalPermittedLeaves++
				}
			}
		}

		response = append(response, question5.Question5{
			Id:                   value.Id,
			Name:                 value.Name,
			Nisn:                 value.Nisn,
			Class:                value.Class,
			School:               value.School,
			TotalSick:            totalSick,
			TotalPermittedLeaves: totalPermittedLeaves,
			TotalPresent:         class.Classes - totalSick - totalPermittedLeaves,
		})
	}

	return
}

func fetchData(q *question5.Class) (err error) {
	absPath, err := filepath.Abs("file/student.json")
	if err != nil {
		return err
	}

	data, err := ioutil.ReadFile(absPath)
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, &q)
	if err != nil {
		return err
	}

	return nil
}
