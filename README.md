# cryo-plugin-echo

这是 [🧊cryo](https://github.com/machinacanis/cryo) 的一个示例插件，通过响应规则过滤实现了一个简单的消息重复功能。

## 快速开始

```bash
go get -u github.com/machinacanis/cryo-plugin-echo
```

在项目中使用插件：

```go
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
```