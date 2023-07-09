package object

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"
)

func Ref(obj interface{}) reflect.Value {
	return reflect.ValueOf(obj).Elem()
}

func Prop(ref reflect.Value, prop string) reflect.Value {
	return ref.FieldByName(prop)
}

func Get(ref reflect.Value) interface{} {
	switch ref.Kind() {
	case reflect.Int, reflect.Int16, reflect.Int8, reflect.Int32, reflect.Int64:
		return ref.Int()
	case reflect.Float32, reflect.Float64:
		return ref.Float()
	case reflect.Bool:
		return ref.Bool()
	}
	return ref.String()
}

func IsEmpty(ref reflect.Value) bool {
	switch ref.Kind() {
	case reflect.Int, reflect.Int16, reflect.Int8, reflect.Int32, reflect.Int64:
		return ref.Int() == 0
	case reflect.Float32, reflect.Float64:
		return ref.Float() == 0
	case reflect.String:
		return ref.String() == ""
	}
	return false
}

func DefaultValue(ref reflect.Value, quoted bool) interface{} {
	switch ref.Kind() {
	case reflect.Int, reflect.Int16, reflect.Int8, reflect.Int32, reflect.Int64:
		return 0
	case reflect.Float32, reflect.Float64:
		return 0.0
	case reflect.Bool:
		return false
	}
	if !quoted {
		return ""
	}
	return "''"
}

func Assign(ref reflect.Value, name string, value interface{}) error {
	switch ref.Interface().(type) {
	case int, int16, int32, int64:
		val, ok := value.(int64)
		if !ok {
			return fmt.Errorf("config failed to parse int : %s", name)
		}
		ref.SetInt(val)
	case float32, float64:
		val, ok := value.(float64)
		if !ok {
			return fmt.Errorf("config failed to parse float : %s", name)
		}
		ref.SetFloat(val)
	case bool:
		val, ok := value.(bool)
		if !ok {
			return fmt.Errorf("config failed to parse bool : %s", name)
		}
		ref.SetBool(val)
	case time.Duration:
		var err error
		var duration time.Duration
		var isSet bool
		duration, ok := value.(time.Duration)
		if !ok {
			if tmp, ok := value.(string); ok {
				duration, err = time.ParseDuration(tmp)
				if err != nil {
					return err
				}
				isSet = true
			} else if tmp, ok := value.(int64); ok {
				duration = time.Duration(tmp)
				isSet = true
			}
		} else {
			isSet = true
		}
		if isSet {
			ref.Set(reflect.ValueOf(duration))
		}
	default:
		val, ok := value.(string)
		if !ok {
			return fmt.Errorf("config failed to parse : %s", name)
		}
		ref.SetString(val)
	}
	return nil
}
func EqualOnNonEmpty(data, filter interface{}) bool {
	el1 := reflect.ValueOf(data).Elem()
	el2 := reflect.ValueOf(filter).Elem()
	ok := true
	for i := 0; i < el2.NumField(); i++ {
		prop := el2.Type().Field(i).Name
		ref1 := el1.FieldByName(prop)
		ref2 := el2.FieldByName(prop)
		if IsEmpty(ref2) {
			continue
		}
		if ref1.Kind() != ref2.Kind() {
			continue
		}
		ok = ok && Get(ref1) == Get(ref2)
	}
	return ok
}

func Patch(data1, data2 interface{}) error {
	el1 := reflect.ValueOf(data1).Elem()
	el2 := reflect.ValueOf(data2).Elem()
	for i := 0; i < el2.NumField(); i++ {
		prop := el2.Type().Field(i).Name
		ref1 := el1.FieldByName(prop)
		ref2 := el2.FieldByName(prop)
		if IsEmpty(ref2) {
			continue
		}
		if ref1.Kind() != ref2.Kind() {
			continue
		}
		if err := Assign(ref1, prop, Get(ref2)); err != nil {
			return err
		}
	}
	return nil
}

func Flatten(data interface{}, filter []string) []string {
	result := make([]string, 0)
	el := reflect.ValueOf(data).Elem()
	for i := 0; i < el.NumField(); i++ {
		prop := el.Type().Field(i).Name
		ref := el.FieldByName(prop)
		if ref.IsValid() {
			result = append(result, fmt.Sprintf("%v", Get(ref)))
		}
	}
	return result
}

func GetStructTags(field reflect.StructField) map[string]string {
	split := func(s string) []string {
		a := []string{}
		sb := &strings.Builder{}
		quoted := false
		for _, r := range s {
			if r == '\'' {
				quoted = !quoted
				sb.WriteRune(r) // keep '"' otherwise comment this line
			} else if !quoted && r == ' ' {
				a = append(a, sb.String())
				sb.Reset()
			} else {
				sb.WriteRune(r)
			}
		}
		if sb.Len() > 0 {
			a = append(a, sb.String())
		}
		return a
	}

	output := make(map[string]string)
	tag := string(field.Tag)
	tag = strings.ReplaceAll(tag, "\"", "")
	if tag != "" {
		tags := split(tag)
		for _, item := range tags {
			tag := strings.SplitN(item, ":", 2)
			if tmp := strings.Split(tag[1], ","); len(tmp) > 0 {
				output[tag[0]] = tmp[0]
			} else {
				output[tag[0]] = tmp[1]
			}

		}
		return output
	}
	return nil
}

func IterateWithDBProp(data interface{}, handler func(key string, value interface{}, tags map[string]string, isdefault bool)) {
	el := reflect.ValueOf(data).Elem()
	to := reflect.TypeOf(data).Elem()
	for i := 0; i < el.NumField(); i++ {
		prop := el.Type().Field(i).Name
		ref := el.FieldByName(prop)
		fields, ok := to.FieldByName(prop)
		if ref.IsValid() && ok {
			tags := GetStructTags(fields)
			var value interface{}
			isempty := IsEmpty(ref)
			if !isempty {
				value = Get(ref)
			}
			handler(prop, value, tags, isempty)
		}
	}
}

func Iterate(data interface{}, handler func(key string, value interface{}, isempty bool)) {
	el := reflect.ValueOf(data).Elem()
	for i := 0; i < el.NumField(); i++ {
		prop := el.Type().Field(i).Name
		ref := el.FieldByName(prop)
		if ref.IsValid() {
			isempty := IsEmpty(ref)
			var value interface{}
			if !isempty {
				value = Get(ref)
			}
			handler(prop, value, isempty)
		}
	}
}

func Call(ref interface{}, name string, request []interface{}) (interface{}, error) {
	methodRef := reflect.ValueOf(ref).MethodByName(name)
	if !methodRef.IsValid() {
		return nil, errors.New("function not exist")
	}
	in := make([]reflect.Value, 0)
	for _, item := range request {
		in = append(in, reflect.ValueOf(item))
	}
	result := methodRef.Call(in)
	if result[1].Interface() == nil {
		return result[0].Interface(), nil
	}
	return result[0].Interface(), result[1].Interface().(error)
}
