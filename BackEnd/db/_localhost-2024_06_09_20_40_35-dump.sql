-- MySQL dump 10.13  Distrib 8.1.0, for Win64 (x86_64)
--
-- Host: 127.0.0.1    Database: pakthani_db
-- ------------------------------------------------------
-- Server version	8.1.0

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `addresses`
--

/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `addresses` (
  `id` int NOT NULL AUTO_INCREMENT,
  `user_id` int NOT NULL,
  `kabupaten` varchar(255) NOT NULL,
  `kecamatan` varchar(255) NOT NULL,
  `kelurahan` varchar(255) NOT NULL,
  `alamat` varchar(255) NOT NULL,
  `note` varchar(255) DEFAULT NULL,
  `is_active` tinyint(1) DEFAULT '1',
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` datetime DEFAULT NULL,
  `nama_penerima` varchar(255) NOT NULL,
  PRIMARY KEY (`id`),
  KEY `addresses_user_id_foreign` (`user_id`),
  CONSTRAINT `addresses_user_id_foreign` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb3;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `addresses`
--

INSERT INTO `addresses` VALUES (1,1,'Banda Aceh','Banda Raya','Lam Lagang','Jl. Tgk Musa No. 34','Rumah cat biru',1,'2023-11-07 15:17:21','2023-11-07 15:17:21',NULL,'');
INSERT INTO `addresses` VALUES (2,25,'Aceh Besar','Peukan Bada','Ajun','jl. Ajun Laksamana','disitu rumahnya',1,'2024-04-26 21:30:54','2024-04-26 21:30:54',NULL,'');

--
-- Table structure for table `cart`
--

/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `cart` (
  `id` int NOT NULL AUTO_INCREMENT,
  `userId` int DEFAULT NULL,
  `product_id` int NOT NULL,
  `quantity` int DEFAULT NULL,
  `price` int DEFAULT NULL,
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `cart_pk` (`id`),
  KEY `cart_product_id_fk` (`product_id`),
  KEY `cart_user_id_fk` (`userId`),
  CONSTRAINT `cart_product_id_fk` FOREIGN KEY (`product_id`) REFERENCES `product` (`id`),
  CONSTRAINT `cart_user_id_fk` FOREIGN KEY (`userId`) REFERENCES `user` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=52 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `cart`
--

INSERT INTO `cart` VALUES (20,26,1,4,NULL,'2024-05-07 17:23:20','2024-05-07 17:23:20',NULL);
INSERT INTO `cart` VALUES (42,25,2,3,NULL,'2024-05-07 17:23:20','2024-05-07 17:23:20','2024-05-07 17:27:51');
INSERT INTO `cart` VALUES (44,25,7,1,NULL,'2024-05-07 17:31:32','2024-05-07 17:31:32','2024-05-07 17:40:06');
INSERT INTO `cart` VALUES (45,25,2,2,NULL,'2024-05-07 17:32:15','2024-05-07 17:32:15','2024-05-07 17:40:06');
INSERT INTO `cart` VALUES (46,25,2,2,NULL,'2024-05-22 09:22:28','2024-05-22 09:22:28','2024-05-22 09:22:58');
INSERT INTO `cart` VALUES (47,25,1,2,NULL,'2024-05-22 10:24:02','2024-05-22 10:24:02','2024-05-22 10:37:59');
INSERT INTO `cart` VALUES (48,25,10,100,NULL,'2024-05-22 10:34:42','2024-05-22 10:34:42','2024-05-22 10:37:59');

--
-- Table structure for table `category`
--

/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `category` (
  `id` int NOT NULL AUTO_INCREMENT,
  `category` varchar(255) NOT NULL,
  `icon` varchar(255) NOT NULL,
  `slug` varchar(255) NOT NULL,
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `slug` (`slug`)
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8mb3;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `category`
--

INSERT INTO `category` VALUES (1,'Terlaris','https://i.ibb.co/JnBrJHL/Bookmark.png','terlaris','2023-11-07 15:17:22','2023-11-07 15:17:22',NULL);
INSERT INTO `category` VALUES (2,'Organik','https://i.ibb.co/x5PxvYy/Natural-Food.png','organik','2023-11-07 15:17:22','2023-11-07 15:17:22',NULL);
INSERT INTO `category` VALUES (3,'Fish','https://i.ibb.co/0QDYc73/Sashimi.png','fish','2023-11-07 15:17:22','2023-11-07 15:17:22',NULL);
INSERT INTO `category` VALUES (4,'Daging','https://i.ibb.co/Lx5zTNN/Thanksgiving.png','daging','2023-11-07 15:17:22','2023-11-07 15:17:22',NULL);
INSERT INTO `category` VALUES (5,'Bread','https://i.ibb.co/9gMP7Kh/Bread.png','bread','2023-11-07 15:17:22','2023-11-07 15:17:22',NULL);
INSERT INTO `category` VALUES (6,'Sayuran','https://i.ibb.co/mqh0Rf3/Group-Of-Vegetables.png','sayuran','2023-11-07 15:17:22','2023-11-07 15:17:22',NULL);
INSERT INTO `category` VALUES (7,'Snacks','https://i.ibb.co/2FwTbjr/Potato-Chips.png','snack','2023-11-07 15:17:22','2023-11-07 15:17:22',NULL);
INSERT INTO `category` VALUES (8,'Buah','https://i.ibb.co/rMMTBd4/Group-Of-Fruits.png','buah','2023-11-07 15:17:22','2023-11-07 15:17:22',NULL);
INSERT INTO `category` VALUES (9,'Promo','https://i.ibb.co/ss2nfcr/Discount.png','promo','2023-11-07 15:17:22','2023-11-07 15:17:22',NULL);
INSERT INTO `category` VALUES (10,'Rempah','https://i.ibb.co/C2V1cMv/Garlic.png','rempah','2023-11-07 15:17:22','2023-11-07 15:17:22',NULL);

--
-- Table structure for table `order_items`
--

/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `order_items` (
  `id` int NOT NULL AUTO_INCREMENT,
  `order_id` int NOT NULL,
  `product_id` int NOT NULL,
  `quantity` int NOT NULL,
  `price` int NOT NULL,
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `order_items_orders_id_fk` (`order_id`),
  KEY `order_items_product_id_fk` (`product_id`),
  CONSTRAINT `order_items_orders_id_fk` FOREIGN KEY (`order_id`) REFERENCES `orders` (`id`),
  CONSTRAINT `order_items_product_id_fk` FOREIGN KEY (`product_id`) REFERENCES `product` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `order_items`
--

INSERT INTO `order_items` VALUES (8,26,2,2,0,'2024-06-09 17:46:14','2024-06-09 17:46:14',NULL);
INSERT INTO `order_items` VALUES (9,26,2,1,0,'2024-06-09 17:46:14','2024-06-09 17:46:14',NULL);
INSERT INTO `order_items` VALUES (10,27,2,1,120000,'2024-06-09 20:33:09','2024-06-09 20:33:09',NULL);

--
-- Table structure for table `orders`
--

/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `orders` (
  `id` int NOT NULL AUTO_INCREMENT,
  `user_id` int NOT NULL,
  `total_items` int DEFAULT NULL,
  `total_price` int DEFAULT NULL,
  `order_status` varchar(255) DEFAULT NULL,
  `payment_status` varchar(255) DEFAULT NULL,
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `orders_user_id_fk` (`user_id`),
  CONSTRAINT `orders_user_id_fk` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=28 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `orders`
--

INSERT INTO `orders` VALUES (1,25,NULL,NULL,'pending','pending','2024-05-07 17:25:08','2024-05-07 17:25:08',NULL);
INSERT INTO `orders` VALUES (2,25,NULL,NULL,'pending','pending','2024-05-07 17:25:08','2024-05-07 17:25:08',NULL);
INSERT INTO `orders` VALUES (3,25,NULL,NULL,'pending','pending','2024-05-07 17:25:08','2024-05-07 17:25:08',NULL);
INSERT INTO `orders` VALUES (4,25,NULL,NULL,'pending','pending','2024-05-07 17:25:08','2024-05-07 17:25:08',NULL);
INSERT INTO `orders` VALUES (5,25,NULL,NULL,'pending','pending','2024-05-07 17:25:08','2024-05-07 17:25:08',NULL);
INSERT INTO `orders` VALUES (7,25,NULL,NULL,'pending','pending','2024-05-07 17:25:08','2024-05-07 17:25:08',NULL);
INSERT INTO `orders` VALUES (8,25,NULL,NULL,'pending','pending','2024-05-07 17:25:08','2024-05-07 17:25:08',NULL);
INSERT INTO `orders` VALUES (9,25,NULL,NULL,'pending','pending','2024-05-07 17:25:08','2024-05-07 17:25:08',NULL);
INSERT INTO `orders` VALUES (10,25,NULL,NULL,'pending','pending','2024-05-07 17:25:08','2024-05-07 17:25:08',NULL);
INSERT INTO `orders` VALUES (11,25,NULL,NULL,'pending','pending','2024-05-07 17:25:08','2024-05-07 17:25:08',NULL);
INSERT INTO `orders` VALUES (12,25,NULL,NULL,'pending','pending','2024-05-07 17:25:08','2024-05-07 17:25:08',NULL);
INSERT INTO `orders` VALUES (13,25,NULL,NULL,'pending','pending','2024-05-07 17:25:08','2024-05-07 17:25:08',NULL);
INSERT INTO `orders` VALUES (14,25,NULL,NULL,'pending','pending','2024-05-07 17:25:08','2024-05-07 17:25:08',NULL);
INSERT INTO `orders` VALUES (15,25,NULL,NULL,'pending','pending','2024-05-07 17:25:08','2024-05-07 17:25:08',NULL);
INSERT INTO `orders` VALUES (16,25,NULL,NULL,'pending','pending','2024-05-07 17:40:06','2024-05-07 17:40:06',NULL);
INSERT INTO `orders` VALUES (17,25,NULL,NULL,'pending','pending','2024-05-22 09:22:58','2024-05-22 09:22:58',NULL);
INSERT INTO `orders` VALUES (18,25,NULL,NULL,'pending','pending','2024-05-22 10:37:59','2024-05-22 10:37:59',NULL);
INSERT INTO `orders` VALUES (26,28,3,0,'pending','pending','2024-06-09 17:46:14','2024-06-09 17:46:14',NULL);
INSERT INTO `orders` VALUES (27,28,1,120000,'pending','pending','2024-06-09 20:33:09','2024-06-09 20:33:09',NULL);

--
-- Table structure for table `product`
--

/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `product` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `category_id` int NOT NULL,
  `image` varchar(255) NOT NULL,
  `description` varchar(255) NOT NULL,
  `quantity` int NOT NULL,
  `price` int NOT NULL,
  `slug` varchar(255) NOT NULL,
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `slug` (`slug`),
  KEY `items_category_id_foreign` (`category_id`),
  CONSTRAINT `items_category_id_foreign` FOREIGN KEY (`category_id`) REFERENCES `category` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=28 DEFAULT CHARSET=utf8mb3;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `product`
--

INSERT INTO `product` VALUES (1,'Dada Ayam Fillet',4,'https://i.ibb.co/rQ0Rjth/dada-ayam-fillet.png','Dada ayam pilihan',0,30000,'dada-ayam-fillet','2023-11-07 15:17:22','2024-01-26 15:57:10',NULL);
INSERT INTO `product` VALUES (2,'Ikan Salmon Fillet',4,'https://i.ibb.co/Hnp2LkJ/ikan-salmon-fillet.png','Selamat dari istri, berakhir disini',4,120000,'ikan-salmon-fillet','2023-11-07 15:17:22','2024-06-05 19:39:41',NULL);
INSERT INTO `product` VALUES (3,'Tenggiri Fillet',4,'https://i.ibb.co/kG6xg8j/tenggiri-fillet.png','Diambil langsung dari nelayan lokal',0,50000,'tenggiri-fillet','2023-11-07 15:17:22','2024-01-26 15:57:10',NULL);
INSERT INTO `product` VALUES (4,'Tuna Fillet',4,'https://i.ibb.co/dgwStyh/tuna-fillet.png','Dulu dia bertemu dengan Nemo',0,80000,'tuna-fillet','2023-11-07 15:17:22','2024-01-26 16:00:13',NULL);
INSERT INTO `product` VALUES (5,'Udang Segar',4,'https://i.ibb.co/JCRcsHJ/udang-segar.png','Udang segar dari balik batu',0,70000,'udang-segar','2023-11-07 15:17:22','2024-01-26 16:00:13',NULL);
INSERT INTO `product` VALUES (6,'Ayam Utuh Potong',4,'https://i.ibb.co/GnJPtp6/ayam-utuh-potong.png','Sepotong ayam sehat yang berakhir di tempat kami',0,25000,'ayam-utuh-potong','2023-11-07 15:17:22','2024-01-26 16:00:13',NULL);
INSERT INTO `product` VALUES (7,'Ceker Ayam',4,'https://i.ibb.co/GJtSQVz/ceker-ayam.png','Ceker ayam pilihan dari pejantan terkuat',0,15000,'ceker-ayam','2023-11-07 15:17:22','2024-01-26 16:00:13',NULL);
INSERT INTO `product` VALUES (8,'Bawang Putih',10,'https://i.ibb.co/82rFJgt/bawang-putih.png','Deskripsi letaknya disini',0,20000,'bawang-putih','2023-11-07 15:17:22','2024-01-26 16:00:13',NULL);
INSERT INTO `product` VALUES (9,'Bumbu Ungkep',10,'https://i.ibb.co/c60yHwf/bumbu-ungkep.png','Deskripsi letaknya disini',0,5000,'bumbu-ungkep','2023-11-07 15:17:22','2024-01-26 16:00:13',NULL);
INSERT INTO `product` VALUES (10,'Jahe Merah',10,'https://i.ibb.co/nL82bJB/jahe-merah.png','Deskripsi letaknya disini',0,10000,'jahe-merah','2023-11-07 15:17:22','2024-01-26 16:00:13',NULL);
INSERT INTO `product` VALUES (11,'Jahe',10,'https://i.ibb.co/ScCJzgD/jahe.png','Deskripsi letaknya disini',0,8000,'jahe','2023-11-07 15:17:22','2024-01-26 16:00:13',NULL);
INSERT INTO `product` VALUES (12,'Kencur',10,'https://i.ibb.co/WxbNzJL/kencur.png','Deskripsi letaknya disini',0,12000,'kencur','2023-11-07 15:17:22','2024-01-26 16:00:13',NULL);
INSERT INTO `product` VALUES (13,'kunyit',10,'https://i.ibb.co/KNBZqVW/kunyit.png','Deskripsi letaknya disini',0,6000,'kunyit','2023-11-07 15:17:22','2024-01-26 16:00:13',NULL);
INSERT INTO `product` VALUES (14,'Lada Putih',10,'https://i.ibb.co/MgYvxQs/lada-putih.png','Deskripsi letaknya disini',0,25000,'lada-putih','2023-11-07 15:17:22','2024-01-26 16:00:13',NULL);
INSERT INTO `product` VALUES (15,'Bawang Bombay',10,'https://i.ibb.co/JBHZJMw/bawang-bombay.png','Deskripsi letaknya disini',0,15000,'bawang-bombay','2023-11-07 15:17:22','2024-01-26 16:00:13',NULL);
INSERT INTO `product` VALUES (16,'Baby Jagung',6,'https://i.ibb.co/fSzsptW/baby-jagung.png','Deskripsi letaknya disini',0,7000,'baby-jagung','2023-11-07 15:17:22','2024-01-26 16:01:33',NULL);
INSERT INTO `product` VALUES (17,'Labu Siam',6,'https://i.ibb.co/d4sjqzm/labu-siam.png','Deskripsi letaknya disini',0,10000,'labu-siam','2023-11-07 15:17:22','2024-01-26 16:02:55',NULL);
INSERT INTO `product` VALUES (18,'Kacang Panjang',6,'https://i.ibb.co/GJRjNfm/kacang-panjang.png','Deskripsi letaknya disini',0,15000,'kacang-panjang','2023-11-07 15:17:22','2024-01-26 16:02:55',NULL);
INSERT INTO `product` VALUES (19,'Buncis',6,'https://i.ibb.co/TH77Kcg/buncis.png','Deskripsi letaknya disini',0,18000,'buncis','2023-11-07 15:17:22','2024-01-26 16:02:55',NULL);
INSERT INTO `product` VALUES (20,'Brokoli',6,'https://i.ibb.co/MCy23km/brocoli.png','Deskripsi letaknya disini',0,20000,'brokoli','2023-11-07 15:17:22','2024-01-26 16:02:55',NULL);
INSERT INTO `product` VALUES (21,'Bayam Merah',6,'https://i.ibb.co/bBW2hmP/bayam-merah.png','Deskripsi letaknya disini',0,5000,'bayam-merah','2023-11-07 15:17:22','2024-01-26 16:02:55',NULL);
INSERT INTO `product` VALUES (22,'Alpukat',8,'https://i.ibb.co/0QSNFf9/alpukat.jpg','Deskripsi letaknya disini',0,30000,'alpukat','2023-11-07 15:17:22','2024-01-26 16:02:55',NULL);
INSERT INTO `product` VALUES (23,'Apel',8,'https://i.ibb.co/kxLrk61/apel.jpg','Deskripsi letaknya disini',0,12000,'apel','2023-11-07 15:17:22','2024-01-26 16:02:55',NULL);
INSERT INTO `product` VALUES (24,'Jeruk',8,'https://i.ibb.co/23G0t9J/jeruk.jpg','Deskripsi letaknya disini',0,15000,'jeruk','2023-11-07 15:17:22','2024-01-26 16:02:55',NULL);
INSERT INTO `product` VALUES (25,'Pepaya',8,'https://i.ibb.co/873pc06/pepaya.jpg','Deskripsi letaknya disini',0,10000,'pepaya','2023-11-07 15:17:22','2024-01-26 16:02:55',NULL);
INSERT INTO `product` VALUES (26,'Pisang',8,'https://i.ibb.co/PFpnSyY/pisang.jpg','Deskripsi letaknya disini',0,15000,'pisang','2023-11-07 15:17:22','2024-01-26 16:02:55',NULL);
INSERT INTO `product` VALUES (27,'Semangka',8,'https://i.ibb.co/5nhGbCZ/semangka.jpg','Deskripsi letaknya disini',0,5000,'semangka','2023-11-07 15:17:22','2024-01-26 16:02:55',NULL);

--
-- Table structure for table `user`
--

/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `user` (
  `id` int NOT NULL AUTO_INCREMENT,
  `email` varchar(255) NOT NULL,
  `username` varchar(255) NOT NULL,
  `no_telepon` varchar(20) NOT NULL,
  `password` varchar(255) NOT NULL,
  `role` tinyint(1) NOT NULL,
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `user_pk` (`email`)
) ENGINE=InnoDB AUTO_INCREMENT=29 DEFAULT CHARSET=utf8mb3;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user`
--

INSERT INTO `user` VALUES (1,'admin@pakthani.com','admin','0','$2y$10$IJG6EdI23/rxM1PaECBWsOEskpLXcz/hZxGpkG1HFytYE9T5PfCA2',0,'2023-11-07 15:17:21','2023-11-07 15:17:21',NULL);
INSERT INTO `user` VALUES (23,'yoo@gmail.com','admin3','0','$2a$10$EEYT7YVeHimritWtCXbeze.cm2thcBG7aTHJpT9NbskJwa9vzIigC',0,'2023-12-02 20:52:56','2023-12-02 20:52:56',NULL);
INSERT INTO `user` VALUES (25,'user@gmail.com','user','0','$2a$10$Qpmy3y5zXynrMJ.kYhhEWOnMs5FK7.SssW49eP6zsH8Ob5eyJX9hO',1,'2023-12-02 20:57:00','2023-12-02 20:57:00',NULL);
INSERT INTO `user` VALUES (26,'user2@gmail.com','user2','0','$2a$10$W.OVwYWtWp/Bo7dAUdASYeyTDICumWQfIuvnYLu2dcGwzCnYYR7Qe',1,'2023-12-11 20:19:09','2023-12-11 20:19:09',NULL);
INSERT INTO `user` VALUES (28,'iniuser@gmail.com','iniuser','082225389171','$2a$10$i/hofx4K8iM2PTJlup87teZM/41/TqBlPgvcUCVjJXX9/IV.gmQ8W',1,'2024-05-26 02:52:28','2024-05-26 02:52:28',NULL);

--
-- Table structure for table `variants_item`
--

/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `variants_item` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `price` int NOT NULL,
  `item_id` int NOT NULL,
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `variants_item_item_id_foreign` (`item_id`),
  CONSTRAINT `variants_item_item_id_foreign` FOREIGN KEY (`item_id`) REFERENCES `product` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=55 DEFAULT CHARSET=utf8mb3;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `variants_item`
--

INSERT INTO `variants_item` VALUES (1,'1kg',15000,1,'2023-11-07 15:17:22','2023-11-07 15:17:22',NULL);
INSERT INTO `variants_item` VALUES (2,'1kg',20000,1,'2023-11-07 15:17:22','2023-11-07 15:17:22',NULL);
INSERT INTO `variants_item` VALUES (3,'500gr',25000,2,'2023-11-07 15:17:22','2023-11-07 15:17:22',NULL);
INSERT INTO `variants_item` VALUES (4,'500gr',30000,2,'2023-11-07 15:17:22','2023-11-07 15:17:22',NULL);
INSERT INTO `variants_item` VALUES (5,'100gr',12000,3,'2023-11-07 15:17:22','2023-11-07 15:17:22',NULL);
INSERT INTO `variants_item` VALUES (6,'100gr',5000,3,'2023-11-07 15:17:22','2023-11-07 15:17:22',NULL);
INSERT INTO `variants_item` VALUES (7,'250gr',27000,4,'2023-11-07 15:17:22','2023-11-07 15:17:22',NULL);
INSERT INTO `variants_item` VALUES (8,'1kg',5000,4,'2023-11-07 15:17:22','2023-11-07 15:17:22',NULL);
INSERT INTO `variants_item` VALUES (9,'750gr',10000,5,'2023-11-07 15:17:22','2023-11-07 15:17:22',NULL);
INSERT INTO `variants_item` VALUES (10,'750gr',5000,5,'2023-11-07 15:17:22','2023-11-07 15:17:22',NULL);
INSERT INTO `variants_item` VALUES (11,'250gr',30000,6,'2023-11-07 15:17:22','2023-11-07 15:17:22',NULL);
INSERT INTO `variants_item` VALUES (12,'1kg',10000,6,'2023-11-07 15:17:22','2023-11-07 15:17:22',NULL);
INSERT INTO `variants_item` VALUES (13,'100gr',5000,7,'2023-11-07 15:17:22','2023-11-07 15:17:22',NULL);
INSERT INTO `variants_item` VALUES (14,'100gr',5000,7,'2023-11-07 15:17:22','2023-11-07 15:17:22',NULL);
INSERT INTO `variants_item` VALUES (15,'100gr',10000,8,'2023-11-07 15:17:22','2023-11-07 15:17:22',NULL);
INSERT INTO `variants_item` VALUES (16,'1kg',22000,8,'2023-11-07 15:17:22','2023-11-07 15:17:22',NULL);
INSERT INTO `variants_item` VALUES (17,'750gr',27000,9,'2023-11-07 15:17:22','2023-11-07 15:17:22',NULL);
INSERT INTO `variants_item` VALUES (18,'100gr',30000,9,'2023-11-07 15:17:22','2023-11-07 15:17:22',NULL);
INSERT INTO `variants_item` VALUES (19,'1kg',27000,10,'2023-11-07 15:17:22','2023-11-07 15:17:22',NULL);
INSERT INTO `variants_item` VALUES (20,'750gr',27000,10,'2023-11-07 15:17:22','2023-11-07 15:17:22',NULL);
INSERT INTO `variants_item` VALUES (21,'750gr',10000,11,'2023-11-07 15:17:22','2023-11-07 15:17:22',NULL);
INSERT INTO `variants_item` VALUES (22,'750gr',5000,11,'2023-11-07 15:17:22','2023-11-07 15:17:22',NULL);
INSERT INTO `variants_item` VALUES (23,'1kg',10000,12,'2023-11-07 15:17:22','2023-11-07 15:17:22',NULL);
INSERT INTO `variants_item` VALUES (24,'100gr',10000,12,'2023-11-07 15:17:22','2023-11-07 15:17:22',NULL);
INSERT INTO `variants_item` VALUES (25,'100gr',12000,13,'2023-11-07 15:17:22','2023-11-07 15:17:22',NULL);
INSERT INTO `variants_item` VALUES (26,'500gr',10000,13,'2023-11-07 15:17:22','2023-11-07 15:17:22',NULL);
INSERT INTO `variants_item` VALUES (27,'500gr',10000,14,'2023-11-07 15:17:22','2023-11-07 15:17:22',NULL);
INSERT INTO `variants_item` VALUES (28,'750gr',27000,14,'2023-11-07 15:17:22','2023-11-07 15:17:22',NULL);
INSERT INTO `variants_item` VALUES (29,'100gr',25000,15,'2023-11-07 15:17:22','2023-11-07 15:17:22',NULL);
INSERT INTO `variants_item` VALUES (30,'500gr',25000,15,'2023-11-07 15:17:22','2023-11-07 15:17:22',NULL);
INSERT INTO `variants_item` VALUES (31,'500gr',10000,16,'2023-11-07 15:17:22','2023-11-07 15:17:22',NULL);
INSERT INTO `variants_item` VALUES (32,'250gr',10000,16,'2023-11-07 15:17:22','2023-11-07 15:17:22',NULL);
INSERT INTO `variants_item` VALUES (33,'250gr',30000,17,'2023-11-07 15:17:22','2023-11-07 15:17:22',NULL);
INSERT INTO `variants_item` VALUES (34,'500gr',20000,17,'2023-11-07 15:17:22','2023-11-07 15:17:22',NULL);
INSERT INTO `variants_item` VALUES (35,'250gr',15000,18,'2023-11-07 15:17:22','2023-11-07 15:17:22',NULL);
INSERT INTO `variants_item` VALUES (36,'500gr',10000,18,'2023-11-07 15:17:22','2023-11-07 15:17:22',NULL);
INSERT INTO `variants_item` VALUES (37,'250gr',5000,19,'2023-11-07 15:17:22','2023-11-07 15:17:22',NULL);
INSERT INTO `variants_item` VALUES (38,'750gr',12000,19,'2023-11-07 15:17:22','2023-11-07 15:17:22',NULL);
INSERT INTO `variants_item` VALUES (39,'500gr',17000,20,'2023-11-07 15:17:22','2023-11-07 15:17:22',NULL);
INSERT INTO `variants_item` VALUES (40,'250gr',10000,20,'2023-11-07 15:17:22','2023-11-07 15:17:22',NULL);
INSERT INTO `variants_item` VALUES (41,'100gr',5000,21,'2023-11-07 15:17:22','2023-11-07 15:17:22',NULL);
INSERT INTO `variants_item` VALUES (42,'1kg',27000,21,'2023-11-07 15:17:22','2023-11-07 15:17:22',NULL);
INSERT INTO `variants_item` VALUES (43,'500gr',20000,22,'2023-11-07 15:17:22','2023-11-07 15:17:22',NULL);
INSERT INTO `variants_item` VALUES (44,'500gr',27000,22,'2023-11-07 15:17:22','2023-11-07 15:17:22',NULL);
INSERT INTO `variants_item` VALUES (45,'750gr',25000,23,'2023-11-07 15:17:22','2023-11-07 15:17:22',NULL);
INSERT INTO `variants_item` VALUES (46,'1kg',15000,23,'2023-11-07 15:17:22','2023-11-07 15:17:22',NULL);
INSERT INTO `variants_item` VALUES (47,'100gr',27000,24,'2023-11-07 15:17:22','2023-11-07 15:17:22',NULL);
INSERT INTO `variants_item` VALUES (48,'500gr',15000,24,'2023-11-07 15:17:22','2023-11-07 15:17:22',NULL);
INSERT INTO `variants_item` VALUES (49,'250gr',17000,25,'2023-11-07 15:17:22','2023-11-07 15:17:22',NULL);
INSERT INTO `variants_item` VALUES (50,'250gr',10000,25,'2023-11-07 15:17:22','2023-11-07 15:17:22',NULL);
INSERT INTO `variants_item` VALUES (51,'100gr',25000,26,'2023-11-07 15:17:22','2023-11-07 15:17:22',NULL);
INSERT INTO `variants_item` VALUES (52,'500gr',8000,26,'2023-11-07 15:17:22','2023-11-07 15:17:22',NULL);
INSERT INTO `variants_item` VALUES (53,'100gr',12000,27,'2023-11-07 15:17:22','2023-11-07 15:17:22',NULL);
INSERT INTO `variants_item` VALUES (54,'500gr',27000,27,'2023-11-07 15:17:22','2023-11-07 15:17:22',NULL);
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2024-06-09 20:40:35
