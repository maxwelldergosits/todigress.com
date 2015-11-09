package main

import (
	"fmt"
	"io/ioutil"
  "io"
	"net/http"
	"path/filepath"
	"regexp"
)

type Templater struct {
  ReplaceFunc func(*RenderContext, string ,func(*RenderContext,string)string)string
  Regex string
}

type RenderContext struct {
  Resolver func(string,string)string
  FindFiles func(string,string)[]string
  Templaters []Templater
  Directory string
  URI string
}

func RenderTemplate(r * RenderContext, templateName string, render func(*RenderContext,string)string) string {
  templateContents := r.Resolver(r.Directory,templateName)
  return render(r,templateContents)
}

func RenderURI(r * RenderContext, conent string, render func(*RenderContext,string)string) string {
  return render(r,"{{"+r.URI+".template}}")
}
func RenderString(r * RenderContext, content string) string {

  for _,templater := range r.Templaters {
    re := regexp.MustCompile(templater.Regex)

    replaceFunc := func(match string) string {
		  match = match[2 : len(match)-2]
      return templater.ReplaceFunc(r,match,RenderString)
    }

    content = re.ReplaceAllStringFunc(content, replaceFunc)
  }

  return content
}

func DefaultResolver(directory, file string)string {
    bytes, err := ioutil.ReadFile(filepath.Join(directory,file))
    if err != nil {
      return err.Error()
    }
    return string(bytes)
}

func DefaultFileFinder(directory, regex string)[]string {
    if (directory == "") {
      directory = "."
    }
    re := regexp.MustCompile(regex)
    files, err := ioutil.ReadDir(directory)
    if err != nil {
      return []string{err.Error()}
    }
    names := []string{}
    for _,f:= range files {
      if re.MatchString(f.Name()) {
        names = append(names, f.Name())
      }
    }
    return names
}

func Render(directory string, w io.Writer, r *http.Request) {

	content_path := filepath.Join(directory, "content/")

	var file string
	var uri string = r.URL.RequestURI()
  if (uri == "/index") {
    uri = "about"
  }
	file = "{{index.header}}"

  var rc RenderContext
  rc.Resolver = DefaultResolver
  rc.FindFiles = DefaultFileFinder
  rc.URI = uri

  rc.Templaters = []Templater{{RenderList, "{{list:.*}}"},
                              {RenderTemplate, "{{.*\\.(template|header)}}"},
                              {RenderURI, "{{uri}}"},
                              {RenderMap, "{{map:.*?}}"},
                             }
  rc.Directory = content_path

	fmt.Fprintln(w, RenderString(&rc,file))

}

