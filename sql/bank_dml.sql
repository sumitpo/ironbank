CREATE TABLE `customers` (
  `customer_id` int PRIMARY KEY AUTO_INCREMENT,
  `first_name` VARCHAR(20),
  `last_name` VARCHAR(40),
  `city` VARCHAR(50),
  `mobile_no` VARCHAR(20) COMMENT 'mobile phone number',
  `IDcard_no` VARCHAR(18),
  `dob` date COMMENT 'date of birth',
  `created_at` timestamp DEFAULT CURRENT_TIMESTAMP,
  `update_at` timestamp DEFAULT CURRENT_TIMESTAMP,
  `deleted_at` timestamp
) ENGINE = InnoDB;

CREATE INDEX `customers_index_fname` ON `customers` (`first_name`);

CREATE INDEX `customers_index_lname` ON `customers` (`last_name`);

CREATE TABLE `branches` (
  `branch_id` int PRIMARY KEY AUTO_INCREMENT,
  `branch_name` VARCHAR(100),
  `branch_location` VARCHAR(200),
  `created_at` timestamp DEFAULT CURRENT_TIMESTAMP,
  `deleted_at` timestamp
) ENGINE = InnoDB;

CREATE TABLE `accounts` (
  `account_id` bigserial PRIMARY KEY,
  `customer_id` int,
  `balance` decimal(25, 2),
  `account_status` enum(
    'Inactive',
    'Withdrawn',
    'Application Approved',
    'Pre-Opening Preparation',
    'In Opening Process',
    'Opened at Bank',
    'Active',
    'Marked for Closing',
    'Closing Request Sent to Bank',
    'Closed at Bank',
    'Closed'
  ) NOT NULL default 'Inactive' COMMENT 'full description in https://help.sap.com/docs/SAP_S4HANA_CLOUD/186460fdc35a4b64a713da9bb00deb1e/c0a95feb9940461a8c2d67f0cb18141d.html',
  `account_type` enum(
    'savings account',
    'current account',
    'fixed desposit'
  ) not null default 'current account' comment 'account type, to be completed',
  `currency` enum('cny', 'usd', 'eur') not null default 'cny' comment 'currency, default chinese currency cny (also rmb)',
  `created_at` timestamp DEFAULT CURRENT_TIMESTAMP,
  `deleted_at` timestamp
) ENGINE = InnoDB;

ALTER TABLE `accounts`
ADD FOREIGN KEY (`customer_id`) REFERENCES `customers` (`customer_id`);

CREATE INDEX `accounts_index_4` ON `accounts` (`account_id`);

CREATE TABLE `transactions` (
  `transaction_id` bigserial PRIMARY KEY,
  `transaction_type` enum('deposit', 'withdrawal', 'transfer') not null default 'deposit',
  `from_account_id` bigint,
  `to_account_id` bigint,
  `date_issued` date,
  `amount` decimal(25, 2),
  `transaction_medium` enum(
    'ATM',
    'mobile',
    'computer',
    'manual service',
    'other'
  ) not null,
  `status` enum("pending", "failed", "success") not null default "pending",
  `reference` bigint,
  `created_at` timestamp DEFAULT CURRENT_TIMESTAMP,
  `deleted_at` timestamp
) ENGINE = InnoDB;

ALTER TABLE `transactions`
ADD FOREIGN KEY (`from_account_id`) REFERENCES `accounts` (`account_id`);

ALTER TABLE `transactions`
ADD FOREIGN KEY (`to_account_id`) REFERENCES `accounts` (`account_id`);

CREATE INDEX `transactions_index_5` ON `transactions` (`from_account_id`);

CREATE INDEX `transactions_index_6` ON `transactions` (`to_account_id`);

CREATE INDEX `transactions_index_7` ON `transactions` (`from_account_id`, `to_account_id`);

CREATE TABLE `loans` (
  `loan_id` bigserial PRIMARY KEY,
  `customer_id` int,
  `branch_id` int,
  `loan_amount` decimal(25, 2),
  `date_issued` date,
  `created_at` timestamp DEFAULT CURRENT_TIMESTAMP,
  `deleted_at` timestamp
) ENGINE = InnoDB;

ALTER TABLE `loans`
ADD FOREIGN KEY (`customer_id`) REFERENCES `customers` (`customer_id`);

ALTER TABLE `loans`
ADD FOREIGN KEY (`branch_id`) REFERENCES `branches` (`branch_id`);
