CREATE TABLE IF NOT EXISTS links (
    id SERIAL PRIMARY KEY,
    fiscal_number VARCHAR(255) NOT NULL,
    factory_number VARCHAR(255) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    UNIQUE(fiscal_number, factory_number)
);

CREATE INDEX idx_links_fiscal_number ON links(fiscal_number);
CREATE INDEX idx_links_factory_number ON links(factory_number);