package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"

    "encoding/csv"
    "os"
)

type Json struct {
    ID              string  `json:"ID"`
    Fixedacidity    string `json:"Fixedacidity"`
    Volatileacidity    string `json:"Volatileacidity"`
      Citricacid    string `json:"Citricacid"`
    Residualsugar    string `json:"Residualsugar"`
      Chlorides    string `json:"Chlorides"`
    Freesulfurdioxide    string `json:"Freesulfurdioxide"`
     Totalsulfurdioxide    string `json:"Totalsulfurdioxide"`
    Density    string `json:"Density"`
      PH    string `json:"PH"`
    Sulphates    string `json:"Sulphates"`
      Alcohol    string `json:"Alcohol"`
    Quality     string `json:"Quality "`
}

func main() {
    fmt.Println("Starting the application...")
    
    // leyendo data del servidor
    data, err := http.Get("http://localhost:3000/gophers")
    if err != nil {
        fmt.Printf("The HTTP request failed with error %s\n", err)
    } else {
        data, _ := ioutil.ReadAll(data.Body)
        fmt.Println(string(data))
    }

    body, readErr := ioutil.ReadAll(data.Body)
    if readErr != nil {
 
    }
    // Unmarshal JSON data
    var d []Json
    err = json.Unmarshal([]byte(string(body)), &d)
    if err != nil {
        fmt.Println(err)
    }
    // Create a csv file
    f, err := os.Create("datasets/wineapi.csv")
    if err != nil {
        fmt.Println(err)
    }
    defer f.Close()
    // Write Unmarshaled json data to CSV file
    w := csv.NewWriter(f)
    for _, obj := range d {
        var record []string
        record = append(record, obj.ID)
        record = append(record, obj.Fixedacidity)
        record = append(record, obj.Volatileacidity)
         record = append(record, obj.Citricacid)
        record = append(record, obj.Residualsugar)
         record = append(record, obj.Chlorides)
        record = append(record, obj.Freesulfurdioxide)
         record = append(record, obj.Totalsulfurdioxide)
        record = append(record, obj.Density)
         record = append(record, obj.PH)
        record = append(record, obj.Sulphates)
         record = append(record, obj.Sulphates)
        record = append(record, "9.0000")
        w.Write(record)
    }
    w.Flush()

     fmt.Println("Terminating the application...")
}



 