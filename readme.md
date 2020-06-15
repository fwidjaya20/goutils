# Go-Utils
The utilities for golang.

**Table of Content**
1. [List](#list)
    - [Create a List](#create-a-list)
    - [Clear](#clear)
    - [Count](#count)
    - [ElementAt](#elementat)
    - [Filter](#filter)
    - [Find](#find)
    - [FindIndex](#findindex)
    - [ForEach](#foreach)
    - [Get](#get)
    - [Map](#map)
    - [Pop](#pop)
    - [Push](#push)
    - [Shift](#shift)
    - [Slice](#slice)
    - [Some](#some)

## List
List are like objects whose prototype has methods to perform a traversal and mutation operations.

**Common Operation**
##### Create a List
> The `NewList()` method create a new `List` object.
```gotemplate
type Person struct {
    Name    string
    Age     string
    Gender  string
}
foo := NewList(Person{})
```

##### Clear
> The `Clear()` method removes the elements stored.
```gotemplate
foo.Clear()
```

##### Count
> The `Count()` method returns the length of the elements stored.
```gotemplate
foo.Count()
```

##### ElementAt
> The `ElementAt()` method returns the element of given index.
```gotemplate
foo.ElementAt(0)
```

##### Filter
> The `Filter()` method creates a new array with all elements that pass the test implemented by the provided function.
```gotemplate
foo := NewList(Person{})
result := foo.Filter(func(i int, el interface{}) interface{} {
    if el.(Person).Age > 20 {
        return el
    }
    return nil
})
```

##### Find
> The `Find()` method returns the value of the first element in the provided array that satisfies the provided testing function.
```gotemplate
result := foo.Find(func(i int, el interface{}) interface{} {
    if el.(Person).Name == "John" {
        return el
    }
    return nil
})
```

##### FindIndex
> The `FindIndex()` method returns the index of the first element in the provided array that satisfies the provided testing function. Otherwise, its return -1, indicating that no element passed the test.
```gotemplate
result := foo.FindIndex(func(i int, el interface{}) interface{} {
    if el.(Person).Name == "John" {
        return el
    }
    return nil
})
```

##### ForEach
> The `ForEach()` method executes a provided function once for each array element.
```gotemplate
foo.ForEach(func(index int, el interface{}) {
    fmt.Println(index, el)
})
```

##### Get
> The `Get()` method returns all the elements.
```gotemplate
foo.ForEach(func(index int, el interface{}) {
    fmt.Println(index, el)
})
```

##### Map
> The `Map()` method creates a new array populated with the results of calling a provided function on every element in the calling array.
```gotemplate
type PersonNameOnly struct {
    Name string
}

result := foo.Map(&[]PersonNameOnly{}, func(index int, el interface{}) interface{} {
    if el.(Person).Age > 16 {
        return PersonNameOnly{
            name: el.(Person).Name,
        }
    }
    return nil
})
```

##### Pop
> The `Pop()` method removes the last element from an array and returns that element. This method changes the length of the array.
```gotemplate
foo.Pop()
```

##### Push
> The `Push()` method adds one or more elements to the end of an array and returns the new length of the array.
```gotemplate
foo.Push(Person{}, Person{}, Person{}, ...)
```

##### Shift
> The `Shift()` method removes the first element from an array and returns that removed element. This method changes the length of the array.
```gotemplate
foo.Shift()
```

##### Slice
> The `Slice()` method returns a shallow copy of a portion of an array into a new array object selected from start to end (end not included) where start and end represent the index of items in that array. The original array will not be modified.
```gotemplate
result := foo.Slice(1:3)
```

##### Some
> The `Some()` method tests whether at least one element in the array passes the test implemented by the provided function. It returns a Boolean value.
```gotemplate
result := foo.Some(func(index int, el interface{}) interface{} {
    if el.(Person).Age == 35 {
        return el
    }
    return nil
})
```