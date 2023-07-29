BEGIN;

ALTER TABLE authors DROP COLUMN bio;

ALTER TABLE authors DROP COLUMN date_of_birth;

END;