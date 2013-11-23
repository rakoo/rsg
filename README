# RSG
RedditStream-Go is a library to parse
[redditanalytic](http://www.reddit.com/r/redditdev/comments/1oc5dt/redditanalytics_progress_and_an_update_on_the/)'s
continuous stream of comments.

# Installation
The go usual:

```sh
$ go get github.com/rakoo/rsg
```

# Usage
Have a look at the `examples` folder. Here's the full listing of
`cli.go`:

```go
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
```

Here's the result:

![No luck for you :(](http://pix.toile-libre.org/upload/original/1385242816.png)
