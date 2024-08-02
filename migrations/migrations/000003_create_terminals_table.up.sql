CREATE TABLE IF NOT EXISTS terminals (
    id SERIAL PRIMARY KEY,
    assembly_number VARCHAR(255) UNIQUE NOT NULL,
    inn VARCHAR(12) NOT NULL,
    company_name VARCHAR(255) NOT NULL,
    address TEXT NOT NULL,
    cash_register_number VARCHAR(255) UNIQUE NOT NULL,
    module_number VARCHAR(255) NOT NULL,
    last_request_date TIMESTAMP WITH TIME ZONE,
    database_update_date TIMESTAMP WITH TIME ZONE,
    status VARCHAR(50) NOT NULL DEFAULT 'active',
    user_id INTEGER REFERENCES users(id),
    free_record_balance INTEGER NOT NULL DEFAULT 0,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

CREATE INDEX idx_terminals_assembly_number ON terminals(assembly_number);
CREATE INDEX idx_terminals_module_number ON terminals(module_number);
CREATE INDEX idx_terminals_user_id ON terminals(user_id);