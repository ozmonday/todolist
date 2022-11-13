CREATE TABLE IF NOT EXISTS todos (
  id int NOT NULL PRIMARY KEY AUTO_INCREMENT,
  title varchar(255) NOT NULL DEFAULT 'unknow',
  activity_group_id int NOT NULL,
  is_active bool NOT NULL DEFAULT true,
  priority varchar(55) NOT NULL DEFAULT 'very-high',
  created_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  deleted_at datetime DEFAULT NULL
) ENGINE=InnoDB ;
