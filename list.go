package goutils

import (
	"reflect"
)

type EachElementCallback = func(index int, el interface{}) interface{}
type ForEachCallback = func(index int, el interface{})

type List interface {
	Clear()
	Count() int
	ElementAt(index int) interface{}
	Filter(fn EachElementCallback) interface{}
	Find(fn EachElementCallback) interface{}
	FindIndex(fn EachElementCallback) int
	ForEach(fn ForEachCallback)
	Get() interface{}
	Map(sample interface{}, fn EachElementCallback) interface{}
	Pop()
	Push(value ...interface{})
	Shift()
	Slice(first int, last int) interface{}
	Some(fn EachElementCallback) bool
}

type list struct {
	t interface{}
	elements []interface{}
}

// The Clear() method removes the elements stored.
func (l *list) Clear() {
	l.elements = l.elements[:0]
}

// The Count() method returns the length of the elements stored.
func (l *list) Count() int {
	return len(l.elements)
}

// The ElementAt() method returns the element of given index.
func (l *list) ElementAt(index int) interface{} {
	return l.elements[index]
}

// The Filter() method creates a new array with all elements that pass the test implemented by the provided function.
func (l *list) Filter(fn EachElementCallback) interface{} {
	typeOf := reflect.TypeOf(l.t)
	sliceOf := reflect.SliceOf(typeOf)
	var result = reflect.ValueOf(reflect.New(sliceOf).Interface()).Elem()

	l.ForEach(func(index int, el interface{}) {
		res := fn(index, el)
		if nil != res {
			result.Set(reflect.Append(result, reflect.ValueOf(res)))
		}
	})

	return result.Interface()
}

// The Find() method returns the value of the first element in the provided array that satisfies the provided testing function.
func (l *list) Find(fn EachElementCallback) interface{} {
	var result interface{}

	for k, v := range l.elements {
		result = fn(k, v)
		if nil != result {
			break
		}
	}

	return result
}

// The FindIndex() method returns the index of the first element in the provided array that satisfies the provided testing function. Otherwise, its return -1, indicating that no element passed the test.
func (l *list) FindIndex(fn EachElementCallback) int {
	var i = -1

	for k, v := range l.elements {
		res := fn(k, v)
		if nil != res {
			i = k
			break
		}
	}

	return i
}

// The ForEach() method executes a provided function once for each array element.
func (l *list) ForEach(fn ForEachCallback) {
	for k, v := range l.elements {
		fn(k, v)
	}
}

// The Get() method returns all the elements.
func (l *list) Get() interface{} {
	typeOf := reflect.TypeOf(l.t)
	sliceOf := reflect.SliceOf(typeOf)
	var result = reflect.ValueOf(reflect.New(sliceOf).Interface()).Elem()

	l.ForEach(func(index int, el interface{}) {
		result.Set(reflect.Append(result, reflect.ValueOf(el)))
	})

	return result.Interface()
}

// The Map() method creates a new array populated with the results of calling a provided function on every element in the calling array.
func (l *list) Map(sample interface{}, fn EachElementCallback) interface{} {
	typeOf := reflect.TypeOf(sample)

	if typeOf.Kind() != reflect.Ptr {
		panic("sample must be pointer")
	}

	valueOf := reflect.ValueOf(sample)
	valueElem := valueOf.Elem()

	l.ForEach(func(index int, el interface{}) {
		res := fn(index, el)
		if res != nil {
			valueElem.Set(reflect.Append(valueElem, reflect.ValueOf(res)))
		}
	})

	return valueElem.Interface()
}

// The Pop() method removes the last element from an array and returns that element. This method changes the length of the array.
func (l *list) Pop() {
	l.elements = l.elements[:len(l.elements)-1]
}

// The Push() method adds one or more elements to the end of an array and returns the new length of the array.
func (l *list) Push(value ...interface{}) {
	for _, v := range value {
		if !InstanceOf(l.t, v) {
			panic(TypeError(l.t, value))
		}

		l.elements = append(l.elements, v)
	}
}

// The Shift() method removes the first element from an array and returns that removed element. This method changes the length of the array.
func (l *list) Shift() {
	l.elements = l.elements[1:]
}

// The Slice() method returns a shallow copy of a portion of an array into a new array object selected from start to end (end not included) where start and end represent the index of items in that array. The original array will not be modified.
func (l *list) Slice(first int, last int) interface{} {
	typeOf := reflect.TypeOf(l.t)
	sliceOf := reflect.SliceOf(typeOf)
	var result = reflect.ValueOf(reflect.New(sliceOf).Interface()).Elem()

	for _, v := range l.elements[first:last] {
		result.Set(reflect.Append(result, reflect.ValueOf(v)))
	}

	return result.Interface()
}

// The Some() method tests whether at least one element in the array passes the test implemented by the provided function. It returns a Boolean value.
func (l *list) Some(fn EachElementCallback) bool {
	var isIncluded = false

	for k, v := range l.elements {
		res := fn(k, v)
		if nil != res {
			isIncluded = true
			break
		}
	}

	return isIncluded
}

func NewList(t interface{}) List {
	return &list{
		t: t,
		elements: []interface{}{},
	}
}