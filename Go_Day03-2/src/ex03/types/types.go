package types

type Location struct {
	Latitude  float64 `json:"lat"`
	Longitude float64 `json:"lon"`
}

type Place struct {
	Id       int      `json:"id"`
	Name     string   `json:"name"`
	Address  string   `json:"address"`
	Phone    string   `json:"phone"`
	Location Location `json:"location"`
}

type RestPage struct {
	TotalNumberOfRests int
	CurPage            int
	PrevPage           int
	NextPage           int
	TotalPages         int
	Rests              []Place
}
