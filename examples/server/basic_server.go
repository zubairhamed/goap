package main

import (
	. "github.com/zubairhamed/canopus"
	"github.com/zubairhamed/go-commons/network"
)

func main() {
	server := NewLocalServer()

	server.NewRoute("{obj}/{inst}/{rsrc}", GET, routeParams)

	server.NewRoute("basic", GET, routeBasic)
	server.NewRoute("basic/json", GET, routeJson)
	server.NewRoute("basic/xml", GET, routeXml)

	server.Start()
}

func routeParams(r network.Request) network.Response {
	req := r.(*CoapRequest)
	msg := NewMessageOfType(TYPE_ACKNOWLEDGEMENT, req.GetMessage().MessageId)
	msg.SetStringPayload("Acknowledged")
	res := NewResponse(msg, nil)

	return res
}

func routeBasic(r network.Request) network.Response {
	req := r.(*CoapRequest)
	msg := NewMessageOfType(TYPE_ACKNOWLEDGEMENT, req.GetMessage().MessageId)
	msg.SetStringPayload("Acknowledged")

	res := NewResponse(msg, nil)

	return res
}

func routeJson(r network.Request) network.Response {
	req := r.(*CoapRequest)
	msg := NewMessageOfType(TYPE_ACKNOWLEDGEMENT, req.GetMessage().MessageId)
	res := NewResponse(msg, nil)

	return res
}

func routeXml(r network.Request) network.Response {
	req := r.(*CoapRequest)
	msg := NewMessageOfType(TYPE_ACKNOWLEDGEMENT, req.GetMessage().MessageId)
	res := NewResponse(msg, nil)

	return res
}

/*
	// canopus.PrintMessage(msg)

	fwOpt := msg.GetOption(canopus.OPTION_PROXY_URI)
	log.Println(fwOpt)

	ack := canopus.NewMessageOfType(canopus.TYPE_ACKNOWLEDGEMENT, msg.MessageId)

	return ack

*/