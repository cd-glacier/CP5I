-- +migrate Up
CREATE TABLE `kitchenware` (
	`id`			int UNSIGNED NOT NULL AUTO_INCREMENT,
	`name` 			varchar(255) NOt NULL,
	`recipe_id`		int NOT NULL,
	PRIMARY KEY(`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;

-- +migrate Down
DROP TABLE kitchenware;
