CREATE TABLE `t_mock` (
                          `id` int NOT NULL AUTO_INCREMENT,
                          `user` varchar(100) NOT NULL,
                          `mid` varchar(100) NOT NULL,
                          `method` varchar(10) NOT NULL,
                          `name` varchar(100) NOT NULL,
                          `mock` varchar(10240) NOT NULL,
                          `path` varchar(100) NOT NULL,
                          `enable` tinyint(1) NOT NULL DEFAULT '1',
                          `params` varchar(1024) NOT NULL,
                          `remark` varchar(1024) NOT NULL,
                          PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE UNIQUE INDEX t_mock_mid_IDX USING BTREE ON t_mock (mid);
