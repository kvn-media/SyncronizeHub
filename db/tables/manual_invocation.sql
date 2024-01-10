CREATE FUNCTION cleanse_all_tables()
RETURNS VOID AS
$$
BEGIN
    PERFORM cleanse_measurements();
    -- Call cleansing functions for other tables;
END;
$$
LANGUAGE plpgsql;