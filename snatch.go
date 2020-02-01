package snatch

import (
	"reflect"
	"unsafe"
)

// Field ...
func Field(
	ob interface{},
	fieldname string,
	replace func(unsafe.Pointer),
) error {
	rv := reflect.ValueOf(ob).Elem()
	rt := rv.Type()

	// TODO: validation

	var rf reflect.Value
	for i := 0; i < rt.NumField(); i++ {
		f := rt.Field(i)
		if f.Name == fieldname {
			rf = rv.FieldByIndex([]int{i})
			break
		}
	}

	ptr := unsafe.Pointer(rf.UnsafeAddr())
	replace(ptr)
	return nil
}

// FieldOrPanic ...
func FieldOrPanic(
	ob interface{},
	fieldname string,
	replace func(unsafe.Pointer),
) {
	err := Field(ob, fieldname, replace)
	if err != nil {
		panic(err)
	}
}
