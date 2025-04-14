SET SQL_MODE='ORACLE';

CREATE SCHEMA `fbla`;
use fbla;
CREATE TABLE `Grades`
(
    `ID`           BIGINT AUTO_INCREMENT,
    `GradeLevel`   TINYINT            NOT NULL,
    `RandomWinner` MEDIUMINT UNSIGNED NOT NULL,
    PRIMARY KEY (`ID`)
);
insert into grades(GradeLevel, RandomWinner)
values
    (9, 1354252),
    (10, 1354252),
    (11, 1354252),
    (12, 1354252),
    (0, 1);
CREATE TABLE `Users`
(
    `UserID`      MEDIUMINT UNSIGNED NOT NULL,
    `StudentName` VARCHAR(255)       NOT NULL,
    `Password`    VARCHAR(255)       NOT NULL,
    `Points` int UNSIGNED NOT NULL,
    `GradeID`     BIGINT not null,
    PRIMARY KEY (`UserID`)
);
insert into users(UserID, StudentName, `Points`, Password, GradeID)
values
    (1354252, 'Michael', 10000, 'ypeBEsobvcr6wjGzmiPcTaeG7_gUfE5yuYB3ha_uSLs=', 2),
    (1, 'Teacher', 1000000, 'ypeBEsobvcr6wjGzmiPcTaeG7_gUfE5yuYB3ha_uSLs=', 5),
    (123, 'Joe', 10300, 'ypeBEsobvcr6wjGzmiPcTaeG7_gUfE5yuYB3ha_uSLs=', 2),
    (223, 'Miguel', 2100, 'ypeBEsobvcr6wjGzmiPcTaeG7_gUfE5yuYB3ha_uSLs=', 2),
    (323, 'Alex', 11000, 'ypeBEsobvcr6wjGzmiPcTaeG7_gUfE5yuYB3ha_uSLs=', 2),
    (113, 'Nathan', 1900, 'ypeBEsobvcr6wjGzmiPcTaeG7_gUfE5yuYB3ha_uSLs=', 1),
    (213, 'Noah', 1800, 'ypeBEsobvcr6wjGzmiPcTaeG7_gUfE5yuYB3ha_uSLs=', 1),
    (313, 'Rae', 9000, 'ypeBEsobvcr6wjGzmiPcTaeG7_gUfE5yuYB3ha_uSLs=', 1),
    (143, 'Hector', 20000, 'ypeBEsobvcr6wjGzmiPcTaeG7_gUfE5yuYB3ha_uSLs=', 4),
    (243, 'Chad', 15000, 'ypeBEsobvcr6wjGzmiPcTaeG7_gUfE5yuYB3ha_uSLs=', 4),
    (343, 'Vishnu', 8000, 'ypeBEsobvcr6wjGzmiPcTaeG7_gUfE5yuYB3ha_uSLs=', 4),
    (133, 'Kevin', 7000, 'ypeBEsobvcr6wjGzmiPcTaeG7_gUfE5yuYB3ha_uSLs=', 3),
    (233, 'Harold', 5000, 'ypeBEsobvcr6wjGzmiPcTaeG7_gUfE5yuYB3ha_uSLs=', 3),
    (332, 'Carlo', 10000, 'ypeBEsobvcr6wjGzmiPcTaeG7_gUfE5yuYB3ha_uSLs=', 3);
CREATE TABLE `Prizes`
(
    `ID`             BIGINT UNSIGNED   NOT NULL AUTO_INCREMENT,
    `PrizeName`      VARCHAR(255)      NOT NULL,
    `PointThreshold` SMALLINT UNSIGNED NOT NULL,
    PRIMARY KEY (`ID`)
);
insert into prizes(PrizeName, PointThreshold)
values
    ('Ice Cream', 100),
    ('School Camera', 500),
    ('School Hoodie', 10000);
CREATE TABLE `UserPrizes`
(
    `PrizeID` BIGINT UNSIGNED    NOT NULL,
    `UserID`  MEDIUMINT UNSIGNED NOT NULL,
    `Attended` varchar(255),
    PRIMARY KEY (PrizeID, UserID),
    FOREIGN KEY (PrizeID) references Prizes (ID),
    FOREIGN KEY (UserID) references Users (UserID)
);
insert into userprizes(PrizeID, UserID, Attended)
values
    (1, 1354252, 'true'),
    (1, 123, 'true'),
    (1, 223, 'false'),
    (1, 323, 'true'),
    (2, 1354252, 'false'),
    (2, 113, 'true'),
    (2, 213, 'false'),
    (2, 313, 'true'),
    (2, 143, 'true'),
    (3, 243, 'true'),
    (3, 343, 'true'),
    (3, 133, 'true'),
    (3, 233, 'true'),

    (3, 332, 'true');
CREATE TABLE `Sports`
(
    `ID`               BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    `SportName`        VARCHAR(255)    NOT NULL,
    `SportDescription` TEXT            NOT NULL,
    PRIMARY KEY (`ID`)
);
insert into sports(SportName, SportDescription)
values
    ('Football', 'Throw ball fast.'),
    ('Soccer', 'Kick ball fast.'),
    ('Rugby', 'Throw ball laterally.'),
    ('Basketball', 'Dribble ball fast'),
    ('Hockey', 'Slide puck while cold'),
    ('Non-Sporting', 'Non-Sporting');
CREATE TABLE `Events`
(
    `EventID`             BIGINT UNSIGNED   NOT NULL AUTO_INCREMENT,
    `EventName`            VARCHAR(255)      NOT NULL,
    `Points`              SMALLINT UNSIGNED NOT NULL,
    `EventDescription`    TEXT              NOT NULL,
    `EventDate`           DATE              NOT NULL,
    `RoomNumber`          SMALLINT          NOT NULL,
    `Advisors` varchar(255) not null,
    `Location`            VARCHAR(255)      NOT NULL,
    `LocationDescription` TEXT              NOT NULL,
    `SportID`             BIGINT UNSIGNED   NOT NULL,
    `Active` bool,
    PRIMARY KEY (`EventID`),
    FOREIGN KEY (`SportID`) REFERENCES Sports (ID)
);
insert into events(EventName, Points, EventDescription, EventDate, RoomNumber, Advisors, Location, LocationDescription, SportID, Active)
values
    ('Matchup 1', 10000, 'Eldo v Atech', '1000-01-02', 1, 'Joe', 'Here', 'Here', 1, true),
    ('Matchup 2', 20000, 'Eldo v Clark', '2000-02-02', 2, 'Joe', 'Here', 'Here', 2, true),
    ('SportEvent3', 30000, 'SE', '3000-03-03', 3, 'Joe', 'Here', 'Here', 3, false),
    ('Matchup 4', 40000, 'Clark V. WCTA', '4000-04-04', 4, 'Joe', 'Here', 'Here', 4, true),
    ('Matchup 5', 50000, 'WCTA V. SECTA', '5000-05-05', 5, 'Joe', 'Here', 'Here', 5, true),
    ('Library Meetup', 10000, 'Wow so fun!', '1000-01-01', 1, 'Joe', 'Here', 'Here', 6, true),
    ('Library Burning', 20000, 'WOW SO FUN!', '2000-02-02', 2, 'Joe', 'Here', 'Here', 6, false),
    ('Library Reconciling', 30000, 'sorry.', '3000-03-03', 3, 'Joe', 'Here', 'Here', 6, false),
    ('Jack Hangout', 40000, 'His mom is onto us', '4000-04-04', 4, 'Joe', 'Here', 'Here', 6, true),
    ('Carlo and Michael Study', 50000, ';)', '5000-05-05', 5, 'Joe', 'Here', 'Here', 6, true);
CREATE TABLE `UserEvents`
(
    `UserID`  MEDIUMINT UNSIGNED NOT NULL,
    `EventID` BIGINT UNSIGNED    NOT NULL,
    `Attended` varchar(255),
    PRIMARY KEY (UserID, EventID),
    FOREIGN KEY (EventID) references Events (EventID),
    FOREIGN KEY (UserID) references Users (UserID)
);
create table bugs(
    `Bugs` text,
    `ID` BIGINT AUTO_INCREMENT,
    PRIMARY KEY (ID)
);
insert into UserEvents(UserID, EventId, Attended)
values
    (1354252, 1, 'true'),
    (123, 1, 'true'),
    (223, 1, 'true'),
    (1354252, 2, 'false'),
    (323, 2, 'false'),
    (113, 2, 'true'),
    (243, 2, 'true'),
    (143, 3, 'true'),
    (243, 3, 'true'),
    (133, 3, 'true'),
    (1354252, 3, 'false'),
    (143, 4, 'true'),
    (123, 4, 'false'),
    (1354252, 4, 'false'),
    (343, 5, 'false'),
    (133, 5, 'false'),
    (332, 5, 'false'),
    (1354252, 5, 'true'),
    (1354252, 6, 'true'),
    (323, 6, 'true'),
    (113, 6, 'false'),
    (213, 6, 'true'),
    (1354252, 7, 'false'),
    (313, 7, 'true'),
    (123, 7, 'true'),
    (243, 7, 'false'),
    (1354252, 8, 'true'),
    (1354252, 9, 'true');
AlTER TABLE `Grades`
    ADD FOREIGN KEY (RandomWinner) REFERENCES Users (UserID);
ALTER TABLE `Users`
    ADD FOREIGN KEY (`GradeID`) REFERENCES Grades (ID);
select * from grades;
select * from users;
select * from prizes;
select * from UserPrizes;
select * from sports;
select * from events;
select * from userevents;
