package main

import (
	"fmt"
	"v2/internal"
)

func main() {

	cust := internal.NewCustomer("dmitry", 23, 10000, 1000, true)

	_ = cust.WrOffDebt()

	fmt.Printf("%+v\n", cust)
}
