
CREATE TABLE IF NOT EXISTS `park_has_feature` ( 
  `id` CHAR(36) NOT NULL  ,
  `details` VARCHAR(255) NOT NULL  ,
  `status` INT NOT NULL  ,
  `created_at` DATETIME NOT NULL  default CURRENT_TIMESTAMP,
  `updated_at` DATETIME NOT NULL  default CURRENT_TIMESTAMP,
  `park_id` CHAR(36) NOT NULL  ,
  `feature_id` CHAR(36) NOT NULL  ,  
  PRIMARY KEY (`id`),
  INDEX `status` (`status` ASC),
  INDEX `created_at` (`created_at` ASC),
  INDEX `updated_at` (`updated_at` ASC),
  INDEX `park_id` (`park_id` ASC),
  INDEX `feature_id` (`feature_id` ASC))  
ENGINE = InnoDB;
