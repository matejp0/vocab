package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/fatih/color"
)

type Term struct {
  question string
  answer string
}

func main() {
  var (
    location = flag.String("f", "vocabulary.txt", "Text file location")
    reverse = flag.Bool("r", false, "Reverse the questions and answers")
  )
  flag.Parse()

  terms := scanFile(location, reverse)
  used := make([]Term, 0)
  
  rand.Seed(time.Now().UnixMilli())
  fmt.Print("\033[H\033[2J") // clear unix terminal

  var (
    incorrectColor = color.New(color.FgRed)
    correctColor = color.New(color.FgGreen)
    termColor = color.New(color.FgYellow, color.Bold)
    scanner = bufio.NewScanner(os.Stdin)
    correctCount int
  )

  for len(used) != len(terms) {
    //if term := terms[rand.Intn(len(terms))]; !contains(used, term) {
    if term := terms[rand.Int() % len(terms)]; !contains(used, term) {
      used = append(used, term)
      fmt.Printf("[%d%%] %s: ", 100*len(used)/len(terms), term.question)
      scanner.Scan()
      response := scanner.Text()

      if response == term.answer {
        correctColor.Printf("✓ YES, it's %s\n", termColor.Sprint(term.answer))
        correctCount++
      } else {
        incorrectColor.Printf("✕ NO, it's %s\n", termColor.Sprint(term.answer))
      }
    } else {
      continue
    }
  }

  correctColor.Printf("Result: %d%%\n", 100*correctCount/len(terms))

}

func scanFile(location *string, reverse *bool) []Term {
  pojmy := make([]Term, 0)

  var reverseInt int8 = 0
  var reverseIntInv int8 = 1
  if *reverse {
    reverseInt = 1
    reverseIntInv = 0
  }
  
  path, _ := os.Getwd()

  file, err := os.Open(fmt.Sprintf("%s/%s", path, *location))
  if err != nil { log.Fatal(err) }
  defer file.Close() 
  scanner := bufio.NewScanner(file)

  for scanner.Scan() {
    arr := strings.Split(scanner.Text(), "=")
    pojmy = append(pojmy, Term{
      question: strings.TrimSpace(arr[reverseInt]),
      answer: strings.TrimSpace(arr[reverseIntInv]),
    })
  }

  if err := scanner.Err(); err != nil { log.Fatal(err) }
  return pojmy
}

func contains(list []Term, term Term) bool {
  for _, v := range list {
    if v == term {
      return true
    }
  }
  return false
}

// use map of struct instead
