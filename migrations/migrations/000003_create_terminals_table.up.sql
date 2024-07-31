CREATE TABLE IF NOT EXISTS terminals (
    id SERIAL PRIMARY KEY,
    inn VARCHAR(12) NOT NULL,
    company_name VARCHAR(255) NOT NULL,
    address TEXT NOT NULL,
    cash_register_number VARCHAR(255) UNIQUE NOT NULL,
    module_number VARCHAR(255) UNIQUE NOT NULL,
    assembly_number VARCHAR(255) UNIQUE NOT NULL,
    last_request_date TIMESTAMP WITH TIME ZONE,
    database_update_date TIMESTAMP WITH TIME ZONE,
    status BOOLEAN DEFAULT true,
    user_id INTEGER REFERENCES users(id),
    free_record_balance INTEGER NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

CREATE INDEX idx_terminals_user_id ON terminals(user_id);
CREATE INDEX idx_terminals_module_number ON terminals(module_number);
CREATE INDEX idx_terminals_cash_register_number ON terminals(cash_register_number);