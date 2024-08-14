package models

import "time"

type Terminal struct {
	ID                 int       `json:"id" db:"id"`
	AssemblyNumber     string    `json:"assembly_number" db:"assembly_number"`
	INN                string    `json:"inn" db:"inn"`
	CompanyName        string    `json:"company_name" db:"company_name"`
	Address            string    `json:"address" db:"address"`
	CashRegisterNumber string    `json:"cash_register_number" db:"cash_register_number"`
	ModuleNumber       string    `json:"module_number" db:"module_number"`
	LastRequestDate    time.Time `json:"last_request_date" db:"last_request_date"`
	DatabaseUpdateDate time.Time `json:"database_update_date" db:"database_update_date"`
	IsActive           bool      `json:"is_active" db:"is_active"`
	UserID             int       `json:"user_id" db:"user_id"`
	FreeRecordBalance  int       `json:"free_record_balance" db:"free_record_balance"`
	CreatedAt          time.Time `json:"created_at" db:"created_at"`
	UpdatedAt          time.Time `json:"updated_at" db:"updated_at"`
}

type TerminalCreateRequest struct {
	AssemblyNumber     string `json:"assembly_number"`
	INN                string `json:"inn"`
	CompanyName        string `json:"company_name"`
	Address            string `json:"address"`
	CashRegisterNumber string `json:"cash_register_number"`
	ModuleNumber       string `json:"module_number"`
	LastRequestDate    string `json:"last_request_date"`
	DatabaseUpdateDate string `json:"database_update_date"`
	IsActive           bool   `json:"is_active"` 
	UserID             int    `json:"user_id"`
	FreeRecordBalance  int    `json:"free_record_balance"`
}

type TerminalUpdateRequest struct {
	AssemblyNumber     string `json:"assembly_number"`
	INN                string `json:"inn"`
	CompanyName        string `json:"company_name"`
	Address            string `json:"address"`
	CashRegisterNumber string `json:"cash_register_number"`
	ModuleNumber       string `json:"module_number"`
	LastRequestDate    string `json:"last_request_date"`
	DatabaseUpdateDate string `json:"database_update_date"`
	IsActive           bool   `json:"is_active"`
	UserID             int    `json:"user_id"`
	FreeRecordBalance  int    `json:"free_record_balance"`
}