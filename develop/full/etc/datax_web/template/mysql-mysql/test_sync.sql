
DROP TABLE IF EXISTS `erp`;
CREATE TABLE `erp` (
                       `id` int NOT NULL AUTO_INCREMENT,
                       `content` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
                       `add_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
                       `erp_custom_content` varchar(255) COLLATE utf8mb4_general_ci DEFAULT NULL,
                       PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

DROP TABLE IF EXISTS `mes`;
CREATE TABLE `mes` (
                       `id` int NOT NULL AUTO_INCREMENT,
                       `content` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
                       `add_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
                       `update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                       `mes_custom_content` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
                       PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

DROP TABLE IF EXISTS `sync`;
CREATE TABLE `sync` (
                        `id` int NOT NULL AUTO_INCREMENT,
                        `erp_content` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
                        `erp_add_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
                        `mes_content` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
                        `mes_add_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
                        `update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                        `sync_custom_content` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL,
                        PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;