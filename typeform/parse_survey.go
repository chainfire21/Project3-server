package typeform

import (
	"net/http"
	"log"
	"fmt"
	// "io/ioutil"
	"os"
	// "bytes"
	"encoding/json"
)

func GetSurveyData(){
	// var prettyJSON bytes.Buffer
	httpClient := &http.Client{}
	
	req, err := http.NewRequest("GET", "https://api.typeform.com/forms/pHdQSo/responses",nil)
	req.Header.Add("Authorization", "Bearer Gr8o49DXvMTTVnDaCNjz86mS2kE283snRA4S25ULogmk")
	response, err := httpClient.Do(req)
	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	} else {
		defer response.Body.Close()
		var u Headers
		errr := json.NewDecoder(response.Body).Decode(&u)
		if errr != nil{
			log.Fatal(errr)
		}
		log.Println(u)
		// if err != nil {
		// 	fmt.Printf("%s", err)
		// 	os.Exit(1)
		// }
		// error := json.Indent(&prettyJSON, contents, "", "\t")
		// if error != nil {
		// 	log.Println("JSON parse error: ", error)
		// 	return
		// }
		// fmt.Printf("%s\n", string(prettyJSON.Bytes()))
	}
}
