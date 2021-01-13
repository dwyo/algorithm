package leetcode

import (
	"math/rand"
)

//设计一个支持在平均 时间复杂度 O(1) 下，执行以下操作的数据结构。
//
// 
// insert(val)：当元素 val 不存在时，向集合中插入该项。 
// remove(val)：元素 val 存在时，从集合中移除该项。 
// getRandom：随机返回现有集合中的一项。每个元素应该有相同的概率被返回。 
// 
//
// 示例 : 
//
// 
//// 初始化一个空的集合。
//RandomizedSet randomSet = new RandomizedSet();
//
//// 向集合中插入 1 。返回 true 表示 1 被成功地插入。
//randomSet.insert(1);
//
//// 返回 false ，表示集合中不存在 2 。
//randomSet.remove(2);
//
//// 向集合中插入 2 。返回 true 。集合现在包含 [1,2] 。
//randomSet.insert(2);
//
//// getRandom 应随机返回 1 或 2 。
//randomSet.getRandom();
//
//// 从集合中移除 1 ，返回 true 。集合现在包含 [2] 。
//randomSet.remove(1);
//
//// 2 已在集合中，所以返回 false 。
//randomSet.insert(2);
//
//// 由于 2 是集合中唯一的数字，getRandom 总是返回 2 。
//randomSet.getRandom();
// 
// Related Topics 设计 数组 哈希表 
// 👍 254 👎 0

//leetcode submit region begin(Prohibit modification and deletion)
type RandomizedSet struct {
	data []int // 存储值
	size int 	// 存储长度
	mp   map[int]int // 存储索引
}

/** Initialize your data structure here. */
func Constructor() RandomizedSet {
	return RandomizedSet{make([]int, 0), 0, make(map[int]int)}
}

/** Inserts a value to the set. Returns true if the set did not already contain the specified element. */
func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.mp[val]; ok {
		return false
	}
	if this.size < len(this.data) {
		this.data[this.size] = val
	} else {
		this.data = append(this.data, val)
	}
	this.mp[val] = this.size
	this.size++

	//fmt.Println("insert: ", val, this.mp, this.data, this.size)

	return true
}

/** Removes a value from the set. Returns true if the set contained the specified element. */
func (this *RandomizedSet) Remove(val int) bool {
	idx, ok := this.mp[val]
	if !ok {
		return false
	}
	// 把要删除元素 和 最后一位元素交换位置，然后删除最后一位
	temp := this.data[this.size - 1]
	this.data[idx], this.data[this.size - 1] = this.data[this.size - 1], this.data[idx]
	this.mp[temp] = idx
	delete(this.mp, val)
	this.size--
	//fmt.Println("remove: ", val, this.mp, this.data, this.size)
	return true
}

/** Get a random element from the set. */
func (this *RandomizedSet) GetRandom() int {
	if this.size == -1 {
		return -1
	}
//fmt.Println("random: ", this.mp, this.data, this.size)
	idx := rand.Intn(this.size)
	return this.data[idx]
}

/**
 * Your RandomizedSet object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
//leetcode submit region end(Prohibit modification and deletion)
