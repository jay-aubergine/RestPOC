package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, r.URL.Path[1:])
		fmt.Fprintln(w, "Welcome!")
	})

	http.HandleFunc("w2ui_rest/something", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Inside Something Handler")
	})

	http.HandleFunc("/w2ui_rest", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {

			r.ParseForm()

			var req RequestMapForm

			var requestData = r.FormValue("request")

			err := json.Unmarshal([]byte(requestData), &req)
			if err != nil {
				fmt.Println("UnmarshallErr")
				fmt.Fprintln(w, "{status:\"error\", message:", err, "}")
				return
				log.Fatal(err)
			}

			if req.Cmd == "save" {
				req, err := http.NewRequest("POST", "http://localhost:8080/rest", bytes.NewBuffer([]byte(requestData)))
				req.Header.Set("Content-Type", "application/json")
				client := &http.Client{}
				resp, err := client.Do(req)
				if err != nil {
					log.Fatal(err)
				}
				defer resp.Body.Close()
				body, _ := ioutil.ReadAll(resp.Body)

				fmt.Fprintln(w, string(body))
			}

			if req.Cmd == "get" {

				recid := strconv.Itoa(req.Recid)
				req, err := http.NewRequest("GET", "http://localhost:8080/rest", bytes.NewBuffer([]byte(requestData)))
				req.Header.Set("Content-Type", "application/json")
				client := &http.Client{}
				resp, err := client.Do(req)
				if err != nil {
					log.Fatal(err)
				}
				defer resp.Body.Close()
				body, _ := ioutil.ReadAll(resp.Body)

				if recid != "0" {
					var tem ResponseData
					err = json.Unmarshal(body, &tem)
					if err != nil {
						fmt.Println("UnmarshallErr")
						log.Fatal(err)
					}
					resStruct := SingleResponseStruct{
						"success", tem}
					res, err := json.Marshal(resStruct)
					if err != nil {
						fmt.Println("MarshallErr")
						log.Fatal(err)
					}
					fmt.Fprintln(w, string(res))
					return
				}

				var tem RestResponseStruct
				//var tem []ResponseData
				err = json.Unmarshal(body, &tem)
				if err != nil {
					fmt.Println("UnmarshallErr")
					log.Fatal(err)
				}

				for i := 0; i < len(tem.ResponseData); i++ {
					tem.ResponseData[i].Recid = i + 1
				}

				resStruct := ResponseStruct{
					"success", len(tem.ResponseData), tem.ResponseData}

				/*resStruct := ResponseStruct{
				"success", len(tem), tem}*/

				res, err := json.Marshal(resStruct)
				if err != nil {
					fmt.Println("MarshallErr")
					log.Fatal(err)
				}
				fmt.Fprintln(w, string(res))
			}

		} else if r.Method == "GET" {

		}
	})

	http.HandleFunc("/rest", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			r.ParseForm()

			defer r.Body.Close()

			body, err := ioutil.ReadAll(r.Body)
			if err != nil {
				panic(err)
			}

			var reqMap RequestMapForm
			err = json.Unmarshal(body, &reqMap)
			if err != nil {
				fmt.Fprintln(w, "{status:\"error\", message:", err, "}")
				return
			}
			fmt.Fprintln(w, "{\"status\": \"success\"}")

		} else if r.Method == "GET" {
			body, _ := ioutil.ReadAll(r.Body)
			fmt.Printf("Bodyyyyy: %+v\n", string(body))

			var reqMap RequestMapForm
			err := json.Unmarshal(body, &reqMap)
			if err != nil {
				fmt.Fprintln(w, "{status:\"error\", message:", err, "}")
				return
			}

			res := getRestData(reqMap.Recid)
			fmt.Fprintln(w, string(res))
		}
	})

	http.HandleFunc("/w2ui", func(w http.ResponseWriter, r *http.Request) {

		if r.Method == "POST" {
			r.ParseForm()
			var req RequestMapForm
			err := json.Unmarshal([]byte(r.FormValue("request")), &req)
			if err != nil {
				fmt.Println("UnmarshallErr")
				log.Fatal(err)
			}

			switch req.Cmd {
			case "get":
				res := getData()
				fmt.Fprintln(w, string(res))

			case "save":
				fmt.Fprintln(w, "save")

			case "delete":
				fmt.Fprintln(w, "del")

			default:
				fmt.Fprintln(w, "default case")
			}
		} else if r.Method == "GET" {
			fmt.Fprintln(w, "W2UI GetMethod")
		}
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}

/*		fmt.Printf("URL: %+v\n", r.URL)
		fmt.Printf("Proto: %+v\n", r.Proto)
		fmt.Printf("Header: %+v\n", r.Header)
		fmt.Printf("Body: %+v\n", r.Body)
		fmt.Printf("GetBody: %+v\n", r.GetBody)
		fmt.Printf("ContentLength: %+v\n", r.ContentLength)
		fmt.Printf("Host: %+v\n", r.Host)
		fmt.Printf("Form: %+v\n", r.Form)
		fmt.Printf("PostForm: %+v\n", r.PostForm)
		fmt.Printf("Trailer: %+v\n", r.Trailer)
		fmt.Printf("RemoteAddr: %+v\n", r.RemoteAddr)
		fmt.Printf("RequestURI: %+v\n", r.RequestURI)
		fmt.Printf("TLS: %+v\n", r.TLS)
		fmt.Println("---------------------------------------------")
*/
