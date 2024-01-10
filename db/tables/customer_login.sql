CREATE TABLE customer_logins (
    id SERIAL PRIMARY KEY,
    customer_id INT REFERENCES dim_customers(id) NOT NULL,
    username VARCHAR(50) NOT NULL UNIQUE,
    password_hash VARCHAR(255) NOT NULL,  -- Store a secure password hash
    last_login TIMESTAMP
);