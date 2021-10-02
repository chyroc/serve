package internal

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/chyroc/mkcert"
	"github.com/gin-gonic/gin"
)

func Listen(handler *gin.Engine, isHttps bool) error {
	listener, port, err := CreateLocalListener()
	if err != nil {
		return err
	}

	hosts, err := getEnableHttpsHost()
	if err != nil {
		return err
	}

	if isHttps {
		certFile, keyFile, err := createCACert(hosts)
		if err != nil {
			return err
		}

		printListener(isHttps, hosts, port)
		return http.ServeTLS(listener, handler, workDirCert+"/"+certFile, workDirCert+"/"+keyFile)
	}

	printListener(isHttps, hosts, port)
	return handler.RunListener(listener)
}

func createCACert(hosts []string) (string, string, error) {
	cache := newCache(workDirCert + "/cache.json")
	key := strings.Join(hosts, "-")
	item := cache.Get(key)
	if item != nil {
		return item.Cert, item.Key, nil
	}

	fmt.Printf("https=tre, should create local CA and create certificate.\n")
	if _, err := mkcert.Install(); err != nil {
		return "", "", err
	}
	info, err := mkcert.CreateCertificate(hosts)
	if err != nil {
		return "", "", err
	}
	os.Rename(info.CertFile, workDirCert+"/"+info.CertFile)
	os.Rename(info.KeyFile, workDirCert+"/"+info.KeyFile)

	cache.Set(key, &cert{Cert: info.CertFile, Key: info.KeyFile}, info.Expiration)
	return info.CertFile, info.KeyFile, nil
}

func printListener(isHttps bool, hosts []string, port int) {
	scheme := "http"
	if isHttps {
		scheme = "https"
	}

	fmt.Printf("Serve Listening at:\n")
	for _, v := range hosts {
		fmt.Printf("\t%s://%s:%d\n", scheme, v, port)
	}
}
