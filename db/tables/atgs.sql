CREATE TABLE atgs (
    id SERIAL PRIMARY KEY,
    technology VARCHAR(50) NOT NULL,
    measurement_interval INT NOT NULL,
    calibration_date DATE
);