package main

import (
	"github.com/machinacanis/cryo"
	"github.com/machinacanis/cryo-plugin-echo"
)

func main() {
	bot := cryo.NewBot()
	bot.Init()

	bot.AddPlugin(CryoPluginEcho.Instance) // 添加插件

	bot.AutoConnect()
	bot.Start()
}
