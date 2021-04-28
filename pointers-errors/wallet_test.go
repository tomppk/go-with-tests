package main

import (
	"testing"
)

func TestWallet(t *testing.T) {

	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("Withdraw with funds", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		err := wallet.Withdraw(Bitcoin(10))
		assertBalance(t, wallet, Bitcoin(10))
		assertNoError(t, err)
	})

	// We want withdraw to return an error if you try to take out more than you have
	// and the balance stays the same. We then check that error has returned by
	// failing the test if it is "nil"
	t.Run("Withdraw insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := Wallet{startingBalance}
		err := wallet.Withdraw(Bitcoin(100))

		assertBalance(t, wallet, startingBalance)
		assertError(t, err, ErrInsufficientFunds)
	})

}

// Helper function to check balance
func assertBalance(t testing.TB, wallet Wallet, want Bitcoin) {
	t.Helper()
	got := wallet.Balance()

	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

func assertNoError(t testing.TB, got error) {
	t.Helper()
	if got != nil {
		t.Fatal("got and error but didnt want one")
	}
}

// Helper function to check error message
func assertError(t testing.TB, got error, want error) {
	t.Helper()
	if got == nil {
		// t.Fatal will stop the test if it is called. Without this the test would
		// carry on to the next step and panic because of a nil pointer.
		t.Fatal("didnt get an error but wanted one")
	}

	if got != want {
		t.Errorf(("got %q, want %q"), got, want)
	}
}
