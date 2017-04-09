package config

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
)

type confStruct struct {
	IntConf     int     `soarConf_required:"INT_CONF"`
	Int32Conf   int32   `soarConf:"INT32_CONF"`
	Int64Conf   int64   `soarConf:"INT64_CONF"`
	HexUintConf uint    `soarConf:"HEX_UINT_CONF"`
	OctUintConf uint    `soarConf:"OCT_UINT_CONF"`
	HexIntConf  int     `soarConf:"HEX_INT_CONF"`
	OctIntConf  int     `soarConf:"OCT_INT_CONF"`
	StrConf     string  `soarConf:"STR_CONF"`
	BoolConf    bool    `soarConf_required:"BOOL_CONF"`
	Float32Conf float32 `soarConf:"FLOAT32_CONF"`
	Float64Conf float64 `soarConf:"FLOAT64_CONF"`
	UintConf    uint    `soarConf:"UINT_CONF"`
	Uint32Conf  uint32  `soarConf:"UINT32_CONF"`
	Uint64Conf  uint64  `soarConf:"UINT64_CONF"`
	OptBoolConf bool    `soarConf:"OPT_BOOL_CONF"`
}

func TestLoadObjFromDictLoadsExpectedConf(t *testing.T) {
	expectedConf := confStruct{
		IntConf:     -2048,
		Int32Conf:   -2147483647,
		Int64Conf:   -9223372036854775808,
		HexUintConf: 0x9aef,
		OctUintConf: 0127,
		HexIntConf:  0xfcd8,
		OctIntConf:  0352,
		StrConf:     "Test_string_conf",
		BoolConf:    true,
		Float32Conf: 1.234,
		Float64Conf: 2.56789,
		UintConf:    4096,
		Uint32Conf:  4294967295,
		Uint64Conf:  18446744073709551615,
	}

	testSource := make(map[string]string)
	ref := reflect.ValueOf(&expectedConf).Elem()
	for i := 0; i != ref.NumField(); i++ {
		valField := ref.Field(i)
		tagField := ref.Type().Field(i).Tag
		confKey := tagField.Get("soarConf_required")
		if confKey == "" {
			confKey = tagField.Get("soarConf")
		}
		confVal := ""
		switch confKey {
		case "HEX_UINT_CONF":
			confVal = fmt.Sprintf("0x%x", valField.Uint())
		case "OCT_UINT_CONF":
			confVal = fmt.Sprintf("0%o", valField.Uint())
		case "HEX_INT_CONF":
			confVal = fmt.Sprintf("0x%x", valField.Int())
		case "OCT_INT_CONF":
			confVal = fmt.Sprintf("0%o", valField.Int())
		default:
			confVal = fmt.Sprintf("%v", valField.Interface())
		}
		testSource[confKey] = confVal
	}

	confObj := confStruct{}
	err := loadObjFromDict(&confObj, testSource)
	if err != nil || !reflect.DeepEqual(&expectedConf, &confObj) {
		t.Error("loadObjFromDict failed to load expected configurations")
	}
}

func TestLoadObjFromDictReturnsErrIfRequiredConfNotSet(t *testing.T) {
	testConfs := map[string]string{
		"INT32_CONF": "255",
		"STR_CONF":   "StringConfgiration",
	}

	confObj := confStruct{}
	err := loadObjFromDict(&confObj, testConfs)
	if err == nil || !strings.Contains(err.Error(), "is required but not set") {
		t.Error("loadObjFromDict should return expected error if required conf")
	}
}

func TestLoadObjFromDictShouldRecognizeBothYesAndTrue(t *testing.T) {
	testConfs := map[string]string{
		"INT_CONF":      "123",
		"BOOL_CONF":     "Yes",
		"OPT_BOOL_CONF": "YES",
	}

	confObj := confStruct{}
	err := loadObjFromDict(&confObj, testConfs)
	if err != nil || confObj.BoolConf != true || confObj.OptBoolConf != true {
		t.Error("loadObjFromDict should recognize \"Yes\" as boolean true")
	}

	testConfs = map[string]string{
		"INT_CONF":      "123",
		"BOOL_CONF":     "True",
		"OPT_BOOL_CONF": "TRuE",
	}

	confObj = confStruct{}
	err = loadObjFromDict(&confObj, testConfs)
	if err != nil || confObj.BoolConf != true || confObj.OptBoolConf != true {
		t.Error("loadObjFromDict should recognize \"True\" as boolean true")
	}
}

func TestLoadObjFromDictShouldRecognizeBothNoAndFalse(t *testing.T) {
	testConfs := map[string]string{
		"INT_CONF":      "123",
		"BOOL_CONF":     "No",
		"OPT_BOOL_CONF": "NO",
	}

	confObj := confStruct{}
	err := loadObjFromDict(&confObj, testConfs)
	if err != nil || confObj.BoolConf != false || confObj.OptBoolConf != false {
		t.Error("loadObjFromDict should recognize \"No\" as boolean false")
	}

	testConfs = map[string]string{
		"INT_CONF":      "123",
		"BOOL_CONF":     "False",
		"OPT_BOOL_CONF": "False",
	}

	confObj = confStruct{}
	err = loadObjFromDict(&confObj, testConfs)
	if err != nil || confObj.BoolConf != false || confObj.OptBoolConf != false {
		t.Error("loadObjFromDict should recognize \"False\" as boolean true")
	}
}

func TestLoadObjFromDictShouldReturnErrIfTypeNotMatch(t *testing.T) {
	testConfs := map[string]string{
		"INT_CONF":  "NOT_AN_INT",
		"BOOL_CONF": "Yes",
	}

	confObj := confStruct{}
	err := loadObjFromDict(&confObj, testConfs)
	if err == nil {
		t.Error("loadObjFromDict should return error if int type not match")
	}

	testConfs = map[string]string{
		"INT_CONF":  "123",
		"BOOL_CONF": "NOT_A_BOOL",
	}

	confObj = confStruct{}
	err = loadObjFromDict(&confObj, testConfs)
	if err == nil {
		t.Error("loadObjFromDict should return error if bool type not match")
	}

	testConfs = map[string]string{
		"INT_CONF":     "123",
		"BOOL_CONF":    "Yes",
		"FLOAT32_CONF": "NOT_A_FLOAT",
	}

	confObj = confStruct{}
	err = loadObjFromDict(&confObj, testConfs)
	if err == nil {
		t.Error("loadObjFromDict should return error if float type not match")
	}

	testConfs = map[string]string{
		"INT_CONF":  "123",
		"BOOL_CONF": "Yes",
		"UINT_CONF": "-256",
	}

	confObj = confStruct{}
	err = loadObjFromDict(&confObj, testConfs)
	if err == nil {
		t.Error("loadObjFromDict should return error if uint type not match")
	}
}

func TestParseLineReturnsExpectedKeyValuePair(t *testing.T) {
	testData := [][]string{
		{"k1=v1", "k1", "v1"},
		{"   k2  = v2 ", "k2", "v2"},
		{"k3 = v3 \r", "k3", "v3"},
		{"k4 = v4 \n", "k4", "v4"},
		{"k5 = v5 \t", "k5", "v5"},
		{"k6 = v6 \t\r\n", "k6", "v6"},
	}
	for _, testCase := range testData {
		k, v, e := parseLine(testCase[0])
        if e != nil {
            t.Errorf("parseLine failed. [%s]", e.Error())
        }
		if k != testCase[1] || v != testCase[2] {
			t.Errorf(
				"parseLine failed. src: [%s],expected [%s]=[%s],got [%s]=[%s]",
                testCase[0],
                testCase[1],
                testCase[2],
                k,
                v,
			)
		}
	}
}
