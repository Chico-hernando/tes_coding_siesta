package model_question1

import (
	"tes_coding_siesta/configs"
	"tes_coding_siesta/init/question1"
)

func Question1() (response []question1.Question1, err error) {

	DB := configs.DB

	query := `select c.charge_id, c.selling_date, ct.charge_type, sum(c.amount*tt.percentage) as 'tax-calculated totals' from charge c
	join charge_type ct on ct.charge_type_id = c.charge_type_id
	join charge_type_tax_list cttl on cttl.charge_type_id = ct.charge_type_id
	join tax_type tt on tt.tax_type_id = cttl.tax_type_id
	where c.is_deleted = false and ct.is_deleted = false
	group by c.charge_type_id, c.selling_date
	order by c.charge_id`

	rows, err := DB.Query(query)

	if err != nil {
		return nil, err
	}

	var res question1.Question1

	defer rows.Close()

	for rows.Next() {

		err = rows.Scan(&res.ChargeId, &res.SellingDate, &res.ChargeType, &res.TaxCalTotal)
		if err != nil {
			return nil, err
		}

		response = append(response, res)

	}

	return
}