package template

import "net/http"

func (t MyTemplateImpl) Login(w http.ResponseWriter, res *http.Request) {
	if res.Method == http.MethodGet {
		t.RenderTemplate(w, "./static/html/login.html", nil)
	}
}

func (t MyTemplateImpl) Register(w http.ResponseWriter, res *http.Request) {
	if res.Method == http.MethodGet {
		t.RenderTemplate(w, "./static/html/register.html", nil)
	}
}
