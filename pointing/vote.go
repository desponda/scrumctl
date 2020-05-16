package pointing

type Vote struct {
	UserId   string `json:"userId"`
	UserName string `json:"userName"`
	Vote     int    `json:"vote"`
}
