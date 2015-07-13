package main
import (
	. "github.com/zubairhamed/canopus"
	"github.com/zubairhamed/go-commons/network"
	"time"
	"log"
)

func main() {
	server := NewLocalServer()
	server.NewRoute("watch/this", GET, routeHandler)

	GenerateRandomChangeNotifications(server)

	server.On(EVT_OBSERVE, func(){
		log.Println("Observe requested")
	})

	server.Start()
}

func GenerateRandomChangeNotifications(server *CoapServer) {
	ticker := time.NewTicker(10 * time.Second)
	go func() {
		for {
			select {
			case <-ticker.C:
				log.Println("Notify Change..")
				server.NotifyChange("watch/this", "Some new value")
			}
		}
	}()
}

func routeHandler(r network.Request) network.Response {
	req := r.(*CoapRequest)
	msg := NewMessageOfType(TYPE_ACKNOWLEDGEMENT, req.GetMessage().MessageId)
	msg.SetStringPayload("Acknowledged")
	res := NewResponse(msg, nil)

	return res
}