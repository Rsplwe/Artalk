package aliyun

import (
	"encoding/json"
	"fmt"
)

type GreenTextConf struct {
	AccessKeyId     string
	AccessKeySecret string
	Region          string
	Content         string
	DataID          string
}

func GreenText(conf GreenTextConf) (bool, error) {
	if conf.Region == "" {
		conf.Region = "cn-shanghai"
	}
	body, err := RequestApi(AliyunApiRequestConf{
		AccessKeyId:     conf.AccessKeyId,
		AccessKeySecret: conf.AccessKeySecret,
		BaseURL:         "https://green-cip." + conf.Region + ".aliyuncs.com",
		Path:            "/",
		Version:         "2022-03-02",
		Data: map[string]interface{}{
			"Service": []string{"comment_detection"},
			"ServiceParameters": map[string]interface{}{
				"content": conf.Content,
			},
		},
	})
	if err != nil {
		return false, err
	}

	var data GreenTextResponse
	err = json.Unmarshal(body, &data)
	if err != nil {
		return false, err
	}

	if data.Code != 200 {
		return false, fmt.Errorf(data.Message)
	}

	if len(data.Data) < 1 {
		return false, fmt.Errorf("not expected response: %s", string(body))
	}

	if data.Data[0].Reason == "" && data.Data[0].Labels == "" {
		return true, nil
	} else {
		return false, nil
	}
}

type GreenTextResponse struct {
	Code int `json:"Code"`
	Data []struct {
		Reason string `json:"reason"`
		Labels string `json:"labels"`
	} `json:"Data"`
	Message   string `json:"Message"`
	RequestID string `json:"RequestId"`
}
