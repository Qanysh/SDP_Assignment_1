package main

import "fmt"

type PaymentStrategy interface {
	Pay(amount float64) string
}

type CreditCardPayment struct {
	CardNumber     string
	ExpirationDate string
	CVV            string
}

func (c *CreditCardPayment) Pay(amount float64) string {
	return fmt.Sprintf("Оплата с использованием кредитной карты: %.2f $. С карты %s", amount, c.CardNumber)
}

type PayPalPayment struct {
	Email string
}

func (p *PayPalPayment) Pay(amount float64) string {
	return fmt.Sprintf("Оплата через PayPal: %.2f $. по адресу %s", amount, p.Email)
}

type ShoppingCart struct {
	PaymentStrategy PaymentStrategy
}

func (sc *ShoppingCart) ProcessPayment(amount float64) string {
	return sc.PaymentStrategy.Pay(amount)
}

func main() {
	creditCardPayment := &CreditCardPayment{
		CardNumber:     "1234-5678-1234-5678",
		ExpirationDate: "12/24",
		CVV:            "123",
	}

	shoppingCart := &ShoppingCart{
		PaymentStrategy: creditCardPayment,
	}

	fmt.Println(shoppingCart.ProcessPayment(100.00))

	payPalPayment := &PayPalPayment{
		Email: "aitu_2023@astanait.edu.kz",
	}

	shoppingCart.PaymentStrategy = payPalPayment
.
	fmt.Println(shoppingCart.ProcessPayment(50.00))
}
