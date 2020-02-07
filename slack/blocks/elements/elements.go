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

package elements

import (
	"encoding/json"

	"sigs.k8s.io/slack-infra/slack/blocks/objects"
)

type Element interface {
	isElement()
}

type SectionElement interface {
	isSectionElement()
}

type ActionElement interface {
	isActionElement()
}

type InputElement interface {
	isInputElement()
}

type ContextElement interface {
	isContextElement()
}

type Button struct {
	Type     buttonType                  `json:"type"`
	Text     objects.PlainText           `json:"text"`
	ActionID string                      `json:"action_id"`
	URL      string                      `json:"url,omitempty"`
	Value    string                      `json:"value,omitempty"`
	Style    ButtonStyle                 `json:"style,omitempty"`
	Confirm  *objects.ConfirmationDialog `json:"confirm,omitempty"`
}

func (Button) isElement()        {}
func (Button) isSectionElement() {}
func (Button) isActionElement()  {}

type ButtonStyle string

const (
	ButtonStylePrimary ButtonStyle = "primary"
	ButtonStyleDanger  ButtonStyle = "danger"
	ButtonStyleDefault ButtonStyle = ""
)

type buttonType string

func (buttonType) MarshalJSON() ([]byte, error) {
	return json.Marshal("button")
}

type CheckboxGroup struct {
	Type           checkboxGroupType           `json:"type"`
	ActionID       string                      `json:"action_id"`
	Options        []objects.Option            `json:"options"`
	InitialOptions []objects.Option            `json:"initial_options,omitempty"`
	Confirm        *objects.ConfirmationDialog `json:"confirm,omitempty"`
}

func (CheckboxGroup) isElement()        {}
func (CheckboxGroup) isActionElement()  {}
func (CheckboxGroup) isSectionElement() {}
func (CheckboxGroup) isInputElement()   {}

type checkboxGroupType string

func (checkboxGroupType) MarshalJSON() ([]byte, error) {
	return json.Marshal("checkboxes")
}

type DatePicker struct {
	Type        datepickerType              `json:"type"`
	ActionID    string                      `json:"action_id"`
	Placeholder *objects.PlainText          `json:"placeholder,omitempty"`
	InitialDate string                      `json:"initial_date"`
	Confirm     *objects.ConfirmationDialog `json:"confirm"`
}

func (DatePicker) isElement()        {}
func (DatePicker) isActionElement()  {}
func (DatePicker) isSectionElement() {}
func (DatePicker) isInputElement()   {}

type datepickerType string

func (datepickerType) MarshalJSON() ([]byte, error) {
	return json.Marshal("datepicker")
}

type Image struct {
	Type     imageType `json:"image"`
	ImageURL string    `json:"image_url"`
	AltText  string    `json:"alt_text"`
}

func (Image) isElement()        {}
func (Image) isSectionElement() {}
func (Image) isContextElement() {}

type imageType string

func (imageType) MarshalJSON() ([]byte, error) {
	return json.Marshal("image")
}

type Overflow struct {
	Type     overflowType                `json:"type"`
	ActionID string                      `json:"action_id"`
	Options  []objects.Option            `json:"options"`
	Confirm  *objects.ConfirmationDialog `json:"confirm,omitempty"`
}

func (Overflow) isElement()        {}
func (Overflow) isSectionElement() {}
func (Overflow) isActionElement()  {}

type overflowType string

func (overflowType) MarshalJSON() ([]byte, error) {
	return json.Marshal("overflow")
}

type PlainTextInput struct {
	Type         plainTextInputType `json:"type"`
	ActionID     string             `json:"action_id"`
	Placeholder  *objects.PlainText `json:"placeholder,omitempty"`
	InitialValue string             `json:"initial_value,omitempty"`
	Multiline    bool               `json:"multiline,omitempty"`
	MinLength    int                `json:"min_length,omitempty"`
	MaxLength    int                `json:"max_length,omitempty"`
}

func (PlainTextInput) isElement()        {}
func (PlainTextInput) isSectionElement() {}
func (PlainTextInput) isActionElement()  {}
func (PlainTextInput) isInputElement()   {}

type plainTextInputType string

func (plainTextInputType) MarshalJSON() ([]byte, error) {
	return json.Marshal("plain_text_input")
}

type RadioButtonGroup struct {
	Type          radioButtonGroupType        `json:"type"`
	ActionID      string                      `json:"action_id"`
	Options       []objects.Option            `json:"options"`
	InitialOption *objects.Option             `json:"initial_option,omitempty"`
	Confirm       *objects.ConfirmationDialog `json:"confirm,omitempty"`
}

func (RadioButtonGroup) isElement()        {}
func (RadioButtonGroup) isSectionElement() {}
func (RadioButtonGroup) isActionElement()  {}
func (RadioButtonGroup) isInputElement()   {}

type radioButtonGroupType string

func (radioButtonGroupType) MarshalJSON() ([]byte, error) {
	return json.Marshal("radio_buttons")
}

type StaticSelectMenu struct {
	Type          staticSelectMenuType        `json:"type'"`
	Placeholder   objects.PlainText           `json:"placeholder"`
	ActionID      string                      `json:"action_id"`
	Options       []objects.Option            `json:"options,omitempty"`
	OptionGroups  []objects.OptionGroup       `json:"option_groups,omitempty"`
	InitialOption *objects.Option             `json:"initial_option,omitempty"`
	Confirm       *objects.ConfirmationDialog `json:"confirm,omitempty"`
}

func (StaticSelectMenu) isElement()        {}
func (StaticSelectMenu) isSectionElement() {}
func (StaticSelectMenu) isActionElement()  {}
func (StaticSelectMenu) isInputElement()   {}

type staticSelectMenuType string

func (staticSelectMenuType) MarshalJSON() ([]byte, error) {
	return json.Marshal("static_select")
}

// Many other types of select menu have been omitted for the time being.
