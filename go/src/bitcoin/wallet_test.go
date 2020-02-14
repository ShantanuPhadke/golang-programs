package bitcoin

import (
	"testing"
)

func TestWallet(t *testing.T) {

	assertBalanceEqual := func(t *testing.T, wallet Wallet, expected Bitcoin) {
		t.Helper()
		actual := wallet.Balance()

		if expected != actual {
			t.Errorf("Actual balance was %s but got %s", actual, expected)
		}
	}

	assertError := func(t *testing.T, err error, expectedError error) {
		t.Helper()
		if err == nil {
			t.Fatal("expected error there was none")
		}
		// If there is an error we check if its the correct one we want
		if err.Error() != expectedError.Error() {
			t.Errorf("got %q, but wanted %q", err, expectedError.Error())
		}
	}

	assertNoError := func(t *testing.T, err error) {
		t.Helper()
		if err != nil {
			t.Fatal("Got an error when we shouldn't have")
		}
	}

	t.Run("Testing Bitcoin Wallet Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10.0))
		expected := Bitcoin(10.0)

		assertBalanceEqual(t, wallet, expected)
	})

	t.Run("Testing Bitcoin Wallet Withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(25.0)}
		err := wallet.Withdraw(Bitcoin(10.0))
		expected := Bitcoin(15.0)

		assertBalanceEqual(t, wallet, expected)
		assertNoError(t, err)
	})

	t.Run("Testing scenario of withdrawing more money than is in the wallet", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(10.0)}
		err := wallet.Withdraw(Bitcoin(20.0))

		assertBalanceEqual(t, wallet, Bitcoin(10.0))
		assertError(t, err, ErrorInsufficientFunds)
	})
}
