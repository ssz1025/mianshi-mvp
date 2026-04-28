-- 热门题目测试数据
-- 先清理旧数据（按依赖顺序删除）
DELETE FROM user_answer_like;
DELETE FROM question_answer;
DELETE FROM exam_paper_question;
DELETE FROM exam_paper;
DELETE FROM user_question_record;
DELETE FROM user_favorite;
DELETE FROM user_search_history;
DELETE FROM question_bank_question;
DELETE FROM question_tag;
DELETE FROM question;
DELETE FROM tag;
DELETE FROM question_bank WHERE id > 0;
DELETE FROM "user" WHERE id > 0;

-- 重置序列
ALTER SEQUENCE user_id_seq RESTART WITH 1;
ALTER SEQUENCE tag_id_seq RESTART WITH 1;
ALTER SEQUENCE question_id_seq RESTART WITH 1;
ALTER SEQUENCE question_tag_id_seq RESTART WITH 1;
ALTER SEQUENCE question_bank_id_seq RESTART WITH 1;
ALTER SEQUENCE question_bank_question_id_seq RESTART WITH 1;
ALTER SEQUENCE question_answer_id_seq RESTART WITH 1;
ALTER SEQUENCE user_question_record_id_seq RESTART WITH 1;
ALTER SEQUENCE user_favorite_id_seq RESTART WITH 1;
ALTER SEQUENCE user_search_history_id_seq RESTART WITH 1;
ALTER SEQUENCE exam_paper_id_seq RESTART WITH 1;
ALTER SEQUENCE exam_paper_question_id_seq RESTART WITH 1;
ALTER SEQUENCE user_answer_like_id_seq RESTART WITH 1;

-- 1. 插入用户
INSERT INTO "user" (id, username, nickname, avatar, password, is_vip, integral) VALUES
(1, 'admin', '管理员', '', '$2a$14$dummyhashforadmin', true, 1000),
(2, 'testuser1', '测试用户1', '', '$2a$14$dummyhashfortest1', false, 100),
(3, 'testuser2', '测试用户2', '', '$2a$14$dummyhashfortest2', false, 50),
(4, 'testuser3', '测试用户3', '', '$2a$14$dummyhashfortest3', true, 500),
(5, 'zhangsan', '张三', '', '$2a$14$dummyhashforzhangsan', false, 200);

-- 2. 插入标签
INSERT INTO "tag" (id, name, category) VALUES
(1, 'Java', 'Java'),
(2, 'Spring Boot', 'Java'),
(3, '集合', 'Java'),
(4, '哈希表', 'Java'),
(5, 'MySQL', '数据库'),
(6, '数据库', '数据库'),
(7, '索引', '数据库'),
(8, 'Redis', '数据库'),
(9, '缓存', '数据库'),
(10, '性能优化', '后端'),
(11, 'TCP', '计算机网络'),
(12, '计算机网络', '计算机网络'),
(13, '协议', '计算机网络'),
(14, '多线程', 'Java'),
(15, '线程池', 'Java'),
(16, 'JVM', 'Java'),
(17, '垃圾回收', 'Java'),
(18, '分布式系统', '后端'),
(19, 'CAP', '后端'),
(20, '理论', '后端'),
(21, 'React', '前端'),
(22, 'Hooks', '前端'),
(23, '前端', '前端'),
(24, '消息队列', '后端'),
(25, '可靠性', '后端'),
(26, '数据结构', '算法'),
(27, '算法', '算法'),
(28, '排序', '算法'),
(29, '时间复杂度', '算法'),
(30, '并发', '数据库'),
(31, '锁', '数据库'),
(32, 'HTTP', '计算机网络'),
(33, 'HTTPS', '计算机网络'),
(34, 'Vue', '前端'),
(35, '前端框架', '前端'),
(36, 'Go', 'Go'),
(37, 'Docker', '运维'),
(38, '操作系统', '操作系统'),
(39, 'Python', 'Python'),
(40, 'C++', 'C++');

-- 3. 插入题库
INSERT INTO question_bank (id, name, description, icon, category, is_vip, question_count, creator_id, sort) VALUES
(1, 'Java 面试题库', '涵盖 Java 基础、集合、多线程、JVM 等核心面试题', '☕', 'Java', false, 6, 1, 100),
(2, '数据库面试题库', 'MySQL、Redis 等数据库相关面试题', '🗄️', '数据库', false, 3, 1, 90),
(3, '计算机网络面试题库', 'TCP/IP、HTTP 等网络协议面试题', '🌐', '计算机网络', false, 2, 1, 80),
(4, '后端面试题库', '分布式系统、消息队列等后端面试题', '⚙️', '后端', false, 2, 1, 70),
(5, '前端面试题库', 'React、Vue 等前端框架面试题', '🖥️', '前端', false, 2, 1, 60),
(6, '算法面试题库', '排序、搜索等算法面试题', '🧮', '算法', false, 1, 1, 50);

-- 4. 插入题目
INSERT INTO question (id, title, content, difficulty, is_vip, view_count, star_count, like_count, creator_id, answer, explanation) VALUES
(1, 'HashMap 的底层实现原理是什么？', 'HashMap 是基于哈希表实现的，JDK 1.8 之后采用数组+链表+红黑树的数据结构。当链表长度超过 8 且数组长度超过 64 时，链表会转换为红黑树以提高查询效率。put 过程：先计算 hash 值定位桶位置，如果桶为空直接插入，否则遍历链表或红黑树进行更新或插入。扩容时默认容量翻倍，重新计算每个元素的位置。', 2, false, 9856, 3200, 2100, 1, 'HashMap 基于哈希表实现，JDK 1.8 之后采用数组+链表+红黑树的数据结构。当链表长度超过 8 且数组长度超过 64 时，链表会转换为红黑树以提高查询效率。put 过程：先计算 hash 值定位桶位置，如果桶为空直接插入，否则遍历链表或红黑树进行更新或插入。扩容时默认容量翻倍，重新计算每个元素的位置。', 'HashMap 是面试高频考点，核心在于理解哈希冲突的解决方式和扩容机制。JDK 1.8 引入红黑树是重要优化，需要掌握链表转红黑树的阈值条件。还需要了解 HashMap 的线程安全问题，为什么推荐使用 ConcurrentHashMap 替代。'),
(2, 'MySQL 索引的底层数据结构为什么用 B+ 树？', 'B+ 树相比 B 树和二叉树，具有更矮的树高，磁盘 I/O 次数更少。叶子节点通过链表相连，支持范围查询。非叶子节点只存储键值，单个节点能存储更多键值，从而降低树的高度。B+ 树的查询性能稳定，每次查询都需要走到叶子节点，路径长度相同。', 2, false, 8734, 2800, 1900, 1, 'B+ 树相比 B 树和二叉树，具有更矮的树高，磁盘 I/O 次数更少。叶子节点通过链表相连，支持范围查询。非叶子节点只存储键值，单个节点能存储更多键值，从而降低树的高度。B+ 树的查询性能稳定，每次查询都需要走到叶子节点，路径长度相同。', 'B+ 树是数据库索引的核心数据结构，面试常考。需要对比 B 树、B+ 树、红黑树、跳表的优劣。重点理解为什么数据库不用红黑树（树太高，磁盘 I/O 多）和为什么不用跳表（范围查询不如 B+ 树高效）。'),
(3, 'Redis 为什么这么快？', 'Redis 基于内存操作，使用单线程避免上下文切换和锁竞争，采用 IO 多路复用模型（epoll），高效的数据结构设计（如 SDS、跳表、压缩列表、整数集合等）。此外 Redis 还有多线程 I/O 读写（6.0+），但命令执行仍是单线程。', 2, false, 8120, 2600, 1800, 1, 'Redis 基于内存操作，使用单线程避免上下文切换和锁竞争，采用 IO 多路复用模型（epoll），高效的数据结构设计（如 SDS、跳表、压缩列表、整数集合等）。此外 Redis 6.0+ 还支持多线程 I/O 读写，但命令执行仍是单线程。', 'Redis 高性能是经典面试题。需要从内存、线程模型、I/O 模型、数据结构四个维度分析。Redis 6.0 的多线程 I/O 是常考点，需要理解多线程只用于网络 I/O，命令执行仍是单线程。'),
(4, 'Spring Boot 自动配置原理是什么？', 'Spring Boot 通过 @EnableAutoConfiguration 注解开启自动配置，利用 SpringFactoriesLoader 加载 META-INF/spring.factories 中的配置类，根据 @Conditional 系列注解（@ConditionalOnClass、@ConditionalOnMissingBean 等）的条件判断来决定是否生效。核心是 spring-boot-autoconfigure 模块。', 3, false, 7650, 2400, 1600, 1, 'Spring Boot 通过 @EnableAutoConfiguration 注解开启自动配置，利用 SpringFactoriesLoader 加载 META-INF/spring.factories 中的配置类，根据 @Conditional 系列注解的条件判断来决定是否生效。核心是 spring-boot-autoconfigure 模块。', 'Spring Boot 自动配置原理是 Spring 全家桶的核心考点。需要理解 @Conditional 系列注解的条件判断机制，以及 spring.factories 的加载过程。建议结合自定义 Starter 来加深理解。'),
(5, 'TCP 三次握手和四次挥手的过程？', '三次握手：客户端发送 SYN=1, seq=x → 服务端回复 SYN=1, ACK=1, seq=y, ack=x+1 → 客户端发送 ACK=1, seq=x+1, ack=y+1。四次挥手：主动方发送 FIN → 被动方回复 ACK → 被动方发送 FIN → 主动方回复 ACK 并进入 TIME_WAIT 状态等待 2MSL。', 2, false, 7230, 2200, 1500, 1, '三次握手：客户端发送 SYN=1, seq=x → 服务端回复 SYN=1, ACK=1, seq=y, ack=x+1 → 客户端发送 ACK=1, seq=x+1, ack=y+1。四次挥手：主动方发送 FIN → 被动方回复 ACK → 被动方发送 FIN → 主动方回复 ACK 并进入 TIME_WAIT 状态等待 2MSL。', 'TCP 三次握手和四次挥手是网络基础必考题。三次握手的核心是确认双方的收发能力，四次挥手是因为 TCP 全双工通信需要双方各自关闭。TIME_WAIT 状态的作用和 2MSL 等待时间也是常见追问。'),
(6, '线程池的核心参数有哪些？', 'corePoolSize 核心线程数、maximumPoolSize 最大线程数、keepAliveTime 空闲线程存活时间、unit 时间单位、workQueue 工作队列、threadFactory 线程工厂、handler 拒绝策略（AbortPolicy、CallerRunsPolicy、DiscardPolicy、DiscardOldestPolicy）。执行流程：核心线程 → 队列 → 非核心线程 → 拒绝策略。', 2, false, 6890, 2100, 1400, 1, 'corePoolSize 核心线程数、maximumPoolSize 最大线程数、keepAliveTime 空闲线程存活时间、unit 时间单位、workQueue 工作队列、threadFactory 线程工厂、handler 拒绝策略。执行流程：核心线程 → 队列 → 非核心线程 → 拒绝策略。', '线程池是 Java 并发的核心考点。七大参数必须牢记，执行流程（核心线程→队列→非核心线程→拒绝策略）是面试重点。四种拒绝策略的区别和应用场景也需要掌握。实际项目中推荐使用自定义线程池而非 Executors 工具类。'),
(7, 'JVM 垃圾回收算法有哪些？', '标记-清除算法：产生内存碎片。复制算法：新生代使用，将内存分为 Eden 和两个 Survivor 区，效率高但浪费空间。标记-整理算法：老年代使用，解决碎片问题。分代收集算法：根据对象存活时间分代，不同代使用不同算法。常见垃圾收集器：Serial、Parallel、CMS、G1、ZGC。', 3, true, 6540, 1900, 1200, 1, '标记-清除算法：产生内存碎片。复制算法：新生代使用，将内存分为 Eden 和两个 Survivor 区，效率高但浪费空间。标记-整理算法：老年代使用，解决碎片问题。分代收集算法：根据对象存活时间分代，不同代使用不同算法。常见垃圾收集器：Serial、Parallel、CMS、G1、ZGC。', 'JVM 垃圾回收是 Java 高级面试必考题。四种算法各有优缺点，需要结合分代模型理解。CMS、G1、ZGC 等收集器的特点和使用场景是区分候选人水平的关键。建议结合 JVM 调优实战经验来回答。'),
(8, '什么是 CAP 定理？', 'CAP 定理指出分布式系统不可能同时满足一致性(Consistency)、可用性(Availability)和分区容错性(Partition tolerance)，最多只能同时满足其中两个。由于网络分区必然存在，所以实际是在 CP 和 AP 之间选择。ZooKeeper 是 CP 系统，Eureka 是 AP 系统。', 2, false, 6120, 1800, 1100, 1, 'CAP 定理指出分布式系统不可能同时满足一致性(Consistency)、可用性(Availability)和分区容错性(Partition tolerance)，最多只能同时满足其中两个。由于网络分区必然存在，所以实际是在 CP 和 AP 之间选择。ZooKeeper 是 CP 系统，Eureka 是 AP 系统。', 'CAP 定理是分布式系统的基础理论。面试中常结合具体系统（如 ZooKeeper、Eureka、Nacos）考察 CP 和 AP 的选择。还需要了解 BASE 理论和最终一致性概念。'),
(9, 'React Hooks 的使用注意事项？', '只能在函数组件顶层调用 Hook，不能在循环、条件或嵌套函数中调用。自定义 Hook 必须以 use 开头。useEffect 的依赖数组需要正确设置，避免闭包陷阱。useCallback 和 useMemo 用于性能优化。useRef 用于保存可变值且不触发重渲染。', 2, false, 5780, 1700, 1000, 1, '只能在函数组件顶层调用 Hook，不能在循环、条件或嵌套函数中调用。自定义 Hook 必须以 use 开头。useEffect 的依赖数组需要正确设置，避免闭包陷阱。useCallback 和 useMemo 用于性能优化。useRef 用于保存可变值且不触发重渲染。', 'React Hooks 是前端面试热点。除了基本使用规则，还需要理解闭包陷阱、依赖数组的正确设置、useCallback/useMemo 的性能优化原理。建议结合实际项目经验说明 Hooks 的最佳实践。'),
(10, '消息队列如何保证消息不丢失？', '生产者端：使用确认机制(confirm)、本地消息表保障消息发出。MQ端：持久化存储、集群部署、主从复制。消费者端：手动ACK确认、幂等性处理防止重复消费。RocketMQ 的事务消息可以保证分布式事务的最终一致性。', 2, false, 5430, 1600, 900, 1, '生产者端：使用确认机制(confirm)、本地消息表保障消息发出。MQ端：持久化存储、集群部署、主从复制。消费者端：手动ACK确认、幂等性处理防止重复消费。RocketMQ 的事务消息可以保证分布式事务的最终一致性。', '消息可靠性是消息队列的核心考点。需要从生产者、MQ、消费者三端分析消息丢失的场景和解决方案。RocketMQ 的事务消息和 Kafka 的 Exactly-Once 语义是进阶考点。'),
(11, 'ArrayList 和 LinkedList 的区别？', 'ArrayList 基于动态数组实现，随机访问 O(1)，插入删除需要移动元素 O(n)。LinkedList 基于双向链表实现，插入删除 O(1)（已知位置），随机访问 O(n)。ArrayList 内存占用更紧凑，LinkedList 每个节点额外存储前后指针。实际开发中 ArrayList 使用更多。', 1, false, 5120, 1500, 850, 1, 'ArrayList 基于动态数组实现，随机访问 O(1)，插入删除需要移动元素 O(n)。LinkedList 基于双向链表实现，插入删除 O(1)（已知位置），随机访问 O(n)。ArrayList 内存占用更紧凑，LinkedList 每个节点额外存储前后指针。实际开发中 ArrayList 使用更多。', 'ArrayList vs LinkedList 是 Java 集合的经典对比题。除了基本区别，面试官可能追问扩容机制、遍历方式、内存占用等细节。实际开发中 ArrayList 使用频率远高于 LinkedList，这是一个重要结论。'),
(12, '什么是乐观锁和悲观锁？', '乐观锁假设冲突概率低，通过版本号或 CAS（Compare And Swap）实现，适合读多写少场景，如数据库 version 字段。悲观锁假设冲突概率高，通过数据库 SELECT FOR UPDATE 或 Java synchronized 实现，适合写多场景。乐观锁可能产生 ABA 问题，可用 AtomicStampedReference 解决。', 1, false, 4890, 1400, 800, 1, '乐观锁假设冲突概率低，通过版本号或 CAS 实现，适合读多写少场景。悲观锁假设冲突概率高，通过 SELECT FOR UPDATE 或 synchronized 实现，适合写多场景。乐观锁可能产生 ABA 问题，可用 AtomicStampedReference 解决。', '乐观锁和悲观锁是并发控制的基础概念。数据库层面需要了解 version 字段和 SELECT FOR UPDATE 的用法。Java 层面需要掌握 CAS 原理和 ABA 问题。分布式锁（Redis、ZooKeeper）也是相关考点。'),
(13, 'HTTP 和 HTTPS 的区别？', 'HTTPS 在 HTTP 基础上增加了 SSL/TLS 加密层，数据传输加密防止窃听篡改。HTTPS 需要 CA 证书，默认端口 443（HTTP 为 80）。HTTPS 握手过程更复杂，需要密钥交换和身份验证。TLS 1.3 简化了握手流程，支持 0-RTT。', 1, false, 4650, 1300, 750, 1, 'HTTPS 在 HTTP 基础上增加了 SSL/TLS 加密层，数据传输加密防止窃听篡改。HTTPS 需要 CA 证书，默认端口 443（HTTP 为 80）。HTTPS 握手过程更复杂，需要密钥交换和身份验证。TLS 1.3 简化了握手流程，支持 0-RTT。', 'HTTP 和 HTTPS 的区别是网络基础考点。需要理解 SSL/TLS 的加密原理、证书验证流程、握手过程。TLS 1.3 的改进（0-RTT、简化握手）是加分项。HTTPS 的性能开销和优化方案也是常见追问。'),
(14, 'Vue 和 React 的区别？', 'Vue 采用模板语法，React 采用 JSX。Vue 数据双向绑定(MVVM)，React 单向数据流。Vue 使用 Composition API，React 使用 Hooks。Vue 内置路由和状态管理，React 需要第三方库。Vue 更易上手，React 生态更丰富。两者都支持虚拟 DOM 和组件化开发。', 2, false, 4420, 1200, 700, 1, 'Vue 采用模板语法，React 采用 JSX。Vue 数据双向绑定(MVVM)，React 单向数据流。Vue 使用 Composition API，React 使用 Hooks。Vue 内置路由和状态管理，React 需要第三方库。Vue 更易上手，React 生态更丰富。两者都支持虚拟 DOM 和组件化开发。', 'Vue vs React 是前端面试的高频话题。回答时避免偏袒，客观对比设计理念、生态、性能等方面的差异。重点展示对两种框架核心思想的理解，而不是简单罗列区别。'),
(15, '快速排序的时间复杂度？', '平均时间复杂度 O(nlogn)，最坏情况 O(n²)（数组已排序时）。空间复杂度 O(logn)（递归栈）。是不稳定排序。可通过随机化选择基准元素避免最坏情况。三路快排可以优化大量重复元素的情况。实际应用中通常结合插入排序，小数组时切换。', 2, false, 4180, 1100, 650, 1, '平均时间复杂度 O(nlogn)，最坏情况 O(n²)（数组已排序时）。空间复杂度 O(logn)（递归栈）。是不稳定排序。可通过随机化选择基准元素避免最坏情况。三路快排可以优化大量重复元素的情况。实际应用中通常结合插入排序，小数组时切换。', '快速排序是算法面试的经典题目。除了基本实现，需要掌握优化方法（随机化、三路快排、小数组切换插入排序）。最坏情况的分析和避免方法也是常考点。建议手写代码并分析复杂度。');

-- 5. 插入题目-标签关联
INSERT INTO question_tag (question_id, tag_id) VALUES
(1, 1), (1, 3), (1, 4),
(2, 5), (2, 6), (2, 7),
(3, 8), (3, 9), (3, 10),
(4, 1), (4, 2), (4, 10),
(5, 11), (5, 12), (5, 13),
(6, 1), (6, 14), (6, 15),
(7, 1), (7, 16), (7, 17),
(8, 18), (8, 19), (8, 20),
(9, 21), (9, 22), (9, 23),
(10, 24), (10, 25), (10, 18),
(11, 1), (11, 3), (11, 26),
(12, 6), (12, 30), (12, 31),
(13, 32), (13, 33), (13, 12),
(14, 34), (14, 21), (14, 35),
(15, 27), (15, 28), (15, 29);

-- 6. 插入题库-题目关联
INSERT INTO question_bank_question (bank_id, question_id) VALUES
(1, 1), (1, 6), (1, 7), (1, 11), (1, 4),
(2, 2), (2, 3), (2, 12),
(3, 5), (3, 13),
(4, 8), (4, 10),
(5, 9), (5, 14),
(6, 15);

-- 7. 插入用户刷题记录
INSERT INTO user_question_record (user_id, question_id, is_master, last_view_time) VALUES
(2, 1, true, '2026-04-16 10:00:00'),
(2, 2, false, '2026-04-16 11:00:00'),
(2, 5, true, '2026-04-16 12:00:00'),
(3, 1, true, '2026-04-15 09:00:00'),
(3, 3, true, '2026-04-15 10:00:00'),
(3, 6, false, '2026-04-15 11:00:00'),
(4, 1, true, '2026-04-14 08:00:00'),
(4, 2, true, '2026-04-14 09:00:00'),
(4, 3, true, '2026-04-14 10:00:00'),
(4, 5, false, '2026-04-14 11:00:00'),
(5, 1, false, '2026-04-16 14:00:00'),
(5, 2, true, '2026-04-16 15:00:00'),
(5, 11, true, '2026-04-16 16:00:00');

-- 8. 插入题目回答
INSERT INTO question_answer (question_id, user_id, content, like_count) VALUES
(1, 2, 'HashMap 底层是数组+链表+红黑树，JDK1.8之后当链表长度>=8时会转红黑树', 56),
(1, 3, 'HashMap 的 put 过程：先计算 hash，然后定位桶位置，如果桶为空直接插入，否则遍历链表或红黑树', 34),
(2, 4, 'B+树的非叶子节点只存键值，叶子节点通过链表连接，非常适合范围查询和磁盘存储', 42),
(3, 2, 'Redis 快的原因：纯内存操作、单线程避免上下文切换、IO多路复用、高效数据结构', 38),
(5, 5, '三次握手确保双方都有收发能力，四次挥手因为TCP是全双工的，需要双方各自关闭', 29),
(6, 3, '线程池7大参数：corePoolSize、maximumPoolSize、keepAliveTime、unit、workQueue、threadFactory、handler', 25),
(8, 4, 'CAP定理：分布式系统最多同时满足C、A、P中的两个，网络分区必然存在，所以实际在CP和AP之间选择', 31),
(9, 2, 'Hooks注意事项：只在顶层调用、只在React函数中调用、useEffect依赖要完整、自定义Hook以use开头', 22),
(10, 5, '消息不丢失三端保障：生产者confirm机制、MQ持久化、消费者手动ACK', 27),
(11, 4, 'ArrayList基于数组随机访问O(1)，LinkedList基于链表插入删除O(1)', 19);

-- 9. 插入用户收藏
INSERT INTO user_favorite (user_id, question_id) VALUES
(2, 1), (2, 3), (2, 5),
(3, 1), (3, 2), (3, 8),
(4, 1), (4, 2), (4, 3), (4, 5),
(5, 1), (5, 11);

-- 10. 插入用户搜索历史
INSERT INTO user_search_history (user_id, keyword) VALUES
(2, 'HashMap'), (2, 'MySQL索引'), (2, 'Redis'),
(3, '线程池'), (3, 'JVM'), (3, 'Spring Boot'),
(4, 'CAP'), (4, '消息队列'), (4, 'Redis'),
(5, 'ArrayList'), (5, 'HTTP');

-- 重置序列到最大ID
SELECT setval('user_id_seq', (SELECT MAX(id) FROM "user"));
SELECT setval('tag_id_seq', (SELECT MAX(id) FROM "tag"));
SELECT setval('question_id_seq', (SELECT MAX(id) FROM question));
SELECT setval('question_tag_id_seq', (SELECT MAX(id) FROM question_tag));
SELECT setval('question_bank_id_seq', (SELECT MAX(id) FROM question_bank));
SELECT setval('question_bank_question_id_seq', (SELECT MAX(id) FROM question_bank_question));
SELECT setval('question_answer_id_seq', (SELECT MAX(id) FROM question_answer));
SELECT setval('user_question_record_id_seq', (SELECT MAX(id) FROM user_question_record));
SELECT setval('user_favorite_id_seq', (SELECT MAX(id) FROM user_favorite));
SELECT setval('user_search_history_id_seq', (SELECT MAX(id) FROM user_search_history));
