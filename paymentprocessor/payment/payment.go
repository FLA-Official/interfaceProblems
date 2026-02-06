package payment

import "errors"

type PaymentProcessor interface {
	pay(amount float64) error
}

type CreditCardPayment struct {
	CardNumber string
	Expiry     string
	Cvv        string
}

func (ccp *CreditCardPayment) pay(amount float64) error {

	if ccp.CardNumber == "" && ccp.Expiry == "" && ccp.Cvv == "" {
		return errors.New("Need valid CardNumber, with proper expiry date and cvv")
	}

	return nil
}

type PaypalPayment struct {
	Mail string
	Pass string
}

func (pp *PaypalPayment) pay(amount float64) error {
	return nil
}

type CashPayment struct {
	receipt bool
}

func (cp *CashPayment) pay(amount float64) error {
	return nil
}
