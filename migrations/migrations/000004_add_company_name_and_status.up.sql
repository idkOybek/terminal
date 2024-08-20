-- Add company_name column to users table
ALTER TABLE users ADD COLUMN company_name VARCHAR(255);

-- Add is_active column to fiscal_modules table
ALTER TABLE fiscal_modules ADD COLUMN is_active BOOLEAN NOT NULL DEFAULT false;

-- Update existing fiscal_modules to set is_active based on whether they have associated terminals
UPDATE fiscal_modules
SET is_active = EXISTS (
    SELECT 1
    FROM terminals
    WHERE terminals.cash_register_number = fiscal_modules.factory_number
);