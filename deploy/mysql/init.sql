CREATE DATABASE IF NOT EXISTS qizhan DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
USE qizhan;

CREATE TABLE IF NOT EXISTS about_contents (
  id BIGINT PRIMARY KEY AUTO_INCREMENT,
  hero_title VARCHAR(255) NOT NULL,
  hero_subtitle VARCHAR(500) NOT NULL,
  company_overview TEXT NOT NULL,
  businesses_json JSON NOT NULL,
  advantages_json JSON NOT NULL,
  milestones_json JSON NOT NULL,
  phone VARCHAR(100) NOT NULL,
  email VARCHAR(255) NOT NULL,
  address VARCHAR(500) NOT NULL
);

INSERT INTO about_contents (
  hero_title,
  hero_subtitle,
  company_overview,
  businesses_json,
  advantages_json,
  milestones_json,
  phone,
  email,
  address
)
VALUES (
  '启展跨境咨询服务',
  '为企业与家庭提供专业、合规、可落地的北美服务方案',
  '我们是一家专注于跨境咨询服务的团队，服务覆盖商业落地、移民规划、留学申请、税务协同与本地生活支持，强调长期陪跑与结果导向。',
  JSON_ARRAY('企业出海咨询', '加拿大移民服务', '留学与签证方案', '税务与财务协同', '本地安家支持'),
  JSON_ARRAY('一对一顾问机制', '透明流程与费用', '本地持牌合作网络', '多语种服务团队', '项目进度实时同步'),
  JSON_ARRAY(
    JSON_OBJECT('year', '2018', 'content', '团队成立，聚焦加拿大市场。'),
    JSON_OBJECT('year', '2020', 'content', '服务升级，建立商业与家庭双线服务体系。'),
    JSON_OBJECT('year', '2022', 'content', '上线数字化客户管理流程。'),
    JSON_OBJECT('year', '2024', 'content', '拓展跨境企业综合服务能力。')
  ),
  '+1-604-123-4567',
  'service@qizhan.ca',
  'Vancouver, BC, Canada'
)
ON DUPLICATE KEY UPDATE hero_title = VALUES(hero_title);
