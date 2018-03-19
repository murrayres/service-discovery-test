package main

import "fmt"
import "net/http"
import "io/ioutil"
import "os"
import "github.com/gin-gonic/gin"

func main() {

    router := gin.Default()
    router.GET("/send", sendrequest)
    router.GET("/respond", respondtorequest)
    router.Run()
}



func sendrequest(c *gin.Context) {
    
//    host := os.Getenv("SDRESPOND_SERVICE_HOST")
//    port := os.Getenv("SDRESPOND_SERVICE_PORT")
//    location := "http://" + host + ":" + port+"/respond"
      location := "http://sdrespond.alamo:"+os.Getenv("SDRESPOND_SERVICE_PORT")

    fmt.Println(location)
    client := http.Client{}
    req, err := http.NewRequest("GET", location, nil)
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

func respondtorequest(c *gin.Context) {

       c.String(200, "") 


}

