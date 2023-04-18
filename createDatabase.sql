CREATE SCHEMA `fbla` ;
use fbla;
CREATE TABLE `Grades`(
    `ID` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    `GradeLevel` TINYINT NOT NULL,
    `RandomWinner` MEDIUMINT UNSIGNED NOT NULL,
    PRIMARY KEY (`ID`)
);
CREATE TABLE `Users`(
    `UserID` MEDIUMINT UNSIGNED NOT NULL,
    `StudentName` VARCHAR(255) NOT NULL,
    `Password` VARCHAR(255) NOT NULL,
    `GradeID` BIGINT UNSIGNED NOT NULL,
    PRIMARY KEY(`UserID`),
    FOREIGN KEY(`GradeID`) REFERENCES Grades(ID)
);
CREATE TABLE `Prizes`(
    `ID` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    `PrizeName` VARCHAR(255) NOT NULL,
    `PointThreshold` SMALLINT UNSIGNED NOT NULL,
    PRIMARY KEY(`ID`)
);
CREATE TABLE `UserPrizes`(
    `PrizeID` BIGINT UNSIGNED NOT NULL,
    `UserID` MEDIUMINT UNSIGNED NOT NULL,
    PRIMARY KEY(PrizeID, UserID),
	FOREIGN KEY(PrizeID) references Prizes(ID),
	FOREIGN KEY(UserID) references Users(UserID)
);
CREATE TABLE `Sports`(
    `ID` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    `SportName` VARCHAR(255) NOT NULL,
    `SportDescription` TEXT NOT NULL,
    PRIMARY KEY(`ID`)
);
CREATE TABLE `Events`(
    `EventID` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    `Points` SMALLINT UNSIGNED NOT NULL,
    `EventDescription` TEXT NOT NULL,
    `EventDate` DATE NOT NULL,
    `RoomNumber` SMALLINT NOT NULL,
    `Location` VARCHAR(255) NOT NULL,
    `LocationDescription` TEXT NOT NULL,
    `SportID` BIGINT UNSIGNED NOT NULL,
    PRIMARY KEY (`EventID`),
    FOREIGN KEY (`SportID`) REFERENCES Sports(ID)
);
CREATE TABLE `UserEvents`(
    `UserID` MEDIUMINT UNSIGNED NOT NULL,
    `EventID` BIGINT UNSIGNED NOT NULL,
    PRIMARY KEY(UserID, EventID),
    FOREIGN KEY(EventID) references Events(EventID),
	FOREIGN KEY(UserID) references Users(UserID)
);
AlTER TABLE `Grades`
ADD FOREIGN KEY (RandomWinner) REFERENCES Users(UserID)