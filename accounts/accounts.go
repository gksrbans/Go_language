package accounts

import (
	"errors"
	"fmt"
)

// Account is ...
// 구조체 선언 => export 해야하기 때문에 대문자로 선언함.
type Account struct {

	// 내부 속성들은 소문자일 경우 private, 대문자일경우 public
	// Owner   string
	// Balance int

	// Public은 아무나 값을 변경할 수 있기 때문에 소문자로 선언함
	owner   string
	balance int
}

var errnoMoney = errors.New("Can't withdraw")

// NewAccount is creates Account
// python이나 java와는 다르게 constructor를 만들수 없음.
// 그렇기 때문에 함수를 이용해 해결함
// 구조체를 복사하지 않고 pointer 이용해 실제 Account의 요소들에 접근함.
func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	return &account // 해당 주소를 return
}

// Deposit x amount on your accounts
// 선언 함수 이름 사이에 'recevier'
// 구조체 Account의 첫번째 글자를 소문자로 시작함.
func (a *Account) Deposit(amount int) {
	a.balance += amount
}

// Balance of your accounts
func (a Account) Balance() int {
	return a.balance
}

// Withdraw x amount from your accounts
// Go에는 try-catch 혹은 exception의 개념이 없음.
// Go는 무조건적으로 에러처리를 해야하는 프로그래밍을 해야함
func (a *Account) Withdraw(amount int) error {
	if a.balance < amount {
		return errnoMoney
	}
	a.balance -= amount
	return nil
}

// ChangeOwner of the account
func (a *Account) ChangeOwner(newOwner string) {
	a.owner = newOwner
}

// Owner of the account
func (a Account) Owner() string {
	return a.owner
}

// Go가 자동적으로 호출해주는 method도 있음. ex) String

func (a Account) String() string {
	return fmt.Sprint(a.Owner(), "'s account. \nHas: ", a.Balance())
}
