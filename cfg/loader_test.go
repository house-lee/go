package cfg

import (
	"bufio"
	"errors"
	"fmt"
	"github.com/house-lee/SoarGO/goinc"
	"io"
	"reflect"
	"strings"
	"testing"
	"os"
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

type mockFS struct{}
type mockFile struct {
	goinc.StubFile
}

var fileOpenErr error

func (mockFS) Open(name string) (goinc.File, error) {
	return &mockFile{}, fileOpenErr
}

type mockBufIO struct{}

var configStr string
var mockBufIOReaderObj goinc.BufIOReader
func (mockBufIO) NewReader(rd io.Reader) goinc.BufIOReader {
	return mockBufIOReaderObj
}

type mockBufIOReader struct {
	goinc.StubBufIO
}
var bufioReadStringError error
func (*mockBufIOReader) ReadString(delim byte) (string, error) {
	return "", bufioReadStringError
}

func TestLoadConfDictFromFileReturnsExpectedMap(t *testing.T) {
	configStr = `#This is an example conf file

PORT = 6379
    MYSQL_HOST = localhost
    MYSQL_PORT = 3306
  RUN_IN_BACKGROUND = true
PI = 3.1415926`
	mockBufIOReaderObj = bufio.NewReader(strings.NewReader(configStr))
	goinc.FS = mockFS{}
	goinc.BufIO = mockBufIO{}
	defer func() {
		goinc.FS = goinc.DefaultFS{}
		goinc.BufIO = goinc.DefaultBufIO{}
		configStr = ""
		mockBufIOReaderObj = nil
	}()


	expectedMap := map[string]string{
		"PORT":              "6379",
		"MYSQL_HOST":        "localhost",
		"MYSQL_PORT":        "3306",
		"RUN_IN_BACKGROUND": "true",
		"PI":                "3.1415926",
	}
	res, err := loadConfDictFromFile("whatever")
	if err != nil {
		t.Errorf("Encountered unexpected error: [%s]", err.Error())
	}
	if !reflect.DeepEqual(res, expectedMap) {
		t.Error("loadConfDictFromFile didn't load expected map")
	}
}

func TestLoadConfDictFromFileReturnsErrorWithParsedMapIfConfigFileIsInvalid(t *testing.T) {
	configStr = `#This is an example conf file

PORT = 6379
    MYSQL_HOST = localhost
    MYSQL_PORT = 3306
  RUN_IN_BACKGROUND = true
PI = 3.1415926

This is NOT a valid line
`
	mockBufIOReaderObj = bufio.NewReader(strings.NewReader(configStr))
	goinc.FS = mockFS{}
	goinc.BufIO = mockBufIO{}
	defer func() {
		goinc.FS = goinc.DefaultFS{}
		goinc.BufIO = goinc.DefaultBufIO{}
		configStr = ""
		mockBufIOReaderObj = nil
	}()
	expectedMap := map[string]string{
		"PORT":              "6379",
		"MYSQL_HOST":        "localhost",
		"MYSQL_PORT":        "3306",
		"RUN_IN_BACKGROUND": "true",
		"PI":                "3.1415926",
	}
	res, err := loadConfDictFromFile("whatever")
	if err == nil || !strings.Contains(err.Error(), "This is NOT a valid line") {
		t.Errorf("Encountered unexpected error: [%s]", err.Error())
	}
	if !reflect.DeepEqual(res, expectedMap) {
		t.Error("loadConfDictFromFile didn't load expected map")
	}
}

func TestLoadConfDictFromFileReturnsErrIfFailedToOpenFile(t *testing.T) {
	goinc.FS = mockFS{}
	fileOpenErr = errors.New("OPEN_FILE_FAILED")
	defer func() {
		goinc.FS = goinc.DefaultFS{}
		fileOpenErr = nil
	}()
	_, err := loadConfDictFromFile("whatever")
	if err == nil || err.Error() != "OPEN_FILE_FAILED" {
		t.Error("loadConfDictFromFile should return error if failed to open the config file")
	}

}
func TestLoadConfDictFromFileReturnsErrIfLineTooLong(t *testing.T) {
	bufioReadStringError = bufio.ErrBufferFull
	mockBufIOReaderObj = &mockBufIOReader{}
	goinc.FS = mockFS{}
	goinc.BufIO = mockBufIO{}
	defer func() {
		goinc.FS = goinc.DefaultFS{}
		goinc.BufIO = goinc.DefaultBufIO{}
		mockBufIOReaderObj = nil
	}()
	_, err :=  loadConfDictFromFile("whatever")
	if !strings.Contains(err.Error(), "is too long") {
		t.Error("loadConfDictFromFile should return error if a line in the config file is too long")
	}
}

func TestLoadConfDictFromEnvReturnsExpectedMap(t *testing.T) {
	backupEnv()
	defer restoreEnv()
	os.Clearenv()
	expectedMap := map[string]string{
		"PORT":              "6379",
		"MYSQL_HOST":        "localhost",
		"MYSQL_PORT":        "3306",
		"RUN_IN_BACKGROUND": "true",
		"PI":                "3.1415926",
	}
	for k,v := range expectedMap {
		os.Setenv(k,v)
	}
	m,err := loadConfDictFromEnv()
	if err != nil {
		t.Errorf("Encountered unexpected error: [%s]", err.Error())
	}
	if !reflect.DeepEqual(m, expectedMap) {
		t.Error("loadConfDictFromEnv didn't load expected map")
	}
}

func TestLoadConfCallLoadDictFromFileAndLoadObjFromDict(t *testing.T) {
	callSequence = ""
	loadConfDictFromFile = mockLoadConfDictFromFile
	loadObjFromDict = mockLoadObjFromDict
	defer func() {
		loadConfDictFromFile = loadConfDictFromFileFunc
		loadObjFromDict = loadObjFromDictFunc
	}()
	LoadConf(nil, "whatever")
	if callSequence != "loadConfDictFromFile,loadObjFromDict" {
		t.Error("LoadConf should call loadConfDictFromFile and loadObjFromDict respectively")
	}
}

func TestLoadEnvCallLoadDictFromFileAndLoadObjFromDict(t *testing.T) {
	callSequence = ""
	loadConfDictFromEnv = mockLoadConfDictFromEnv
	loadObjFromDict = mockLoadObjFromDict
	defer func() {
		loadConfDictFromEnv = loadConfDictFromEnvFunc
		loadObjFromDict = loadObjFromDictFunc
	}()
	LoadEnv(nil)
	if callSequence != "loadConfDictFromEnv,loadObjFromDict" {
		t.Error("LoadConf should call loadConfDictFromEnv and loadObjFromDict respectively")
	}
}

var callSequence string = ""
func mockLoadConfDictFromFile(confFile string) (map[string]string, error) {
	callSequence += "loadConfDictFromFile,"
	return nil,nil
}
func mockLoadConfDictFromEnv () (map[string]string, error) {
	callSequence += "loadConfDictFromEnv,"
	return nil, nil
}
func mockLoadObjFromDict(confObj interface{}, source map[string]string) error {
	callSequence += "loadObjFromDict"
	return nil
}

var envBuffer []string
func backupEnv() {
	envBuffer = os.Environ()
}
func restoreEnv() {
	os.Clearenv()
	for _,env := range envBuffer {
		item := strings.Split(env, "=")
		os.Setenv(item[0], item[1])
	}
}