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

type TransUserFormatter struct {
	ID        int               `json: "id"`
	Amount    int               `json : "amount"`
	Status    string            `json: "status"`
	CreatedAt time.Time         `json : "cretaed_at"`
	Campaign  CampaignFormatter `json : "campaign"`
}

type CampaignFormatter struct {
	Title    string `json: "title"`
	ImageUrl string `json: "image_url"`
}

func FormatTransUser(transaction Transaction) TransUserFormatter {
	TransUserFormatter := TransUserFormatter{}
	TransUserFormatter.ID = transaction.ID
	TransUserFormatter.Amount = transaction.Amount
	TransUserFormatter.Status = transaction.Status
	TransUserFormatter.CreatedAt = transaction.CreatedAt

	campaignFormatter := CampaignFormatter{}
	campaignFormatter.Title = transaction.Campaign.Title
	if len(transaction.Campaign.CampaignImages[0].FileName) > 0 {
		campaignFormatter.ImageUrl = ""
	}
	campaignFormatter.ImageUrl = transaction.Campaign.CampaignImages[0].FileName

	TransUserFormatter.Campaign = campaignFormatter

	return TransUserFormatter
}

func FormatListTransUser(transaction []Transaction) []TransUserFormatter {
	if len(transaction) == 0 {
		return []TransUserFormatter{}
	}

	var transUserFormatter []TransUserFormatter

	for _, transactions := range transaction {
		formatter := FormatTransUser(transactions)
		transUserFormatter = append(transUserFormatter, formatter)
	}

	return transUserFormatter
}

type SuccessFormat struct {
	CampaignID int       `json: "campaign_id"`
	UserID     int       `json: "user_id"`
	Amount     int       `json : "amount"`
	Code       string    `json: code`
	PaymentURL string    `json: paymeny_url`
	CreatedAt  time.Time `json : "cretaed_at"`
}

func SuccessFormatter(transaction Transaction) SuccessFormat {
	SuccessTrans := SuccessFormat{}
	SuccessTrans.CampaignID = transaction.CampaignID
	SuccessTrans.UserID = transaction.UserID
	SuccessTrans.Amount = transaction.Amount
	SuccessTrans.CreatedAt = transaction.CreatedAt
	SuccessTrans.Code = transaction.Code
	SuccessTrans.PaymentURL = transaction.PaymentURL

	return SuccessTrans
}
