package rsg

import (
	"bufio"
	"encoding/json"
	"log"
	"net/http"
)


type feed struct {
  handlers []HandlerFunc

  newhandlers chan HandlerFunc
  comments chan *Comment
}
func NewFeed() (f *feed) {
  f = &feed{
    handlers: make([]HandlerFunc, 0),
    newhandlers: make(chan HandlerFunc),
    comments: make(chan *Comment),
  }

  go f.start()
  go f.listen()

  return
}

/* Example:
{
  "author_flair_text":null,
  "author":"Ninja-512",
  "parent_id":"t1_cdl9fux",
  "link_id":"t3_1rajhm",
  "author_flair_css_class":null,
  "body":"I've never played on hardcore.  This will be fun and I'm excited!",
  "id":"cdla5wi",
  "created_utc":1385230736,
  "subreddit_id":"t5_2qnzu",
  "link_title":"Fun FNV playthroughs? (DON'T UPVOTE)",
  "created":1385259536,
  "subreddit":"Fallout"
}
*/
type Comment struct {
  AuthorFlairtext     string `json:"author_flair_text"`
  ParentId            string `json:"parent_id"`
  LinkId              string `json:"link_id"`
  Body                string `json:"body"`
  SubredditId         string `json:"subreddit_id"`
  Created             uint64 `json:"created"`
  Subreddit           string `json:"subreddit"`
  Author              string `json:"author"`
  AuthorFlairCssClass string `json:"author_flair_css_class"`
  Id                  string `json:"id"`
  CreatedUtc          uint64 `json:"created_utc"`
  LinkTitle           string `json:"link_title"`
}

type HandlerFunc func(c *Comment)
func (f *feed) Handle(fu HandlerFunc) {
  f.newhandlers <- fu
}

func (f *feed) start() {
  for {
    select {
    case h := <-f.newhandlers:
      f.handlers = append(f.handlers, h)
    case c := <-f.comments:
      for _, h := range f.handlers {
        go h(c)
      }

    }
  }
}

func (f *feed) listen() {
	log.Println("Starting reading feed")

	feed := `http://stream.redditanalytics.com`
	resp, err := http.Get(feed)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	scanner := bufio.NewScanner(resp.Body)
	for scanner.Scan() {
		if err := scanner.Err(); err != nil {
			log.Println("Error reading line: ", err)
			continue
		}

		line := scanner.Text()

		var c Comment
		err = json.Unmarshal([]byte(line), &c)
		if err != nil {
			log.Println("Error when parsing to json: ", err)
		}

    f.comments <- &c
	}
}
