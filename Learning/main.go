package main

import (
	"html/template"
	"net/http"

	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
)

var viewTemplate = template.Must(template.New("view").Parse(`
<!DOCTYPE html>
<html>
<head>
    <title>Record</title>
</head>
<body>
    <h1>{{.Title}}</h1>
    <p>{{.Content}}</p>
</body>
</html>
`))

type Record struct {
	Title   string
	Content string
}

func init() {
	http.HandleFunc("/view", viewRecord)
}

func viewRecord(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	key := datastore.NewKey(c, "Record", r.FormValue("id"), 0, nil)
	record := new(Record)
	if err := datastore.Get(c, key, record); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err := viewTemplate.Execute(w, record); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	appengine.Main()
}
