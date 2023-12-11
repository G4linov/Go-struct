package internal

import "errors"

const DEFAULT_DISCOUNT = 500

type Debtor interface {
	WrOffDebt() error
}

type Discounter interface {
	CalcDiscount() (int, error)
}

type Overduer struct {
	balance int
	debt    int
}

type Customer struct {
	*Overduer
	Name     string
	Age      int
	discount bool
}

func (c *Customer) WrOffDebt() error {
	if c.debt >= c.balance {
		return errors.New("not possible write off")
	}

	c.balance -= c.debt
	c.debt = 0

	return nil
}

func (c *Customer) CalcDiscount() (int, error) {
	if !c.discount {
		return 0, errors.New("discount not available")
	}
	result := DEFAULT_DISCOUNT - c.debt
	if result < 0 {
		return 0, nil
	}
	return result, nil
}

func NewOverDeur(balance int, debt int) *Overduer {
	return &Overduer{
		balance: balance,
		debt:    debt,
	}
}

func NewCustomer(name string, age int, balance int, debt int, discount bool) *Customer {
	return &Customer{
		Overduer: NewOverDeur(balance, debt),
		Name:     name,
		Age:      age,
		discount: discount,
	}
}

type Partner struct {
	Name    string
	Age     int
	balance int
	debt    int
}

func (c *Partner) WrOffDebt() error {
	c.debt = 0

	return nil
}

func NewPartner(name string, age int, balance int, debt int) *Partner {
	return &Partner{
		Name:    name,
		Age:     age,
		balance: balance,
		debt:    debt,
	}
}
