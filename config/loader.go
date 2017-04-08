package config

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"reflect"
	"strconv"
	"strings"
)


func LoadEnv(envObj interface{}) error {
    envs, err := loadConfDictFromEnv()
    if err != nil {
        return err
    }
    return loadObjFromDict(envObj, envs)
}

func LoadConf(confObj interface{}, confFile string) error {
    confDict, err := loadConfDictFromFile(confFile)
    if err != nil {
        return err
    }
    return loadObjFromDict(confObj, confDict)
}

func loadConfDictFromEnv() (map[string]string, error) {
    osEnvs := os.Environ()
    confs := make(map[string]string)
    for _,env := range osEnvs {
        if key,val, err := parseLine(env); err != nil {
            return confs, err
        } else {
            confs[key] = val
        }
    }
    return confs, nil
}

func loadConfDictFromFile(confFile string) (map[string]string, error) {
	fp, err := os.Open(confFile)
	if err != nil {
		return nil, err
	}
	defer fp.Close()
	fileBuf := bufio.NewReader(fp)
	confs := make(map[string]string)
	for lineNum := 0; ; {
		lineNum++
		line, err := fileBuf.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			} else if err == bufio.ErrBufferFull {
				return confs, fmt.Errorf("Line [%d] is too long", lineNum)
			} else {
				return confs, err
			}
		}

		line = strings.TrimSpace(line)
		if line[0] == '#' {
			continue
		}

        if key,val,err := parseLine(line); err != nil {
            return confs, fmt.Errorf("Line [%d]:%s", lineNum, err.Error())
        } else {
            confs[key] = val
        }
	}
	return confs, nil
}

func parseLine(line string) (key,value string, err error) {
    delimPos := strings.Index(line, "=")
    if delimPos == -1 {
        return "","", fmt.Errorf("Invalid line:[%s]", line)
    }
    key = strings.TrimSpace(line[:delimPos])
    value = strings.Trim(line[delimPos+1:], " \r\n\t")
    err = nil
    return
}

func loadObjFromDict(confObj interface{}, source map[string]string) error {
	if confObj == nil || source == nil || len(source) == 0 {
		return errors.New("Invalid Arguments")
	}
	refl := reflect.ValueOf(confObj).Elem()
	for i := 0; i != refl.NumField(); i++ {
		valField := refl.Field(i)
		tagField := refl.Type().Field(i).Tag

		isRequiredConf := false
		confKey := ""
		confVal := ""

		if tmpKey, ok := tagField.Lookup("soarConf_required"); ok {
			isRequiredConf = true
			confKey = tmpKey
		} else if tmpKey, ok := tagField.Lookup("soarConf"); ok {
			confKey = tmpKey
		} else {
			continue
		}
		if tmpVal, ok := source[confKey]; ok && tmpVal != "" {
			confVal = tmpVal
		}
		if confVal == "" {
			if isRequiredConf {
				return fmt.Errorf("Conf [%s] is required but not set", confKey)
			}
			continue
		}

		k := valField.Kind()
		switch k {
		case reflect.Int, reflect.Int32, reflect.Int64:
			bitSize := 32
			if k == reflect.Int64 {
				bitSize = 64
			}
			base := 10
			if strings.HasPrefix(confVal, "0x") {
				base = 16
			} else if confVal[0] == '0' {
				base = 8
			}
			tmpInt, err := strconv.ParseInt(confVal, base, bitSize)
			if err != nil {
				return err
			}
			valField.SetInt(tmpInt)
		case reflect.String:
			valField.SetString(confVal)
		case reflect.Bool:
			tmpBoolStr := strings.ToLower(confVal)
			if "true" == tmpBoolStr || "yes" == tmpBoolStr {
				valField.SetBool(true)
			} else if "false" == tmpBoolStr || "no" == tmpBoolStr {
				valField.SetBool(false)
			} else {
				return fmt.Errorf("Cannot convert [%s] to boolean", confVal)
			}
		case reflect.Float32, reflect.Float64:
			bitSize := 32
			if k == reflect.Float64 {
				bitSize = 64
			}
			tmpFloat, err := strconv.ParseFloat(confVal, bitSize)
			if err != nil {
				return err
			}
			valField.SetFloat(tmpFloat)
		case reflect.Uint, reflect.Uint32, reflect.Uint64:
			bitSize := 32
			if k == reflect.Uint64 {
				bitSize = 64
			}
			base := 10
			if strings.HasPrefix(confVal, "0x") {
				base = 16
			} else if confVal[0] == '0' {
				base = 8
			}
			tmpUint, err := strconv.ParseUint(confVal, base, bitSize)
			if err != nil {
				return err
			}
			valField.SetUint(tmpUint)
		default:
			return fmt.Errorf(
				"Not Supported Type [%s] for Conf:[%s]",
				k.String(),
				confKey,
			)
		}

	}
	return nil
}
