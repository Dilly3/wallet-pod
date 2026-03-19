-- Update timestamps to BIGINT (Unix seconds) and add soft delete columns

-- Update users table
DO $$ BEGIN
    ALTER TABLE users 
        ALTER COLUMN created_at DROP DEFAULT;
EXCEPTION 
    WHEN OTHERS THEN NULL;
END $$;

DO $$ BEGIN
    ALTER TABLE users 
        ALTER COLUMN created_at TYPE BIGINT USING EXTRACT(EPOCH FROM created_at)::BIGINT;
EXCEPTION 
    WHEN OTHERS THEN NULL;
END $$;

DO $$ BEGIN
    ALTER TABLE users 
        ALTER COLUMN created_at SET DEFAULT EXTRACT(EPOCH FROM NOW())::BIGINT;
EXCEPTION 
    WHEN OTHERS THEN NULL;
END $$;

DO $$ BEGIN
    ALTER TABLE users 
        ADD COLUMN updated_at BIGINT NOT NULL DEFAULT EXTRACT(EPOCH FROM NOW())::BIGINT;
EXCEPTION 
    WHEN duplicate_column THEN NULL;
END $$;

DO $$ BEGIN
    ALTER TABLE users 
        ADD COLUMN deleted_at BIGINT;
EXCEPTION 
    WHEN duplicate_column THEN NULL;
END $$;

-- Update wallets table
DO $$ BEGIN
    ALTER TABLE wallets 
        ALTER COLUMN created_at DROP DEFAULT;
EXCEPTION 
    WHEN OTHERS THEN NULL;
END $$;

DO $$ BEGIN
    ALTER TABLE wallets 
        ALTER COLUMN created_at TYPE BIGINT USING EXTRACT(EPOCH FROM created_at)::BIGINT;
EXCEPTION 
    WHEN OTHERS THEN NULL;
END $$;

DO $$ BEGIN
    ALTER TABLE wallets 
        ALTER COLUMN created_at SET DEFAULT EXTRACT(EPOCH FROM NOW())::BIGINT;
EXCEPTION 
    WHEN OTHERS THEN NULL;
END $$;

DO $$ BEGIN
    ALTER TABLE wallets 
        ADD COLUMN updated_at BIGINT NOT NULL DEFAULT EXTRACT(EPOCH FROM NOW())::BIGINT;
EXCEPTION 
    WHEN duplicate_column THEN NULL;
END $$;

DO $$ BEGIN
    ALTER TABLE wallets 
        ADD COLUMN deleted_at BIGINT;
EXCEPTION 
    WHEN duplicate_column THEN NULL;
END $$;

-- Update transactions table
DO $$ BEGIN
    ALTER TABLE transactions 
        ALTER COLUMN created_at DROP DEFAULT;
EXCEPTION 
    WHEN OTHERS THEN NULL;
END $$;

DO $$ BEGIN
    ALTER TABLE transactions 
        ALTER COLUMN created_at TYPE BIGINT USING EXTRACT(EPOCH FROM created_at)::BIGINT;
EXCEPTION 
    WHEN OTHERS THEN NULL;
END $$;

DO $$ BEGIN
    ALTER TABLE transactions 
        ALTER COLUMN created_at SET DEFAULT EXTRACT(EPOCH FROM NOW())::BIGINT;
EXCEPTION 
    WHEN OTHERS THEN NULL;
END $$;

DO $$ BEGIN
    ALTER TABLE transactions 
        ADD COLUMN updated_at BIGINT NOT NULL DEFAULT EXTRACT(EPOCH FROM NOW())::BIGINT;
EXCEPTION 
    WHEN duplicate_column THEN NULL;
END $$;

DO $$ BEGIN
    ALTER TABLE transactions 
        ADD COLUMN deleted_at BIGINT;
EXCEPTION 
    WHEN duplicate_column THEN NULL;
END $$;
