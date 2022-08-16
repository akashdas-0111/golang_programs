package main
import "go-tutorails/internal/network/tcp"
func main() {
  go tcp.Server()
  tcp.Client()
}

