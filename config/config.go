package config

import (
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

var BaseDir, LocalisationDir, NationalIdeaPath, LanguageSuffix string

const (
	configFile = "config.yml"
)

var langPrefix = map[string]string{
	"eng": "_l_english.yml",
	"fra": "_l_french.yml",
	"deu": "_l_german.yml",
	"spa": "_l_spanish.yml",
}

type Config struct {
	BaseDir string `yaml:"base_dir"`
	Lang    string `yaml:"lang"`
}

func init() {
	c, _ := LoadConfig()
	SetConfig(c)
}

func AvailableLanguages() []string {
	return []string{"eng", "fra", "deu", "spa"}
}

func LoadConfig() (Config, error) {
	c := Config{}
	b, err := os.ReadFile(configFile)
	if err != nil {
		return c, err
	}

	err = yaml.Unmarshal(b, &c)
	if err != nil {
		return c, err
	}

	return c, nil
}

func SetConfig(c Config) {
	BaseDir = c.BaseDir
	NationalIdeaPath = filepath.Join(BaseDir, "common", "ideas", "00_country_ideas.txt")
	LocalisationDir = filepath.Join(BaseDir, "localisation")
	LanguageSuffix = langPrefix[c.Lang]
}

func WriteConfig(c Config) error {
	b, err := yaml.Marshal(c)
	if err != nil {
		return err
	}

	f, err := os.Create(configFile)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.Write(b)
	return err
}

func VerifyGameFiles() bool {
	paths := []string{BaseDir, LocalisationDir, NationalIdeaPath}
	for _, path := range paths {
		f, err := os.Open(path)
		if err != nil {
			return false
		}
		f.Close()
	}

	return true
}
