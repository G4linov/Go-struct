package main

import (
	"fmt"
	"v2/internal"
)

func CalcPrice(discounter internal.Discounter, price int) (int, error) {
	result, err := discounter.CalcDiscount()
	if err != nil {
		return result, err
	}
	return price - result, nil
}

func main() {

	cust := internal.NewCustomer("dmitry", 23, 10000, 1000, true)
	partner := internal.NewPartner("dmitrty", 23, 10000, 1000)

	startTransaction(cust)
	startTransaction(partner)

	fmt.Printf("%+v\n", cust)
	fmt.Printf("%+v\n", partner)
}

func startTransaction(debtor internal.Debtor) error {
	return debtor.WrOffDebt()
}
