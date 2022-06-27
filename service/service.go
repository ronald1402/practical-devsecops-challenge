package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"practical-devsecops-technical/config"
	"practical-devsecops-technical/dto"
)

func ExecuteJDoodle(input dto.JDoodleDto) ([]byte, int, error) {
	request := dto.JDoodleRequestDto{
		ClientId:     config.Config.JDoodle.ClientId,
		ClientSecret: config.Config.JDoodle.ClientSecret,
		ClientScript: input.Script,
		Language:     input.Language,
		VersionIndex: input.VersionIndex,
	}
	byteBody, _ := json.Marshal(request)
	url := fmt.Sprintf("%s/v1/execute", config.Config.JDoodle.Host)
	r, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(byteBody))
	if err != nil {
		return nil, 0, err
	}
	r.Header.Add("Content-Type", "application/json")
	cli := http.Client{}
	resp, err := cli.Do(r)
	if err != nil {
		return nil, 0, err
	}
	responseData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, 0, err
	}
	return responseData, resp.StatusCode, err
}
