package internal

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"time"
)

type cache struct {
	file string
}

func newCache(file string) *cache {
	return &cache{file: file}
}

type cert struct {
	Cert string `json:"cert"`
	Key  string `json:"key"`
}

type item struct {
	Cert    *cert     `json:"cert"`
	Expired time.Time `json:"expired"`
}

func (r *cache) Get(key string) *cert {
	bs, err := ioutil.ReadFile(r.file)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		_ = os.Remove(r.file)
		return nil
	}
	resp := map[string]*item{}
	if err = json.Unmarshal(bs, &resp); err != nil {
		_ = os.Remove(r.file)
		return nil
	}

	v, ok := resp[key]
	if !ok {
		return nil
	}
	if v.Expired.Before(time.Now()) {
		delete(resp, key)
		bs, _ = json.Marshal(resp)
		_ = ioutil.WriteFile(r.file, bs, 0o666)
		return nil
	}

	return v.Cert
}

func (r *cache) Set(key string, cert *cert, expired time.Time) {
	bs, err := ioutil.ReadFile(r.file)
	if err != nil {
		if os.IsNotExist(err) {
			bs, _ = json.Marshal(map[string]*item{
				key: {
					Cert:    cert,
					Expired: expired,
				},
			})
			_ = ioutil.WriteFile(r.file, bs, 0o666)
			return
		}
		_ = os.Remove(r.file)
		return
	}
	resp := map[string]*item{}
	if err = json.Unmarshal(bs, &resp); err != nil {
		_ = os.Remove(r.file)
		return
	}

	resp[key] = &item{
		Cert:    cert,
		Expired: expired,
	}
	bs, _ = json.Marshal(resp)
	_ = ioutil.WriteFile(r.file, bs, 0o666)
	return
}
