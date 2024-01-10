CREATE TABLE dim_flow_meters (
    id SERIAL PRIMARY KEY,
    type VARCHAR(50) NOT NULL,
    location VARCHAR(100) NOT NULL,
    fluid_type VARCHAR(50) NOT NULL,
    calibration_date DATE,
    customer_id INT REFERENCES dim_customers(id)  -- Updated foreign key
);

CREATE TABLE dim_atgs (
    id SERIAL PRIMARY KEY,
    technology VARCHAR(50) NOT NULL,
    measurement_interval INT NOT NULL,
    calibration_date DATE
);

CREATE TABLE dim_tanks (
    id SERIAL PRIMARY KEY,
    location VARCHAR(100) NOT NULL,
    capacity DECIMAL(10,2) NOT NULL,
    fluid_type VARCHAR(50) NOT NULL
);

CREATE TABLE dim_customers (  -- Replaced dim_pipelines with customers
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    -- ... other customer attributes
);