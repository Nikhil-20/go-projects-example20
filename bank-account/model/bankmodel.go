package model

import "fmt"

type AccountType string

const (
	Savings AccountType = "Savings"
	Current AccountType = "Current"
)

func (a AccountType) IsValid() bool {
	switch a {
	case Savings, Current:
		return true
	}
	return false
}

// ValidateAccountType parses and validates a string input
func ValidateAccountType(input string) (AccountType, error) {
	at := AccountType(input)
	if !at.IsValid() {
		return "", fmt.Errorf("invalid account type %q: must be 'Savings' or 'Current'", input)
	}
	return at, nil
}

type AccountManager struct {
	Name           string      `json:"name"`
	Age            int         `json:"age"`
	Address        string      `json:"address"`
	Nationality    string      `json:"nationality"`
	Accounttype    AccountType `json:"accounttype"`
	Aadhaar_number string      `json:"aadhaar"`
	Accountnumber  int         `json:"accountnumber"`
	Amount         float64     `json:"amount"`
}

type Deposit struct {
	Amount_deposit int
	//Balance_amount  int
	Withdraw_amount int
	Accountnumber   int
	//mu              sync.Mutex
}
