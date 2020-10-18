package resource

import (
	"strings"
	"time"
)

var (
	domainName = []string{
		"image", "api", "www", "media", "photo", "music", "tech", "doc", "document", "pic", "zh", "en", "html", "help", "code", "git", "github", "view",
		"video", "text", "txt", "computer", "entertainment", "happy", "coding", "blog", "mobile", "phone", "shop", "app", "banner", "buffer", "bug", "button",
		"cache", "crash", "code", "cover", "deadline", "delay", "demo", "feature", "idea", "icon", "loading", "log", "logo", "native", "push", "review",
		"slogan", "tab", "tag", "timeline", "tips", "title", "toast", "alter", "burst", "dispose", "blast", "consume", "split", "spit", "spill", "slip", "slide",
		"bacteria", "breed", "budget", "candidate", "campus", "liberal", "transform", "transmit", "transplant", "transport", "shift", "vary", "vanish", "swallow",
		"suspicion", "suspicious", "mild", "tender", "nuisance", "insignificant", "accelerate", "absolute", "boundary", "brake", "catalog", "vague", "vain", "extinct",
		"extraordinary", "extreme", "agent", "alcohol", "appeal", "appreciate", "approve", "stimulate", "acquire", "accomplish", "network", "tide", "tidy",
		"trace", "torture", "wander", "wax", "weave", "preserve", "abuse", "academic", "academy", "battery", "barrier", "cargo", "career", "vessel", "vertical",
		"oblige", "obscure", "extent", "exterior", "external", "petrol", "petroleum", "delay", "decay", "decent", "route", "ruin", "sake", "satellite", "scale",
		"temple", "tedious", "tend", "tendency", "ultimate", "undergo", "abundant", "adopt", "adapt", "bachelor", "casual", "trap", "vacant", "vacuum", "oral",
		"optics", "organ", "excess", "expel", "expend", "expenditure", "expense", "expensive", "expand", "expansion", "grant", "grand", "invade", "acid",
		"acknowledg", "balcon", "calculat", "calenda", "optimisti", "optiona", "outstandin", "expor", "importimpos", "religio", "religiou", "victi", "vide",
		"videotap", "offen", "bothe", "interfer", "interna", "beforehan", "racia", "radiatio", "radica", "rang", "wonde", "isolat", "issu", "hollo", "hoo",
	}
	topDomain = []string{
		".com", ".edu", ".gov", ".int", ".mil", ".net", ".org", ".biz", ".info", ".pro", ".name", ".museum", ".coop", ".aero", ".xxx", ".idv",
	}
)

func RandomUrl() string {
	unixNano := time.Now().UnixNano()
	builder := new(strings.Builder)
	if unixNano%2 == 0 {
		builder.WriteString("https")
	} else {
		builder.WriteString("http")
	}
	builder.WriteString("://")
	thirdLevelDomainLength := len(domainName)
	topDomainLength := len(topDomain)
	if thirdLevelDomainLength%2 == 0 {
		thirdLevelDomainLength -= 1
	}
	if topDomainLength%2 == 0 {
		topDomainLength -= 1
	}
	index1 := unixNano % int64(thirdLevelDomainLength)
	index2 := unixNano % int64(thirdLevelDomainLength-1)
	builder.WriteString(domainName[index1])
	builder.WriteString(".")
	builder.WriteString(domainName[index2])
	builder.WriteString(topDomain[unixNano%int64(topDomainLength)])
	return builder.String()
}
