package goutils

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

type FooObject struct {
	val int
}
type BarObject struct {
	res int
}

var myData = NewList(FooObject{})

func TestList(t *testing.T) {
	t.Run("List", func(t *testing.T) {
		t.Run("Clear()", TestList_Clear)
		t.Run("Count()", TestList_Count)
		t.Run("ElementAt()", TestList_ElementAt)
		t.Run("Filter()", TestList_Filter)
		t.Run("Find()", TestList_Find)
		t.Run("FindIndex()", TestList_FindIndex)
		t.Run("Get()", TestList_Get)
		t.Run("Map()", TestList_Map)
		t.Run("Pop()", TestList_Pop)
		t.Run("Push()", TestList_Push)
		t.Run("Shift()", TestList_Shift)
		t.Run("Slice()", TestList_Slice)
		t.Run("Some()", TestList_Some)
	})
}

func initSampleData() {
	myData.Push(
		FooObject{val: 10},
		FooObject{val: 20},
		FooObject{val: 30},
		FooObject{val: 40},
		FooObject{val: 50},
	)
}

func TestList_Clear(t *testing.T) {
	initSampleData()
	myData.Clear()
	assert.Empty(t, myData.Get().([]FooObject))
}

func TestList_Count(t *testing.T) {
	initSampleData()
	assert.Equal(t, myData.Count(), 5)
}

func TestList_ElementAt(t *testing.T) {
	initSampleData()
	assert.EqualValues(t, myData.ElementAt(0), FooObject{
		val: 10,
	})
}

func TestList_Filter(t *testing.T) {
	initSampleData()

	result := myData.Filter(func(index int, el interface{}) interface{} {
		if el.(FooObject).val >= 20 && el.(FooObject).val <= 30 {
			return el
		}
		return nil
	})

	for _, v := range result.([]FooObject) {
		if v.val < 20 || v.val > 30 {
			assert.True(t, false, fmt.Sprintf("%v : filtered value not match.", result.([]FooObject)))
			break
		}
	}

	assert.True(t, true)
}

func TestList_Find(t *testing.T) {
	initSampleData()

	result := myData.Find(func(index int, el interface{}) interface{} {
		if el.(FooObject).val == 40 {
			return el
		}
		return nil
	})

	assert.EqualValues(t, result, FooObject{
		val: 40,
	})
}

func TestList_FindIndex(t *testing.T) {
	initSampleData()
	i1 := myData.FindIndex(func(index int, el interface{}) interface{} {
		if el.(FooObject).val == 30 {
			return el
		}
		return nil
	})
	i2 := myData.FindIndex(func(index int, el interface{}) interface{} {
		if el.(FooObject).val == 99 {
			return el
		}
		return nil
	})
	assert.Equal(t, i1, 2)
	assert.Equal(t, i2, -1)
}

func TestList_Get(t *testing.T) {
	initSampleData()
	assert.NotPanics(t, func() {
		_ = myData.Get().([]FooObject)
	})
}

func TestList_Map(t *testing.T) {
	initSampleData()

	assert.Panics(t, func() {
		_ = myData.Map([]BarObject{}, func(index int, el interface{}) interface{} {
			if el.(FooObject).val > 30 {
				return BarObject{
					res: el.(FooObject).val,
				}
			}
			return nil
		})
	})
	assert.NotPanics(t, func() {
		_ = myData.Map(&[]BarObject{}, func(index int, el interface{}) interface{} {
			if el.(FooObject).val > 30 {
				return BarObject{
					res: el.(FooObject).val,
				}
			}
			return nil
		})
	})

	result := myData.Map(&[]BarObject{}, func(index int, el interface{}) interface{} {
		if el.(FooObject).val > 30 {
			return BarObject{res: el.(FooObject).val}
		}
		return nil
	})
	assert.EqualValues(t, result, []BarObject{
		{res: 40},
		{res: 50},
	})
}

func TestList_Pop(t *testing.T) {
	initSampleData()
	totalBeforePop := myData.Count()
	myData.Pop()
	totalAfterPop := myData.Count()
	assert.NotEqual(t, totalAfterPop, totalBeforePop)
	assert.Equal(t, totalAfterPop, totalBeforePop - 1)
}

func TestList_Push(t *testing.T) {
	assert.NotPanics(t, func() {
		myData.Push(FooObject{
			val: 99,
		})
	})

	assert.Panics(t, func() {
		myData.Push(&BarObject{
			res: 99,
		})
	})
}

func TestList_Shift(t *testing.T) {
	initSampleData()
	totalBeforePop := myData.Count()
	myData.Shift()
	totalAfterPop := myData.Count()
	assert.NotEqual(t, totalAfterPop, totalBeforePop)
	assert.Equal(t, totalAfterPop, totalBeforePop - 1)
}

func TestList_Slice(t *testing.T) {
	initSampleData()
	result := myData.Slice(2, 4)
	assert.EqualValues(t, result, []FooObject{
		{val: 30},
		{val: 40},
	})
}

func TestList_Some(t *testing.T) {
	initSampleData()
	isExist := myData.Some(func(index int, el interface{}) interface{} {
		if el.(FooObject).val == 30 {
			return el
		}
		return nil
	})
	assert.True(t, isExist, true)
}