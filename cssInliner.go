package cssInliner

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

type Inliner struct {
	c http.Client
}

type ResponseCampaignMonitor struct {
	Html string `json:HTML`
}

const UrlCampaignMonitor = "https://inliner.cm/inline.php"

func NewInliner() *Inliner {
	return &Inliner{
		c: http.Client{},
	}
}

func (i *Inliner) Inline(dst io.Writer, src io.Reader) error {
	b, err := ioutil.ReadAll(src)
	if err != nil {
		return err
	}

	resp, err := i.c.PostForm(UrlCampaignMonitor, url.Values{
		"code": {string(b)},
	})
	if err != nil {
		return err
	}

	data, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		return err
	}

	res := ResponseCampaignMonitor{}
	err = json.Unmarshal(data, &res)
	if err != nil {
		return err
	}

	_, err = io.Copy(dst, strings.NewReader(res.Html))
	if err != nil {
		return err
	}

	return nil
}
