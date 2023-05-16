package main

import (
	"errors"
	"q11_answer/bank"
)

type bankService struct {
	users map[string]bank.User
}

var _ bank.BankService = (*bankService)(nil)
var _ bank.UserService = (*bankService)(nil)

func (s *bankService) Login(username, password string) (bank.User, error) {
	u, ok := s.users[username]
	if !ok || u.Password != password {
		return bank.User{}, errors.New("invalid username or password")
	}

	return u, nil
}

func (s *bankService) Deposit(user bank.User, amount int) error {
	if amount < 0 {
		return errors.New("amount must be positive")
	}

	u, ok := s.users[user.Username]
	if !ok {
		return errors.New("user not found")
	}

	u.Balance += amount
	s.users[user.Username] = u
	return nil
}

func (s *bankService) Withdraw(user bank.User, amount int) error {
	if amount < 0 {
		return errors.New("amount must be positive")
	}

	u, ok := s.users[user.Username]
	if !ok {
		return errors.New("user not found")
	}

	u.Balance -= amount
	s.users[user.Username] = u
	return nil
}
