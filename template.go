package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
  "strings"
)

func Render(directory string, w http.ResponseWriter, r *http.Request) {

	content_path := filepath.Join(directory, "content/")
	config_path := filepath.Join(directory, "config/")

	header_path := filepath.Join(config_path, "header.html")
	footer_path := filepath.Join(config_path, "footer.html")

	header, err := os.Open(header_path)

	if err != nil {
		fmt.Fprintln(w, "error:", err.Error())
		return
	}

	footer, err := os.Open(footer_path)

	if err != nil {
		fmt.Fprintln(w, "error:", err.Error())
		return
	}

	io.Copy(w, header)
	var file string
	var uri string = r.URL.RequestURI()
	uri = "/index"
	file = fmt.Sprintf("%s%s", filepath.Base(uri), ".template")

	fmt.Fprintln(w, RenderMatch(r,content_path, file))

	io.Copy(w, footer)

}

func RenderMatch(r *http.Request, dir string, filename string) string {

  if strings.HasPrefix(filename,"key:") {
	  var uri string = r.URL.RequestURI()
    return uri
  }
	content_bytes, err := ioutil.ReadFile(filepath.Join(dir, filename))
	content := string(content_bytes)

	if err != nil {
		return fmt.Sprintf("<<Error::%s>>", err.Error())
	}

	re := regexp.MustCompile("{{.*}}")

	replaceFunc := func(match string) string {
		match = match[2 : len(match)-2]
		return RenderMatch(r,dir, match)
	}

	return re.ReplaceAllStringFunc(content, replaceFunc)
}
