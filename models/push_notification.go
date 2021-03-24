package models

import (
	"strings"
	"sync"

	"github.com/appleboy/go-fcm"
	"github.com/msalihkarakasli/go-hms-push/push/model"
)

type Object map[string]interface{}

const (
	// PlatFormIos constant is 1 for iOS
	PlatFormIos = iota + 1
	// PlatFormAndroid constant is 2 for Android
	PlatFormAndroid
	// PlatFormHuawei constant is 3 for Huawei
	PlatFormHuawei
)

// PushNotification is single notification request
type PushNotification struct {
	wg *sync.WaitGroup

	// Common
	ID               string      `json:"notif_id,omitempty"`
	Tokens           []string    `json:"tokens" binding:"required"`
	Platform         int         `json:"platform" binding:"required"`
	Message          string      `json:"message,omitempty"`
	Title            string      `json:"title,omitempty"`
	Image            string      `json:"image,omitempty"`
	Priority         string      `json:"priority,omitempty"`
	ContentAvailable bool        `json:"content_available,omitempty"`
	MutableContent   bool        `json:"mutable_content,omitempty"`
	Sound            interface{} `json:"sound,omitempty"`
	Data             Object      `json:"data,omitempty"`
	Retry            int         `json:"retry,omitempty"`

	// Android
	APIKey                string            `json:"api_key,omitempty"`
	To                    string            `json:"to,omitempty"`
	CollapseKey           string            `json:"collapse_key,omitempty"`
	DelayWhileIdle        bool              `json:"delay_while_idle,omitempty"`
	TimeToLive            *uint             `json:"time_to_live,omitempty"`
	RestrictedPackageName string            `json:"restricted_package_name,omitempty"`
	DryRun                bool              `json:"dry_run,omitempty"`
	Condition             string            `json:"condition,omitempty"`
	Notification          *fcm.Notification `json:"notification,omitempty"`

	// Huawei
	APPId              string                     `json:"app_id,omitempty"`
	HuaweiNotification *model.AndroidNotification `json:"huawei_notification,omitempty"`
	HuaweiData         string                     `json:"huawei_data,omitempty"`
	HuaweiCollapseKey  int                        `json:"huawei_collapse_key,omitempty"`
	HuaweiTTL          string                     `json:"huawei_ttl,omitempty"`
	BiTag              string                     `json:"bi_tag,omitempty"`
	FastAppTarget      int                        `json:"fast_app_target,omitempty"`

	// iOS
	Expiration  *int64   `json:"expiration,omitempty"`
	ApnsID      string   `json:"apns_id,omitempty"`
	CollapseID  string   `json:"collapse_id,omitempty"`
	Topic       string   `json:"topic,omitempty"`
	PushType    string   `json:"push_type,omitempty"`
	Badge       *int     `json:"badge,omitempty"`
	Category    string   `json:"category,omitempty"`
	ThreadID    string   `json:"thread-id,omitempty"`
	URLArgs     []string `json:"url-args,omitempty"`
	Alert       Alert    `json:"alert,omitempty"`
	Production  bool     `json:"production,omitempty"`
	Development bool     `json:"development,omitempty"`
	SoundName   string   `json:"name,omitempty"`
	SoundVolume float32  `json:"volume,omitempty"`
	Apns        Object   `json:"apns,omitempty"`
}

// WaitDone decrements the WaitGroup counter.
func (p *PushNotification) WaitDone() {
	if p.wg != nil {
		p.wg.Done()
	}
}

// AddWaitCount increments the WaitGroup counter.
func (p *PushNotification) AddWaitCount() {
	if p.wg != nil {
		p.wg.Add(1)
	}
}

// IsTopic check if message format is topic for FCM
// ref: https://firebase.google.com/docs/cloud-messaging/send-message#topic-http-post-request
func (p *PushNotification) IsTopic() bool {
	if p.Platform == PlatFormAndroid {
		return p.To != "" && strings.HasPrefix(p.To, "/topics/") || p.Condition != ""
	}

	if p.Platform == PlatFormHuawei {
		return p.Topic != "" || p.Condition != ""
	}

	return false
}
