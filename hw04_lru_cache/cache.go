package hw04lrucache

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type lruCache struct {
	Cache // Remove me after realization.

	capacity int
	queue    List
	items    map[Key]*ListItem
}

type cacheItem struct {
	key   Key
	value interface{}
}

func NewCache(capacity int) Cache {
	return &lruCache{
		capacity: capacity,
		queue:    NewList(),
		items:    make(map[Key]*ListItem, capacity),
	}
}

func (c *lruCache) Set(key Key, value interface{}) (exists bool) {
	item, exists := c.items[key]
	if exists {
		item.Value = value
		c.queue.MoveToFront(item)
		return
	}

	item = c.queue.PushFront(value)
	c.items[key] = item

	if c.queue.Len() > c.capacity {
		last := c.queue.Back()
		var lastKey Key
		for k, v := range c.items {
			if v == last.Value {
				lastKey = k
				break
			}
		}
		delete(c.items, lastKey)
		c.queue.Remove(c.queue.Back())
	}

	return
}

func (c *lruCache) Get(key Key) (val interface{}, exists bool) {
	item, exists := c.items[key]
	if !exists {
		val = nil
		return
	}

	val = item.Value

	c.queue.MoveToFront(item)
	return
}

func (c *lruCache) Clear() {
	c.items = make(map[Key]*ListItem, c.capacity)
	c.queue = NewList()
}
