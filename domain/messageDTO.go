package domain

type Entity interface {
	Validate() error
}

type MessageDTO struct {
	Entity        Entity `json:"entity"`
	Action        string `json:"action"`
	Functionality string `json:"functionality"`
	Company       string `json:"company"`
}
