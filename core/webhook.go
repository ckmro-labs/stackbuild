package core

import "context"

//manage repository's webhook
type (
	// Webhook defines an integration endpoint.
	Webhook struct {
		Endpoint string `json:"endpoint,omitempty"` // like this: 'https://xxx/api/builds/5c81c0cb345204781944106e'
	}

	// WebhookData provides the webhook data.
	WebhookData struct {
		Event  string      `json:"-"`
		Action string      `json:"action"`
		User   *User       `json:"user,omitempty"`
		Repo   *Repository `json:"repo,omitempty"`
		Build  *Build      `json:"build,omitempty"`
	}

	// WebhookSender sends the webhook payload.
	WebhookSender interface {
		// Send sends the webhook to the global endpoint.
		Send(context.Context, *WebhookData) error
	}
)
