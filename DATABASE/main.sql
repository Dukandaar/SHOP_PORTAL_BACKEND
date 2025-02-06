
begin;

-- create schema
CREATE SCHEMA shop;

-- create table
CREATE TABLE shop.shop_owner (
    id serial4 PRIMARY KEY,
    shop_name varchar(255) NOT NULL,
    owner_name varchar(255) NOT NULL,
    reg_date date NULL,
    ph_no varchar(10) NULL,
    address text NULL
    is_active is_active NOT NULL,
    remarks text NULL,
);

CREATE TABLE shop.partner (
    id SERIAL PRIMARY KEY,
    full_name VARCHAR(255) NOT NULL,
    company_name VARCHAR(255),
    reg_date DATE,
    ph_no VARCHAR(10),
    address VARCHAR(255),
    is_active is_active NOT NULL,

    UNIQUE (ph_no, full_name)
);

CREATE TABLE shop.balance (
    id SERIAL PRIMARY KEY,
    partner_id INTEGER NOT NULL,
    gold_due FLOAT,
    silver_due FLOAT,
    cash_due FLOAT,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    FOREIGN KEY (partner_id) REFERENCES shop.partner(id)
);

CREATE TABLE shop.bill (
    id SERIAL PRIMARY KEY,
    partner_id INTEGER NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,  -- Updated on every update

    FOREIGN KEY (partner_id) REFERENCES shop.partner(id)  -- Assuming 'partner' table exists
);

CREATE TABLE shop.transaction (
    id SERIAL PRIMARY KEY,
    bill_id INTEGER NOT NULL,
    type transaction_type NOT NULL,
    item_name VARCHAR(500),
    weight FLOAT,
    less FLOAT,
    net_weight FLOAT,
    tunch FLOAT,
    fine FLOAT,
    discount FLOAT,
    amount FLOAT,

    FOREIGN KEY (bill_id) REFERENCES shop.bill(id)
);

CREATE TABLE shop.stock (
    id SERIAL PRIMARY KEY,
    partner_id INTEGER NOT NULL,
    type stock_type NOT NULL,
    item_name VARCHAR(255) NOT NULL,
    weight FLOAT,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    FOREIGN KEY (partner_id) REFERENCES shop.partner(id)
);

-- data type
CREATE TYPE transaction_type AS ENUM ('Gold', 'Silver', 'Cash');
CREATE TYPE stock_type AS ENUM ('Gold', 'Silver', 'Cash');

-- functions
CREATE OR REPLACE FUNCTION update_bill_updated_at()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = now();
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION update_stock_updated_at()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = now();
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- create trigger
-- Trigger to call the function before each update on bill table
CREATE TRIGGER bill_updated_at_trigger
BEFORE UPDATE ON shop.bill
FOR EACH ROW
EXECUTE PROCEDURE update_bill_updated_at();

-- Trigger to call the function before each update on stock table
CREATE TRIGGER stock_updated_at_trigger
BEFORE UPDATE ON shop.stock
FOR EACH ROW
EXECUTE PROCEDURE update_stock_updated_at();

-- add index
-- Add an index for faster queries on partner_id
CREATE INDEX idx_stock_partner_id ON shop.stock (partner_id);

-- Index for bill_id (important for performance)
CREATE INDEX idx_transaction_bill_id ON shop.transaction (bill_id);

commit;