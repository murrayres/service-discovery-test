package main

import "fmt"
import "bytes"
import "net/http"
import "io/ioutil"
import "os"
import "github.com/gin-gonic/gin"

func main() {

    router := gin.Default()
    router.POST("/", publishEvent)
    router.Run()
}

func publishEvent(c *gin.Context) {
    event, err := ioutil.ReadAll(c.Request.Body)
    if err != nil {
        fmt.Println(err)
        c.String(500, err.Error())
        return
    }
    bs := bytes.NewBuffer(event)
    publisherhost := os.Getenv("EVENTS_SERVICE_HOST")
    publisherport := os.Getenv("EVENTS_SERVICE_PORT")
    publisheruri := "/"
    fulllocation := "http://" + publisherhost + ":" + publisherport + publisheruri
    fmt.Println(fulllocation)
    client := http.Client{}
    req, err := http.NewRequest("POST", fulllocation, bs)
    if err != nil {
        fmt.Println(err)
        c.String(500, err.Error())
        return
    }
    req.Header.Add("Content-type", "application/json")
    resp, err := client.Do(req)
    if err != nil {
        fmt.Println("Error sending request")
        fmt.Println(err)
        c.String(500, err.Error())
        return
    }
    defer resp.Body.Close()

    bodybytes, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        fmt.Println(err)
        c.String(500, err.Error())

        return
    }
    c.String(resp.StatusCode, string(bodybytes))
}
