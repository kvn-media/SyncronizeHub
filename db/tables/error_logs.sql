CREATE TABLE error_logs (
    id SERIAL PRIMARY KEY,
    timestamp TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    function_name VARCHAR(100) NOT NULL,
    error_message TEXT NOT NULL,
    stack_trace TEXT
);