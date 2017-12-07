package order

type SpecificationRequest struct {
	Id string
}

type Characteristic struct {
	Name  string
	Value string
}

type ProductRequest struct {
	Characteristics []Characteristic
}

type OrderItemRequest struct {
	Action        string
	Specification SpecificationRequest
	Product       ProductRequest
}

type CreateOrderRequest struct {
	Priority           int
	Category           string
	RequestedStartDate string
	OrderItem          []OrderItemRequest
}
