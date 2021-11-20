package main

import (
	"fmt"
	"net/http"
	"time"

	"io/ioutil"
	"log"
	"strconv"

	"github.com/nats-io/nats.go"
	"gitlab.com/avarf/getenvs"
)

func main() {

	conn, err := nats.Connect(getenvs.GetEnvString("NATS_SERVER", "nats://nats:4222"))
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

		if len(mc) > 0 {
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
	log.Printf("Starting server on " + getenvs.GetEnvString("NATS_PROXY_PORT", ":80"))
	log.Fatal(http.ListenAndServe(getenvs.GetEnvString("NATS_PROXY_PORT", ":80"), nil))
}
