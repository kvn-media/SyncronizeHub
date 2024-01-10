CREATE TRIGGER cleanse_measurements_after_insert
AFTER INSERT ON measurements
FOR EACH ROW
EXECUTE PROCEDURE cleanse_measurements();

-- Create similar triggers for other tables.