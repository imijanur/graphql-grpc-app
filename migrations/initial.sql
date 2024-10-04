-- Create the database if it doesn't exist
CREATE DATABASE IF NOT EXISTS `graphql_grpc_app`
  CHARACTER SET = utf8mb4
  COLLATE = utf8mb4_unicode_ci;

-- Select the database
USE `graphql_grpc_app`;

-- Create `users` table
CREATE TABLE IF NOT EXISTS `users` (
    `id` INT(11) NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `email` VARCHAR(255) NOT NULL UNIQUE,
    `status` ENUM('active', 'inactive') NOT NULL DEFAULT 'active',
    `created_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `modified_at` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Create `user_contact` table (One-to-one relationship with `users`)
CREATE TABLE IF NOT EXISTS `user_contact` (
    `id` INT(11) NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `first_name` VARCHAR(100) NOT NULL,
    `last_name` VARCHAR(100) NOT NULL,
    `phone` VARCHAR(20) NOT NULL,
    `user_id` INT(11) NOT NULL,
    FOREIGN KEY (`user_id`) REFERENCES `users`(`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- Create `user_address` table (One-to-many relationship with `users`)
CREATE TABLE IF NOT EXISTS `user_address` (
    `id` INT(11) NOT NULL AUTO_INCREMENT PRIMARY KEY,
    `name` VARCHAR(255) NOT NULL,
    `prefix` VARCHAR(50),
    `street_address_1` VARCHAR(255) NOT NULL,
    `street_address_2` VARCHAR(255),
    `city` VARCHAR(100) NOT NULL,
    `state` VARCHAR(100) NOT NULL,
    `zip_code` VARCHAR(20) NOT NULL,
    `user_id` INT(11) NOT NULL,
    FOREIGN KEY (`user_id`) REFERENCES `users`(`id`) ON DELETE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

ALTER TABLE user_contact
ADD CONSTRAINT fk_user_contact_user
FOREIGN KEY (user_id) REFERENCES users(id),
ADD UNIQUE (user_id);

INSERT INTO users (email, status, created_at, modified_at)
VALUES
    ('john.doe@example.com', 'active', NOW(), NOW()),
    ('jane.smith@example.com', 'inactive', NOW(), NOW()),
    ('michael.brown@example.com', 'active', NOW(), NOW()),
    ('alice.williams@example.com', 'active', NOW(), NOW());


INSERT INTO user_contact (first_name, last_name, phone, user_id)
VALUES
    ('John', 'Doe', '1234567890', 1),
    ('Jane', 'Smith', '9876543210', 2),
    ('Michael', 'Brown', '1231231234', 3),
    ('Alice', 'Williams', '5556667777', 4);


INSERT INTO user_address (name, prefix, street_address_1, street_address_2, city, state, zip_code, user_id)
VALUES
    ('Home', 'Mr.', '123 Main St', 'Apt 4B', 'New York', 'NY', '10001', 1),
    ('Office', 'Ms.', '456 Oak Ave', 'Suite 120', 'San Francisco', 'CA', '94103', 1),
    ('Home', 'Mr.', '789 Pine Dr', '', 'Los Angeles', 'CA', '90001', 2),
    ('Office', 'Ms.', '135 Maple St', '', 'Chicago', 'IL', '60601', 3),
    ('Home', 'Mr.', '246 Cedar Blvd', '', 'Miami', 'FL', '33101', 4);

