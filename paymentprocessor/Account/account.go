package account

import "errors"

type Account interface {
	NewUser(variable ...interface{}) error
}

type CreditCardPayment struct {
	Id         int
	CardNumber string
	Expiry     string
	Cvv        string
	Amount     float64
}

func (ccp *CreditCardPayment) NewUser(CardNumber string, Expiry string, Cvv string, initDepo float64) error {

	if ccp.CardNumber == "" && ccp.Expiry == "" && ccp.Cvv == "" {
		return errors.New("Need valid CardNumber, with proper expiry date and cvv")
	}

	return nil
}

type PaypalPayment struct {
	Id     int
	Mail   string
	Pass   string
	Amount float64
}

func (pp *PaypalPayment) NewUser(Mail string, Pass string, initDepo float64) error {
	return nil
}

type CashPayment struct {
	receipt bool
}

func (cp *CashPayment) pay(amount float64) error {
	return nil
}
