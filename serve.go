//+build !withsource
//默认使用http运行

package main

func serve() (string, error) {
	// umi默认端口为8000
	return "http://localhost:8000", nil
}