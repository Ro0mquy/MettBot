package main

import (
	"./ircclient"
	"./plugins"
	"log"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	log.SetFlags(log.Lshortfile)

	s := ircclient.NewIRCClient("mettbot.cfg")
	s.RegisterPlugin(new(plugins.KexecPlugin))
	s.RegisterPlugin(new(plugins.ListPlugins))
	s.RegisterPlugin(new(plugins.LoggerPlugin))
	s.RegisterPlugin(new(plugins.QuitHandler))
	s.RegisterPlugin(new(plugins.ChannelsPlugin))
	s.RegisterPlugin(new(plugins.AdminPlugin))
	s.RegisterPlugin(new(plugins.TwitterPlugin))
	s.RegisterPlugin(new(plugins.DongPlugin))
	s.RegisterPlugin(new(plugins.BadewannePlugin))
	s.RegisterPlugin(new(plugins.TopicDiffPlugin))
	s.RegisterPlugin(new(plugins.MumblePlugin))
	s.RegisterPlugin(new(plugins.QuoteDBPlugin))
	s.RegisterPlugin(new(plugins.MettDBPlugin))
	s.RegisterPlugin(new(plugins.XKCDPlugin))
	//s.RegisterPlugin(new(plugins.AltPlugin))
	s.RegisterPlugin(new(plugins.TemperaturPlugin))
	//s.RegisterPlugin(new(plugins.CorrectionPlugin))

	err := s.Connect()
	if err != nil {
		log.Fatal(err.Error())
	}

	err = s.InputLoop()
	if err != nil {
		log.Fatal(err.Error())
	}
}
