package main

type GameEvent struct {
	Game  string `json:"game"`
	Event string `json:"event"`
	Data  struct {
		Value int `json:"value"`
		Frame interface{} `json:"frame"`
	} `json:"data"`
}

type gameEventBuilder struct {
	gameEvent GameEvent
}

func BuildEvent() *gameEventBuilder {
	return &gameEventBuilder{}
}

func (b*gameEventBuilder) Get() GameEvent {
	return b.gameEvent
}

func (b*gameEventBuilder) BuildGameName(game string) *gameEventBuilder {
	b.gameEvent.Game=game
	return b
}

func (b*gameEventBuilder) BuildEventName(event string) *gameEventBuilder {
	b.gameEvent.Event=event
	return b
}

func (b*gameEventBuilder) BuildBitmap(bitmap [132][3]int, excludedEvents *[]string) *gameEventBuilder {
	bitmapFrame := struct {
		Bitmap [132][3]int `json:"bitmap"`
		ExcludedEvents *[]string `json:"excluded-events,omitempty"`
	}{Bitmap: bitmap, ExcludedEvents: excludedEvents}

	b.gameEvent.Data.Frame=bitmapFrame
	return b
}

func (b* gameEventBuilder) BuildValue(value int) *gameEventBuilder {
	b.gameEvent.Data.Value=value
	return b
}
