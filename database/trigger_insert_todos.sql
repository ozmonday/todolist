CREATE TRIGGER IF NOT EXISTS insert_todos AFTER INSERT ON todos
FOR EACH ROW 
BEGIN 
	SET @id_todo = NEW.id;
	SET @created_at_todo = NEW.created_at;
	SET @updated_at_todo = NEW.updated_at;
	SET @is_active_todo = NEW.is_active;
	SET @priority_todo = NEW.priority;
END
