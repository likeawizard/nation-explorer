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

		ReadLangFile(fp)
	}

}

func Get(key string) string {
	if v, ok := tokens[key]; ok {
		return v
	}

	return key
}

func ReadLangFile(filePath string) {
	f, err := os.Open(filePath)
	if err != nil {
		return
	}
	defer f.Close()
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

func snakeCaseToPretty(s string) string {
	return strings.Title(strings.ReplaceAll(s, "_", " "))
}
