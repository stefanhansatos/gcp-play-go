package functions

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func Receive(w http.ResponseWriter, r *http.Request) {

	// Read the request body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to read request: %s", err), http.StatusInternalServerError)
		return
	}
	fmt.Printf("Body: %s\n", body)

	// Unmarshal request body
	bytes := []byte(string(body))
	var jsonMap map[string]interface{}

	err = json.Unmarshal(bytes, &jsonMap)
	if err != nil {
		fmt.Printf("cannot unmarshal JSON input: %s", err)
		return
	}

	fmt.Printf("jsonMap[%q]: %s (%t)\n", "message", jsonMap["message"], jsonMap["message"])

	if msgMap, ok := jsonMap["message"].(map[string]interface{}); ok {

		fmt.Printf("msgMap[%q]: %q\n", "data", msgMap["data"])

		if data, ok := msgMap["data"].(string); ok {
			fmt.Printf("data: %q\n", data)

			//decoded, err := base64.StdEncoding.DecodeString(data)
			//if err != nil {
			//	fmt.Println("decode data error:", err)
			//}
			//
			//fmt.Printf("decoded data: %q\n", string(decoded))

			var registration Registration

			if registration, err = PubsubDecodeRegistration(data); err != nil {
				fmt.Printf("cannot decode pubsub body: %s", err)
				return
			}
			fmt.Printf("registration: %s\n", registration)
		}

		//for key, value := range msgMap {
		//	fmt.Printf("jsonMap[%q]: %s (%t)\n", key, value, value)
		//}
	}

	//var msg pubsub.Message
	//err = json.Unmarshal(bytes, &msg)
	//if err != nil {
	//	http.Error(w, fmt.Sprintf("cannot unmarshal JSON input: %s", err), http.StatusInternalServerError)
	//	return
	//}
	//
	//
	//
	//fmt.Printf("msg.ID: %s\n", string(msg.ID))

	//decoded, err := base64.StdEncoding.DecodeString(string(msg.Data))
	//if err != nil {
	//	fmt.Println("decode data error:", err)
	//
	//	decoded, err := base64.StdEncoding.DecodeString("SGVsbG8gV29ybGQ=")
	//	if err != nil {
	//		fmt.Println("decode error:", err)
	//		return
	//	}
	//	fmt.Println(string(decoded))
	//	return
	//}
	//fmt.Println(string(decoded))

	return
}

/*
gcloud functions deploy receive --region europe-west3 \
--entry-point Receive --runtime go113 --trigger-http \
--allow-unauthenticated
*/
