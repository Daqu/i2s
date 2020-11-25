package i2s

import (
	"errors"
	"fmt"
	"reflect"
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
	return nil
}
