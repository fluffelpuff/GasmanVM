package utils

import "reflect"

// IsSupportedDataType überprüft, ob der Wert eines unterstützten Datentyps entspricht.
func isSupportedDataType(value interface{}) bool {
	// Verwenden von Reflektion, um den genauen Typ zu ermitteln
	valueType := reflect.TypeOf(value)
	switch valueType.Kind() {
	case reflect.Int64, reflect.Uint64, reflect.Float64, reflect.String, reflect.Bool:
		// Unterstützte Datentypen
		return true
	case reflect.Slice:
		// Überprüfen, ob es sich um einen Slice ([]byte) handelt
		if valueType.Elem().Kind() == reflect.Uint8 {
			return true
		}
	case reflect.Map:
		// Überprüfen, ob es sich um eine Map (map[string]interface{}) handelt
		if valueType.Key().Kind() == reflect.String && valueType.Elem().Kind() == reflect.Interface {
			return true
		}
	case reflect.Struct:
		// Unterstützte Datentypen
		return true
	}
	return false
}

// CheckDataValues überprüft einen Wert oder eine Liste von Werten auf unterstützte Datentypen.
func CheckDataValues(data interface{}) bool {
	switch data := data.(type) {
	case []interface{}:
		// Wenn es sich um eine Liste handelt, einzelne Einträge prüfen.
		for _, entry := range data {
			if !CheckDataValues(entry) {
				return false
			}
		}
	case map[string]interface{}:
		// Wenn es sich um eine map handelt, rekursiv überprüfen.
		if !checkMapValues(data) {
			return false
		}
	default:
		// Wenn es sich nicht um eine Liste oder Map handelt, direkt überprüfen, ob der Wert unterstützt wird.
		if !isSupportedDataType(data) {
			return false
		}
	}
	return true
}

// CheckMapValues rekursiv überprüft die Werte in einer map[string]interface{}.
func checkMapValues(data map[string]interface{}) bool {
	for _, value := range data {
		if isMap, ok := value.(map[string]interface{}); ok {
			// Wenn es sich um eine weitere map handelt, rekursiv fortsetzen.
			if checkMapValues(isMap) {
				return true
			}
		} else if !isSupportedDataType(value) {
			// Wenn es sich nicht um einen unterstützten Datentyp handelt, return false.
			return false
		}
	}
	return true
}
