package banking

import (
	"errors"
)

// Account struct
type Account struct {
	owner   string
	balance int
}

// NewAccount creates Account
func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	return &account // 새로 만들어진 object 를 리턴
}

// Deposit x amount on your account
// between the func and name(Deposit)
// a is called as Receiver
// pointer receiver -> don't make a copy of object, just use the object.
func (a *Account) Deposit(amount int) {
	a.balance += amount
}

// Balance of your Account
func (a Account) Balance() int {
	return a.balance
}

func (a *Account) Withdraw(amount int) error {
	if a.balance < amount {
		return errors.New("cant withdraw you are poor")
	}
	a.balance -= amount
	return nil
}

// java 처럼 클래스(struct) 내부에 method 를 작성하지 않는다.

// ChangeOwner of the account
func (a *Account) ChangeOwner(newOwner string) {
	a.owner = newOwner
}

func (a Account) Owner() string {
	return a.owner
}

//func (a Account) String() string {
//	return fmt.Sprint(a.Owner(), "'s account.\nHas: ", a.Balance())
//}
