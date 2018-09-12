package WafLibrary

import (
	"net/http"
	"io/ioutil"
	"strconv"
	_ "github.com/max75025/go-sqlite3"
	"encoding/json"
)

const constEndTime = 2000000000

const testApiKey = "5a9ebd7d5f7c8cc17f385f2b36b26181a03fb3dfe78c512cb71f538869a7ea8d6b803385245dfcb698d47be097c82d4759eed12ad106021e2cfa646f905cacfc"
const testApiStartTime = 1532449279
const monthInSecond = 2592000


type event struct {
	DateTime    int
	TypeTrace   []string
	ResultTypes []string
	IpAddr      string
	Country     string
}

type av struct {
	ApiKey               string
	EventTime            int
	EventType            string
	FileName             string
	FileExt              string
	FilePath             string
	SuspiciousType       string
	SuspiciousDescripton string
}



func GetJsonEvent(apiKey string, startTime int, endTime int) (string, error) {
	url := "http://wafwaf.tech/eventclient/" + apiKey + "/" + strconv.Itoa(startTime) + "/" + strconv.Itoa(endTime)
	//fmt.Println(url)
	resp, err := http.Get(url)
	if err != nil {
		//log.Println(err)
		return "", err
	}
	defer resp.Body.Close()
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		//log.Println(err)
		return "", err
	}

	return string(content), nil
}

func GetJsonAV(apiKey string, startTime int, endTime int) (string, error) {
	url := "http://wafwaf.tech/eventav/" + apiKey + "/" + strconv.Itoa(startTime) + "/" + strconv.Itoa(endTime)
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(content), nil
}

func GetAV(apiKey string, startTime int, endTime int) ([]av,error){
	var result []av
	jsonStr,err:= GetJsonAV(apiKey,startTime, endTime)
	if err!=nil{
		return nil,err
	}
	err=json.Unmarshal([]byte(jsonStr),&result)
	if err!=nil{
		return nil,err
	}
	return result,nil
}

func GetEvent(apiKey string, startTime int, endTime int) ([]event,error){
	var result []event

	jsonStr,err:= GetJsonEvent(apiKey,startTime,endTime)
	if err!=nil{
		return nil,err
	}
	err = json.Unmarshal([]byte(jsonStr),&result)
	if err!=nil{
		return nil,err
	}
	return result,nil
}


func main() {
	//fmt.Println(GetEvent(testApiKey,constEndTime,5))
}
