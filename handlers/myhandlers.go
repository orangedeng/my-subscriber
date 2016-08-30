package handlers

import (
	log "github.com/Sirupsen/logrus"
	"github.com/rancher/event-subscriber/events"
	"github.com/rancher/go-rancher/client"
)

// TestHandler first handler
func TestHandler(event *events.Event, apiClient *client.RancherClient) error {
	log.WithFields(log.Fields{
		"resourceId": event.ResourceID,
		"eventId":    event.ID,
		"name":       event.Name,
	}).Info("Event")

	return nil
}
