package handler

import (
	"api-gateway/pkg/util"
	"errors"
)

// PanicIfUserError 包装错误
func PanicIfUserError(err error) {
	if err != nil {
		err = errors.New("userService--" + err.Error())
		util.LogrusObj.Info(err)
		panic(err)
	}
}

// PanicIfPythonError 包装Python服务错误
func PanicIfPythonError(err error) {
	if err != nil {
		err = errors.New("PythonService" + err.Error())
		util.LogrusObj.Info(err)
		panic(err)
	}
}

// PanicIfJavaError 包装Java服务错误
func PanicIfJavaError(err error) {
	if err != nil {
		err = errors.New("JavaService" + err.Error())
		util.LogrusObj.Info(err)
		panic(err)
	}
}

func PanicIfTaskError(err error) {
	if err != nil {
		err = errors.New("taskService--" + err.Error())
		util.LogrusObj.Info(err)
		panic(err)
	}
}
