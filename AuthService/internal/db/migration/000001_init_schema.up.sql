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
INSERT INTO cho_tot.categories (id,cat_name) VALUES (1,'B???t ?????ng s???n');
INSERT INTO `cho_tot`.`categories` (`id`,`cat_name`) VALUES (2,'Xe c???');
INSERT INTO `cho_tot`.`categories` (`id`,`cat_name`) VALUES (3,'????? ??i???n t???');
INSERT INTO `cho_tot`.`categories` (`id`,`cat_name`) VALUES (4,'Vi???c l??m');
INSERT INTO `cho_tot`.`categories` (`id`,`cat_name`) VALUES (5,'Th?? c??ng');
INSERT INTO `cho_tot`.`categories` (`id`,`cat_name`) VALUES (6,'????? ??n, th???c ph???m v?? c??c lo???i kh??c');
INSERT INTO `cho_tot`.`categories` (`id`,`cat_name`) VALUES (7,'T??? l???nh, m??y l???nh, m??y gi???t');
INSERT INTO `cho_tot`.`categories` (`id`,`cat_name`) VALUES (8,'????? gia d???ng, n???i th???t, c??y c???nh');
INSERT INTO `cho_tot`.`categories` (`id`,`cat_name`) VALUES (9,'M??? v?? b??');
INSERT INTO `cho_tot`.`categories` (`id`,`cat_name`) VALUES (10,'Th???i trang, ????? d??ng c?? nh??n');
INSERT INTO `cho_tot`.`categories` (`id`,`cat_name`) VALUES (11,'Gi???i tr??, Th??? thao, S??? th??ch');
INSERT INTO `cho_tot`.`categories` (`id`,`cat_name`) VALUES (12,'????? d??ng v??n ph??ng, c??ng n??ng nghi???p');
INSERT INTO `cho_tot`.`categories` (`id`,`cat_name`) VALUES (13,'D???ch v???, Du l???ch');

-- add subproduct
SELECT * FROM `cho_tot`.`sub_categories` LIMIT 1000;
INSERT INTO `cho_tot`.`sub_categories` (`id`,`type_name`,`cat_id`) VALUES (1,'C??n h???/Chung c??', 1);
INSERT INTO `cho_tot`.`sub_categories` (`id`,`type_name`,`cat_id`) VALUES (2,'Nh?? ???', 1);
INSERT INTO `cho_tot`.`sub_categories` (`id`,`type_name`,`cat_id`) VALUES (3,'??????t', 1);
INSERT INTO `cho_tot`.`sub_categories` (`id`,`type_name`,`cat_id`) VALUES (4,'V??n ph??ng, M???t b???ng kinh doanh', 1);
INSERT INTO `cho_tot`.`sub_categories` (`id`,`type_name`,`cat_id`) VALUES (5,'Ph??ng tr???', 1);

INSERT INTO `cho_tot`.`sub_categories` (`id`,`type_name`,`cat_id`) VALUES (6,'?? t??', '2');
INSERT INTO `cho_tot`.`sub_categories` (`id`,`type_name`,`cat_id`) VALUES (7,'Xe m??y', '2');
INSERT INTO `cho_tot`.`sub_categories` (`id`,`type_name`,`cat_id`) VALUES (8,'Xe t???i, xe ben', '2');
INSERT INTO `cho_tot`.`sub_categories` (`id`,`type_name`,`cat_id`) VALUES (9,'Xe ??i???n', '2');
INSERT INTO `cho_tot`.`sub_categories` (`id`,`type_name`,`cat_id`) VALUES (10,'Xe ?????p', '2');
INSERT INTO `cho_tot`.`sub_categories` (`id`,`type_name`,`cat_id`) VALUES (11,'Ph????ng ti???n kh??c', '2');
INSERT INTO `cho_tot`.`sub_categories` (`id`,`type_name`,`cat_id`) VALUES (12,'Ph??? t??ng xe', '2');

INSERT INTO `cho_tot`.`sub_categories` (`id`,`type_name`,`cat_id`) VALUES (13,'??i???n tho???i', '3');
INSERT INTO `cho_tot`.`sub_categories` (`id`,`type_name`,`cat_id`) VALUES (14,'M??y t??nh b???ng', '3');
INSERT INTO `cho_tot`.`sub_categories` (`id`,`type_name`,`cat_id`) VALUES (15,'Laptop', '3');
INSERT INTO `cho_tot`.`sub_categories` (`id`,`type_name`,`cat_id`) VALUES (16,'M??y t??nh ????? b??n', '3');
INSERT INTO `cho_tot`.`sub_categories` (`id`,`type_name`,`cat_id`) VALUES (17,'M??y ???nh, M??y quay', '3');
INSERT INTO `cho_tot`.`sub_categories` (`id`,`type_name`,`cat_id`) VALUES (18,'Tivi, ??m thanh', '3');
INSERT INTO `cho_tot`.`sub_categories` (`id`,`type_name`,`cat_id`) VALUES (19,'Thi???t b??? ??eo th??ng minh', '3');
INSERT INTO `cho_tot`.`sub_categories` (`id`,`type_name`,`cat_id`) VALUES (20,'Ph??? ki???n (M??n h??nh, Chu???t...)', '3');
INSERT INTO `cho_tot`.`sub_categories` (`id`,`type_name`,`cat_id`) VALUES (21,'Linh ki???n (RAM, Card...)', '3');

INSERT INTO `cho_tot`.`sub_categories` (`id`,`type_name`,`cat_id`) VALUES (22,'Vi????c la??m', '4');

INSERT INTO `cho_tot`.`sub_categories` (`id`,`type_name`,`cat_id`) VALUES (23,'Ga??', '5');
INSERT INTO `cho_tot`.`sub_categories` (`id`,`type_name`,`cat_id`) VALUES (24,'Cho??', '5');
INSERT INTO `cho_tot`.`sub_categories` (`id`,`type_name`,`cat_id`) VALUES (25,'Chim', '5');
INSERT INTO `cho_tot`.`sub_categories` (`id`,`type_name`,`cat_id`) VALUES (26,'Me??o', '5');
INSERT INTO `cho_tot`.`sub_categories` (`id`,`type_name`,`cat_id`) VALUES (27,'Th?? c??ng kh??c', '5');
INSERT INTO `cho_tot`.`sub_categories` (`id`,`type_name`,`cat_id`) VALUES (28,'Ph??? ki???n, Th???c ??n, D???ch v???', '5');

INSERT INTO `cho_tot`.`sub_categories` (`id`,`type_name`,`cat_id`) VALUES (29,'????? ??n, th???c ph???m v?? c??c lo???i kh??c', '6');

INSERT INTO `cho_tot`.`sub_categories` (`id`,`type_name`,`cat_id`) VALUES (30,'Ma??y la??nh, ??i????u ho??a', '7');
INSERT INTO `cho_tot`.`sub_categories` (`id`,`type_name`,`cat_id`) VALUES (31,'Ma??y gi????t', '7');
INSERT INTO `cho_tot`.`sub_categories` (`id`,`type_name`,`cat_id`) VALUES (32,'Tu?? la??nh', '7');

INSERT INTO `cho_tot`.`sub_categories` (`id`,`type_name`,`cat_id`) VALUES (33,'Gi?????ng, ch??n ga g???i n???m', '8');
INSERT INTO `cho_tot`.`sub_categories` (`id`,`type_name`,`cat_id`) VALUES (34,'B???p, l??, ????? ??i???n nh?? b???p', '8');
INSERT INTO `cho_tot`.`sub_categories` (`id`,`type_name`,`cat_id`) VALUES (35,'D???ng c??? nh?? b???p', '8');
INSERT INTO `cho_tot`.`sub_categories` (`id`,`type_name`,`cat_id`) VALUES (36,'Qua??t', '8');
INSERT INTO `cho_tot`.`sub_categories` (`id`,`type_name`,`cat_id`) VALUES (37,'??e??n', '8');
INSERT INTO `cho_tot`.`sub_categories` (`id`,`type_name`,`cat_id`) VALUES (38,'C??y c???nh, ????? trang tr??', '8');
INSERT INTO `cho_tot`.`sub_categories` (`id`,`type_name`,`cat_id`) VALUES (39,'Thi???t b??? v??? sinh, nh?? t???m', '8');
INSERT INTO `cho_tot`.`sub_categories` (`id`,`type_name`,`cat_id`) VALUES (40,'N???i th???t, ????? gia d???ng kh??c', '8');
INSERT INTO `cho_tot`.`sub_categories` (`id`,`type_name`,`cat_id`) VALUES (41,'B??n gh???', '8');

INSERT INTO `cho_tot`.`sub_categories` (`id`,`type_name`,`cat_id`) VALUES (42,'Me?? va?? be??', '9');

INSERT INTO `cho_tot`.`sub_categories` (`id`,`type_name`,`cat_id`) VALUES (43,'Qu???n ??o', '10');
INSERT INTO `cho_tot`.`sub_categories` (`id`,`type_name`,`cat_id`) VALUES (44,'?????ng h???', '10');
INSERT INTO `cho_tot`.`sub_categories` (`id`,`type_name`,`cat_id`) VALUES (45,'Gi??y d??p', '10');
INSERT INTO `cho_tot`.`sub_categories` (`id`,`type_name`,`cat_id`) VALUES (46,'T??i x??ch', '10');
INSERT INTO `cho_tot`.`sub_categories` (`id`,`type_name`,`cat_id`) VALUES (47,'N?????c hoa', '10');
INSERT INTO `cho_tot`.`sub_categories` (`id`,`type_name`,`cat_id`) VALUES (48,'Ph??? ki???n th???i trang kh??c', '10');

INSERT INTO `cho_tot`.`sub_categories` (`id`,`type_name`,`cat_id`) VALUES (49,'Nh???c c???', '11');
INSERT INTO `cho_tot`.`sub_categories` (`id`,`type_name`,`cat_id`) VALUES (50,'S??ch', '11');
INSERT INTO `cho_tot`.`sub_categories` (`id`,`type_name`,`cat_id`) VALUES (51,'????? th??? thao, D?? ngo???i', '11');
INSERT INTO `cho_tot`.`sub_categories` (`id`,`type_name`,`cat_id`) VALUES (52,'Thi???t b??? ch??i game', '11');
INSERT INTO `cho_tot`.`sub_categories` (`id`,`type_name`,`cat_id`) VALUES (53,'S??? th??ch kh??c', '11');

INSERT INTO `cho_tot`.`sub_categories` (`id`,`type_name`,`cat_id`) VALUES (54,'????? d??ng v??n ph??ng', '12');
INSERT INTO `cho_tot`.`sub_categories` (`id`,`type_name`,`cat_id`) VALUES (55,'????? chuy??n d???ng, Gi???ng nu??i tr???ng', '12');

INSERT INTO `cho_tot`.`sub_categories` (`id`,`type_name`,`cat_id`) VALUES (56,'D???ch v???', '13');
INSERT INTO `cho_tot`.`sub_categories` (`id`,`type_name`,`cat_id`) VALUES (57,'Du l???ch', '13');