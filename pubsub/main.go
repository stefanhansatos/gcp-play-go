package main

import (
	"cloud.google.com/go/pubsub"
	"context"
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type Registration struct {
	ReporterID           string    `json:"reporter"`
	Contagious           bool      `json:"contagious"`
	TimeContagionUpdated time.Time `json:"time-contagion-updated"`
}

func main() {

	gcpProject := os.Getenv("GCP_PROJECT")
	topicName := "cf-sink"

	ctx := context.Background()
	pubsubClient, err := pubsub.NewClient(ctx, gcpProject)
	if err != nil {
		fmt.Printf("failed to create new pubsub client: %v", err)
	}
	defer func() {
		pubsubClient.Close()
	}()

	gre := &Registration{
		ReporterID:           "me",
		Contagious:           true,
		TimeContagionUpdated: time.Now(),
	}

	msg, err := json.Marshal(gre)
	if err != nil {

		fmt.Printf("failed to marshal message: %v", msg, err)
	}

	fmt.Printf("message: %s", msg)

	//bytes := []byte(string(msg))
	//var registration Registration
	//err = json.Unmarshal(bytes, &registration)
	//if err != nil {
	//
	//	fmt.Printf("failed to unmarshal msg: %v", err)
	//	return
	//}
	//
	//var regi *Registration
	//
	//err = json.Unmarshal(msg, regi)
	//if err != nil {
	//
	//	fmt.Printf("failed to unmarshal message: %v", err)
	//	return
	//}
	//
	//return

	//msg := fmt.Sprintf("{%q: false,%q: %q}", "contagious", "time-contagion-updated", "2020-04-05T11:38:36.899503Z")

	result := pubsubClient.Topic(topicName).Publish(context.Background(), &pubsub.Message{
		Data: msg,
	})

	// Block until the result is returned and a ws-server-generated
	// ID is returned for the published message.
	id, err := result.Get(ctx)
	if err != nil {
		fmt.Printf("failed to publish message %q: %v", msg, err)
	}
	fmt.Printf("Published message %q (ID: %v)\n", msg, id)

}

//
//func publish(w io.Writer, projectID, topicID, msg string) error {
//	// projectID := "my-project-id"
//	// topicID := "my-topic"
//	// msg := "Hello World"
//	ctx := context.Background()
//	client, err := pubsub.NewClient(ctx, projectID)
//	if err != nil {
//		return fmt.Errorf("pubsub.NewClient: %v", err)
//	}
//
//	t := client.Topic(topicID)
//	result := t.Publish(ctx, &pubsub.Message{
//		Data: []byte(msg),
//	})
//	// Block until the result is returned and a ws-server-generated
//	// ID is returned for the published message.
//	id, err := result.Get(ctx)
//	if err != nil {
//		return fmt.Errorf("Get: %v", err)
//	}
//	fmt.Fprintf(w, "Published a message; msg ID: %v\n", id)
//	return nil
//}
