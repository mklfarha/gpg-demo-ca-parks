
CREATE TABLE IF NOT EXISTS `event` ( 
  `id` CHAR(36) NOT NULL  ,
  `name` VARCHAR(255) NOT NULL  ,
  `description` TEXT NOT NULL  ,
  `main_image` VARCHAR(255) NOT NULL  ,
  `start_date` DATE NOT NULL  default '2022-02-02',
  `end_date` DATE NOT NULL  default '2022-02-02',
  `status` INT NOT NULL  ,
  `created_at` DATETIME NOT NULL  default CURRENT_TIMESTAMP,
  `updated_at` DATETIME NOT NULL  default CURRENT_TIMESTAMP,
  `park_id` CHAR(36) NOT NULL  ,
  `user_id` CHAR(36) NOT NULL  ,  
  PRIMARY KEY (`id`),
  INDEX `status` (`status` ASC),
  INDEX `created_at` (`created_at` ASC),
  INDEX `updated_at` (`updated_at` ASC),
  INDEX `park_id` (`park_id` ASC))  
ENGINE = InnoDB;
