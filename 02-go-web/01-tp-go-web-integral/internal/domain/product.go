package domain

type Product struct {
	Id           int     `json:"id,omitempty"`
	Name         string  `json:"name,omitempty"`
	Quantity     int     `json:"quantity,omitempty"`
	Code_Value   string  `json:"code___value,omitempty"`
	Is_Published bool    `json:"is___published,omitempty"`
	Expiration   string  `json:"expiration,omitempty"`
	Price        float64 `json:"price,omitempty"`
}
type ProductRequest struct {
	Name       string  `json:"name"`
	Quantity   int     `json:"quantity"`
	Code       string  `json:"code"`
	Published  bool    `json:"published,omitempty"`
	Expiration string  `json:"expiration"`
	Price      float64 `json:"price"`
}

type ProductResponse struct {
	Id         int     `json:"id,omitempty"`
	Name       string  `json:"name,omitempty"`
	Quantity   int     `json:"quantity,omitempty"`
	Code       string  `json:"code,omitempty"`
	Published  bool    `json:"published,omitempty"`
	Expiration string  `json:"expiration,omitempty"`
	Price      float64 `json:"price,omitempty"`
}
