package plugins

import (
	"../ircclient"
	"strings"
)

type BadewannePlugin struct {
	ic *ircclient.IRCClient
}

func (q *BadewannePlugin) String() string {
	return "badewanne"
}

func (q *BadewannePlugin) Info() string {
	return `sends back a smiling badewanne for every \badewanne`
}

func (q *BadewannePlugin) Usage(cmd string) string {
	// just for interface saturation
	return ""
}

func (q *BadewannePlugin) Register(cl *ircclient.IRCClient) {
	q.ic = cl
}

func (q *BadewannePlugin) Unregister() {
	return
}

func (q *BadewannePlugin) ProcessLine(msg *ircclient.IRCMessage) {
	if msg.Command != "PRIVMSG" {
		//only process messages from chatrooms
		return
	}

	count := strings.Count(msg.Args[0], `\badewanne`)
	count -= strings.Count(msg.Args[0], `\\badewanne`)
	count += strings.Count(msg.Args[0], `\lb`)
	count -= strings.Count(msg.Args[0], `\\lb`)
	if count < 1 {
		return
	}

	str := "¯\\_(ツ)_/¯"
	message := ""

	for i := 0; i < count; i++ {
		message += str + " "
	}

	q.ic.ReplyMsg(msg, message)
}

func (q *BadewannePlugin) ProcessCommand(cmd *ircclient.IRCCommand) {
	return
}
