package question1

type Question1 struct {
	ChargeId    int    `json:"charge_id"`
	SellingDate string `json:"selling_date"`
	ChargeType  string `json:"charge_type"`
	TaxCalTotal string `json:"tax_calculated_total"`
}
