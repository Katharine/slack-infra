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

package blocks

import (
	"encoding/json"

	"sigs.k8s.io/slack-infra/slack/blocks/elements"
	"sigs.k8s.io/slack-infra/slack/blocks/objects"
)

type Block interface {
	isBlock()
}

type Section struct {
	BlockID   string                  `json:"block_id,omitempty"`
	Type      sectionBlockType        `json:"type"`
	Text      objects.Text            `json:"text,omitempty"`
	Fields    []objects.Text          `json:"fields,omitempty"`
	Accessory elements.SectionElement `json:"accessory,omitempty"`
}

func (Section) isBlock() {}

type sectionBlockType string

func (sectionBlockType) MarshalJSON() ([]byte, error) {
	return json.Marshal("section")
}

type Divider struct {
	Block
	Type dividerBlockType `json:"type"`
}

func (Divider) isBlock() {}

type dividerBlockType string

func (dividerBlockType) MarshalJSON() ([]byte, error) {
	return json.Marshal("divider")
}

type Image struct {
	BlockID  string             `json:"block_id,omitempty"`
	Type     imageBlockType     `json:"type"`
	ImageURL string             `json:"image_url"`
	AltText  string             `json:"alt_text"`
	Title    *objects.PlainText `json:"title,omitempty"`
}

func (Image) isBlock() {}

type imageBlockType string

func (imageBlockType) MarshalJSON() ([]byte, error) {
	return json.Marshal("image")
}

type Actions struct {
	BlockID  string                   `json:"block_id,omitempty"`
	Type     actionsBlockType         `json:"type"`
	Elements []elements.ActionElement `json:"elements"`
}

func (Actions) isBlock() {}

type actionsBlockType string

func (actionsBlockType) MarshalJSON() ([]byte, error) {
	return json.Marshal("actions")
}

type Context struct {
	BlockID  string                    `json:"block_id,omitempty"`
	Type     contextBlockType          `json:"type"`
	Elements []elements.ContextElement `json:"elements"`
}

func (Context) isBlock() {}

type contextBlockType string

func (contextBlockType) MarshalJSON() ([]byte, error) {
	return json.Marshal("context")
}

type Input struct {
	BlockID  string                `json:"block_id,omitempty"`
	Type     inputBlockType        `json:"type"`
	Label    objects.PlainText     `json:"label"`
	Element  elements.InputElement `json:"element"`
	Hint     *objects.PlainText    `json:"hint,omitempty"`
	Optional bool                  `json:"optional,omitempty"`
}

func (Input) isBlock() {}

type inputBlockType string

func (inputBlockType) MarshalJSON() ([]byte, error) {
	return json.Marshal("input")
}

type File struct {
	BlockID    string        `json:"block_id,omitempty"`
	Type       fileBlockType `json:"type"`
	ExternalID string        `json:"external_id"`
	Source     string        `json:"source"`
}

func (File) isBlock() {}

type fileBlockType string

func (fileBlockType) MarshalJSON() ([]byte, error) {
	return json.Marshal("file")
}
