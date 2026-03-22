package object

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"
)

func Ref(obj any) reflect.Value {
	return reflect.ValueOf(obj).Elem()
}

func Prop(ref reflect.Value, prop string) reflect.Value {
	return ref.FieldByName(prop)
}

func Get(ref reflect.Value) any {
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

func DefaultValue(ref reflect.Value, quoted bool) any {
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

func Assign(ref reflect.Value, name string, value any) error {
	switch ref.Interface().(type) {
	case int, int16, int32, int64:
		val, ok := value.(int64)
		if !ok {
			return fmt.Errorf("config failed to parse int: %s", name)
		}
		ref.SetInt(val)
	case float32, float64:
		val, ok := value.(float64)
		if !ok {
			return fmt.Errorf("config failed to parse float: %s", name)
		}
		ref.SetFloat(val)
	case bool:
		val, ok := value.(bool)
		if !ok {
			return fmt.Errorf("config failed to parse bool: %s", name)
		}
		ref.SetBool(val)
	case time.Duration:
		d, err := parseDuration(value)
		if err != nil {
			return fmt.Errorf("config failed to parse duration: %s: %w", name, err)
		}
		if d != nil {
			ref.Set(reflect.ValueOf(*d))
		}
	case map[string]string:
		// no-op: value is still empty and will be filled later
	default:
		val, ok := value.(string)
		if !ok {
			return fmt.Errorf("config failed to parse: %s", name)
		}
		ref.SetString(val)
	}
	return nil
}

func parseDuration(value any) (*time.Duration, error) {
	switch v := value.(type) {
	case time.Duration:
		return &v, nil
	case string:
		d, err := time.ParseDuration(v)
		if err != nil {
			return nil, err
		}
		return &d, nil
	case int64:
		d := time.Duration(v)
		return &d, nil
	default:
		return nil, nil
	}
}

func EqualOnNonEmpty(data, filter any) bool {
	el1 := reflect.ValueOf(data).Elem()
	el2 := reflect.ValueOf(filter).Elem()
	for i := 0; i < el2.NumField(); i++ {
		ref2 := el2.Field(i)
		if IsEmpty(ref2) {
			continue
		}
		ref1 := el1.FieldByName(el2.Type().Field(i).Name)
		if ref1.Kind() != ref2.Kind() {
			continue
		}
		if Get(ref1) != Get(ref2) {
			return false
		}
	}
	return true
}

func Patch(data1, data2 any) error {
	el1 := reflect.ValueOf(data1).Elem()
	el2 := reflect.ValueOf(data2).Elem()
	for i := 0; i < el2.NumField(); i++ {
		ref2 := el2.Field(i)
		if IsEmpty(ref2) {
			continue
		}
		prop := el2.Type().Field(i).Name
		ref1 := el1.FieldByName(prop)
		if ref1.Kind() != ref2.Kind() {
			continue
		}
		if err := Assign(ref1, prop, Get(ref2)); err != nil {
			return err
		}
	}
	return nil
}

func Flatten(data any, filter []string) []string {
	result := make([]string, 0)
	el := reflect.ValueOf(data).Elem()
	filterSet := make(map[string]bool, len(filter))
	for _, f := range filter {
		filterSet[f] = true
	}
	for i := 0; i < el.NumField(); i++ {
		prop := el.Type().Field(i).Name
		if len(filterSet) > 0 && !filterSet[prop] {
			continue
		}
		ref := el.Field(i)
		if ref.IsValid() {
			result = append(result, fmt.Sprintf("%v", Get(ref)))
		}
	}
	return result
}

func GetStructTags(field reflect.StructField) map[string]string {
	raw := string(field.Tag)
	raw = strings.ReplaceAll(raw, "\"", "")
	if raw == "" {
		return nil
	}

	output := make(map[string]string)
	for _, item := range splitQuoted(raw) {
		parts := strings.SplitN(item, ":", 2)
		if len(parts) < 2 {
			continue
		}
		output[parts[0]] = strings.Split(parts[1], ",")[0]
	}
	return output
}

func splitQuoted(s string) []string {
	var result []string
	var sb strings.Builder
	quoted := false
	for _, r := range s {
		if r == '\'' {
			quoted = !quoted
			sb.WriteRune(r)
		} else if !quoted && r == ' ' {
			result = append(result, sb.String())
			sb.Reset()
		} else {
			sb.WriteRune(r)
		}
	}
	if sb.Len() > 0 {
		result = append(result, sb.String())
	}
	return result
}

func IterateWithDBProp(data any, handler func(key string, value any, tags map[string]string, isdefault bool)) {
	el := reflect.ValueOf(data).Elem()
	tp := reflect.TypeOf(data).Elem()
	for i := 0; i < el.NumField(); i++ {
		ref := el.Field(i)
		field := tp.Field(i)
		if ref.IsValid() {
			tags := GetStructTags(field)
			var value any
			isempty := IsEmpty(ref)
			if !isempty {
				value = Get(ref)
			}
			handler(field.Name, value, tags, isempty)
		}
	}
}

func Iterate(data any, handler func(key string, value any, isempty bool)) {
	el := reflect.ValueOf(data).Elem()
	for i := 0; i < el.NumField(); i++ {
		ref := el.Field(i)
		if ref.IsValid() {
			name := el.Type().Field(i).Name
			isempty := IsEmpty(ref)
			var value any
			if !isempty {
				value = Get(ref)
			}
			handler(name, value, isempty)
		}
	}
}

func Call(ref any, name string, request []any) (any, error) {
	methodRef := reflect.ValueOf(ref).MethodByName(name)
	if !methodRef.IsValid() {
		return nil, errors.New("method not found")
	}
	in := make([]reflect.Value, len(request))
	for i, item := range request {
		in[i] = reflect.ValueOf(item)
	}
	result := methodRef.Call(in)
	if result[1].Interface() == nil {
		return result[0].Interface(), nil
	}
	return result[0].Interface(), result[1].Interface().(error)
}
