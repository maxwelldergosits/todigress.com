package main

import "testing"

func TestRenderTemplate(t * testing.T) {
  input := "{{go.template}}"
  output := "go.template.go"

  testResolver := func(in string)string {
    return in+".go"
  }

  var r RenderContext
  r.Resolver = testResolver
  r.Templaters = []Templater{{RenderTemplate, "{{.*\\.template}}"}}

  returned := RenderString(&r, input)
  if (returned != output) {
    t.Fail()
  }
}

func TestRenderStringResolver(t * testing.T) {
  input := "{{list::.*:<:>}}"
  output := "<hello><world><goodbye>"

  var r RenderContext

  r.FindFiles = func(dir, regex string) []string {
    return []string{"hello.template","world.template","goodbye.template"}
  }

  r.Templaters = []Templater{{RenderList, "{{list:.*"}}

  returned := RenderString(&r, input)
  if (returned != output) {
    t.Fail()
  }
}
