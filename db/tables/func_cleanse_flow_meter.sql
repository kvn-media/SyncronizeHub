CREATE OR REPLACE FUNCTION cleanse_flow_meters()
RETURNS VOID AS
$$
BEGIN
    -- Handle missing values:
    UPDATE flow_meters
    SET type = 'Unknown' WHERE type IS NULL;
    UPDATE flow_meters
    SET location = 'Unknown' WHERE location IS NULL;
    UPDATE flow_meters
    SET fluid_type = 'Unknown' WHERE fluid_type IS NULL;

    -- Correct invalid data types:
    UPDATE flow_meters
    SET calibration_date = CAST(calibration_date AS DATE) WHERE calibration_date::DATE IS NULL;

    -- Validate data against business rules (add specific rules here):
    -- ...
EXCEPTION
    WHEN OTHERS THEN
        INSERT INTO error_logs (function_name, error_message, stack_trace)
        VALUES ('cleanse_flow_meters', SQLERRM, pg_exception_stacktrace());
        RAISE NOTICE 'Error during flow_meters cleansing: %', SQLERRM;
END;
$$
LANGUAGE plpgsql;
