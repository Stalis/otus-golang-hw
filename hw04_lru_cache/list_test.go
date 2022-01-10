package hw04lrucache

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestList(t *testing.T) {
	t.Run("empty list", func(t *testing.T) {
		l := NewList()

		require.Equal(t, 0, l.Len())
		require.Nil(t, l.Front())
		require.Nil(t, l.Back())
	})

	t.Run("complex", func(t *testing.T) {
		l := NewList()

		l.PushFront(10) // [10]
		l.PushBack(20)  // [10, 20]
		l.PushBack(30)  // [10, 20, 30]
		require.Equal(t, 3, l.Len())

		middle := l.Front().Next // 20
		l.Remove(middle)         // [10, 30]
		require.Equal(t, 2, l.Len())

		for i, v := range [...]int{40, 50, 60, 70, 80} {
			if i%2 == 0 {
				l.PushFront(v)
			} else {
				l.PushBack(v)
			}
		} // [80, 60, 40, 10, 30, 50, 70]

		require.Equal(t, 7, l.Len())
		require.Equal(t, 80, l.Front().Value)
		require.Equal(t, 70, l.Back().Value)

		l.MoveToFront(l.Front()) // [80, 60, 40, 10, 30, 50, 70]
		l.MoveToFront(l.Back())  // [70, 80, 60, 40, 10, 30, 50]

		elems := make([]int, 0, l.Len())
		for i := l.Front(); i != nil; i = i.Next {
			elems = append(elems, i.Value.(int))
		}
		require.Equal(t, []int{70, 80, 60, 40, 10, 30, 50}, elems)
	})

	t.Run("test PushFront()", func(t *testing.T) {
		l := NewList()

		l.PushFront(10)
		require.Equal(t, 10, l.Front().Value)
		require.Equal(t, 10, l.Back().Value)
		require.Equal(t, 1, l.Len())

		l.PushFront(20)
		require.Equal(t, 20, l.Front().Value)
		require.Equal(t, 10, l.Back().Value)
		require.Equal(t, 2, l.Len())

		l.PushFront(30)
		require.Equal(t, 30, l.Front().Value)
		require.Equal(t, 10, l.Back().Value)
		require.Equal(t, 3, l.Len())
	})

	t.Run("test PushBack()", func(t *testing.T) {
		l := NewList()

		l.PushBack(10)
		require.Equal(t, 10, l.Back().Value)
		require.Equal(t, 10, l.Front().Value)
		require.Equal(t, 1, l.Len())

		l.PushBack(20)
		require.Equal(t, 20, l.Back().Value)
		require.Equal(t, 10, l.Front().Value)
		require.Equal(t, 2, l.Len())

		l.PushBack(30)
		require.Equal(t, 30, l.Back().Value)
		require.Equal(t, 10, l.Front().Value)
		require.Equal(t, 3, l.Len())
	})

	t.Run("test Remove()", func(t *testing.T) {
		l := NewList()

		first := l.PushBack(10)
		second := l.PushBack(20)
		third := l.PushBack(30)
		require.Equal(t, 3, l.Len())

		l.Remove(first)
		require.Equal(t, 2, l.Len())

		require.Equal(t, second, l.Front())
		require.Equal(t, third, l.Back())

		require.Nil(t, first.Next)
		require.Nil(t, first.Prev)
		require.Nil(t, second.Prev)

		first = l.PushFront(first.Value)
		require.Equal(t, 3, l.Len())

		l.Remove(second)
		require.Equal(t, 2, l.Len())

		require.Equal(t, first, l.Front())
		require.Equal(t, third, l.Back())

		require.Nil(t, second.Next)
		require.Nil(t, second.Prev)
		require.Equal(t, third, first.Next)
		require.Equal(t, first, third.Prev)

		l.Remove(third)
		require.Equal(t, 1, l.Len())

		require.Equal(t, first, l.Front())
		require.Equal(t, first, l.Back())
		require.Nil(t, l.Front().Next)
		require.Nil(t, l.Back().Prev)

		l.Remove(first)
		require.Equal(t, 0, l.Len())

		require.Nil(t, l.Front())
		require.Nil(t, l.Back())
	})
}
