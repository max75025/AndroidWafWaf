package main

import (
	"net/http"

	"io/ioutil"
	"fmt"
	"time"
	"strconv"
	"log"
	_"github.com/max75025/go-sqlite3"
	"encoding/json"
	"os"
	"database/sql"
)
const constEndTime = 9999999999999999

type event struct{
	DateTime 	int
	TypeTrace 	[]string
	ResultTypes []string
	IpAddr		string
	Country		string
}

func saveEventToDB(jsonStr string) error{
	var events []event
	json.Unmarshal([]byte(jsonStr),&events)
	//fmt.Println(events)


	return nil
}

func getEventClient(apiKey string, startTime int, endTime int)(string,error)  {
	url := "http://wafwaf.tech/eventclient/" + apiKey + "/" + strconv.Itoa(startTime)+"/"+ strconv.Itoa(endTime)
	//fmt.Println(url)
	resp,err:= http.Get(url)
	if err!= nil {
		//log.Println(err)
		return "",err
	}
	defer resp.Body.Close()
	content,err:= ioutil.ReadAll(resp.Body)
	if err!=nil{
		//log.Println(err)
		return "",err
	}

	return string(content), nil
}

func autoCheckNewEventClient(apiKey string){

	for  range time.Tick(10 *time.Second){
		currentTime := int(time.Now().Unix())
		//fmt.Println("check...")
		result,err:= getEventClient(apiKey,currentTime-10,constEndTime )
		if err!=nil{
			log.Println(err)
		}
		fmt.Println(result)
	}
}



func main(){
	const testApiKey  = "5a9ebd7d5f7c8cc17f385f2b36b26181a03fb3dfe78c512cb71f538869a7ea8d6b803385245dfcb698d47be097c82d4759eed12ad106021e2cfa646f905cacfc"
	const testApiStartTime = 1532449279
	const monthInSecond = 2592000

	 newDB:= false
	if _, err := os.Stat("./db.db"); os.IsNotExist(err) {
		_,fileErr:=os.Create("./db.db")
		if fileErr!=nil{log.Println(err)}else{newDB = true}
	}

	db,err:= sql.Open("sqlite3","./db.db" )
	if err!=nil{log.Println(err)}
	if newDB{
		_,err:= db.Exec("CREATE TABLE `event`( `DateTime` INTEGER, `TypeTrace` TEXT , `ResultTypes` TEXT,`IpAddr`	TEXT,`Country` TEXT)")
		if err!=nil{log.Println(err)}
	}

	result,err:=getEventClient(testApiKey,int(time.Now().Unix())-monthInSecond,constEndTime)
	if err!= nil{
		log.Println(err)
	}else{
		fmt.Println(result)
	}




	 autoCheckNewEventClient(testApiKey)

}