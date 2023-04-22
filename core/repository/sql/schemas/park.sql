
CREATE TABLE IF NOT EXISTS `park` ( 
  `id` CHAR(36) NOT NULL  ,
  `name` VARCHAR(255) NOT NULL  ,
  `main_image` VARCHAR(255) NOT NULL  ,
  `phone` VARCHAR(255) NOT NULL  ,
  `hours` VARCHAR(255) NOT NULL  ,
  `allows_dogs` TINYINT(1) NOT NULL  ,
  `links` JSON NOT NULL  ,
  `status` INT NOT NULL  ,
  `created_at` DATETIME NOT NULL  default CURRENT_TIMESTAMP,
  `updated_at` DATETIME NOT NULL  default CURRENT_TIMESTAMP,
  `recreation_area_id` CHAR(36) NOT NULL  ,
  `user_id` CHAR(36) NOT NULL  ,  
  PRIMARY KEY (`id`),
  INDEX `status` (`status` ASC),
  INDEX `created_at` (`created_at` ASC),
  INDEX `updated_at` (`updated_at` ASC),
  INDEX `recreation_area_id` (`recreation_area_id` ASC))  
ENGINE = InnoDB;
