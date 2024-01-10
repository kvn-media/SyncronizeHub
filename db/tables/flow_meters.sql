CREATE TABLE flow_meters (
    id SERIAL PRIMARY KEY,
    type VARCHAR(50) NOT NULL,
    location VARCHAR(100) NOT NULL,
    fluid_type VARCHAR(50) NOT NULL,
    calibration_date DATE,
    pipeline_id INT REFERENCES pipelines(id) -- Assuming a separate table for pipelines
);