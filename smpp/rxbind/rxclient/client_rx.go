package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/ajankovic/smpp"
	"github.com/ajankovic/smpp/pdu"
)

var (
	serverAddr string
	dstAddr    string
	srcAddr    string
	msg        string
	msgID      int
)

func main() {
	flag.StringVar(&serverAddr, "addr", "localhost:2775", "server will listen on this address.")
	flag.StringVar(&dstAddr, "dst_addr", "111111", "destination to which you are sending the message.")
	flag.StringVar(&srcAddr, "src_addr", "222222", "source from which the message is comming from.")
	flag.StringVar(&msg, "msg", "example", "contents of the message.")
	flag.Parse()

	bc := smpp.BindConf{
		Addr:     serverAddr,
		SystemID: "ExampleClient",
	}
	sc := smpp.SessionConf{}
	_, err := smpp.BindRx(sc, bc)
	if err != nil {
		fail("Can't bind: %v", err)
	}
	_ = smpp.SessionConf{

		Handler: smpp.HandlerFunc(func(ctx *smpp.Context) {
			switch ctx.CommandID() {
			case pdu.DeliverSmID:
				ctx.DeliverSm()
			}
		}),
	}

	// if err := smpp.Unbind(context.Background(), sess); err != nil {
	// 	fmt.Fprintln(os.Stderr, err.Error())
	// }

}

func fail(msg string, params ...interface{}) {
	fmt.Fprintf(os.Stderr, msg+"\n", params...)
	os.Exit(1)
}
