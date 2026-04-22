package docsify_transfer

import (
	"fmt"
	"testing"
)

func TestDocsTabsReplace(t *testing.T) {
	// 原始数据
	//data, _ := os.ReadFile("/Users/stevenobelia/Desktop/_scratch/clean/archive-docsify/README.md")
	data := []byte(`
<ul>
<li>
<h2 id="hello">
  hello
  <a class="anchor" href="#hello">#</a>
</h2>
  <!-- tabs:start -->
<h4 id="nasm">
  <strong>nasm</strong>
  <a class="anchor" href="#nasm">#</a>
</h4>
<div class="outer yosemite"></div>
<div class="code-toolbar"><pre data-lang="nasm"><code class="language-nasm">section .text
    msg db 'Hello, world!'
</code></pre></div>
  <!-- tabs:end --></li>
</ul>
`)

	// 应用所有处理器
	result := ApplyPostProcessors(data, replaceTabs())
	//result := test()

	// 输出结果
	fmt.Println(string(result))
}

func TestDocsTabsRegReplace(t *testing.T) {
	input := "  <!-- tabs:start -->\n<h4 id=\"nasm\">\n  <strong>nasm</strong>\n  <a class=\"anchor\" href=\"#nasm\">#</a>\n</h4>\n<div class=\"outer yosemite\"></div>\n<div class=\"code-toolbar\"><pre data-lang=\"nasm\"><code class=\"language-nasm\">section .text\n    msg db 'Hello, world!$$'\n</code></pre></div>\n  <!-- tabs:end -->"
	replace := "e!$$'"
	//compile := regexp2.MustCompile(`(?:[\s\S]+)(a)(?=t)`, regexp2.Multiline)

	s, err := tabDetailsMarkup.Replace(input, replace, 0, 1)

	// https://learn.microsoft.com/en-us/dotnet/standard/base-types/substitutions-in-regular-expressions?#:~:text=$'
	if err != nil {
		fmt.Println(s)
	}
}
