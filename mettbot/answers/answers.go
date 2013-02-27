package answers

import "math/rand"

func RandStr(str []string) string {
	index := rand.Intn(len(str))
	return str[index]
}

var Syntax []string = []string{
	"Wrong Syntax. Try !help",
	"Not at all. You may try !help",
	"m( You should try the !help command",
	"Ask some people around here how to talk with me",
	"Hey, I'm a sensible MettBot. You can't talk to me that way",
}

var AddedQuote []string = []string{ // needs %v which is a number
	"Added quote #%v to database",
	"The %vth quote which will be never forgotten",
	"Another quote from the endless depths of #mett (#%v)",
	"And we've got quote number %v",
}

var AddedMett []string = []string{ // needs %v which is a number
	"Added mett #%v to database",
	"The %vth entry for those offtopic people",
	"Now I've already %v entries of mettcontent",
}

var OffendNick []string = []string{ // (ACTION) needs %v which is a string
	"slaps %v",
	"warns %v",
	"kicks %v in his mettom",
}

var QuoteNotFound []string = []string{
	"Quote not found",
	"There is no such quote",
	"The Mettlers around here aren't that funny",
}

var Warning []string = []string{
	"You guys are a bit offtopic",
	"Have ou looked at the name of this channel recently?",
	"I warn you",
	"Back to Mett",
	"You should talk more about mett",
}

var MettTime []string = []string{ // needs %v which is a string
	"It's time for moar mett: %v",
	"Hey guys, I think you will need this: %v",
	"I warned you: %v",
	"I think this channel needs some of this: %v",
}

var IgnoreCmd []string = []string{
	"Bother someone else!",
	"You're annoying me",
	"No way",
	"Ask me again later",
	"No.",
	"Talk with my non-existing hand",
}

var Mention []string = []string{
	"Yes.",
	"No.",
	"Ask _vincent",
	"May the mett be with you",
	"Hmm, what?",
}

var Firebird []string = []string{
	"Wer ist überhaupt dieser firebird?",
	"Gibt es firebird überhaupt?",
	"firebird: ist das ipv6 im HaWo schon gefixt?",
}

var Dong []string = []string{
	"DONG",
	"DÖNG",
	"DÄNG",
	"DING",
	"PLÖNK",
	"BÄM",
	"KLONK",
}
