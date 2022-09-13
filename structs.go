package main

type Registrable interface {
	Registrable()
}

const (
	ZKeyIndex = 90
	XKeyIndex = 91
	ShiftKeyIndex = 88
	CtrlKeyIndex = 110

	ArrowUpIndex = 104
	ArrowLeftIndex = 125
	ArrowDownIndex = 126
	ArrowRightIndex = 127
)

type GameEventRegister struct {
	Game          string `json:"game"`
	Event         string `json:"event"`
	MinValue      int    `json:"min_value,omitempty"`
	MaxValue      int    `json:"max_value,omitempty"`
	IconId        int    `json:"icon_id,omitempty"`
	ValueOptional bool   `json:"value_optional,omitempty"`
}

func (_ GameBindMetadata) Registrable() {}

type gameMetadata struct {
	Game            string `json:"game"`
	GameDisplayName string `json:"game_display_name,omitempty"`
	Developer       string `json:"developer,omitempty"`
	TimerMs 		int    `json:"deinitialize_timer_length_ms,omitempty"`
}
type GameBindMetadata struct {
	Game     string `json:"game"`
	Event    string `json:"event"`
	MinValue int    `json:"min_value,omitempty"`
	MaxValue int    `json:"max_value,omitempty"`
	IconId   int    `json:"icon_id,omitempty"`
}

type Config struct {
	Address            string `json:"address"`
	EncryptedAddress   string `json:"encryptedAddress"`
	GgEncryptedAddress string `json:"ggEncryptedAddress"`
	MercstealthAddress string `json:"mercstealthAddress"`
}

