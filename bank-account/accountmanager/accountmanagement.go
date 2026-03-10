package accountmanager

import (
	"bank-account-management/model"
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"sync"
)

type StoreAccountData struct {
	StoreAccData map[int]model.AccountManager
	mu           sync.Mutex
}

func InitializeStoreData() *StoreAccountData {

	storedata := make(map[int]model.AccountManager)

	return &StoreAccountData{StoreAccData: storedata}

}

func (sa *StoreAccountData) CreateAccount(accdetails model.AccountManager) (int, error) {

	sa.mu.Lock()
	defer sa.mu.Unlock()
	id := rand.Intn(10000)
	accdetails.Accountnumber = id
	sa.StoreAccData[id] = accdetails

	json.MarshalIndent(accdetails, "", " ")
	data, err := json.Marshal(accdetails)
	if err != nil {
		return 0, err
	} else {
		filename := accdetails.Name
		os.WriteFile(filename, data, 0777)
	}

	return id, nil
}

func (sa *StoreAccountData) GetAccountDetails(id int) (model.AccountManager, error) {

	sa.mu.Lock()
	defer sa.mu.Unlock()

	accmanager, exist := sa.StoreAccData[id]

	if !exist {
		return model.AccountManager{}, fmt.Errorf("Account number does not exist %d", id)
	} else {
		return accmanager, nil
	}

}

func (sa *StoreAccountData) DepositAmount(id int, amount float64) (float64, error) {

	sa.mu.Lock()
	defer sa.mu.Unlock()

	amttobedeposited := amount

	accmanager, exist := sa.StoreAccData[id]
	if !exist {
		return 0, fmt.Errorf("account number %d does not exit", id)
	}
	accmanager.Amount += amttobedeposited
	sa.StoreAccData[id] = accmanager

	return amttobedeposited, nil
}

func (sa *StoreAccountData) WithdrawAmount(id int, amount float64) (float64, error) {

	sa.mu.Lock()
	defer sa.mu.Unlock()

	accmanager, exist := sa.StoreAccData[id]

	if !exist {
		return 0, fmt.Errorf("given account number %d does not exist", id)
	}
	amountinacc := accmanager.Amount

	if amount > amountinacc {
		return 0, fmt.Errorf("withdraw amount %d is greater than the amount in the account", amount)
	}

	accmanager.Amount -= amount

	sa.StoreAccData[id] = accmanager

	return accmanager.Amount, nil
}
