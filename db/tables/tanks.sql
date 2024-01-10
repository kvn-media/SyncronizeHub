CREATE TABLE tanks (
    id SERIAL PRIMARY KEY,
    location VARCHAR(100) NOT NULL,
    capacity DECIMAL(10,2) NOT NULL,  -- Assuming capacity in liters
    fluid_type VARCHAR(50) NOT NULL
);