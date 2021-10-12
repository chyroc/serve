package internal

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func ReadPath(dir, path string) (string, string, error) {
	if dir == "." {
		dir = "./"
	}
	path = strings.TrimLeft(path, "/")
	if !strings.HasPrefix(path, "./") {
		path = "./" + path
	}

	// path: ./a/b
	path = filepath.Join(dir, path)

	fileinfo, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return htmlContentType, BuildTemplate(Html404Template, map[string]interface{}{"Path": path}), nil
		}
		return txtContentType, "", err
	}

	if fileinfo.IsDir() {
		if !strings.HasSuffix(path, "/") {
			path = path + "/"
		}
		if _, err := os.Stat(path + "index.html"); err == nil {
			return readFile(path + "index.html")
		}
		content, err := readDir(dir, path)
		if err != nil {
			return txtContentType, "", err
		}
		return htmlContentType, content, nil
	}

	return readFile(path)
}

func readDir(dir, path string) (string, error) {
	dirs, err := os.ReadDir(path)
	if err != nil {
		return "", err
	}

	res := []string{}
	if strings.TrimRight(path, "/") != strings.TrimRight(dir, "/") {
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

func readFile(path string) (string, string, error) {
	bs, err := ioutil.ReadFile(path)
	if err != nil {
		return txtContentType, "", err
	}

	return pathToContentType(path), string(bs), err
}
