package main

import (
  "log"
  "os"
  "text/template"
)

var tpl *template.Template

type sage struct {
  Name string
  Motto string
}

type car struct {
  Manufacturer string
  Model string
  Doors int
}



func init() {
  tpl = template.Must(template.ParseFiles("tpl.gohtml"))
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

  f := car {
    Manufacturer: "Ford",
    Model: "T",
    Doors: 4,
  }

  c := car {
    Manufacturer: "Comy",
    Model: "c",
    Doors: 2,
  }
  // store structs in a slice of structs of type sage in this case
  sages := []sage{buddha, gandhi, mlk, jesus, muhammed}
  cars := []car{f, c}

  // store slices of structs in another struct
  // use a composit literal instead of declaring a new literal
  // this format uses anonymous struct
  data := struct{
    Wisdom []sage
    Transport []car
  }{
    sages,
    cars,
  }


  err := tpl.Execute(os.Stdout, data)
  if err != nil {
    log.Fatalln(err)
  }
}
