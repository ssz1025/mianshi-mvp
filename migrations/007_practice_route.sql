-- 创建刷题路线表
CREATE TABLE IF NOT EXISTS "practice_route" (
  "id" bigserial PRIMARY KEY,
  "name" varchar(100),
  "title" varchar(100),
  "icon" varchar(255),
  "color" varchar(100),
  "description" text,
  "target_level" varchar(100),
  "suitable_for" varchar(255)[],
  "skills" varchar(255)[],
  "interview_weight" varchar(50),
  "sort" int DEFAULT 0,
  "is_deleted" boolean DEFAULT false,
  "create_time" timestamp DEFAULT CURRENT_TIMESTAMP,
  "update_time" timestamp DEFAULT CURRENT_TIMESTAMP
);
COMMENT ON TABLE "practice_route" IS '刷题路线表';

-- 创建刷题路线阶段表
CREATE TABLE IF NOT EXISTS "route_phase" (
  "id" bigserial PRIMARY KEY,
  "route_id" bigint REFERENCES "practice_route"(id),
  "phase" varchar(100),
  "duration" varchar(50),
  "topics" varchar(255)[],
  "description" text,
  "sort" int DEFAULT 0,
  "create_time" timestamp DEFAULT CURRENT_TIMESTAMP,
  "update_time" timestamp DEFAULT CURRENT_TIMESTAMP
);
COMMENT ON TABLE "route_phase" IS '刷题路线阶段表';

-- 插入前端工程师路线
INSERT INTO "practice_route" ("name", "title", "icon", "color", "description", "target_level", "suitable_for", "skills", "interview_weight", "sort") VALUES
('frontend', '前端工程师', '🎨', 'from-orange-400 to-pink-500', '专注于用户界面开发和用户体验优化，掌握前端主流框架和工程化工具。', '中高级前端工程师 (15-30K)', ARRAY['零基础入门', '转行前端', '提升进阶'], ARRAY['React', 'Vue', 'TypeScript', 'Webpack', 'Node.js', '性能优化'], '⭐⭐⭐⭐⭐', 1);

INSERT INTO "route_phase" ("route_id", "phase", "duration", "topics", "description", "sort") VALUES
(1, '第一阶段：基础入门', '2-3周', ARRAY['HTML基础', 'CSS布局', 'JavaScript语法', 'DOM操作', 'BOM操作'], '掌握网页结构、样式和基础交互逻辑', 1),
(1, '第二阶段：进阶提升', '3-4周', ARRAY['ES6+特性', '异步编程', '原型链', '闭包', '模块化', 'TypeScript基础'], '深入理解JavaScript核心概念和现代语法', 2),
(1, '第三阶段：框架学习', '4-6周', ARRAY['React/Vue核心', '组件化开发', '状态管理', '虚拟DOM', 'Diff算法', 'Hooks'], '掌握主流前端框架的开发模式', 3),
(1, '第四阶段：工程化与实战', '3-4周', ARRAY['Webpack/Vite', '项目构建', '性能优化', '单元测试', 'CI/CD'], '提升工程化能力和项目实战经验', 4);

-- 插入后端工程师路线
INSERT INTO "practice_route" ("name", "title", "icon", "color", "description", "target_level", "suitable_for", "skills", "interview_weight", "sort") VALUES
('backend', '后端工程师', '⚙️', 'from-blue-400 to-cyan-500', '负责服务器端逻辑、数据库设计和API开发，构建稳定的业务系统。', '中高级后端工程师 (20-40K)', ARRAY['计算机专业', '转行后端', '技术深耕'], ARRAY['Java', 'Go', 'Spring Boot', 'MySQL', 'Redis', '微服务', 'Docker'], '⭐⭐⭐⭐⭐', 2);

INSERT INTO "route_phase" ("route_id", "phase", "duration", "topics", "description", "sort") VALUES
(2, '第一阶段：语言基础', '3-4周', ARRAY['Java/Go/Python基础', '面向对象', '集合框架', 'I/O操作', '异常处理'], '掌握一门后端编程语言的核心语法', 1),
(2, '第二阶段：数据库', '3-4周', ARRAY['MySQL索引', '事务隔离', 'Redis缓存', 'SQL优化', '分库分表'], '深入理解数据存储和缓存技术', 2),
(2, '第三阶段：框架与微服务', '4-6周', ARRAY['Spring Boot', '微服务架构', '消息队列', 'RPC通信', '服务治理'], '掌握企业级开发框架和分布式架构', 3),
(2, '第四阶段：架构设计', '3-4周', ARRAY['系统设计', '高并发', 'CAP定理', '分布式事务', '限流熔断'], '具备系统架构设计和性能优化能力', 4);

-- 插入算法工程师路线
INSERT INTO "practice_route" ("name", "title", "icon", "color", "description", "target_level", "suitable_for", "skills", "interview_weight", "sort") VALUES
('algorithm', '算法工程师', '🧮', 'from-purple-400 to-indigo-500', '专注于算法设计和优化，解决复杂的工程问题，是晋升大厂的核心竞争力。', '算法工程师 / 架构师 (25-50K)', ARRAY['冲刺大厂', '算法竞赛', '提升编程能力'], ARRAY['数据结构', '动态规划', '二叉树', '图论', '贪心', '回溯', '位运算'], '⭐⭐⭐⭐⭐', 3);

INSERT INTO "route_phase" ("route_id", "phase", "duration", "topics", "description", "sort") VALUES
(3, '第一阶段：基础算法', '2-3周', ARRAY['时间空间复杂度', '数组链表', '栈队列', '哈希表', '字符串'], '掌握算法基础概念和基本数据结构', 1),
(3, '第二阶段：核心算法', '4-5周', ARRAY['二分查找', '双指针', '滑动窗口', '回溯算法', '贪心算法', '动态规划基础'], '学习高频面试算法题型', 2),
(3, '第三阶段：高级专题', '4-5周', ARRAY['动态规划进阶', '二叉树', '图论', '堆与优先队列', '位运算'], '攻克算法难点和高级题型', 3),
(3, '第四阶段：综合提升', '3-4周', ARRAY['综合刷题', '面经实战', '模拟面试', '算法思维'], '系统提升解题能力和面试表现', 4);

-- 插入全栈工程师路线
INSERT INTO "practice_route" ("name", "title", "icon", "color", "description", "target_level", "suitable_for", "skills", "interview_weight", "sort") VALUES
('fullstack', '全栈工程师', '🌐', 'from-green-400 to-emerald-500', '前后端全能型开发者，具备独立完成完整项目的能力。', '高级全栈工程师 (25-45K)', ARRAY['独立开发者', '创业需求', '快速成长'], ARRAY['React', 'Node.js', 'Go', 'MySQL', 'Docker', 'AWS', '系统设计'], '⭐⭐⭐⭐', 4);

INSERT INTO "route_phase" ("route_id", "phase", "duration", "topics", "description", "sort") VALUES
(4, '第一阶段：前端基础', '3-4周', ARRAY['HTML/CSS/JS', 'React/Vue', '前端工程化', '状态管理'], '掌握前端开发核心技能', 1),
(4, '第二阶段：后端基础', '3-4周', ARRAY['Node.js/Go', 'RESTful API', '数据库设计', '认证授权'], '掌握后端开发核心技能', 2),
(4, '第三阶段：全栈实战', '4-6周', ARRAY['项目开发', '部署上线', '性能优化', 'DevOps', 'SaaS开发'], '完成完整的全栈项目实战', 3),
(4, '第四阶段：架构提升', '2-3周', ARRAY['系统设计', '微服务', '云原生', '团队协作'], '提升架构设计和团队协作能力', 4);

-- 插入DevOps工程师路线
INSERT INTO "practice_route" ("name", "title", "icon", "color", "description", "target_level", "suitable_for", "skills", "interview_weight", "sort") VALUES
('devops', 'DevOps工程师', '🚀', 'from-cyan-400 to-blue-500', '负责构建自动化流水线、容器化部署和系统运维，提升开发效率。', '高级DevOps工程师 (25-45K)', ARRAY['运维转型', '效率提升', '云原生方向'], ARRAY['Kubernetes', 'Docker', 'Jenkins', 'Prometheus', 'Linux', 'Ansible', 'Terraform'], '⭐⭐⭐⭐', 5);

INSERT INTO "route_phase" ("route_id", "phase", "duration", "topics", "description", "sort") VALUES
(5, '第一阶段：基础入门', '2-3周', ARRAY['Linux系统', 'Shell脚本', '网络基础', 'Docker基础'], '掌握Linux和容器基础', 1),
(5, '第二阶段：容器编排', '3-4周', ARRAY['Docker进阶', 'Kubernetes', 'Helm', 'Service Mesh'], '掌握容器编排和管理', 2),
(5, '第三阶段：CI/CD', '3-4周', ARRAY['Jenkins', 'GitLab CI', 'ArgoCD', '自动化测试', '代码质量'], '构建完整的持续集成/部署流程', 3),
(5, '第四阶段：监控与运维', '2-3周', ARRAY['Prometheus', 'Grafana', '日志系统', '故障排查', 'SRE实践'], '建立完整的监控和运维体系', 4);

-- 插入移动端工程师路线
INSERT INTO "practice_route" ("name", "title", "icon", "color", "description", "target_level", "suitable_for", "skills", "interview_weight", "sort") VALUES
('mobile', '移动端工程师', '📱', 'from-pink-400 to-rose-500', '专注于iOS或Android移动应用开发，打造优质的移动端用户体验。', '高级移动端工程师 (20-40K)', ARRAY['移动开发', 'App方向', 'iOS/Android'], ARRAY['Swift', 'Kotlin', 'Flutter', 'iOS', 'Android', '性能优化', '混合开发'], '⭐⭐⭐⭐', 6);

INSERT INTO "route_phase" ("route_id", "phase", "duration", "topics", "description", "sort") VALUES
(6, '第一阶段：基础入门', '3-4周', ARRAY['Swift/Kotlin', 'UI基础', '数据存储', '网络请求'], '掌握移动端开发基础', 1),
(6, '第二阶段：进阶开发', '3-4周', ARRAY['多线程', '性能优化', '动画特效', 'Framework深入'], '深入理解移动端核心机制', 2),
(6, '第三阶段：架构设计', '3-4周', ARRAY['MVVM/MVP', '组件化', '插件化', '热更新', '混合开发'], '掌握移动端架构设计模式', 3),
(6, '第四阶段：实战与优化', '2-3周', ARRAY['项目实战', '上架发布', '性能监控', '包体积优化'], '完成项目并优化上线', 4);

-- 插入数据工程师路线
INSERT INTO "practice_route" ("name", "title", "icon", "color", "description", "target_level", "suitable_for", "skills", "interview_weight", "sort") VALUES
('data', '数据工程师', '📊', 'from-teal-400 to-cyan-500', '负责数据采集、清洗、存储和分析，构建数据中台和数据流水线。', '高级数据工程师 (25-50K)', ARRAY['数据方向', 'BI工程师', '数仓建设'], ARRAY['Spark', 'Flink', 'Hive', 'Kafka', 'Python', 'SQL', '数据建模'], '⭐⭐⭐⭐', 7);

INSERT INTO "route_phase" ("route_id", "phase", "duration", "topics", "description", "sort") VALUES
(7, '第一阶段：基础技能', '2-3周', ARRAY['SQL', 'Python', 'Excel', 'Linux', '数据仓库概念'], '掌握数据处理基础工具', 1),
(7, '第二阶段：数据开发', '4-5周', ARRAY['Hive', 'Spark', 'Flink', 'ETL开发', '数据建模'], '掌握大数据处理技术', 2),
(7, '第三阶段：数据平台', '3-4周', ARRAY['离线计算', '实时计算', '数据湖', '数据治理'], '构建数据平台和流水线', 3),
(7, '第四阶段：数据应用', '2-3周', ARRAY['BI可视化', '指标体系', 'A/B测试', '数据挖掘基础'], '数据驱动业务决策', 4);

-- 插入测试工程师路线
INSERT INTO "practice_route" ("name", "title", "icon", "color", "description", "target_level", "suitable_for", "skills", "interview_weight", "sort") VALUES
('test', '测试工程师', '🧪', 'from-amber-400 to-orange-500', '负责软件质量保障，包括功能测试、自动化测试和性能测试。', '高级测试工程师 (20-35K)', ARRAY['功能测试', '自动化测试', '测开方向'], ARRAY['Selenium', 'JMeter', 'Postman', 'Python', 'Appium', '性能测试', 'CI/CD'], '⭐⭐⭐⭐', 8);

INSERT INTO "route_phase" ("route_id", "phase", "duration", "topics", "description", "sort") VALUES
(8, '第一阶段：测试基础', '2-3周', ARRAY['测试理论', '测试用例', '缺陷管理', '测试报告'], '掌握软件测试基本概念', 1),
(8, '第二阶段：测试技术', '3-4周', ARRAY['功能测试', '接口测试', '数据库测试', 'Linux', 'Shell'], '掌握测试核心技能', 2),
(8, '第三阶段：自动化', '4-5周', ARRAY['Selenium/Appium', 'Postman', 'JMeter', 'CI集成', '测试框架'], '构建自动化测试体系', 3),
(8, '第四阶段：测试开发', '2-3周', ARRAY['测试平台', '性能测试', '安全测试', '质量度量'], '提升测试开发能力', 4);
