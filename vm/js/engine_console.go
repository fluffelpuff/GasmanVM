package jsengine

import "github.com/dop251/goja"

// Gibt ein Array von Interfaces[] aus
func convertoToInterfacesArray(parms []goja.Value) []interface{} {
	newArray := make([]interface{}, len(parms))
	for h, i := range parms {
		if i == nil {
			newArray[h] = "null"
		} else if i.String() == "undefined" {
			newArray[h] = "undefined"
		} else if i.String() == "null" {
			newArray[h] = "null"
		} else {
			newArray[h] = i.Export()
		}
	}
	return newArray
}
