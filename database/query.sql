CREATE DATABASE IF NOT EXISTS todolist;
USE todolist;

DROP TABLE IF EXISTS activities;
DROP TABLE IF EXISTS todos;

CREATE TABLE IF NOT EXISTS activities (
  id int NOT NULL PRIMARY KEY AUTO_INCREMENT,
  email varchar(255) NOT NULL,
  title varchar(255) NOT NULL,
  created_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  deleted_at datetime DEFAULT NULL
) ENGINE=InnoDB;

CREATE TABLE IF NOT EXISTS todos (
  id int NOT NULL PRIMARY KEY AUTO_INCREMENT,
  title varchar(255) NOT NULL,
  activity_group_id int NOT NULL,
  is_active bool NOT NULL DEFAULT true,
  priority varchar(55) NOT NULL DEFAULT 'very-high',
  created_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  deleted_at datetime DEFAULT NULL
) ENGINE=InnoDB ;


CREATE TRIGGER IF NOT EXISTS insert_act AFTER INSERT ON activities
FOR EACH ROW 
BEGIN 
	SET @id = NEW.id;
	SET @created_at = NEW.created_at;
	SET @updated_at = NEW.updated_at;
END

CREATE TRIGGER IF NOT EXISTS insert_todo AFTER INSERT ON todos
FOR EACH ROW 
BEGIN 
	SET @id = NEW.id;
	SET @created_at = NEW.created_at;
	SET @updated_at = NEW.updated_at;
	SET @is_active = NEW.is_active;
	SET @priority = NEW.priority;
END
