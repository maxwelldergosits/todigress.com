package main
import (
  "strings"
  "path/filepath"
  "fmt"
)
func RenderList(r * RenderContext, match string, render func(*RenderContext,string)string) string {
  content := strings.TrimPrefix(match, "list:")
  tokens := strings.SplitN(content, ":", 3)
  if len(tokens) < 3 {
    fmt.Println(content)
    return "Error: improperly formatted list specifier"
  }
  directory := tokens[0]
  fileRegex := tokens[1]
  itemFormatter := tokens[2]
  files := r.FindFiles(filepath.Join(r.Directory,directory), fileRegex)
  output := ""
  for _,f := range files {
    replaced := strings.Replace(itemFormatter, "{}", strings.TrimSuffix(f, ".template"),-1)
    output = fmt.Sprintf("%s%s",output,replaced)
  }
  return render(r,output)
}
