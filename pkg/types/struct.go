package types

import (
	"fmt"
	"reflect"
	"strings"

	"google.golang.org/protobuf/types/known/structpb"
)

// ToPBArrayStruct - converts a []map[string]any to a []pb.Struct
func ToPBArrayStruct(a []map[string]any) []*structpb.Struct {
	var s []*structpb.Struct
	for _, m := range a {
		s = append(s, ToPBStruct(m))
	}
	return s
}

// ToPBStruct - converts a map[string]any to a pb.Struct.
func ToPBStruct(v map[string]any) *structpb.Struct {
	size := len(v)
	if size == 0 {
		return nil
	}
	fields := make(map[string]*structpb.Value, size)
	for k, v := range v {
		fields[k] = ToValue(v)
	}
	return &structpb.Struct{
		Fields: fields,
	}
}

// ToPBArrayValue - converts a []any to a []pb.Value.
func ToPBArrayValue(a []any) []*structpb.Value {
	var s []*structpb.Value
	for _, v := range a {
		s = append(s, ToValue(v))
	}
	return s
}

// ToValue - converts an any to a pb.Value.
func ToValue(v any) *structpb.Value {
	switch v := v.(type) {
	case nil:
		return nil
	case bool:
		return &structpb.Value{
			Kind: &structpb.Value_BoolValue{
				BoolValue: v,
			},
		}
	case int:
		return &structpb.Value{
			Kind: &structpb.Value_NumberValue{
				NumberValue: float64(v),
			},
		}
	case int8:
		return &structpb.Value{
			Kind: &structpb.Value_NumberValue{
				NumberValue: float64(v),
			},
		}
	case int32:
		return &structpb.Value{
			Kind: &structpb.Value_NumberValue{
				NumberValue: float64(v),
			},
		}
	case int64:
		return &structpb.Value{
			Kind: &structpb.Value_NumberValue{
				NumberValue: float64(v),
			},
		}
	case uint:
		return &structpb.Value{
			Kind: &structpb.Value_NumberValue{
				NumberValue: float64(v),
			},
		}
	case uint8:
		return &structpb.Value{
			Kind: &structpb.Value_NumberValue{
				NumberValue: float64(v),
			},
		}
	case uint32:
		return &structpb.Value{
			Kind: &structpb.Value_NumberValue{
				NumberValue: float64(v),
			},
		}
	case uint64:
		return &structpb.Value{
			Kind: &structpb.Value_NumberValue{
				NumberValue: float64(v),
			},
		}
	case float32:
		return &structpb.Value{
			Kind: &structpb.Value_NumberValue{
				NumberValue: float64(v),
			},
		}
	case float64:
		return &structpb.Value{
			Kind: &structpb.Value_NumberValue{
				NumberValue: v,
			},
		}
	case string:
		return &structpb.Value{
			Kind: &structpb.Value_StringValue{
				StringValue: v,
			},
		}
	case error:
		return &structpb.Value{
			Kind: &structpb.Value_StringValue{
				StringValue: v.Error(),
			},
		}
	default:
		// Fallback to reflection for other types
		return toValue(reflect.ValueOf(v))
	}
}

func toValue(v reflect.Value) *structpb.Value {
	switch v.Kind() {
	case reflect.Bool:
		return &structpb.Value{
			Kind: &structpb.Value_BoolValue{
				BoolValue: v.Bool(),
			},
		}
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return &structpb.Value{
			Kind: &structpb.Value_NumberValue{
				NumberValue: float64(v.Int()),
			},
		}
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return &structpb.Value{
			Kind: &structpb.Value_NumberValue{
				NumberValue: float64(v.Uint()),
			},
		}
	case reflect.Float32, reflect.Float64:
		return &structpb.Value{
			Kind: &structpb.Value_NumberValue{
				NumberValue: v.Float(),
			},
		}
	case reflect.Ptr:
		if v.IsNil() {
			return nil
		}
		return toValue(reflect.Indirect(v))
	case reflect.Array, reflect.Slice:
		size := v.Len()
		if size == 0 {
			return nil
		}
		values := make([]*structpb.Value, size)
		for i := 0; i < size; i++ {
			values[i] = toValue(v.Index(i))
		}
		return &structpb.Value{
			Kind: &structpb.Value_ListValue{
				ListValue: &structpb.ListValue{
					Values: values,
				},
			},
		}
	case reflect.Struct:
		t := v.Type()
		size := v.NumField()
		if size == 0 {
			return nil
		}
		fields := make(map[string]*structpb.Value, size)
		for i := 0; i < size; i++ {
			// name := t.Field(i).Name
			val, ok := t.Field(i).Tag.Lookup("json")
			if ok && len(val) > 0 {
				name := strings.Split(val, ",")[0]
				// Better way?
				if len(name) > 0 && 'A' <= name[0] && name[0] <= 'z' {
					fields[name] = toValue(v.Field(i))
				}
			}
		}
		if len(fields) == 0 {
			return nil
		}
		return &structpb.Value{
			Kind: &structpb.Value_StructValue{
				StructValue: &structpb.Struct{
					Fields: fields,
				},
			},
		}
	case reflect.Map:
		keys := v.MapKeys()
		if len(keys) == 0 {
			return nil
		}
		fields := make(map[string]*structpb.Value, len(keys))
		for _, k := range keys {
			if k.Kind() == reflect.String {
				fields[k.String()] = toValue(v.MapIndex(k))
			}
		}
		if len(fields) == 0 {
			return nil
		}
		return &structpb.Value{
			Kind: &structpb.Value_StructValue{
				StructValue: &structpb.Struct{
					Fields: fields,
				},
			},
		}
	case reflect.Interface:
		return toValue(v.Elem())
	default:
		// Last resort
		return &structpb.Value{
			Kind: &structpb.Value_StringValue{
				StringValue: fmt.Sprint(v),
			},
		}
	}
}

// DecodePBToArrayMap - converts a []pb.struct to array map form strings to Go types.
func DecodePBToArrayMap(a []*structpb.Struct) []map[string]any {
	var s []map[string]any
	for _, m := range a {
		s = append(s, DecodePBToMap(m))
	}
	return s
}

// DecodePBToMap - converts a pb.Struct to a map from strings to Go types.
// DecodePBToMap panics if s is invalid.
func DecodePBToMap(s *structpb.Struct) map[string]any {
	if s == nil {
		return nil
	}
	m := map[string]any{}
	for k, v := range s.Fields {
		m[k] = decodeValue(v)
	}
	return m
}

func decodeValue(v *structpb.Value) any {
	switch k := v.Kind.(type) {
	case *structpb.Value_NullValue:
		return nil
	case *structpb.Value_NumberValue:
		return k.NumberValue
	case *structpb.Value_StringValue:
		return k.StringValue
	case *structpb.Value_BoolValue:
		return k.BoolValue
	case *structpb.Value_StructValue:
		return DecodePBToMap(k.StructValue)
	case *structpb.Value_ListValue:
		s := make([]any, len(k.ListValue.Values))
		for i, e := range k.ListValue.Values {
			s[i] = decodeValue(e)
		}
		return s
	default:
		panic("proto struct: unknown kind")
	}
}
