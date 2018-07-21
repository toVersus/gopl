package xkcd

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

func GetComic(index int) (Comic, error) {
	var comic Comic
	urlStr := strings.Join([]string{ApiURL, strconv.Itoa(index), "info.0.json"}, "/")
	fmt.Printf("%#v\n", urlStr)

	resp, err := http.Get(urlStr)
	if err != nil {
		return comic, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return comic, fmt.Errorf("can(t get comic %d: %s", index, err)
	}
	if err = json.NewDecoder(resp.Body).Decode(&comic); err != nil {
		return comic, err
	}
	return comic, nil
}
