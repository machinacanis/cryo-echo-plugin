package CryoEcho

import (
	"github.com/machinacanis/cryo"
)

var GroupAtMeRule = cryo.RuleFor(func(e *cryo.GroupMessageEvent) bool {
	msg := *e.GetMessage()
	newmsg := cryo.Message{}
	atFound := false

	for i := 0; i < len(msg); i++ {
		// 检查当前元素是否为at类型且目标是当前用户
		if msg[i].GetType() == cryo.AtType {
			at, ok := msg[i].(*cryo.At)
			if ok && at.TargetUin == e.GetClient().Uin {
				// 找到了at当前用户的元素，跳过它不添加到newmsg中
				atFound = true

				// 检查下一个元素是否存在且为文本类型
				if i+1 < len(msg) && msg[i+1].GetType() == cryo.TextType {
					text, ok := msg[i+1].(*cryo.Text)
					if ok && len(text.Content) > 0 && text.Content[0] == ' ' {
						// 如果文本内容以空格开头，则去除开头的空格
						// log.Info("发现文本以空格开头，去除首个空格")
						text.Content = text.Content[1:]
					}
				}
				continue
			}
		}

		// 将当前元素添加到新消息中
		newmsg = append(newmsg, msg[i])
	}

	// 如果找到了at当前用户的元素，替换事件中的消息并返回true
	if atFound {
		// log.Info("发现At，已从消息中移除")
		*e.GetMessage() = newmsg
		return true
	}

	return false
})
