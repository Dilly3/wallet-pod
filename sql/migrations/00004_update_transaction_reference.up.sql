-- Replace reference_id with reference VARCHAR column and add UNIQUE constraint
ALTER TABLE transactions DROP COLUMN IF EXISTS reference_id;

ALTER TABLE transactions 
ADD COLUMN IF NOT EXISTS reference VARCHAR(255) UNIQUE;
