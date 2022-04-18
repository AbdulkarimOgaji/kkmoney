CREATE TABLE `accounts` (
  `id` bigint PRIMARY KEY AUTO_INCREMENT,
  `holder` varchar(255) NOT NULL,
  `balance` bigint NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT (now()),
  `currency` varchar(255) NOT NULL
);

CREATE TABLE `entries` (
  `id` bigint PRIMARY KEY AUTO_INCREMENT,
  `account_id` bigint NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT (now()),
  `amount` bigint NOT NULL COMMENT 'amount can be positive or negative'
);

CREATE TABLE `transactions` (
  `id` bigint PRIMARY KEY AUTO_INCREMENT,
  `from_account_id` bigint NOT NULL,
  `to_account_id` bigint NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT (now()),
  `amount` bigint NOT NULL COMMENT 'amount can only be positive'
);

ALTER TABLE `transactions` ADD FOREIGN KEY (`from_account_id`) REFERENCES `accounts` (`id`);

ALTER TABLE `transactions` ADD FOREIGN KEY (`to_account_id`) REFERENCES `accounts` (`id`);

ALTER TABLE `entries` ADD FOREIGN KEY (`account_id`) REFERENCES `accounts` (`id`);

CREATE INDEX `accounts_index_0` ON `accounts` (`holder`);

CREATE INDEX `entries_index_1` ON `entries` (`account_id`);

CREATE INDEX `transactions_index_2` ON `transactions` (`from_account_id`);

CREATE INDEX `transactions_index_3` ON `transactions` (`to_account_id`);

CREATE INDEX `transactions_index_4` ON `transactions` (`from_account_id`, `to_account_id`);
