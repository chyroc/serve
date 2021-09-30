package internal

import (
	"io/ioutil"
	"os"
	"strings"
)

func ReadPath(path string) (bool, string, error) {
	path = strings.TrimLeft(path, "/")
	if !strings.HasPrefix(path, "./") {
		path = "./" + path
	}
	fileinfo, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return true, BuildTemplate(Html404Template, map[string]interface{}{"Path": path}), nil
		}
		return false, "", err
	}

	if fileinfo.IsDir() {
		if _, err := os.Stat(path + "index.html"); err == nil {
			return readFile(path + "index.html")
		}
		content, err := readDir(path)
		if err != nil {
			return false, "", err
		}
		return true, content, nil
	}

	return readFile(path)
}

func readDir(path string) (string, error) {
	dirs, err := os.ReadDir(path)
	if err != nil {
		return "", err
	}

	res := []string{}
	if path != "./" {
		res = append(res, "../")
	}
	for _, dir := range dirs {
		if !dir.IsDir() {
			res = append(res, dir.Name())
		} else {
			res = append(res, dir.Name()+"/")
		}
	}

	return BuildTemplate(htmlDirTemplate, map[string]interface{}{
		"Path": path,
		"Dirs": res,
	}), nil
}

func readFile(path string) (bool, string, error) {
	bs, err := ioutil.ReadFile(path)
	if err != nil {
		return false, "", err
	}

	return strings.HasSuffix(path, ".html"), string(bs), err
}
