CREATE DATABASE IF NOT EXISTS cho_tot
character set "utf8mb4"
collate "utf8mb4_general_ci";
USE cho_tot;

CREATE TABLE IF NOT EXISTS users
(
    id        INT          NOT NULL PRIMARY KEY AUTO_INCREMENT,
    phone     VARCHAR(12)  NOT NULL UNIQUE,
    username  VARCHAR(50)  NOT NULL,
    passwd    VARCHAR(255) NOT NULL,
    address   VARCHAR(255) NOT NULL DEFAULT '',
    email     VARCHAR(50) NOT NULL DEFAULT '',
    isAdmin   BOOLEAN NOT NULL DEFAULT 0
);
CREATE TABLE IF NOT EXISTS products
(
    id           INT          NOT NULL PRIMARY KEY AUTO_INCREMENT,
    product_name VARCHAR(255) NOT NULL,
    user_id      INT,
    cat_id       VARCHAR(10),
    type_id      VARCHAR(10),
    price        DOUBLE(15, 2),
    state        BOOLEAN NOT NULL DEFAULT 0,
    created_time DATETIME DEFAULT CURRENT_TIMESTAMP,
    expired_time DATETIME,
    address      VARCHAR(255),
    content      VARCHAR(255),
    priority     BOOLEAN DEFAULT 0
);

CREATE TABLE IF NOT EXISTS categories
(
    id       VARCHAR(10) NOT NULL PRIMARY KEY,
    cat_name VARCHAR(50) NOT NULL UNIQUE
);

CREATE TABLE IF NOT EXISTS sub_categories
(
    id        VARCHAR(10) NOT NULL PRIMARY KEY,
    type_name VARCHAR(50) NOT NULL UNIQUE,
    cat_id    VARCHAR(10)
);

CREATE TABLE IF NOT EXISTS images
(
    id         INT NOT NULL PRIMARY KEY AUTO_INCREMENT,
    product_id INT,
    link       VARCHAR(255)
);

ALTER TABLE products
    ADD CONSTRAINT FK_Products_Users_UserId FOREIGN KEY (user_id) REFERENCES users (id);
ALTER TABLE products
    ADD CONSTRAINT FK_Products_Users_CatId FOREIGN KEY (cat_id) REFERENCES categories (id);
ALTER TABLE products
    ADD CONSTRAINT FK_Products_Users_TypeId FOREIGN KEY (type_id) REFERENCES sub_categories (id);
ALTER TABLE sub_categories
    ADD CONSTRAINT FK_SubCategories_Categories_CatId FOREIGN KEY (cat_id) REFERENCES categories (id);
ALTER TABLE images
    ADD CONSTRAINT FK_Images_Products_ProductId FOREIGN KEY (product_id) REFERENCES products (id);


-- add categories
SELECT * FROM cho_tot.categories LIMIT 1000;
INSERT INTO cho_tot.categories (id,cat_name) VALUES (1,'Bất động sản');
INSERT INTO `cho_tot`.`categories` (`id`,`cat_name`) VALUES (2,'Xe cộ');
INSERT INTO `cho_tot`.`categories` (`id`,`cat_name`) VALUES (3,'Đồ điện tử');
INSERT INTO `cho_tot`.`categories` (`id`,`cat_name`) VALUES (4,'Việc làm');
INSERT INTO `cho_tot`.`categories` (`id`,`cat_name`) VALUES (5,'Thú cưng');
INSERT INTO `cho_tot`.`categories` (`id`,`cat_name`) VALUES (6,'Đồ ăn, thực phẩm và các loại khác');
INSERT INTO `cho_tot`.`categories` (`id`,`cat_name`) VALUES (7,'Tủ lạnh, máy lạnh, máy giặt');
INSERT INTO `cho_tot`.`categories` (`id`,`cat_name`) VALUES (8,'Đồ gia dụng, nội thất, cây cảnh');
INSERT INTO `cho_tot`.`categories` (`id`,`cat_name`) VALUES (9,'Mẹ và bé');
INSERT INTO `cho_tot`.`categories` (`id`,`cat_name`) VALUES (10,'Thời trang, Đồ dùng cá nhân');
INSERT INTO `cho_tot`.`categories` (`id`,`cat_name`) VALUES (11,'Giải trí, Thể thao, Sở thích');
INSERT INTO `cho_tot`.`categories` (`id`,`cat_name`) VALUES (12,'Đồ dùng văn phòng, công nông nghiệp');
INSERT INTO `cho_tot`.`categories` (`id`,`cat_name`) VALUES (13,'Dịch vụ, Du lịch');

-- add subproduct
SELECT * FROM `cho_tot`.`sub_categories` LIMIT 1000;
INSERT INTO `cho_tot`.`sub_categories` (`id`,`type_name`,`cat_id`) VALUES (1,'Căn hộ/Chung cư', 1);
INSERT INTO `cho_tot`.`sub_categories` (`id`,`type_name`,`cat_id`) VALUES (2,'Nhà ở', 1);
INSERT INTO `cho_tot`.`sub_categories` (`id`,`type_name`,`cat_id`) VALUES (3,'Đất', 1);
INSERT INTO `cho_tot`.`sub_categories` (`id`,`type_name`,`cat_id`) VALUES (4,'Văn phòng, Mặt bằng kinh doanh', 1);
INSERT INTO `cho_tot`.`sub_categories` (`id`,`type_name`,`cat_id`) VALUES (5,'Phòng trọ', 1);

INSERT INTO `cho_tot`.`sub_categories` (`id`,`type_name`,`cat_id`) VALUES (6,'Ô tô', '2');
INSERT INTO `cho_tot`.`sub_categories` (`id`,`type_name`,`cat_id`) VALUES (7,'Xe máy', '2');
INSERT INTO `cho_tot`.`sub_categories` (`id`,`type_name`,`cat_id`) VALUES (8,'Xe tải, xe ben', '2');
INSERT INTO `cho_tot`.`sub_categories` (`id`,`type_name`,`cat_id`) VALUES (9,'Xe điện', '2');
INSERT INTO `cho_tot`.`sub_categories` (`id`,`type_name`,`cat_id`) VALUES (10,'Xe đạp', '2');
INSERT INTO `cho_tot`.`sub_categories` (`id`,`type_name`,`cat_id`) VALUES (11,'Phương tiện khác', '2');
INSERT INTO `cho_tot`.`sub_categories` (`id`,`type_name`,`cat_id`) VALUES (12,'Phụ tùng xe', '2');

INSERT INTO `cho_tot`.`sub_categories` (`id`,`type_name`,`cat_id`) VALUES (13,'Điện thoại', '3');
INSERT INTO `cho_tot`.`sub_categories` (`id`,`type_name`,`cat_id`) VALUES (14,'Máy tính bảng', '3');
INSERT INTO `cho_tot`.`sub_categories` (`id`,`type_name`,`cat_id`) VALUES (15,'Laptop', '3');
INSERT INTO `cho_tot`.`sub_categories` (`id`,`type_name`,`cat_id`) VALUES (16,'Máy tính để bàn', '3');
INSERT INTO `cho_tot`.`sub_categories` (`id`,`type_name`,`cat_id`) VALUES (17,'Máy ảnh, Máy quay', '3');
INSERT INTO `cho_tot`.`sub_categories` (`id`,`type_name`,`cat_id`) VALUES (18,'Tivi, Âm thanh', '3');
INSERT INTO `cho_tot`.`sub_categories` (`id`,`type_name`,`cat_id`) VALUES (19,'Thiết bị đeo thông minh', '3');
INSERT INTO `cho_tot`.`sub_categories` (`id`,`type_name`,`cat_id`) VALUES (20,'Phụ kiện (Màn hình, Chuột...)', '3');
INSERT INTO `cho_tot`.`sub_categories` (`id`,`type_name`,`cat_id`) VALUES (21,'Linh kiện (RAM, Card...)', '3');

INSERT INTO `cho_tot`.`sub_categories` (`id`,`type_name`,`cat_id`) VALUES (22,'Việc làm', '4');

INSERT INTO `cho_tot`.`sub_categories` (`id`,`type_name`,`cat_id`) VALUES (23,'Gà', '5');
INSERT INTO `cho_tot`.`sub_categories` (`id`,`type_name`,`cat_id`) VALUES (24,'Chó', '5');
INSERT INTO `cho_tot`.`sub_categories` (`id`,`type_name`,`cat_id`) VALUES (25,'Chim', '5');
INSERT INTO `cho_tot`.`sub_categories` (`id`,`type_name`,`cat_id`) VALUES (26,'Mèo', '5');
INSERT INTO `cho_tot`.`sub_categories` (`id`,`type_name`,`cat_id`) VALUES (27,'Thú cưng khác', '5');
INSERT INTO `cho_tot`.`sub_categories` (`id`,`type_name`,`cat_id`) VALUES (28,'Phụ kiện, Thức ăn, Dịch vụ', '5');

INSERT INTO `cho_tot`.`sub_categories` (`id`,`type_name`,`cat_id`) VALUES (29,'Đồ ăn, thực phẩm và các loại khác', '6');

INSERT INTO `cho_tot`.`sub_categories` (`id`,`type_name`,`cat_id`) VALUES (30,'Máy lạnh, điều hòa', '7');
INSERT INTO `cho_tot`.`sub_categories` (`id`,`type_name`,`cat_id`) VALUES (31,'Máy giặt', '7');
INSERT INTO `cho_tot`.`sub_categories` (`id`,`type_name`,`cat_id`) VALUES (32,'Tủ lạnh', '7');

INSERT INTO `cho_tot`.`sub_categories` (`id`,`type_name`,`cat_id`) VALUES (33,'Giường, chăn ga gối nệm', '8');
INSERT INTO `cho_tot`.`sub_categories` (`id`,`type_name`,`cat_id`) VALUES (34,'Bếp, lò, đồ điện nhà bếp', '8');
INSERT INTO `cho_tot`.`sub_categories` (`id`,`type_name`,`cat_id`) VALUES (35,'Dụng cụ nhà bếp', '8');
INSERT INTO `cho_tot`.`sub_categories` (`id`,`type_name`,`cat_id`) VALUES (36,'Quạt', '8');
INSERT INTO `cho_tot`.`sub_categories` (`id`,`type_name`,`cat_id`) VALUES (37,'Đèn', '8');
INSERT INTO `cho_tot`.`sub_categories` (`id`,`type_name`,`cat_id`) VALUES (38,'Cây cảnh, đồ trang trí', '8');
INSERT INTO `cho_tot`.`sub_categories` (`id`,`type_name`,`cat_id`) VALUES (39,'Thiết bị vệ sinh, nhà tắm', '8');
INSERT INTO `cho_tot`.`sub_categories` (`id`,`type_name`,`cat_id`) VALUES (40,'Nội thất, đồ gia dụng khác', '8');
INSERT INTO `cho_tot`.`sub_categories` (`id`,`type_name`,`cat_id`) VALUES (41,'Bàn ghế', '8');

INSERT INTO `cho_tot`.`sub_categories` (`id`,`type_name`,`cat_id`) VALUES (42,'Mẹ và bé', '9');

INSERT INTO `cho_tot`.`sub_categories` (`id`,`type_name`,`cat_id`) VALUES (43,'Quần áo', '10');
INSERT INTO `cho_tot`.`sub_categories` (`id`,`type_name`,`cat_id`) VALUES (44,'Đồng hồ', '10');
INSERT INTO `cho_tot`.`sub_categories` (`id`,`type_name`,`cat_id`) VALUES (45,'Giày dép', '10');
INSERT INTO `cho_tot`.`sub_categories` (`id`,`type_name`,`cat_id`) VALUES (46,'Túi xách', '10');
INSERT INTO `cho_tot`.`sub_categories` (`id`,`type_name`,`cat_id`) VALUES (47,'Nước hoa', '10');
INSERT INTO `cho_tot`.`sub_categories` (`id`,`type_name`,`cat_id`) VALUES (48,'Phụ kiện thời trang khác', '10');

INSERT INTO `cho_tot`.`sub_categories` (`id`,`type_name`,`cat_id`) VALUES (49,'Nhạc cụ', '11');
INSERT INTO `cho_tot`.`sub_categories` (`id`,`type_name`,`cat_id`) VALUES (50,'Sách', '11');
INSERT INTO `cho_tot`.`sub_categories` (`id`,`type_name`,`cat_id`) VALUES (51,'Đồ thể thao, Dã ngoại', '11');
INSERT INTO `cho_tot`.`sub_categories` (`id`,`type_name`,`cat_id`) VALUES (52,'Thiết bị chơi game', '11');
INSERT INTO `cho_tot`.`sub_categories` (`id`,`type_name`,`cat_id`) VALUES (53,'Sở thích khác', '11');

INSERT INTO `cho_tot`.`sub_categories` (`id`,`type_name`,`cat_id`) VALUES (54,'Đồ dùng văn phòng', '12');
INSERT INTO `cho_tot`.`sub_categories` (`id`,`type_name`,`cat_id`) VALUES (55,'Đồ chuyên dụng, Giống nuôi trồng', '12');

INSERT INTO `cho_tot`.`sub_categories` (`id`,`type_name`,`cat_id`) VALUES (56,'Dịch vụ', '13');
INSERT INTO `cho_tot`.`sub_categories` (`id`,`type_name`,`cat_id`) VALUES (57,'Du lịch', '13');