package sbparam

import (
	"gopkg.in/sensorbee/sensorbee.v0/data"
	"math"
	"reflect"
	"testing"
)

type testStrValue struct {
	StrField1 string `sbparam:"str_field1"` // = required
	StrField2 string `sbparam:"str_field2,,a"`
	StrField3 string `sbparam:",omitempty"`
	StrField4 string `sbparam:""` // = required
	StrField5 string // = required
}

func TestUnmarshalString(t *testing.T) {
	d1 := data.Map{
		"str_field1": data.String("あ"),
		"StrField4":  data.String("b"),
		"StrField5":  data.String("c"),
	}

	actual := &testStrValue{}

	if err := Unmarshal(d1, actual); err != nil {
		t.Fatalf("failed to unmarshal: %v\n", err)
	}
	expected := &testStrValue{
		StrField1: "あ",
		StrField2: "a",
		StrField4: "b",
		StrField5: "c",
	}
	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("faild to unmarshal\n%v\nis expected to equals\n%v\n",
			actual, expected)
	}
}

func TestUnmarshalStringError(t *testing.T) {
	type testSet struct {
		title  string
		in     data.Map
		format interface{}
	}
	testCases := []testSet{
		testSet{
			title:  "lack required field set by no tag",
			in:     data.Map{},
			format: &struct{ StrField string }{},
		},
		testSet{
			title: "lack required field set by empty tag",
			in:    data.Map{},
			format: &struct {
				StrField string `sbparam:""`
			}{},
		},
		testSet{
			title: "lack required field set by sbparam tag",
			in:    data.Map{},
			format: &struct {
				StrField string `sbparam:"str_field"`
			}{},
		},
		testSet{
			title: "value type mismatch",
			in: data.Map{
				"str_field1": data.Int(1),
			},
			format: &struct {
				StrField string
			}{},
		},
	}

	for _, c := range testCases {
		actual := c.format
		if err := Unmarshal(c.in, actual); err == nil {
			t.Fatalf("test case '%v' should occur an error but nothing", c.title)
		}
	}
}

type testIntValue struct {
	IntField   int
	IntField2  int `sbparam:",,-1"`
	Int8Field  int8
	Int16Field int16
	Int32Field int32
	Int64Field int64
}

func TestUnmarshalInt(t *testing.T) {
	d1 := data.Map{
		"IntField":   data.Int(0),
		"Int8Field":  data.Int(math.MaxInt8),
		"Int16Field": data.Int(math.MinInt16),
		"Int32Field": data.Int(math.MaxInt32),
		"Int64Field": data.Int(math.MinInt64),
	}

	actual := &testIntValue{}

	if err := Unmarshal(d1, actual); err != nil {
		t.Fatalf("failed to unmarshal: %v\n", err)
	}
	expected := &testIntValue{
		IntField:   0,
		IntField2:  -1,
		Int8Field:  math.MaxInt8,
		Int16Field: math.MinInt16,
		Int32Field: math.MaxInt32,
		Int64Field: math.MinInt64,
	}
	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("faild to unmarshal\n%v\nis expected to equals\n%v\n",
			actual, expected)
	}
}

func TestUnmarshalIntError(t *testing.T) {
	type testSet struct {
		title  string
		in     data.Map
		format interface{}
	}
	testCases := []testSet{
		testSet{
			title: "data type mismatch",
			in: data.Map{
				"IntField": data.String("error"),
			},
			format: &struct{ IntField int }{},
		},
		testSet{
			title: "default value type mismatch",
			in:    data.Map{},
			format: &struct {
				IntField int `sbparam:",,a"`
			}{},
		},
		testSet{
			title: "int8 value overflow",
			in: data.Map{
				"IntField": data.Int(math.MaxInt8 + 1),
			},
			format: &struct{ IntField int8 }{},
		},
		testSet{
			title: "int16 value overflow",
			in: data.Map{
				"IntField": data.Int(math.MinInt16 - 1),
			},
			format: &struct{ IntField int16 }{},
		},
		testSet{
			title: "int32 value overflow",
			in: data.Map{
				"IntField": data.Int(math.MaxInt32 + 1),
			},
			format: &struct{ IntField int32 }{},
		},
	}

	for _, c := range testCases {
		actual := c.format
		if err := Unmarshal(c.in, actual); err == nil {
			t.Fatalf("test case '%v' should occur an error but nothing", c.title)
		}
	}
}

type testUintValue struct {
	UintField   uint
	UintField2  uint `sbparam:",,1"`
	Uint8Field  uint8
	Uint16Field uint16
	Uint32Field uint32
	Uint64Field uint64
}

func TestUnmarshalUint(t *testing.T) {
	d1 := data.Map{
		"UintField":   data.Int(0),
		"Uint8Field":  data.Int(math.MaxUint8),
		"Uint16Field": data.Int(math.MaxUint16),
		"Uint32Field": data.Int(math.MaxUint32),
		"Uint64Field": data.Int(math.MaxUint32 + 1),
	}

	actual := &testUintValue{}

	if err := Unmarshal(d1, actual); err != nil {
		t.Fatalf("failed to unmarshal: %v\n", err)
	}
	expected := &testUintValue{
		UintField:   0,
		UintField2:  1,
		Uint8Field:  math.MaxUint8,
		Uint16Field: math.MaxUint16,
		Uint32Field: math.MaxUint32,
		Uint64Field: math.MaxUint32 + 1,
	}
	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("faild to unmarshal\n%v\nis expected to equals\n%v\n",
			actual, expected)
	}
}

func TestUnmarshalUintError(t *testing.T) {
	type testSet struct {
		title  string
		in     data.Map
		format interface{}
	}
	testCases := []testSet{
		testSet{
			title: "data type mismatch",
			in: data.Map{
				"UintField": data.String("error"),
			},
			format: &struct{ UintField uint }{},
		},
		testSet{
			title: "default value type mismatch",
			in:    data.Map{},
			format: &struct {
				UintField uint `sbparam:",,a"`
			}{},
		},
		testSet{
			title: "minus value",
			in: data.Map{
				"UintField": data.Int(-1),
			},
			format: &struct{ UintField uint }{},
		},
		testSet{
			title: "uint8 value overflow",
			in: data.Map{
				"UintField": data.Int(math.MaxUint8 + 1),
			},
			format: &struct{ UintField uint8 }{},
		},
		testSet{
			title: "uint16 value overflow",
			in: data.Map{
				"UintField": data.Int(math.MaxUint16 + 1),
			},
			format: &struct{ UintField uint16 }{},
		},
		testSet{
			title: "uint32 value overflow",
			in: data.Map{
				"UintField": data.Int(math.MaxUint32 + 1),
			},
			format: &struct{ UintField uint32 }{},
		},
	}

	for _, c := range testCases {
		actual := c.format
		if err := Unmarshal(c.in, actual); err == nil {
			t.Fatalf("test case '%v' should occur an error but nothing", c.title)
		}
	}
}
