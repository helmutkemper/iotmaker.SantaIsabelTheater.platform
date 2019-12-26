package fps

import (
	"fmt"
	"math/rand"
	"time"
)

type funcRunner struct {
	Func  func()
	Async bool
}

const kUIdSize = 10

var fps = 120
var kUIdCharList []string

// en: Warning! stopTicker should be a channel, however, conflict with webassembly <-done channer
// pt_br: Cuidado! stopTicker deveria ser um channel, porém, deu conflito com o webassembly <-done channer
var stopTicker bool
var ticker *time.Ticker
var funcListToRunner map[string]funcRunner

func Set(value int) {
	fps = value
	stopTicker = true
}

func Get() int {
	return fps
}

func AddRunner(runnerFunc func(), async bool) string {
	UId := getUId()
	funcListToRunner[UId] = funcRunner{
		Func:  runnerFunc,
		Async: async,
	}

	return UId
}

func DeleteRunner(UId string) {
	delete(funcListToRunner, UId)
}

func init() {
	kUIdCharList = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s",
		"t", "u", "v", "w", "x", "y", "z", "A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P",
		"Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z", "0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "_", "!", "@",
		"#", "$", "%", "&", "*", "(", ")", "-", "_", "+", "=", "[", "{", "}", "]", "/", "?", ";", ":", ".", ",", "<", ">",
		"|"}
	funcListToRunner = make(map[string]funcRunner)
	tickerStart()
}

func getUId() string {
	var UId = ""
	for i := 0; i != kUIdSize; i += 1 {
		UId += kUIdCharList[rand.Intn(len(kUIdCharList)-1)]
	}

	return UId
}

func tickerStart() {
	ticker = time.NewTicker(time.Second / time.Duration(fps))
	go func() { tickerRunner() }()
}

func tickerRunner() {
	for {
		select {
		case <-ticker.C:
			if stopTicker == true {
				stopTicker = false
				fmt.Printf("fps done entrou\n")
				func() { tickerStart() }()
				return
			}

			for _, runnerFunc := range funcListToRunner {
				if runnerFunc.Async == false && runnerFunc.Func != nil {
					runnerFunc.Func()
				} else if runnerFunc.Async == true && runnerFunc.Func != nil {
					go runnerFunc.Func()
				}
			}
		}
	}
}
