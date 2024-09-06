package main

import (
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"strconv"
)

type post struct {
	title   string
	url     string
	summary string
	date    string
}

var baseURL string = "https://tech.kakaopay.com"
var pageURL string = baseURL + "/page/"

func main() {

	// 어케 totalPage 를 파악하지
	// page 범위를 넘어가면 404 를 뱉는다.
	for i := 1; i < 18; i++ {
		pages := getPages(i)
		log.Println(pages)
	}

}

func getPages(page int) []post {

	var posts []post
	res, err := http.Get(pageURL + strconv.Itoa(page))

	checkErr(err)
	checkCode(res)
	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	doc.Find("._postListItem_1cl5f_66>a").Each(func(i int, selection *goquery.Selection) {

		title := selection.Find("._postInfo_1cl5f_99>strong")
		href, _ := selection.Attr("href")
		summary := selection.Find("p")
		date := selection.Find("time")

		post := post{title: title.Text(), url: baseURL + href, summary: summary.Text(), date: date.Text()}
		posts = append(posts, post)

	})
	return posts
}

func checkErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func checkCode(res *http.Response) {
	if res.StatusCode != 200 {
		log.Fatalln("Request failed with Status:", res.StatusCode)
	}
}
