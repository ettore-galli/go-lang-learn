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

func assertStrings(t *testing.T, got string, want string) {
	if got != want {
		t.Errorf("\nwant: [%s]; \ngot.: [%s]", want, got)
	}
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

func TestSearchOnMap(t *testing.T) {
	testMap := map[string]string{"test": "questo è un test"}
	got := SearchOnMap(testMap, "test")
	want := "questo è un test"
	if got != want {
		t.Errorf("want: %s; got: %s", want, got)
	}
	assertStrings(t, got, want)
}

func TestSearchOnDictionary(t *testing.T) {
	t.Run("found", func(t *testing.T) {
		dict := Dictionary{"test": "questo è un test"}
		got, _ := dict.Search("test")
		want := "questo è un test"

		assertStrings(t, got, want)

	})

	t.Run("not found", func(t *testing.T) {
		dict := Dictionary{"test": "questo è un test"}
		_, err := dict.Search("xxx")
		wantErr := "non trovato xxx"

		assertStrings(t, err.Error(), wantErr)
	})
}

func TestAddDictionaryEntry(t *testing.T) {
	t.Run("add - ok", func(t *testing.T) {
		dict := Dictionary{}
		dict.AddEntry("one", "The number one")
		got, _ := dict.Search("one")
		want := "The number one"

		assertStrings(t, got, want)

	})

	t.Run("add / exixts", func(t *testing.T) {
		dict := Dictionary{"one": "The number one"}
		err := dict.AddEntry("one", "The number one")
		want := "entry [one] exists"

		assertStrings(t, err.Error(), want)

	})

	t.Run("update - ok", func(t *testing.T) {
		dict := Dictionary{"one": "The number one"}
		dict.UpdateEntry("one", "nr.1")
		got, _ := dict.Search("one")
		want := "nr.1"

		assertStrings(t, got, want)

	})

	t.Run("update / not exixts", func(t *testing.T) {
		dict := Dictionary{"one": "The number one"}
		err := dict.UpdateEntry("ONE", "1")
		want := "entry [ONE] does not exist"

		assertStrings(t, err.Error(), want)

	})

}

func TestDeleteDictionary(t *testing.T) {
	t.Run("delete", func(t *testing.T) {
		dict := Dictionary{"test": "questo è un test"}
		dict.DeleteEntry("test")
		_, err := dict.Search("test")
		wantErr := "non trovato test"

		assertStrings(t, err.Error(), wantErr)

	})

}
