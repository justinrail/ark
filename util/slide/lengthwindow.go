package slide

import (
	"container/list"
	"sync"
)

//LengthWindow 长度滑动窗口，（有限长度的可抛弃数据队列）, 默认删除最旧数据
type LengthWindow struct {
	data    *list.List
	Length  int
	lock    sync.Mutex
	alerter func(interface{})
	Total   int
	Lost    int
}

//NewLengthWindow 创建新对象,有最大长度限制
func NewLengthWindow(length int) *LengthWindow {
	q := new(LengthWindow)
	q.data = list.New()
	q.Length = length
	return q
}

//Enqueue Enqueue
func (q *LengthWindow) Enqueue(v interface{}) {
	defer q.lock.Unlock()
	q.lock.Lock()
	q.Total++
	q.data.PushBack(v)
	if q.data.Len() > q.Length {
		e := q.data.Front()
		if e != nil {
			q.data.Remove(e)
			q.Lost++
			if q.alerter != nil {
				q.alerter(e)
			}
		}
	}
}

//Dequeue Dequeue
func (q *LengthWindow) Dequeue() interface{} {
	defer q.lock.Unlock()
	q.lock.Lock()
	iter := q.data.Front()
	if iter == nil {
		return nil
	}

	v := iter.Value
	q.data.Remove(iter)
	return v
}

//ForEach 遍历窗口的队列
func (q *LengthWindow) ForEach(hooker func(interface{})) {
	defer q.lock.Unlock()
	q.lock.Lock()
	for iter := q.data.Back(); iter != nil; iter = iter.Prev() {
		hooker(iter.Value)
	}
}

//SetAlerter 设置删除数据的警告函数
func (q *LengthWindow) SetAlerter(hooker func(interface{})) {
	defer q.lock.Unlock()
	q.lock.Lock()
	q.alerter = hooker
}

//Count window当前的数据量
func (q *LengthWindow) Count() int {
	defer q.lock.Unlock()
	q.lock.Lock()
	return q.data.Len()
}
