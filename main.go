package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ipfs-cluster/ipfs-cluster/api"
	"github.com/ipfs-cluster/ipfs-cluster/api/rest/client"
	files "github.com/ipfs/go-libipfs/files"
	"github.com/multiformats/go-multiaddr"
)

func main() {

	apiAddr, err := multiaddr.NewMultiaddr("/ip4/127.0.0.1/tcp/9094")
	if err != nil {
		fmt.Println(err)
		return
	}
	cfg := &client.Config{
		APIAddr: apiAddr,
	}

	c, err := client.NewDefaultClient(cfg)
	if err != nil {
		fmt.Println(err)
		return
	}

	router := gin.Default()
	// sh := shell.NewShell("localhost:5001") // ipfs node api address

	router.POST("/upload", func(g *gin.Context) {
		form, _ := g.MultipartForm()
		filesHeaders := form.File["file"]

		for _, fileHeader := range filesHeaders {

			file, err := fileHeader.Open()
			if err != nil {
				g.String(http.StatusBadRequest, fmt.Sprintf("unable to open file: %s", err.Error()))
				return
			}

			defer file.Close()
			data, err := ioutil.ReadAll(file)
			if err != nil {
				g.String(http.StatusBadRequest, fmt.Sprintf("unable to read file: %s", err.Error()))
				return
			}

			params := api.DefaultAddParams()

			fileNode := files.NewBytesFile(data)
			dir := map[string]files.Node{fileHeader.Filename: fileNode}
			fileDir := files.NewMapDirectory(dir)
			out := make(chan api.AddedOutput)

			go func() {
				err = c.AddMultiFile(context.Background(), files.NewMultiFileReader(fileDir, true), params, out)
				if err != nil {
					g.String(http.StatusInternalServerError, fmt.Sprintf("error adding file: %s", err.Error()))
					return
				}
			}()

			for {
				select {
				case added, ok := <-out:
					if !ok {
						return
					}
					g.String(http.StatusOK, fmt.Sprintf("Added file: %s", added.Cid))
					return
				case <-time.After(10 * time.Second):
					g.String(http.StatusInternalServerError, "Add operation timed out")
					return
				}
			}
		}
	})

	// Retrieve a file by its CID
	router.GET("/file/:cid", func(g *gin.Context) {
		cid := g.Param("cid")
		resp, err := http.Get("http://127.0.0.1:8080/ipfs/" + cid)

		if err != nil {
			g.String(http.StatusInternalServerError, fmt.Sprintf("error retrieving file: %s", err.Error()))
			return
		}
		defer resp.Body.Close()

		data, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			g.String(http.StatusInternalServerError, fmt.Sprintf("error reading file: %s", err.Error()))
			return
		}

		g.Data(http.StatusOK, resp.Header.Get("Content-Type"), data)
	})

	router.Run(":8090")
}
