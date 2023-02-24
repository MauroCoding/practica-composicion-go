package customers

type customer struct {
	name, city, country, email string
}

type Customer = customer

func NewCustomer(name, city, country, email string) customer {
	return customer{name, city, country, email}
}
