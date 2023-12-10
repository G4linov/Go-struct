package main

import (
	"errors"
	"fmt"
	"v2/internal"
)

func main() {

	cust := internal.NewCustomer("dmitry", 23, 10000, 1000, true)

	const DEFAULT_DISCOUNT = 500

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

	fmt.Printf("%+v\n", cust)
}
