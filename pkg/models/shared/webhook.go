// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

type WebhookEventType string

const (
	WebhookEventTypeGroup WebhookEventType = "group"
	WebhookEventTypeList  WebhookEventType = "list"
)

type WebhookEvent struct {
	EventGroup *EventGroup
	EventList  *EventList

	Type WebhookEventType
}

func CreateWebhookEventGroup(group EventGroup) WebhookEvent {
	typ := WebhookEventTypeGroup
	typStr := EventGroupTag(typ)
	group.DotTag = typStr

	return WebhookEvent{
		EventGroup: &group,
		Type:       typ,
	}
}

func CreateWebhookEventList(list EventList) WebhookEvent {
	typ := WebhookEventTypeList
	typStr := EventListTag(typ)
	list.DotTag = typStr

	return WebhookEvent{
		EventList: &list,
		Type:      typ,
	}
}

func (u *WebhookEvent) UnmarshalJSON(data []byte) error {
	var d *json.Decoder

	type discriminator struct {
		DotTag string
	}

	dis := new(discriminator)
	if err := json.Unmarshal(data, &dis); err != nil {
		return fmt.Errorf("could not unmarshal discriminator: %w", err)
	}

	switch dis.DotTag {
	case "group":
		d = json.NewDecoder(bytes.NewReader(data))
		eventGroup := new(EventGroup)
		if err := d.Decode(&eventGroup); err != nil {
			return fmt.Errorf("could not unmarshal expected type: %w", err)
		}

		u.EventGroup = eventGroup
		u.Type = WebhookEventTypeGroup
		return nil
	case "list":
		d = json.NewDecoder(bytes.NewReader(data))
		eventList := new(EventList)
		if err := d.Decode(&eventList); err != nil {
			return fmt.Errorf("could not unmarshal expected type: %w", err)
		}

		u.EventList = eventList
		u.Type = WebhookEventTypeList
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u WebhookEvent) MarshalJSON() ([]byte, error) {
	if u.EventGroup != nil {
		return json.Marshal(u.EventGroup)
	}

	if u.EventList != nil {
		return json.Marshal(u.EventList)
	}

	return nil, errors.New("could not marshal union type: all fields are null")

}

type Webhook struct {
	// The time at which the webhook was created
	CreatedAt *time.Time   `json:"created_at,omitempty"`
	Event     WebhookEvent `json:"event"`
	// The webhook's ID
	ID *string `json:"id,omitempty"`
	// The webhook's URL
	URL string `json:"url"`
}

func (o *Webhook) GetCreatedAt() *time.Time {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *Webhook) GetEvent() WebhookEvent {
	if o == nil {
		return WebhookEvent{}
	}
	return o.Event
}

func (o *Webhook) GetEventGroup() *EventGroup {
	return o.GetEvent().EventGroup
}

func (o *Webhook) GetEventList() *EventList {
	return o.GetEvent().EventList
}

func (o *Webhook) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *Webhook) GetURL() string {
	if o == nil {
		return ""
	}
	return o.URL
}

type WebhookInput struct {
	Event WebhookEvent `json:"event"`
	// The webhook's URL
	URL string `json:"url"`
}

func (o *WebhookInput) GetEvent() WebhookEvent {
	if o == nil {
		return WebhookEvent{}
	}
	return o.Event
}

func (o *WebhookInput) GetEventGroup() *EventGroup {
	return o.GetEvent().EventGroup
}

func (o *WebhookInput) GetEventList() *EventList {
	return o.GetEvent().EventList
}

func (o *WebhookInput) GetURL() string {
	if o == nil {
		return ""
	}
	return o.URL
}
