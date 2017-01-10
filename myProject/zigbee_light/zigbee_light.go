package main
import(
	"log"
	"socketClient"
	"encoding/json"
	"os"
	"runtime"
	"time"
)

var sock_zigbee socketClient.SockClient
var device_lists []device_info

type device_info struct{
	device_name string
	device_id uint16
	device_online uint8
	device_mac uint64
}

func main(){
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	log.Println("zigbee light test")

	GetDeviceLists()
	for i := 0; i < 10; i++{
		SetLightOn()
		time.Sleep(10000)
		SetLightOff()
		time.Sleep(10000)
	}
	defer sock_zigbee.Finished()
}

func GetDeviceLists(){
	err := sock_zigbee.Init("127.0.0.1", 6667)
	if err != nil{
		log.Fatal("Error:", err)
		os.Exit(1)
	}

	type searchCmd struct {
		Command int `json:"command"`
		Sequence int `json:"sequence"`
	}
	cmd, _ := json.Marshal(searchCmd{0x0011,0})
	ret, sta := sock_zigbee.SendMsgWithResp(string(cmd))
	if sta != 0{
		log.Fatal("recv msg error")
		os.Exit(1)
	}

	type description struct{
		Device_name string `json:"device_name"`
		Device_id uint16 `json:"device_id"`
		Device_online uint8 `json:"device_online"`
		Device_mac uint64 `json:"device_mac_address"`
	}
	type resp struct {
		Status uint8 `json:"status"`
		Sequence int `json:"sequence"`
		Desc []description `json:"description"`
	}
	var r resp;
	err = json.Unmarshal([]byte(ret), &r)
	log.Println(err)
	if _, _, line, _ := runtime.Caller(0); err != nil {
		log.Println(line, err)
	}

	if r.Status == 0{
		for _,dev := range r.Desc{
			device_lists = append(device_lists, device_info{dev.Device_name,
				dev.Device_id, dev.Device_online, dev.Device_mac})
		}
	}
	log.Println(device_lists)
}

func SetLightOn(){
	type jsoncmd struct {
		Cmd 	uint16 	`json:"command"`
		Seq 	int 	`json:"sequence"`
		Addr 	uint64 	`json:"device_address"`
		Group 	uint8 	`json:"group_id"`
		Mod 	uint8 	`json:"mode"`
	}
	for i := range device_lists{
		if device_lists[i].device_id == 0x0101{
			cmd_on,_ := json.Marshal(jsoncmd{0x0020, 0, device_lists[i].device_mac, 0, 1})
			sock_zigbee.SendMsgWithResp(string(cmd_on))
		}
	}
}

func SetLightOff(){
	type jsoncmd struct {
		Cmd 	uint16 	`json:"command"`
		Seq 	int 	`json:"sequence"`
		Addr 	uint64 	`json:"device_address"`
		Group 	uint8 	`json:"group_id"`
		Mod 	uint8 	`json:"mode"`
	}
	for i := range device_lists{
		if device_lists[i].device_id == 0x0101{
			cmd_on,_ := json.Marshal(jsoncmd{0x0020, 0, device_lists[i].device_mac, 0, 0})
			sock_zigbee.SendMsgWithResp(string(cmd_on))
		}
	}
}