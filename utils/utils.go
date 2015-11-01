package utils
import "io/ioutil"
import "log"
import "strings"

func ReadMap(filename string)  map[string]string {
  ret  := make(map[string]string)
  bytes, err := ioutil.ReadFile(filename)
  if (err != nil) {
    log.Println("unable to read ", filename, ":", err.Error())
    return ret
  }
  file_contents := string(bytes)
  lines := strings.Split(file_contents, "\n")
  for i,line := range lines {
    if len(line) <= 0 || line[0:1] == "#" {
      continue
    }
    contents := strings.SplitN(line, ":", 2)
    if (len(contents) < 2) {
      log.Println("line",i,"invalid:",line)
      continue
    }
    ret[contents[0]] = contents[1]
  }
  return ret
}
