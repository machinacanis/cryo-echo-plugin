package main

import (
	"github.com/machinacanis/cryo"
	"github.com/machinacanis/cryo-echo-plugin"
)

func main() {
	bot := cryo.NewBot()
	bot.Init()

	bot.AddPlugin(CryoEcho.Plugin()) // 添加插件

	bot.AutoConnect()
	bot.Start()
}
