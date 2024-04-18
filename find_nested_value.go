package find_nested_value

import (
	"fmt"
	"reflect"
	"strings"
)

// FindNestedValue extract nested values from map and struct based on the path
func FindNestedValue(data any, path string) (any, error) {
	value := reflect.ValueOf(data)
	parts := strings.Split(path, ".")

	for i, part := range parts {
		// Handle pointers and interfaces
		if value.Kind() == reflect.Ptr || value.Kind() == reflect.Interface {
			value = value.Elem()
			if !value.IsValid() {
				return nil, fmt.Errorf("dereferencing pointer/interface resulted in invalid value")
			}
		}

		switch value.Kind() {
		case reflect.Struct:
			field := value.FieldByName(part)
			if !field.IsValid() {
				return nil, fmt.Errorf("field %s not found", part)
			}
			value = field

		case reflect.Map:
			if value.Type().Key().Kind() != reflect.String {
				return nil, fmt.Errorf("map keys must be of type string")
			}
			val := value.MapIndex(reflect.ValueOf(part))
			if !val.IsValid() {
				return nil, fmt.Errorf("key %s not found in map", part)
			}
			value = val

		default:
			if i != len(parts)-1 {
				return nil, fmt.Errorf("unsupported type %v for path navigation", value.Kind())
			}
			// If we're at the last part and the type is unsupported but valid, we still return the value
			// This allows for returning values like basic types (int, string) at the end of the path
		}
	}

	return value.Interface(), nil
}
