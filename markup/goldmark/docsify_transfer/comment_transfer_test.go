package docsify_transfer

import (
	"fmt"
	"testing"
)

func TestCommentReplace(t *testing.T) {
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
