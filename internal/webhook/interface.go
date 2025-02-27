package webhook

// src/NzbDrone.Core/Notifications/Webhook/WebhookPayload.cs
type Event struct {
	EventType      EventType `json:"eventType"`
	InstanceName   string    `json:"instanceName"`
	ApplicationUrl string    `json:"applicationUrl"`
}

// src/NzbDrone.Core/Notifications/Webhook/WebhookEventType.cs
type EventType string

const (
	Test                      EventType = "Test"
	Grab                      EventType = "Grab"
	Download                  EventType = "Download"
	Rename                    EventType = "Rename"
	SeriesAdd                 EventType = "SeriesAdd"
	SeriesDelete              EventType = "SeriesDelete"
	EpisodeFileDelete         EventType = "EpisodeFileDelete"
	Health                    EventType = "Health"
	ApplicationUpdate         EventType = "ApplicationUpdate"
	HealthRestored            EventType = "HealthRestored"
	ManualInteractionRequired EventType = "ManualInteractionRequired"
	MovieDelete               EventType = "MovieDelete"
	MovieFileDelete           EventType = "MovieFileDelete"
	MovieAdded                EventType = "MovieAdded"
)
