package wallet

import (
	"testing"
)

func TestWallet(t *testing.T) {
	wallet := Wallet{}

	assertBalance := func(t testing.TB, wallet Wallet, want Bitcoin) {
		if wallet.Balance() != want {
			t.Errorf("want: %s; got: %s", want, wallet.Balance())
		}
	}

	assertError := func(t testing.TB, err error, wantMessage string) {
		if err == nil {
			t.Errorf("Expected an error")
		}
		if err != nil && err.Error() != wantMessage {
			t.Errorf("Error not expected: got: [%s], expected: [%s]", err, wantMessage)
		}
	}
	t.Run("deposit", func(t *testing.T) {
		wallet.Deposit(10)
		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("withdraw", func(t *testing.T) {
		err := wallet.Withdraw(3)
		assertBalance(t, wallet, Bitcoin(7))
		if err != nil {
			t.Errorf("Error %e is unexpected", err)
		}
	})

	t.Run("withdraw too much", func(t *testing.T) {
		startingBalance := Bitcoin(10)
		wallet := Wallet{startingBalance}
		assertBalance(t, wallet, startingBalance)
		err := wallet.Withdraw(Bitcoin(11))
		assertBalance(t, wallet, startingBalance)
		assertError(t, err, "cannot withdraw 11 BTC from a balance of 10 BTC")
	})
}

func TestWalletWithdraw(t *testing.T) {
	wallet := Wallet{}
	wallet.Deposit(10)
	err := wallet.Withdraw(3)
	if err != nil {
		t.Errorf("Error %e is unexpected", err)
	}
	want := Bitcoin(7)
	got := wallet.Balance()
	if got != want {
		t.Errorf("want: %s; got: %s", want, got)
	}
}

func TestBitcoinString(t *testing.T) {
	btc := Bitcoin(10)
	got := btc.String()
	want := "10 BTC"
	if got != want {
		t.Errorf("want: %s; got: %s", want, got)
	}
}
