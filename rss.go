package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/feeds"
)

func RSSHandler(w http.ResponseWriter, r *http.Request) {
	baseURL := "https://binx.page"

	posts := getSortedPosts()

	var created, updated time.Time
	if len(posts) > 0 {
		created = posts[len(posts)-1].Date // oldest post
		updated = posts[0].Date            // newest post
	} else {
		created = time.Now()
		updated = time.Now()
	}

	feed := &feeds.Feed{
		Title:       "BinxBytes",
		Link:        &feeds.Link{Href: baseURL + "/blog"},
		Description: "Thoughts on technology, learning, and building things.",
		Author:      &feeds.Author{Name: "Zac Bagley"},
		Created:     created,
		Updated:     updated,
		Image: &feeds.Image{
			Url:   baseURL + "/static/favicon.ico",
			Title: "BinxBytes",
			Link:  baseURL,
		},
	}

	for _, post := range posts {
		item := &feeds.Item{
			Title:       post.Title,
			Link:        &feeds.Link{Href: fmt.Sprintf("%s/blog/%s", baseURL, post.Slug)},
			Description: post.Description,
			Author:      feed.Author,
			Created:     post.Date,
			Updated:     post.Date,
			Id:          fmt.Sprintf("%s/blog/%s", baseURL, post.Slug),
			Content:     string(post.Content),
		}

		if post.Category != "" {
			item.Description = fmt.Sprintf("Category: %s\n\n%s", post.Category, post.Description)
		}

		feed.Items = append(feed.Items, item)
	}

	w.Header().Set("Content-Type", "application/rss+xml; charset=utf-8")

	rss, err := feed.ToRss()
	if err != nil {
		http.Error(w, "Failed to generate RSS feed", http.StatusInternalServerError)
		log.Printf("Error generating RSS feed: %v", err)
		return
	}

	_, err = w.Write([]byte(rss))
	if err != nil {
		log.Printf("Error writing RSS feed to response: %v", err)
	}
}
