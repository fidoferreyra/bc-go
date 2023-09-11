package pkg

type ProductDTO struct {
	Name         string  `json:"name,omitempty"`
	Quantity     int     `json:"quantity,omitempty"`
	Code_Value   string  `json:"code___value,omitempty"`
	Is_Published bool    `json:"is___published,omitempty"`
	Expiration   string  `json:"expiration,omitempty"`
	Price        float64 `json:"price,omitempty"`
}