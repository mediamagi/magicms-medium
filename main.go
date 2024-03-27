package main

import (
	"bytes"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/mediamagi/magicms/templates"
	fences "github.com/stefanfritsch/goldmark-fences"
	"github.com/yuin/goldmark"
	meta "github.com/yuin/goldmark-meta"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer/html"
	"mvdan.cc/xurls/v2"
)

func createMarkdownParser() goldmark.Markdown {
	return goldmark.New(
		goldmark.WithExtensions(
			&fences.Extender{},
			extension.GFM,
			meta.Meta,
			extension.NewLinkify(
				extension.WithLinkifyAllowedProtocols([]string{"http:", "https:"}),
				extension.WithLinkifyURLRegexp(xurls.Strict()),
			),
		),
		goldmark.WithParserOptions(
			parser.WithAutoHeadingID(),
		),
		goldmark.WithRendererOptions(
			html.WithHardWraps(),
			html.WithXHTML(),
		),
	)
}

var md = createMarkdownParser()

func main() {
	router := http.NewServeMux()

	serveStaticFiles(router)
	router.HandleFunc("/robots.txt", serveRobotsTxt)
	router.HandleFunc("/{path...}", serveContent)

	port := getPort()
	log.Printf("Server starting on port %s", port)
	if err := http.ListenAndServe(port, router); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}

func serveStaticFiles(router *http.ServeMux) {
	fs := http.FileServer(http.Dir("./static"))
	router.Handle("/static/", http.StripPrefix("/static/", fs))
}

func serveRobotsTxt(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("User-agent: *"))
}

func serveContent(w http.ResponseWriter, r *http.Request) {
	path := "./content" + r.URL.Path

	// Prevent serving files or directories starting with '_'
	if strings.HasPrefix(filepath.Base(path), "_") {
		http.NotFound(w, r)
		return
	}

	filePath, fileExtension := resolvePath(path)
	if filePath == "" {
		http.NotFound(w, r)
		return
	}

	file, err := os.ReadFile(filePath)
	if err != nil {
		log.Printf("Error reading file %s: %v", filePath, err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	if fileExtension == ".md" {
		renderMarkdown(file, w, r, filePath)
	} else {
		templates.Page(string(file), "Title", "Description").Render(r.Context(), w)
	}
}

func renderMarkdown(file []byte, w http.ResponseWriter, r *http.Request, filePath string) {
	var buf bytes.Buffer
	context := parser.NewContext()

	if err := md.Convert(file, &buf, parser.WithContext(context)); err != nil {
		log.Printf("Error converting markdown: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	metaData := meta.Get(context)
	title, _ := metaData["title"].(string)
	description, _ := metaData["description"].(string)
	relation, _ := metaData["relation"].(string)

	renderString := buf.String()

	if relation != "" {
		// Assuming 'relation' is a relative path from the "content" folder
		relationPath := filepath.Join("content", relation) // Ensure it points correctly within the "content" folder
		resolvedPath, fileExtension := resolvePath(relationPath)
		if resolvedPath != "" {
			relationFile, err := os.ReadFile(resolvedPath)
			if err != nil {
				log.Printf("Error reading relation file %s: %v", resolvedPath, err)
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				return
			}

			if fileExtension == ".md" {
				var relationBuf bytes.Buffer
				md.Convert(relationFile, &relationBuf, parser.WithContext(parser.NewContext()))
				renderString += "\n" + relationBuf.String()
			} else if fileExtension == ".html" {
				// Append HTML content directly.
				renderString += "\n" + string(relationFile)
			}
		}
	}

	templates.Page(renderString, title, description).Render(r.Context(), w)
}

func resolvePath(path string) (string, string) {
	if info, err := os.Stat(path); err == nil {
		if info.IsDir() {
			return checkIndexFiles(path)
		}
		return path, filepath.Ext(path)
	}
	return checkIfFileExtensionExists(path)
}

func checkIndexFiles(dirPath string) (string, string) {
	return checkIfFileExtensionExists(filepath.Join(dirPath, "index"))
}

func checkIfFileExtensionExists(path string) (string, string) {
	for _, ext := range []string{".html", ".md"} {
		fullPath := path + ext
		if _, err := os.Stat(fullPath); err == nil {
			return fullPath, ext
		}
	}
	return "", ""
}

func getPort() string {
	if port := os.Getenv("PORT"); port != "" {
		return ":" + port
	}
	return ":8088"
}
