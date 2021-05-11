// Code generated by go2go; DO NOT EDIT.


//line linkedlist_test.go2:1
package collection

//line linkedlist_test.go2:1
import (
//line linkedlist_test.go2:1
 "math/rand"
//line linkedlist_test.go2:1
 "reflect"
//line linkedlist_test.go2:1
 "testing"
//line linkedlist_test.go2:1
 "time"
//line linkedlist_test.go2:1
)

//line linkedlist_test.go2:10
func TestLinkedListAPI(t *testing.T) {

	l := instantiate୦୦CreateLinkedList୦int()

	rand.Seed(time.Now().UnixNano())
	size := 10

	array := make([]int, size)

	for i := 0; i < size; i++ {
		array[i] = rand.Intn(10000)
	}

//line linkedlist_test.go2:24
 for i := 0; i < 10; i++ {
		l.Append(array[i])
	}

	if l.Empty() {
		t.Fatalf("unexpected empty!")
	}

	if l.Len() != 10 {
		t.Fatalf("unexpected length!")
	}

	if l.Front() != array[0] {
		t.Fatalf("unexpected Front!")
	}

	if l.Back() != array[9] {
		t.Fatalf("unexpected Back!")
	}

//line linkedlist_test.go2:45
 if !reflect.DeepEqual(l.Traverse(), array[:10]) {
		t.Fatalf("unexpected Traverse!")
	}

//line linkedlist_test.go2:50
 l.Set(1050, 0)
	if l.Front() != 1050 {
		t.Fatalf("unexpected Set!")
	}

	l.Set(1050, -1)
	if l.Back() != 1050 {
		t.Fatalf("unexpected Set!")
	}

//line linkedlist_test.go2:61
 l.Remove(-1)
	if l.Len() != 9 {
		t.Fatalf("unexpected Remove!")
	}

//line linkedlist_test.go2:67
 l.Add(1232, 3)
	if l.Len() != 10 {
		t.Fatalf("unexpected Add!")
	}

//line linkedlist_test.go2:73
 l.Clear()
	if !l.Empty() {
		t.Fatalf("unexpected Clear!")
	}
}
//line linkedlist.go2:17
func instantiate୦୦CreateLinkedList୦int() *instantiate୦୦linkedList୦int {
	h, t := &instantiate୦୦linkedListNode୦int{}, &instantiate୦୦linkedListNode୦int{}
	h.Prev, h.Next, t.Prev, t.Next = t, t, h, h
	return &instantiate୦୦linkedList୦int{
		h,
		t,
		0,
	}
}

//line linkedlist.go2:25
type instantiate୦୦linkedList୦int struct {
//line linkedlist.go2:12
 head, tail *instantiate୦୦linkedListNode୦int

				len int
}

//line linkedlist.go2:27
func (l *instantiate୦୦linkedList୦int,) Len() int {
	return l.len
}

func (l *instantiate୦୦linkedList୦int,) Empty() bool {
	return l.len == 0
}

func (l *instantiate୦୦linkedList୦int,) Clear() {
	l.len = 0
	l.head.Next, l.head.Prev, l.tail.Prev, l.tail.Next = l.tail, l.tail, l.head, l.head
}

func (l *instantiate୦୦linkedList୦int,) Traverse() []int {
	ret := make([]int, l.len)
	cur := l.head.Next
	for i := 0; i < l.Len(); i++ {
		ret[i], cur = cur.Val, cur.Next
	}
	return ret
}

//line linkedlist.go2:56
func (l *instantiate୦୦linkedList୦int,) Add(val int,

//line linkedlist.go2:56
 index int) {

	if index > l.len || index < -l.len-1 {
		return
	}

	if index >= 0 {
		l.addFromFirst(val, index)
	} else {
		l.addFromLast(val, -index)
	}

	l.len++
}

func (l *instantiate୦୦linkedList୦int,) addFromFirst(val int,

//line linkedlist.go2:71
 index int) {
	cur := l.head
	for i := 0; i < index; i++ {
		cur = cur.Next
	}
	node := &instantiate୦୦linkedListNode୦int{
		Val: val,
	}
	node.Prev, node.Next, cur.Next, cur.Next.Prev = cur, cur.Next, node, node
}

func (l *instantiate୦୦linkedList୦int,) addFromLast(val int,

//line linkedlist.go2:82
 index int) {
	cur := l.tail
	for i := 0; i < index; i++ {
		cur = cur.Prev
	}
	node := &instantiate୦୦linkedListNode୦int{
		Val: val,
	}
	node.Prev, node.Next, cur.Next, cur.Next.Prev = cur, cur.Next, node, node
}

//line linkedlist.go2:100
func (l *instantiate୦୦linkedList୦int,) Set(val int,

//line linkedlist.go2:100
 index int) {

	if l.Empty() || index >= l.len || index <= -l.len-1 {
		return
	}

	if index >= 0 {
		l.setFromFirst(val, index)
	} else {
		l.setFromLast(val, -index)
	}
}

func (l *instantiate୦୦linkedList୦int,) setFromFirst(val int,

//line linkedlist.go2:113
 index int) {
	cur := l.head
	for i := 0; i < index; i++ {
		cur = cur.Next
	}
	cur.Next.Val = val
}

func (l *instantiate୦୦linkedList୦int,) setFromLast(val int,

//line linkedlist.go2:121
 index int) {
	cur := l.tail
	for i := 0; i < index-1; i++ {
		cur = cur.Prev
	}
	cur.Prev.Val = val
}

//line linkedlist.go2:136
func (l *instantiate୦୦linkedList୦int,) Remove(index int) {

	if l.Empty() || index >= l.len || index <= -l.len-1 {
		return
	}
	if index < 0 {
		l.removeFromLast(-index)
	} else {
		l.removeFromFirst(index)
	}
	l.len--
}

func (l *instantiate୦୦linkedList୦int,) removeFromLast(index int) {
	cur := l.tail
	for i := 0; i < index; i++ {
		cur = cur.Prev
	}
	cur.Prev, cur.Next, cur.Prev.Next, cur.Next.Prev = nil, nil, cur.Next, cur.Prev
}

func (l *instantiate୦୦linkedList୦int,) removeFromFirst(index int) {
	cur := l.head
	for i := 0; i < index+1; i++ {
		cur = cur.Next
	}
	cur.Prev, cur.Next, cur.Prev.Next, cur.Next.Prev = nil, nil, cur.Next, cur.Prev
}

//line linkedlist.go2:166
func (l *instantiate୦୦linkedList୦int,) Append(val int,

//line linkedlist.go2:166
) {
	node := &instantiate୦୦linkedListNode୦int{Val: val}
	l.tail.Prev.Next, l.tail.Prev, node.Prev, node.Next = node, node, l.tail.Prev, l.tail
	l.len++
}

//line linkedlist.go2:174
func (l *instantiate୦୦linkedList୦int,) Front() (e int,

//line linkedlist.go2:174
) {
	if !l.Empty() {
		return l.head.Next.Val
	}
	return
}

//line linkedlist.go2:183
func (l *instantiate୦୦linkedList୦int,) Back() (e int,

//line linkedlist.go2:183
) {
	if !l.Empty() {
		return l.tail.Prev.Val
	}
	return
}

//line linkedlist.go2:188
type instantiate୦୦linkedListNode୦int struct {
//line linkedlist.go2:5
 Val        int
				Prev, Next *instantiate୦୦linkedListNode୦int
}

//line linkedlist.go2:7
var _ = rand.ExpFloat64
//line linkedlist.go2:7
var _ = reflect.Append
//line linkedlist.go2:7
var _ = testing.AllocsPerRun

//line linkedlist.go2:7
const _ = time.ANSIC
