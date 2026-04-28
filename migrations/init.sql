-- 创建用户
-- CREATE USER ssz WITH SUPERUSER PASSWORD 'postgres';
-- 直接在当前数据库中创建表结构
-- 创建用户表
CREATE TABLE "user" (
  "id" bigserial PRIMARY KEY,
  "username" varchar(50) UNIQUE,
  "nickname" varchar(50),
  "avatar" varchar(255),
  "password" varchar(100),
  "phone" varchar(20) UNIQUE,
  "openid" varchar(100) UNIQUE,
  "is_vip" boolean DEFAULT false,
  "vip_expire_time" timestamp,
  "integral" int DEFAULT 0,
  "is_deleted" boolean DEFAULT false,
  "create_time" timestamp DEFAULT CURRENT_TIMESTAMP,
  "update_time" timestamp DEFAULT CURRENT_TIMESTAMP
);
COMMENT ON TABLE "user" IS '用户表';
COMMENT ON COLUMN "user".id IS '用户ID';
COMMENT ON COLUMN "user".username IS '登录用户名';
COMMENT ON COLUMN "user".nickname IS '用户昵称';
COMMENT ON COLUMN "user".avatar IS '头像URL';
COMMENT ON COLUMN "user".password IS '加密密码';
COMMENT ON COLUMN "user".phone IS '手机号';
COMMENT ON COLUMN "user".openid IS '微信OpenID';
COMMENT ON COLUMN "user".is_vip IS '是否VIP会员';
COMMENT ON COLUMN "user".vip_expire_time IS 'VIP过期时间';
COMMENT ON COLUMN "user".integral IS '用户积分';
COMMENT ON COLUMN "user".is_deleted IS '是否删除';
-- 创建标签表
CREATE TABLE "tag" (
  "id" bigserial PRIMARY KEY,
  "name" varchar(50) UNIQUE,
  "category" varchar(50),
  "is_deleted" boolean DEFAULT false,
  "create_time" timestamp DEFAULT CURRENT_TIMESTAMP,
  "update_time" timestamp DEFAULT CURRENT_TIMESTAMP
);
COMMENT ON TABLE "tag" IS '标签表';
COMMENT ON COLUMN "tag".id IS '标签ID';
COMMENT ON COLUMN "tag".name IS '标签名称';
COMMENT ON COLUMN "tag".category IS '标签分类';
-- 创建题目表
CREATE TABLE "question" (
  "id" bigserial PRIMARY KEY,
  "title" text,
  "content" text,
  "difficulty" smallint,
  "answer" text DEFAULT '',
  "explanation" text DEFAULT '',
  "bank_id" bigint DEFAULT 0,
  "heat" int DEFAULT 0,
  "is_vip" boolean DEFAULT false,
  "view_count" int DEFAULT 0,
  "star_count" int DEFAULT 0,
  "like_count" int DEFAULT 0,
  "answered_count" int DEFAULT 0,
  "correct_rate" int DEFAULT 0,
  "creator_id" bigint REFERENCES "user"(id),
  "is_deleted" boolean DEFAULT false,
  "create_time" timestamp DEFAULT CURRENT_TIMESTAMP,
  "update_time" timestamp DEFAULT CURRENT_TIMESTAMP
);
COMMENT ON TABLE "question" IS '题目表';
COMMENT ON COLUMN "question".id IS '题目ID';
COMMENT ON COLUMN "question".title IS '题干';
COMMENT ON COLUMN "question".content IS '题目解析';
COMMENT ON COLUMN "question".difficulty IS '难度：1-简单 2-中等 3-困难';
COMMENT ON COLUMN "question".is_vip IS '是否VIP专属题目';
COMMENT ON COLUMN "question".view_count IS '浏览次数';
COMMENT ON COLUMN "question".star_count IS '收藏次数';
COMMENT ON COLUMN "question".like_count IS '点赞次数';
COMMENT ON COLUMN "question".creator_id IS '创建者ID';
-- 题目全文搜索索引
CREATE INDEX idx_question_title_fts ON question USING gin(to_tsvector('english', title));
CREATE INDEX idx_question_content_fts ON question USING gin(to_tsvector('english', content));
-- 创建题目-标签关联表
CREATE TABLE "question_tag" (
  "id" bigserial PRIMARY KEY,
  "question_id" bigint REFERENCES "question"(id),
  "tag_id" bigint REFERENCES "tag"(id),
  "create_time" timestamp DEFAULT CURRENT_TIMESTAMP,
  UNIQUE(question_id, tag_id)
);
COMMENT ON TABLE "question_tag" IS '题目-标签关联表';
-- 创建题库表
CREATE TABLE "question_bank" (
  "id" bigserial PRIMARY KEY,
  "name" varchar(100),
  "description" text,
  "icon" varchar(255),
  "category" varchar(50),
  "is_vip" boolean DEFAULT false,
  "question_count" int DEFAULT 0,
  "creator_id" bigint REFERENCES "user"(id),
  "sort" int DEFAULT 0,
  "is_deleted" boolean DEFAULT false,
  "create_time" timestamp DEFAULT CURRENT_TIMESTAMP,
  "update_time" timestamp DEFAULT CURRENT_TIMESTAMP
);
COMMENT ON TABLE "question_bank" IS '题库表';
COMMENT ON COLUMN "question_bank".id IS '题库ID';
COMMENT ON COLUMN "question_bank".name IS '题库名称';
COMMENT ON COLUMN "question_bank".description IS '题库描述';
COMMENT ON COLUMN "question_bank".icon IS '题库图标';
COMMENT ON COLUMN "question_bank".category IS '题库分类';
COMMENT ON COLUMN "question_bank".is_vip IS '是否VIP专属题库';
COMMENT ON COLUMN "question_bank".question_count IS '题目数量';
COMMENT ON COLUMN "question_bank".sort IS '排序权重';
-- 创建题库-题目关联表
CREATE TABLE "question_bank_question" (
  "id" bigserial PRIMARY KEY,
  "bank_id" bigint REFERENCES "question_bank"(id),
  "question_id" bigint REFERENCES "question"(id),
  "create_time" timestamp DEFAULT CURRENT_TIMESTAMP,
  UNIQUE(bank_id, question_id)
);
COMMENT ON TABLE "question_bank_question" IS '题库-题目关联表';
-- 创建学习路线表
CREATE TABLE "study_route" (
  "id" bigserial PRIMARY KEY,
  "name" varchar(100),
  "description" text,
  "category" varchar(50),
  "creator_id" bigint REFERENCES "user"(id),
  "sort" int DEFAULT 0,
  "is_deleted" boolean DEFAULT false,
  "create_time" timestamp DEFAULT CURRENT_TIMESTAMP,
  "update_time" timestamp DEFAULT CURRENT_TIMESTAMP
);
COMMENT ON TABLE "study_route" IS '学习路线表';
-- 创建学习路线步骤表
CREATE TABLE "study_route_step" (
  "id" bigserial PRIMARY KEY,
  "route_id" bigint REFERENCES "study_route"(id),
  "step" smallint,
  "name" varchar(100),
  "description" text,
  "bank_id" bigint REFERENCES "question_bank"(id),
  "sort" int DEFAULT 0,
  "create_time" timestamp DEFAULT CURRENT_TIMESTAMP,
  "update_time" timestamp DEFAULT CURRENT_TIMESTAMP
);
COMMENT ON TABLE "study_route_step" IS '学习路线步骤表';
-- 创建用户收藏表
CREATE TABLE "user_favorite" (
  "id" bigserial PRIMARY KEY,
  "user_id" bigint REFERENCES "user"(id),
  "question_id" bigint REFERENCES "question"(id),
  "create_time" timestamp DEFAULT CURRENT_TIMESTAMP,
  UNIQUE(user_id, question_id)
);
COMMENT ON TABLE "user_favorite" IS '用户收藏表';
-- 创建用户刷题记录表
CREATE TABLE "user_question_record" (
  "id" bigserial PRIMARY KEY,
  "user_id" bigint REFERENCES "user"(id),
  "question_id" bigint REFERENCES "question"(id),
  "is_master" boolean DEFAULT false,
  "last_view_time" timestamp,
  "create_time" timestamp DEFAULT CURRENT_TIMESTAMP,
  "update_time" timestamp DEFAULT CURRENT_TIMESTAMP,
  UNIQUE(user_id, question_id)
);
COMMENT ON TABLE "user_question_record" IS '用户刷题记录表';
-- 创建用户搜索历史表
CREATE TABLE "user_search_history" (
  "id" bigserial PRIMARY KEY,
  "user_id" bigint REFERENCES "user"(id),
  "keyword" varchar(100),
  "create_time" timestamp DEFAULT CURRENT_TIMESTAMP
);
COMMENT ON TABLE "user_search_history" IS '用户搜索历史表';
CREATE INDEX idx_user_search_history_user_time ON user_search_history(user_id, create_time DESC);
-- 创建试卷表
CREATE TABLE "exam_paper" (
  "id" bigserial PRIMARY KEY,
  "name" varchar(100),
  "user_id" bigint REFERENCES "user"(id),
  "is_public" boolean DEFAULT false,
  "question_count" int DEFAULT 0,
  "is_deleted" boolean DEFAULT false,
  "create_time" timestamp DEFAULT CURRENT_TIMESTAMP,
  "update_time" timestamp DEFAULT CURRENT_TIMESTAMP
);
COMMENT ON TABLE "exam_paper" IS '试卷表';
-- 创建试卷-题目关联表
CREATE TABLE "exam_paper_question" (
  "id" bigserial PRIMARY KEY,
  "paper_id" bigint REFERENCES "exam_paper"(id),
  "question_id" bigint REFERENCES "question"(id),
  "sort" int DEFAULT 0,
  "create_time" timestamp DEFAULT CURRENT_TIMESTAMP
);
COMMENT ON TABLE "exam_paper_question" IS '试卷-题目关联表';
-- 创建题目回答表
CREATE TABLE "question_answer" (
  "id" bigserial PRIMARY KEY,
  "question_id" bigint REFERENCES "question"(id),
  "user_id" bigint REFERENCES "user"(id),
  "content" text,
  "like_count" int DEFAULT 0,
  "is_deleted" boolean DEFAULT false,
  "create_time" timestamp DEFAULT CURRENT_TIMESTAMP,
  "update_time" timestamp DEFAULT CURRENT_TIMESTAMP
);
COMMENT ON TABLE "question_answer" IS '题目回答表';
-- 创建回答点赞表
CREATE TABLE "user_answer_like" (
  "id" bigserial PRIMARY KEY,
  "user_id" bigint REFERENCES "user"(id),
  "answer_id" bigint REFERENCES "question_answer"(id),
  "create_time" timestamp DEFAULT CURRENT_TIMESTAMP,
  UNIQUE(user_id, answer_id)
);
COMMENT ON TABLE "user_answer_like" IS '回答点赞表';
-- 创建简历模板表
CREATE TABLE "resume_template" (
  "id" bigserial PRIMARY KEY,
  "name" varchar(100),
  "description" text,
  "download_url" varchar(255),
  "download_count" int DEFAULT 0,
  "is_vip" boolean DEFAULT false,
  "create_time" timestamp DEFAULT CURRENT_TIMESTAMP,
  "update_time" timestamp DEFAULT CURRENT_TIMESTAMP
);
COMMENT ON TABLE "resume_template" IS '简历模板表';
-- 创建面试视频表
CREATE TABLE "interview_video" (
  "id" bigserial PRIMARY KEY,
  "title" varchar(100),
  "description" text,
  "cover_url" varchar(255),
  "video_url" varchar(255),
  "play_count" int DEFAULT 0,
  "is_vip" boolean DEFAULT false,
  "create_time" timestamp DEFAULT CURRENT_TIMESTAMP,
  "update_time" timestamp DEFAULT CURRENT_TIMESTAMP
);
COMMENT ON TABLE "interview_video" IS '面试视频表';
