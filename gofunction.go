package main

import "log"

func registerFunc() {
	ui.Bind("query", func(any string) string {
		log.Println("被调用" + any)
		ui.Eval(`
			console.log("go invoke success")
		`)
		return "query ok"
	})
}
