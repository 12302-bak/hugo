package docsify_transfer

import (
	"github.com/dlclark/regexp2"
)

var (
	// error parsing regexp: invalid or unsupported Perl syntax: `(?=`
	//invalid regex       = regexp.MustCompile(`<!-+\s+div:\s*(.*)\s+-+>[\r\n]+([\s\S]*?)[\r\n]+((?=\s*<!-+\s+div:?)|(?=\s*<!-+\s+panels?))`)
	panels = regexp2.MustCompile(`[\r\n]*(\s*)(<!-+\s+panels:\s*?start\s+-+>)[\r\n]+([\s|\S]*?)[\r\n\s]+(<!-+\s+panels:\s*?end\s+-+>)`, regexp2.Multiline)
	//panelMakeup = regexp2.MustCompile(`<!-+\s+div:\s*(.*)\s+-+>[\r\n]+([\s\S]*?)[\r\n]+((?=\s*<!-+\s+div:?)|(?=\s*<!-+\s+panels?))`, regexp2.Multiline)
	panelMakeup = regexp2.MustCompile(`<!-+\s+div:\s*(.*)\s+-+>[\r\n]+([\s\S]*?)[\r\n]*?((?=\s*<!-+\s+div:?)|(?=\s*<!-+\s+panels?))`, regexp2.Multiline)
	panelAttr   = regexp2.MustCompile(`(title|left|right)-panel(-(\d*))?`, regexp2.Singleline)
)

func replacePanelsOuter(re *regexp2.Regexp, replacement string) Processor {
	return func(data []byte) []byte {
		for {
			match, _ := panels.FindStringMatch(string(data))
			if match == nil {
				break
			}

			wrapper := match.String()
			//panelWrapperIndent := match.Groups()[1].String()
			for {
				subMatch, _ := panelMakeup.FindStringMatch(wrapper)
				if subMatch == nil {
					break
				}

				panelName := subMatch.Groups()[1].String()
				panelContent := subMatch.Groups()[2].String()

				name := ""
				width := ""
				attr, _ := panelAttr.FindStringMatch(panelName)
				if attr != nil {
					name = attr.Groups()[1].String() + "-panel"
					width = attr.Groups()[3].String()
					if width != "" {
						width = "style=\"max-width: " + width + "%; width: " + width + "%;\""
					}
				}

				repl := "<div class=\"docsify-example-panel " + name + "\"" + width + ">\n" +
					escapeCaptureRef(panelContent) +
					"</div>"

				body, err := panelMakeup.Replace(wrapper, repl, 0, 1)

				if err != nil {
					println(err)
				}
				wrapper = body
			}

			rWrapper, _ := panels.Replace(wrapper, "$1"+"<div class=\"docsify-example-panels\">$3"+
				"</div>\n", 0, 1)

			temp, _ := panels.Replace(string(data), escapeCaptureRef(rWrapper), 0, 1)
			data = []byte(temp)
		}

		return data
	}
}

// replaceAllWithPrecompiled 使用预编译的正则表达式替换所有匹配项
//func replaceAllWithPrecompiled(re *regexp2.Regexp, replacement string) Processor {
//	return func(data []byte) []byte {
//		return re.ReplaceAll(data, []byte(replacement))
//	}
//}
//
//// replaceGroupWithPrecompiled 使用预编译的正则表达式替换匹配项的分组
//func replaceGroupWithPrecompiled(re *regexp2.Regexp, replacement string) Processor {
//	return func(data []byte) []byte {
//		return re.ReplaceAllFunc(data, func(match []byte) []byte {
//			subMatches := re.FindSubmatch(match)
//			if len(subMatches) > 1 {
//				return []byte(fmt.Sprintf(replacement, subMatches[1:]))
//			}
//			return match
//		})
//	}
//}
