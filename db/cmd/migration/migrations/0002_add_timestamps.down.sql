DROP TRIGGER IF EXISTS set_timestamp_on_roles ON roles;
DROP TRIGGER IF EXISTS set_timestamp_on_users ON users;
DROP TRIGGER IF EXISTS set_timestamp_on_user_roles ON user_roles;

DROP FUNCTION IF EXISTS update_updated_at_column();

ALTER TABLE roles DROP COLUMN IF EXISTS created_at, DROP COLUMN IF EXISTS updated_at;
ALTER TABLE users DROP COLUMN IF EXISTS created_at, DROP COLUMN IF EXISTS updated_at;
ALTER TABLE user_roles DROP COLUMN IF EXISTS created_at, DROP COLUMN IF EXISTS updated_at;