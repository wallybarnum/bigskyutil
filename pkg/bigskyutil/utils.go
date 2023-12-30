package bigskyutil

import "fmt"



func dumpByteSlice(b []byte) {
	var a [16]byte
	n := (len(b) + 15) &^ 15
	for i := 0; i < n; i++ {
		if i%16 == 0 {
			fmt.Printf("%4d", i)
		}
		if i%8 == 0 {
			fmt.Print(" ")
		}
		if i < len(b) {
			fmt.Printf(" %02X", b[i])
		} else {
			fmt.Print("   ")
		}
		if i >= len(b) {
			a[i%16] = ' '
		} else if b[i] < 32 || b[i] > 126 {
			a[i%16] = '.'
		} else {
			a[i%16] = b[i]
		}
		if i%16 == 15 {
			fmt.Printf("  %s\n", string(a[:]))
		}
	}
}

/*
const fid string = "18"
const pid string = "01"
const xfer_opcode string = "50"
const xfer_cmd_opcode string = "00"
const xfer_data_opcode string = "01"

//file transfer protocol header
const hdr string = "F0 00 01 55 " + fid + " " + pid + " " + xfer_opcode + " "

// the header for a command is:
// F0 00 01 55 [FID] [PID] 50 00 [JSON PAYLOAD] F7
const cmd_hdr string = hdr + xfer_cmd_opcode + " "

//the header for a data block is:
//F0 00 01 55 [FID] [PID] 50 01 [SIZE HI] [SIZE LO] [DATA] F7
//
// SIZE HI, SIZE LO, and DATA are not always needed depending on the command.
const data_hdr string = hdr + xfer_data_opcode + " "
*/
/*
func syx8to7(data []byte){
    
    //convert an array of 8 bit data bytes to 7 bit midi bytes
    
    var outputBytes = 0
    var u7 = 0

    //start on byte 1 
    var input_ix int = 0
    var i int = 1
    var msbs int = 0
    var num = len(data)
    
    out := []byte{0}

    for num > 0{
		//the first byte in a string of 8 will be the MSBs of the following 7
        if i % 8 == 0{
            out[i - 8] = msbs & 0x7F
            msbs = 0
		} else {
            //g et the input byte and save the MSB
            u7 = data[input_ix]
            input_ix = input_ix + 1
            
            msbs = msbs << 1
            msbs = msbs | ((u7 >> 7) & 0x1)
            out[i] = u7 & 0x7F
            num = num - 1
		}
        i = i + 1
        outputBytes = outputBytes + 1
	}
    // get the msbs if things don't line up right
    if i%8 != 1{
        out[i - (i%8)] = (msbs << (8 - (i % 8))) & 0x7F
        outputBytes = outputBytes + 1
	}
    return out[:outputBytes]
}
*/
/*
func syx7to8(data []byte){
    //convert an array of 7 bit midi bytes to 8 bit data bytes
    outputBytes = 0
    i = 0
    ix = 0
	out := []byte{0}

    msbs = 0
    for i < len(data){
        if i % 8 == 0{
            // most significant bits of the following 7 bytes
            msbs = data[i]
            msbs = msbs << 1
		} else{
            // get the data byte and or in the MSB
            out[ix] = (data[i] | (msbs & 0x80))
            ix = ix + 1
            
            msbs = msbs << 1
            outputBytes = outputBytes + 1
        i = i + 1
		}
	}
    return out[:ix]
}
*/
/*
func _json2str(s){}
    s= s.encode('ascii')
    s = s.hex().upper()

    split_strings = []
    n  = 2
    for index in range(0, len(s), n):
        split_strings.append(s[index : index + n])

    return " ".join(split_strings)

def _bytes2str(b):
    """
    Convert a byte array to a string
    e.g. b'\x0a\x02\x1f' to '0A 02 1F'
    """
    dat = ['{:02X}'.format(elem) for elem in b]
    return " ".join(dat)

def _str2bytes(s):
    """
    Convert a string to byte array
    e.g. '0A 02 1F' -> b'\x0a\x02\x1f'
    """
    spacesRemoved = "".join(s.split())
    return bytearray.fromhex(spacesRemoved)
*/

