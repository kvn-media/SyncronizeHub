CREATE TABLE level_access (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL UNIQUE,  -- Name of the access level
    permissions VARCHAR(255) NOT NULL  -- Permissions associated with the level
);

CREATE TABLE user_level_access (
    user_id INT NOT NULL,  -- References either customer_logins.id or admin_logins.id
    level_id INT REFERENCES level_access(id),
    PRIMARY KEY (user_id, level_id)
);