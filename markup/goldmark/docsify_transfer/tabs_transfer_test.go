package docsify_transfer

import (
	"fmt"
	"testing"
)

func TestDocsTabsReplace(t *testing.T) {
	// 原始数据
	//data, _ := os.ReadFile("/Users/stevenobelia/Desktop/_scratch/clean/archive-docsify/README.md")
	data := []byte(`
  <!-- tabs:start -->
<h5 id="ubuntu184">
  <strong>Ubuntu:18.4</strong>
  <a class="anchor" href="#ubuntu184">#</a>
</h5>
<p>End of standard support for 18.04 LTS - 31 May 2023</p>
<h5 id="ubuntu204">
  <strong>Ubuntu:20.4</strong>
  <a class="anchor" href="#ubuntu204">#</a>
</h5>
<p>嗨！你知道我们有中文站吗？立即带我去！ ›</p>
  <!-- tabs:end -->
`)

	// 应用所有处理器
	result := ApplyPostProcessors(data)
	//result := test()

	// 输出结果
	fmt.Println(string(result))
}
