package route

import "net/http"

func (r MyRouteImpl) Article() {
	//http.HandleFunc("/article/getArticleByKeyword", r.getArticleByKeyword)
	http.HandleFunc("/article/getArticleByCategory", MyTemplate.GetArticleByCategory)
	http.HandleFunc("/article/getArticleByID", MyTemplate.GetArticleByID)
	http.HandleFunc("/article/createArticle", MyTemplate.CreateArticle)
}
