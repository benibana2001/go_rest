DROP TABLE IF EXISTS `users`;

CREATE TABLE users(
id int not null auto_increment,
name varchar(255) not null,
email varchar(255) not null,
primary key (id)
);

INSERT INTO users (id, name, email)
VALUES (null, 'Takeru Satou', 'takeru@mail.jp'),
 (null, 'Hanako Yamada', 'hanako@mail.jp'),
 (null, 'Sathoshi Tajima', 'satoshi@mail.jp');
