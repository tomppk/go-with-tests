package main

import (
	"errors"
	"fmt"
)

// Stringer interface defined in fmt package. Lets you define how your type
// is printed when used with the %s format string in prints.
type Stringer interface {
	String() string
}

// Go lets you create new types from existing ones.
// Syntax is type MyName OriginalType
type Bitcoin int

// Custom string representation when printing type Bitcoin
func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

// In Go if a symbol (so variables, types, functions etc.) starts with a lowercase
// symbol then it is private outside the package it is defined in.
// We want our methods to be able to manipulate balance state but nothing else.
// We can access the internal balance field in the struct using the "receiver"
// variable "w" in this case
type Wallet struct {
	balance Bitcoin
}

// In Go, when you call a function or a method the arguments are copied.
// So "w" is a copy of whatever we called the method from.
// Rather than taking a copy of the Wallet, we take a pointer to the wallet
// so we can change it.
// The difference is the receiver type is *Wallet rather than Wallet, which
// can be read as "a pointer to a wallet".
// This refers to the original Wallet instance address in memory rather than a copy
func (w *Wallet) Deposit(amount Bitcoin) {
	fmt.Printf("address of balance in Deposit is %v \n", &w.balance)
	w.balance += amount
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

// In Go, errors are values. We define error variable to have a single source of
// truth for it.
// The "var" keyword allows us to define values global to the package.
var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient funds")

// errors.New creates a new error with a custom message
func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return ErrInsufficientFunds
	}
	w.balance -= amount
	return nil
}
