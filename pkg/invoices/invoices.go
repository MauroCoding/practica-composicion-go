package invoices

import (
	"github.com/MauroCoding/practica-composicion-go/pkg/customers"
	"github.com/MauroCoding/practica-composicion-go/pkg/items"
)

type invoice struct {
	customer customers.Customer
	items    items.Items
	total    float64
}

func NewInvoice(customer customers.Customer, items items.Items) invoice {
	return invoice{
		customer, items, 0,
	}
}

func SetTotal(i *invoice) {
	for _, v := range i.items {
		i.total += v.GetValue()
	}
}
