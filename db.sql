DROP DATABASE IF EXISTS library;

CREATE DATABASE library;

USE library;

CREATE TABLE `tbl_book` (
  `book_id` int(11) NOT NULL AUTO_INCREMENT,
  `book_name` varchar(45) NOT NULL UNIQUE,
  `book_author` varchar(45) NOT NULL,
  PRIMARY KEY (`book_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `tbl_student` (
  `student_id` int(11) NOT NULL AUTO_INCREMENT,
  `student_name` varchar(45) NOT NULL UNIQUE,
  `student_age` int(11) NOT NULL,
  `street_address` varchar(45) DEFAULT NULL,
  `street_address_2` varchar(45) DEFAULT NULL,
  `city` varchar(45) DEFAULT NULL,
  `state` varchar(45) DEFAULT NULL,
  `zip_code` varchar(45) DEFAULT NULL,
  `country` varchar(45) DEFAULT NULL,
  PRIMARY KEY (`student_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `tbl_student_book` (
  `student_book_id` int(11) NOT NULL AUTO_INCREMENT,
  `student_id` int(11) NOT NULL,
  `book_id` int(11) NOT NULL,
  KEY `fk_student` (`student_id`),
  KEY `fk_book` (`book_id`),
  PRIMARY KEY (`student_book_id`),
  CONSTRAINT `fk_student` FOREIGN KEY (`student_id`) REFERENCES `tbl_student` (`student_id`) ON UPDATE CASCADE,
  CONSTRAINT `fk_book` FOREIGN KEY (`book_id`) REFERENCES `tbl_book` (`book_id`) ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
