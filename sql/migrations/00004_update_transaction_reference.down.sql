-- Revert to reference_id INT column
ALTER TABLE transactions DROP COLUMN IF EXISTS reference;

ALTER TABLE transactions 
ADD COLUMN reference_id INT;
