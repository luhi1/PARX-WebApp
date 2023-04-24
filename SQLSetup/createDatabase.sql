DROP SCHEMA `fbla`;
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
    (1, 'Teacher', 1000000, 'ypeBEsobvcr6wjGzmiPcTaeG7_gUfE5yuYB3ha_uSLs=', 5);
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
    (2, 1354252, 'true'),
    (3, 1354252, 'true');
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
    ('SportEvent1', 10000, 'SE', '1000-01-02', 1, 'Joe', 'Here', 'Here', 1, true),
    ('SportEvent2', 20000, 'SE', '2000-02-02', 2, 'Joe', 'Here', 'Here', 2, true),
    ('SportEvent3', 30000, 'SE', '3000-03-03', 3, 'Joe', 'Here', 'Here', 3, false),
    ('SportEvent4', 40000, 'SE', '4000-04-04', 4, 'Joe', 'Here', 'Here', 4, true),
    ('SportEvent5', 50000, 'SE', '5000-05-05', 5, 'Joe', 'Here', 'Here', 5, true),
    ('RegEvent1', 10000, 'RE', '1000-01-01', 1, 'Joe', 'Here', 'Here', 6, true),
    ('RegEvent2', 20000, 'RE', '2000-02-02', 2, 'Joe', 'Here', 'Here', 6, false),
    ('RegEvent3', 30000, 'RE', '3000-03-03', 3, 'Joe', 'Here', 'Here', 6, false),
    ('RegEvent4', 40000, 'RE', '4000-04-04', 4, 'Joe', 'Here', 'Here', 6, true),
    ('RegEvent5', 50000, 'RE', '5000-05-05', 5, 'Joe', 'Here', 'Here', 6, true);
CREATE TABLE `UserEvents`
(
    `UserID`  MEDIUMINT UNSIGNED NOT NULL,
    `EventID` BIGINT UNSIGNED    NOT NULL,
    `Attended` varchar(255),
    PRIMARY KEY (UserID, EventID),
    FOREIGN KEY (EventID) references Events (EventID),
    FOREIGN KEY (UserID) references Users (UserID)
);
insert into UserEvents(UserID, EventId, Attended)
values
    (1354252, 1, 'true'),
    (1, 1, 'true'),
    (1354252, 2, 'false'),
    (1354252, 3, 'false'),
    (1354252, 4, 'false'),
    (1354252, 5, 'true'),
    (1354252, 6, 'true'),
    (1354252, 7, 'false'),
    (1354252, 8, 'true'),
    (1354252, 9, 'true'),
    (1354252, 10, 'true');
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