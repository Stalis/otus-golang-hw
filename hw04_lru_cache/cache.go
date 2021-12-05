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

func (c *lruCache) Set(key Key, value interface{}) bool {
	if c == nil {
		return false
	}
	return c.set(cacheItem{key, value})
}

func (c *lruCache) set(i cacheItem) (exists bool) {
	item, exists := c.items[i.key]
	if exists {
		item.Value = i
		c.queue.MoveToFront(item)
		return
	}

	item = c.queue.PushFront(i)
	c.items[i.key] = item

	if c.queue.Len() > c.capacity {
		last := c.queue.Back()

		delete(c.items, last.Value.(cacheItem).key)
		c.queue.Remove(last)
	}

	return
}

func (c *lruCache) Get(key Key) (val interface{}, exists bool) {
	item, exists := c.items[key]
	if exists {
		val = item.Value.(cacheItem).value
		c.queue.MoveToFront(item)
	}

	return
}

func (c *lruCache) Clear() {
	c.items = make(map[Key]*ListItem, c.capacity)
	c.queue = NewList()
}
