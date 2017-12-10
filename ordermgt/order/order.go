package order

import "github.com/satori/go.uuid"

type specification struct {
	Id string
}

type characteristic struct {
	Name  string
	Value string
}

type product struct {
	Characteristics []characteristic
}

type orderItem struct {
	Action        string
	Specification specification
	Product       product
}

type order struct {
	Id                 uuid.UUID `json:"id"`
	Priority           int `json:"priority"`
	Category           string `json:"category"`
	RequestedStartDate string `json:"requestedStarDate"`
	OrderItem          []orderItem `json:"orderItem"`
}

func NewOrder(priority int, category string) *order {
	return &order{
		Id:       uuid.NewV4(),
		Priority: priority,
		Category: category,
	}
}

func (o *order) GetId() uuid.UUID {
	return o.Id
}
