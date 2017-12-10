package order

import "fmt"

type Service struct {
	OrderRepository RedisRepository `inject:"inline"`
}

func (service *Service) CreateOrderService(request CreateOrderRequest) (*order, error) {

	order := NewOrder(request.Priority, request.Category)

	fmt.Println(order)

	service.OrderRepository.save(order)

	return order, nil
}

func (service *Service) GetOrderByIdService(orderI string) (*order, error) {

	order, err := service.OrderRepository.getById(orderI)

	return order, err
}
