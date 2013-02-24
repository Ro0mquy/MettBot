package mettbot

import (
	a "./answers"
	"fmt"
	irc "github.com/fluffle/goirc/client"
	"math/rand"
	"regexp"
	"strings"
	"time"
)

func (bot *Mettbot) HandlerConnected()    { bot.Join(*Channel) }
func (bot *Mettbot) HandlerDisconnected() { bot.Quitted <- true }

func (bot *Mettbot) HandlerJoin(line *irc.Line) {
	time.Sleep(1000 * time.Millisecond)
	actChannel := line.Args[0]
	bot.Topics[actChannel] = bot.ST.GetChannel(actChannel).Topic
}

func (bot *Mettbot) HandlerTopic(line *irc.Line) {
	actChannel := line.Args[0]
	newTopic := line.Args[1]
	oldTopic := bot.Topics[actChannel]
	bot.Topics[actChannel] = newTopic
	bot.Notice(actChannel, bot.diffTopic(oldTopic, newTopic))
}

func (bot *Mettbot) HandlerPrivmsg(line *irc.Line) {
	actChannel := line.Args[0]
	if actChannel == *Nick {
		actChannel = line.Nick
	}
	msg := line.Args[1]
	matchedTwitter, _ := regexp.MatchString(*Twitterregex, msg)

	if line.Nick == "firebird" {
		if rand.Float64() < *Firebird {
			go bot.firebird(actChannel)
		}
	}

	switch {
	case strings.Index(msg, "!") == 0:
		if rand.Float64() < *Probability {
			bot.Notice(actChannel, a.RandStr(a.IgnoreCmd))
		} else {
			bot.Command(actChannel, msg, line)
		}
	case strings.Index(msg, *Nick+":") == 0:
		bot.Mett()
		bot.Mentioned(actChannel)
	case matchedTwitter:
		bot.GetTweet(actChannel, msg)
	case strings.Contains(msg, "mett") || strings.Contains(msg, "Mett") || strings.Contains(msg, "METT"):
		bot.Mett()
	case strings.Contains(msg, "\\a"):
		bot.Dong()
	default:
		bot.MsgSinceMett++
		if bot.MsgSinceMett >= *Offmessages {
			bot.Mett()
			bot.Notice(*Channel, fmt.Sprintf(a.RandStr(a.MettTime), bot.GetMett(*Channel)))
		} else if bot.MsgSinceMett == int(float32(*Offmessages)*0.95) {
			bot.Notice(*Channel, a.RandStr(a.Warning))
		}
	}
}
