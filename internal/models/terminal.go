package models

import "time"

type Terminal struct {
	ID                 int       `json:"id" db:"id"`
	AssemblyNumber     string    `json:"assembly_number" db:"assembly_number"`
	INN                string    `json:"inn" db:"inn"`
	CompanyName        string    `json:"company_name" db:"company_name"`
	Address            string    `json:"address" db:"address"`
	CashRegisterNumber string    `json:"cash_register_number" db:"cash_register_number"`
	Status             bool      `json:"status" db:"status"`
	UserID             int       `json:"user_id" db:"user_id"`
	CreatedAt          time.Time `json:"created_at" db:"created_at"`
	UpdatedAt          time.Time `json:"updated_at" db:"updated_at"`
}

type TerminalCreateRequest struct {
	AssemblyNumber     string `json:"assembly_number"`
	INN                string `json:"inn"`
	CompanyName        string `json:"company_name"`
	Address            string `json:"address"`
	CashRegisterNumber string `json:"cash_register_number"`
	UserID             int    `json:"user_id"`
}

type TerminalUpdateRequest struct {
	INN                string `json:"inn"`
	CompanyName        string `json:"company_name"`
	Address            string `json:"address"`
	CashRegisterNumber string `json:"cash_register_number"`
	Status             bool   `json:"status"`
}
