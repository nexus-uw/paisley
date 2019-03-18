package controllers

import (
	"bytes"
	"log"
	"math"
	"os"
	"strings"
	"time"

	"github.com/nexus-uw/paisley/app/models"

	"github.com/revel/revel"

	"text/template"

	"github.com/gorilla/feeds"
	"github.com/turnage/graw/reddit"
)

type Feed struct {
	Application
}

func createBody(post *reddit.Post) string {
	check := func(err error) {
		if err != nil {
			log.Fatal(err)
		}
	}
	funcMap := template.FuncMap{

		// TODO: do the same but for video

		// check if given link is an image
		"isImageLink": func(link string) bool {
			for _, ext := range []string{
				".jpg", ".jpeg", ".jpe", ".jif", ".jfif", ".jfi", // https://en.wikipedia.org/wiki/JPEG
				".gif", ".png", ".apng", ".svg", ".bmp", ".dib", ".ico",
			} {
				if strings.HasSuffix(link, ext) {
					return true
				}
			}
			return false
		},
	}
	out := new(bytes.Buffer)
	t, err := template.New("foo").Funcs(funcMap).Parse(`
	<!DOCTYPE html>
	<body>
		
		{{if ( isImageLink .URL) }}
		<img  src="{{.URL}}" alt="post image"/>
		{{else if .URL}}
		{{if .Thumbnail}}
		<img src="{{.Thumbnail}}" alt="thumbnail"/>
		{{end}}
		<a href="{{.URL}}"> Link </a>
		{{end}}
		
		
		<div>
			{{.SelfText}}
		</div>
	</body>
	</html>
	`)

	check(err)
	err = t.Execute(out, post)
	check(err)
	return out.String()
}

// ff
func (c Feed) GetFeed(SubscriptionID string) revel.Result {
	c.Log.Info("Fetching index")

	cfg := reddit.BotConfig{
		Agent: "graw:doc_demo_bot:0.3.1 by /u/ammobin",
		// Your registered app info from following:
		// https://github.com/reddit/reddit/wiki/OAuth2
		App: reddit.App{
			ID:       os.Getenv("REDDIT_ID"),
			Secret:   os.Getenv("REDDIT_SECRET"),
			Username: os.Getenv("REDDIT_USER"),
			Password: os.Getenv("REDDIT_PASSWORD"),
		},
	}

	obj, err := c.Db.Get(models.Subscription{}, SubscriptionID)
	Subscription := obj.(*models.Subscription)
	if err != nil {
		panic(err)
	}

	bot, _ := reddit.NewBot(cfg)
	subredditName := "/r/" + Subscription.Subredit
	harvest, err := bot.Listing(subredditName, "")
	if err != nil {
		return c.RenderError(err)
	}

	
	now := time.Now()
	feed := &feeds.Feed{
		Title:       subredditName,
		Link:        &feeds.Link{Href: "https://reddit.com" + subredditName},
		Description: "rss feed generated by Paisley for " + subredditName,
		Author:      &feeds.Author{Name: "Pasiley", Email: "lol@example.com"},
		Created:     now,
	}

	for _, post := range harvest.Posts[:] {

		var postVoteRatio = 100
		if post.Downs > 0 {
			postVoteRatio = int(math.Round((float64(post.Ups) / float64(post.Downs)) * 1000))
		}
		if post.Deleted ||
			Subscription.MinVoteRatio < postVoteRatio ||
			int32(Subscription.MinVoteCount) < (post.Ups+post.Downs) ||
			int32(Subscription.MinNumberOfComments) < post.NumComments {
			continue
		}

		// things to do
		// put link in description
		// if link is to media asset -> create image tag for description
		// if media thumbnail is present -> add to body
		// if include comments it turned on, append to body => https://golang.org/pkg/html/template/
		feed.Items = append(feed.Items,
			&feeds.Item{
				Title:       post.Title,
				Link:        &feeds.Link{Href: "https://reddit.com" + post.Permalink},
				Description: createBody(post),
				Author:      &feeds.Author{Name: post.Author},
				Created:     time.Unix(int64(post.CreatedUTC), 0),
			},
		)
	}

	rss, err := feed.ToRss()
	if err != nil {
		log.Fatal(err)
	}
	c.Response.ContentType = "text/xml; charset=utf-8"
	return c.RenderText(rss)
}
