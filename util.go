package goutils

import (
	"fmt"
	"reflect"
)

func InstanceOf(sample interface{}, given interface{}) bool {
	foo := reflect.TypeOf(sample)
	bar := reflect.TypeOf(given)

	return foo == bar
}

func TypeError(sample interface{}, given interface{}) string {
	sampleTypeOf := reflect.TypeOf(sample)
	givenTypeOf := reflect.TypeOf(given)

	return fmt.Sprintf("invalid type, cannot process %s as %s.", givenTypeOf, sampleTypeOf)
}
