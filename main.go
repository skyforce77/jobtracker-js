package main

import (
	"github.com/IDerr/jobtracker/providers"
	"github.com/gopherjs/gopherjs/js"
	"net/http"
	"strings"
)

func main() {
	js.Module.Get("exports").Set("getProviders", getProviders)

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

func getProviders() []*js.Object {
	jsProviders := make([]*js.Object, 0)
	for _, p := range providers.GetProviders() {
		j := p
		jsp := js.Global.Get("Object").New()
		jsp.Set("retrieveJobs", func(fn func(interface{})) {
			go j.RetrieveJobs(getCallback(fn))
		})
		jsProviders = append(jsProviders, jsp)
	}
	return jsProviders
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
