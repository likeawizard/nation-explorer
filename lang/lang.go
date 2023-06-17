package lang

import (
	"bufio"
	"fmt"
	"likeawizard/nation-explorer/config"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

var tokens map[string]string = make(map[string]string)

var tokenMatch = regexp.MustCompile(`(?P<key>[^\s]+):([0-9]*)?\s"(?P<value>.*)"`)

func init() {
	files, err := os.ReadDir(config.LocalisationDir)

	if err != nil {
		fmt.Println(err)
	}

	for _, file := range files {
		if file.IsDir() || !strings.Contains(file.Name(), config.LanguageSuffix) {
			continue
		}
		fp := filepath.Join(config.LocalisationDir, file.Name())

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
