CREATE TABLE attendances(
  id int(11) unsigned AUTO_INCREMENT PRIMARY KEY NOT NULL,
  user_id int(11) unsigned NOT NULL,
  type varchar(255) NOT NULL,
  created_at datetime
);

CREATE INDEX attedance_user_id_idx ON attendances(user_id);