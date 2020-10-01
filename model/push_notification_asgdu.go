// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See License.txt for license information.

package model

import (
	"encoding/json"
	"time"
)

// ASGDU format message
type ResponseASGDU struct {
	ErrorCode       int    `json:"errorCode"`
	ErrorMessage    string `json:"errorMessage"`
	MessageUniqueId string `json:"messageUniqueId"`
}

// ASGDU format message
type PushNotificationASGDU struct {
	From            string          `json:"from"`
	MessageID       string          `json:"messageId"`
	To              string          `json:"to"`
	Category        string          `json:"category"`
	Async           bool            `json:"async"`
	DeliveryChannel DeliveryChannel `json:"channel"`
	Timestamp       time.Time       `json:"timestamp"`
}

type DeliveryChannel struct {
	ServerChannelID         string `json:"serverChannelId"`
	DeliveryChannelUniqueID string `json:"deliveryChannelUniqueId,omitempty"`
	Data                    Data   `json:"data"`
}

type Data struct {
	Title        string           `json:"title"`
	Body         string           `json:"body"`
	ExtendedData PushNotification `json:"extendedData"`
}

func (me *PushNotificationASGDU) ToJson() string {
	b, _ := json.Marshal(me)
	return string(b)
}
