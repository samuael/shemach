package pdf

import (
	"html/template"
	"strings"
	"time"
)

func GetDate() string {
	time := time.Now()
	return strings.Split(time.String(), ".")[0]
}

// funcMap
var MyFuncMap = template.FuncMap{
	"getDate": GetDate,
}
