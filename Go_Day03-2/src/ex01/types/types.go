package types

type Place struct {
	Name    string `json:"name"`
	Address string `json:"address"`
	Phone   string `json:"phone"`
	Id      int    `json:"id"`
}

type RestPage struct {
	TotalNumberOfRests int
	CurPage            int
	PrevPage           int
	NextPage           int
	TotalPages         int
	Rests              []Place
}
