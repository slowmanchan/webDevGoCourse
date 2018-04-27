package main

import (
  "log"
  "text/template"
  "os"
  "strings"
)

var tpl *template.Template

type sage struct {
  Name string
  Motto string
}

var fm = template.FuncMap{
  "uc": strings.ToUpper,
  "ft": firstThree,
}

func init() {
  tpl = template.Must(template.New("").Funcs(fm).ParseFiles("tpl.gohtml"))
}

func firstThree(s string) string {
  s = strings.TrimSpace(s)
  s = s[:3]
  return s
}

func main() {
  buddha := sage {
    Name: "Buddha",
    Motto: "The belief of no beliefs",
  }

  gandhi := sage {
    Name: "Gandhi",
    Motto: "Be the change",
  }

  mlk := sage {
    Name: "Martin Luther King",
    Motto: "Hatred never ceases with hatred",
  }

  jesus := sage {
    Name: "Jesus",
    Motto: "Love all",
  }

  muhammed := sage {
    Name: "Muhammed",
    Motto: "To overcome evil with good is good",
  }
  sages := []sage{mlk, jesus, muhammed, buddha, gandhi}

  err := tpl.ExecuteTemplate(os.Stdout,"tpl.gohtml", sages)

  if err != nil {
    log.Fatalln(err)
  }
}
