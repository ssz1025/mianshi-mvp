ALTER TABLE question ADD COLUMN IF NOT EXISTS answer text DEFAULT '';
ALTER TABLE question ADD COLUMN IF NOT EXISTS explanation text DEFAULT '';
ALTER TABLE question ADD COLUMN IF NOT EXISTS bank_id bigint DEFAULT 0;
ALTER TABLE question ADD COLUMN IF NOT EXISTS heat int DEFAULT 0;
ALTER TABLE question ADD COLUMN IF NOT EXISTS answered_count int DEFAULT 0;
ALTER TABLE question ADD COLUMN IF NOT EXISTS correct_rate int DEFAULT 0;

UPDATE question SET
  answer = CASE id
    WHEN 1 THEN 'HashMap 基于哈希表实现，JDK 1.8 之后采用数组+链表+红黑树的数据结构。当链表长度超过 8 且数组长度超过 64 时，链表会转换为红黑树以提高查询效率。put 过程：先计算 hash 值定位桶位置，如果桶为空直接插入，否则遍历链表或红黑树进行更新或插入。扩容时默认容量翻倍，重新计算每个元素的位置。'
    WHEN 2 THEN 'B+ 树相比 B 树和二叉树，具有更矮的树高，磁盘 I/O 次数更少。叶子节点通过链表相连，支持范围查询。非叶子节点只存储键值，单个节点能存储更多键值，从而降低树的高度。B+ 树的查询性能稳定，每次查询都需要走到叶子节点，路径长度相同。'
    WHEN 3 THEN 'Redis 基于内存操作，使用单线程避免上下文切换和锁竞争，采用 IO 多路复用模型（epoll），高效的数据结构设计（如 SDS、跳表、压缩列表、整数集合等）。此外 Redis 6.0+ 还支持多线程 I/O 读写，但命令执行仍是单线程。'
    WHEN 4 THEN 'Spring Boot 通过 @EnableAutoConfiguration 注解开启自动配置，利用 SpringFactoriesLoader 加载 META-INF/spring.factories 中的配置类，根据 @Conditional 系列注解的条件判断来决定是否生效。核心是 spring-boot-autoconfigure 模块。'
    WHEN 5 THEN '三次握手：客户端发送 SYN=1, seq=x → 服务端回复 SYN=1, ACK=1, seq=y, ack=x+1 → 客户端发送 ACK=1, seq=x+1, ack=y+1。四次挥手：主动方发送 FIN → 被动方回复 ACK → 被动方发送 FIN → 主动方回复 ACK 并进入 TIME_WAIT 状态等待 2MSL。'
    WHEN 6 THEN 'corePoolSize 核心线程数、maximumPoolSize 最大线程数、keepAliveTime 空闲线程存活时间、unit 时间单位、workQueue 工作队列、threadFactory 线程工厂、handler 拒绝策略。执行流程：核心线程 → 队列 → 非核心线程 → 拒绝策略。'
    WHEN 7 THEN '标记-清除算法：产生内存碎片。复制算法：新生代使用，将内存分为 Eden 和两个 Survivor 区，效率高但浪费空间。标记-整理算法：老年代使用，解决碎片问题。分代收集算法：根据对象存活时间分代，不同代使用不同算法。常见垃圾收集器：Serial、Parallel、CMS、G1、ZGC。'
    WHEN 8 THEN 'CAP 定理指出分布式系统不可能同时满足一致性(Consistency)、可用性(Availability)和分区容错性(Partition tolerance)，最多只能同时满足其中两个。由于网络分区必然存在，所以实际是在 CP 和 AP 之间选择。ZooKeeper 是 CP 系统，Eureka 是 AP 系统。'
    WHEN 9 THEN '只能在函数组件顶层调用 Hook，不能在循环、条件或嵌套函数中调用。自定义 Hook 必须以 use 开头。useEffect 的依赖数组需要正确设置，避免闭包陷阱。useCallback 和 useMemo 用于性能优化。useRef 用于保存可变值且不触发重渲染。'
    WHEN 10 THEN '生产者端：使用确认机制(confirm)、本地消息表保障消息发出。MQ端：持久化存储、集群部署、主从复制。消费者端：手动ACK确认、幂等性处理防止重复消费。RocketMQ 的事务消息可以保证分布式事务的最终一致性。'
    WHEN 11 THEN 'ArrayList 基于动态数组实现，随机访问 O(1)，插入删除需要移动元素 O(n)。LinkedList 基于双向链表实现，插入删除 O(1)（已知位置），随机访问 O(n)。ArrayList 内存占用更紧凑，LinkedList 每个节点额外存储前后指针。实际开发中 ArrayList 使用更多。'
    WHEN 12 THEN '乐观锁假设冲突概率低，通过版本号或 CAS 实现，适合读多写少场景。悲观锁假设冲突概率高，通过 SELECT FOR UPDATE 或 synchronized 实现，适合写多场景。乐观锁可能产生 ABA 问题，可用 AtomicStampedReference 解决。'
    WHEN 13 THEN 'HTTPS 在 HTTP 基础上增加了 SSL/TLS 加密层，数据传输加密防止窃听篡改。HTTPS 需要 CA 证书，默认端口 443（HTTP 为 80）。HTTPS 握手过程更复杂，需要密钥交换和身份验证。TLS 1.3 简化了握手流程，支持 0-RTT。'
    WHEN 14 THEN 'Vue 采用模板语法，React 采用 JSX。Vue 数据双向绑定(MVVM)，React 单向数据流。Vue 使用 Composition API，React 使用 Hooks。Vue 内置路由和状态管理，React 需要第三方库。Vue 更易上手，React 生态更丰富。两者都支持虚拟 DOM 和组件化开发。'
    WHEN 15 THEN '平均时间复杂度 O(nlogn)，最坏情况 O(n²)（数组已排序时）。空间复杂度 O(logn)（递归栈）。是不稳定排序。可通过随机化选择基准元素避免最坏情况。三路快排可以优化大量重复元素的情况。实际应用中通常结合插入排序，小数组时切换。'
    ELSE ''
  END,
  explanation = CASE id
    WHEN 1 THEN 'HashMap 是面试高频考点，核心在于理解哈希冲突的解决方式和扩容机制。JDK 1.8 引入红黑树是重要优化，需要掌握链表转红黑树的阈值条件。还需要了解 HashMap 的线程安全问题，为什么推荐使用 ConcurrentHashMap 替代。'
    WHEN 2 THEN 'B+ 树是数据库索引的核心数据结构，面试常考。需要对比 B 树、B+ 树、红黑树、跳表的优劣。重点理解为什么数据库不用红黑树（树太高，磁盘 I/O 多）和为什么不用跳表（范围查询不如 B+ 树高效）。'
    WHEN 3 THEN 'Redis 高性能是经典面试题。需要从内存、线程模型、I/O 模型、数据结构四个维度分析。Redis 6.0 的多线程 I/O 是常考点，需要理解多线程只用于网络 I/O，命令执行仍是单线程。'
    WHEN 4 THEN 'Spring Boot 自动配置原理是 Spring 全家桶的核心考点。需要理解 @Conditional 系列注解的条件判断机制，以及 spring.factories 的加载过程。建议结合自定义 Starter 来加深理解。'
    WHEN 5 THEN 'TCP 三次握手和四次挥手是网络基础必考题。三次握手的核心是确认双方的收发能力，四次挥手是因为 TCP 全双工通信需要双方各自关闭。TIME_WAIT 状态的作用和 2MSL 等待时间也是常见追问。'
    WHEN 6 THEN '线程池是 Java 并发的核心考点。七大参数必须牢记，执行流程（核心线程→队列→非核心线程→拒绝策略）是面试重点。四种拒绝策略的区别和应用场景也需要掌握。实际项目中推荐使用自定义线程池而非 Executors 工具类。'
    WHEN 7 THEN 'JVM 垃圾回收是 Java 高级面试必考题。四种算法各有优缺点，需要结合分代模型理解。CMS、G1、ZGC 等收集器的特点和使用场景是区分候选人水平的关键。建议结合 JVM 调优实战经验来回答。'
    WHEN 8 THEN 'CAP 定理是分布式系统的基础理论。面试中常结合具体系统（如 ZooKeeper、Eureka、Nacos）考察 CP 和 AP 的选择。还需要了解 BASE 理论和最终一致性概念。'
    WHEN 9 THEN 'React Hooks 是前端面试热点。除了基本使用规则，还需要理解闭包陷阱、依赖数组的正确设置、useCallback/useMemo 的性能优化原理。建议结合实际项目经验说明 Hooks 的最佳实践。'
    WHEN 10 THEN '消息可靠性是消息队列的核心考点。需要从生产者、MQ、消费者三端分析消息丢失的场景和解决方案。RocketMQ 的事务消息和 Kafka 的 Exactly-Once 语义是进阶考点。'
    WHEN 11 THEN 'ArrayList vs LinkedList 是 Java 集合的经典对比题。除了基本区别，面试官可能追问扩容机制、遍历方式、内存占用等细节。实际开发中 ArrayList 使用频率远高于 LinkedList，这是一个重要结论。'
    WHEN 12 THEN '乐观锁和悲观锁是并发控制的基础概念。数据库层面需要了解 version 字段和 SELECT FOR UPDATE 的用法。Java 层面需要掌握 CAS 原理和 ABA 问题。分布式锁（Redis、ZooKeeper）也是相关考点。'
    WHEN 13 THEN 'HTTP 和 HTTPS 的区别是网络基础考点。需要理解 SSL/TLS 的加密原理、证书验证流程、握手过程。TLS 1.3 的改进（0-RTT、简化握手）是加分项。HTTPS 的性能开销和优化方案也是常见追问。'
    WHEN 14 THEN 'Vue vs React 是前端面试的高频话题。回答时避免偏袒，客观对比设计理念、生态、性能等方面的差异。重点展示对两种框架核心思想的理解，而不是简单罗列区别。'
    WHEN 15 THEN '快速排序是算法面试的经典题目。除了基本实现，需要掌握优化方法（随机化、三路快排、小数组切换插入排序）。最坏情况的分析和避免方法也是常考点。建议手写代码并分析复杂度。'
    ELSE ''
  END,
  bank_id = CASE id
    WHEN 1 THEN 1
    WHEN 2 THEN 2
    WHEN 3 THEN 2
    WHEN 4 THEN 1
    WHEN 5 THEN 3
    WHEN 6 THEN 1
    WHEN 7 THEN 1
    WHEN 8 THEN 4
    WHEN 9 THEN 5
    WHEN 10 THEN 4
    WHEN 11 THEN 1
    WHEN 12 THEN 2
    WHEN 13 THEN 3
    WHEN 14 THEN 5
    WHEN 15 THEN 6
    ELSE 0
  END,
  heat = view_count,
  answered_count = like_count,
  correct_rate = 80
WHERE id <= 15;
