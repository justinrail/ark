/*ark init data to MYSQL*/

CREATE TABLE IF NOT EXISTS `Gateway` (`Id` INT(11) AUTO_INCREMENT PRIMARY KEY NOT NULL,  `UUID` VARCHAR(128) NOT NULL,`Name` VARCHAR(128) NOT NULL, `Collector` VARCHAR(128) NULL, `IP` VARCHAR(64) NULL, `State` VARCHAR(64) NULL, `Joined` TINYINT(1) NULL) ENGINE=InnoDB DEFAULT CHARSET utf8;
INSERT INTO `Gateway` (`Id`, `UUID`, `Name`, `Collector`, `IP`, `State`, `Joined`) VALUES (-1, '0000', 'Mock Gateway 1', 'stub', '127.0.0.1', 'offline',0);

CREATE TABLE IF NOT EXISTS `CoreSource` (`CoreSourceId` INT(11) AUTO_INCREMENT PRIMARY KEY NOT NULL, `GatewayId` INT(11) NOT NULL, `SourceName` VARCHAR(128) NOT NULL, `State` VARCHAR(128) NOT NULL, `UniqueId` VARCHAR(128) NULL) ENGINE=InnoDB DEFAULT CHARSET utf8;
INSERT INTO `CoreSource` (`CoreSourceId`, `GatewayId`, `State`, `SourceName`,`UniqueId`) VALUES (1, -1, "on", '1#温湿度传感器', '1');
INSERT INTO `CoreSource` (`CoreSourceId`, `GatewayId`, `State`, `SourceName`,`UniqueId`) VALUES (2, -1, "on", '2#温湿度传感器', '2');
INSERT INTO `CoreSource` (`CoreSourceId`, `GatewayId`, `State`, `SourceName`,`UniqueId`) VALUES (3, -1, "on", '3#温湿度传感器', '3');
INSERT INTO `CoreSource` (`CoreSourceId`, `GatewayId`, `State`, `SourceName`,`UniqueId`) VALUES (4, -1, "on", '4#温湿度传感器', '4');

CREATE TABLE IF NOT EXISTS `CoreDataType` (`DataTypeId` INT(11) PRIMARY KEY NOT NULL, `DataTypeName` VARCHAR(128) NOT NULL) ENGINE=InnoDB DEFAULT CHARSET utf8;
CREATE UNIQUE INDEX `UQE_coredatatype_DataTypeId` ON `CoreDataType` (`DataTypeId`);
INSERT INTO `CoreDataType` (`DataTypeId`, `DataTypeName`) VALUES (1, '整数');
INSERT INTO `CoreDataType` (`DataTypeId`, `DataTypeName`) VALUES (2, '浮点数');
INSERT INTO `CoreDataType` (`DataTypeId`, `DataTypeName`) VALUES (3, '字符串');
INSERT INTO `CoreDataType` (`DataTypeId`, `DataTypeName`) VALUES (4, '日期');
INSERT INTO `CoreDataType` (`DataTypeId`, `DataTypeName`) VALUES (5, 'JSON');

CREATE TABLE IF NOT EXISTS `CoreEventSeverity` (`EventSeverityId` INT(11) PRIMARY KEY NOT NULL, `SeverityName` VARCHAR(128) NOT NULL) ENGINE=InnoDB DEFAULT CHARSET utf8;
CREATE UNIQUE INDEX `UQE_coreeventseverity_EventSeverityId` ON `CoreEventSeverity` (`EventSeverityId`);
INSERT INTO `CoreEventSeverity` (`EventSeverityId`, `SeverityName`) VALUES (0, '未知');
INSERT INTO `CoreEventSeverity` (`EventSeverityId`, `SeverityName`) VALUES (1, '一级告警');
INSERT INTO `CoreEventSeverity` (`EventSeverityId`, `SeverityName`) VALUES (2, '二级告警');
INSERT INTO `CoreEventSeverity` (`EventSeverityId`, `SeverityName`) VALUES (3, '三级告警');
INSERT INTO `CoreEventSeverity` (`EventSeverityId`, `SeverityName`) VALUES (4, '四级告警');

CREATE TABLE IF NOT EXISTS `CorePoint` (`CorePointId` INT(11) AUTO_INCREMENT PRIMARY KEY NOT NULL, `PointName` VARCHAR(128) NOT NULL, `Accuracy` VARCHAR(32) NULL, `Unit` VARCHAR(32) NULL, `Max` VARCHAR(32) NULL, `Min` VARCHAR(32) NULL, `CoreSourceId` INT(11) NOT NULL, `CoreDataTypeId` INT(11) NULL, `EventSeverity` INT(11) NULL, `StateRuleId` INT(11) NULL, `Readable` TINYINT(1) NULL, `Writable` TINYINT(1) NULL, `DefaultValue` VARCHAR(32) NULL, `Step` FLOAT NULL, `Masked` TINYINT(1) NULL, `OriginStandardId` VARCHAR(64) NULL, `Cron` VARCHAR(128) NULL, `Expression` VARCHAR(255) NULL, `UniqueId` VARCHAR(128) NULL, `StandardId` INT(11) NULL) ENGINE=InnoDB DEFAULT CHARSET utf8;
CREATE UNIQUE INDEX `UQE_corepoint_CorePointId` ON `CorePoint` (`CorePointId`);
CREATE INDEX `IDX_Corepoint_1` ON `CorePoint` (`CoreSourceId`);
INSERT INTO `CorePoint` (`CorePointId`, `PointName`, `Accuracy`, `Unit`, `Max`, `Min`, `CoreSourceId`, `CoreDataTypeId`, `EventSeverity`, `StateRuleId`, `Readable`, `Writable`, `DefaultValue`, `Step`, `Masked`, `OriginStandardId`,`Cron`, `Expression`,`UniqueId`,`StandardId`) VALUES (1, '温度', '00.00', '℃', '35', '15', 1, 2, 0, 0, 1, 0, NULL, 0.1, 0, '1006001001',NULL,NULL,NULL,1001001);
INSERT INTO `CorePoint` (`CorePointId`, `PointName`, `Accuracy`, `Unit`, `Max`, `Min`, `CoreSourceId`, `CoreDataTypeId`, `EventSeverity`, `StateRuleId`, `Readable`, `Writable`, `DefaultValue`, `Step`, `Masked`, `OriginStandardId`,`Cron`, `Expression`,`UniqueId`,`StandardId`) VALUES (2, '温度', '00.00', '℃', '35', '15', 2, 2, 0, 0, 1, 0, NULL, 0.1, 0, '1006001001','@every 3s',NULL,NULL,1001002);
INSERT INTO `CorePoint` (`CorePointId`, `PointName`, `Accuracy`, `Unit`, `Max`, `Min`, `CoreSourceId`, `CoreDataTypeId`, `EventSeverity`, `StateRuleId`, `Readable`, `Writable`, `DefaultValue`, `Step`, `Masked`, `OriginStandardId`,`Cron`, `Expression`,`UniqueId`,`StandardId`) VALUES (3, '温度', '00.00', '℃', '35', '15', 3, 2, 0, 0, 1, 0, NULL, 0.1, 0, '1006001001','@every 3s',NULL,NULL,1001003);
INSERT INTO `CorePoint` (`CorePointId`, `PointName`, `Accuracy`, `Unit`, `Max`, `Min`, `CoreSourceId`, `CoreDataTypeId`, `EventSeverity`, `StateRuleId`, `Readable`, `Writable`, `DefaultValue`, `Step`, `Masked`, `OriginStandardId`,`Cron`, `Expression`,`UniqueId`,`StandardId`) VALUES (4, '温度', '00.00', '℃', '35', '15', 4, 2, 0, 0, 1, 0, NULL, 0.1, 0, '1006001001',NULL,NULL,NULL,1001004);
INSERT INTO `CorePoint` (`CorePointId`, `PointName`, `Accuracy`, `Unit`, `Max`, `Min`, `CoreSourceId`, `CoreDataTypeId`, `EventSeverity`, `StateRuleId`, `Readable`, `Writable`, `DefaultValue`, `Step`, `Masked`, `OriginStandardId`,`Cron`, `Expression`,`UniqueId`,`StandardId`) VALUES (51, '湿度', '00.00', 'RH', '100', '20', 1, 2, 0, 0, 1, 0, NULL, 0.1, 0, '1006003001','@every 5s',NULL,NULL,1001005);
INSERT INTO `CorePoint` (`CorePointId`, `PointName`, `Accuracy`, `Unit`, `Max`, `Min`, `CoreSourceId`, `CoreDataTypeId`, `EventSeverity`, `StateRuleId`, `Readable`, `Writable`, `DefaultValue`, `Step`, `Masked`, `OriginStandardId`,`Cron`, `Expression`,`UniqueId`,`StandardId`) VALUES (52, '湿度', '00.00', 'RH', '100', '20', 2, 2, 0, 0, 1, 0, NULL, 0.1, 0, '1006003001',NULL,NULL,NULL,1004205);
INSERT INTO `CorePoint` (`CorePointId`, `PointName`, `Accuracy`, `Unit`, `Max`, `Min`, `CoreSourceId`, `CoreDataTypeId`, `EventSeverity`, `StateRuleId`, `Readable`, `Writable`, `DefaultValue`, `Step`, `Masked`, `OriginStandardId`,`Cron`, `Expression`,`UniqueId`,`StandardId`) VALUES (53, '湿度', '00.00', 'RH', '100', '20', 3, 2, 0, 0, 1, 0, NULL, 0.1, 0, '1006003001',NULL,NULL,NULL,1004205);
INSERT INTO `CorePoint` (`CorePointId`, `PointName`, `Accuracy`, `Unit`, `Max`, `Min`, `CoreSourceId`, `CoreDataTypeId`, `EventSeverity`, `StateRuleId`, `Readable`, `Writable`, `DefaultValue`, `Step`, `Masked`, `OriginStandardId`,`Cron`, `Expression`,`UniqueId`,`StandardId`) VALUES (54, '湿度', '00.00', 'RH', '100', '20', 4, 2, 0, 0, 1, 0, NULL, 0.1, 0, '1006003001',NULL,NULL,NULL,1004205);
INSERT INTO `CorePoint` (`CorePointId`, `PointName`, `Accuracy`, `Unit`, `Max`, `Min`, `CoreSourceId`, `CoreDataTypeId`, `EventSeverity`, `StateRuleId`, `Readable`, `Writable`, `DefaultValue`, `Step`, `Masked`, `OriginStandardId`,`Cron`, `Expression`,`UniqueId`,`StandardId`) VALUES (26, '压缩机告警状态', '00.00', '℃', '1', '0', 1, 1, 3, 2, 1, 0, NULL, 0.1, 0, '702900002',NULL,NULL,NULL,1304205);
INSERT INTO `CorePoint` (`CorePointId`, `PointName`, `Accuracy`, `Unit`, `Max`, `Min`, `CoreSourceId`, `CoreDataTypeId`, `EventSeverity`, `StateRuleId`, `Readable`, `Writable`, `DefaultValue`, `Step`, `Masked`, `OriginStandardId`,`Cron`, `Expression`,`UniqueId`,`StandardId`) VALUES (29, '压缩机告警状态', '00.00', '℃', '1', '0', 2, 1, 3, 2, 1, 0, NULL, 0.1, 0, '702900002',NULL,NULL,NULL,1304205);
INSERT INTO `CorePoint` (`CorePointId`, `PointName`, `Accuracy`, `Unit`, `Max`, `Min`, `CoreSourceId`, `CoreDataTypeId`, `EventSeverity`, `StateRuleId`, `Readable`, `Writable`, `DefaultValue`, `Step`, `Masked`, `OriginStandardId`,`Cron`, `Expression`,`UniqueId`,`StandardId`) VALUES (32, '压缩机告警状态', '00.00', '℃', '1', '0', 3, 1, 3, 2, 1, 0, NULL, 0.1, 0, '702900002',NULL,NULL,NULL,1300522);
INSERT INTO `CorePoint` (`CorePointId`, `PointName`, `Accuracy`, `Unit`, `Max`, `Min`, `CoreSourceId`, `CoreDataTypeId`, `EventSeverity`, `StateRuleId`, `Readable`, `Writable`, `DefaultValue`, `Step`, `Masked`, `OriginStandardId`,`Cron`, `Expression`,`UniqueId`,`StandardId`) VALUES (35, '压缩机告警状态', '00.00', '℃', '1', '0', 4, 1, 3, 2, 1, 0, NULL, 0.1, 0, '702900002',NULL,NULL,NULL,1300522);

CREATE TABLE IF NOT EXISTS `ComplexIndex` (
`ComplexIndexId` INT(11) PRIMARY KEY NOT NULL, 
`ComplexIndexName` VARCHAR(128) NOT NULL,
`Category` VARCHAR(128) NULL,
`Title` VARCHAR(128) NULL,
`Label` VARCHAR(128) NULL,
`BusinessID` VARCHAR(128) NULL,
`ObjectTypeId` INT(11) NULL,
`GlobalResourceID` INT(11) NULL,
`CalcCron` VARCHAR(128) NOT NULL,
`CalcType` INT(11) NULL,
`AfterCalc` VARCHAR(3096) NULL,
`SaveCron` VARCHAR(128) NULL,
`Expression` VARCHAR(3096) NOT NULL,
`Unit` VARCHAR(128) NULL,
`Remark` VARCHAR(255) NULL
) ENGINE=InnoDB DEFAULT CHARSET utf8;

INSERT INTO `ComplexIndex` (`ComplexIndexId`, `ComplexIndexName`,`Category`,`Title`,`Label`,`BusinessID`, `ObjectTypeId`, `GlobalResourceID`,`CalcCron`, `CalcType`,`AfterCalc`,`SaveCron`,`Expression`, `Unit`, `Remark` ) 
VALUES (-1, 'AVGTEMP','模拟Collector','平均温度','测试','powerfee',7,null, '@every 5s',12,'powerfee(selfval())','@every 30s','avg(cp(1),cp(2),cp(3))','摄氏度','' );