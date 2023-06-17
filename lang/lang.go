package lang

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

const PATH = "/home/arturs/.local/share/Steam/steamapps/common/Europa Universalis IV/localisation/"

var tokens map[string]string = make(map[string]string)

var tokenMatch = regexp.MustCompile(`(?P<key>[^\s]+):([0-9]*)?\s"(?P<value>.*)"`)

func init() {
	dir, err := os.Open(PATH)

	if err != nil {
		fmt.Println(err)
		return
	}
	defer dir.Close()
	dirInfo, err := dir.Stat()

	if err != nil {
		fmt.Println(err)
		return
	}

	if !dirInfo.IsDir() {
		fmt.Println("not a directory")
		return
	}

	files, err := os.ReadDir(PATH)

	if err != nil {
		fmt.Println(err)
	}

	for _, file := range files {
		if file.IsDir() || !strings.Contains(file.Name(), "_l_english.yml") {
			continue
		}
		fp := filepath.Join(PATH, file.Name())

		f, err := os.Open(fp)

		if err != nil {
			fmt.Println(err)
			continue
		}
		s := bufio.NewScanner(f)

		for s.Scan() {
			m := tokenMatch.FindStringSubmatch(s.Text())
			if m == nil {
				continue
			}
			tokens[m[tokenMatch.SubexpIndex("key")]] = m[tokenMatch.SubexpIndex("value")]
		}
		f.Close()
	}

}

func Get(key string) string {
	if v, ok := tokens[key]; ok {
		return v
	}

	return key
}
