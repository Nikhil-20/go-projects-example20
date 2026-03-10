package bankstore

import "bank-account-management/model"

type BankManagement interface {
	CreateAccount(account model.AccountManager) (int, error)
	GetAccountDetails(id int) (model.AccountManager, error)
	DepositAmount(id int, amount int) (int, error)
	WithdrawAmount(id int, amount int) (int, error)
	//TransactionHistroy(id int)
}
