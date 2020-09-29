package pointers

// Bitcoin add domain functionality on existing int type.
type Bitcoin int

// Wallet represents bitcoin wallet here.
type Wallet struct {
	balance Bitcoin
}

// Deposit the amount into wallet
func (w *Wallet) Deposit(amount Bitcoin) {
	// Below is also correct
	// (*w).balance += amount
	w.balance += amount
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
 */
