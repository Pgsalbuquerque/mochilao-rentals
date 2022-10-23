package types

type Rental struct {
	Fields Fields `json:"fields"`
}

type Fields struct {
	City        string  `json:"city"`
	Country     string  `json:"country"`
	Name        string  `json:"name"`
	Summary     string  `json:"summary"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Street      string  `json:"street"`
}
