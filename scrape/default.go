/*
Copyright © 2020 D Ther Htun <dtherhtun.cw@gmail.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package scrape

import (
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/olekukonko/tablewriter"
)

//Scrape default
func Scrape() {
	var total []string
	var headers []string
	// Request the HTML page.
	res, err := http.Get("https://www.worldometers.info/coronavirus/")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	doc.Find("div#maincounter-wrap").Each(func(i int, s *goquery.Selection) {

		s.Find("span").Each(func(t int, count *goquery.Selection) {
			total = append(total, count.Text())
		})
		s.Find("h1").Each(func(h int, header *goquery.Selection) {
			title := strings.Replace(header.Text(), ":", "", -1)
			headers = append(headers, title)
		})

	})
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader(headers)
	table.SetBorder(false)
	table.SetHeaderColor(tablewriter.Colors{tablewriter.Bold, tablewriter.BgYellowColor, tablewriter.FgBlackColor}, tablewriter.Colors{tablewriter.BgRedColor, tablewriter.FgWhiteColor, tablewriter.Bold}, tablewriter.Colors{tablewriter.BgGreenColor, tablewriter.FgWhiteColor})
	table.SetColumnColor(tablewriter.Colors{tablewriter.Bold, tablewriter.FgYellowColor}, tablewriter.Colors{tablewriter.Bold, tablewriter.FgRedColor}, tablewriter.Colors{tablewriter.Bold, tablewriter.FgGreenColor})
	table.Append(total)
	table.Render()
}
