package bigskyutil

import (
	"encoding/json"
	"log"
	"time"

	"gitlab.com/gomidi/midi/v2"
	_ "gitlab.com/gomidi/midi/v2/drivers/rtmididrv"
)

func ReadFile(directory string) ([]byte, error) {

	defer midi.CloseDriver()

	in, err := midi.FindInPort("BigSkyMX")
	if err != nil {
		log.Println("can't find BigSkyMX")
		return nil, err
	}
	// listens to the in port and calls eachMessage for every message.
	// any running status bytes are converted and only complete messages are passed to the eachMessage.
	stop, err := midi.ListenTo(in, EachMessage, midi.UseSysEx())
	if err != nil {
		log.Println("can't listen to in port")
		return nil, err
	}

	var out, _ = midi.FindOutPort("BigSkyMX")
	if err != nil {
		log.Println("can't find BigSkyMX")
		return nil, err
	}

	// creates a sender function to the out port
	send, err := midi.SendTo(out)
	if err != nil {
		log.Println("can't make send function")
		return nil, err
	}

	Rxbytes = Rxbytes[:0] // reset rx slice - TODO: use channel instead

	// make a dir request sysex message
	m, err := ReadFileRequest(directory)
	if err != nil {
		log.Println("can't build dir request")
		return nil, err
	}
	err = send(m) // send the read request sysex message
	if err != nil {
		log.Println("can't send sysex message err = ", err)
		return nil, err
	}
	// wait for a response
	for Eof == false {
		if Sig == false {
			//log.Println("waiting for response")
			time.Sleep(time.Millisecond * 200)
		} else {
			Sig = false
			// make a data block request sysex message
			m, err = DataBlockRequest(256)
			err = send(m) // send the datablock sysex message
			if err != nil {
				log.Println("can't send data block requeste sysex message err = ", err)
				return nil, err
			}
		}
	}
	// stops listening
	stop()
	time.Sleep(time.Millisecond * 100)
	return Rxbytes, nil
}


func ReadFileRequest(file string) ([]byte, error ){
    /*
	To read a file you will first send a command with a json payload of:
    {
     "cmd" : "read",
     "path" : "/path/to/file"
    }
    
    followed by data blocks
    F0 00 01 55 [FID] [PID] 50 01 [SIZE HI] [SIZE LO] F7
    where SIZE is the number of 8-bit bytes to read
    
    e.g. to request 256 bytes
    F0 00 01 55 18 01 50 01 02 00 F7

	*/

	type Message struct {
		Cmd  string `json:"cmd"`
		Path string `json:"path"`
	}

	data := Message{"read", file}
	// make a json object
	json_obj, err := json.Marshal(data)
	if err != nil {
		log.Printf("ERROR: %s\n", err)
		return nil, err
	}
	// json_obj is a byte array
	//log.Printf("json_obj: %s\n", json_obj)
    
    bt := append(cmd_hdr, xfer_cmd_opcode...)
    bt = append(cmd_hdr, json_obj...)
	// make a sysex message, which will prepend F0 and append F7 to byte slice
	m := midi.SysEx(bt)
	//log.Println("sysex msg to send: ", m)
    return m, err
}

