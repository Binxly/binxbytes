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

// NOTE: throwing blog posts in memory - just load once at startup
// LoadAllPosts() is called, stores generated html in postCache
// https://go.dev/blog/maps - look up by slug
var postCache map[string]Post

var tmplPost *template.Template
var tmplIndex *template.Template
var tmplAbout *template.Template
var tmplBlog *template.Template
var tmplContact *template.Template
var baseDir string
var devMode bool

func init() {
	flag.BoolVar(&devMode, "dev", false, "Run in development mode (use relative paths)")
	flag.Parse()

	// NOTE: use relative paths when the -dev flag is set
	// otherwise, uses absolute paths, which may not work in dev
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
	tmplContact = template.Must(template.ParseFiles(filepath.Join(baseDir, "templates/contact.gohtml")))
}

func main() {
	postCache = LoadAllPosts()
	mux := http.NewServeMux()

	mux.HandleFunc("GET /blog/{slug}", PostHandler)
	mux.HandleFunc("GET /", IndexHandler)
	mux.HandleFunc("GET /about", AboutHandler)
	mux.HandleFunc("GET /blog", BlogHandler)
	mux.HandleFunc("GET /contact", ContactHandler)
	mux.HandleFunc("GET /favicon.ico", GetFavicon)
	mux.HandleFunc("GET /rss", RSSHandler)

	fs := http.FileServer(http.Dir(filepath.Join(baseDir, "static")))
	mux.Handle("GET /static/", http.StripPrefix("/static/", fs))

	// NOTE: -dev flag for local development only
	// without it, routes rely on algnhsa's Lambda adapter
	if devMode {
		port := os.Getenv("PORT")
		if port == "" {
			port = "3000"
		}
		log.Printf("Local dev on http://localhost:%s", port)
		log.Fatal(http.ListenAndServe(":"+port, mux))
	} else {
		algnhsa.ListenAndServe(mux, nil)
	}
}

func PostHandler(w http.ResponseWriter, r *http.Request) {
	slug := r.PathValue("slug")

	// look up posts by slug in postCache map
	post, ok := postCache[slug]
	if !ok {
		http.NotFound(w, r)
		return
	}

	// execute post.gohtml template
	err := tmplPost.Execute(w, post)
	if err != nil {
		http.Error(w, "Error executing template", http.StatusInternalServerError)
	}
}

// getSortedPosts returns a slice of all post structs sorted by date,
// instead of sorting the slice in both IndexHandler and BlogHandler
func getSortedPosts() []Post {
	var posts []Post
	for _, post := range postCache {
		posts = append(posts, post)
	}
	sort.Slice(posts, func(i, j int) bool {
		return posts[i].Date.After(posts[j].Date)
	})
	return posts
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	posts := getSortedPosts()
	if len(posts) > 3 {
		posts = posts[:3]
	}

	data := struct {
		Title       string
		Description string
		Posts       []Post
	}{
		Title:       "BinxBytes",
		Description: "Another dev blog on the internet.",
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
	posts := getSortedPosts()

	data := struct {
		Title       string
		Description string
		Posts       []Post
	}{
		Title:       "Blog | BinxBytes",
		Description: "Thoughts on technology, learning, and building things.",
		Posts:       posts,
	}
	err := tmplBlog.Execute(w, data)
	if err != nil {
		http.Error(w, "Error executing template", http.StatusInternalServerError)
	}
}

func ContactHandler(w http.ResponseWriter, r *http.Request) {
	err := tmplContact.Execute(w, nil)
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
	// TODO: set default date to file mod time, or log a warning
	// https://pkg.go.dev/os#FileInfo.ModTime
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

	// goldmark renderer configuration
	mdRenderer := goldmark.New(
		goldmark.WithExtensions(
			meta.Meta, // enables frontmatter parsing
			highlighting.NewHighlighting(
				highlighting.WithStyle("dracula"),
				// Supported styles defined under
				// https://github.com/alecthomas/chroma/tree/master/formatters.
			),
		),
	)

	// NOTE: naive implementation, zipping all posts in blog/
	// with the binary for now
	// TODO: look at S3 for this and stylesheets?
	postsDir := filepath.Join(baseDir, "blog")

	// filepath.Glob finds all .md files in the blog directory
	// returns a slice of file paths
	// https://pkg.go.dev/path/filepath#Glob
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

		// NOTE: getting metadata from the markdown file's frontmatter
		// https://github.com/yuin/goldmark-meta
		// meta.Get(context) returns a map[string]interface{} of frontmatter metadata
		// look up the key in that map, check if ok,
		// type assert to convert interface{} to string, ok check,
		// assign to post struct
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
			// NOTE: author is unique because it links to the author's email
			// using <a href="mailto:..."> over the name in the template
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

			// NOTE: excerpt
			if description, ok := metaData["description"]; ok {
				if descriptionStr, ok := description.(string); ok {
					post.Description = descriptionStr
				}
			}

			// NOTE: tags
			if category, ok := metaData["category"]; ok {
				if categoryStr, ok := category.(string); ok {
					post.Category = categoryStr
				}
			}
		}

		// NOTE: no slug in frontmatter, use filename instead
		if post.Slug == "" {
			post.Slug = strings.TrimSuffix(filepath.Base(file), ".md")
		}

		post.Content = template.HTML(buf.String())
		posts[post.Slug] = post
	}
	return posts
}
