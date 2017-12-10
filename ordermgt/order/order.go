package order

type specification struct {
	id string
}

type characteristic struct {
	name  string
	value string
}

type product struct {
	characteristics []characteristic
}

type orderItem struct {
	action        string
	specification specification
	product       product
}

type order struct {
	id int
	priority           int
	category           string
	requestedStartDate string
	orderItem          []orderItem
}

func NewOrder() (*order) {
	return &order{}
}

func (o *order) GetId() int {
	return o.id
}