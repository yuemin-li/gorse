package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"

    "github.com/go-martini/martini"
    // "github.com/user/gorse-cache-redis"
    // "github.com/user/gorse-handler-auth"
    // "github.com/user/gorse-storage-kafka"
    // "github.com/user/gorse-storage-redis"
)

func main() {
    m := martini.Classic()

    m.Handlers()

    m.Get("/**", func(request *http.Request, params martini.Params) string {
        topic := params["_1"]
        marker, marker_present := request.URL.Query()["marker"]

        // Get message from storage
        // if marker_present {
        //     storage.Get(topic, marker)
        // } else {
        //     storage.Get(topic)
        // }

        tempReturn := "Getting messages on " + topic
        if marker_present {
            tempReturn += " with marker of " + marker[0]
        }
        return tempReturn + "\r\n"
    })

    m.Post("/**", func(request *http.Request, params martini.Params, response http.ResponseWriter) {
        topic := params["_1"]

        body, err := ioutil.ReadAll(request.Body)
        if err != nil {
            panic("Unable to read body")
        }
        var valid_json interface{}
        err = json.Unmarshal(body, &valid_json)
        if err != nil {
            panic("Invalid JSON")
        }

        fmt.Println("Topic: " + topic)
        fmt.Println("Message: " + string(body))
        // Send message to storage
        // storage.Insert(topic, string(body))

        response.WriteHeader(http.StatusCreated)
    })

    m.Run()
}
