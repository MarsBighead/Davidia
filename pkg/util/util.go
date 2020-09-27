package util

import (
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"

	yaml "gopkg.in/yaml.v2"
)

//GetUsageReportFiles get usage report file from the input
func GetUsageReportFiles(pathName string) ([]string, error) {
	fs, err := os.Stat(pathName)
	if err != nil {
		return nil, err
	}
	var files []string
	if fs.IsDir() {
		re := regexp.MustCompile(`^bp\_2\d{3}(0[1-9]|1[0-2])\.xlsx$`)
		pathName = regexp.MustCompile(`\/$`).ReplaceAllString(pathName, ``)
		filenames, err := ioutil.ReadDir(pathName)
		if err != nil {
			return nil, err
		}
		for _, f := range filenames {
			if re.MatchString(f.Name()) {
				files = append(files, pathName+"/"+f.Name())
			}
		}
	} else {
		if regexp.MustCompile(`/bp\_2\d{3}(0[1-9]|1[0-2])\.xlsx$`).MatchString(pathName) {
			files = append(files, pathName)
		}

	}

	return files, nil
}

//GetCurrentPath Get current path, default supoort only `go run main.go`
func GetCurrentPath(filename string) (string, error) {
	if regexp.MustCompile(`[Dd]avidia`).MatchString(filename) {
		file, err := exec.LookPath(filename)
		if err != nil {
			return "", err
		}
		path, err := filepath.Abs(file)
		if err != nil {
			return "", err
		}
		return filepath.Dir(path), nil
	}
	return ".", nil
}

// Config configure information struct
type Config struct {
	Home    string  `yaml:"-"`
	Project string  `yaml:"project"`
	Source  *Source `yaml:"source"`
	Output  *Output `yaml:"output"`
}

// Source data source file
type Source struct {
	Class string `yaml:"class"`
	Name  string `yaml:"name"`
}

// Output report common information
type Output struct {
	Base string `yaml:"base"`
}

// Parse config.yaml to data struct
func Parse(dir, filename string) (cfg *Config, err error) {
	cfg = new(Config)
	body, err := ioutil.ReadFile(dir + "/" + filename)
	if err != nil {
		return
	}
	err = yaml.Unmarshal(body, cfg)
	if err != nil {
		return
	}
	cfg.Home = dir
	return
}
