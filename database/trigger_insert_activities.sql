CREATE TRIGGER IF NOT EXISTS insert_activities AFTER INSERT ON activities
FOR EACH ROW 
BEGIN 
	SET @id_activity = NEW.id;
	SET @created_at_activity = NEW.created_at;
	SET @updated_at_activity = NEW.updated_at;
END