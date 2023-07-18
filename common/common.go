package common

import (
	"encoding/json"
	"go-blog/config"
	"go-blog/models"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
)

var Template models.HtmlTemplate

//帮助处理常见的操作，如加载模板、解析 JSON 参数以及返回错误和成功的结果
func LoadTemplate() {
	//初始化
	w := sync.WaitGroup{}
	w.Add(1)
	go func() {
		var err error
		Template, err = models.InitTemplate(config.Cfg.System.CurrentDir + "/template/")
		if err != nil {
			panic(err)
		}
		w.Done()
	}()
	w.Wait()
}

func GetRequestJsonParam(r *http.Request) map[string]interface{} {
	var params map[string]interface{}
	body, _ := ioutil.ReadAll(r.Body)
	_ = json.Unmarshal(body, &params)
	//fmt.Printf("Request Body: %s\n", string(body))
	return params
}

func Error(w http.ResponseWriter, err error) {
	var result models.Result
	result.Code = -999
	result.Error = err.Error()
	resultJson, _ := json.Marshal(result)
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(resultJson)
	if err != nil {
		log.Println(err)
	}
}

func Success(w http.ResponseWriter, data interface{}) {
	var result models.Result
	result.Code = 200
	result.Error = ""
	result.Data = data
	resultJson, _ := json.Marshal(result)
	w.Header().Set("Content-Type", "application/json")
	_, err := w.Write(resultJson)
	if err != nil {
		log.Println(err)
	}
}
