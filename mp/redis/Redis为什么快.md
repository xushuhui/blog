## 前言
我们对Redis的深刻印象就是"快"，它在接收到一个键值对操作指令后在微妙内完成操作。
为什么它能这么快，一方面它是在内存中进行操作，内存访问本身速度快，另一方面是它有高效的数据结构。键值对是按一定的数据结构存储，操作键值对就是对数据结构增删改查，高效的数据结构是Redis快速处理数据的基础。

## 底层数据结构
![底层数据结构](https://cdn.guojiang.club/FrtB8lP_rbaqaxk3hcOfmbo01TFY)
Redis的底层数据结构有六种，简单动态字符串、双向链表、压缩列表、哈希表、跳表和整数数组，String的底层实现是简单动态字符串，List、Hash、Set和SortedSet都有两种底层实现结构，这四种类型被称为集合类型，特点是一个key对应一个集合数据

### 键和值的数据结构是什么
Redis用一个哈希表保存所有键值对，实现key-value快速访问。

一个哈希表就是一个数组，数组每个元素叫哈希桶，每个哈希桶保存键值对数据。然而哈希桶中的元素不是value本身，而是指向value的指针，即value存储的内存地址。

![全局哈希表](https://cdn.guojiang.club/Fuu5tihTIyL9ypByeYs3ZitCCIQG)

如图，这个哈希表保存了所有键值对，哈希桶中的entry元素保存*key和*value指针，哈希表能在O(1)时间复杂度快速查找键值对，所以我们只需要计算key的哈希值就能找到对应的哈希桶位置，进而找到对应的entry元素。不同类型的value都能被找到，不论是String、List、Set、Hash。

这种查找方式只需要进行一次哈希计算，不论数据规模多少，然而，在Redis中写入大量数据后，操作有时候会变慢，因为出现了哈希表的冲突以及rehash带来的操作阻塞。

### 哈希冲突
当哈希表中数据增加，新增的数据key哈希计算出的哈希值和老数据key的哈希值会在同一个哈希桶中，也就是说多个key对应同一个哈希桶。
#### 链式哈希
Redis中，同一个哈希桶中多个元素用一个链表保存，它们之间用指针连接，这就是链式哈希。

如图所示，entry1、entry2和entry3都保存在哈希桶3中，导致哈希冲突。entry1增加个*next指针指向entry2，entry2增加*next指针指向entry3，不论哈希桶3元素有多少个，都可以通过指针连接起来，形成一个链表，叫做哈希冲突链。

![](https://cdn.guojiang.club/FtHcgEDV22B5bNyKPQnNDkkK_PPF)

链式哈希会产生一个问题，随着哈希表数据越来越多，哈希冲突越来越多，单个哈希桶链表上数据越来越多，查找时间复杂度退化到O(n)，查找耗时增加，效率降低。
#### rehash
为解决这个问题，Redis会对哈希表做rehash操作。rehash 也就是增加现有的哈希桶数量，让逐渐增多的 entry 元素能在更多的桶之间分散保存，减少单个桶中的元素数量，从而减少单个桶中的冲突。

Redis使用两个全局哈希表：哈希表1和哈希表2，最开始新增数据默认存到哈希表1，哈希表2没有被分配空间，当数据增加，Redis开始执行Rehash操作：
1. 给哈希表2分配更大空间，可以是当前哈希表1大小的两倍
2. 把哈希表1的数据重新映射并拷贝到哈希表2
3. 释放哈希表1空间

rehash后，从哈希表1切换到哈希表2，哈希表2空间更多，哈希冲突更少，原来哈希表1留做下次rehash扩容备用，按同样的步骤把哈希表2的数据迁移到哈希表1。

在第二步涉及大量数据拷贝，如果一次性把哈希表1迁移完，耗时很长，会造成线程阻塞，无法处理其他请求，Redis是怎么处理这个问题呢？它采用渐进式rehash

#### 渐进式rehash
在第二步中，Redis正常处理客户端请求，每处理一个请求，从哪哈希表1的第一个索引位置开始，把这个位置上的所有entry拷贝到哈希表2中。处理下一个请求时，把下一个索引位置的entry做同样操作。

![](https://cdn.guojiang.club/FkWDvfrCfFPxwpAjC5p9luENfiN9)


![](https://cdn.guojiang.club/Fi2ZO7lmxpHAsP0u1Cdm7jjsZ7KY)

![](https://cdn.guojiang.club/FmSmiFmTJYeWNa3kD0cBp2Bb0pMl)

![](https://cdn.guojiang.club/FujbEZ-MkzawIUyXzkykPr-nLQIn)


