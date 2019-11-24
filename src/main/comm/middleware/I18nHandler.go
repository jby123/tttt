package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/text/language"
	"golang.org/x/text/language/display"
)

func I18nMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		locale := ctx.Query("locale")
		if locale != "" {
			ctx.Request.Header.Set("Accept-Language", locale)
		}
		language := ctx.GetHeader("Accept-Language")
		getAcceptLanguage(language)

		// NOTE: On June 2012, the deprecation of recommendation to use the "X-" prefix has become official as RFC 6648.
		// https://stackoverflow.com/questions/3561381/custom-http-headers-naming-conventions
		ctx.Request.Header.Set("I18n-Language", language)
		ctx.Next()
	}
}

func getAcceptLanguage(acceptLanguage string) {
	var serverLanguages = []language.Tag{
		language.SimplifiedChinese, // zh-Hans fallback
		language.AmericanEnglish,   // en-US
		language.Korean,            // de
	}

	// 也可以不定义 serverLangs 用下面一行选择支持所有语种。
	// var matcher = language.NewMatcher(message.DefaultCatalog.Languages())
	var matcher = language.NewMatcher(serverLanguages)
	t, _, _ := language.ParseAcceptLanguage(acceptLanguage)
	tag, index, confidence := matcher.Match(t...)

	fmt.Printf("best match: %s (%s) index=%d confidence=%v\n",
		display.English.Tags().Name(tag),
		display.Self.Name(tag),
		index, confidence)

	str := fmt.Sprintf("tag is %s", tag)
	fmt.Println(str)
	fmt.Printf("best match: %s\n", display.Self.Name(tag))
}
