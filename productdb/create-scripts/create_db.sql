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

SET character_set_client = utf8;
SET character_set_connection = utf8;
SET character_set_results = utf8;
SET collation_connection = utf8_general_ci;

INSERT INTO products(id, name, type, quantity) VALUES ('c58b4e85-600b-4171-91ff-80af251ddeab', 'Camisa do Grêmio', 'clothing', 100);
INSERT INTO products(id, name, type, quantity) VALUES ('8a6a92c3-9417-4d34-b98f-9763894cdc8a', 'Capim Dourado', 'plant', 42);
INSERT INTO products(id, name, type, quantity) VALUES ('d9c0e7d1-113f-47d8-b33d-3f184b2902e3', 'CD do Atitude 67', 'music', 146);
INSERT INTO products(id, name, type, quantity) VALUES ('4f040e09-ac9c-4a53-8fa0-0afbad8d17f6', 'Flash 165', 'boat', 67);
INSERT INTO products(id, name, type, quantity) VALUES ('5bd235f3-2d1c-4270-aa12-b2266ac30fad', 'Bandana Dazaranha', 'clothing', 98);
INSERT INTO products(id, name, type, quantity) VALUES ('827c4788-a0ba-4431-b7b6-6c127855c1ec', 'Motul 5w40', 'oil', 34);