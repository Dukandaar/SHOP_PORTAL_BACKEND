-- Schema
CREATE SCHEMA shop AUTHORIZATION postgres;
-- Enums
CREATE TYPE is_active AS ENUM ('Y', 'N');
CREATE TYPE stock_type AS ENUM ('Gold', 'Silver', 'Cash');
CREATE TYPE bill_type AS ENUM ('WholeSale', 'Retail');
CREATE TYPE metal_type AS ENUM ('Gold', 'Silver');
CREATE TYPE payment_type AS ENUM ('Gold', 'Silver', 'Cash', 'UPI', 'Cheque', 'NEFT', 'RTGS', 'Card', 'Other');
CREATE TYPE entry_factor AS ENUM ('Fine', 'Amount');
CREATE TYPE reason AS ENUM ('Sell', 'Buy');



-- Tables
CREATE TABLE shop.owner (
    id SERIAL,
    shop_name VARCHAR(255) NOT NULL,
    owner_name VARCHAR(255) NOT NULL,
    reg_id VARCHAR(10) NOT NULL,
    gst_in VARCHAR(15) DEFAULT NULL,
    phone_no VARCHAR(10) NOT NULL,
    is_active is_active NOT NULL,
    reg_date DATE NOT NULL,
    address VARCHAR(255),
    key VARCHAR(255) NOT NULL,
    remarks TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE shop.stock (
    id SERIAL,
    owner_id BIGINT NOT NULL,
    type stock_type NOT NULL,
    item_name VARCHAR(255) NOT NULL,
    tunch FLOAT,
    weight FLOAT DEFAULT 0.0,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE shop.customer (
    id SERIAL,
    owner_id BIGINT NOT NULL,
    shop_name VARCHAR(255),
    name VARCHAR(255) NOT NULL,
    reg_id VARCHAR(12) NOT NULL,
    gst_in VARCHAR(15) DEFAULT NULL,
    reg_date DATE,
    phone_no VARCHAR(10),
    is_active is_active NOT NULL,
    address VARCHAR(255),
    remarks TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE shop.balance (
    id SERIAL,
    owner_id BIGINT,
    customer_id BIGINT,
    gold FLOAT DEFAULT 0.0,
    silver FLOAT DEFAULT 0.0,
    cash FLOAT DEFAULT 0.0,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE shop.bill (
    id SERIAL,
    bill_no BIGINT NOT NULL,
    customer_id BIGINT NOT NULL,
    type bill_type NOT NULL,
    metal metal_type NOT NULL,
    metal_rate FLOAT NOT NULL,
    date DATE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE shop.owner_bill_count (
    id SERIAL,
    owner_id BIGINT NOT NULL,
    bill_cnt INT DEFAULT 0,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE shop.transaction (
    id SERIAL,
    bill_id BIGINT NOT NULL,
    is_active is_active NOT NULL,
    item_name VARCHAR(255) NOT NULL,
    weight FLOAT NOT NULL,
    less FLOAT NOT NULL,
    net_weight FLOAT NOT NULL,
    tunch FLOAT NOT NULL,
    fine FLOAT NOT NULL,
    discount FLOAT DEFAULT 0.0,
    amount FLOAT
);

CREATE TABLE shop.payment (
    id SERIAL,
    bill_id BIGINT,
    customer_id BIGINT,
    factor entry_factor NOT NULL,
    new FLOAT DEFAULT 0.0,
    prev FLOAT DEFAULT 0.0,
    total FLOAT DEFAULT 0.0,
    paid FLOAT DEFAULT 0.0,
    rem FLOAT DEFAULT 0.0,
    type payment_type NOT NULL,
    date DATE NOT NULL,
    remarks TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE shop.stock_history(
    id SERIAL,
    stock_id BIGINT NOT NULL,
    prev_balance FLOAT NOT NULL,
    new_balance FLOAT NOT NULL,
    reason reason NOT NULL,
    transaction_id BIGINT,
    remarks TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Primary Key
ALTER TABLE shop.owner ADD CONSTRAINT owner_pkey PRIMARY KEY (id);
ALTER TABLE shop.stock ADD CONSTRAINT stock_pkey PRIMARY KEY (id);
ALTER TABLE shop.customer ADD CONSTRAINT customer_pkey PRIMARY KEY (id);
ALTER TABLE shop.balance ADD CONSTRAINT balance_pkey PRIMARY KEY (id);
ALTER TABLE shop.bill ADD CONSTRAINT bill_pkey PRIMARY KEY (id);
ALTER TABLE shop.transaction ADD CONSTRAINT transaction_pkey PRIMARY KEY (id);
ALTER TABLE shop.payment ADD CONSTRAINT payment_pkey PRIMARY KEY (id);
ALTER TABLE shop.stock_history ADD CONSTRAINT stock_history_pkey PRIMARY KEY (id);
ALTER TABLE owner_bill_count ADD CONSTRAINT owner_bill_count_pkey PRIMARY KEY (id);



-- Foreign Key
ALTER TABLE shop.stock ADD CONSTRAINT fk_owner_id FOREIGN KEY (owner_id) REFERENCES shop.owner(id);
ALTER TABLE shop.customer ADD CONSTRAINT fk_owner_id FOREIGN KEY (owner_id) REFERENCES shop.owner(id);
ALTER TABLE shop.balance ADD CONSTRAINT fk_balance_owner FOREIGN KEY (owner_id) REFERENCES shop.owner(id);
ALTER TABLE shop.balance ADD CONSTRAINT fk_balance_customer FOREIGN KEY (customer_id) REFERENCES shop.customer(id);
ALTER TABLE shop.bill ADD CONSTRAINT fk_bill_customer FOREIGN KEY (customer_id) REFERENCES shop.customer(id);
ALTER TABLE shop.transaction ADD CONSTRAINT fk_transaction_bill FOREIGN KEY (bill_id) REFERENCES shop.bill(id);
ALTER TABLE shop.payment ADD CONSTRAINT fk_payment_bill FOREIGN KEY (bill_id) REFERENCES shop.bill(id); 
ALTER TABLE shop.payment ADD CONSTRAINT fk_payment_customer_id FOREIGN KEY (customer_id) REFERENCES shop.customer(id); 
ALTER TABLE shop.stock_history ADD CONSTRAINT fk_stock_history_stock FOREIGN KEY (stock_id) REFERENCES shop.stock(id);
ALTER TABLE shop.stock_history ADD CONSTRAINT fk_stock_history_transaction_id FOREIGN KEY (transaction_id) REFERENCES shop.transaction(id);
ALTER TABLE shop.owner_bill_count ADD CONSTRAINT fk_owner_bill_count FOREIGN KEY (owner_id) REFERENCES shop.owner(id);

-- Sequences
CREATE SEQUENCE IF NOT EXISTS shop.owner_id_seq;
ALTER TABLE shop.owner ALTER COLUMN id SET DEFAULT nextval('shop.owner_id_seq');
ALTER SEQUENCE shop.owner_id_seq OWNED BY shop.owner.id;

CREATE SEQUENCE IF NOT EXISTS shop.stock_id_seq;
ALTER TABLE shop.stock ALTER COLUMN id SET DEFAULT nextval('shop.stock_id_seq');
ALTER SEQUENCE shop.stock_id_seq OWNED BY shop.stock.id;

CREATE SEQUENCE IF NOT EXISTS shop.customer_id_seq;
ALTER TABLE shop.customer ALTER COLUMN id SET DEFAULT nextval('shop.customer_id_seq');
ALTER SEQUENCE shop.customer_id_seq OWNED BY shop.customer.id;

CREATE SEQUENCE IF NOT EXISTS shop.balance_id_seq;
ALTER TABLE shop.balance ALTER COLUMN id SET DEFAULT nextval('shop.balance_id_seq');
ALTER SEQUENCE shop.balance_id_seq OWNED BY shop.balance.id;

CREATE SEQUENCE IF NOT EXISTS shop.bill_id_seq;
ALTER TABLE shop.bill ALTER COLUMN id SET DEFAULT nextval('shop.bill_id_seq');
ALTER SEQUENCE shop.bill_id_seq OWNED BY shop.bill.id;

CREATE SEQUENCE IF NOT EXISTS shop.transaction_id_seq;
ALTER TABLE shop.transaction ALTER COLUMN id SET DEFAULT nextval('shop.transaction_id_seq');
ALTER SEQUENCE shop.transaction_id_seq OWNED BY shop.transaction.id;

CREATE SEQUENCE IF NOT EXISTS shop.payment_id_seq;
ALTER TABLE shop.payment ALTER COLUMN id SET DEFAULT nextval('shop.payment_id_seq');
ALTER SEQUENCE shop.payment_id_seq OWNED BY shop.payment.id;

CREATE SEQUENCE IF NOT EXISTS shop.stock_history_id_seq;
ALTER TABLE shop.stock_history ALTER COLUMN id SET DEFAULT nextval('shop.stock_history_id_seq');
ALTER SEQUENCE shop.stock_history_id_seq OWNED BY shop.stock_history.id;

CREATE SEQUENCE IF NOT EXISTS shop.owner_bill_count_id_seq;
ALTER TABLE shop.owner_bill_count ALTER COLUMN id SET DEFAULT nextval('shop.owner_bill_count_id_seq');
ALTER SEQUENCE shop.owner_bill_count_id_seq OWNED BY shop.owner_bill_count.id;

-- Indexes
CREATE INDEX idx_owner_reg_id ON shop.owner (reg_id);
CREATE INDEX idx_owner_phone_no ON shop.owner (phone_no);
CREATE INDEX idx_stock_owner_id ON shop.stock (owner_id);
CREATE INDEX idx_customer_owner_id ON shop.customer (owner_id);
CREATE INDEX idx_customer_reg_id ON shop.customer (reg_id);
CREATE INDEX idx_balance_owner_id ON shop.balance (owner_id);
CREATE INDEX idx_balance_customer_id ON shop.balance (customer_id);
CREATE INDEX idx_bill_customer_id ON shop.bill (customer_id);
CREATE INDEX idx_transaction_bill_id ON shop.transaction (bill_id);
CREATE INDEX idx_payment_bill_id ON shop.payment (bill_id);
CREATE INDEX idx_payment_customer_id ON shop.payment (customer_id);
CREATE INDEX idx_stock_history_stock_id ON shop.stock_history (stock_id);
CREATE INDEX idx_stock_history_transaction_id ON shop.stock_history (transaction_id);
CREATE INDEX idx_owner_bill_count_owner_id ON owner_bill_count (owner_id);

-- Constraint
ALTER TABLE shop.owner ADD CONSTRAINT owner_reg_id UNIQUE (reg_id);
ALTER TABLE shop.owner ADD CONSTRAINT unique_name_ph_no UNIQUE (shop_name, owner_name, phone_no);
ALTER TABLE shop.stock ADD CONSTRAINT unique_owner_type_item_tunch UNIQUE (owner_id, type, item_name, tunch);
ALTER TABLE shop.customer ADD CONSTRAINT unique_reg_id UNIQUE (reg_id);
ALTER TABLE shop.customer ADD CONSTRAINT unique_name_ph_no_oId UNIQUE (shop_name, name, phone_no, owner_id);
ALTER TABLE shop.balance ADD CONSTRAINT check_either_owner_or_customer CHECK ((owner_id IS NULL AND customer_id IS NOT NULL) OR (owner_id IS NOT NULL AND customer_id IS NULL));



--Trigger function
CREATE OR REPLACE FUNCTION update_updated_at_column() RETURNS TRIGGER AS $$ 
BEGIN
   NEW.updated_at = now();
   RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- Trigger on table
CREATE TRIGGER update_stock_updated_at BEFORE UPDATE ON shop.owner FOR EACH ROW EXECUTE PROCEDURE update_updated_at_column();
CREATE TRIGGER update_stock_updated_at BEFORE UPDATE ON shop.stock FOR EACH ROW EXECUTE PROCEDURE update_updated_at_column();
CREATE TRIGGER update_customer_updated_at BEFORE UPDATE ON shop.customer FOR EACH ROW EXECUTE PROCEDURE update_updated_at_column();
CREATE TRIGGER update_balance_updated_at BEFORE UPDATE ON shop.balance FOR EACH ROW EXECUTE PROCEDURE update_updated_at_column();
CREATE TRIGGER update_bill_updated_at BEFORE UPDATE ON shop.bill FOR EACH ROW EXECUTE PROCEDURE update_updated_at_column();
CREATE TRIGGER update_transaction_updated_at BEFORE UPDATE ON shop.transaction FOR EACH ROW EXECUTE PROCEDURE update_updated_at_column();
CREATE TRIGGER update_payment_updated_at BEFORE UPDATE ON shop.payment FOR EACH ROW EXECUTE PROCEDURE update_updated_at_column();
CREATE TRIGGER update_owner_bill_count_updated_at BEFORE UPDATE ON owner_bill_count FOR EACH ROW EXECUTE PROCEDURE update_updated_at_column();