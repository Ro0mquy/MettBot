package plugins

import (
	"../ircclient"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"strconv"
	"strings"
	"time"
)

const (
	hawo_address  = "hawo.net:7337"
	izibi_address = "http://tempi.0x4a42.net:7337/"
	matti_address = "http://temp.notmatti.me"
)

type TemperaturPlugin struct {
	ic *ircclient.IRCClient
}

func (q *TemperaturPlugin) String() string {
	return "temperatur"
}

func (q *TemperaturPlugin) Info() string {
	return "Prints information gathered by temperatur sensors"
}

func (q *TemperaturPlugin) Usage(cmd string) string {
	switch cmd {
	case "ht":
		return "prints the temperatur at HaWo"
	case "it":
		return "prints the temperatur at izibis home"
	case "mt":
		return "prints the temperatur at mattis home"
	}
	return ""
}

func (q *TemperaturPlugin) Register(ic *ircclient.IRCClient) {
	q.ic = ic
	q.ic.RegisterCommandHandler("ht", 0, 0, q)
	q.ic.RegisterCommandHandler("it", 0, 0, q)
	q.ic.RegisterCommandHandler("mt", 0, 0, q)
}

func (q *TemperaturPlugin) Unregister() {
	// nothing to do here
}

func (q *TemperaturPlugin) ProcessLine(msg *ircclient.IRCMessage) {
	return
}

func (q *TemperaturPlugin) ProcessCommand(cmd *ircclient.IRCCommand) {
	switch cmd.Command {
	case "ht":
		q.ic.Reply(cmd, q.getHaWo())
	case "it":
		q.ic.Reply(cmd, q.getIzibi())
	case "mt":
		q.ic.Reply(cmd, q.getMatti())
	}
}

// does a dial to <address> and returns the answer as a string
func (q *TemperaturPlugin) getTCP(address string) (answer string, err error) {
	conn, err := net.DialTimeout("tcp", address, 10*time.Second)
	if err != nil {
		return
	}
	defer conn.Close()
	buffer, err := ioutil.ReadAll(conn)
	answer = string(buffer)
	return
}

// does a http request to <address> and returns the answer as a string
func (q *TemperaturPlugin) getHttp(address string) (answer string, err error) {
	resp, err := http.Get(address)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	answer = string(body)
	return
}

func (q *TemperaturPlugin) getHaWo() string {
	answer, err := q.getTCP(hawo_address)
	if err != nil {
		return "Error accessing \"" + hawo_address + "\": " + err.Error()
	}
	return "Aktuelle Temperatur am HaWo: " + answer
}

func (q *TemperaturPlugin) getIzibi() string {
	answer, err := q.getHttp(izibi_address + "aussen")
	if err != nil {
		return "Error accessing \"" + izibi_address + "aussen\": " + err.Error()
	}
	aussen, err := strconv.ParseFloat(strings.TrimSuffix(answer, "\n"), 64)
	if err != nil {
		return "Error accessing \"" + izibi_address + "aussen\": " + err.Error()
	}

	answer, err = q.getHttp(izibi_address + "innen")
	if err != nil {
		return "Error accessing \"" + izibi_address + "innen\": " + err.Error()
	}
	innen, err := strconv.ParseFloat(strings.TrimSuffix(answer, "\n"), 64)
	if err != nil {
		return "Error accessing \"" + izibi_address + "innen\": " + err.Error()
	}

	return fmt.Sprint("Aktuelle Temperatur vor izibis Fenster: ", aussen, "°C und in seinem Zimmer: ", innen, "°C")
}

func (q *TemperaturPlugin) getMatti() string {
	answer, err := q.getHttp(matti_address)
	if err != nil {
		return "Error accessing \"" + matti_address + "\": " + err.Error()
	}
	return "Aktuelle Temperatur bei Matti: " + answer
}
