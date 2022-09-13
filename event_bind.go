package main

type GameBindEvent struct {
	GameBindMetadata
	Handlers []interface{} `json:"handlers"`
}

type gameBindEventBuilder struct {
	gameBindEvent GameBindEvent
}

func BuildGameBindEvent() *gameBindEventBuilder {
	return &gameBindEventBuilder{}
}

func (bb* gameBindEventBuilder) Get() GameBindEvent {
	return bb.gameBindEvent
}

func (bb* gameBindEventBuilder) BuildMetadata(metadata GameBindMetadata) *gameBindEventBuilder {
	bb.gameBindEvent.GameBindMetadata=metadata
	return bb
}

func (bb* gameBindEventBuilder) BuildColor(deviceType , mode, zone string, colorZero, colorHundred Color, frequency, repeatLimit *[]Frequency) *gameBindEventBuilder {
	type colorHandlerStruct struct {
		DeviceType string `json:"device-type"`
		Zone       string `json:"zone"`
		Mode       string `json:"mode"`
		Color      struct {
			Gradient struct {
				Zero Color `json:"zero"`
				Hundred Color `json:"hundred"`
			} `json:"gradient"`
		} `json:"color"`
		Rate struct {
			Frequency *[]Frequency `json:"frequency,omitempty"`
			RepeatLimit *[]Frequency `json:"repeat_limit,omitempty"`
		} `json:"rate"`
	}

	handler := colorHandlerStruct{}
	handler.Zone=zone
	handler.Mode=mode
	handler.DeviceType=deviceType
	handler.Rate.RepeatLimit=repeatLimit
	handler.Rate.Frequency=frequency
	handler.Color.Gradient.Hundred=colorHundred
	handler.Color.Gradient.Zero=colorZero

	bb.gameBindEvent.Handlers = append(bb.gameBindEvent.Handlers, handler)
	return bb
}

func (bb* gameBindEventBuilder) BuildPartialBitmap(strings *[]string) *gameBindEventBuilder {
	type partialBitmapHandler struct {
		DeviceType     string   `json:"device-type"`
		Mode           string   `json:"mode"`
		ExcludedEvents *[]string `json:"excluded-events,omitempty"`
	}
	
	handler := partialBitmapHandler{
		DeviceType:    "rgb-per-key-zones",
		Mode:           "partial-bitmap",
		ExcludedEvents: strings,
	}

	bb.gameBindEvent.Handlers = append(bb.gameBindEvent.Handlers, handler)
	return bb
}



type Frequency struct {
	Low       int `json:"low"`
	High      int `json:"high"`
	Frequency int `json:"frequency"`
}

type Color struct {
	Red   int `json:"red"`
	Green int `json:"green"`
	Blue  int `json:"blue"`
}

func (bb* gameBindEventBuilder) BuildScreenEventHandler(datas []ScreenData)  *gameBindEventBuilder{
	type screenEventBindHandler struct {
		DeviceType string `json:"device-type"`
		Mode       string `json:"mode"`
		Zone       string `json:"zone"`
		Datas []ScreenData`json:"datas"`
	}
	
	handler := screenEventBindHandler{
		DeviceType: "screened",
		Mode:       "screen",
		Zone:       "one",
		Datas: datas,
	}

	bb.gameBindEvent.Handlers = append(bb.gameBindEvent.Handlers, handler)
	return bb
}

type ScreenData struct {
	HasText         bool   `json:"has-text"`
	ContextFrameKey string `json:"context-frame-key,omitempty"`
	Prefix       string `json:"prefix,omitempty"`
	Suffix       string `json:"suffix,omitempty"`
	LengthMillis int    `json:"length-millis,omitempty"`
	Arg          string `json:"arg,omitempty"`
	IconId       int    `json:"icon-id,omitempty"`
}

