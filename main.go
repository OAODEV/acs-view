package main

import(
  "encoding/json"
  "fmt"
  "io/ioutil"
  "log"
  "net/http"
  "os"
  "strings"
  )

type Member struct {
  Name string
  Score int
}

func indexHandler( w http.ResponseWriter, r *http.Request) {
  api_uri := os.Getenv("api_uri")
  resp, err := http.Get(api_uri)
  if err != nil {
    log.Fatal(err)
  }
  defer resp.Body.Close()
  body, err := ioutil.ReadAll(resp.Body)
  if err != nil {
    log.Fatal(err)
  }

  var members []Member
  err = json.Unmarshal(body, &members)

  var names []string
  total := 0
  for _, member := range members {
    names = append(names, member.Name)
    total += member.Score
  }

  name_string := strings.Join(names, ", ")

  fmt.Fprintf(w, "<h1>Team members</h1><p>%s</p><h2>Total Score</h2><p>%d</p>",
              name_string, total)
}

func main() {
  http.HandleFunc("/", indexHandler)
  http.ListenAndServe("0.0.0.0:8080",nil)
}
