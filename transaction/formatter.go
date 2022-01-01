package transaction

import "time"

type TransactionFormatter struct {
	ID        int       `json: "id"`
	Name      string    `json: "name"`
	Amount    int       `json : "amount"`
	CreatedAt time.Time `json : "cretaed_at"`
}

func FormatTransaction(transaction Transaction) TransactionFormatter {
	TransactionFormatter := TransactionFormatter{}
	TransactionFormatter.ID = transaction.ID
	TransactionFormatter.Name = transaction.User.Name
	TransactionFormatter.Amount = transaction.Amount
	TransactionFormatter.CreatedAt = transaction.CreatedAt

	return TransactionFormatter
}

func FormatListTrans(transaction []Transaction) []TransactionFormatter {
	if len(transaction) == 0 {
		return []TransactionFormatter{}
	}

	var transactionFormatter []TransactionFormatter

	for _, transactions := range transaction {
		formatter := FormatTransaction(transactions)
		transactionFormatter = append(transactionFormatter, formatter)
	}

	return transactionFormatter
}
