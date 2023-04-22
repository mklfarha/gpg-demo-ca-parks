
CREATE TABLE IF NOT EXISTS `feature` ( 
  `id` CHAR(36) NOT NULL  ,
  `name` VARCHAR(255) NOT NULL  ,
  `status` INT NOT NULL  ,
  `created_at` DATETIME NOT NULL  default CURRENT_TIMESTAMP,
  `updated_at` DATETIME NOT NULL  default CURRENT_TIMESTAMP,
  `user_id` CHAR(36) NOT NULL  ,  
  PRIMARY KEY (`id`),
  INDEX `status` (`status` ASC),
  INDEX `created_at` (`created_at` ASC),
  INDEX `updated_at` (`updated_at` ASC))  
ENGINE = InnoDB;
