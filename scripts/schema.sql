CREATE DATABASE `triviadb` /*!40100 DEFAULT CHARACTER SET latin1 */;

USE triviadb;

-- triviadb.`number` definition

CREATE TABLE `numbers` (
                           `number` varchar(200) DEFAULT NULL,
                           `type` varchar(200) NOT NULL,
                           `text` text,
                           `found` tinyint(1) DEFAULT '0'
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
