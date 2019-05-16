package reflectex

import (
	"errors"
	"reflect"
)

type ReflectX struct {
	T    reflect.Type
	V    reflect.Value
	THIS interface{}
}

func New(val interface{}) *ReflectX {
	t := reflect.TypeOf(val)
	t = DeRef(t)
	v := reflect.Indirect(reflect.ValueOf(val))
	return &ReflectX{
		T:    t,
		V:    v,
		THIS: val,
	}
}

// SetValueWithAny 设置值
func (rx *ReflectX) SetValueWithAny(fieldIndex int, value interface{}) {
	if fieldIndex < 0 {
		panic(errors.New("index must > 0"))
	}
	field := rx.V.Field(fieldIndex)
	field.Set(reflect.Indirect(reflect.ValueOf(value)))
	//switch field.Kind() {
	//case reflect.String:
	//	field.SetString(*value.(*string))
	//	break
	//case reflect.Int64:
	//	field.SetInt(*value.(*int64))
	//	break
	//default:
	//	log.Println("不支持的数据类型: ", field.Kind())
	//	break
	//}
}

// TStruct 获得T得Struct
func (rx *ReflectX) TStruct(pos int) reflect.StructField {
	return rx.T.Field(pos)
}

// FieldName 获得T得FieldName
func (rx *ReflectX) FieldName(pos int) string {
	return rx.TStruct(pos).Name
}

// GetDBEntity 生成当前对象的实体,可直接用在rows.scan(xxx...)里
// 大概的样子是[*int64, *string, *int64, *string] 这种
func (rx *ReflectX) GetDBEntity() []interface{} {
	// 新的对象，如果拥有同一个对象会出现错乱
	newVal := reflect.New(rx.T)
	newVal = reflect.Indirect(newVal)

	fNum := newVal.NumField() // 结构体内部属性个数
	dbEntity := make([]interface{}, fNum)
	for i := 0; i < fNum; i++ {
		field := newVal.Field(i)
		if field.CanAddr() {
			dbEntity[i] = field.Addr().Interface()
		} else {
			dbEntity[i] = field.Interface()
		}
	}
	return dbEntity
}

// DeRef 还原type的本质
func DeRef(t reflect.Type) reflect.Type {
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
	}
	return t
}
