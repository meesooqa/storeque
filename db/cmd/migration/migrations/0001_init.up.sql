CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    telegram_id BIGINT NOT NULL UNIQUE,
    username VARCHAR(255),
    first_name VARCHAR(255),
    last_name VARCHAR(255)
);

CREATE TABLE roles (
    id SERIAL PRIMARY KEY,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    code VARCHAR(31) NOT NULL UNIQUE
);

INSERT INTO roles (code) VALUES
    ('admin'),
    ('editor'),
    ('manager'),
    ('customer');

CREATE TABLE user_settings (
    user_id INT PRIMARY KEY
        REFERENCES users(id)
            ON UPDATE CASCADE
            ON DELETE CASCADE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    lang  VARCHAR(15) NOT NULL DEFAULT 'en',
    role_id INT NOT NULL DEFAULT 4, -- 4: 'customer' role
    CONSTRAINT fk_user_role
        FOREIGN KEY (role_id)
            REFERENCES roles(id)
            ON UPDATE CASCADE
            ON DELETE RESTRICT
);

CREATE OR REPLACE FUNCTION update_updated_at_column()
    RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = now();
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER set_timestamp_on_users
    BEFORE UPDATE ON users
    FOR EACH ROW
EXECUTE FUNCTION update_updated_at_column();

CREATE TRIGGER set_timestamp_on_roles
    BEFORE UPDATE ON roles
    FOR EACH ROW
EXECUTE FUNCTION update_updated_at_column();

CREATE TRIGGER set_timestamp_on_user_settings
    BEFORE UPDATE ON user_settings
    FOR EACH ROW
EXECUTE FUNCTION update_updated_at_column();