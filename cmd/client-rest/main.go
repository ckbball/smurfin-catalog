package main

import (
  "encoding/json"
  "flag"
  "fmt"
  "io/ioutil"
  "log"
  "net/http"
  "strings"
  "time"
)

func main() {
  // get configuration
  address := flag.String("server", "http://localhost:8080", "HTTP gateway url, e.g. http://localhost:8080")
  flag.Parse()

  t := time.Now().In(time.UTC)
  pfx := t.Format(time.RFC3339Nano)

  var body string

  // Call Create
  resp, err := http.Post(*address+"/v1/catalogs", "application/json", strings.NewReader(fmt.Sprintf(`
    {
      "api":"v1",
      "item": {
        "vendor_id":"3",
        "blue_essence":10000,
        "solo":3
      }
    }
  `, pfx, pfx, pfx)))
  if err != nil {
    log.Fatalf("failed to call Create method: %v", err)
  }
  bodyBytes, err := ioutil.ReadAll(resp.Body)
  resp.Body.Close()
  if err != nil {
    body = fmt.Sprintf("failed read Create response body: %v", err)
  } else {
    body = string(bodyBytes)
  }
  log.Printf("Create response: Code=%d, Body=%s\n\n", resp.StatusCode, body)

  // parse ID of created ToDo
  var created struct {
    API string `json:"api"`
    ID  string `json:"id"`
  }
  err = json.Unmarshal(bodyBytes, &created)
  if err != nil {
    log.Fatalf("failed to unmarshal JSON response of Create method: %v", err)
    fmt.Println("error:", err)
  }

  // Call Search
  resp, err = http.Post(*address+"/v1/catalogs/search", "application/json", strings.NewReader(fmt.Sprintf(`
    {
      "api":"v1",
      "solo": 3,
      "flex": 2
    }
  `, pfx, pfx, pfx)))
  if err != nil {
    log.Fatalf("failed to call Create method: %v", err)
  }
  bodyBytes, err = ioutil.ReadAll(resp.Body)
  resp.Body.Close()
  if err != nil {
    body = fmt.Sprintf("failed read ReadAll response body: %v", err)
  } else {
    body = string(bodyBytes)
  }
  log.Printf("ReadAll response: Code=%d, Body=%s\n\n", resp.StatusCode, body)

  // Call ReadAll
  resp, err = http.Post(*address+"/v1/catalogs/search", "application/json", strings.NewReader(fmt.Sprintf(`
    {
      "api":"v1",
      "solo": 3,
      "flex": 2
    }
  `, pfx, pfx, pfx)))
  bodyBytes, err = ioutil.ReadAll(resp.Body)
  resp.Body.Close()
  if err != nil {
    body = fmt.Sprintf("failed read ReadAll response body: %v", err)
  } else {
    body = string(bodyBytes)
  }
  log.Printf("ReadAll response: Code=%d, Body=%s\n\n", resp.StatusCode, body)

  // Call Delete
  req, err = http.NewRequest("DELETE", fmt.Sprintf("%s%s/%s", *address, "/v1/catalogs", created.ID), nil)
  resp, err = http.DefaultClient.Do(req)
  if err != nil {
    log.Fatalf("failed to call Delete method: %v", err)
  }
  bodyBytes, err = ioutil.ReadAll(resp.Body)
  resp.Body.Close()
  if err != nil {
    body = fmt.Sprintf("failed read Delete response body: %v", err)
  } else {
    body = string(bodyBytes)
  }
  log.Printf("Delete response: Code=%d, Body=%s\n\n", resp.StatusCode, body)
}
