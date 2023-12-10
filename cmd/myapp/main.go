package main

import (
	"errors"
	"fmt"
	"v2/internal"
)

func main() {

	const DEFAULT_DISCOUNT = 500

	cust := internal.Customer{
		Age:     23,
		Balance: 10000,
		Debt:    1000,
		Name:    "Dmitry",
	}

	cust.CalcDiscount = func() (int, error) {
		if !cust.Discount {
			return 0, errors.New("discount not available")
		}
		result := DEFAULT_DISCOUNT - cust.Debt
		if result < 0 {
			return 0, nil
		}
		return result, nil
	}

	discount, _ := cust.CalcDiscount()

	fmt.Printf("%d", discount)
}
