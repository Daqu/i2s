package i2s

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
)

/*Decode function
* data interface{} the data you want to use for init
* s interface{} a struct
 */
func Decode(data interface{}, s interface{}) error {
	switch i := data.(type) {
	case nil:
		fmt.Println("get nil value")
	case map[string]interface{}:
		t := reflect.TypeOf(s).Elem()
		for i := 0; i < t.NumField(); i++ {
			k := t.Field(i).Tag.Get("json")
			if _, ok := data.(map[string]interface{})[k]; ok {
				continue
			} else {
				return errors.New("Your data is not match the struct you send in")
			}
		}
	default:
		fmt.Printf("unsupport type:%v\n", i)
	}
	SetStructFieldByJsonName(s, data.(map[string]interface{}))
	return nil
}

/*SetStructFieldByJsonName init strcut by json tag, ref:https://www.cnblogs.com/fwdqxl/p/7789162.html
* ptr the pointer of struct which want to be init
* data use for init
 */
func SetStructFieldByJsonName(ptr interface{}, data map[string]interface{}) {

	v := reflect.ValueOf(ptr).Elem() // the struct variable

	for i := 0; i < v.NumField(); i++ {

		fieldInfo := v.Type().Field(i) // a reflect.StructField
		tag := fieldInfo.Tag           // a reflect.StructTag
		name := tag.Get("json")

		if name == "" {
			name = strings.ToLower(fieldInfo.Name)
		}
		//去掉逗号后面内容 如 `json:"voucher_usage,omitempty"`
		name = strings.Split(name, ",")[0]

		if value, ok := data[name]; ok {
			//给结构体赋值
			//保证赋值时数据类型一致
			if reflect.ValueOf(value).Type() == v.FieldByName(fieldInfo.Name).Type() {
				v.FieldByName(fieldInfo.Name).Set(reflect.ValueOf(value))
			}

		}
	}

	return
}
