package some

const (
	hostBit = uint64(^uint(0)) == ^uint64(0)
	LENGTH  = 100
)

/**
LRU 节点
*/
type LruNode struct {
	pre  *LruNode
	next *LruNode

	key int
	val int

	hNext *LruNode
}

type LruCache struct {
	node []LruNode // 存放容器

	head *LruNode
	tail *LruNode

	capacity int // 容量
	used     int // 使用数量
}

/**
构造函数
*/
func Constructor(capacity int) *LruCache {
	return &LruCache{
		node:     make([]LruNode, LENGTH),
		head:     nil,
		tail:     nil,
		capacity: capacity,
		used:     0,
	}
}

func (c *LruCache) Get(key int) int {
	if c.tail == nil {
		return -1
	}

	if tmp := c.Search(key); tmp != nil {
		c.moveToTail(tmp)
		return tmp.val
	}

	return -1
}

func (c *LruCache) Put(key, val int) {

	// 需要满足一下条件
	// 1. 首次插入数据
	// 2. 插入数据不在 LRU 中
	// 3. 插入数据在 LRU 中
	// 4. 插入数据不在 LRU 中, 并且 LRU 已满

	if tmp := c.Search(key); tmp != nil {
		tmp.val = val
		c.moveToTail(tmp)
		return
	}

	c.AddNode(key, val)

	if c.used > c.capacity {
		c.DelNode()
	}
}

/**
最新的数据放在尾部
*/
func (c *LruCache) AddNode(key, val int) {
	newNode := &LruNode{
		key: key,
		val: val,
	}

	tmp := c.node[hash(key)]

	newNode.hNext = tmp.hNext
	tmp.hNext = newNode
	c.used++

	if c.tail == nil {
		c.head, c.tail = newNode, newNode
	}
	c.tail.next = newNode
	newNode.pre = c.tail
	c.tail = newNode
}

/**
删除头部节点
 */
func (c *LruCache) DelNode() {
	if c.head == nil {
		return
	}

	pre := &c.node[hash(c.head.key)]
	tmp := pre.hNext
	
	for tmp != nil && tmp.key != c.head.key {
		pre = tmp
		tmp = tmp.hNext
	}
	if tmp == nil {
		return
	}
	pre.hNext = tmp.hNext
	c.head = c.head.next
	c.head.pre = nil
	c.used--
}

func (c *LruCache) Search(key int) *LruNode {
	if c.tail == nil {
		return nil
	}

	tmp := c.node[hash(key)].hNext
	if tmp != nil {
		if tmp.key == key {
			return tmp
		}
		tmp = tmp.hNext
	}
	return nil
}

/**
每访问一次，都把数据挪到尾部，则最少使用的数据在头部
*/
func (c *LruCache) moveToTail(node *LruNode) {
	if c.tail == nil {
		return
	}
	if c.head == node {
		c.head = node.next
		c.head.pre = nil
	} else {
		node.next.pre = node.pre
		node.pre.next = node.next
	}
	node.next = nil
	c.tail.next = node
	node.pre = c.tail

	c.tail = node
}

func hash(key int) int {
	if hostBit {
		// 对 LENGTH 进行取模运算
		return (key ^ (key >> 32)) & (LENGTH - 1)
	}
	return (key ^ (key >> 16)) & (LENGTH - 1)
}
