CREATE TABLE `categories`(
  `id` INTEGER PRIMARY KEY AUTOINCREMENT,
  `name` VARCHAR(16) NOT NULL
);
CREATE TABLE `sensors`(
  `id` INTEGER PRIMARY KEY AUTOINCREMENT,
  `name` VARCHAR(16) NOT NULL,
  `description` VARCHAR(400) NULL,
  `unit` VARCHAR(16) NOT NULL,
  `category_id` INTEGER NOT NULL,
  FOREIGN KEY(`category_id`) REFERENCES categories(`id`)
);
CREATE TABLE `readings`(
  `id` INTEGER PRIMARY KEY AUTOINCREMENT,
  `value` VARCHAR(16) NOT NULL,
  `sensor_id` INTEGER NOT NULL,
  `created_at` DATE DEFAULT(datetime('now','localtime')),
  FOREIGN KEY(`sensor_id`) REFERENCES sensors(`id`)
);
