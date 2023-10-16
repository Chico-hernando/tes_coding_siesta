package model_question2

import (
	"tes_coding_siesta/configs"
	"tes_coding_siesta/init/question2"
)

func Question2(lat, long string) (response []question2.Question2, err error) {

	DB := configs.DB

	query := `select *,
	Abs(geo_districts.lat - ` + lat + `) + Abs(geo_districts.long - ` + long + `) AS distance 
	from geo_districts
	ORDER BY distance`

	rows, err := DB.Query(query)

	if err != nil {
		return nil, err
	}

	var res question2.Question2

	defer rows.Close()

	for rows.Next() {

		err = rows.Scan(&res.Id, &res.CityId, &res.Name, &res.Lat, &res.Long, &res.CreatedOn, &res.ModifiedOn, &res.Deleted, &res.Distance)
		if err != nil {
			return nil, err
		}

		response = append(response, res)

	}

	return
}
