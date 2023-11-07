package utils

import "runtime"

// 获取调用栈中最后一个函数名
func GetCurrentFunctionName() string {
	pc, _, _, _ := runtime.Caller(1)
	return runtime.FuncForPC(pc).Name()
}
