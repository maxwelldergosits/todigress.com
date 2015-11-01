package main
import "strings"
import "path/filepath"
import "log"


func RenderMap(r * RenderContext, match string, render func(*RenderContext, string)string) string {
  log.Println("mapping:",match)
  content := strings.TrimPrefix(match, "map:")
  tokens := strings.SplitN(content, ":",2)
  if (len(tokens) <2) {
    return "Error: improperly formatted map: "+content
  }
  file := tokens[0]
  key := tokens[1]
  values := utils.ReadMap(filepath.Join(r.Directory,file))
  log.Println("key value=:",key,values[key])
  return values[key]
}

