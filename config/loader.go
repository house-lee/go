package config

import (
    "reflect"
    "errors"
)

func loadObjFromDict(configObj interface{}, source map[string]string) error {
    if configObj == nil || source == nil || len(source) == 0 {
        return errors.New("Invalid Arguments")
    }
    refl := reflect.ValueOf(configObj).Elem()
    for i := 0; i != refl.NumField(); i++ {
        valField := refl.Field(i)
        tagField := refl.Type().Field(i).Tag

        required := false
        configKey := ""

        if tmpKey,ok := tagField.Lookup("soar_go_config_required")
    }
}