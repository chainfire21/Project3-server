package typeform

import (
	"net/http"
	"log"

	"io/ioutil"
	"os"

	"github.com/Jeffail/gabs"

)

func GetSurveyDataCoach() *gabs.Container{
	httpClient := &http.Client{}
	req, err := http.NewRequest("GET", "https://api.typeform.com/forms/"+os.Getenv("TYPEFORM_COACH")+"/responses",nil)
	if err != nil{
		log.Fatal(err)
	}


	req.Header.Add("Authorization", "Bearer "+os.Getenv("TYPEFORM_AUTH"))
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


func GetSurveyDataClient() *gabs.Container{
	httpClient := &http.Client{}
		// req, err := http.NewRequest("GET", "https://api.typeform.com/forms/"+os.Getenv("TYPEFORM_CLIENT")+"/responses",nil)
		req, err := http.NewRequest("GET", "https://api.typeform.com/forms/GA1xBQ/responses",nil)
		if err != nil{
			log.Fatal(err)
		}


	// req.Header.Add("Authorization", "Bearer "+os.Getenv("TYPEFORM_AUTH"))
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

