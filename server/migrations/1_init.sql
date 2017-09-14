-- +migrate Up
CREATE TABLE `recipe` (
	`id`			int UNSIGNED NOT NULL AUTO_INCREMENT,
	`name` 		varchar(255) NOt NULL,
	`time` 	varchar(255) NOT NULL,
	`producer_id`		int NOT NULL,
	`difficulty`		int NOT NULL,
	PRIMARY KEY(`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;

-- +migrate Down
DROP TABLE recipe;
