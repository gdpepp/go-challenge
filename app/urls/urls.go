package urls

type (
	Urls struct {
		SHOW_PATH string
		GENEALOGY string
	}
)

func ReturnURLS() Urls {
	var urlPatterns Urls
	urlPatterns.SHOW_PATH = "/show/:itemID"
	urlPatterns.GENEALOGY = "/genealogy/:catID"
	return urlPatterns
}
