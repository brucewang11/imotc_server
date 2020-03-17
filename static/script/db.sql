/*
 Navicat Premium Data Transfer

 Source Server         : 116.62.71.106
 Source Server Type    : MySQL
 Source Server Version : 50646
 Source Host           : 116.62.71.106:3306
 Source Schema         : db_myotc

 Target Server Type    : MySQL
 Target Server Version : 50646
 File Encoding         : 65001

 Date: 11/03/2020 12:39:14
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for tb_ask
-- ----------------------------
DROP TABLE IF EXISTS `tb_ask`;
CREATE TABLE `tb_ask` (
  `id` varchar(40) COLLATE utf8_bin NOT NULL,
  `user_name` varchar(100) COLLATE utf8_bin DEFAULT NULL,
  `amount` double(18,8) DEFAULT NULL,
  `min_limit` double(18,2) DEFAULT NULL,
  `max_limit` double(18,2) DEFAULT NULL,
  `price` varchar(80) COLLATE utf8_bin DEFAULT NULL,
  `pay_type` varchar(20) COLLATE utf8_bin DEFAULT NULL,
  `bank_card` varchar(50) COLLATE utf8_bin DEFAULT NULL,
  `coin_type` int(2) DEFAULT NULL COMMENT '1.usdt 2.btc .3eth 4.ltc',
  `bank_address` varchar(200) COLLATE utf8_bin DEFAULT NULL,
  `bank_name` varchar(100) COLLATE utf8_bin DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

-- ----------------------------
-- Records of tb_ask
-- ----------------------------
BEGIN;
INSERT INTO `tb_ask` VALUES ('12', '张卡卡', 10000.00000000, 100.00, 10000.00, '7.14', '银行卡', '2234 2234 3343 2234', 1, '河南省郑州市建行支行', '中国建设银行');
INSERT INTO `tb_ask` VALUES ('2', '张拉拉', 200000.00000000, 100.00, 10000.00, '7.13', '银行卡', '2234 1123 9983 3434', 1, '河南省郑州市建行支行', '中国建设银行');
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;


/*
 Navicat Premium Data Transfer

 Source Server         : 116.62.71.106
 Source Server Type    : MySQL
 Source Server Version : 50646
 Source Host           : 116.62.71.106:3306
 Source Schema         : db_myotc

 Target Server Type    : MySQL
 Target Server Version : 50646
 File Encoding         : 65001

 Date: 11/03/2020 12:39:29
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for tb_order
-- ----------------------------
DROP TABLE IF EXISTS `tb_order`;
CREATE TABLE `tb_order` (
  `id` varchar(40) COLLATE utf8_bin NOT NULL,
  `ask_id` varchar(40) COLLATE utf8_bin DEFAULT NULL,
  `amount` double(18,8) DEFAULT NULL,
  `verify_code` varchar(10) COLLATE utf8_bin DEFAULT NULL,
  `status` int(2) DEFAULT NULL COMMENT '1.支付中 2.支付成功 3.支付失败',
  `price` double(18,2) DEFAULT NULL,
  `coin_type` int(2) DEFAULT NULL COMMENT '1.usdt 2.btc.3eth.4.ltc',
  `user_id` varchar(40) COLLATE utf8_bin DEFAULT NULL,
  `user_name` varchar(100) COLLATE utf8_bin DEFAULT NULL,
  `create_time` bigint(20) DEFAULT NULL,
  `unit_price` double(10,5) DEFAULT NULL,
  `fee` double(18,8) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

-- ----------------------------
-- Records of tb_order
-- ----------------------------
BEGIN;
INSERT INTO `tb_order` VALUES ('00fce01d-b4ae-42a5-a8c9-f93b6e600673', '12', 300.00000000, 'PWYT5N', 2, 1000.00, 1, '0c468b05-c9e0-411a-b07c-0789681e46d5', '120779726@qq.com', 1583823696288, NULL, NULL);
INSERT INTO `tb_order` VALUES ('01de7638-a101-4783-9795-7e202c4ab7c3', '12', 300.00000000, 'XWGKZI', 2, 1000.00, 1, '0c468b05-c9e0-411a-b07c-0789681e46d5', '120779726@qq.com', NULL, NULL, NULL);
INSERT INTO `tb_order` VALUES ('094acf9c-4504-4529-b3f4-af2b8d32a7c8', '12', 300.00000000, 'IRLTQW', 1, 1000.00, 1, '0c468b05-c9e0-411a-b07c-0789681e46d5', '120779726@qq.com', 1583823697787, NULL, NULL);
INSERT INTO `tb_order` VALUES ('0bf66bf1-f8e4-489e-b570-39eafc8c5d4f', '12', 300.00000000, '1E7ZKV', 1, 1000.00, 1, '0c468b05-c9e0-411a-b07c-0789681e46d5', '120779726@qq.com', NULL, NULL, NULL);
INSERT INTO `tb_order` VALUES ('1004ab30-149f-471a-9831-f055934fc5d1', '12', 300.00000000, 'S7FNCR', 1, 1000.00, 1, '0c468b05-c9e0-411a-b07c-0789681e46d5', '120779726@qq.com', NULL, NULL, NULL);
INSERT INTO `tb_order` VALUES ('2a119e1c-e2da-42df-877e-3dfc5bda65c1', '12', 300.00000000, 'XNYW07', 1, 1000.00, 1, '0c468b05-c9e0-411a-b07c-0789681e46d5', '120779726@qq.com', NULL, NULL, NULL);
INSERT INTO `tb_order` VALUES ('2a446fda-8922-491a-9136-c27be20da65f', '12', 300.00000000, 'L889VY', 1, 1000.00, 1, '0c468b05-c9e0-411a-b07c-0789681e46d5', '120779726@qq.com', NULL, NULL, NULL);
INSERT INTO `tb_order` VALUES ('55be1011-c796-4d63-b0b4-3046b047222d', '12', 300.00000000, 'FFVO6T', 1, 1000.00, 1, '', '120779726@qq.com', NULL, NULL, NULL);
INSERT INTO `tb_order` VALUES ('57b400be-2403-418c-84bd-23694a6fd5b2', '12', 300.00000000, 'G1BX18', 1, 1000.00, 1, '0c468b05-c9e0-411a-b07c-0789681e46d5', '120779726@qq.com', NULL, NULL, NULL);
INSERT INTO `tb_order` VALUES ('5e137103-135d-4d30-bb1a-529744a59911', '12', 300.00000000, 'PPDC3K', 1, 1000.00, 1, '0c468b05-c9e0-411a-b07c-0789681e46d5', '120779726@qq.com', NULL, NULL, NULL);
INSERT INTO `tb_order` VALUES ('650a6d53-a6f8-449d-953b-7b78fb759a14', '12', 300.00000000, 'FMXGD1', 1, 1000.00, 1, '0c468b05-c9e0-411a-b07c-0789681e46d5', '120779726@qq.com', NULL, NULL, NULL);
INSERT INTO `tb_order` VALUES ('6a051aee-2d1f-494e-a05c-fa5396b031ea', '12', 300.00000000, '9G0U2B', 1, 1000.00, 1, '0c468b05-c9e0-411a-b07c-0789681e46d5', '120779726@qq.com', NULL, NULL, NULL);
INSERT INTO `tb_order` VALUES ('7c28c8cf-da13-4ac3-a005-5e21c61ad0f4', '12', 300.00000000, 'HM0RE8', 1, 1000.00, 1, '', '120779726@qq.com', NULL, NULL, NULL);
INSERT INTO `tb_order` VALUES ('7f46f7f5-a212-4c5d-8a45-56a165bb2188', '12', 300.00000000, 'MY996B', 1, 1000.00, 1, '', '120779726@qq.com', NULL, NULL, NULL);
INSERT INTO `tb_order` VALUES ('89dbd8a9-0559-4a47-9ab7-eeccdfedc578', '12', 300.00000000, 'ZY0ZVI', 1, 1000.00, 1, '0c468b05-c9e0-411a-b07c-0789681e46d5', '120779726@qq.com', NULL, NULL, NULL);
INSERT INTO `tb_order` VALUES ('919c1d4a-1d48-466d-8bfe-a4218c513d49', '12', 300.00000000, 'Y9HYYK', 1, 1000.00, 1, '0c468b05-c9e0-411a-b07c-0789681e46d5', '120779726@qq.com', 1583823694871, NULL, NULL);
INSERT INTO `tb_order` VALUES ('a8bf7c14-228a-478c-913c-fb72059b02fa', '12', 300.00000000, 'UR0HYN', 1, 1000.00, 1, '0c468b05-c9e0-411a-b07c-0789681e46d5', '120779726@qq.com', NULL, NULL, NULL);
INSERT INTO `tb_order` VALUES ('e455cef2-73a3-4a1e-a5cf-3bd06a1c69e6', '12', 300.00000000, '8WX4BR', 1, 1000.00, 1, '', '', NULL, NULL, NULL);
INSERT INTO `tb_order` VALUES ('e6dad9cf-4a43-4547-9876-926bc6e9b90d', '12', 300.00000000, 'QF51S7', 1, 1000.00, 1, '0c468b05-c9e0-411a-b07c-0789681e46d5', '120779726@qq.com', NULL, NULL, NULL);
INSERT INTO `tb_order` VALUES ('ee32eec7-49a0-446e-8184-c7ad3112bba4', '12', 300.00000000, 'V9SJ71', 1, 1000.00, 1, '0c468b05-c9e0-411a-b07c-0789681e46d5', '120779726@qq.com', NULL, NULL, NULL);
INSERT INTO `tb_order` VALUES ('f3cd253c-d35f-4bd8-8139-af9b90e4e26b', '12', 300.00000000, 'PGZKLY', 1, 1000.00, 1, '', '', NULL, NULL, NULL);
INSERT INTO `tb_order` VALUES ('f5c19108-c5d1-4220-b4a8-f16c3c87ce8a', '12', 300.00000000, '5LGVCB', 1, 1000.00, 1, '', '', NULL, NULL, NULL);
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;

/*
 Navicat Premium Data Transfer

 Source Server         : 116.62.71.106
 Source Server Type    : MySQL
 Source Server Version : 50646
 Source Host           : 116.62.71.106:3306
 Source Schema         : db_myotc

 Target Server Type    : MySQL
 Target Server Version : 50646
 File Encoding         : 65001

 Date: 11/03/2020 12:39:39
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for tb_sys_user
-- ----------------------------
DROP TABLE IF EXISTS `tb_sys_user`;
CREATE TABLE `tb_sys_user` (
  `id` varchar(40) COLLATE utf8_bin NOT NULL,
  `account` varchar(100) COLLATE utf8_bin DEFAULT NULL,
  `pwd` varchar(255) COLLATE utf8_bin DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

-- ----------------------------
-- Records of tb_sys_user
-- ----------------------------
BEGIN;
INSERT INTO `tb_sys_user` VALUES ('1', 'admin', 'ecc5172cb9149a295e895ff749792e203cc3ba3c63f5546eb776171c8ba9bcb6');
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;




/*
 Navicat Premium Data Transfer

 Source Server         : 116.62.71.106
 Source Server Type    : MySQL
 Source Server Version : 50646
 Source Host           : 116.62.71.106:3306
 Source Schema         : db_myotc

 Target Server Type    : MySQL
 Target Server Version : 50646
 File Encoding         : 65001

 Date: 11/03/2020 12:39:48
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for tb_user
-- ----------------------------
DROP TABLE IF EXISTS `tb_user`;
CREATE TABLE `tb_user` (
  `id` varchar(40) COLLATE utf8_bin NOT NULL,
  `account` varchar(120) COLLATE utf8_bin DEFAULT NULL,
  `pwd` varchar(120) COLLATE utf8_bin DEFAULT NULL,
  `usdt_balance` double(18,8) DEFAULT NULL,
  `create_time` bigint(20) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_bin;

-- ----------------------------
-- Records of tb_user
-- ----------------------------
BEGIN;
INSERT INTO `tb_user` VALUES ('0c468b05-c9e0-411a-b07c-0789681e46d5', '120779726@qq.com', 'ecc5172cb9149a295e895ff749792e203cc3ba3c63f5546eb776171c8ba9bcb6', 300.00000000, NULL);
INSERT INTO `tb_user` VALUES ('316a61c9-56a5-41b6-a1b6-49411d3b028f', '120779728@qq.com', 'ecc5172cb9149a295e895ff749792e203cc3ba3c63f5546eb776171c8ba9bcb6', 0.00000000, NULL);
INSERT INTO `tb_user` VALUES ('580a8af1-661b-48d0-959a-cbb9caccb9fd', 'admin@qq.com', 'ecc5172cb9149a295e895ff749792e203cc3ba3c63f5546eb776171c8ba9bcb6', 0.00000000, 1583847999854);
INSERT INTO `tb_user` VALUES ('5cb84c56-3ef9-4e03-b7b1-bcb0710e9d45', 'admin1@qq.com', 'e11eb64fb56ea35d7db221946da7d0c6bb5f398cbe6304cbf9d258e62589bea0', 0.00000000, 1583848251697);
INSERT INTO `tb_user` VALUES ('6f19fb35-23a8-47ee-8c0c-2b2384565e3c', 'value1', '2abe3863125de0d2f58fd3b41538eb87c8e6bd22adc87c6be0739434f08b41ad', 0.00000000, 1583848120439);
INSERT INTO `tb_user` VALUES ('87603bac-99ad-46a0-89c0-be8cdf42225f', 'admin2@qq.com', 'e11eb64fb56ea35d7db221946da7d0c6bb5f398cbe6304cbf9d258e62589bea0', 0.00000000, 1583848278366);
INSERT INTO `tb_user` VALUES ('d453d6fb-5ccf-4f69-90bf-b27f457d443e', '120779727@qq.com', 'ecc5172cb9149a295e895ff749792e203cc3ba3c63f5546eb776171c8ba9bcb6', 0.00000000, NULL);
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;




