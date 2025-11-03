package main

import (
	"container/list"
	"fmt"
)

// CacheEntry 用于存储键值对，以便在链表中引用 map 中的键
type CacheEntry struct {
	key   any
	value any
}

// LRUCache 定义了 LRU 缓存的结构
type LRUCache struct {
	capacity int
	ll       *list.List
	cache    map[any]*list.Element
}

// NewLRUCache 创建一个指定容量的 LRUCache 实例
// capacity: 缓存的最大容量
func NewLRUCache(capacity int) *LRUCache {
	if capacity <= 0 {
		// 保证容量是正数
		capacity = 1
	}
	return &LRUCache{
		capacity: capacity,
		ll:       list.New(),
		cache:    make(map[any]*list.Element),
	}
}

// Get 从缓存中获取一个值
// 如果键存在，则将该条目移动到链表头部并返回值
func (c *LRUCache) Get(key any) (value any, ok bool) {
	// 检查键是否存在于 map 中
	if elem, hit := c.cache[key]; hit {
		// 如果存在，将其移动到链表头部，表示最近使用过
		c.ll.MoveToFront(elem)
		// 返回找到的值
		return elem.Value.(*CacheEntry).value, true
	}
	return nil, false
}

// Put 向缓存中添加或更新一个键值对
// 如果键已存在，则更新其值并将其移动到链表头部
// 如果键不存在，则添加新条目；若缓存已满，则淘汰最久未使用的条目
func (c *LRUCache) Put(key any, value any) {
	// 检查键是否已存在
	if elem, hit := c.cache[key]; hit {
		// 更新值
		elem.Value.(*CacheEntry).value = value
		// 移动到链表头部，标记为最近使用
		c.ll.MoveToFront(elem)
	} else {
		// 如果缓存已满
		if c.ll.Len() >= c.capacity {
			// 获取并移除链表尾部的元素（最久未使用的）
			lastElem := c.ll.Back()
			if lastElem != nil {
				// 从 map 中删除对应的键
				delete(c.cache, lastElem.Value.(*CacheEntry).key)
				// 从链表中删除该元素
				c.ll.Remove(lastElem)
			}
		}
		// 创建新的缓存条目
		newEntry := &CacheEntry{key: key, value: value}
		// 将新条目添加到链表头部
		newElem := c.ll.PushFront(newEntry)
		// 在 map 中建立键与链表元素的映射
		c.cache[key] = newElem
	}
}

func main() {
	// 创建一个容量为 2 的 LRU 缓存
	lru := NewLRUCache(2)

	fmt.Println("Putting key1: value1")
	lru.Put("key1", "value1")
	fmt.Println("Putting key2: value2")
	lru.Put("key2", "value2")

	// 访问 key1，使其变为最近使用的
	fmt.Println("Getting key1...")
	value, ok := lru.Get("key1")
	if ok {
		fmt.Printf("Got key1: %v\n", value)
	}

	// 添加 key3，此时缓存已满，key2 应该被淘汰
	fmt.Println("Putting key3: value3 (should evict key2)")
	lru.Put("key3", "value3")

	// 尝试获取 key2，应该找不到了
	fmt.Println("Getting key2...")
	value, ok = lru.Get("key2")
	if !ok {
		fmt.Println("key2 not found (as expected)")
	}

	// 获取 key1 和 key3，应该都能找到
	fmt.Println("Getting key1...")
	value, ok = lru.Get("key1")
	if ok {
		fmt.Printf("Got key1: %v\n", value)
	}

	fmt.Println("Getting key3...")
	value, ok = lru.Get("key3")
	if ok {
		fmt.Printf("Got key3: %v\n", value)
	}
}
