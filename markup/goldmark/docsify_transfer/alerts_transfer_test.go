package docsify_transfer

import (
	"fmt"
	"github.com/dlclark/regexp2"
	"testing"
)

func TestReplaceFunc(t *testing.T) {
	data := `this is a test
for regexp2.ReplaceFunc(),
replace something with func,
and check result, verify.
`
	regexp := regexp2.MustCompile(`(.*)a(.*)`, regexp2.Multiline)

	result, err := regexp.ReplaceFunc(data, func(match regexp2.Match) string {
		return match.Groups()[1].String() + "x" + match.Groups()[2].String()
	}, -1, -1)

	if err != nil {
		panic(err)
	}
	fmt.Printf("\n%s\n", result)
}

func TestFunction(t *testing.T) {
	data := []byte(`
<blockquote><p>[!NOTE|asd:234] An alert of type ’note’ using global style ‘callout’.</p></blockquote>
<blockquote><p>[!TIP|style:flat|label:中华人名共和国] An alert of type ’note’ using global style ‘callout’.</p></blockquote>
<blockquote><p>[!NOTE|icon:234] An alert of type ’note’ using global style ‘callout’.</p></blockquote>
<blockquote><p>[!CAUTION|iconVisibility:false] An alert of type ’note’ using global style ‘callout’.</p></blockquote>
<blockquote><p>[!ATTENTION|labelVisibility:false] An alert of type ’note’ using global style ‘callout’.</p></blockquote>
`)
	bytes := replaceAlerts()(data)

	fmt.Printf("\n\n%s\n", string(bytes))
}
