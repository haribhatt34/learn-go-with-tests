package pointers

import "errors"

// Bitcoin add domain functionality on existing int type.
type Bitcoin int

// Wallet represents bitcoin wallet here.
type Wallet struct {
	balance Bitcoin
}

// Deposit amount into wallet
func (w *Wallet) Deposit(amount Bitcoin) {
	// Below is also correct
	// (*w).balance += amount
	w.balance += amount
}

// ErrInsufficientFunds contains error message
var ErrInsufficientFunds = errors.New("can't withdraw, insufficient funds")

// Withdraw amount from wallet
func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return ErrInsufficientFunds
	}
	w.balance -= amount
	return nil
}

// Balance return the amount in wallet
func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

/* Notes
 * 1)
 * Whey are we using balance() instead of directly accessing balance
 * using dot operator?
 * For security, we don't want to expose our inner state to rest of the
 * word, so we access data using methods to control the access.
 *
 * 2)
 * In go, if variables, types, functions, types, etc starts with a lowercase
 * symbol, then it is private outside the package it's defined in.
 *
 * 3)
 * We can access the internal 'balance' field in the struct using the
 * receiver variable .i.e 'w Wallet'.
 *
 * 4)
 * In go, calling a function/methods, passes the arguments as copy.
 *
 * 5)
 * Most lovely thing about pointers in go, we don't have to deference it
 * to get the value, although deferencing will also yield the correct ans.
 * The makers of go deemed, (*w) like notation cumbersome, rightly so.
 *
 * 6)
 * We can create new types from existing ones.
 * type Bitcoin int
 *
 * 7)
 * Above creates a new type 'Bitcoin', we can declare methods on it,
 * which adds functionality specific to this type like 'deposit', 'withdraw'.
 *
 * 8)
 * Cannot compare to errors.New() as each new call to New() method a new value
 * even if the text is same.
 * Call to errors.New() returns an error object.
 *
 * 9)
 * The 'var' keyword allows us to define values global to the package.
 */
