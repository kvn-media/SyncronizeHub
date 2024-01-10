CREATE TABLE customer_downloads (
    id SERIAL PRIMARY KEY,
    customer_id INT REFERENCES dim_customers(id) NOT NULL,
    timestamp TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    data_type VARCHAR(50) NOT NULL,  -- Type of data downloaded (e.g., measurements, atgs, tanks)
    file_format VARCHAR(10)  -- Optional format of downloaded file
);