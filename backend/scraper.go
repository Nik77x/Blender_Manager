package backend

import (
	"fmt"
	"github.com/Nik77x/Blender_Manager/backend/Data"
	"github.com/gocolly/colly/v2"
	"runtime"
	"strings"
)

func ScrapeVersions() []Data.BlendInfo {

	c := colly.NewCollector(
		colly.AllowedDomains("builder.blender.org", "www.builder.blender.org"),
	)

	var blendInfos = make([]Data.BlendInfo, 0, 5)

	c.OnRequest(func(request *colly.Request) {
		request.Headers.Set("Accept-Language", "en-US;q=0.9")
	})

	c.OnError(func(response *colly.Response, err error) {
		// TODO: Display error boxes don't just print to console
		println("Scraping failed with error:")
		println(err)
		return
	})

	selector := fmt.Sprintf(".platform-%v .builds-list", runtime.GOOS)

	c.OnHTML(selector, func(element *colly.HTMLElement) {

		sel := element.DOM

		// Filter versions
		sel = sel.Children()
		sel = sel.Not("[style=display\\:none\\;]")

		for i := range sel.Nodes {

			listItem := sel.Eq(i)

			tmpInfo := Data.BlendInfo{}

			tmpInfo.FileExtension = listItem.Find("div.build-info ul.build-details [title=File\\ extension]").Text()

			if strings.Contains(tmpInfo.FileExtension, "msi") {
				continue
			}

			tmpInfo.BuildType = listItem.Find("div.build-info a.build-title span.build-var").Text()
			tmpInfo.DownloadLink, _ = listItem.Find("div.build-info [ga_cat=download]").Attr("href")
			tmpInfo.CommitLink, _ = listItem.Find("a[title=See\\ commit]").Attr("href")
			tmpInfo.CommitsHistoryLink, _ = listItem.Find("a[title=See\\ history]").Attr("href")
			tmpInfo.CommitHash = listItem.Find("a[title=See\\ history]").Text()
			tmpInfo.FileSize = listItem.Find("div.build-info ul.build-details [title=File\\ size]").Text()
			tmpInfo.Sha256DownloadLink, _ = listItem.Find("a.sha[title=SHA-256\\ for\\ this\\ file]").Attr("href")
			tmpInfo.VersionTitle = listItem.Find("a.build-title").Text()
			tmpInfo.UploadDate = listItem.Find("div.build-info ul.build-details").Children().First().Text()

			blendInfos = append(blendInfos, tmpInfo)

		}

	})

	err := c.Visit("https://builder.blender.org/download/experimental/")
	// TODO: Handle them errors!
	c.Visit("https://builder.blender.org/download/daily/")
	c.Visit("https://builder.blender.org/download/patch/")

	if err != nil {
		println("cant visit URL")
		panic(err)
	}

	return blendInfos

}
