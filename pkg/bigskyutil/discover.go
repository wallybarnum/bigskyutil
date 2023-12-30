package bigskyutil

import (
	"gitlab.com/gomidi/midi/v2"
	_ "gitlab.com/gomidi/midi/v2/drivers/rtmididrv"
)


func Discover()(string, string) {
	defer midi.CloseDriver()
	inports := midi.GetInPorts().String()
	outports := midi.GetOutPorts().String()
	return inports, outports
}
