package main

import "C"
import (
	"fmt"
	"net/http"
	_ "runtime/cgo"
	"time"
	"unsafe"
)


const GameName = "TH15"
const (
	EventBackground = "BITMAP_BACK"
	EventLives = "LIVES"
	EventBombs = "BOMBS"
	EventScore = "SCORE"
)

func init() {
	fmt.Println("Okay, attached")

	client := http.Client{}
	config := ReadFile("C:\\ProgramData\\SteelSeries\\SteelSeries Engine 3\\coreProps.json")
	addr := fmt.Sprintf("http://%s",config.Address)
	gameInit(client, gameMetadata{Game: GameName,
		GameDisplayName: "Touhou 15",
		Developer: "Team Shanghai Alice",
		TimerMs: 3000,
	}, addr)

	life := GameBindMetadata{Game: GameName,Event: EventLives,MinValue: 0, MaxValue: 10}
	lifeBind := BuildGameBindEvent().BuildMetadata(life).BuildColor(
		"keyboard",
		"count",
		"function-keys",
		Color{Red: 255},
		Color{Red: 255},
		nil,
		nil).Get()

	bindGameEvent(client,lifeBind,addr)
	go registerPeriodicEvent(client,addr, func() int {
		ptr := (*C.int)(unsafe.Pointer(uintptr(0x4E745C)))
		return int(*ptr)
	},EventLives,time.Second)

	bombs := GameBindMetadata{Game: GameName,Event: EventBombs,MinValue: 0, MaxValue: 10}
	bombBind := BuildGameBindEvent().BuildMetadata(bombs).BuildColor("keyboard",
		"count",
		"number-keys",
		Color{Green: 255},
		Color{Green: 255},
		nil,
		nil).Get()

	bindGameEvent(client,bombBind,addr)
	go registerPeriodicEvent(client, addr, func() int {
		ptr := (*C.int)(unsafe.Pointer(uintptr(0x4E745C)))
		return int(*ptr)
	}, EventBombs,time.Second)

	scoreMeta := GameBindMetadata{Game: GameName, Event: EventScore}
	scoreBind := BuildGameBindEvent().BuildMetadata(scoreMeta).BuildScreenEventHandler([]ScreenData{
		{
			HasText: true,
			Prefix: "Score: ",
		},
	}).Get()
	bindGameEvent(client,scoreBind,addr)
	go registerPeriodicEvent(client,addr, func() int {
		ptr := (*C.int)(unsafe.Pointer(uintptr(0x4E740C)))
		return int(*ptr)*10
	},EventScore,time.Second/10)

	backMetadata := GameBindMetadata{Game: GameName, Event: EventBackground}
	backBind := BuildGameBindEvent().BuildMetadata(backMetadata).BuildPartialBitmap(&[]string{EventBombs,EventLives}).Get()
	bindGameEvent(client, backBind, addr)

	bitmap := [132][3]int{{}}
	for i := range bitmap{
		bitmap[i] = [3]int{50,50,0}
	}

	bitmap[ZKeyIndex] = [3]int{255,255,255}
	bitmap[XKeyIndex] = [3]int{255,255,255}
	bitmap[ShiftKeyIndex] = [3]int{255,255,255}
	bitmap[CtrlKeyIndex] = [3]int{255,255,255}
	bitmap[ArrowUpIndex] = [3]int{255,255,255}
	bitmap[ArrowLeftIndex] = [3]int{255,255,255}
	bitmap[ArrowDownIndex] = [3]int{255,255,255}
	bitmap[ArrowRightIndex] = [3]int{255,255,255}

	bitmapEvent := BuildEvent().BuildEventName(EventBackground).BuildGameName(GameName).BuildBitmap(bitmap, &[]string{EventBombs,EventLives}).Get()
	postEvent(client,addr,bitmapEvent)
	go heartbeat(client, addr, false)
}

func main()  {
	
}
