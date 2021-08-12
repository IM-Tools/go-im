/**
  @author:panliang
  @data:2021/7/9
  @note
**/
package utils

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
)

type Header struct {
	Authorization string
	Token string
}

//POST /api/UploadImg HTTP/1.1
//Host: 127.0.0.1:9502
//Authorization: nk4NHLbNuERuGompX
//Content-Length: 247
//Content-Type: multipart/form-data; boundary=----WebKitFormBoundary7MA4YWxkTrZu0gW
//
//----WebKitFormBoundary7MA4YWxkTrZu0gW
//Content-Disposition: form-data; name="smfile"; filename="aHR0cHM6Ly9xaW5pdS5nYW9iaW56aGFuLmNvbS8yMDIwLzA0LzI3LzFkZGQ2ODIxY2UwYTkucG5n.png"
//Content-Type: image/png
//
//(data)
//----WebKitFormBoundary7MA4YWxkTrZu0gW

func PostFile(filename string, target_url string,headers *Header) (*http.Response, error) {
		body_buf := bytes.NewBufferString("")
		body_writer := multipart.NewWriter(body_buf)
		_, err := body_writer.CreateFormFile("smfile", filename)
		if err != nil {
			fmt.Println("error writing to buffer")
			return nil, err
		}
		fh, err := os.Open(filename)
		if err != nil {
			fmt.Println("error opening file")
			return nil, err
		}
		boundary := body_writer.Boundary()
		close_buf := bytes.NewBufferString(fmt.Sprintf("\r\n--%s--\r\n", boundary))

		request_reader := io.MultiReader(body_buf, fh, close_buf)
		fi, err := fh.Stat()
		if err != nil {
			fmt.Printf("Error Stating file: %s", filename)
			return nil, err
		}
		req, err := http.NewRequest("POST", target_url, request_reader)
		if err != nil {
			return nil, err
		}
		req.Header.Add(headers.Authorization,headers.Token)
		req.Header.Add("Content-Type", "multipart/form-data; boundary="+boundary)
		req.ContentLength = fi.Size() + int64(body_buf.Len()) + int64(close_buf.Len())
		return http.DefaultClient.Do(req)
}
