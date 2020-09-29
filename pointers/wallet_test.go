package pointers

import (
	"fmt"
	"testing"
)

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

func TestWallet(t *testing.T) {

	assertBalance := func(t *testing.T, wallet Wallet, want Bitcoin) {
		t.Helper()
		balance := wallet.Balance()
		if balance != want {
			t.Errorf("want %s got %s", want, balance)
		}
	}

	assertError := func(t *testing.T, got error, want error) {
		t.Helper()
		if got == nil {
			t.Fatal("should return error, insufficient funds")
		}
		if got != want {
			t.Errorf("want %v got %v", got, want)
		}
	}

	assertNoError := func(t *testing.T, got error) {
		t.Helper()
		if got != nil {
			t.Fatal("should not return error")
		}
	}

	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		err := wallet.Withdraw(Bitcoin(10))
		assertNoError(t, err)
		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("Withdraw insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(10)
		wallet := Wallet{balance: startingBalance}
		err := wallet.Withdraw(Bitcoin(20))
		assertError(t, err, ErrInsufficientFunds)
		assertBalance(t, wallet, Bitcoin(10))
	})
}

// Notes:
/*
* 1)
 * type Stringer interface {
 *	 String() string
 * }
 *
 * the above interface defines how %s is interpreted with strings in print.
 * we can implement String() to be used with our specific types.
 *
 * 2)
 * Return type of withdraw is error which is an interface so it can be nill.
 * All interfaces are nillable.
 * When a nil value is accessed, it throws 'runtime panic'.
 *
 * 3) Use errorVar.Error() method when comparing with string, as it returns
 * the error as string.
 *
 * 4) t.Fatal will stop the test there only, and code after it won't be
 * executed.
 *
 * 5)
 * For unchecked erros use install errcheck
 * go get -u github.com/kisielk/errcheck
 * to run use, go to source folder-
 * errcheck .
 * for ex - it return below for the plain withdraw function when
 * assertNoError is not used.
 * wallet_test.go:48:18:   wallet.Withdraw(Bitcoin(10))
 * it means, we have not checked the error returned by Withdraw.
*/
