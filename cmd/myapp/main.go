package main

import (
	"fmt"
	"v2/internal"
)

func main() {
	cust := internal.Customer{
		Age:     23,
		Balance: 10000,
		Debt:    1000,
		Name:    "Dmitry",
	}

	fmt.Printf("%+v\n", cust)
}
