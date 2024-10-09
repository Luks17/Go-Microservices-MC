// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package sqlc

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
	"time"
)

type Currencies string

const (
	CurrenciesUSD Currencies = "USD"
	CurrenciesEUR Currencies = "EUR"
	CurrenciesBRL Currencies = "BRL"
)

func (e *Currencies) Scan(src interface{}) error {
	switch s := src.(type) {
	case []byte:
		*e = Currencies(s)
	case string:
		*e = Currencies(s)
	default:
		return fmt.Errorf("unsupported scan type for Currencies: %T", src)
	}
	return nil
}

type NullCurrencies struct {
	Currencies Currencies `json:"currencies"`
	Valid      bool       `json:"valid"` // Valid is true if Currencies is not NULL
}

// Scan implements the Scanner interface.
func (ns *NullCurrencies) Scan(value interface{}) error {
	if value == nil {
		ns.Currencies, ns.Valid = "", false
		return nil
	}
	ns.Valid = true
	return ns.Currencies.Scan(value)
}

// Value implements the driver Valuer interface.
func (ns NullCurrencies) Value() (driver.Value, error) {
	if !ns.Valid {
		return nil, nil
	}
	return string(ns.Currencies), nil
}

func (e Currencies) Valid() bool {
	switch e {
	case CurrenciesUSD,
		CurrenciesEUR,
		CurrenciesBRL:
		return true
	}
	return false
}

type Account struct {
	ID        int64      `json:"id"`
	Owner     string     `json:"owner"`
	Balance   string     `json:"balance"`
	Currency  Currencies `json:"currency"`
	CreatedAt time.Time  `json:"created_at"`
}

type Entry struct {
	ID        int64     `json:"id"`
	AccountID int64     `json:"account_id"`
	Amount    string    `json:"amount"`
	CreatedAt time.Time `json:"created_at"`
}

type Transfer struct {
	ID            int64 `json:"id"`
	FromAccountID int64 `json:"from_account_id"`
	ToAccountID   int64 `json:"to_account_id"`
	// cannot be negative
	Amount    string    `json:"amount"`
	CreatedAt time.Time `json:"created_at"`
}

type User struct {
	Username              string       `json:"username"`
	Password              string       `json:"password"`
	FullName              string       `json:"full_name"`
	Email                 string       `json:"email"`
	PasswordLastChangedAt sql.NullTime `json:"password_last_changed_at"`
	CreatedAt             time.Time    `json:"created_at"`
}
