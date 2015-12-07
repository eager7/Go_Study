package main
import (
	"fmt"
	"socketClient"
	//"time"
	"encoding/json"
	"os"
	"runtime"
	//"bytes"
)

type Device struct {
	DeviceName		string	
	DeviceId 		uint16
	DeviceAddr		uint64
}
var Devices []Device = make([]Device, 0, 10)
var sock socketClient.SockClient

func main(){
	fmt.Println("zigbee daemon test, test the performance")

	initSocket()
	searchZigbee()
	fmt.Println("Devices:", Devices)

	for i := 0; i < 10; i++ {
		for _, d := range Devices {
			fmt.Println(d.DeviceAddr)
		}
	}

	defer sock.Finished()
}

func initSocket(){
	err := sock.Init("127.0.0.1", 6667)
	checkError(err)
	fmt.Println("Init Finished")
}

func searchZigbee() {
	type searchJson struct {
		Command 	int `json:"command"`
		Sequence 	int `json:"sqeuence"`
	}

	search := searchJson{
		Command: 0,
		Sequence:0,
	}
	SerchCommand, err := json.Marshal(search)
	checkError(err)
	fmt.Println(string(SerchCommand))

	DeviceList, sta := sock.SendMsgWithResp(string(SerchCommand))
	if sta != 0 {
		fmt.Println("can't recv msg")
		return
	}

	type Description struct {
		DeviceName	string	`json:"device_name"`
		DeviceId 	uint16	`json:"device_id"`
		DeviceAddr	uint64	`json:"device_mac_address"`
	}
	type DeviceListInf struct {
		Status 		int 	`json:"status"`
		Sequence 	int 	`json:"sequence"`
		Descript 	[]Description `json:"description"`
	}
	var devicelists DeviceListInf
	err = json.Unmarshal([]byte(DeviceList), &devicelists)
	if _, _, line, _ := runtime.Caller(0); err != nil {
		fmt.Println(line, err)
	}
	fmt.Println("devicelists:", devicelists)
	if devicelists.Status == 0{
		for _, d := range devicelists.Descript {
			Devices = append(Devices, Device{d.DeviceName, d.DeviceId, d.DeviceAddr})
		}
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}