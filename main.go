package main

import (
	"bytes"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/adrg/frontmatter"
	"github.com/yuin/goldmark"
	highlighting "github.com/yuin/goldmark-highlighting/v2"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /posts/{slug}", PostHandler(FileReader{}))
	mux.HandleFunc("GET /", IndexHandler(LoadAllPosts()))
	mux.HandleFunc("GET /about", AboutHandler)

	err := http.ListenAndServe(":3030", mux)
	if err != nil {
		log.Fatal(err)
	}
}

type SlugReader interface {
	Read(slug string) (string, error)
}

type FileReader struct{}

func (fsr FileReader) Read(slug string) (string, error) {
	f, err := os.Open("posts/" + slug + ".md")
	if err != nil {
		return "", err
	}
	defer f.Close()
	b, err := io.ReadAll(f)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func PostHandler(sl SlugReader) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var post Post
		post.Slug = r.PathValue("slug")
		postMarkdown, err := sl.Read(post.Slug)
		if err != nil {
			// TODO: Handle different errors in the future
			http.Error(w, "Post not found", http.StatusNotFound)
			return
		}
		rest, err := frontmatter.Parse(strings.NewReader(postMarkdown), &post)
		if err != nil {
			http.Error(w, "Error parsing frontmatter", http.StatusInternalServerError)
			return
		}
		mdRenderer := goldmark.New(
			goldmark.WithExtensions(
				highlighting.NewHighlighting(
					highlighting.WithStyle("dracula"),
				),
			),
		)
		var buf bytes.Buffer
		err = mdRenderer.Convert(rest, &buf)
		if err != nil {
			http.Error(w, "Error converting markdown", http.StatusInternalServerError)
			return
		}
		// TODO: Parse the template once, not every page load.
		tpl, err := template.ParseFiles("templates/post.gohtml")
		if err != nil {
			http.Error(w, "Error parsing template", http.StatusInternalServerError)
			return
		}
		post.Content = template.HTML(buf.String())
		err = tpl.Execute(w, post)
	}
}

func IndexHandler(posts []Post) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tpl, err := template.ParseFiles("templates/index.gohtml")
		if err != nil {
			http.Error(w, "Error parsing template", http.StatusInternalServerError)
			return
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
		tpl.Execute(w, data)
	}
}

func AboutHandler(w http.ResponseWriter, r *http.Request) {
	tpl, err := template.ParseFiles("templates/about.gohtml")
	if err != nil {
		http.Error(w, "Error parsing template", http.StatusInternalServerError)
		return
	}
	tpl.Execute(w, nil)
}

type Post struct {
	Title   string `toml:"title"`
	Slug    string `toml:"slug"`
	Content template.HTML
	Author  Author `toml:"author"`
}

type Author struct {
	Name  string `toml:"name"`
	Email string `toml:"email"`
}

// type PostData struct {
// 	Title   string
// 	Content template.HTML
// 	Author  string
// }

// LoadAllPosts reads all .md files in the posts directory, parses their frontmatter, and returns a slice of Post structs.
func LoadAllPosts() []Post {
	var posts []Post
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
		var post Post
		b, err := io.ReadAll(f)
		f.Close()
		if err != nil {
			log.Printf("Error reading file %s: %v", file, err)
			continue
		}
		_, err = frontmatter.Parse(strings.NewReader(string(b)), &post)
		if err != nil {
			log.Printf("Error parsing frontmatter in %s: %v", file, err)
			continue
		}
		// Set the slug based on the filename (without .md extension) if not present
		if post.Slug == "" {
			post.Slug = strings.TrimSuffix(filepath.Base(file), ".md")
		}
		posts = append(posts, post)
	}
	return posts
}
