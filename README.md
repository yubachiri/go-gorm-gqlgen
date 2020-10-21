たぶん `docker-compose up` だけで動く
初回は
`docker exec -it go-gorm-gqlgen_db_1 sh`
して、
`mysql -u localuser -p`
`password`

```
CREATE TABLE IF NOT EXISTS `todo` (
  `id` varchar(64) NOT NULL,
  `text` varchar(256) NOT NULL,
  `done` bool NOT NULL,
  `user_id` varchar(64) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE utf8_bin;

CREATE TABLE IF NOT EXISTS `user` (
  `id` varchar(64) NOT NULL,
  `name` varchar(256) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE utf8_bin;

```

する必要があるかも
