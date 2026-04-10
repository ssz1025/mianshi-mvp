# 面试鸭PostgreSQL数据库设计

本文档为模仿面试鸭（[https://www.mianshiya.com/](https://www.mianshiya.com/)）网站的 PostgreSQL 数据库设计，可用于支撑网站全功能接口服务的开发，覆盖用户、题库、题目、刷题、收藏、路线、组卷等核心业务场景。

## 整体架构说明

本设计采用关系型数据库模型，针对面试刷题场景做了针对性优化，支持：

- 多维度的题目、题库管理

- 用户刷题、收藏、记录等个性化功能

- 学习路线、组卷、搜索等扩展功能

- VIP 会员权限控制

- 数据软删除，保证数据可追溯

---

## 核心表结构设计

### 1. 用户表 `user`

存储用户基础信息，支持账号密码、微信登录两种方式，同时包含会员、积分等信息。

|字段名|类型|说明|约束|
|---|---|---|---|
|id|bigserial|用户 ID|主键|
|username|varchar(50)|登录用户名|唯一|
|nickname|varchar(50)|用户昵称|-|
|avatar|varchar(255)|头像 URL|-|
|password|varchar(100)|加密密码（BCrypt）|-|
|phone|varchar(20)|手机号|唯一|
|openid|varchar(100)|微信 OpenID（微信登录用）|唯一|
|is_vip|boolean|是否为 VIP 会员|默认 false|
|vip_expire_time|timestamp|VIP 过期时间|-|
|integral|int|用户积分|默认 0|
|is_deleted|boolean|是否删除|默认 false|
|create_time|timestamp|创建时间|默认 now ()|
|update_time|timestamp|更新时间|默认 now ()|
### 2. 标签表 `tag`

存储题目标签，用于题目的多维度分类筛选，支持技术、公司、岗位等多类标签。

|字段名|类型|说明|约束|
|---|---|---|---|
|id|bigserial|标签 ID|主键|
|name|varchar(50)|标签名称|唯一|
|category|varchar(50)|标签分类（如：技术、公司、岗位）|-|
|is_deleted|boolean|是否删除|默认 false|
|create_time|timestamp|创建时间|默认 now ()|
|update_time|timestamp|更新时间|默认 now ()|
### 3. 题目表 `question`

存储面试题的核心信息，包含题干、解析、统计数据等。

|字段名|类型|说明|约束|
|---|---|---|---|
|id|bigserial|题目 ID|主键|
|title|text|题干 / 问题|-|
|content|text|题目解析 / 答案|-|
|difficulty|smallint|难度：1 - 简单 2 - 中等 3 - 困难|-|
|is_vip|boolean|是否 VIP 专属题目|默认 false|
|view_count|int|浏览次数|默认 0|
|star_count|int|收藏次数|默认 0|
|like_count|int|点赞次数|默认 0|
|creator_id|bigint|创建者 ID（关联 [user.id](user.id)）|外键|
|is_deleted|boolean|是否删除|默认 false|
|create_time|timestamp|创建时间|默认 now ()|
|update_time|timestamp|更新时间|默认 now ()|
### 4. 题目 - 标签关联表 `question_tag`

题目与标签为多对多关系，通过中间表实现关联。

|字段名|类型|说明|约束|
|---|---|---|---|
|id|bigserial|关联 ID|主键|
|question_id|bigint|题目 ID（关联 [question.id](question.id)）|外键|
|tag_id|bigint|标签 ID（关联 [tag.id](tag.id)）|外键|
|create_time|timestamp|创建时间|默认 now ()|
|-|-|联合唯一|(question_id, tag_id)|
### 5. 题库表 `question_bank`

存储题库信息，题库是题目的集合，比如 Java 热门题库、美团面经题库等。

|字段名|类型|说明|约束|
|---|---|---|---|
|id|bigserial|题库 ID|主键|
|name|varchar(100)|题库名称|-|
|description|text|题库描述|-|
|icon|varchar(255)|题库图标 URL|-|
|category|varchar(50)|题库分类（如：热门、Java、真实面经）|-|
|is_vip|boolean|是否 VIP 专属题库|默认 false|
|question_count|int|题库内题目数量|默认 0|
|creator_id|bigint|创建者 ID（关联 [user.id](user.id)）|外键|
|sort|int|展示排序权重|默认 0|
|is_deleted|boolean|是否删除|默认 false|
|create_time|timestamp|创建时间|默认 now ()|
|update_time|timestamp|更新时间|默认 now ()|
### 6. 题库 - 题目关联表 `question_bank_question`

题库与题目为多对多关系，一个题目可属于多个题库，一个题库包含多个题目。

|字段名|类型|说明|约束|
|---|---|---|---|
|id|bigserial|关联 ID|主键|
|bank_id|bigint|题库 ID（关联 question\[_bank.id](_bank.id)）|外键|
|question_id|bigint|题目 ID（关联 [question.id](question.id)）|外键|
|create_time|timestamp|创建时间|默认 now ()|
|-|-|联合唯一|(bank_id, question_id)|
### 7. 学习路线表 `study_route`

存储刷题学习路线，比如 Java1-3 年刷题路线、春招刷题路线等。

|字段名|类型|说明|约束|
|---|---|---|---|
|id|bigserial|路线 ID|主键|
|name|varchar(100)|路线名称|-|
|description|text|路线描述|-|
|category|varchar(50)|路线分类（如：Java、前端）|-|
|creator_id|bigint|创建者 ID（关联 [user.id](user.id)）|外键|
|sort|int|展示排序权重|默认 0|
|is_deleted|boolean|是否删除|默认 false|
|create_time|timestamp|创建时间|默认 now ()|
|update_time|timestamp|更新时间|默认 now ()|
### 8. 学习路线步骤表 `study_route_step`

学习路线的步骤，每个步骤对应一个题库，引导用户按顺序刷题。

|字段名|类型|说明|约束|
|---|---|---|---|
|id|bigserial|步骤 ID|主键|
|route_id|bigint|路线 ID（关联 study\[_route.id](_route.id)）|外键|
|step|smallint|步骤序号|-|
|name|varchar(100)|步骤名称|-|
|description|text|步骤描述|-|
|bank_id|bigint|步骤对应题库 ID（关联 question\[_bank.id](_bank.id)）|外键|
|sort|int|步骤内排序|默认 0|
|create_time|timestamp|创建时间|默认 now ()|
|update_time|timestamp|更新时间|默认 now ()|
### 9. 用户收藏表 `user_favorite`

用户收藏的题目记录。

|字段名|类型|说明|约束|
|---|---|---|---|
|id|bigserial|收藏 ID|主键|
|user_id|bigint|用户 ID（关联 [user.id](user.id)）|外键|
|question_id|bigint|题目 ID（关联 [question.id](question.id)）|外键|
|create_time|timestamp|收藏时间|默认 now ()|
|-|-|联合唯一|(user_id, question_id)|
### 10. 用户刷题记录表 `user_question_record`

用户的刷题记录，包含题目掌握状态标记。

|字段名|类型|说明|约束|
|---|---|---|---|
|id|bigserial|记录 ID|主键|
|user_id|bigint|用户 ID（关联 [user.id](user.id)）|外键|
|question_id|bigint|题目 ID（关联 [question.id](question.id)）|外键|
|is_master|boolean|是否标记为已掌握|默认 false|
|last_view_time|timestamp|最后查看时间|-|
|create_time|timestamp|记录创建时间|默认 now ()|
|update_time|timestamp|记录更新时间|默认 now ()|
|-|-|联合唯一|(user_id, question_id)|
### 11. 用户搜索历史表 `user_search_history`

用户的搜索记录，用于热门搜索、搜索推荐。

|字段名|类型|说明|约束|
|---|---|---|---|
|id|bigserial|历史 ID|主键|
|user_id|bigint|用户 ID（关联 [user.id](user.id)）|外键|
|keyword|varchar(100)|搜索关键词|-|
|create_time|timestamp|搜索时间|默认 now ()|
### 12. 试卷表 `exam_paper`

用户自定义组卷的试卷信息，支持一键组卷、试卷分享。

|字段名|类型|说明|约束|
|---|---|---|---|
|id|bigserial|试卷 ID|主键|
|name|varchar(100)|试卷名称|-|
|user_id|bigint|创建用户 ID（关联 [user.id](user.id)）|外键|
|is_public|boolean|是否公开试卷|默认 false|
|question_count|int|试卷内题目数量|默认 0|
|is_deleted|boolean|是否删除|默认 false|
|create_time|timestamp|创建时间|默认 now ()|
|update_time|timestamp|更新时间|默认 now ()|
### 13. 试卷 - 题目关联表 `exam_paper_question`

试卷与题目的多对多关联，记录题目在试卷中的顺序。

|字段名|类型|说明|约束|
|---|---|---|---|
|id|bigserial|关联 ID|主键|
|paper_id|bigint|试卷 ID（关联 exam\[_paper.id](_paper.id)）|外键|
|question_id|bigint|题目 ID（关联 [question.id](question.id)）|外键|
|sort|int|题目在试卷中的排序|默认 0|
|create_time|timestamp|创建时间|默认 now ()|
### 14. 题目回答表 `question_answer`

用户针对题目发布的补充回答 / 讨论，支持社区互动。

|字段名|类型|说明|约束|
|---|---|---|---|
|id|bigserial|回答 ID|主键|
|question_id|bigint|题目 ID（关联 [question.id](question.id)）|外键|
|user_id|bigint|用户 ID（关联 [user.id](user.id)）|外键|
|content|text|回答内容|-|
|like_count|int|点赞数|默认 0|
|is_deleted|boolean|是否删除|默认 false|
|create_time|timestamp|创建时间|默认 now ()|
|update_time|timestamp|更新时间|默认 now ()|
### 15. 回答点赞表 `user_answer_like`

用户对回答的点赞记录。

|字段名|类型|说明|约束|
|---|---|---|---|
|id|bigserial|点赞 ID|主键|
|user_id|bigint|用户 ID（关联 [user.id](user.id)）|外键|
|answer_id|bigint|回答 ID（关联 question\[_answer.id](_answer.id)）|外键|
|create_time|timestamp|点赞时间|默认 now ()|
|-|-|联合唯一|(user_id, answer_id)|
### 16. 简历模板表 `resume_template`

简历模板资源，供用户下载使用。

|字段名|类型|说明|约束|
|---|---|---|---|
|id|bigserial|模板 ID|主键|
|name|varchar(100)|模板名称|-|
|description|text|模板描述|-|
|download_url|varchar(255)|模板下载链接|-|
|download_count|int|下载次数|默认 0|
|is_vip|boolean|是否 VIP 专属|默认 false|
|create_time|timestamp|创建时间|默认 now ()|
|update_time|timestamp|更新时间|默认 now ()|
### 17. 面试视频表 `interview_video`

面试经验视频资源。

|字段名|类型|说明|约束|
|---|---|---|---|
|id|bigserial|视频 ID|主键|
|title|varchar(100)|视频标题|-|
|description|text|视频描述|-|
|cover_url|varchar(255)|视频封面 URL|-|
|video_url|varchar(255)|视频播放链接|-|
|play_count|int|播放次数|默认 0|
|is_vip|boolean|是否 VIP 专属|默认 false|
|create_time|timestamp|创建时间|默认 now ()|
|update_time|timestamp|更新时间|默认 now ()|
---

## 完整建表 SQL 语句

以下是可直接在 PostgreSQL 中执行的建表语句，包含外键、索引、注释：

```sql

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
    "is_vip" boolean DEFAULT false,
    "view_count" int DEFAULT 0,
    "star_count" int DEFAULT 0,
    "like_count" int DEFAULT 0,
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
```

---

## 扩展说明

1. **全文搜索**：本设计中为题目标题和内容创建了 PostgreSQL 原生的全文搜索索引，可直接支持题目搜索功能，无需额外引入 Elasticsearch（如果需要更复杂的搜索可扩展）。

2. **软删除**：所有核心表都增加了`is_deleted`字段，实现软删除，避免数据误删，同时保留数据追溯能力。

3. **权限控制**：通过`is_vip`字段实现 VIP 资源的权限控制，可支撑会员业务。

4. **扩展性**：表结构预留了扩展空间，可根据业务需求增加更多字段，比如题目编辑历史、用户消息等。
> （注：文档部分内容可能由 AI 生成）