package main

import (
	"fmt"

	. "github.com/codingXiang/gogo-i18n"
	"golang.org/x/text/language"
)

func main() {
	GGi18n = NewGoGoi18n(language.TraditionalChinese)
	GGi18n.SetFileType("yaml")
	GGi18n.LoadTranslationFile("./locale",
		language.TraditionalChinese,
		language.English)

	msg := GGi18n.GetMessage("welcome", map[string]interface{}{
		"username": "阿翔",
	})
	fmt.Println(msg)

	GGi18n.SetUseLanguage(language.English)
	msg = GGi18n.GetMessage("welcome", map[string]interface{}{
		"username": "阿翔",
	})
	fmt.Println(msg)
}
