package main

import (
    "os"
    "log"
    "fmt"
    "time"
    "encoding/json"

    "github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/session"
    "github.com/aws/aws-sdk-go/service/sfn"
)

const RFC3339 = "2006-01-02T15:04:05+07:00"


type Payload struct {
    Name string `json:"name"`
    Job string `json:"job"`
}

type Task struct {
    Url string `json:"url"`
    Method string `json:"method"`
    Payload *Payload `json:"payload,omitempty"`
    Schedule string `json:"scheduleTime,omitempty"`
}

func main(){
    now := time.Now().Local()
    delay := 60 // seconds

    task := &Task{
        Url: "https://reqres.in/api/users",
        Method: "POST",
        Payload: &Payload{
            Name: "Joko",
            Job: "Dukun",
        },
        Schedule: now.Add(time.Second * time.Duration(delay)).Format(RFC3339),
    }

    t, err := json.Marshal(task)
    if err != nil {
        log.Fatal(err)
    }

    sess, err := session.NewSession(&aws.Config{
        Region: aws.String("ap-southeast-3")},
    )
    if err != nil {
        log.Fatal(err)
    }

    svc := sfn.New(sess)

    output, err := svc.StartExecution(&sfn.StartExecutionInput{
    	Input: aws.String(string(t)),
    	StateMachineArn: aws.String(os.Getenv("STEP_FUNCTION_ARN")),
    })

    if err != nil {
        log.Fatal(err)
    }

    msg := "Execution triggered " + aws.StringValue(output.ExecutionArn)
    log.Println(msg)
}

