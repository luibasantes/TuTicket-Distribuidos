-- MySQL dump 10.13  Distrib 8.0.14, for Linux (x86_64)
--
-- Host: localhost    Database: tu_ticket
-- ------------------------------------------------------
-- Server version	8.0.14

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
 SET NAMES utf8mb4 ;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `asiento_models`
--

DROP TABLE IF EXISTS `asiento_models`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
 SET character_set_client = utf8mb4 ;
CREATE TABLE `asiento_models` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  `numero` int(10) unsigned DEFAULT NULL,
  `disponible` tinyint(1) DEFAULT NULL,
  `evento_id` int(10) unsigned DEFAULT NULL,
  `localidad_id` int(10) unsigned DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_asiento_models_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=18 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `asiento_models`
--

LOCK TABLES `asiento_models` WRITE;
/*!40000 ALTER TABLE `asiento_models` DISABLE KEYS */;
INSERT INTO `asiento_models` VALUES (1,NULL,NULL,NULL,1,1,1,1),(2,NULL,NULL,NULL,2,1,1,1),(3,NULL,NULL,NULL,3,1,1,1),(4,NULL,NULL,NULL,4,1,1,1),(5,NULL,NULL,NULL,5,1,1,1),(6,NULL,NULL,NULL,1,1,1,2),(7,NULL,NULL,NULL,2,1,1,2),(8,NULL,NULL,NULL,3,1,1,2),(9,NULL,NULL,NULL,4,1,1,2),(10,NULL,NULL,NULL,5,1,1,2),(11,NULL,NULL,NULL,1,1,1,3),(12,NULL,NULL,NULL,2,1,1,3),(13,NULL,NULL,NULL,3,1,1,3),(14,NULL,NULL,NULL,4,1,1,3),(15,NULL,NULL,NULL,5,1,1,3),(16,NULL,NULL,NULL,1,1,1,4),(17,NULL,NULL,NULL,2,1,1,4);
/*!40000 ALTER TABLE `asiento_models` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `boleto_models`
--

DROP TABLE IF EXISTS `boleto_models`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
 SET character_set_client = utf8mb4 ;
CREATE TABLE `boleto_models` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  `asiento_id` int(10) unsigned DEFAULT NULL,
  `usuario_id` int(10) unsigned DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_boleto_models_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `boleto_models`
--

LOCK TABLES `boleto_models` WRITE;
/*!40000 ALTER TABLE `boleto_models` DISABLE KEYS */;
/*!40000 ALTER TABLE `boleto_models` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `evento_models`
--

DROP TABLE IF EXISTS `evento_models`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
 SET character_set_client = utf8mb4 ;
CREATE TABLE `evento_models` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  `nombre` varchar(255) DEFAULT NULL,
  `descripcion` varchar(255) DEFAULT NULL,
  `fecha` varchar(255) DEFAULT NULL,
  `sede_id` int(11) DEFAULT NULL,
  `agotado` tinyint(1) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_evento_models_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `evento_models`
--

LOCK TABLES `evento_models` WRITE;
/*!40000 ALTER TABLE `evento_models` DISABLE KEYS */;
INSERT INTO `evento_models` VALUES (1,'2019-01-23 23:24:55','2019-01-23 23:24:55',NULL,'Concierto Mana','Se realizara el concierto de Mana el 29 de febrero','29/02/2029',1,0),(2,'2019-01-23 23:25:13','2019-01-23 23:25:13',NULL,'Concierto Paramore','Se realizara el concierto de Mana el 22 de febrero','22/02/2029',1,0),(3,'2019-01-23 23:25:35','2019-01-23 23:25:35',NULL,'Concierto Twenty One Pilots','Se realizara el concierto de TOP el 31 de febrero','31/02/2029',1,0);
/*!40000 ALTER TABLE `evento_models` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `localidad_models`
--

DROP TABLE IF EXISTS `localidad_models`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
 SET character_set_client = utf8mb4 ;
CREATE TABLE `localidad_models` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  `nombre` varchar(255) DEFAULT NULL,
  `precio` double DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_localidad_models_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `localidad_models`
--

LOCK TABLES `localidad_models` WRITE;
/*!40000 ALTER TABLE `localidad_models` DISABLE KEYS */;
INSERT INTO `localidad_models` VALUES (1,NULL,NULL,NULL,'General',10),(2,NULL,NULL,NULL,'Palco',20),(3,NULL,NULL,NULL,'VIP',50),(4,NULL,NULL,NULL,'GoldenBox',100);
/*!40000 ALTER TABLE `localidad_models` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `sede_models`
--

DROP TABLE IF EXISTS `sede_models`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
 SET character_set_client = utf8mb4 ;
CREATE TABLE `sede_models` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  `nombre` varchar(255) DEFAULT NULL,
  `tipo` varchar(255) DEFAULT NULL,
  `ciudad` varchar(255) DEFAULT NULL,
  `pais` varchar(255) DEFAULT NULL,
  `direccion` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_sede_models_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `sede_models`
--

LOCK TABLES `sede_models` WRITE;
/*!40000 ALTER TABLE `sede_models` DISABLE KEYS */;
/*!40000 ALTER TABLE `sede_models` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `usuario_models`
--

DROP TABLE IF EXISTS `usuario_models`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
 SET character_set_client = utf8mb4 ;
CREATE TABLE `usuario_models` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  `usuario` varchar(255) DEFAULT NULL,
  `nombre` varchar(255) DEFAULT NULL,
  `contrasena` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_usuario_models_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `usuario_models`
--

LOCK TABLES `usuario_models` WRITE;
/*!40000 ALTER TABLE `usuario_models` DISABLE KEYS */;
/*!40000 ALTER TABLE `usuario_models` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2019-01-24  6:42:25
