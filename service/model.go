package service

type Request struct {
	UserName  string `json:"userName"`
	Phone     string `json:"phone"`
	FirstName string `json:"firstName"`
	Lastname  string `json:"lastname"`
}

type Response struct {
	Status   string `json:"status"`
	UserName string `json:"userName"`
}
