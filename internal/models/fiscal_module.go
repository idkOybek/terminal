package models

import "time"

type FiscalModule struct {
	ID            int       `json:"id" db:"id"`
	FiscalNumber  string    `json:"fiscal_number" db:"fiscal_number"`
	FactoryNumber string    `json:"factory_number" db:"factory_number"`
	UserID        int       `json:"user_id" db:"user_id"`
	CreatedAt     time.Time `json:"created_at" db:"created_at"`
	UpdatedAt     time.Time `json:"updated_at" db:"updated_at"`
}

type FiscalModuleCreateRequest struct {
	FiscalNumber  string `json:"fiscal_number"`
	FactoryNumber string `json:"factory_number"`
	UserID        int    `json:"user_id"`
}

type FiscalModuleUpdateRequest struct {
	FiscalNumber  *string `json:"fiscal_number,omitempty"`
	FactoryNumber *string `json:"factory_number,omitempty"`
	UserID        *int    `json:"user_id,omitempty"`
}

type FiscalModuleResponse struct {
	ID            int    `json:"id"`
	FiscalNumber  string `json:"fiscal_number"`
	FactoryNumber string `json:"factory_number"`
	UserID        int    `json:"user_id"`
}
