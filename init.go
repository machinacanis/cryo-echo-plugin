package CryoEcho

import (
	"github.com/machinacanis/cryo"
)

type CryoPluginEcho struct {
	bot        *cryo.Bot           // bot实例
	responsers []*cryo.OnResponser // 响应器列表

	pluginName        string // 插件名称
	pluginVersion     string // 插件版本
	pluginDescription string // 插件描述
	pluginAuthor      string // 插件作者
	isEnable          bool   // 是否启用插件
}

func (p *CryoPluginEcho) GetPluginName() string {
	return p.pluginName
}

func (p *CryoPluginEcho) GetPluginVersion() string {
	return p.pluginVersion
}

func (p *CryoPluginEcho) GetPluginDescription() string {
	return p.pluginDescription
}

func (p *CryoPluginEcho) GetPluginAuthor() string {
	return p.pluginAuthor
}

func (p *CryoPluginEcho) Enable() {
	for _, responser := range p.responsers {
		responser.Register()
	}

	p.responsers = nil
	p.isEnable = true
}

func (p *CryoPluginEcho) Disable() {
	for _, responser := range p.responsers {
		responser.Remove()
	}

	p.responsers = nil
	p.isEnable = false
}

func (p *CryoPluginEcho) IsEnable() bool {
	return p.isEnable
}

func (p *CryoPluginEcho) Init(bot *cryo.Bot) error {
	// 初始化插件
	p.bot = bot
	p.pluginName = "cryo_plugin_echo"
	p.pluginVersion = "1.0.0"
	p.pluginDescription = "一个简单的示例插件"
	p.pluginAuthor = "MachinaCanis"
	p.isEnable = false

	groupResponser := p.bot.
		OnType(cryo.GroupMessageEventType).
		AddRule(GroupAtMeRule).
		Handle(func(e *cryo.GroupMessageEvent) {
			msg := e.GetMessage()
			e.Send(msg)
		})
	p.responsers = append(p.responsers, groupResponser)
	privateResponser := p.bot.
		OnType(cryo.PrivateMessageEventType).
		Handle(func(e *cryo.PrivateMessageEvent) {
			msg := e.GetMessage()
			e.Send(msg)
		})
	p.responsers = append(p.responsers, privateResponser)

	return nil
}

func Plugin() *CryoPluginEcho {
	return &CryoPluginEcho{}
}
