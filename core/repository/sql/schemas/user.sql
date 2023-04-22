
CREATE TABLE IF NOT EXISTS `user` ( 
  `id` CHAR(36) NOT NULL  ,
  `name` VARCHAR(255) NOT NULL  ,
  `email` VARCHAR(255) NOT NULL UNIQUE ,
  `password` VARCHAR(255) NOT NULL  ,
  `status` INT NOT NULL  ,
  `created_at` DATETIME NOT NULL  default CURRENT_TIMESTAMP,
  `updated_at` DATETIME NOT NULL  default CURRENT_TIMESTAMP,  
  PRIMARY KEY (`id`),
  INDEX `status` (`status` ASC))  
ENGINE = InnoDB;
