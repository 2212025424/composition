package customer

type Customer struct {
	name string
	address string
	phone string
}

func New (name, addressm, phone string) Customer {
	return Customer{name, address, phone}
}