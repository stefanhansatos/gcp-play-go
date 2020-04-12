package functions

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"time"
)

type Registration struct {
	ReporterID           string    `json:"reporter"`
	Contagious           bool      `json:"contagious"`
	TimeContagionUpdated time.Time `json:"time-contagion-updated"`
}

func PubsubDecodeRegistration(body string) (Registration, error) {

	fmt.Printf("PubsubBodyDecode: body: %s (%t)\n", body, body)

	decoded, err := base64.StdEncoding.DecodeString(body)
	if err != nil {
		fmt.Println("PubsubBodyDecode: decode data error:", err)
	}
	fmt.Printf("PubsubBodyDecode: decoded: %s (%t)\n", decoded, decoded)

	bytes := []byte(string(decoded))
	var registration Registration
	err = json.Unmarshal(bytes, &registration)

	// Unmarshal request body
	//err = json.Unmarshal(bytes, &registration)

	return registration, err
}
