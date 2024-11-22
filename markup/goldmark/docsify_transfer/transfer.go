package docsify_transfer

import (
	"fmt"
	"github.com/dlclark/regexp2"
	"strings"
)

// Processor 定义一个处理函数类型
type Processor func([]byte) []byte

var (
	_processors = []Processor{
		replacePanelsOuter(panels, "$1<div class=\"docsify-example-panels\">$3</div>"),
		replaceAlerts(),
		replaceTabs(),
	}
)

func escapeCaptureRef(content string) string {
	return strings.ReplaceAll(content, "$", "$$")
}

func stringToBool(s string) (bool, error) {
	switch strings.ToLower(strings.TrimSpace(s)) {
	case "true", "1":
		return true, nil
	case "false", "0":
		return false, nil
	default:
		return false, fmt.Errorf("invalid boolean value: %s", s)
	}
}

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
