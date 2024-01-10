CREATE TABLE alarms_events (
    id SERIAL PRIMARY KEY,
    timestamp TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    flow_meter_id INT REFERENCES flow_meters(id),
    atg_id INT REFERENCES atgs(id),
    event_type VARCHAR(50) NOT NULL,
    description VARCHAR(255) NOT NULL
);