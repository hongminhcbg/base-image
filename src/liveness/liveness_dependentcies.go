package liveness

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/urfave/cli/v3"
)

func LivenessDependencies(ctx *cli.Context) error {
	urlsRaw := os.Getenv("HEALTH_DEPENDENCY")
	log.Println("startServiceWithConfig", "urlRaws", urlsRaw)
	urls := strings.Split(urlsRaw, ",")

	started := time.Now()
	log.Println("starting server with port 8080")
	http.HandleFunc("/started", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		data := (time.Since(started)).String()
		w.Write([]byte(data))
	})

	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		log.Println("pong")
		w.WriteHeader(200)
		w.Write([]byte(`{"message":"pong"}`))
	})

	http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		duration := time.Since(started)
		if duration.Seconds() < 10 {
			w.WriteHeader(500)
			w.Write([]byte(fmt.Sprintf("error: %v", duration.Seconds())))
			return
		}

		for _, url := range urls {
			if len(strings.TrimSpace(url)) == 0 {
				continue
			}

			log.Println("startPingToService: ", url)
			resp, err := http.Get(url)
			if err != nil || resp.StatusCode != http.StatusOK {
				log.Println("pingError: ", url, "error", err, "resp", resp)
				w.WriteHeader(500)
				w.Write([]byte(fmt.Sprintf("error: %v, health with resp %v", err, resp)))
				return
			}

			log.Println("pingToServiceSuccess", "url", url)
		}

		w.WriteHeader(200)
		w.Write([]byte("ok"))
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
	return nil
}
