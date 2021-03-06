package main

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"time"
)

var (
	site  string
	token string
	b     string
)

func main() {
	site = os.Getenv("URL")
	token = os.Getenv("TOKEN")
	b = os.Getenv("b")

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

	// fmt.Println("response Status:", response.Status)
	// fmt.Println("response Headers:", response.Header)
	body, _ := ioutil.ReadAll(response.Body)
	// fmt.Println("response Body:", string(body))
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
		log.Panic(err)
	}
	time.Sleep(30 * time.Second)

	log.Println("??????????????????")

	id1 := id.Data.FileOperation.ID

	downloadUrl := site + "/api/fileOperations.redirect?id=" + id1

	requestdwon, err := http.NewRequest("GET", downloadUrl, nil)

	requestdwon.Header.Set("Authorization", "Bearer "+token)

	clientdown := &http.Client{}
	resp, err := clientdown.Do(requestdwon)

	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	//???????????????????????????
	// ??????????????????????????????
	var filename = "/backup/" + time.Now().Format("2006-1-2-150405") + ".zip"
	out, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer out.Close()

	// ??????????????????????????????????????????
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		panic(err)
	}

	if b != "" {
		command := `./deleteold.sh`
		cmd := exec.Command("/bin/sh", "-c", command)

		output, err := cmd.Output()
		if err != nil {
			log.Printf("Execute Shell:%s failed with error:%s", command, err.Error())
			return
		}
		log.Printf("Execute Shell:%s finished with output:\n%s", command, string(output))

	}

	log.Println("????????????????????????", filename)
}
