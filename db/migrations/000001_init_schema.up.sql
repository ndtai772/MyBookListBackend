CREATE TABLE `accounts` (
  `id` BIGINT UNSIGNED PRIMARY KEY,
  `username` VARCHAR(255) UNIQUE NOT NULL,
  `email` VARCHAR(255) UNIQUE NOT NULL,
  `avatar_uri` VARCHAR(255),
  `role` ENUM ('admin', 'user') NOT NULL DEFAULT 'user',
  `modified_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE `books` (
  `id` INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
  `title` VARCHAR(255) NOT NULL,
  `author` VARCHAR(255) NOT NULL,
  `description` VARCHAR(10000) NOT NULL,
  `modified_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE `categories` (
  `id` INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
  `name` VARCHAR(255) NOT NULL,
  `description` VARCHAR(10000),
  `modified_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE `book_category` (
  `book_id` INT UNSIGNED NOT NULL,
  `category_id` INT UNSIGNED NOT NULL,
  `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE `bookmarks` (
  `id` BIGINT UNSIGNED PRIMARY KEY,
  `book_id` INT UNSIGNED NOT NULL,
  `account_id` BIGINT UNSIGNED NOT NULL,
  `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE `rates` (
  `id` BIGINT UNSIGNED PRIMARY KEY,
  `book_id` INT UNSIGNED NOT NULL,
  `account_id` BIGINT UNSIGNED NOT NULL,
  `modified_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE `comments` (
  `id` BIGINT UNSIGNED PRIMARY KEY,
  `content` VARCHAR(1000),
  `book_id` INT UNSIGNED NOT NULL,
  `created_by` BIGINT UNSIGNED NOT NULL,
  `modified_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE `feedbacks` (
  `id` BIGINT UNSIGNED PRIMARY KEY,
  `content` VARCHAR(1000) NOT NULL,
  `created_by` BIGINT UNSIGNED NOT NULL,
  `message` VARCHAR(1000),
  `status` ENUM ('processing', 'resolved', 'not_resolved') NOT NULL,
  `modified_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `created_at` TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

ALTER TABLE `book_category` ADD FOREIGN KEY (`book_id`) REFERENCES `books` (`id`);

ALTER TABLE `book_category` ADD FOREIGN KEY (`category_id`) REFERENCES `categories` (`id`);

ALTER TABLE `bookmarks` ADD FOREIGN KEY (`book_id`) REFERENCES `books` (`id`);

ALTER TABLE `bookmarks` ADD FOREIGN KEY (`account_id`) REFERENCES `accounts` (`id`);

ALTER TABLE `rates` ADD FOREIGN KEY (`book_id`) REFERENCES `books` (`id`);

ALTER TABLE `rates` ADD FOREIGN KEY (`account_id`) REFERENCES `accounts` (`id`);

ALTER TABLE `comments` ADD FOREIGN KEY (`book_id`) REFERENCES `books` (`id`);

ALTER TABLE `comments` ADD FOREIGN KEY (`created_by`) REFERENCES `accounts` (`id`);

ALTER TABLE `feedbacks` ADD FOREIGN KEY (`created_by`) REFERENCES `accounts` (`id`);

CREATE INDEX `accounts_index_0` ON `accounts` (`username`);

CREATE INDEX `accounts_index_1` ON `accounts` (`email`);

CREATE INDEX `books_index_2` ON `books` (`author`);

CREATE INDEX `books_index_3` ON `books` (`modified_at`);

CREATE INDEX `categories_index_4` ON `categories` (`name`);

CREATE INDEX `book_category_index_5` ON `book_category` (`book_id`);

CREATE INDEX `book_category_index_6` ON `book_category` (`category_id`);

CREATE UNIQUE INDEX `book_category_index_7` ON `book_category` (`book_id`, `category_id`);

CREATE INDEX `bookmarks_index_8` ON `bookmarks` (`book_id`);

CREATE INDEX `bookmarks_index_9` ON `bookmarks` (`account_id`);

CREATE UNIQUE INDEX `bookmarks_index_10` ON `bookmarks` (`book_id`, `account_id`);

CREATE INDEX `comments_index_11` ON `comments` (`book_id`);

CREATE INDEX `comments_index_12` ON `comments` (`created_by`);
