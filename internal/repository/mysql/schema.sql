DROP TABLE IF EXISTS `discounts`;
DROP TABLE IF EXISTS `orders`;
DROP TABLE IF EXISTS `users`;

CREATE TABLE `users` (
  `user_id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(100) NOT NULL,
  `email` varchar(255) NOT NULL,
  `location` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`user_id`),
  UNIQUE KEY `email` (`email`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `orders` (
  `order_id` int NOT NULL AUTO_INCREMENT,
  `user_id` int NOT NULL,
  `item` varchar(255) NOT NULL,
  `quantity` int NOT NULL,
  `price` decimal(10,2) NOT NULL,
  PRIMARY KEY (`order_id`),
  KEY `user_id` (`user_id`),
  CONSTRAINT `orders_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `users` (`user_id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE `discounts` (
  `discount_id` int NOT NULL AUTO_INCREMENT,
  `order_id` int NOT NULL,
  `discount_percent` decimal(5,2) NOT NULL,
  `description` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`discount_id`),
  KEY `order_id` (`order_id`),
  CONSTRAINT `discounts_ibfk_1` FOREIGN KEY (`order_id`) REFERENCES `orders` (`order_id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

-- Seed data
INSERT INTO `users` VALUES (1,'Alice Johnson','alice@example.com','New York'),(2,'Bob Smith','bob@example.com','Los Angeles'),(3,'Charlie Brown','charlie@example.com','Chicago'),(4,'Diana Prince','diana@example.com','Miami');

INSERT INTO `orders` VALUES (1,1,'Laptop',1,1200.00),(2,1,'Mouse',2,25.50),(3,2,'Keyboard',1,75.00),(4,3,'Monitor',2,300.00),(5,4,'Headphones',1,150.00),(6,2,'Webcam',1,90.00);

INSERT INTO `discounts` VALUES (1,1,10.00,'Promo on laptop'),(2,3,5.00,'Keyboard sale'),(3,4,15.00,'Bulk monitor discount');
