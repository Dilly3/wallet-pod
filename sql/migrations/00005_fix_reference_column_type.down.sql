-- Revert reference column changes
DO $$ BEGIN
    DROP INDEX IF EXISTS idx_transactions_reference;
EXCEPTION 
    WHEN OTHERS THEN NULL;
END $$;

DO $$ BEGIN
    ALTER TABLE transactions DROP CONSTRAINT IF EXISTS uk_transactions_reference;
EXCEPTION 
    WHEN OTHERS THEN NULL;
END $$;
