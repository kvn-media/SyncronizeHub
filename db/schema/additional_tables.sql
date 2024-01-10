CREATE TABLE dim_alarms_events (
    id SERIAL PRIMARY KEY,
    event_type VARCHAR(50) NOT NULL,
    description VARCHAR(255) NOT NULL
);

CREATE TABLE dim_calibration_records (
    id SERIAL PRIMARY KEY,
    technician_name VARCHAR(50),
    notes VARCHAR(255)
);

CREATE TABLE dim_maintenance_records (
    id SERIAL PRIMARY KEY,
    work_performed VARCHAR(255) NOT NULL,
    technician_name VARCHAR(50)
);