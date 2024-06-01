package route

import "net/http"

func (r MyRouteImpl) Article() {
	//http.HandleFunc("/article/getArticleByKeyword", r.getArticleByKeyword)
	http.HandleFunc("/article/GetArticleByCategory", ArticleApi.GetArticleByCategory)
	http.HandleFunc("/article/GetArticleByID", ArticleApi.GetArticleByID)
}
