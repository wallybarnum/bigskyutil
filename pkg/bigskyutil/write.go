package bigskyutil

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"gitlab.com/gomidi/midi/v2"
	_ "gitlab.com/gomidi/midi/v2/drivers/rtmididrv"
)

var block_size = 64

func WriteFile(path string, data []byte) error {

	defer midi.CloseDriver()

	in, err := midi.FindInPort("BigSkyMX")
	if err != nil {
		log.Println("can't find BigSkyMX")
		return err
	}
	// listens to the in port and calls eachMessage for every message.
	// any running status bytes are converted and only complete messages are passed to the eachMessage.
	stop, err := midi.ListenTo(in, EachMessage, midi.UseSysEx())
	if err != nil {
		log.Println("can't listen to in port")
		return err
	}
	out, err := midi.FindOutPort("BigSkyMX")
	if err != nil {
		log.Println("can't find BigSkyMX")
		return err
	}

	// creates a sender function to the out port
	send, err := midi.SendTo(out)
	if err != nil {
		log.Println("can't make send function")
		return err
	}

	// make a write file sysex message
	m, err := WriteFileRequest(path)
	if err != nil {
		log.Println("can't build write request")
		return err
	}
	err = send(m) // send the write request sysex message
	if err != nil {
		log.Println("can't send sysex message err = ", err)
		return err
	}
	time.Sleep(time.Millisecond * 200) // delay for the heck of it
	Sig = true
	// send data blocks
	for i := 0; i < len(data); i += block_size {
		if Sig == false {
			log.Println("waiting for response")
			time.Sleep(time.Millisecond * 200)
		} else {
			Sig = false
			//log.Println("sending block ", i)
			if i+block_size > len(data) {
				m, err = DataBlockTransfer(data[i:], len(data[i:]))
			} else {
				m, err = DataBlockTransfer(data[i:i+block_size], block_size)
			}
			if err != nil {
				log.Println("can't build data block transfer sysex message")
				return err
			}
			err = send(m) // send the data block sysex message
			if err != nil {
				log.Println("can't send sysex message err = ", err)
				return err
			}
			time.Sleep(time.Millisecond * 500) // delay for the heck of it
		}
	}

	// send close cmd
	m, err = CloseFileRequest(path)
	if err != nil {
		log.Println("can't build write request")
		return err
	}
	err = send(m) // send the close file sysex message
	if err != nil {
		log.Println("can't send sysex message err = ", err)
		return err
	}
	time.Sleep(time.Millisecond * 200) // delay for the heck of it
	// stops listening
	stop()
	return err
}

func WriteFileRequest(file string) ([]byte, error) {

	// read a local file from disk and create a sysex file that will write that file to a remote device.
	//
	// To write a file you will first send a command with a json payload of:
	// {
	//  "cmd" : "write",
	//  "path" : "/path/to/file"
	// }
	//
	// followed by data blocks
	// F0 00 01 55 [FID] [PID] 50 01 [SIZE HI] [SIZE LO] [DATA] F7
	// where SIZE is the number of 8-bit bytes to write
	//
	// after writing is complete a close command should be sent
	// {
	//  "cmd" : "close"
	// }

	type Message struct {
		Cmd  string `json:"cmd"`
		Path string `json:"path"`
	}

	cmd := Message{"write", file}
	// make a json object
	json_obj, err := json.Marshal(cmd)
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

func DataBlockTransfer(data []byte, size int) ([]byte, error) {

	//    data block
	//    F0 00 01 55 [FID] [PID] 50 01 [SIZE HI] [SIZE LO] [DATA] F7
	//    where SIZE is the number of 8-bit bytes to write
	//

	// the data is 8 bit bytes, but sent as 7 bit bytes
	// by spreading 7 data bytes over 8 transmitted bytes
	// the first transmitted byte contains the msbs of the following
	// 7 data bytes.
	// if the data is all ascii < 128, then their msbs will all be 0
	// so the data can be converted to 8 bit bytes by dropping
	// the msbs and setting the first transmitted byte to zero
	// sending non-ascii data will require a more complicated
	// conversion.
	// TODO: properly convert 8 bit bytes to 7 bit bytes

	// append length of data
	bt := append(data_hdr, byte(size>>8))
	bt = append(bt, byte(size))
	// append data - convert 8 bit bytes to 7 bit bytes
	escaped_data := []byte{}
	for i := 0; i < len(data); i++ {
		if (i+7)%7 == 0 {
			escaped_data = append(escaped_data, 0x00) // t
			escaped_data = append(escaped_data, data[i])
		} else {
			escaped_data = append(escaped_data, data[i])
		}
	}
	bt = append(bt, escaped_data...)
	fmt.Println("data transfer msg to send: ")
	DumpByteSlice(bt)

	// make a sysex message, which will prepend F0 and append F7 to byte slice
	m := midi.SysEx(bt)
	return m, nil
}
func CloseFileRequest(file string) ([]byte, error) {

	// create a sysex file that will close file on a remote device.
	//
	// after writing is complete, a close command should be sent
	// {
	//  "cmd" : "close"
	// }

	type Message struct {
		Cmd string `json:"cmd"`
	}

	data := Message{"close"}
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
