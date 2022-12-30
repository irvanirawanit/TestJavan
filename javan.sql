-- DROP TABLE IF EXISTS assets;
CREATE TABLE assets  (
  id int(0) NOT NULL AUTO_INCREMENT,
  name varchar(255),
  price int(0),
  PRIMARY KEY (id)
);


INSERT INTO assets (`name`,`price`) VALUES ('Samsung Universe 9', 1249);
INSERT INTO assets (name,price) VALUES ('Samsung Galaxy Book', 1499);
INSERT INTO assets (name,price) VALUES ('iPhone 9', 549);
INSERT INTO assets (name,price) VALUES ('iPhone X', 899);
INSERT INTO assets (name,price) VALUES ('Huawei P30', 499);

-- ----------------------------
-- Table structure for family
-- ----------------------------
-- DROP TABLE IF EXISTS family;
CREATE TABLE family  (
  id int(0) NOT NULL AUTO_INCREMENT,
  name varchar(255),
  gender varchar(255),
  parent int(0) NULL DEFAULT NULL,
  created_at varchar(255),
  update_at datetime(0) NULL DEFAULT NULL,
  PRIMARY KEY (id)
);

-- ----------------------------
-- Records of family
-- ----------------------------
INSERT INTO family(id,name,gender,parent,created_at,update_at) VALUES (1, 'Bani', 'laki-laki', NULL, NULL, NULL);
INSERT INTO family(id,name,gender,parent,created_at,update_at) VALUES (2, 'Budi', 'laki-laki', 1, NULL, NULL);
INSERT INTO family(id,name,gender,parent,created_at,update_at) VALUES (3, 'Nida', 'perempuan', 1, NULL, NULL);
INSERT INTO family(id,name,gender,parent,created_at,update_at) VALUES (4, 'Andi', 'laki-laki', 1, NULL, NULL);
INSERT INTO family(id,name,gender,parent,created_at,update_at) VALUES (5, 'Sigit', 'laki-laki', 1, NULL, NULL);
INSERT INTO family(id,name,gender,parent,created_at,update_at) VALUES (6, 'Hari', 'laki-laki', 2, NULL, NULL);
INSERT INTO family(id,name,gender,parent,created_at,update_at) VALUES (7, 'Siti', 'perempuan', 2, NULL, NULL);
INSERT INTO family(id,name,gender,parent,created_at,update_at) VALUES (8, 'Bila', 'perempuan', 3, NULL, NULL);
INSERT INTO family(id,name,gender,parent,created_at,update_at) VALUES (9, 'Lesti', 'perempuan', 3, NULL, NULL);
INSERT INTO family(id,name,gender,parent,created_at,update_at) VALUES (10, 'Diki', 'laki-laki', 4, NULL, NULL);
INSERT INTO family(id,name,gender,parent,created_at,update_at) VALUES (11, 'Doni', 'laki-laki', 5, NULL, NULL);
INSERT INTO family(id,name,gender,parent,created_at,update_at) VALUES (12, 'Toni', 'laki-laki', 5, NULL, NULL);

-- ----------------------------
-- Table structure for family_assets
-- ----------------------------
-- DROP TABLE IF EXISTS family_assets;
CREATE TABLE family_assets  (
  id int(0) NOT NULL AUTO_INCREMENT,
  familyid int(0) NULL DEFAULT NULL,
  assetid int(0) NULL DEFAULT NULL,
  PRIMARY KEY (id)
);

-- ----------------------------
-- Records of family_assets
-- ----------------------------
INSERT INTO family_assets(id,familyid,assetid) VALUES (1, 2, 1);
INSERT INTO family_assets(id,familyid,assetid) VALUES (2, 2, 2);
INSERT INTO family_assets(id,familyid,assetid) VALUES (3, 6, 3);
INSERT INTO family_assets(id,familyid,assetid) VALUES (4, 7, 4);
INSERT INTO family_assets(id,familyid,assetid) VALUES (5, 3, 5);
INSERT INTO family_assets(id,familyid,assetid) VALUES (6, 8, 1);
INSERT INTO family_assets(id,familyid,assetid) VALUES (7, 9, 5);
INSERT INTO family_assets(id,familyid,assetid) VALUES (8, 9, 4);
INSERT INTO family_assets(id,familyid,assetid) VALUES (9, 4, 1);
INSERT INTO family_assets(id,familyid,assetid) VALUES (10, 10, 2);
INSERT INTO family_assets(id,familyid,assetid) VALUES (11, 5, 5);
INSERT INTO family_assets(id,familyid,assetid) VALUES (12, 11, 4);

