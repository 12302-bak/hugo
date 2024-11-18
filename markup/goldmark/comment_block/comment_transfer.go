package comment_block

import (
	"fmt"
	"github.com/dlclark/regexp2"
	"strings"
)

var (
	// error parsing regexp: invalid or unsupported Perl syntax: `(?=`
	//invalid regex       = regexp.MustCompile(`<!-+\s+div:\s*(.*)\s+-+>[\r\n]+([\s\S]*?)[\r\n]+((?=\s*<!-+\s+div:?)|(?=\s*<!-+\s+panels?))`)
	panels = regexp2.MustCompile(`[\r\n]*(\s*)(<!-+\s+panels:\s*?start\s+-+>)[\r\n]+([\s|\S]*?)[\r\n\s]+(<!-+\s+panels:\s*?end\s+-+>)`, regexp2.Multiline)
	//panelMakeup = regexp2.MustCompile(`<!-+\s+div:\s*(.*)\s+-+>[\r\n]+([\s\S]*?)[\r\n]+((?=\s*<!-+\s+div:?)|(?=\s*<!-+\s+panels?))`, regexp2.Multiline)
	panelMakeup = regexp2.MustCompile(`<!-+\s+div:\s*(.*)\s+-+>[\r\n]+([\s\S]*?)[\r\n]*?((?=\s*<!-+\s+div:?)|(?=\s*<!-+\s+panels?))`, regexp2.Multiline)
	panelAttr   = regexp2.MustCompile(`(title|left|right)-panel(-(\d*))?`, regexp2.Singleline)
	_processors = []Processor{
		replacePanelsOuter(panels, "$1<div class=\"docsify-example-panels\">$3</div>"),
	}
)

// Processor 定义一个处理函数类型
type Processor func([]byte) []byte

// ApplyPostProcessors 依次应用多个处理器
func ApplyPostProcessors(data []byte, processors ...Processor) []byte {
	if processors == nil {
		processors = _processors
	}
	for _, p := range processors {
		data = p(data)
	}
	return data
}

func escapeCaptureRef(content string) string {
	return strings.ReplaceAll(content, "$", "$$")
}

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

func main() {
	// 原始数据
	//data, _ := os.ReadFile("/Users/stevenobelia/Desktop/_scratch/clean/archive-docsify/README.md")
	data := []byte(`
<ul>
<li>
<h2 id="intropanels">
  Intro(PANELS)
  <a class="anchor" href="#intropanels">#</a>
</h2>
</li>
</ul>
<p>hello</p>
<ul>
<li>
<h2 id="面板布局">
  面板布局
  <a class="anchor" href="#%e9%9d%a2%e6%9d%bf%e5%b8%83%e5%b1%80">#</a>
</h2>
  <!-- panels:start -->
  <!-- div:title-panel -->
<h5 id="hello-world">
  Hello World
  <a class="anchor" href="#hello-world">#</a>
</h5>
  <!-- div:left-panel-40 -->
<blockquote>
<p>[?] If you are on widescreen, checkout the <em>right</em> panel, <em>right</em> there →</p>
</blockquote>
  <!-- div:right-panel-60 -->
<blockquote>
<p>[?] This is an example panel.
<br>You can see it&rsquo;s usage in practice in the docs listed below:</p>
</blockquote>
<ul>
<li>
<a href="https://fairlay.com/api" rel="noopener">Fairlay API</a></li>
<li>
<a href="https://docs.flap.cloud/#/create_new_service?id=special-files" rel="noopener">FLAP services</a></li>
</ul>
<p><small>please contact me if you use docsify-example-panels. i would like to display it here too.</small></p>
  <!-- panels:end --></li>
</ul>
`)

	// 应用所有处理器
	result := ApplyPostProcessors(data)
	//result := test()

	// 输出结果
	fmt.Println(string(result))
}

func test() string {
	data := []byte("hello world")
	compile := regexp2.MustCompile(`(hello) (world)`, regexp2.Multiline)

	replace, err := compile.Replace(string(data), "$1+$$2", 0, -1)
	if err != nil {
		return ""
	}
	// 输出结果
	return replace
}
