package tgbot

import (
	"bytes"
	"fmt"
	"strings"

	"HautBot/jwglxt"

	"text/template"

	tb "gopkg.in/tucnak/telebot.v2"
)

var scoreTemplate = template.Must(template.New("score").Parse(
	"您本学期的成绩如下:\n\n{{range .Items}}{{.Name}}: {{.Score}}\n{{end}}",
))

func start(m *tb.Message) {
	botServer.bot.Send(m.Sender, "你好，我是 haut 教务系统机器人，请以以下格式输入您的学号和密码：\n\n/login 学号 密码")
}

func login(m *tb.Message) {
	infos := strings.Split(m.Payload, " ")
	if len(infos) == 2 {
		username := infos[0]
		password := infos[1]
		stu := jwglxt.New(username, password)
		if err := stu.Login(); err != nil {
			botServer.bot.Send(m.Sender, "登录失败，请检查学号和密码是否正确")
			return
		}
		info, err := stu.GetUserInfo()
		if err != nil {
			botServer.bot.Send(m.Sender, "获取信息失败，请稍后再试")
			return
		}
		botServer.bot.Send(m.Sender, fmt.Sprintf("你好，%s\n欢迎使用 haut 教务系统机器人", info.Name))
		botServer.cache.SetDefault(fmt.Sprintf("%d", m.Sender.ID), stu)
		botServer.db.SetUser(m.Sender.ID, username, password)
	}
}

func score(m *tb.Message) {
	stu_interface, ok := botServer.cache.Get(fmt.Sprintf("%d", m.Sender.ID))
	if !ok {
		username, password, err := botServer.db.GetUser(m.Sender.ID)
		if err != nil {
			fmt.Println(username, password, err)
			botServer.bot.Send(m.Sender, "数据库中没有您的信息，请先登录")
			return
		}
		stu := jwglxt.New(username, password)
		if err := stu.Login(); err != nil {
			botServer.bot.Send(m.Sender, "学号或密码错误")
			return
		}
		botServer.cache.SetDefault(fmt.Sprintf("%d", m.Sender.ID), stu)
		stu_interface = stu
	}
	stu := stu_interface.(*jwglxt.STU)
	score, err := stu.GetScore()
	if err != nil {
		botServer.bot.Send(m.Sender, "获取成绩失败，请稍后再试")
		return
	}
	msg := bytes.NewBuffer(nil)
	if err = scoreTemplate.Execute(msg, score); err != nil {
		fmt.Println(err)
	}
	botServer.bot.Send(m.Sender, msg.String(), &tb.SendOptions{ParseMode: tb.ModeMarkdown})
}
