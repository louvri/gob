package object_test

import (
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/louvri/gob/object"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// --- test helpers ---

type sample struct {
	Name    string
	Age     int64
	Score   float64
	Active  bool
	Comment string
}

type tagged struct {
	ID   int    `json:"id" db:"id,autoincrement"`
	Name string `json:"name" db:"name"`
}

type callable struct{}

func (c *callable) Add(a, b int) (int, error) {
	return a + b, nil
}

func (c *callable) Fail(msg string) (string, error) {
	return "", fmt.Errorf("%s", msg)
}

// --- Ref ---

func TestRef(t *testing.T) {
	s := &sample{Name: "alice"}
	ref := object.Ref(s)
	assert.Equal(t, "alice", ref.FieldByName("Name").String())
}

// --- Prop ---

func TestProp(t *testing.T) {
	s := &sample{Name: "bob"}
	ref := object.Ref(s)
	prop := object.Prop(ref, "Name")
	assert.Equal(t, "bob", prop.String())
}

// --- Get ---

func TestGet_Int(t *testing.T) {
	s := &sample{Age: 25}
	ref := object.Ref(s)
	val := object.Get(ref.FieldByName("Age"))
	assert.Equal(t, int64(25), val)
}

func TestGet_Float(t *testing.T) {
	s := &sample{Score: 3.14}
	ref := object.Ref(s)
	val := object.Get(ref.FieldByName("Score"))
	assert.Equal(t, 3.14, val)
}

func TestGet_Bool(t *testing.T) {
	s := &sample{Active: true}
	ref := object.Ref(s)
	val := object.Get(ref.FieldByName("Active"))
	assert.Equal(t, true, val)
}

func TestGet_String(t *testing.T) {
	s := &sample{Name: "hello"}
	ref := object.Ref(s)
	val := object.Get(ref.FieldByName("Name"))
	assert.Equal(t, "hello", val)
}

// --- IsEmpty ---

func TestIsEmpty_ZeroInt(t *testing.T) {
	s := &sample{}
	ref := object.Ref(s)
	assert.True(t, object.IsEmpty(ref.FieldByName("Age")))
}

func TestIsEmpty_NonZeroInt(t *testing.T) {
	s := &sample{Age: 1}
	ref := object.Ref(s)
	assert.False(t, object.IsEmpty(ref.FieldByName("Age")))
}

func TestIsEmpty_EmptyString(t *testing.T) {
	s := &sample{}
	ref := object.Ref(s)
	assert.True(t, object.IsEmpty(ref.FieldByName("Name")))
}

func TestIsEmpty_NonEmptyString(t *testing.T) {
	s := &sample{Name: "a"}
	ref := object.Ref(s)
	assert.False(t, object.IsEmpty(ref.FieldByName("Name")))
}

func TestIsEmpty_ZeroFloat(t *testing.T) {
	s := &sample{}
	ref := object.Ref(s)
	assert.True(t, object.IsEmpty(ref.FieldByName("Score")))
}

func TestIsEmpty_NonZeroFloat(t *testing.T) {
	s := &sample{Score: 0.1}
	ref := object.Ref(s)
	assert.False(t, object.IsEmpty(ref.FieldByName("Score")))
}

// --- DefaultValue ---

func TestDefaultValue_Int(t *testing.T) {
	s := &sample{}
	ref := object.Ref(s)
	assert.Equal(t, 0, object.DefaultValue(ref.FieldByName("Age"), false))
}

func TestDefaultValue_Float(t *testing.T) {
	s := &sample{}
	ref := object.Ref(s)
	assert.Equal(t, 0.0, object.DefaultValue(ref.FieldByName("Score"), false))
}

func TestDefaultValue_Bool(t *testing.T) {
	s := &sample{}
	ref := object.Ref(s)
	assert.Equal(t, false, object.DefaultValue(ref.FieldByName("Active"), false))
}

func TestDefaultValue_StringNotQuoted(t *testing.T) {
	s := &sample{}
	ref := object.Ref(s)
	assert.Equal(t, "", object.DefaultValue(ref.FieldByName("Name"), false))
}

func TestDefaultValue_StringQuoted(t *testing.T) {
	s := &sample{}
	ref := object.Ref(s)
	assert.Equal(t, "''", object.DefaultValue(ref.FieldByName("Name"), true))
}

// --- Assign ---

func TestAssign_DurationFromString(t *testing.T) {
	type dur struct {
		Duration time.Duration
	}
	obj := &dur{}
	el := reflect.ValueOf(obj).Elem()
	ref := el.FieldByName("Duration")
	err := object.Assign(ref, "Duration", "1h")
	require.NoError(t, err)
	assert.Equal(t, time.Hour, obj.Duration)
}

func TestAssign_DurationFromInt64(t *testing.T) {
	type dur struct {
		Duration time.Duration
	}
	obj := &dur{}
	el := reflect.ValueOf(obj).Elem()
	ref := el.FieldByName("Duration")
	err := object.Assign(ref, "Duration", int64(5000000000))
	require.NoError(t, err)
	assert.Equal(t, 5*time.Second, obj.Duration)
}

func TestAssign_DurationFromDuration(t *testing.T) {
	type dur struct {
		Duration time.Duration
	}
	obj := &dur{}
	el := reflect.ValueOf(obj).Elem()
	ref := el.FieldByName("Duration")
	err := object.Assign(ref, "Duration", 30*time.Minute)
	require.NoError(t, err)
	assert.Equal(t, 30*time.Minute, obj.Duration)
}

func TestAssign_Int(t *testing.T) {
	s := &sample{}
	el := reflect.ValueOf(s).Elem()
	ref := el.FieldByName("Age")
	err := object.Assign(ref, "Age", int64(42))
	require.NoError(t, err)
	assert.Equal(t, int64(42), s.Age)
}

func TestAssign_IntWrongType(t *testing.T) {
	s := &sample{}
	el := reflect.ValueOf(s).Elem()
	ref := el.FieldByName("Age")
	err := object.Assign(ref, "Age", "not a number")
	assert.Error(t, err)
}

func TestAssign_Float(t *testing.T) {
	s := &sample{}
	el := reflect.ValueOf(s).Elem()
	ref := el.FieldByName("Score")
	err := object.Assign(ref, "Score", float64(9.5))
	require.NoError(t, err)
	assert.Equal(t, 9.5, s.Score)
}

func TestAssign_FloatWrongType(t *testing.T) {
	s := &sample{}
	el := reflect.ValueOf(s).Elem()
	ref := el.FieldByName("Score")
	err := object.Assign(ref, "Score", "wrong")
	assert.Error(t, err)
}

func TestAssign_Bool(t *testing.T) {
	s := &sample{}
	el := reflect.ValueOf(s).Elem()
	ref := el.FieldByName("Active")
	err := object.Assign(ref, "Active", true)
	require.NoError(t, err)
	assert.True(t, s.Active)
}

func TestAssign_BoolWrongType(t *testing.T) {
	s := &sample{}
	el := reflect.ValueOf(s).Elem()
	ref := el.FieldByName("Active")
	err := object.Assign(ref, "Active", "wrong")
	assert.Error(t, err)
}

func TestAssign_String(t *testing.T) {
	s := &sample{}
	el := reflect.ValueOf(s).Elem()
	ref := el.FieldByName("Name")
	err := object.Assign(ref, "Name", "alice")
	require.NoError(t, err)
	assert.Equal(t, "alice", s.Name)
}

func TestAssign_StringWrongType(t *testing.T) {
	s := &sample{}
	el := reflect.ValueOf(s).Elem()
	ref := el.FieldByName("Name")
	err := object.Assign(ref, "Name", 123)
	assert.Error(t, err)
}

func TestAssign_DurationInvalidString(t *testing.T) {
	type dur struct {
		Duration time.Duration
	}
	obj := &dur{}
	el := reflect.ValueOf(obj).Elem()
	ref := el.FieldByName("Duration")
	err := object.Assign(ref, "Duration", "not-a-duration")
	assert.Error(t, err)
}

// --- EqualOnNonEmpty ---

func TestEqualOnNonEmpty_Match(t *testing.T) {
	data := &sample{Name: "alice", Age: 25, Score: 3.14}
	filter := &sample{Name: "alice"}
	assert.True(t, object.EqualOnNonEmpty(data, filter))
}

func TestEqualOnNonEmpty_Mismatch(t *testing.T) {
	data := &sample{Name: "alice", Age: 25}
	filter := &sample{Name: "bob"}
	assert.False(t, object.EqualOnNonEmpty(data, filter))
}

func TestEqualOnNonEmpty_EmptyFilter(t *testing.T) {
	data := &sample{Name: "alice", Age: 25}
	filter := &sample{}
	assert.True(t, object.EqualOnNonEmpty(data, filter))
}

func TestEqualOnNonEmpty_MultipleFieldsMatch(t *testing.T) {
	data := &sample{Name: "alice", Age: 25}
	filter := &sample{Name: "alice", Age: 25}
	assert.True(t, object.EqualOnNonEmpty(data, filter))
}

func TestEqualOnNonEmpty_MultipleFieldsMismatch(t *testing.T) {
	data := &sample{Name: "alice", Age: 25}
	filter := &sample{Name: "alice", Age: 30}
	assert.False(t, object.EqualOnNonEmpty(data, filter))
}

// --- Patch ---

func TestPatch_AppliesNonEmptyFields(t *testing.T) {
	data := &sample{Name: "alice", Age: 25, Score: 1.0}
	patch := &sample{Age: 30, Score: 9.5}
	err := object.Patch(data, patch)
	require.NoError(t, err)
	assert.Equal(t, "alice", data.Name)
	assert.Equal(t, int64(30), data.Age)
	assert.Equal(t, 9.5, data.Score)
}

func TestPatch_EmptyPatchChangesNothing(t *testing.T) {
	data := &sample{Name: "alice", Age: 25}
	patch := &sample{}
	err := object.Patch(data, patch)
	require.NoError(t, err)
	assert.Equal(t, "alice", data.Name)
	assert.Equal(t, int64(25), data.Age)
}

// --- Flatten ---

func TestFlatten_AllFields(t *testing.T) {
	s := &sample{Name: "alice", Age: 25, Score: 3.14, Active: true, Comment: "hi"}
	result := object.Flatten(s, nil)
	assert.Equal(t, []string{"alice", "25", "3.14", "true", "hi"}, result)
}

func TestFlatten_WithFilter(t *testing.T) {
	s := &sample{Name: "alice", Age: 25, Score: 3.14}
	result := object.Flatten(s, []string{"Name", "Score"})
	assert.Equal(t, []string{"alice", "3.14"}, result)
}

func TestFlatten_EmptyFilter(t *testing.T) {
	s := &sample{Name: "alice"}
	result := object.Flatten(s, []string{})
	assert.Equal(t, 5, len(result))
}

// --- GetStructTags ---

func TestGetStructTags(t *testing.T) {
	typ := reflect.TypeOf(tagged{})
	field, _ := typ.FieldByName("ID")
	tags := object.GetStructTags(field)
	assert.Equal(t, "id", tags["json"])
	assert.Equal(t, "id", tags["db"])
}

func TestGetStructTags_EmptyTag(t *testing.T) {
	type noTags struct {
		Field string
	}
	typ := reflect.TypeOf(noTags{})
	field, _ := typ.FieldByName("Field")
	tags := object.GetStructTags(field)
	assert.Nil(t, tags)
}

func TestGetStructTags_MalformedTag(t *testing.T) {
	field := reflect.StructField{
		Name: "Field",
		Type: reflect.TypeOf(""),
		Tag:  reflect.StructTag("nocolon"),
	}
	tags := object.GetStructTags(field)
	assert.Empty(t, tags)
}

// --- Iterate ---

func TestIterate(t *testing.T) {
	s := &sample{Name: "alice", Age: 25}
	var keys []string
	var empties []bool
	object.Iterate(s, func(key string, value any, isempty bool) {
		keys = append(keys, key)
		empties = append(empties, isempty)
	})
	assert.Contains(t, keys, "Name")
	assert.Contains(t, keys, "Age")
	assert.Equal(t, 5, len(keys))
}

// --- IterateWithDBProp ---

func TestIterateWithDBProp(t *testing.T) {
	s := &tagged{ID: 1, Name: "alice"}
	var keys []string
	var tagsList []map[string]string
	object.IterateWithDBProp(s, func(key string, value any, tags map[string]string, isdefault bool) {
		keys = append(keys, key)
		tagsList = append(tagsList, tags)
	})
	assert.Equal(t, []string{"ID", "Name"}, keys)
	assert.Equal(t, "id", tagsList[0]["json"])
	assert.Equal(t, "name", tagsList[1]["json"])
}

// --- Call ---

func TestCall_ValidMethod(t *testing.T) {
	c := &callable{}
	result, err := object.Call(c, "Add", []any{3, 4})
	require.NoError(t, err)
	assert.Equal(t, 7, result)
}

func TestCall_InvalidMethod(t *testing.T) {
	c := &callable{}
	_, err := object.Call(c, "NonExistent", nil)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "method not found")
}

func TestCall_MethodReturnsError(t *testing.T) {
	c := &callable{}
	result, err := object.Call(c, "Fail", []any{"something went wrong"})
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "something went wrong")
	assert.Equal(t, "", result)
}

// --- SliceContains / SliceNotContains ---

func TestSliceContains_Found(t *testing.T) {
	assert.True(t, object.SliceContains([]string{"a", "b", "c"}, "b"))
}

func TestSliceContains_NotFound(t *testing.T) {
	assert.False(t, object.SliceContains([]string{"a", "b"}, "z"))
}

func TestSliceContains_Empty(t *testing.T) {
	assert.False(t, object.SliceContains([]string{}, "a"))
}

func TestSliceNotContains_Found(t *testing.T) {
	assert.False(t, object.SliceNotContains([]string{"a", "b"}, "b"))
}

func TestSliceNotContains_NotFound(t *testing.T) {
	assert.True(t, object.SliceNotContains([]string{"a", "b"}, "z"))
}

// --- GenerateRandomText ---

func TestGenerateRandomText_Length(t *testing.T) {
	for _, n := range []int{0, 1, 10, 100} {
		text := object.GenerateRandomText(n)
		assert.Equal(t, n, len(text))
	}
}

func TestGenerateRandomText_OnlyLetters(t *testing.T) {
	text := object.GenerateRandomText(1000)
	for _, c := range text {
		assert.True(t, (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z'), "unexpected char: %c", c)
	}
}

func TestGenerateRandomText_Unique(t *testing.T) {
	text1 := object.GenerateRandomText(16)
	text2 := object.GenerateRandomText(16)
	assert.NotEqual(t, text1, text2)
}

// --- GenerateRunningNumbers ---

func TestGenerateRunningNumbers_NonZero(t *testing.T) {
	n := object.GenerateRunningNumbers()
	assert.NotZero(t, n)
}

func TestGenerateRunningNumbers_Unique(t *testing.T) {
	seen := make(map[uint32]bool)
	for i := 0; i < 100; i++ {
		n := object.GenerateRunningNumbers()
		assert.False(t, seen[n], "duplicate running number: %d", n)
		seen[n] = true
	}
}

// --- ConvertDateTimeTextWithTimezone ---

func TestConvertDateTimeTextWithTimezone_Valid(t *testing.T) {
	loc, _ := time.LoadLocation("Asia/Jakarta")
	result, err := object.ConvertDateTimeTextWithTimezone(
		"2024-01-15 10:00:00",
		"2006-01-02 15:04:05",
		"02/01/2006 15:04",
		loc,
	)
	require.NoError(t, err)
	assert.Equal(t, "15/01/2024 17:00", result)
}

func TestConvertDateTimeTextWithTimezone_Empty(t *testing.T) {
	loc, _ := time.LoadLocation("UTC")
	result, err := object.ConvertDateTimeTextWithTimezone("", "2006-01-02", "02/01/2006", loc)
	require.NoError(t, err)
	assert.Equal(t, "", result)
}

func TestConvertDateTimeTextWithTimezone_InvalidFormat(t *testing.T) {
	loc, _ := time.LoadLocation("UTC")
	_, err := object.ConvertDateTimeTextWithTimezone("not-a-date", "2006-01-02", "02/01/2006", loc)
	assert.Error(t, err)
}

// --- GetTimeVariableValue (deprecated) ---

func TestGetTimeVariableValue_DelegatesToConvert(t *testing.T) {
	loc, _ := time.LoadLocation("UTC")
	result, err := object.GetTimeVariableValue("2024-01-15", "2006-01-02", "01/02/2006", loc)
	require.NoError(t, err)
	assert.Equal(t, "01/15/2024", result)
}

// --- ConvertDateWithTimezone ---

func TestConvertDateWithTimezone_Valid(t *testing.T) {
	loc, _ := time.LoadLocation("America/New_York")
	result, err := object.ConvertDateWithTimezone("2024-01-15", "2006-01-02", loc)
	require.NoError(t, err)
	assert.NotEmpty(t, result)
}

func TestConvertDateWithTimezone_Empty(t *testing.T) {
	loc, _ := time.LoadLocation("UTC")
	result, err := object.ConvertDateWithTimezone("", "2006-01-02", loc)
	require.NoError(t, err)
	assert.Equal(t, "", result)
}

func TestConvertDateWithTimezone_InvalidFormat(t *testing.T) {
	loc, _ := time.LoadLocation("UTC")
	_, err := object.ConvertDateWithTimezone("bad-date", "2006-01-02", loc)
	assert.Error(t, err)
}
