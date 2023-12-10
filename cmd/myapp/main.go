package main

import (
	"errors"
	"fmt"
	"v2/internal"
)

func CalcPrice(cust internal.Customer, price int) (int, error) {
	if cust.Discount {
		return (price - 100), nil
	}
	return 0, errors.New("discount not available")
}

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

	cust.Discount = false

	newPirce, err := CalcPrice(*cust, 1000)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("New price %d\n", newPirce)
	}
}
