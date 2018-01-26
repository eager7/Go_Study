package main

import . "github.com/eager7/go/mlog"
import (
	"github.com/PuerkitoBio/goquery"
	"fmt"
)

func main(){
	Debug.Println("豆瓣读书抓取程序...")

	doc, err := goquery.NewDocument("http://metalsucks.net")
	if err != nil {
		Error.Fatal(err)
	}

	//find the review items
	doc.Find(".sidebar-reviews article .content-block").Each(func(i int, s *goquery.Selection){
		fmt.Println(s.Text())
		//for each item found, get the band and title
		band := s.Find("a").Text()
		title := s.Find("i").Text()
		Info.Printf("Review[%d]:%s-%s\n", i, band, title)
	})
}
