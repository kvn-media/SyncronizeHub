CREATE TABLE fact_measurements (
    id SERIAL PRIMARY KEY,
    flow_meter_id INT REFERENCES dim_flow_meters(id),
    atg_id INT REFERENCES dim_atgs(id),
    tank_id INT REFERENCES dim_tanks(id),
    customer_id INT REFERENCES dim_customers(id),  -- Updated foreign key
    timestamp TIMESTAMP NOT NULL,
    flow_rate DECIMAL(10,2),
    temperature DECIMAL(5,2),
    pressure DECIMAL(10,2),
    tank_level DECIMAL(5,2)
);