package process_transaction

type TransactionInputDTO struct {
	ID                        string  `json:"id"`
	AccountID                 string  `json:"account_id"`
	Amount                    float64 `json:"amount"`
	CreditCardNumber          string  `json:"credit_card_number"`
	CreditCardName            string  `json:"credit_card_name"`
	CreditCardExpirationMonth int     `json:"credit_card_expiration_month"`
	CreditCardExpirationYear  int     `json:"credit_card_expiration_year"`
	CreditCardCVV             string  `json:"credit_card_cvv"`
}

type TransactionOutputDTO struct {
	ID           string `json:"id"`
	Status       string `json:"status"`
	ErrorMessage string `json:"error_message"`
}
