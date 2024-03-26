package main

import (
	"bytes"
	"log"
	"net/http"
	"os"
	"path/filepath"

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
		renderMarkdown(file, w, r)
	} else {
		templates.Page(string(file), "Title", "Description").Render(r.Context(), w)
	}
}

func renderMarkdown(file []byte, w http.ResponseWriter, r *http.Request) {
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

	templates.Page(buf.String(), title, description).Render(r.Context(), w)
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
		if _, err := os.Stat(path + ext); err == nil {
			return path + ext, ext
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
