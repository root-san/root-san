DROP DATABASE IF EXISTS root_san;
CREATE DATABASE root_san;
USE root_san;

CREATE TABLE IF NOT EXISTS `rooms` (
  `id` char(36) NOT NULL,
  `name` varchar(300) NOT NULL,
  `created_at` timestamp DEFAULT CURRENT_TIMESTAMP NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

CREATE TABLE IF NOT EXISTS `room_members` (
  `member_id` char(36) NOT NULL,
  `room_id` char(36) NOT NULL,
  `name` varchar(300) NOT NULL,
  `created_at` timestamp DEFAULT CURRENT_TIMESTAMP NOT NULL,
  PRIMARY KEY (`member_id`, `room_id`),
  FOREIGN KEY (`room_id`) REFERENCES `rooms`(`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

CREATE TABLE IF NOT EXISTS `transactions` (
  `id` char(36) NOT NULL,
  `room_id` char(36) NOT NULL,
  `amount` int NOT NULL,
  `payer_id` char(36) NOT NULL,
  `paid_at` datetime NOT NULL,
  `description` varchar(300) NOT NULL,
  `created_at` timestamp DEFAULT CURRENT_TIMESTAMP NOT NULL,
  `updated_at` timestamp DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP NOT NULL,
  PRIMARY KEY (`id`),
  FOREIGN KEY (`room_id`) REFERENCES `rooms`(`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

CREATE TABLE IF NOT EXISTS `transaction_receivers` (
    `member_id` char(36) NOT NULL,
    `transaction_id` char(36) NOT NULL,
    `created_at` timestamp DEFAULT CURRENT_TIMESTAMP NOT NULL,
    PRIMARY KEY (`member_id`, `transaction_id`),
    FOREIGN KEY (`transaction_id`) REFERENCES `transactions`(`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
