package slack

import (
	"encoding/json"

	"sigs.k8s.io/slack-infra/slack/blocks"
	"sigs.k8s.io/slack-infra/slack/blocks/objects"
)

type ViewWrapper struct {
	TriggerID string `json:"trigger_id"`
	View      Modal  `json:"view"`
}

type Modal struct {
	Type            modalType          `json:"type"`
	Title           objects.PlainText  `json:"title"`
	Blocks          []blocks.Block     `json:"blocks"`
	Close           *objects.PlainText `json:"close,omitempty"`
	Submit          *objects.PlainText `json:"close,omitempty"`
	PrivateMetadata string             `json:"private_metadata,omitempty"`
	CallbackID      string             `json:"callback_id,omitempty"`
	ClearOnClose    bool               `json:"clear_on_close,omitempty"`
	NotifyOnClose   bool               `json:"notify_on_close,omitempty"`
	ExternalID      string             `json:"external_id,omitempty"`
}

type modalType string

func (modalType) MarshalJSON() ([]byte, error) {
	return json.Marshal("modal")
}
