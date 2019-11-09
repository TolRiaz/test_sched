package main

import (
	"encoding/json"
	"fmt"
	"gopkg.in/resty.v1"
	"strconv"
)

const (
	LatURL string = "http://220.149.11.93:30303/api/v1/query?query=golang_latency_ovs_"
)

func LatRequest(nodeNum int) (string, error) {
	NodeNum := strconv.Itoa(nodeNum)
	resp, err := resty.R().Get(LatURL + NodeNum)
	if err != nil {
		return "", err
	}

	result, err := parseStringtoStruct(resp.Body())
	if err != nil {
		return "", err
	}

	value := result.Data.Result[0].Value[1]
	strValue := fmt.Sprintf("%v", value)
	return strValue, nil
}

func parseStringtoStruct(body []byte) (*NodeLat, error) {
	var nl NodeLat
	err := json.Unmarshal(body, &nl)
	if err != nil {
		return nil, err
	}
	if nl.Status != "success" {
		return nil, err
	}
	return &nl, nil
}

type NodeLat struct {
	Status string `json:"status"`
	Data   struct {
		ResultType string `json:"resultType"`
		Result     []struct {
			Metric struct {
				PodName             string `json:"__name__"`
				App                 string `json:"app"`
				Instance            string `json:"instance"`
				Job                 string `json:"job"`
				KubernetesName      string `json:"kubernetes_name"`
				KubernetesNamespace string `json:"kubernetes_namespace"`
				Name                string `json:"name"`
			} `json:"metric"`
			Value []interface{} `json:"value"`
		} `json:"result"`
	} `json:"data"`
}

func main() {
	fmt.Println(LatRequest(1))
}
