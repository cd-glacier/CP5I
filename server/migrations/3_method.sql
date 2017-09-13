-- +migrate Up
CREATE TABLE `method` (
	`id`			int UNSIGNED NOT NULL AUTO_INCREMENT,
	`recipe_id`		int NOT NULL,
	`method_order`		int NOT NULL,
	`content` 			varchar(255) NOt NULL,
	PRIMARY KEY(`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;

-- +migrate Down
DROP TABLE method;
