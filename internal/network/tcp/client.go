package tcp

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
	"time"
)


func Client() {
	os.Setenv("HOSTNAME", "localhost")
	c, err := net.Dial("tcp", os.Getenv("HOSTNAME")+os.Getenv("PORT"))
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println("Connected from:", c.LocalAddr().Network(), c.LocalAddr().String())
	}
	for {
		text := "HeartBeat"
		fmt.Print(">>", text)
		fmt.Fprintf(c, text+"\n")
		message, _ := bufio.NewReader(c).ReadString('\n')
		fmt.Print("\n->:\n " + message)
		fmt.Print("Sent from :", c.LocalAddr().Network(), " ", c.LocalAddr().String(), "\n")
		if strings.TrimSpace(string(text)) == "STOP" {
			fmt.Println("TCP client exiting...")
			return
		}
		time.Sleep(2 * time.Second)
	}
}

func print(){
	fmt.Println("Hello")
}
func printn(){
	fmt.Println("Hello")
}
func printnn(){
	fmt.Println("Hello")
}