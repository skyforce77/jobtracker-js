package main

import (
	"github.com/skyforce77/jobtracker/providers"
	"github.com/gopherjs/gopherjs/js"
	"net/http"
	"reflect"
	"strings"
)

func main() {
	js.Module.Get("exports").Set("getProviders", getProviders)
	generateConstructors()

	httpFix()
}

func isNode() bool {
	if process := js.Global.Get("process"); process != js.Undefined {
		argv := process.Get("argv")
		return strings.Contains(argv.String(), "node")
	}
	return false
}

func httpFix() {
	if isNode() {
		js.Global.Set("XMLHttpRequest", js.Global.Call("require", "xhr2"))
		http.DefaultTransport = &http.XHRTransport{}
	}
}

func convertProvider(p providers.Provider) *js.Object {
	jsp := js.Global.Get("Object").New()
	jsp.Set("retrieveJobs", func(fn func(interface{})) {
		go p.RetrieveJobs(getCallback(fn))
	})
	return jsp
}

func getProviders() []*js.Object {
	jsProviders := make([]*js.Object, 0)
	for _, p := range providers.GetProviders() {
		jsp := convertProvider(p)
		jsProviders = append(jsProviders, jsp)
	}
	return jsProviders
}

func generateConstructors() {
	for _, p := range providers.GetProviders() {
		typ := reflect.TypeOf(p)
		name := typ.String()[11:]
		if name[0] == '_' {
			name = name[1:]
		}
		name = "new" + strings.Title(name)

		js.Module.Get("exports").Set(name, func() *js.Object {
			return convertProvider(p)
		})
	}
}

func getCallback(fn func(interface{})) func(job *providers.Job) {
	return func(job *providers.Job) {
		jsJob := js.Global.Get("Object").New()
		jsJob.Set("title", job.Title)
		jsJob.Set("company", job.Company)
		jsJob.Set("location", job.Location)
		jsJob.Set("type", job.Type)
		jsJob.Set("desc", job.Desc)
		jsJob.Set("link", job.Link)
		jsJob.Set("misc", job.Misc)

		fn(jsJob)
	}
}
