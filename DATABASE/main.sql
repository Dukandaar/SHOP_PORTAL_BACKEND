CREATE SCHEMA shop AUTHORIZATION postgres;

-- define enums
CREATE TYPE is_active AS ENUM ('Y', 'N');
CREATE TYPE stock_type AS ENUM ('Gold', 'Silver', 'Cash');
CREATE TYPE balance_type AS ENUM ('Gold', 'Silver', 'Cash');
CREATE TYPE transaction_type AS ENUM ('Gold', 'Silver', 'Cash');
CREATE TYPE payment_type AS ENUM ('Gold', 'Silver', 'Cash', 'Upi', 'Cheque', 'NEFT', 'RTGS', 'Card', 'Other');
CREATE TYPE reason AS ENUM ('Sell', 'Buy');

-- create tables
CREATE TABLE shop.owner (
    id SERIAL,
    shop_name varchar(255) NOT NULL,
    owner_name varchar(255) NOT NULL,
    reg_id varchar(10) NOT NULL,
    phone_no varchar(10) NOT NULL,
    is_active is_active NOT NULL,
    reg_date date NOT NULL,
    address VARCHAR(255),
    remarks text NULL,
    key VARCHAR(500) NOT NULL;
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE shop.stock (
    id SERIAL,
    owner_id INTEGER NOT NULL,
    type stock_type NOT NULL,
    item_name VARCHAR(255) NOT NULL,
    tunch FLOAT,
    weight FLOAT DEFAULT 0.0,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE shop.customer (
    id SERIAL,
    name VARCHAR(255) NOT NULL,
    company_name VARCHAR(255),
    reg_id VARCHAR(10) NOT NULL,
    reg_date DATE,
    ph_no VARCHAR(10),
    address VARCHAR(255),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE shop.owner_customer (
    id SERIAL,
    owner_id INTEGER NOT NULL,
    customer_id INTEGER NOT NULL,
    is_active is_active NOT NULL,
    remark TEXT
);

CREATE TABLE shop.balance (
    id SERIAL,
    owner_id INTEGER,
    customer_id INTEGER,
    type balance_type NOT NULL,
    balance FLOAT DEFAULT 0.0,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);


CREATE TABLE shop.bill (
    id SERIAL,
    customer_id INTEGER NOT NULL,
    date DATE,
    remark TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);


CREATE TABLE shop.transaction (
    id SERIAL,
    bill_id INTEGER NOT NULL,
    type transaction_type NOT NULL,
    item_name VARCHAR(500),
    weight FLOAT,
    less FLOAT,
    net_weight FLOAT,
    tunch FLOAT,
    fine FLOAT,
    discount FLOAT,
    gold_rate FLOAT,
    amount FLOAT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE shop.payment (
    id SERIAL,
    customer_id INTEGER NOT NULL,
    bill_id INTEGER,
    type payment_type NOT NULL,
    amount FLOAT,
    date DATE NOT NULL,
    remark TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE shop.stock_transactions(
    id SERIAL,
    stock_id INTEGER NOT NULL,
    prev_balance FLOAT NOT NULL,
    new_balance FLOAT NOT NULL,
    reason reason NOT NULL,
    transaction_id INTEGER,
    remark TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Primary Key
ALTER TABLE shop.owner ADD CONSTRAINT owner_pkey PRIMARY KEY (id);
ALTER TABLE shop.stock ADD CONSTRAINT stock_pkey PRIMARY KEY (id);
ALTER TABLE shop.customer ADD CONSTRAINT customer_pkey PRIMARY KEY (id);
ALTER TABLE shop.owner_customer ADD CONSTRAINT owner_customer_pkey PRIMARY KEY (id);
ALTER TABLE shop.balance ADD CONSTRAINT balance_pkey PRIMARY KEY (id);
ALTER TABLE shop.bill ADD CONSTRAINT bill_pkey PRIMARY KEY (id);
ALTER TABLE shop.transaction ADD CONSTRAINT transaction_pkey PRIMARY KEY (id);
ALTER TABLE shop.payment ADD CONSTRAINT payment_pkey PRIMARY KEY (id);
ALTER TABLE shop.stock_transactions ADD CONSTRAINT stock_transactions_pkey PRIMARY KEY (id);

-- Foreign Key
ALTER TABLE shop.stock ADD CONSTRAINT fk_owner_id FOREIGN KEY (owner_id) REFERENCES shop.owner(id);
ALTER TABLE shop.owner_customer ADD CONSTRAINT fk_owner_customer_owner FOREIGN KEY (owner_id) REFERENCES shop.owner(id);
ALTER TABLE shop.owner_customer ADD CONSTRAINT fk_owner_customer_customer FOREIGN KEY (customer_id) REFERENCES shop.customer(id);
ALTER TABLE shop.balance ADD CONSTRAINT fk_balance_owner FOREIGN KEY (owner_id) REFERENCES shop.owner(id);
ALTER TABLE shop.balance ADD CONSTRAINT fk_balance_customer FOREIGN KEY (customer_id) REFERENCES shop.customer(id);
ALTER TABLE shop.bill ADD CONSTRAINT fk_bill_customer FOREIGN KEY (customer_id) REFERENCES shop.customer(id);
ALTER TABLE shop.transaction ADD CONSTRAINT fk_transaction_bill FOREIGN KEY (bill_id) REFERENCES shop.bill(id);
ALTER TABLE shop.payment ADD CONSTRAINT fk_payment_customer FOREIGN KEY (customer_id) REFERENCES shop.customer(id);
ALTER TABLE shop.payment ADD CONSTRAINT fk_payment_bill FOREIGN KEY (bill_id) REFERENCES shop.bill(id);
ALTER TABLE shop.stock_transactions ADD CONSTRAINT fk_stock_transactions_stock FOREIGN KEY (stock_id) REFERENCES shop.stock(id);
ALTER TABLE shop.stock_transactions ADD CONSTRAINT fk_stock_transactions_transaction FOREIGN KEY (transaction_id) REFERENCES shop.transaction(id);

-- Constraint
ALTER TABLE shop.owner ADD CONSTRAINT owner_reg_id UNIQUE (reg_id);
ALTER TABLE shop.owner ADD CONSTRAINT unique_name_ph_no UNIQUE (shop_name, owner_name, phone_no);
ALTER TABLE shop.stock ADD CONSTRAINT unique_type_item_tunch UNIQUE (type, item_name, tunch); 
ALTER TABLE shop.customer ADD CONSTRAINT unique_reg_id UNIQUE (reg_id);
ALTER TABLE shop.customer ADD CONSTRAINT unique_phone_name UNIQUE (ph_no, name);
ALTER TABLE shop.balance ADD CONSTRAINT unique_owner_type UNIQUE (owner_id, type);
ALTER TABLE shop.balance ADD CONSTRAINT unique_customer_type UNIQUE (customer_id, type);
ALTER TABLE shop.balance ADD CONSTRAINT check_either_owner_or_customer CHECK ((owner_id IS NULL AND customer_id IS NOT NULL) OR (owner_id IS NOT NULL AND customer_id IS NULL));



-- Trigger functions
CREATE OR REPLACE FUNCTION update_updated_at_column() RETURNS TRIGGER AS $$ 
BEGIN
   NEW.updated_at = now();
   RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- Create Trigger
CREATE TRIGGER update_stock_updated_at BEFORE UPDATE ON shop.owner FOR EACH ROW EXECUTE PROCEDURE update_updated_at_column();
CREATE TRIGGER update_stock_updated_at BEFORE UPDATE ON shop.stock FOR EACH ROW EXECUTE PROCEDURE update_updated_at_column();
CREATE TRIGGER update_customer_updated_at BEFORE UPDATE ON shop.customer FOR EACH ROW EXECUTE PROCEDURE update_updated_at_column();
CREATE TRIGGER update_balance_updated_at BEFORE UPDATE ON shop.balance FOR EACH ROW EXECUTE PROCEDURE update_updated_at_column();
CREATE TRIGGER update_bill_updated_at BEFORE UPDATE ON shop.bill FOR EACH ROW EXECUTE PROCEDURE update_updated_at_column();
CREATE TRIGGER update_transaction_updated_at BEFORE UPDATE ON shop.transaction FOR EACH ROW EXECUTE PROCEDURE update_updated_at_column();


-- create index
CREATE INDEX idx_owner_id ON shop.owner (id);
CREATE INDEX idx_stock_id ON shop.stock (id);
CREATE INDEX idx_customer_id ON shop.customer (id);
CREATE INDEX idx_owner_customer_id ON shop.owner_customer (id);
CREATE INDEX idx_balance_id ON shop.balance (id);
CREATE INDEX idx_bill_id ON shop.bill (id);
CREATE INDEX idx_transaction_id ON shop.transaction (id);
CREATE INDEX idx_payment_id ON shop.payment (id);
CREATE INDEX idx_stock_transactions_id ON shop.stock_transactions (id);

CREATE INDEX idx_reg_id ON shop.owner (reg_id);
CREATE INDEX idx_owner_type ON shop.balance (owner_id, type);

