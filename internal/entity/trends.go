package entity

import (
	"encoding/xml"
	"fmt"
	"strings"
	"time"
)

// TODO: Store chat country in Trends struct for reactive trend messages

type Trends struct {
	XML  xml.Name `xml:"rss"`
	Data *data    `xml:"channel"`
}

type data struct {
	ItemList []item `xml:"item"`
}

type item struct {
	Title    string `xml:"title"`
	NewsList []news `xml:"news_item"`
}

type news struct {
	Headline string `xml:"news_item_title"`
	Link     string `xml:"news_item_url"`
}

func (t *Trends) String() string {
	trendsArr := make([]string, 0, len(t.Data.ItemList))
	numEmojis := []string{"1️⃣", "2️⃣", "3️⃣", "4️⃣", "5️⃣", "6️⃣", "7️⃣", "8️⃣", "9️⃣", "🔟"}

	header := fmt.Sprintf("🔍 What's discussed on %s\n", time.Now().Format("Jan 02 2006"))

	trendsArr = append(trendsArr, header)

	for i, item := range t.Data.ItemList {
		itemTitle := numEmojis[i]+" "+item.Title
		trendsArr = append(trendsArr, itemTitle)
		for _, news := range item.NewsList {
			str := fmt.Sprintf("%s\n", news.Link)
			trendsArr = append(trendsArr, str)
		}
		i++
		if i > len(numEmojis)-1 {
			break
		}
	}

	return strings.Join(trendsArr, "\n")
}