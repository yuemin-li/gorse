package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"

    "github.com/dlintw/goconf"
    "github.com/go-martini/martini"
    "github.com/kuwagata/martini-keystone-auth"
    "github.com/user/gorse-cache-redis"
    // "github.com/user/gorse-storage-kafka"
    // "github.com/user/gorse-storage-redis"
)


func main() {
    cfg, err := goconf.ReadConfigFile("rse.config")
    if err != nil {
        panic(err)
    }
    
    m := martini.Classic()

    auth_handler := setupAuthHandler(cfg)

    m.Handlers(
        auth_handler,
        martini.Recovery(),
    )

    m.Get("/**", func(request *http.Request, params martini.Params, response http.ResponseWriter) string {
        topic := params["_1"]
        marker, marker_present := request.URL.Query()["marker"]

        // Get message from storage
        // if marker_present {
        //     storage.Get(topic, marker)
        // } else {
        //     storage.Get(topic)
        // }

        // Example return
        events := make([]map[string]string, 0)
        event := make(map[string]string)
        event["property"] = "value"
        events = append(events, event)
        jsonString, _ := json.Marshal(events)
        fmt.Println("Events: " + string(jsonString))

        tempReturn := "Getting messages on " + topic
        if marker_present {
            tempReturn += " with marker of " + marker[0]
        }
        // fmt.Println(tempReturn)
        // return tempReturn + "\r\n"

        response.Header()["Content-Type"] = []string{"application/json"}
        return string(jsonString) + "\r\n"
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


func setupAuthHandler(cfg goconf.ConfigFile) martini.Handler {
    auth_url, _ := cfg.GetString("Auth", "endpoint")
    hostname, _ := cfg.GetString("Cache", "hostname")
    port, _ := cfg.GetString("Cache", "port")
    password, _ := cfg.GetString("Cache", "password")

    return auth.Keystone(
        auth.IdentityValidator{AuthUrl: auth_url},
        auth.Redis{
            Hostname: hostname,
            Port:     port,
            Password: password,
            Database: int64(0)})
}
