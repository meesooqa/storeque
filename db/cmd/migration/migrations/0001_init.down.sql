DROP TRIGGER IF EXISTS set_timestamp_on_user_settings ON user_settings;
DROP TRIGGER IF EXISTS set_timestamp_on_roles ON roles;
DROP TRIGGER IF EXISTS set_timestamp_on_users ON users;

DROP TABLE IF EXISTS user_settings;
DROP TABLE IF EXISTS roles;
DROP TABLE IF EXISTS users;

DROP FUNCTION IF EXISTS update_updated_at_column();