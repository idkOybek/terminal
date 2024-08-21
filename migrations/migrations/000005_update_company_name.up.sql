-- Обновляем существующие NULL значения на пустую строку
UPDATE users SET company_name = '' WHERE company_name IS NULL;

-- Изменяем схему таблицы, делая company_name NOT NULL с пустой строкой по умолчанию
ALTER TABLE users 
    ALTER COLUMN company_name SET NOT NULL,
    ALTER COLUMN company_name SET DEFAULT '';