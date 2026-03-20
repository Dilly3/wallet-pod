-- Ensure reference column is VARCHAR(255) and add proper constraints
DO $$ BEGIN
    ALTER TABLE transactions 
        ALTER COLUMN reference TYPE VARCHAR(255);
EXCEPTION 
    WHEN OTHERS THEN NULL;
END $$;

-- Remove old integer constraint if it exists, then add UNIQUE
DO $$ BEGIN
    ALTER TABLE transactions DROP CONSTRAINT uk_transactions_reference;
EXCEPTION 
    WHEN OTHERS THEN NULL;
END $$;

DO $$ BEGIN
    ALTER TABLE transactions 
        ADD CONSTRAINT uk_transactions_reference UNIQUE (reference);
EXCEPTION 
    WHEN OTHERS THEN NULL;
END $$;

-- Create index on reference column
CREATE INDEX IF NOT EXISTS idx_transactions_reference ON transactions(reference);
