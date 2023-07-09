package object_test

import (
	"fmt"
	"github.com/louvri/gob/object"
	"reflect"
	"testing"
	"time"
)

func TestAssign(t *testing.T) {
	type test struct {
		Duration time.Duration
	}
	objectTest := &test{}
	durationString := "1h"
	el := reflect.ValueOf(objectTest).Elem()
	for i := 0; i < el.NumField(); i++ {
		prop := el.Type().Field(i).Name
		ref := el.FieldByName(prop)
		err := object.Assign(ref, prop, durationString)
		if err != nil {
			t.Fatal(err)
		}
	}
	fmt.Println(objectTest.Duration)
}
