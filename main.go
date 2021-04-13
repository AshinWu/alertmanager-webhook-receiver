package main

import (
	"encoding/json"
	"flag"
	"log"
	"net/http"
)

type (
	// HookMessage is a JSON body from the alertmanager send HTTP POST requests.
	// Learn more to see: #https://prometheus.io/docs/alerting/latest/configuration/#webhook_config
	HookMessage struct {
		Version           string            `json:"version"`
		GroupKey          string            `json:"groupKey"`
		Status            string            `json:"status"`
		Receiver          string            `json:"receiver"`
		GroupLabels       map[string]string `json:"groupLabels"`
		CommonLabels      map[string]string `json:"commonLabels"`
		CommonAnnotations map[string]string  `json:"commonAnnotations"`
		ExternalURL       string            `json:"externalURL"`
		Alerts            []Alert           `json:"alerts"`
	}

	// Alert is a single alert.
	Alert struct {
		Status       string            `json:"status"`
		Labels       map[string]string `json:"labels"`
		Annotations  map[string]string `json:"annotations"`
		StartsAt     string            `json:"startsAt,omitempty"`
		EndsAt       string            `json:"EndsAt,omitempty"`
		GeneratorURL string            `json:"generatorURL"`
	}
)

func main() {
	addr := flag.String("addr", ":8090", "address to listen for webhook")
	flag.Parse()

	http.HandleFunc("/alerts", alertsHandler)
	http.HandleFunc("/health", healthHandler)

	log.Printf("start server address=%s", *addr)
	log.Fatal(http.ListenAndServe(*addr, nil))
}

func healthHandler(w http.ResponseWriter, r *http.Request)  {
	log.Printf("ping from %v", r.Host)
	w.Write([]byte("pong"))
}

func alertsHandler(w http.ResponseWriter, r *http.Request) {
	dec := json.NewDecoder(r.Body)
	defer r.Body.Close()

	var m HookMessage
	if err := dec.Decode(&m); err != nil {
		log.Printf("error decoding message: %v", err)
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}
	// do something here
	log.Printf("[Body]：%#v\n", &m)

	w.WriteHeader(http.StatusNoContent)
}

