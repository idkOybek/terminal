// internal/models/fiscal_module.go

package models

import "time"

type FiscalModule struct {
	ID            int       `json:"id" db:"id"`
	FactoryNumber string    `json:"factory_number" db:"factory_number"`
	FiscalNumber  string    `json:"fiscal_number" db:"fiscal_number"`
	UserID        int       `json:"user_id" db:"user_id"`
	CreatedAt     time.Time `json:"created_at" db:"created_at"`
	UpdatedAt     time.Time `json:"updated_at" db:"updated_at"`
}

type FiscalModuleCreateRequest struct {
	FactoryNumber string `json:"factory_number"`
	FiscalNumber  string `json:"fiscal_number"`
	UserID        int    `json:"user_id"`
}

type FiscalModuleUpdateRequest struct {
	FactoryNumber string `json:"factory_number"`
	FiscalNumber  string `json:"fiscal_number"`
	UserID        int    `json:"user_id"`
}
