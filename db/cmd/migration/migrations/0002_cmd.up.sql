CREATE TABLE commands (
    id SERIAL PRIMARY KEY,
    created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT now(),
    code VARCHAR(31) NOT NULL UNIQUE
);

CREATE TABLE role_commands (
    role_id INT REFERENCES roles(id),
    command_id INT REFERENCES commands(id),
    PRIMARY KEY (role_id, command_id)
);
