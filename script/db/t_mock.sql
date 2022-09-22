CREATE TABLE `t_mock` (
                          `id` int NOT NULL AUTO_INCREMENT,
                          `name` varchar(100) NOT NULL,
                          `mock` varchar(10240) NOT NULL,
                          `enable` tinyint(1) NOT NULL DEFAULT '1',
                          `remark` varchar(1024) NOT NULL,
                          PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;