package logging

import "log"

var _isEnable bool = false

func LogInfo(data ...interface{}) {
	log.Println(data...)
}

func LogWarning(data ...interface{}) {

}
