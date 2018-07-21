package xkcd

const ApiURL = "https://xkcd.com"

type Comic struct {
	Num              int
	Year, Month, Day string
	Title            string
	Transcript       string
	Alt              string
	Img              string
}
