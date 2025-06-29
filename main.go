package main

import (
	"bytes"
	"flag"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"github.com/akrylysov/algnhsa"

	"github.com/yuin/goldmark"
	highlighting "github.com/yuin/goldmark-highlighting/v2"
	meta "github.com/yuin/goldmark-meta"
	"github.com/yuin/goldmark/parser"
)

var postCache map[string]Post
var tmplPost *template.Template
var tmplIndex *template.Template
var tmplAbout *template.Template
var tmplBlog *template.Template
var baseDir string
var devMode bool

func init() {
	flag.BoolVar(&devMode, "dev", false, "Run in development mode (use relative paths)")
	flag.Parse()

	if devMode {
		baseDir = "."
	} else {
		ex, _ := os.Executable()
		baseDir = filepath.Dir(ex)
	}

	tmplPost = template.Must(template.ParseFiles(filepath.Join(baseDir, "templates/post.gohtml")))
	tmplIndex = template.Must(template.ParseFiles(filepath.Join(baseDir, "templates/index.gohtml")))
	tmplAbout = template.Must(template.ParseFiles(filepath.Join(baseDir, "templates/about.gohtml")))
	tmplBlog = template.Must(template.ParseFiles(filepath.Join(baseDir, "templates/blog.gohtml")))
}

func main() {
	postCache = LoadAllPosts()
	mux := http.NewServeMux()

	mux.HandleFunc("GET /blog/{slug}", PostHandler)
	mux.HandleFunc("GET /", IndexHandler)
	mux.HandleFunc("GET /about", AboutHandler)
	mux.HandleFunc("GET /blog", BlogHandler)
	mux.HandleFunc("GET /favicon.ico", GetFavicon)

	fs := http.FileServer(http.Dir(filepath.Join(baseDir, "static")))
	mux.Handle("GET /static/", http.StripPrefix("/static/", fs))

	// -dev flag for local, otherwise needs algnhsa for lambda
	if devMode {
		port := os.Getenv("PORT")
		if port == "" {
			port = "3000"
		}
		log.Printf("Local dev on :%s", port)
		log.Fatal(http.ListenAndServe(":"+port, mux))
	} else {
		log.Printf("Running with algnhsa")
		algnhsa.ListenAndServe(mux, nil)
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

	// recent posts
	if len(posts) > 3 {
		posts = posts[:3]
	}

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

func BlogHandler(w http.ResponseWriter, r *http.Request) {
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
		Title:       "Blog | BinxBytes",
		Description: "Thoughts, tutorials, and explorations in technology and development.",
		Posts:       posts,
	}
	err := tmplBlog.Execute(w, data)
	if err != nil {
		http.Error(w, "Error executing template", http.StatusInternalServerError)
	}
}

func GetFavicon(w http.ResponseWriter, r *http.Request) {
	faviconPath := filepath.Join(baseDir, "static", "favicon.ico")
	http.ServeFile(w, r, faviconPath)
}

type Post struct {
	Title       string `toml:"title"`
	Slug        string `toml:"slug"`
	Description string `toml:"description"`
	Category    string `toml:"category"`
	Content     template.HTML
	Date        time.Time `toml:"date"`
	Author      Author    `toml:"author"`
}

func (p Post) FormattedDate() string {
	if p.Date.IsZero() {
		return ""
	}
	return p.Date.Format("January 2, 2006")
}

type Author struct {
	Name  string `toml:"name"`
	Email string `toml:"email"`
}

func LoadAllPosts() map[string]Post {
	posts := make(map[string]Post)
	mdRenderer := goldmark.New(
		goldmark.WithExtensions(
			meta.Meta,
			highlighting.NewHighlighting(
				highlighting.WithStyle("dracula"),
			),
		),
	)
	postsDir := filepath.Join(baseDir, "blog")
	files, err := filepath.Glob(filepath.Join(postsDir, "*.md"))
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

		b, err := io.ReadAll(f)
		if err != nil {
			log.Printf("Error reading file %s: %v", file, err)
			continue
		}

		context := parser.NewContext()
		var buf bytes.Buffer
		err = mdRenderer.Convert(b, &buf, parser.WithContext(context))
		if err != nil {
			log.Printf("Error converting markdown in %s: %v", file, err)
			continue
		}

		post := Post{}
		if metaData := meta.Get(context); metaData != nil {
			if title, ok := metaData["title"]; ok {
				if titleStr, ok := title.(string); ok {
					post.Title = titleStr
				}
			}
			if slug, ok := metaData["slug"]; ok {
				if slugStr, ok := slug.(string); ok {
					post.Slug = slugStr
				}
			}
			if dateStr, ok := metaData["date"]; ok {
				if date, ok := dateStr.(string); ok {
					if parsedDate, err := time.Parse("2006-01-02", date); err == nil {
						post.Date = parsedDate
					}
				}
			}
			if author, ok := metaData["author"]; ok {
				if authorMap, ok := author.(map[string]interface{}); ok {
					if name, ok := authorMap["name"]; ok {
						if nameStr, ok := name.(string); ok {
							post.Author.Name = nameStr
						}
					}
					if email, ok := authorMap["email"]; ok {
						if emailStr, ok := email.(string); ok {
							post.Author.Email = emailStr
						}
					}
				} else if authorStr, ok := author.(string); ok {
					post.Author.Name = authorStr
				}
			}
			if email, ok := metaData["email"]; ok {
				if emailStr, ok := email.(string); ok {
					post.Author.Email = emailStr
				}
			}
			if description, ok := metaData["description"]; ok {
				if descriptionStr, ok := description.(string); ok {
					post.Description = descriptionStr
				}
			}
			if category, ok := metaData["category"]; ok {
				if categoryStr, ok := category.(string); ok {
					post.Category = categoryStr
				}
			}
		}

		if post.Slug == "" {
			post.Slug = strings.TrimSuffix(filepath.Base(file), ".md")
		}

		post.Content = template.HTML(buf.String())
		posts[post.Slug] = post
	}
	return posts
}
