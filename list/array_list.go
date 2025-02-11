// Copyright 2021 ecodeclub
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package list

import (
	"github.com/ecodeclub/ekit/internal/errs"
	"github.com/ecodeclub/ekit/internal/iterator"
	"github.com/ecodeclub/ekit/internal/slice"
)

var (
	_ List[any] = &ArrayList[any]{}
)

// ArrayList 基于切片的简单封装
type ArrayList[T any] struct {
	vals     []T
	modCount int
}

// NewArrayList 初始化一个len为0，cap为cap的ArrayList
func NewArrayList[T any](cap int) *ArrayList[T] {
	return &ArrayList[T]{vals: make([]T, 0, cap)}
}

// NewArrayListOf 直接使用 ts，而不会执行复制
func NewArrayListOf[T any](ts []T) *ArrayList[T] {
	return &ArrayList[T]{
		vals: ts,
	}
}

func (a *ArrayList[T]) Get(index int) (t T, e error) {
	l := a.Len()
	if index < 0 || index >= l {
		return t, errs.NewErrIndexOutOfRange(l, index)
	}
	return a.vals[index], e
}

// Append 往ArrayList里追加数据
func (a *ArrayList[T]) Append(ts ...T) error {
	a.Increment()
	a.vals = append(a.vals, ts...)
	return nil
}

// Add 在ArrayList下标为index的位置插入一个元素
// 当index等于ArrayList长度等同于append
func (a *ArrayList[T]) Add(index int, t T) (err error) {
	a.Increment()
	a.vals, err = slice.Add(a.vals, t, index)
	return
}

// Set 设置ArrayList里index位置的值为t
func (a *ArrayList[T]) Set(index int, t T) error {
	a.Increment()
	length := len(a.vals)
	if index >= length || index < 0 {
		return errs.NewErrIndexOutOfRange(length, index)
	}
	a.vals[index] = t
	return nil
}

// Delete 方法会在必要的时候引起缩容，其缩容规则是：
// - 如果容量 > 2048，并且长度小于容量一半，那么就会缩容为原本的 5/8
// - 如果容量 (64, 2048]，如果长度是容量的 1/4，那么就会缩容为原本的一半
// - 如果此时容量 <= 64，那么我们将不会执行缩容。在容量很小的情况下，浪费的内存很少，所以没必要消耗 CPU去执行缩容
func (a *ArrayList[T]) Delete(index int) (T, error) {
	a.Increment()
	res, t, err := slice.Delete(a.vals, index)
	if err != nil {
		return t, err
	}
	a.vals = res
	a.shrink()
	return t, nil
}

// shrink 数组缩容
func (a *ArrayList[T]) shrink() {
	a.Increment()
	a.vals = slice.Shrink(a.vals)
}

func (a *ArrayList[T]) Len() int {
	return len(a.vals)
}

func (a *ArrayList[T]) Cap() int {
	return cap(a.vals)
}

func (a *ArrayList[T]) Range(fn func(index int, t T) error) error {
	for key, value := range a.vals {
		e := fn(key, value)
		if e != nil {
			return e
		}
	}
	return nil
}

func (a *ArrayList[T]) AsSlice() []T {
	res := make([]T, len(a.vals))
	copy(res, a.vals)
	return res
}

func (p *ArrayList[T]) GetModCount() int {
	return p.modCount
}

// Increment 这里更倾向对于修改为内部方法
func (p *ArrayList[T]) Increment() {
	if p.modCount == 10000000 {
		p.modCount = 0
	}
	p.modCount++
}

func (a *ArrayList[T]) Iterator() iterator.Iterator[T] {
	return &iter[T]{
		cur:         0,
		prev:        -1,
		source:      a,
		curModCount: a.GetModCount(),
	}
}

type iter[T any] struct {
	source      *ArrayList[T]
	err         error
	cur         int
	prev        int
	curModCount int
}

func (i *iter[T]) Get() (T, error) {
	var t T
	err := i.CheckMod()
	if err != nil {
		return t, err
	}
	return i.source.Get(i.cur)
}

func (i *iter[T]) Err() error {
	return i.err
}

func (i *iter[T]) Next() (T, error) {
	var t T
	err := i.CheckMod()
	if err != nil {
		return t, err
	}

	t, err = i.source.Get(i.cur)
	if err != nil {
		return t, err
	}
	i.prev = i.cur
	i.cur++
	return t, nil
}

func (i *iter[T]) HasNext() bool {
	err := i.CheckMod()
	if err != nil {
		i.err = err
		return false
	}
	return i.cur < i.source.Len()
}

func (i *iter[T]) Delete() error {
	err := i.CheckMod()
	if err != nil {
		return err
	}
	if i.prev == -1 {
		return iterator.ErrNoSuchData
	}
	_, err = i.source.Delete(i.cur)
	if err != nil {
		return err
	}
	i.curModCount = i.source.GetModCount()
	i.cur = i.prev
	i.prev = -1
	return nil
}

// CheckMod 由于有的数据结构支持遍历修改，所有没有加到迭代器里
func (i *iter[T]) CheckMod() error {
	if i.curModCount != i.source.GetModCount() {
		return iterator.ErrStructHasChange
	}
	return nil
}
