package main

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
)

func main() {
	// 定义一个包含注册信息的映射
	params := map[string]interface{}{
		"passport":  "",        // 用户的账号
		"password":  "123456",  // 用户的密码
		"password2": "1234567", // 用户确认的密码
	}
	// 定义验证规则
	rules := map[string]string{
		"passport":  "required|length:6,16",                // 账号是必须的，长度在6到16之间
		"password":  "required|length:6,16|same:password2", // 密码是必须的，长度在6到16之间，且必须和password2字段相同
		"password2": "required|length:6,16",                // password2字段是必须的，长度在6到16之间
	}
	// 定义错误消息
	messages := map[string]interface{}{
		"passport": "账号不能为空|账号长度应当在{min}到{max}之间", // 账号相关的错误消息
		"password": map[string]string{ // 密码相关的错误消息
			"required": "密码不能为空",
			"same":     "两次密码输入不相等",
		},
	}
	// 创建一个新的验证器，设置错误消息，验证规则和数据，然后运行验证
	err := g.Validator().Messages(messages).Rules(rules).Data(params).Run(gctx.New())
	// 如果验证出错，打印错误信息
	if err != nil {
		g.Dump(err.Maps())
	}
}
