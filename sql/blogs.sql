-- --------------------------------------------------------
-- Database: blog_db
-- --------------------------------------------------------
CREATE DATABASE IF NOT EXISTS `blog_db` DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
USE `blog_db`;

-- --------------------------------------------------------
-- Table structure for table `users`
-- --------------------------------------------------------
CREATE TABLE IF NOT EXISTS `users` (
  `id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
  `username` VARCHAR(50) NOT NULL,
  `email` VARCHAR(100) NOT NULL,
  `password` VARCHAR(255) NOT NULL,
  `registered_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE INDEX `username` (`username`),
  UNIQUE INDEX `email` (`email`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- --------------------------------------------------------
-- Table structure for table `posts`
-- --------------------------------------------------------
CREATE TABLE IF NOT EXISTS `posts` (
  `id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
  `title` VARCHAR(255) NOT NULL,
  `slug` VARCHAR(255) NOT NULL,
  `content` LONGTEXT NOT NULL,
  `excerpt` VARCHAR(500) NULL,
  `author_id` INT UNSIGNED NOT NULL,
  `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` DATETIME NULL,
  `status` ENUM('draft', 'publish', 'trash') NOT NULL DEFAULT 'draft',
  `featured_image` VARCHAR(255) NULL,
  PRIMARY KEY (`id`),
  UNIQUE INDEX `slug` (`slug`),
  INDEX `author_id` (`author_id`),
  CONSTRAINT `fk_posts_author`
    FOREIGN KEY (`author_id`)
    REFERENCES `users` (`id`)
    ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- --------------------------------------------------------
-- Table structure for table `categories`
-- --------------------------------------------------------
CREATE TABLE IF NOT EXISTS `categories` (
  `id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(50) NOT NULL,
  `slug` VARCHAR(50) NOT NULL,
  `description` VARCHAR(500) NULL,
  `parent_id` INT UNSIGNED NULL,
  PRIMARY KEY (`id`),
  UNIQUE INDEX `slug` (`slug`),
  INDEX `parent_id` (`parent_id`),
  CONSTRAINT `fk_categories_parent`
    FOREIGN KEY (`parent_id`)
    REFERENCES `categories` (`id`)
    ON DELETE SET NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- --------------------------------------------------------
-- Table structure for table `posts_categories`
-- --------------------------------------------------------
CREATE TABLE IF NOT EXISTS `posts_categories` (
  `post_id` INT UNSIGNED NOT NULL,
  `category_id` INT UNSIGNED NOT NULL,
  PRIMARY KEY (`post_id`, `category_id`),
  INDEX `category_id` (`category_id`),
  CONSTRAINT `fk_posts_categories_post`
    FOREIGN KEY (`post_id`)
    REFERENCES `posts` (`id`)
    ON DELETE CASCADE,
  CONSTRAINT `fk_posts_categories_category`
    FOREIGN KEY (`category_id`)
    REFERENCES `categories` (`id`)
    ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- --------------------------------------------------------
-- Table structure for table `tags`
-- --------------------------------------------------------
CREATE TABLE IF NOT EXISTS `tags` (
  `id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(50) NOT NULL,
  `slug` VARCHAR(50) NOT NULL,
  `description` VARCHAR(500) NULL,
  PRIMARY KEY (`id`),
  UNIQUE INDEX `slug` (`slug`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- --------------------------------------------------------
-- Table structure for table `posts_tags`
-- --------------------------------------------------------
CREATE TABLE IF NOT EXISTS `posts_tags` (
  `post_id` INT UNSIGNED NOT NULL,
  `tag_id` INT UNSIGNED NOT NULL,
  PRIMARY KEY (`post_id`, `tag_id`),
  INDEX `tag_id` (`tag_id`),
  CONSTRAINT `fk_posts_tags_post`
    FOREIGN KEY (`post_id`)
    REFERENCES `posts` (`id`)
    ON DELETE CASCADE,
  CONSTRAINT `fk_posts_tags_tag`
    FOREIGN KEY (`tag_id`)
    REFERENCES `tags` (`id`)
    ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- --------------------------------------------------------
-- Table structure for table `menus`
-- --------------------------------------------------------
CREATE TABLE IF NOT EXISTS `menus` (
  `id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(50) NOT NULL,
  `location` VARCHAR(50) NULL,
  PRIMARY KEY (`id`),
  UNIQUE INDEX `location` (`location`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- --------------------------------------------------------
-- Table structure for table `menu_items`
-- --------------------------------------------------------
CREATE TABLE IF NOT EXISTS `menu_items` (
  `id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
  `menu_id` INT UNSIGNED NOT NULL,
  `parent_id` INT UNSIGNED NULL,
  `title` VARCHAR(100) NOT NULL,
  `url` VARCHAR(255) NOT NULL,
  `order` INT NOT NULL DEFAULT 0,
  `class` VARCHAR(50) NULL,
  `target` VARCHAR(20) NULL,
  PRIMARY KEY (`id`),
  INDEX `menu_id` (`menu_id`),
  INDEX `parent_id` (`parent_id`),
  CONSTRAINT `fk_menu_items_menu`
    FOREIGN KEY (`menu_id`)
    REFERENCES `menus` (`id`)
    ON DELETE CASCADE,
  CONSTRAINT `fk_menu_items_parent`
    FOREIGN KEY (`parent_id`)
    REFERENCES `menu_items` (`id`)
    ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- --------------------------------------------------------
-- Table structure for table `settings`
-- --------------------------------------------------------
CREATE TABLE IF NOT EXISTS `settings` (
  `id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
  `key` VARCHAR(50) NOT NULL,
  `value` TEXT NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE INDEX `key` (`key`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- --------------------------------------------------------
-- Table structure for table `widgets`
-- --------------------------------------------------------
CREATE TABLE IF NOT EXISTS `widgets` (
  `id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
  `type` VARCHAR(50) NOT NULL COMMENT 'sidebar, footer, etc.',
  `title` VARCHAR(255) NULL,
  `content` TEXT NOT NULL,
  `position` VARCHAR(50) NOT NULL,
  `sort_order` INT NOT NULL DEFAULT 0,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;