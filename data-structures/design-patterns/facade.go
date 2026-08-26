package designpatterns

import "fmt"

type Account struct {
	id          int32
	accountType string
}

func (account *Account) createAccount(accountType string) *Account {
	fmt.Println("account creation with type")
	account.accountType = accountType
	return account
}

func (account *Account) getById(id string) *Account {
	fmt.Println("getting account by id")
	return account
}

func (account *Account) deleteById(id string) {
	fmt.Println("deleting account by id")
}

type Customer struct {
	name string
	id   string
}

func (customer *Customer) create(name string) *Customer {
	fmt.Println("customer creation")
	customer.name = name
	return customer
}

type Transaction struct {
	id            string
	amount        float32
	srcAccountId  string
	destAccountId string
}

func (transaction *Transaction) create(srcAccountId string, destAccountId string, amount float32) *Transaction {
	fmt.Println("creating Transaction")
	transaction.srcAccountId = srcAccountId
	transaction.destAccountId = destAccountId
	transaction.amount = amount
	return transaction
}

type BranchManagerFacade struct {
	account     *Account
	customer    *Customer
	transaction *Transaction
}

func NewBranchManagerFacade() *BranchManagerFacade {
	return &BranchManagerFacade{&Account{}, &Customer{}, &Transaction{}}
}

func (facade *BranchManagerFacade) createCustomerAccount(customerName string, accountType string) (*Customer, *Account) {
	var customer = facade.customer.create(customerName)
	var account = facade.account.createAccount(accountType)
	return customer, account
}

func (facade *BranchManagerFacade) createTransaction(srcAccountId string, destAccountId string, amount float32) *Transaction {
	var transaction = facade.transaction.create(srcAccountId, destAccountId, amount)
	return transaction
}

func RunFacadePatterrn() {
	fmt.Println("--> Facade Pattern")
	var facade = NewBranchManagerFacade()
	var customer *Customer
	var account *Account
	customer, account = facade.createCustomerAccount("Thomas Smith", "Savings")
	fmt.Println(customer.name)
	fmt.Println(account.accountType)
	var transaction *Transaction
	transaction = facade.createTransaction("23455", "12897", 1000)
	fmt.Println(transaction.amount)
}
