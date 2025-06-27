package main

import (
	"bytes"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"github.com/adrg/frontmatter"
	"github.com/yuin/goldmark"
	highlighting "github.com/yuin/goldmark-highlighting/v2"
)

var postCache map[string]Post
var tmplPost *template.Template
var tmplIndex *template.Template
var tmplAbout *template.Template

func init() {
	tmplPost = template.Must(template.ParseFiles("templates/post.gohtml"))
	tmplIndex = template.Must(template.ParseFiles("templates/index.gohtml"))
	tmplAbout = template.Must(template.ParseFiles("templates/about.gohtml"))
}

func main() {
	postCache = LoadAllPosts()
	mux := http.NewServeMux()

	mux.HandleFunc("GET /posts/{slug}", PostHandler)
	mux.HandleFunc("GET /", IndexHandler)
	mux.HandleFunc("GET /about", AboutHandler)

	err := http.ListenAndServe(":3030", mux)
	if err != nil {
		log.Fatal(err)
	}
}

func PostHandler(w http.ResponseWriter, r *http.Request) {
	slug := r.PathValue("slug")
	post, ok := postCache[slug]
	if !ok {
		http.NotFound(w, r)
		return
	}
	err := tmplPost.Execute(w, post)
	if err != nil {
		http.Error(w, "Error executing template", http.StatusInternalServerError)
	}
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	var posts []Post
	for _, post := range postCache {
		posts = append(posts, post)
	}
	sort.Slice(posts, func(i, j int) bool {
		return posts[i].Date.After(posts[j].Date)
	})
	data := struct {
		Title       string
		Description string
		Posts       []Post
	}{
		Title:       "BinxBytes",
		Description: "Welcome to BinxBytes! Explore my latest posts.",
		Posts:       posts,
	}
	err := tmplIndex.Execute(w, data)
	if err != nil {
		http.Error(w, "Error executing template", http.StatusInternalServerError)
	}
}

func AboutHandler(w http.ResponseWriter, r *http.Request) {
	err := tmplAbout.Execute(w, nil)
	if err != nil {
		http.Error(w, "Error executing template", http.StatusInternalServerError)
	}
}

type Post struct {
	Title   string `toml:"title"`
	Slug    string `toml:"slug"`
	Content template.HTML
	Date    time.Time `toml:"date"`
	Author  Author    `toml:"author"`
}

type Author struct {
	Name  string `toml:"name"`
	Email string `toml:"email"`
}

func LoadAllPosts() map[string]Post {
	posts := make(map[string]Post)
	mdRenderer := goldmark.New(
		goldmark.WithExtensions(
			highlighting.NewHighlighting(
				highlighting.WithStyle("dracula"),
			),
		),
	)
	files, err := filepath.Glob("posts/*.md")
	if err != nil {
		log.Printf("Error globbing md files: %v", err)
		return posts
	}
	for _, file := range files {
		f, err := os.Open(file)
		if err != nil {
			log.Printf("Error opening file %s: %v", file, err)
			continue
		}
		defer f.Close()
		var post Post
		b, err := io.ReadAll(f)
		if err != nil {
			log.Printf("Error reading file %s: %v", file, err)
			continue
		}
		rest, err := frontmatter.Parse(strings.NewReader(string(b)), &post)
		if err != nil {
			log.Printf("Error parsing frontmatter in %s: %v", file, err)
			continue
		}
		if post.Slug == "" {
			post.Slug = strings.TrimSuffix(filepath.Base(file), ".md")
		}
		var buf bytes.Buffer
		err = mdRenderer.Convert(rest, &buf)
		if err != nil {
			log.Printf("Error converting markdown in %s: %v", file, err)
			continue
		}
		post.Content = template.HTML(buf.String())
		posts[post.Slug] = post
	}
	return posts
}
