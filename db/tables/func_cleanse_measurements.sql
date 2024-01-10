CREATE OR REPLACE FUNCTION cleanse_measurements()
RETURNS VOID AS
$$
BEGIN
    -- Remove duplicates based on timestamp and flow_meter_id:
    DELETE FROM measurements m1
    WHERE EXISTS (
        SELECT 1 FROM measurements m2
        WHERE m2.timestamp = m1.timestamp AND m2.flow_meter_id = m1.flow_meter_id
        AND m2.id < m1.id
    );

    -- Handle missing values:
    UPDATE measurements
    SET flow_rate = 0.0 WHERE flow_rate IS NULL;
    UPDATE measurements
    SET temperature = 0.0 WHERE temperature IS NULL;
    UPDATE measurements
    SET pressure = 0.0 WHERE pressure IS NULL;
    UPDATE measurements
    SET tank_level = 0.0 WHERE tank_level IS NULL;

    -- Correct invalid data types:
    UPDATE measurements
    SET flow_rate = CAST(flow_rate AS DECIMAL(10,2)) WHERE flow_rate::DECIMAL(10,2) IS NULL;
    UPDATE measurements
    SET temperature = CAST(temperature AS DECIMAL(5,2)) WHERE temperature::DECIMAL(5,2) IS NULL;
    UPDATE measurements
    SET pressure = CAST(pressure AS DECIMAL(10,2)) WHERE pressure::DECIMAL(10,2) IS NULL;
    UPDATE measurements
    SET tank_level = CAST(tank_level AS DECIMAL(5,2)) WHERE tank_level::DECIMAL(5,2) IS NULL;

    -- Identify and correct outliers (adjust thresholds as needed):
    UPDATE measurements
    SET flow_rate = 0.0 WHERE flow_rate < 0 OR flow_rate > 1000;
    UPDATE measurements
    SET temperature = 0.0 WHERE temperature < -50 OR temperature > 200;
    UPDATE measurements
    SET pressure = 0.0 WHERE pressure < 0 OR pressure > 1000000;
    UPDATE measurements
    SET tank_level = 0.0 WHERE tank_level < 0 OR tank_level > 100;

    -- Validate data against business rules (add specific rules here):
    -- ...
EXCEPTION
    WHEN OTHERS THEN
        INSERT INTO error_logs (function_name, error_message, stack_trace)
        VALUES ('cleanse_measurements', SQLERRM, pg_exception_stacktrace());
        RAISE NOTICE 'Error during measurements cleansing: %', SQLERRM;
END;
$$
LANGUAGE plpgsql;