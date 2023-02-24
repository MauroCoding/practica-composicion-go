package main

import (
	"fmt"

	"github.com/MauroCoding/practica-composicion-go/pkg/customers"
	"github.com/MauroCoding/practica-composicion-go/pkg/invoices"
	"github.com/MauroCoding/practica-composicion-go/pkg/items"
)

func main() {
	name := "Mauricio González"
	city := "Yerbas Buenas"
	country := "Chile"
	email := "alguien@algo.net"

	coca_cola := items.NewItem("Coca-Cola", 100, 0.99)
	pepsi := items.NewItem("Pepsi", 101, 0.97)
	doritos := items.NewItem("Doritos", 102, 0.40)

	customer := customers.NewCustomer(name, city, country, email)

	invoice := invoices.NewInvoice(customer, items.NewItemSlice(coca_cola, pepsi, doritos))

	invoices.SetTotal(&invoice)

	fmt.Printf("%T\n%+v", invoice, invoice)
}
