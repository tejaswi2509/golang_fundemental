package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	"github.com/ajankovic/smpp"
	"github.com/ajankovic/smpp/pdu"
)

var (
	ServerAddr string
	dstAddr    string
	srcAddr    string
	msg        string
)

func main() {
	flag.StringVar(&ServerAddr, "addr", "localhost:2775", "server will listen on this address.")
	flag.StringVar(&dstAddr, "dst_addr", "111111", "destination to which you are sending the message.")
	flag.StringVar(&srcAddr, "src_addr", "222222", "source from which the message is comming from.")
	flag.StringVar(&msg, "msg", "example", "contents of the message.")
	flag.Parse()

	bc := smpp.BindConf{
		Addr:     ServerAddr,
		SystemID: "ExampleClient",
	}
	sc := smpp.SessionConf{}
	sess, err := smpp.BindTRx(sc, bc)
	if err != nil {
		fail("Can't bind: %v", err)
	}
	sm := &pdu.SubmitSm{
		SourceAddr:      srcAddr,
		DestinationAddr: dstAddr,
		ShortMessage:    msg,
	}
	resp, err := sess.Send(context.Background(), sm)
	if err != nil {
		fail("Can't send message: %+v", err)
	}
	em := &pdu.EnquireLink{}
	eresp, er := sess.Send(context.Background(), em)
	if er != nil {
		fail("can't send enquire: %+v", er)
	}
	fmt.Fprintf(os.Stderr, "Message sent\n")
	fmt.Fprintf(os.Stderr, "Received response %s %+v\n", resp.CommandID(), resp)
	fmt.Fprintf(os.Stderr, "enquired response %s %+v\n", eresp.CommandID(), eresp)
	if err := smpp.Unbind(context.Background(), sess); err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
	}
}

func fail(msg string, params ...interface{}) {
	fmt.Fprintf(os.Stderr, msg+"\n", params...)
	os.Exit(1)
}
