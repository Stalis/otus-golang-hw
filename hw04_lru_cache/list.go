package hw04lrucache

type List interface {
	Len() int
	Front() *ListItem
	Back() *ListItem
	PushFront(v interface{}) *ListItem
	PushBack(v interface{}) *ListItem
	Remove(i *ListItem)
	MoveToFront(i *ListItem)
}

type ListItem struct {
	Value interface{}
	Next  *ListItem
	Prev  *ListItem
}

type list struct {
	first *ListItem
	last  *ListItem
	count int
}

func NewList() List {
	return new(list)
}

func (l list) Len() int {
	return l.count
}

func (l list) Front() *ListItem {
	return l.first
}

func (l list) Back() *ListItem {
	return l.last
}

func (l *list) PushFront(v interface{}) *ListItem {
	if l == nil {
		return nil
	}
	return l.pushFront(&ListItem{Value: v})
}

func (l *list) pushFront(item *ListItem) *ListItem {
	if l.first == nil {
		l.first = item
		l.last = item
	} else {
		item.Next = l.first
		l.first.Prev = item
		l.first = item
	}

	l.count++
	return item
}

func (l *list) PushBack(v interface{}) *ListItem {
	if l == nil {
		return nil
	}
	return l.pushBack(&ListItem{Value: v})
}

func (l *list) pushBack(item *ListItem) *ListItem {
	if l.last == nil {
		l.last = item
		l.first = item
	} else {
		item.Prev = l.last
		l.last.Next = item
		l.last = item
	}

	l.count++
	return l.last
}

func (l *list) Remove(i *ListItem) {
	if l == nil || i == nil {
		return
	}
	l.remove(i)
}

func (l *list) remove(i *ListItem) {
	if l.first == i {
		l.first = i.Next
	}
	if l.last == i {
		l.last = i.Prev
	}

	if i.Prev != nil {
		i.Prev.Next = i.Next
	}
	if i.Next != nil {
		i.Next.Prev = i.Prev
	}

	i.Next = nil
	i.Prev = nil

	l.count--
}

func (l *list) MoveToFront(i *ListItem) {
	if l == nil || i == nil {
		return
	}
	l.moveToFront(i)
}

func (l *list) moveToFront(i *ListItem) {
	if i == l.first {
		return
	}
	l.remove(i)
	l.pushFront(i)
}
