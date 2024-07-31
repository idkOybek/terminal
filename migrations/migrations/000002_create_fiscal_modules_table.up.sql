CREATE TABLE IF NOT EXISTS fiscal_modules (
    id SERIAL PRIMARY KEY,
    factory_number VARCHAR(255) UNIQUE NOT NULL,
    fiscal_number VARCHAR(255) UNIQUE NOT NULL,
    user_id INTEGER REFERENCES users(id),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

CREATE INDEX idx_fiscal_modules_user_id ON fiscal_modules(user_id);
CREATE INDEX idx_fiscal_modules_factory_number ON fiscal_modules(factory_number);
CREATE INDEX idx_fiscal_modules_fiscal_number ON fiscal_modules(fiscal_number);