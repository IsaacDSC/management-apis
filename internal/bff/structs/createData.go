package structs

type CreateData struct {
	Teste   float64 `json:"teste"`
	Complex struct {
		KeyComplex string `json:"keyComplex"`
	} `json:"complex"`
	List []float64 `json:"list"`
}
