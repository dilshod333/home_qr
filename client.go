package main

import (
    "bytes"
    "fmt"
    "io"
    "mime/multipart"
    "net/http"
    "os"
)

func main() {
    url := "http://localhost:8080/generate"
    size := "256"
    content := "https://example.com"

 
    var b bytes.Buffer
    w := multipart.NewWriter(&b)
  
	
    fw, err := w.CreateFormField("size")
    if err != nil {
        fmt.Println("error creating :", err)
        return
    }
    _, err = fw.Write([]byte(size))
    if err != nil {
        fmt.Println("error writing :", err)
        return
    }
 
    fw, err = w.CreateFormField("content")
    if err != nil {
        fmt.Println("eerror creating", err)
        return
    }
    _, err = fw.Write([]byte(content))
    if err != nil {
        fmt.Println("errror writing :", err)
        return
    }
   
    w.Close()


    req, err := http.NewRequest("POST", url, &b)
    if err != nil {
        fmt.Println("error creating HTTP request:", err)
        return
    }
 
    req.Header.Set("Content-Type", w.FormDataContentType())

 
    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        fmt.Println("errror  HTTP request:", err)
        return
    }
    defer resp.Body.Close()

  
    if resp.StatusCode != http.StatusOK {
        fmt.Println("wronggg status code:", resp.StatusCode)
        return
    }

    outFile, err := os.Create("qrcode.png")
    if err != nil {
        fmt.Println("Error creating output file:", err)
        return
    }
    defer outFile.Close()

   
    _, err = io.Copy(outFile, resp.Body)
    if err != nil {
        fmt.Println("error saving the qr code image:", err)
        return
    }

    fmt.Println("qr code saved qr_code.png")
}
