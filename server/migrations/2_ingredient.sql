-- +migrate Up
CREATE TABLE `ingredient` (
	`id`			int UNSIGNED NOT NULL AUTO_INCREMENT,
	`recipe_id`		int NOT NULL,
	`name` 			varchar(255) NOt NULL,
	`quantity` 	varchar(255) NOT NULL,
	PRIMARY KEY(`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;

-- +migrate Down
DROP TABLE ingredient;
