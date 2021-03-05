package application

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/projectxpolaris/youtrans/config"
	"github.com/projectxpolaris/youtrans/service"
	"net/http"
)

var DefaultYouVideoClient = YouVideoClient{}

type YouVideoClient struct {
}

func (c *YouVideoClient) GetUrl(path string) string {
	return fmt.Sprintf("%s/%s", config.DefaultConfig.YouVideoUrl, path)
}

func (c *YouVideoClient) makePOSTRequest(url string, data interface{}, responseBody interface{}) error {
	rawData, err := json.Marshal(data)
	if err != nil {
		return err
	}
	response, err := http.Post(c.GetUrl(url), "application/json", bytes.NewBuffer(rawData))
	if responseBody != nil {
		err = json.NewDecoder(response.Body).Decode(&responseBody)
	}
	return err
}

func (c *YouVideoClient) makeGETRequest(url string, responseBody interface{}) error {
	response, err := http.Get(c.GetUrl(url))
	err = json.NewDecoder(response.Body).Decode(&responseBody)
	return err
}

func (c YouVideoClient) SendCompleteTask(task *service.Task) error {
	template := TaskTemplate{}
	template.Assign(task)
	err := c.makePOSTRequest("callback/tran/complete", template, nil)
	return err
}
