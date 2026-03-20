-- Add unique composite constraint on wallet_id and reference
-- A wallet cannot have the same reference twice
ALTER TABLE transactions
ADD CONSTRAINT uk_transactions_wallet_reference UNIQUE (wallet_id, reference);
