-- Remove the composite unique constraint
ALTER TABLE transactions
DROP CONSTRAINT IF EXISTS uk_transactions_wallet_reference;
