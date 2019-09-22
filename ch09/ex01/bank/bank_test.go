// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package bank

import (
	"fmt"
	"testing"
)

func TestBank(t *testing.T) {
	done := make(chan struct{})

	// 残金が0の時引き出せない
	if Withdraw(100) {
		t.Errorf("Balance() = %d, Withdraw(100) must be false", Balance())
	}
	// 引き出しに失敗した後残高が変わらない
	if Balance() != 0 {
		t.Errorf("Balance = %d, want 0", Balance())
	}

	// Alice
	go func() {
		Deposit(200)
		fmt.Println("=", Balance())
		done <- struct{}{}
	}()

	// Bob
	go func() {
		Deposit(100)
		done <- struct{}{}
	}()

	// Wait for both transactions.
	<-done
	<-done

	if got, want := Balance(), 300; got != want {
		t.Errorf("Balance = %d, want %d", got, want)
	}

	// 残金以上は引き出せない
	if Withdraw(301) {
		t.Errorf("Balance() = %d, Withdraw(301) must be false", Balance())
	}
	// 引き出しに失敗した後残高が変わらない
	if Balance() != 300 {
		t.Errorf("Balance = %d, want 300", Balance())
	}

	// Alice
	var aliceWithdraw bool
	go func() {
		aliceWithdraw = Withdraw(200)
		fmt.Println("Alice withdraw: ", aliceWithdraw)
		done <- struct{}{}
	}()

	var bobWithdraw bool
	// Bob
	go func() {
		bobWithdraw = Withdraw(200)
		fmt.Println("Bob withdraw: ", bobWithdraw)
		done <- struct{}{}
	}()

	<-done
	<-done

	// 片方の取引のみが成功していることを期待している
	if aliceWithdraw == bobWithdraw { // <=> A ^ B = true <=> 両方取引が成功したか両方失敗したか
		t.Errorf("Either Alice or Bob's withdrwing must have failed.")
	}
	if got, want := Balance(), 100; got != want {
		t.Errorf("Balance = %d, want %d", got, want)
	}
}
