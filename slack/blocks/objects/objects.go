/*
Copyright 2020 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package objects

import "encoding/json"

type Text interface {
	isTextObject()
}

type PlainText struct {
	Type  plainTextType `json:"type"`
	Text  string        `json:"text"`
	Emoji bool          `json:"emoji,omitempty"`
}

func (PlainText) isTextObject() {}

type plainTextType string

func (plainTextType) MarshalJSON() ([]byte, error) {
	return json.Marshal("plain_text")
}

type Mrkdwn struct {
	Type     mrkdwnType `json:"type"`
	Text     string     `json:"text"`
	Verbatim bool       `json:"verbatim,omitempty"`
}

func (Mrkdwn) isTextObject() {}

type mrkdwnType string

func (mrkdwnType) MarshalJSON() ([]byte, error) {
	return json.Marshal("mrkdwn")
}

type ConfirmationDialog struct {
	Title   PlainText `json:"title"`
	Text    Text      `json:"text"`
	Confirm PlainText `json:"confirm"`
	Deny    PlainText `json:"deny"`
}

type Option struct {
	Text        PlainText `json:"text"`
	Value       string    `json:"value"`
	Description PlainText `json:"description"`
	URL         string    `json:"url"`
}

type OptionGroup struct {
	Label   PlainText `json:"label"`
	Options []Option  `json:"options"`
}
