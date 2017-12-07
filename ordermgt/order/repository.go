package order

type OrderRepository interface {
	save(order order)
}
