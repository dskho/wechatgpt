package bootstrap

import (
	"github.com/eatmoreapple/openwechat"
	log "github.com/sirupsen/logrus"
	"github.com/wechatgpt/wechatbot/handler/wechat"
)

func StartWebChat() {
	bot := openwechat.DefaultBot(openwechat.Desktop)
	bot.MessageHandler = wechat.Handler
	bot.UUIDCallback = openwechat.PrintlnQrcodeUrl

	reloadStorage := openwechat.NewJsonFileHotReloadStorage("token.json")
	err := bot.HotLogin(reloadStorage)
	if err != nil {
		if err = bot.Login(); err != nil {
			log.Fatal(err)
			return
		}
	}

	// 获取登陆的用户
	self, err := bot.GetCurrentUser()
	if err != nil {
		log.Fatal(err)
		return
	}

	friends, err := self.Friends()

	for i, friend := range friends {
		log.Println(i, friend)
	}
	groups, err := self.Groups()
	for i, group := range groups {
		log.Println(i, group)
	}

	err = bot.Block()
	if err != nil {
		log.Fatal(err)
		return
	}

}
