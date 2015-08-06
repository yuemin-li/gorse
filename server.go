package main

import (
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

    m.Get("/**", func(r *http.Request, params martini.Params) string {
        topic := params["_1"]
        marker, marker_present := r.URL.Query()["marker"]

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

    m.Run()
}
