package worker

import (
	"fmt"
	"regexp"
	"bufio"
	"strings"
	"strconv"
)

func Lofter(name string, cnum chan int) {

	r := regexp.MustCompile(`<img[^>]+\bsrc=["']([^"']+)["']`)
	url := fmt.Sprintf("http://%s.lofter.com/", name)
	f := GetFile("lofter_"+name)
	w := bufio.NewWriter(f)
	num := 0
	pageNum := 1
	defer f.Close()
	for {
		page_url := fmt.Sprintf("%s?page=%d", url, pageNum)
		fmt.Println(page_url)
		content, err := ReadContent(page_url)
		if err != nil {
			continue
		}
		match := r.FindAllStringSubmatch(content, -1)
		if len(match) < 2 {
			break
		}

		for i := range match {
			u := match[i][1]
			u = strings.Split(u, "?")[0] + "\n"
			w.WriteString(u)
			num += 1
		}
		pageNum += 1
	}
	w.Flush()
	fmt.Println("Finished! Crawl %d link!", num)
	cnum <- 1
}

func Douban(name string, cnum chan int)  {
	urlCmp := regexp.MustCompile(`<a href="(.+?)"[^>]class="photolst_photo"`)
	picCmp := regexp.MustCompile(`<a href="#" class="view-zoom view-zoom-out"><img src="(.+?)"`)
	albumUrl := fmt.Sprintf("https://www.douban.com/photos/album/%s/?start=", name)
	start := 0
	var lis []string

	for  {
		url := albumUrl + strconv.Itoa(start)
		fmt.Println(url)
		content, err := ReadContent(url)
		if err != nil {
			continue
		}
		match := urlCmp.FindAllStringSubmatch(content, -1)
		for i := range match {
			newUrl := match[i][1] + "large"
			pic_cont, err := ReadContent(newUrl)
			if err != nil {
				continue
			}
			large_url := picCmp.FindAllStringSubmatch(pic_cont,-1)
			lis = append(lis, large_url[0][1])
		}
		if len(match) < 18{
			break
		}
		start += 18
	}
	f := GetFile("douban_"+name)
	defer f.Close()
	w := bufio.NewWriter(f)
	for _, v := range lis {
		w.WriteString(v+"\n")
	}
	w.Flush()
	fmt.Println("Finished! Crawl %d link!", len(lis))
	cnum <- 1
}

func Tumblr(name string, cnum chan int) {
	r := regexp.MustCompile(`<photo-url max-width="1280">(.+?)</photo-url>`)
	api_url := fmt.Sprintf("https://%s.tumblr.com/api/read?type=photo&num=50&start=", name)
	f := GetFile("tumblr_"+name)
	w := bufio.NewWriter(f)
	start, num := 0, 0
	defer f.Close()
	for  {
		url := api_url + strconv.Itoa(start)
		fmt.Println(url)
		content, err := ReadContent(url)
		if err != nil {
			continue
		}
		match := r.FindAllStringSubmatch(content, -1)

		out := make(map[string]int)
		for i := range match {
			fmt.Println(match[i])
			u := match[i][1] + "\n"
			out[u] = 0
		}
		num += len(out)
		for k := range out {
			w.WriteString(k)
		}
		if len(match) < 50{
			break
		}
		start += 50
	}
	w.Flush()
	fmt.Println("Finished! Crawl %d link!", num)
	cnum <- 1
}
