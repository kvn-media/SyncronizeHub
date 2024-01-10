CREATE TABLE maintenance_records (
    id SERIAL PRIMARY KEY,
    flow_meter_id INT REFERENCES flow_meters(id),
    atg_id INT REFERENCES atgs(id),
    maintenance_date DATE NOT NULL,
    work_performed VARCHAR(255) NOT NULL,
    technician_name VARCHAR(50)
);