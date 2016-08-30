package main

import (
	"os"

	"github.com/rancher/my-subcriber/handlers"

	log "github.com/Sirupsen/logrus"
	"github.com/rancher/event-subscriber/events"
)

func main() {
	log.Info("test start")
	apiURL := os.Getenv("CATTLE_URL")
	accessKey := os.Getenv("CATTLE_ACCESS_KEY")
	secretKey := os.Getenv("CATTLE_SECRET_KEY")

	eventChan := make(chan bool)
	done := make(chan error)
	go func() {
		eventHandlers := map[string]events.EventHandler{
			"resource.change": handlers.TestHandler,
		}
		router, err := events.NewEventRouter("", 2000, apiURL, accessKey, secretKey,
			nil, eventHandlers, "resource", 10, events.DefaultPingConfig)
		if err == nil {
			err = router.Start(eventChan)
		}
		done <- err
	}()
	go func() {
		log.Infof("my subcriber starting")
		<-eventChan
	}()
	err := <-done
	if err == nil {
		log.Infof("exiting")
	}
}
