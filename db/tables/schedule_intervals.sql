CREATE FUNCTION schedule_cleansing()
RETURNS TRIGGER AS
$$
BEGIN
    PERFORM cleanse_measurements();
    -- PERFORM cleansing functions for other tables;
    RETURN NULL;
END;
$$
LANGUAGE plpgsql;

CREATE EVENT scheduled_cleansing
    ON SCHEDULE EVERY 1 HOUR -- Adjust interval as needed
    DO
    EXECUTE PROCEDURE schedule_cleansing();