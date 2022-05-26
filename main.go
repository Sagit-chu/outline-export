package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {

	var site string
	var token string
	site = os.Getenv("URL")
	token = os.Getenv("TOKEN")

	httpposturl := site + "/api/collections.export_all"

	var jsonData = []byte(`{
		"name": "morpheus",
		"job": "leader"
	}`)
	request, error := http.NewRequest("POST", httpposturl, bytes.NewBuffer(jsonData))
	request.Header.Set("Authorization", "Bearer "+token)

	client := &http.Client{}
	response, error := client.Do(request)
	if error != nil {
		panic(error)
	}
	defer response.Body.Close()

	fmt.Println("response Status:", response.Status)
	fmt.Println("response Headers:", response.Header)
	body, _ := ioutil.ReadAll(response.Body)
	fmt.Println("response Body:", string(body))

	type FruitBasket struct {
		Success bool `json:"success"`
		Data    struct {
			FileOperation struct {
				ID           string      `json:"id"`
				Type         string      `json:"type"`
				Name         string      `json:"name"`
				State        string      `json:"state"`
				Error        interface{} `json:"error"`
				Size         string      `json:"size"`
				CollectionID interface{} `json:"collectionId"`
				User         struct {
					ID           string    `json:"id"`
					Name         string    `json:"name"`
					AvatarURL    string    `json:"avatarUrl"`
					Color        string    `json:"color"`
					IsAdmin      bool      `json:"isAdmin"`
					IsSuspended  bool      `json:"isSuspended"`
					IsViewer     bool      `json:"isViewer"`
					CreatedAt    time.Time `json:"createdAt"`
					UpdatedAt    time.Time `json:"updatedAt"`
					LastActiveAt time.Time `json:"lastActiveAt"`
				} `json:"user"`
				CreatedAt time.Time `json:"createdAt"`
				UpdatedAt time.Time `json:"updatedAt"`
			} `json:"fileOperation"`
		} `json:"data"`
		Status int  `json:"status"`
		Ok     bool `json:"ok"`
	}

	var id FruitBasket
	err := json.Unmarshal(body, &id)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("ID是 ", id.Data.FileOperation.ID)

	time.Sleep(30 * time.Second)

	id1 := id.Data.FileOperation.ID

	downloadUrl := site + "/api/fileOperations.redirect?id=" + id1

	println("下载地址是：", downloadUrl)

	requestdwon, err := http.NewRequest("GET", downloadUrl, nil)

	requestdwon.Header.Set("Authorization", "Bearer "+token)

	clientdown := &http.Client{}
	resp, err := clientdown.Do(requestdwon)

	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	//新建时间，为文件名
	// 创建一个文件用于保存
	var filename = "/backup/" + time.Now().Format("2006-1-2-150405") + ".zip"
	out, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer out.Close()

	// 然后将响应流和文件流对接起来
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		panic(err)
	}
}
