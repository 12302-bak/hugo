package docsify_transfer

import (
	"github.com/dlclark/regexp2"
	"strings"
)

var (
	alerts = regexp2.MustCompile(`<\s*blockquote[^>]*>[\s]*?(?:<p>)?\[!(\w*)((?:\|[\w*:  \u00A0-\uD7FF\uF900-\uFDCF\uFDF0-\uFFEF-  ]*)*?)\]([\s\S]*?)(?:<\/p>)?<\s*\/\s*blockquote>`, regexp2.Multiline)

	// |style:tip|label:zhonghuarenminggongheguo |iconVisibility:visible|labelVisibility:visible|icon:icon-note|className:note
	repStyle           = regexp2.MustCompile(`style:([\w\s\u00A0-\uD7FF\uF900-\uFDCF\uFDF0-\uFFEF-]*)`, regexp2.Singleline)
	repLabel           = regexp2.MustCompile(`label:([\w\s\u00A0-\uD7FF\uF900-\uFDCF\uFDF0-\uFFEF-]*)`, regexp2.Singleline)
	repIconVisibility  = regexp2.MustCompile(`iconVisibility:([\w\s\u00A0-\uD7FF\uF900-\uFDCF\uFDF0-\uFFEF-]*)`, regexp2.Singleline)
	repLabelVisibility = regexp2.MustCompile(`labelVisibility:([\w\s\u00A0-\uD7FF\uF900-\uFDCF\uFDF0-\uFFEF-]*)`, regexp2.Singleline)
	repIcon            = regexp2.MustCompile(`icon:([\w\s\u00A0-\uD7FF\uF900-\uFDCF\uFDF0-\uFFEF-]*)`, regexp2.Singleline)
	repClassName       = regexp2.MustCompile(`className:([\w\s\u00A0-\uD7FF\uF900-\uFDCF\uFDF0-\uFFEF-]*)`, regexp2.Singleline)
)

func replaceAlerts() Processor {
	return func(data []byte) []byte {
		s := string(data)
		tmp, err := alerts.ReplaceFunc(s, func(m regexp2.Match) string {
			//match, key, settings, value
			match := m.String()
			key := strings.ToLower(m.Groups()[1].String())
			settings := m.Groups()[2].String()
			value := m.Groups()[3].String()

			var style = "callout"
			var label, icon, className string
			var isIconVisible, isLabelVisible = true, true
			switch key {
			case "note":
				label = "Note"
				icon = "icon-note"
				className = "note"
			case "comment":
				label = "Comment"
				icon = "fas fa-comment"
				className = "note"
			case "tip":
				label = "Tip"
				icon = "icon-tip"
				className = "tip"
			case "warning":
				label = "Warning"
				icon = "icon-warning"
				className = "warning"
			case "attention":
				label = "Attention"
				icon = "icon-attention"
				className = "attention"
			case "caution":
				label = "Caution"
				icon = "icon icon-attention"
				className = "attention"
			default:
				return match
			}

			styleM, _ := repStyle.FindStringMatch(settings)
			if styleM != nil {
				style = styleM.Groups()[1].String()
			}

			labelM, _ := repLabel.FindStringMatch(settings)
			if labelM != nil {
				label = labelM.Groups()[1].String()
			}

			iconM, _ := repIcon.FindStringMatch(settings)
			if iconM != nil {
				icon = iconM.Groups()[1].String()
			}

			classNameM, _ := repClassName.FindStringMatch(settings)
			if classNameM != nil {
				className = classNameM.Groups()[1].String()
			}

			iconVisibilityM, _ := repIconVisibility.FindStringMatch(settings)
			if iconVisibilityM != nil {
				isIconVisible, _ = stringToBool(iconVisibilityM.Groups()[1].String())
			}

			labelVisibilityM, _ := repLabelVisibility.FindStringMatch(settings)
			if labelVisibilityM != nil {
				isLabelVisible, _ = stringToBool(labelVisibilityM.Groups()[1].String())
			}

			iconTag := ""
			if isIconVisible {
				iconTag = `<span class="icon ` + icon + `"></span> `
			}
			labelStr := ""
			if isLabelVisible {
				labelStr = label
			}

			titleTag := ""
			if isIconVisible || isLabelVisible {
				titleTag = `<p class="title">` + iconTag + labelStr + ` </p>`
			}
			return `<div class="alert ` + style + ` ` + className + `">` + titleTag + `<p>` + value + `</p></div>`
		}, -1, -1)
		if err != nil {
			return nil
		}
		return []byte(tmp)
	}
}
