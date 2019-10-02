package urls

type (
	Urls struct {
		SHOW_PATH string
		GENEALOGY string
		GENEALOGY_EMPTY string
		SHOW_EMPTY string
	}
)

func ReturnURLS() Urls {
	var urlPatterns Urls
	urlPatterns.SHOW_PATH = "/show/:itemID"
	urlPatterns.GENEALOGY = "/genealogy/:catID"
	urlPatterns.GENEALOGY_EMPTY = "/genealogy/"
	urlPatterns.SHOW_EMPTY = "/show/"
	return urlPatterns
}
