CREATE TABLE measurements (
    id SERIAL PRIMARY KEY,
    flow_meter_id INT REFERENCES flow_meters(id),
    atg_id INT REFERENCES atgs(id),
    tank_id INT REFERENCES tanks(id),
    timestamp TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    flow_rate DECIMAL(10,2),  -- Assuming flow rate in liters per minute
    temperature DECIMAL(5,2), -- Assuming temperature in degrees Celsius
    pressure DECIMAL(10,2),  -- Assuming pressure in Pascals
    tank_level DECIMAL(5,2)   -- Assuming tank level as a percentage
);