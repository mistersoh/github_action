package main

import (
   "net/http"
   "fmt"
   "os"
   "time"
   "html/template"
)

type Welcome struct {
   Name string
   Time string
   Pod string
}

func getPort() string {
  p := os.Getenv("APP_PORT")
  if p != "" {
    return ":" + p
  }
  return ":8080"
}

func main() {
   welcome := Welcome{"WWSOH123", time.Now().Format(time.Stamp), os.Getenv("HOSTNAME")}

   templates := template.Must(template.ParseFiles("templates/welcome-template.html"))

   http.Handle("/static/",
      http.StripPrefix("/static/",
         http.FileServer(http.Dir("static"))))

   http.HandleFunc("/" , func(w http.ResponseWriter, r *http.Request) {

      if name := r.FormValue("name"); name != "" {
         welcome.Name = name;
      }
      if err := templates.ExecuteTemplate(w, "welcome-template.html", welcome); err != nil {
         http.Error(w, err.Error(), http.StatusInternalServerError)
      }
   })

   fmt.Println("Listening on port " + getPort());
   fmt.Println(http.ListenAndServe(getPort(), nil));
}
