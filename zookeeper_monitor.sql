-- --------------------------------------------------------
-- 主机:                           127.0.0.1
-- 服务器版本:                        5.7.11-log - MySQL Community Server (GPL)
-- 服务器操作系统:                      Win64
-- HeidiSQL 版本:                  9.3.0.5051
-- --------------------------------------------------------

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET NAMES utf8 */;
/*!50503 SET NAMES utf8mb4 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;


-- 导出 zabbix_monitor 的数据库结构
CREATE DATABASE IF NOT EXISTS `zabbix_monitor` /*!40100 DEFAULT CHARACTER SET utf8 */;
USE `zabbix_monitor`;

-- 导出  表 zabbix_monitor.cluster 结构
CREATE TABLE IF NOT EXISTS `cluster` (
  `ID` int(11) NOT NULL AUTO_INCREMENT,
  `Name` varchar(200) NOT NULL,
  `Description` varchar(500) DEFAULT NULL,
  `InUser` varchar(50) NOT NULL,
  `InDate` datetime NOT NULL,
  `EditUser` varchar(50) DEFAULT NULL,
  `EditDate` datetime DEFAULT NULL,
  PRIMARY KEY (`ID`)
) ENGINE=InnoDB AUTO_INCREMENT=0 DEFAULT CHARSET=utf8;

-- 数据导出被取消选择。
-- 导出  表 zabbix_monitor.mail 结构
CREATE TABLE IF NOT EXISTS `mail` (
  `ID` int(11) NOT NULL AUTO_INCREMENT,
  `Address` varchar(100) NOT NULL DEFAULT '0',
  PRIMARY KEY (`ID`)
) ENGINE=InnoDB AUTO_INCREMENT=0 DEFAULT CHARSET=utf8;

-- 数据导出被取消选择。
-- 导出  表 zabbix_monitor.server 结构
CREATE TABLE IF NOT EXISTS `server` (
  `ID` int(11) NOT NULL AUTO_INCREMENT,
  `ClusterID` int(11) NOT NULL,
  `IP` varchar(50) NOT NULL,
  `Port` varchar(50) NOT NULL DEFAULT '8481',
  `Name` varchar(200) NOT NULL,
  `Description` varchar(500) DEFAULT NULL,
  `IsRunning` tinyint(4) DEFAULT NULL,
  `Mode` varchar(50) DEFAULT NULL,
  `InUser` varchar(50) NOT NULL,
  `InDate` datetime NOT NULL,
  `EditUser` varchar(50) DEFAULT NULL,
  `EditDate` datetime DEFAULT NULL,
  PRIMARY KEY (`ID`)
) ENGINE=InnoDB AUTO_INCREMENT=0 DEFAULT CHARSET=utf8;

-- 数据导出被取消选择。
-- 导出  表 zabbix_monitor.status 结构
CREATE TABLE IF NOT EXISTS `status` (
  `ID` int(11) NOT NULL AUTO_INCREMENT,
  `ServerID` int(11) NOT NULL,
  `Version` varchar(100) NOT NULL COMMENT '版本号',
  `AvgLatency` int(11) NOT NULL COMMENT '平均延时',
  `MaxLatency` int(11) NOT NULL COMMENT '最大延时',
  `MinLatency` int(11) NOT NULL COMMENT '最小延时',
  `PacketsReceived` int(11) NOT NULL COMMENT '收包数',
  `PacketsSend` int(11) NOT NULL COMMENT '发包数',
  `NumAliveConnections` int(11) NOT NULL COMMENT '连接数',
  `OutstandingRequests` int(11) NOT NULL COMMENT '堆积请求数',
  `ServerState` varchar(100) NOT NULL COMMENT '状态',
  `ZnodeCount` int(11) NOT NULL COMMENT 'znode数量',
  `WatchCount` int(11) NOT NULL DEFAULT '-1' COMMENT 'watch数量',
  `EphemeralsCount` int(11) NOT NULL DEFAULT '-1' COMMENT '临时节点数量',
  `ApproximateDataSize` int(11) NOT NULL DEFAULT '-1' COMMENT '数据大小',
  `OpenFileDescriptorCount` int(11) NOT NULL DEFAULT '-1' COMMENT '打开的文件描述符数量  ',
  `MaxFileDescriptorCount` int(11) NOT NULL DEFAULT '-1' COMMENT '最大文件描述符数量  ',
  `Followers` int(11) NOT NULL DEFAULT '-1' COMMENT 'follower数量  ',
  `SyncedFollowers` int(11) NOT NULL DEFAULT '-1' COMMENT '同步的follower数量  ',
  `PendingSyncs` int(11) NOT NULL DEFAULT '-1' COMMENT '准备同步数  ',
  `InDate` datetime NOT NULL COMMENT '写入时间',
  PRIMARY KEY (`ID`)
) ENGINE=InnoDB AUTO_INCREMENT=0 DEFAULT CHARSET=utf8;

-- 导出  表 zookeeper_monitor.log 结构
CREATE TABLE IF NOT EXISTS `log` (
  `ID` int(11) NOT NULL AUTO_INCREMENT,
  `Operate` varchar(500) DEFAULT NULL,
  `InUser` varchar(50) NOT NULL,
  `InDate` datetime NOT NULL,
  PRIMARY KEY (`ID`)
) ENGINE=InnoDB AUTO_INCREMENT=0 DEFAULT CHARSET=utf8;

-- 导出  事件 zookeeper_monitor.Delete_Status 结构
DELIMITER //
CREATE EVENT `Delete_Status` ON SCHEDULE EVERY 1 DAY STARTS '2016-03-08 16:00:00' ON COMPLETION PRESERVE ENABLE DO BEGIN  
	START TRANSACTION;
		DELETE FROM status WHERE InDate < (CURRENT_TIMESTAMP()+INTERVAL -1 DAY);
		INSERT INTO log(Operate, InUser, InDate) VALUES('Delete status', 'job', now());
	COMMIT;
END//
DELIMITER ;

-- 数据导出被取消选择。
/*!40101 SET SQL_MODE=IFNULL(@OLD_SQL_MODE, '') */;
/*!40014 SET FOREIGN_KEY_CHECKS=IF(@OLD_FOREIGN_KEY_CHECKS IS NULL, 1, @OLD_FOREIGN_KEY_CHECKS) */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
