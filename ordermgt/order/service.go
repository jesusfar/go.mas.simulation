package order

type Service struct {
	OrderRepository RedisRepository `inject:"inline"`
}

func (service *Service) CreateOrderService(request CreateOrderRequest) (CreateOrderRequest, error) {

	conn := service.OrderRepository.DBPoolConn.Get()

	order := NewOrder()

	conn.Do("SET", order.GetId(), order)

	return request, nil
}
