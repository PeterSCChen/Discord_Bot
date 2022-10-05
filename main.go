package main

import (
	"fmt"
	"os"
	"io/ioutil"
	"os/signal"
	"syscall"
	"strings"
	"net/http"
	"log"

	//"github.com/KnutZuidema/golio"
  //"github.com/KnutZuidema/golio/api"
	"github.com/bwmarrin/discordgo"
)

// Variables used for command line parameters

var Discord_Token

func main() {

	// section for discord bot
	dg, err := discordgo.New("Bot " + Discord_Token)
	if err != nil {
		fmt.Println("Error in creating discord bot session!")
		return
	}

	dg.AddHandler(messageCreate)

	dg.Identify.Intents = discordgo.IntentsGuildMessages

	err = dg.Open()
	if err != nil {
		fmt.Println("Error in opening discord bot connection!")
		return
	}

	fmt.Println("Now running. Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	dg.Close()

	// session for http get request
	resp, err := http.Get("Https://jsonplaceholder.typicode.com/posts")
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	sb := string(body)
	log.Printf(sb)
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate){

	if m.Author.ID == s.State.User.ID {
		return
	}

	if m.Content == "ping pong" {
		s.ChannelMessageSend(m.ChannelID, "Pong!")
	}
	sp := strings.Split(m.Content, " ")

	if(len(sp) < 2){
		fmt.Println("Not length 2 ")
	} else if sp[0] != "!live"{
		fmt.Println("not !live")
	}else{
		var id = sp[1]
		s.ChannelMessageSend(m.ChannelID, id)
	}

}
