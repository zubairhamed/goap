package goap

import (
    "net"
    "time"
    "log"
)

/*
SendInternalServerError()
SendAcknowledge()
SendConfirmable()
SendNonConfirmable()
SendContent()
*/

func SendError402BadOption(messageId uint16, conn *net.UDPConn, addr *net.UDPAddr) {
    msg := NewMessage(TYPE_NONCONFIRMABLE, COAPCODE_501_NOT_IMPLEMENTED, messageId)
    msg.SetStringPayload("Bad Option: An unknown option of type critical was encountered")

    SendMessageTo(msg, conn, addr)
}

// Sends a CoAP Message to a UDP Connection


// Sends a CoAP Message to UDP address
func SendMessageTo(msg *Message, conn *net.UDPConn, addr *net.UDPAddr) (*CoapResponse, error) {
    b, _ := MessageToBytes(msg)
    _, err := conn.WriteToUDP(b, addr)

    if err != nil {
        log.Println(err)

        return nil, err
    }

    if msg.MessageType == TYPE_NONCONFIRMABLE {
        return nil, err
    } else {
        var buf []byte = make([]byte, 1500)
        // conn.SetReadDeadline(time.Now().Add(time.Second * 2))
        n, _, err := conn.ReadFromUDP(buf)

        if err != nil {
            return nil, err
        }

        msg, err := BytesToMessage(buf[:n])

        resp := NewResponse(msg, err)

        return resp, err
    }
    return nil, nil
}

func SendMessage(msg *Message, conn *net.UDPConn) (*CoapResponse, error) {
    b, _ := MessageToBytes(msg)
    _, err := conn.Write(b)

    if err != nil {
        log.Println(err)

        return nil, err
    }

    if msg.MessageType == TYPE_NONCONFIRMABLE {
        return nil, err
    } else {
        var buf []byte = make([]byte, 1500)
        conn.SetReadDeadline(time.Now().Add(time.Second * 2))
        n, _, err := conn.ReadFromUDP(buf)

        if err != nil {
            return nil, err
        }

        msg, err := BytesToMessage(buf[:n])

        resp := NewResponse(msg, err)

        return resp, err
    }
}