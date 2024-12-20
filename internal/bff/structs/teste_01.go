package structs

type ResponseTestExample struct {
	StringKey string   `json:"stringKey"`
	IntKey    int      `json:"intKey"`
	FloatKey  float64  `json:"floatKey"`
	BoolKey   bool     `json:"boolKey"`
	SliceKey  []string `json:"sliceKey"`
	SliceKey2 []any    `json:"sliceKey2"`
	Complex   struct {
		Complex_keyValue string `json:"complex_keyValue"`
	} `json:"complex"`
	MapKey struct {
		NestedKey string `json:"nestedKey"`
	} `json:"mapKey"`
}
