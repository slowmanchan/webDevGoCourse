package main

import (
  "log"
  "os"
  "text/template"
)

type course struct {
  Number, Name, Units string
}

type semester struct {
  Term string
  Courses []course
}

type year struct {
  Fall, Spring, Summer semester
}

var tpl *template.Template

func init() {
  tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
  y := year{
    Fall: semester{
      Term: "Fall",
      Courses: []course{
        course{"CSCI-40", "INtro to computers", "4"},
        course{"CSCI-130", "Intro to bugs", "5"},
        course{"CSCI_333", "Intro to life", "33"},
      },
    },
    Spring: semester{
      Term: "Spring",
      Courses: []course{
        course{"CSCI-50", "Advanced Go", "5"},
				course{"CSCI-190", "Advanced Web Programming with Go", "5"},
				course{"CSCI-191", "Advanced Mobile Apps With Go", "5"},
      },
    },
  }

  err := tpl.Execute(os.Stdout, y)

  if err != nil {
    log.Fatalln(err)
  }
}
