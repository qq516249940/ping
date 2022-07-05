package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"time"

	"github.com/tatsushid/go-fastping"
	wxworkbot "github.com/vimsucks/wxwork-bot-go"
)

func ping(x string) {
	p := fastping.NewPinger()
	ra, err := net.ResolveIPAddr("ip4:icmp", x)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	p.AddIPAddr(ra)
	p.OnRecv = func(addr *net.IPAddr, rtt time.Duration) {
		fmt.Printf("IP Addr: %s receive, RTT: %v\n", addr.String(), rtt)
		bot()
	}
	p.OnIdle = func() {
		fmt.Println("finish")

	}
	err = p.Run()
	if err != nil {
		fmt.Println(err)
	}

}

func bot() {
	bot := wxworkbot.New(os.Args[2])
	// or Markdown, Image, News

	// 文本消息
	// text := wxworkbot.Text{
	// 	Content:             "Hello World",
	// 	MentionedList:       []string{"foo;", "bar"},
	// 	MentionedMobileList: []string{"@all"},
	// }
	// err := bot.Send(text)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// // Markdown 消息
	content := "# " + os.Args[1] + "ping通了"

	markdown := wxworkbot.Markdown{
		Content: content,
	}
	err := bot.Send(markdown)
	if err != nil {
		log.Fatal(err)
	}

}

func main() {
	ping(os.Args[1])

}
