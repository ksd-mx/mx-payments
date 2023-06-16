package entity

import (
	"errors"
	"regexp"
	"time"
)

type CreditCard struct {
	number          string
	name            string
	expirationMonth int
	expirationYear  int
	cvv             string
}

func NewCreditCard(number string, name string, expirationMonth int, expirationYear int, cvv string) (*CreditCard, error) {
	cc := &CreditCard{
		number:          number,
		name:            name,
		expirationMonth: expirationMonth,
		expirationYear:  expirationYear,
		cvv:             cvv,
	}

	err := cc.IsValid()
	if err != nil {
		return nil, err
	}
	return cc, nil
}

func (c *CreditCard) IsValid() error {
	err := c.validateNumber()
	if err != nil {
		return err
	}
	err = c.validateExpirationMonth()
	if err != nil {
		return err
	}
	err = c.validateExpirationYear()
	if err != nil {
		return err
	}
	err = c.validateCvv()
	if err != nil {
		return err
	}
	return nil
}

func (c *CreditCard) validateNumber() error {
	re := regexp.MustCompile(`^(?:4[0-9]{12}(?:[0-9]{3})?|[25][1-7][0-9]{14}|6(?:011|5[0-9][0-9])[0-9]{12}|3[47][0-9]{13}|3(?:0[0-5]|[68][0-9])[0-9]{11}|(?:2131|1800|35\d{3})\d{11})$`)
	if !re.MatchString(c.number) {
		return errors.New("Invalid credit card number")
	}
	return nil
}

func (c *CreditCard) validateExpirationMonth() error {
	if c.expirationMonth < 01 || c.expirationMonth > 12 {
		return errors.New("Invalid expiration month")
	}
	return nil
}

func (c *CreditCard) validateExpirationYear() error {
	if c.expirationYear < time.Now().Year() {
		return errors.New("Invalid expiration year")
	}
	return nil
}

func (c *CreditCard) validateCvv() error {
	if len(c.cvv) != 3 {
		return errors.New("Invalid CVV")
	}
	return nil
}
