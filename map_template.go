package main
import "strings"
import "path/filepath"


func RenderMap(r * RenderContext, match string, render func(*RenderContext, string)string) string {
  content := strings.TrimPrefix(match, "map:")
  tokens := strings.SplitN(content, ":",2)
  if (len(tokens) <2) {
    return "Error: improperly formatted map: "+content
  }
  file := tokens[0]
  key := tokens[1]
  values := ReadMap(filepath.Join(r.Directory,file))
  return values[key]
}

