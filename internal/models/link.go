// internal/models/link.go

package models

import "time"

type Link struct {
	ID            int       `json:"id" db:"id"`
	FiscalNumber  string    `json:"fiscal_number" db:"fiscal_number"`
	FactoryNumber string    `json:"factory_number" db:"factory_number"`
	CreatedAt     time.Time `json:"created_at" db:"created_at"`
}

type LinkCreateRequest struct {
	FiscalNumber  string `json:"fiscal_number"`
	FactoryNumber string `json:"factory_number"`
}
