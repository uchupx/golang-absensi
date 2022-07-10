CREATE TABLE activities(
  id int(11) unsigned AUTO_INCREMENT PRIMARY KEY NOT NULL,
  user_id int(11) unsigned NOT NULL,
  title varchar(255) not null,
  description text null,
  created_at datetime not null,
  updated_at datetime null,
  deleted_at datetime null
);

CREATE INDEX activity_user_id_idx ON activities(user_id);
CREATE INDEX activity_id_idx ON activities(id);
CREATE INDEX activity_userid_id_idx ON activities(id, user_id);