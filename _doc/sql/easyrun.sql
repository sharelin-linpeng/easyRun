

-- ----------------------------
-- Table structure for application
-- ----------------------------
DROP TABLE IF EXISTS `application`;
CREATE TABLE `application` (
  `id` varchar(100) NOT NULL,
  `app_name` varchar(100) DEFAULT NULL COMMENT '应用名称',
  `app_file` varchar(100) DEFAULT NULL COMMENT '应用运行文件',
  `app_workpath` varchar(100) DEFAULT NULL COMMENT '应用运行路径',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for codebranch
-- ----------------------------
DROP TABLE IF EXISTS `codebranch`;
CREATE TABLE `codebranch` (
  `id` varchar(30) NOT NULL,
  `branch_name` varchar(30) DEFAULT NULL,
  `git_url` varchar(30) DEFAULT NULL,
  `branch` varchar(30) DEFAULT NULL,
  `dir` varchar(30) DEFAULT NULL,
  `commond` varchar(30) DEFAULT NULL,
  `repo_local` varchar(30) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for machine
-- ----------------------------
DROP TABLE IF EXISTS `machine`;
CREATE TABLE `machine` (
  `id` varchar(30) NOT NULL,
  `ip` varchar(30) DEFAULT NULL,
  `machine_name` varchar(30) DEFAULT NULL,
  `login_name` varchar(30) DEFAULT NULL,
  `password` varchar(30) DEFAULT NULL,
  `env` varchar(30) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for publishInfo
-- ----------------------------
DROP TABLE IF EXISTS `publishInfo`;
CREATE TABLE `publishInfo` (
  `id` varchar(30) NOT NULL,
  `application_id` varchar(30) DEFAULT NULL,
  `machine_id` varchar(30) DEFAULT NULL,
  `branch_id` varchar(30) DEFAULT NULL,
  `status` varchar(30) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

