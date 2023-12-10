package blogposts

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/fs"
	"strings"
)

type Post struct {
	Title       string
	Description string
	Tags        []string
	Body        string
}

const (
	titleSeparator       = "Title: "
	descriptionSeparator = "Description: "
	tagsSeparator        = "Tags: "
)

func NewPostsFromFS(fileSystem fs.FS) ([]Post, error) {
	dir, err := fs.ReadDir(fileSystem, ".")
	if err != nil {
		return nil, err //nolint: wrapcheck
	}

	var posts []Post //nolint: prealloc

	for _, file := range dir {
		post, err := getPost(fileSystem, file.Name())
		if err != nil {
			return nil, err
		}

		posts = append(posts, post)
	}

	return posts, nil
}

func getPost(fileSystem fs.FS, fileName string) (Post, error) {
	postFile, err := fileSystem.Open(fileName)
	if err != nil {
		return Post{}, err //nolint: wrapcheck
	}
	defer postFile.Close()

	return newPost(postFile)
}

func newPost(postFile io.Reader) (Post, error) {
	scanner := bufio.NewScanner(postFile)

	readMetaLine := func(tagName string) string {
		scanner.Scan()

		return strings.TrimPrefix(scanner.Text(), tagName)
	}

	readBody := func() string {
		scanner.Scan()

		buf := bytes.Buffer{}

		for scanner.Scan() {
			fmt.Fprintln(&buf, scanner.Text())
		}

		return strings.TrimSuffix(buf.String(), "\n")
	}

	return Post{
			Title:       readMetaLine(titleSeparator),
			Description: readMetaLine(descriptionSeparator),
			Tags:        strings.Split(readMetaLine(tagsSeparator), ", "),
			Body:        readBody(),
		},
		nil
}
