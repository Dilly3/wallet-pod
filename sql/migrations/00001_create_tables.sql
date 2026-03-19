-- sql/create_tables.sql

CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    email TEXT NOT NULL UNIQUE,
    full_name TEXT NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS wallets (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    balance NUMERIC(12,2) NOT NULL DEFAULT 0.00,
    currency VARCHAR(10) NOT NULL DEFAULT 'USD',
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    CONSTRAINT fk_wallet_user FOREIGN KEY (user_id)
        REFERENCES users(id) ON DELETE CASCADE
);

-- Create transaction type enum using PL/pgSQL anonymous block for safe, idempotent creation
-- PostgreSQL doesn't support "CREATE TYPE IF NOT EXISTS", so we use DO $$ BEGIN ... END $$ 
-- to wrap the creation in error handling that allows the migration to run multiple times
DO $$ BEGIN
    -- Define an ENUM type with the three allowed transaction statuses
    CREATE TYPE transaction_type AS ENUM ('deposit', 'withdrawal', 'transfer');
EXCEPTION
    -- If the type already exists from a previous run, catch the duplicate_object error
    -- and do nothing (null;), allowing the migration to complete successfully
    WHEN duplicate_object THEN null;
END $$;

CREATE TABLE IF NOT EXISTS transactions (
    id SERIAL PRIMARY KEY,
    wallet_id INT NOT NULL,
    amount NUMERIC(12,2) NOT NULL CHECK (amount > 0),
    txn_type transaction_type NOT NULL,
    reference_id INT,
    description TEXT,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    CONSTRAINT fk_txn_wallet FOREIGN KEY (wallet_id)
        REFERENCES wallets(id) ON DELETE CASCADE
);
