CREATE TABLE `parking_details` (
  `parking_floor` varchar(20) DEFAULT NULL,
  `block` varchar(10) DEFAULT NULL,
  `occupancy_type` varchar(24) DEFAULT NULL,
  `units` int(5) DEFAULT '1',
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `created_date` datetime DEFAULT NULL,
  `last_modified` datetime DEFAULT NULL,
  `priority` int(4) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;


CREATE TABLE `parking_entries` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `vehicle_no` varchar(20) NOT NULL,
  `entry_time` datetime DEFAULT NULL,
  `alloted_slot` int(11) DEFAULT NULL,
  `exit_time` datetime DEFAULT NULL,
  `status` varchar(20) DEFAULT NULL,
  `is_active` tinyint(1) DEFAULT '1',
  `vehicle_type` varchar(24) DEFAULT NULL,
  `created_date` datetime DEFAULT NULL,
  `last_modified` datetime DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `alloted_slot` (`alloted_slot`),
  CONSTRAINT `entries_ibfk_1` FOREIGN KEY (`alloted_slot`) REFERENCES `parking_details` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=latin1;


/* Sample Data for testing */

INSERT INTO `parking_details` ( `parking_floor`, `block`, `occupancy_type`, `units`, `created_date`, `last_modified`)
VALUES
	( 'level1', 'blockA', 'two_wheeler', 10, curtime(), curtime()),
	( 'level1', 'blockB', 'four_wheeler', 15, curtime(),curtime()),
	('level1', 'blockC', 'two_wheeler',5, curtime(),curtime()),
	('level1', 'blockC', 'four_wheeler', 8, curtime(),curtime()),
	('level2', 'blockA', 'four_wheeler', 20, curtime(),curtime())



INSERT INTO `parking_entries` ( `vehicle_no`, `entry_time`, `alloted_slot`, `exit_time`, `status`, `is_active`, `vehicle_type`, `created_date`, `last_modified`)
VALUES
	( 'BC1235678', '2020-11-19 16:27:36', 1, NULL, 'active', 1, 'two_wheeler', '2020-11-19 16:27:36', '2020-11-19 16:27:36');

INSERT INTO `parking_entries` (`vehicle_no`, `entry_time`, `alloted_slot`, `exit_time`, `status`, `is_active`, `vehicle_type`, `created_date`, `last_modified`)
VALUES
	( 'BC123567', '2020-11-19 14:17:43', 1, NULL, 'active', 1, 'two_wheeler', '2020-11-19 14:17:43', '2020-11-19 14:17:43');

INSERT INTO `parking_entries` ( `vehicle_no`, `entry_time`, `alloted_slot`, `exit_time`, `status`, `is_active`, `vehicle_type`, `created_date`, `last_modified`)
VALUES
	( 'AB53EC2', '2020-11-19 14:16:44', 1, NULL, 'active', 1, 'two_wheeler', '2020-11-19 14:17:40', NULL);







