CREATE TABLE `charge` (
  `charge_id` int(11) NOT NULL AUTO_INCREMENT,
  `selling_date` date DEFAULT NULL,
  `amount` decimal(10,2) DEFAULT NULL,
  `is_deleted` tinyint(1) NOT NULL DEFAULT '0',
  `charge_type_id` int(11) DEFAULT NULL,
  PRIMARY KEY (`charge_id`)
);

CREATE TABLE `charge_type` (
  `charge_type_id` int(11) NOT NULL AUTO_INCREMENT,
  `charge_type` varchar(20) NOT NULL,
  `is_deleted` tinyint(1) DEFAULT '0',
  PRIMARY KEY (`charge_type_id`)
);

CREATE TABLE `charge_type_tax_list` (
  `tax_type_id` int(11) NOT NULL,
  `charge_type_id` int(11) NOT NULL,
  PRIMARY KEY (`tax_type_id`,`charge_type_id`)
);

CREATE TABLE `tax_type` (
  `tax_type_id` int(11) NOT NULL AUTO_INCREMENT,
  `tax_type` varchar(35) NOT NULL,
  `percentage` decimal(5,4) NOT NULL DEFAULT '0.0000',
  `is_deleted` tinyint(1) DEFAULT '0',
  PRIMARY KEY (`tax_type_id`)
);

INSERT INTO charge (
  `selling_date`,
  `amount`,
  `is_deleted`,
  `charge_type_id`
)
VALUES
("2013-12-01", 50, 0, 1), 
("2013-12-01", 20, 0, 2), 
("2013-12-02", 40, 0, 1), 
("2013-12-02", 30, 0, 3), 
("2013-12-02", 30, 1, 1), 
("2013-12-03", 10, 0, 1);


INSERT INTO charge_type (
  `charge_type_id`,
  `charge_type`,
  `is_deleted`
)
VALUES
(1, "room charge", 0), 
(2, "snack charge", 0), 
(3, "deleted charge", 1);


INSERT INTO charge_type_tax_list (
  `tax_type_id`,
  `charge_type_id`
)
VALUES
(1, 1),
(2, 1),
(3, 1),
(1, 2),
(1, 3);



INSERT INTO tax_type (
  `tax_type_id`,
  `tax_type` ,
  `percentage`,
  `is_deleted`
  
)
VALUES
(1, "GST", 0.05, 0),
(2, "HRT", 0.04, 0),
(3, "DELETED TAX", 0.10, 1);



