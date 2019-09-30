package url

type (
	urls struct {
		SHOW_PATH string
		GENEALOGY string
	}
)

func ReturnURLS() urls {
	var urlPatterns urls
	urlPatterns.SHOW_PATH = "/show/:itemID"
	urlPatterns.GENEALOGY = "/genealogy/:catID"
	return urlPatterns
}
