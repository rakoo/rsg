package main

// A simple cli app that will list all contributions

import (
  "fmt"

  "github.com/rakoo/rsg"
)

func main() {
  feed := rsg.NewFeed()
  feed.Handle(handleComment)
  select {}
}

func handleComment(c *rsg.Comment) {
  end := len(c.Body) - 1
  suffix := ""

  if end > 30 {
    end = 30
    suffix = "[...]"
  }

  fmt.Printf("%25s [%20s] %q\n", "/r/"+c.Subreddit, c.Author, c.Body[:end] + suffix)
}
