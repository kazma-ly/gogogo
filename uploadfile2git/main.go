package main

import (
	"context"
	"crypto/sha1"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/google/go-github/v31/github"
	"golang.org/x/oauth2"
)

var (
	// Need Write
	accessToken = "Github AccessToken"

	// Info
	branch           = "master"
	repoName         = "kazma233.github.io"
	owner            = "kazma233"
	defaultCommitMsg = "upload file via go client"
	email            = "zly.private@hotmail.com"
	pathDateFormat   = "20060102"
	walkPath         = "./images"
)

func main() {

	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: accessToken},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)

	rootPath := time.Now().Format(pathDateFormat)

	filepath.Walk(walkPath, func(path string, info os.FileInfo, err error) error {

		if err != nil {
			panic(err)
		}

		if !info.IsDir() {
			uploadPath := rootPath + "/" + info.Name()

			file, err := os.OpenFile(path, os.O_RDONLY, 0755)
			if err != nil {
				panic(err)
			}

			fileByte, err := ioutil.ReadAll(file)
			if err != nil {
				panic(err)
			}

			_sha := sha1.New()
			_, err = _sha.Write(fileByte)
			if err != nil {
				panic(err)
			}

			date := time.Now()
			_, _, err = client.Repositories.CreateFile(ctx, owner, repoName, uploadPath, &github.RepositoryContentFileOptions{
				Message: github.String(defaultCommitMsg),
				Content: fileByte,
				Branch:  github.String(branch),
				Author: &github.CommitAuthor{
					Date:  &date,
					Name:  github.String(owner),
					Email: github.String(email),
				},
				Committer: &github.CommitAuthor{
					Date:  &date,
					Name:  github.String(owner),
					Email: github.String(email),
				},
				// SHA: github.String(string(_sha.Sum(nil))),
			})

			if err != nil {
				panic(err)
			}

			log.Printf("%v upload success, url is: %v", info.Name(), "https://raw.githubusercontent.com/"+owner+"/"+repoName+"/"+branch+"/"+uploadPath)
		}

		return nil
	},
	)

}
