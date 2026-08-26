package designpatterns

import (
	"encoding/json"
	"fmt"
)

type AccountDetails struct {
	id          string
	accountType string
}

type BankAccount struct {
	details      *AccountDetails
	CustomerName string
}

func (account *BankAccount) setDetails(id string, accountType string) {
	account.details = &AccountDetails{id, accountType}
}

func (account *BankAccount) getId() string {
	return account.details.id
}

func (account *BankAccount) getAccountType() string {
	return account.details.accountType
}

func RunPrivateClassPattern() {
	var account *BankAccount = &BankAccount{CustomerName: "John smith"}
	account.setDetails("4532", "current")
	jsonAccount, _ := json.Marshal(account)
	fmt.Println("Private Class hidden", string(jsonAccount))
	fmt.Println("Account Id", account.getId())
	fmt.Println("Account Type", account.getAccountType())
}
