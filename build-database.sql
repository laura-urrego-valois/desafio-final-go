-- MySQL Workbench Forward Engineering

SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0;
SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0;
SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='ONLY_FULL_GROUP_BY,STRICT_TRANS_TABLES,NO_ZERO_IN_DATE,NO_ZERO_DATE,ERROR_FOR_DIVISION_BY_ZERO,NO_ENGINE_SUBSTITUTION';

-- -----------------------------------------------------
-- Schema turnos-odontologia
-- -----------------------------------------------------
CREATE SCHEMA IF NOT EXISTS `turnos-odontologia` DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci ;
USE `turnos-odontologia`;

-- -----------------------------------------------------
-- Table `turnos-odontologia`.`dentists`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `turnos-odontologia`.`dentists` (
  `Id` INT NOT NULL AUTO_INCREMENT,
  `FirstName` VARCHAR(45) NULL DEFAULT NULL,
  `LastName` VARCHAR(45) NULL DEFAULT NULL,
  `License` VARCHAR(45) NULL DEFAULT NULL,
  PRIMARY KEY (`Id`)
)
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8mb4;

-- -----------------------------------------------------
-- Table `turnos-odontologia`.`patients`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `turnos-odontologia`.`patients` (
  `Id` INT NOT NULL AUTO_INCREMENT,
  `FirstName` VARCHAR(45) NULL DEFAULT NULL,
  `LastName` VARCHAR(45) NULL DEFAULT NULL,
  `Address` VARCHAR(45) NULL DEFAULT NULL,
  `DNI` VARCHAR(45) NULL DEFAULT NULL,
  `ReleaseDate` VARCHAR(45) NULL DEFAULT NULL,
  PRIMARY KEY (`Id`)
)
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8mb4;

-- -----------------------------------------------------
-- Table `turnos-odontologia`.`appointments`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `turnos-odontologia`.`appointments` (
  `Id` INT NOT NULL AUTO_INCREMENT,
  `Date` VARCHAR(45) NULL DEFAULT NULL,
  `Hour` VARCHAR(45) NULL DEFAULT NULL,
  `Description` VARCHAR(45) NULL DEFAULT NULL,
  `patients_Id` INT NOT NULL,
  `dentists_Id` INT NOT NULL,
  PRIMARY KEY (`Id`),
  CONSTRAINT `fk_appointments_patients`
    FOREIGN KEY (`patients_Id`)
    REFERENCES `turnos-odontologia`.`patients` (`Id`),
  CONSTRAINT `fk_appointments_dentists1`
    FOREIGN KEY (`dentists_Id`)
    REFERENCES `turnos-odontologia`.`dentists` (`Id`)
)
ENGINE = InnoDB
DEFAULT CHARACTER SET = utf8mb4;

SET SQL_MODE=@OLD_SQL_MODE;
SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS;
SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS;
