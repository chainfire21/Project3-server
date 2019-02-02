package typeform

import (
	"net/http"
	"log"
	// "fmt"
	"io/ioutil"
	// "os"
	// "bytes"
	// "encoding/json"
	"github.com/Jeffail/gabs"

)

func GetMatches() {
	surveys := GetSurveyData()
	log.Print(surveys)
	// S is shorthand for Search
	// fmt.Println(surveys.Path("items.value").String())
}

func GetSurveyData() *gabs.Container{
	httpClient := &http.Client{}
	// req, err := http.NewRequest("GET", "https://api.typeform.com/forms/"+os.Getenv(TYPEFORM_COACH)+"/responses",nil)
	req, err := http.NewRequest("GET", "https://api.typeform.com/forms/pHdQSo/responses",nil)
	if err != nil{
		log.Fatal(err)
	}
	// req.Header.Add("Authorization", "Bearer "+os.Getenv(TYPEFORM_AUTH))
	req.Header.Add("Authorization", "Bearer Gr8o49DXvMTTVnDaCNjz86mS2kE283snRA4S25ULogmk")
	response, err := httpClient.Do(req)
	if err != nil{
		log.Fatal(err)
	}
	defer response.Body.Close()
	// var surveys Headers
	// json.Unmarshal(response.Body, &surveys)
	// log.Println(response.Body)
	body, err := ioutil.ReadAll(response.Body)
	if err != nil{
		log.Fatal(err)
	}
	jsonParsed, err := gabs.ParseJSON([]byte(body))
	return jsonParsed
	

}
