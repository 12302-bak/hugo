package docsify_transfer

import (
	"github.com/dlclark/regexp2"
)

var (
	tabBlockMarkup = regexp2.MustCompile(`( *)(<!-+\s+tabs:\s*?start\s+-+>)(?:(?!(<!-+\s+tabs:\s*?(?:start|end)\s+-+>))[\s\S])*(<!-+\s+tabs:\s*?end\s+-+>)`, regexp2.Multiline)
	//tabDetailsMarkup = regexp2.MustCompile(`[\r]*(\s*)(<h[1-6].*>\s*<strong>\s*(.*[^\s])\s*<\/strong>[\s\S]*?<\/h[1-6]>)[\r]*?([\s\S]*?)(?=<h[1-6].*>\s*<strong>|<!-+\s+tabs:\s*?end\s+-+>)`, regexp2.Multiline)
	tabDetailsMarkup = regexp2.MustCompile(`[\r]*(\s*)(<h[1-6](?:\s+id="(.*)")?>\s*<strong>\s*(.*[^\s])\s*<\/strong>[\s\S]*?<\/h[1-6]>)[\r]*?([\s\S]*?)(?=<h[1-6].*>\s*<strong>|<!-+\s+tabs:\s*?end\s+-+>)`, regexp2.Multiline)
)

func replaceTabs() Processor {
	return func(data []byte) []byte {
		for {
			match, _ := tabBlockMarkup.FindStringMatch(string(data))
			if match == nil {
				break
			}

			tabBlock := match.String()
			itemAppend := ""
			anchorAppend := ""
			for {
				subMatch, _ := tabDetailsMarkup.FindStringMatch(tabBlock)
				if subMatch == nil {
					break
				}

				tabId := subMatch.Groups()[3].String()
				tabTitle := subMatch.Groups()[4].String()
				tabContent := subMatch.Groups()[5].String()

				if tabTitle == "" {
					tabTitle = "tab"
				}

				tabTitleHtml := "<button anchorId=\"" + tabId + "\" class=\"docsify-tabs__tab\" data-tab=\"" + tabTitle + "\">" + tabTitle + "</button>"
				tabContentHtml := ""
				if tabContent != "" {
					tabContentHtml = "<div class=\"docsify-tabs__content\" data-tab-content=\"" + tabId + "\">" + tabContent + "</div>"
				}

				item := tabTitleHtml + tabContentHtml
				partial, err := tabDetailsMarkup.Replace(tabBlock, item, 0, 1)
				itemAppend += item
				anchorAppend += "<h6 id=\"" + tabId + "\" class=\"anchor_hr\" style=\"font-size: 0rem; margin: 0 0;\"></h6>"

				if err != nil {
					println(err)
				}
				tabBlock = partial
			}

			eachTab := anchorAppend + "<div class=\"docsify-tabs docsify-tabs--classic\">" + itemAppend + "</div>"
			temp, _ := tabBlockMarkup.Replace(string(data), eachTab, 0, 1)
			data = []byte(temp)
		}
		return data
	}
}
