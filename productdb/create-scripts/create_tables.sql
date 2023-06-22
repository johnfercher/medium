USE ProductDb;

CREATE TABLE IF NOT EXISTS products(
    `id` varchar(36) NOT NULL,
    `name` varchar(100) NOT NULL,
    `type` varchar(50) NOT NULL,
    `quantity` integer NOT NULL,
    PRIMARY KEY (`id`)
    ) ENGINE = InnoDB
    DEFAULT CHARACTER SET = UTF8MB4
    COLLATE = utf8mb4_unicode_520_ci;