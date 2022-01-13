package tgbot

import (
	"HautBot/database"
	"encoding/json"
	"flag"
	"log"
	"net/http"
	"net/url"
	"time"

	cache "github.com/patrickmn/go-cache"
	tb "gopkg.in/tucnak/telebot.v2"
)

const (
	USUAL = iota
	USERNAME
	PASSWORD
)

type BotServer struct {
	bot   *tb.Bot
	cache *cache.Cache
	db    *database.Database
}

type Proxy string

func (p *Proxy) String() string {
	return string(*p)
}

func (p *Proxy) Set(s string) error {
	*p = Proxy(s)
	return nil
}

func ProxyFlag(name string, val Proxy, usage string) *Proxy {
	flag.Var(&val, name, usage)
	return &val
}

var client = http.DefaultClient
var botServer = func() *BotServer {
	bot, err := tb.NewBot(tb.Settings{
		Token:   "",
		Poller:  &tb.LongPoller{Timeout: 10 * time.Second},
		Client:  client,
		Offline: true,
	})
	if err != nil {
		panic(err)
	}
	c := cache.New(5*time.Minute, 10*time.Minute)
	return &BotServer{
		bot:   bot,
		cache: c,
		db:    &database.Database{},
	}
}()

func New(token string, v ...interface{}) *BotServer {
	for _, val := range v {
		switch val := val.(type) {
		case Proxy:
			url, err := url.Parse(string(val))
			if err != nil {
				panic(err)
			}
			client.Transport = &http.Transport{
				Proxy: http.ProxyURL(url),
			}
		}
	}
	botServer.bot.Token = token

	data, err := botServer.bot.Raw("getMe", nil)
	if err != nil {
		panic(err)
	}
	var resp struct {
		Result *tb.User
	}
	if err := json.Unmarshal(data, &resp); err != nil {
		panic(err)
	}
	botServer.bot.Me = resp.Result

	botServer.db, err = database.Open("database.db")
	if err != nil {
		panic(err)
	}

	return botServer
}

func (bs *BotServer) Run() {
	log.Printf("%s start", bs.bot.Me.Username)
	bs.bot.Handle("/start", start)
	bs.bot.Handle("/login", login)
	bs.bot.Handle("/score", score)
	bs.bot.Start()
}
