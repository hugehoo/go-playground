package accounts

import "errors"

// Account struct
type Account struct {
	owner   string
	balance int
}

// struct 를 private 하게 만들면, function 을 통해 struct 를 생성할 수 있다.

// NewAccount creates Account
func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	return &account // object 를 리턴, 실제 메모리 주소를 리턴 (복사본을 만들지 않는다)
}

// this is method (different with function) (a Account) is called as receiver.

// DepositCopy : copies the account object, which is different with method caller. that's why the balance doesn't change.
func (account Account) DepositCopy(amount int) {
	account.balance += amount
}

// Deposit Pointer : don't copy the method caller, just receive the method caller object.
func (account *Account) Deposit(amount int) {
	account.balance += amount
}

// Balance of your account
func (account Account) Balance() int {
	return account.balance
}

// Withdraw amount from account
func (account *Account) Withdraw(amount int) error {
	if account.balance < amount {
		//return errors.New("Cant withdraw")
		return errNoMoney
	}
	account.balance -= amount
	return nil
}

var errNoMoney = errors.New("can't withdraw")
