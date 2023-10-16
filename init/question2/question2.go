package question2

type Question2 struct {
	Id         int         `json:"id"`
	CityId     int         `json:"city_id"`
	Name       string      `json:"name"`
	Lat        float64     `json:"lat"`
	Long       float64     `json:"long"`
	CreatedOn  string      `json:"created_on"`
	ModifiedOn interface{} `json:"modified_on"`
	Deleted    int         `json:"deleted"`
	Distance   float64     `json:"-"`
}
