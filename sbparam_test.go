package sbparam

import (
	"gopkg.in/sensorbee/sensorbee.v0/data"
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
		title string
		in    data.Map
	}
	testCases := []testSet{
		testSet{
			title: "lack required field, 'StrField1'",
			in: data.Map{
				"StrField4": data.String("b"),
				"StrField5": data.String("c"),
			},
		},
		testSet{
			title: "lack required field, 'StrField4'",
			in: data.Map{
				"str_field1": data.String("a"),
				"StrField5":  data.String("c"),
			},
		},
		testSet{
			title: "lack required field, 'StrField5'",
			in: data.Map{
				"str_field1": data.String("a"),
				"StrField4":  data.String("b"),
			},
		},
		testSet{
			title: "key mismatch",
			in: data.Map{
				"str_field1": data.Int(1),
				"StrField4":  data.String("b"),
				"StrField5":  data.String("c"),
			},
		},
	}

	for _, c := range testCases {
		actual := &testStrValue{}
		if err := Unmarshal(c.in, actual); err == nil {
			t.Fatalf("test case '%v' should occur an error but nothing", c.title)
		}
	}
}
