package socketClient

import "testing"
import "time"
import "fmt"
import "strconv"

func TestInit(t *testing.T){
	addr := "127.0.0.1"
	port := 6667
	s := new(SockClient)
	s.Init(addr, port)

	result, sta := s.SendMsgWithResp("{\"command\":20,\"sequence\":2}")

	if sta != 0 && result != "{\"status\":0,\"sequence\":2,\"description\":\"1.2\"}" {
		t.Error("Get Version Error")
		return
	}

	defer s.Finished()
}

func TestZigbee(t *testing.T) {
	addr := "127.0.0.1"
	port := 6667
	s := new(SockClient)
	s.Init(addr, port)

	const total int = 10
	const interval = time.Millisecond*100
	var num_succ int
	var sequence int = 0
	for i := 0; i < total; i++ {

		sequence++
		openLight  := "{\"command\":6,\"sequence\":" +strconv.Itoa(sequence) + ",\"device_address\":\"6066005655905860\",\"group_id\":0}"
		result, sta := s.SendMsgWithResp(openLight)
		verify := "{\"status\":0,\"sequence\":" +strconv.Itoa(sequence) +",\"description\":\"success\"}"
		if sta != 0 && result != verify{
			t.Error("Get Version Error")
		} else {
			num_succ ++
		}
		time.Sleep(interval)

		sequence++
		getLight   := "{\"command\":8,\"sequence\":" +strconv.Itoa(sequence) +",\"device_address\":\"6066005655905860\"}"
		result, sta = s.SendMsgWithResp(getLight)
		verify = "{\"status\":0,\"sequence\":" +strconv.Itoa(sequence) +",\"description\":{\"light_status\":1}}"
		if sta != 0 && result != verify{
			t.Error("Get Version Error")
		} else {
			num_succ ++
		}
		time.Sleep(interval)
				
		sequence++
		closeLight := "{\"command\":7,\"sequence\":" +strconv.Itoa(sequence) +",\"device_address\":\"6066005655905860\",\"group_id\":0}"
		result, sta = s.SendMsgWithResp(closeLight)
		verify = "{\"status\":0,\"sequence\":" +strconv.Itoa(sequence) +",\"description\":\"success\"}"
		if sta != 0 && result != verify{
			t.Error("Get Version Error")
		} else {
			num_succ ++
		}	
		time.Sleep(interval)

		sequence++
		getLight = "{\"command\":8,\"sequence\":" +strconv.Itoa(sequence) +",\"device_address\":\"6066005655905860\"}"
		result, sta = s.SendMsgWithResp(getLight)
		verify = "{\"status\":0,\"sequence\":" +strconv.Itoa(sequence) +",\"description\":{\"light_status\":0}}"
		if sta != 0 && result != verify {
			t.Error("Get Version Error")
		} else {
			num_succ ++
		}
		time.Sleep(interval)
	}
	fmt.Println("The Success num is", num_succ)
	var pre float32 = (float32(num_succ)/(float32(total)*4))*100
	fmt.Printf("The Success pre is %f\n", pre)

	defer s.Finished()
}