package main

import (
	"log"
	"strconv"
)

var count = 0

func registerFunc() {
	ui.Bind("query", func(any string) string {
		count += 1
		log.Printf("被调用 传参%s 调用%d次", any, count)
		ui.Eval(`
			console.log("go invoke success");
			document.getElementById("callbackPrint").innerHTML="golang 回调测试` + strconv.Itoa(count) + `";
		`)
		return "query ok"
	})
}
