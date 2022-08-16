package tcp

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
	"time"
)
func Server(){
	os.Setenv("PORT",":1234")
	l, err := net.Listen("tcp",os.Getenv("PORT"))
	if err != nil {
			fmt.Println(err)
			return
	}else{
			fmt.Println("Connected to",l.Addr().Network(),l.Addr().String())
		}
	defer l.Close()
	c, err := l.Accept()
	if err != nil {
			fmt.Println(err)
			return
	}
	for {
			netData, err := bufio.NewReader(c).ReadString('\n')
			if err != nil {
					fmt.Println(err)
					return
			}
			if strings.TrimSpace(string(netData)) == "STOP" {
					fmt.Println("Exiting TCP server!")
					return
			}
			fmt.Print("-> ", string(netData))
			fmt.Print("Recieved from",c.LocalAddr())
			t := time.Now()
			myTime := t.Format(time.RFC3339) + "\n"
			c.Write([]byte(myTime))
	}
}

func printp(){
	fmt.Println("Hell")
}
func printpp(){
	fmt.Println("Hell")
}