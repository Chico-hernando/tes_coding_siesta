package model_question4

import (
	"sync"
	"tes_coding_siesta/init/question4"
)

func Question4() (response []question4.Question4, err error) {

	users := []question4.Question4{
		{Name: "John", Gender: "F"},
		{Name: "Jack", Gender: "F"},
		{Name: "May", Gender: "M"},
		{Name: "Sonia", Gender: "M"},
	}

	wg := new(sync.WaitGroup)
	for i, value := range users {
		wg.Add(1)

		go func(i int, value question4.Question4) {
			defer wg.Done()
			if value.Gender == "F" {
				users[i].Gender = "Females"
			} else if value.Gender == "M" {
				users[i].Gender = "Male"
			}
		}(i, value)
	}

	wg.Wait()

	response = users

	return
}
