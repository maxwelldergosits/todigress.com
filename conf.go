package main
import "io/ioutil"
import "log"
import "strings"

type Conf map[string]string
func (c * Conf) Directory() string {
  return (*c)["directory"]
}
//test
func GetDefaultConf() * Conf {
  return ReadConf("/home/git/deployments/todigress-server/.todigress.conf")
}

func ReadConf(filename string) * Conf {
  config := Conf(make(map[string]string))
  bytes, err := ioutil.ReadFile(filename)
  if (err != nil) {
    log.Println("unable to read ", filename, ":", err.Error())
  }
  file_contents := string(bytes)
  lines := strings.Split(file_contents, "\n")
  for i,line := range lines {
    if len(line) <= 0 || line[0:1] == "#" {
      continue
    }
    log.Println("line:",line, len(line))
    contents := strings.SplitN(line, ":", 2)
    if (len(contents) < 2) {
      log.Println("line",i,"invalid:",line)
      continue
    }
    config[contents[0]] = contents[1]
  }
  return &config
}
