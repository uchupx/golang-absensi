CREATE TABLE users(
  id int(11) unsigned primary key AUTO_INCREMENT not null,
  name varchar(255) not null,
  email varchar(255) not null UNIQUE,
  password text not null,
  created_at datetime not null,
  updated_at datetime null
);

CREATE INDEX user_id_idx on users(id);