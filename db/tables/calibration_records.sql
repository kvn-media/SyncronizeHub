CREATE TABLE calibration_records (
    id SERIAL PRIMARY KEY,
    flow_meter_id INT REFERENCES flow_meters(id),
    atg_id INT REFERENCES atgs(id),
    calibration_date DATE NOT NULL,
    technician_name VARCHAR(50),
    notes VARCHAR(255)
);