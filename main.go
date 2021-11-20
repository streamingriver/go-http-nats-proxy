package main

import (
	"fmt"
	"net/http"
	"time"

	// "errors"
	"flag"
	"io/ioutil"
	"log"
	"strconv"

	"github.com/nats-io/nats.go"
)

var (
	flagNatsServer = flag.String("nats-server", "nats://nats:4222", "")
	flagBindTo     = flag.String("bind-to", ":80", "")
)

func main() {

	conn, err := nats.Connect(*flagNatsServer)
	if err != nil {
		panic(err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		topic := r.FormValue("topic")
		timeout := r.FormValue("timeout")
		mc := r.FormValue("mc")

		if topic == "" {
			fmt.Fprintf(w, "%s", "www.streamingriveriptv.com services")
			return
		}
		if r.Method != "POST" {
			fmt.Fprintf(w, "%s", "permission denied")
			return
		}

		defer r.Body.Close()
		data, e := ioutil.ReadAll(r.Body)
		if e != nil {
			log.Printf("error: %v", e)
			fmt.Fprintf(w, "%s", e)
			return
		}

		var msg *nats.Msg
		var err error

		if len(mc) == 0 {
			if timeout != "" {
				timeoutINT, _ := strconv.Atoi(timeout)
				msg, err = conn.Request(topic, data, time.Duration(timeoutINT)*time.Second)
			} else {
				msg, err = conn.Request(topic, data, 3*time.Second)
			}
		} else {
			conn.Publish(topic, data)
			return
		}

		if err != nil {
			log.Printf("%v", err)
			fmt.Fprintf(w, "%s", err)
			return
		}
		fmt.Fprintf(w, "%s", msg.Data)
	})
	log.Printf("Starting server on "+*flagBindTo)
	log.Fatal(http.ListenAndServe(*flagBindTo, nil))
}
