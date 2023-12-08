package blogposts

import (
	"io"
	"io/fs"
)

type Post struct {
	Title string
}

func NewPostsFromFS(fileSystem fs.FS) ([]Post, error) {
	dir, err := fs.ReadDir(fileSystem, ".")
	if err != nil {
		return nil, err //nolint: wrapcheck
	}

	var posts []Post //nolint: prealloc

	for _, file := range dir {
		post, err := getPost(fileSystem, file)
		if err != nil {
			return nil, err
		}

		posts = append(posts, post)
	}

	return posts, nil
}

func getPost(fileSystem fs.FS, file fs.DirEntry) (Post, error) {
	postFile, err := fileSystem.Open(file.Name())
	if err != nil {
		return Post{}, err //nolint: wrapcheck
	}
	defer postFile.Close()

	postData, err := io.ReadAll(postFile)
	if err != nil {
		return Post{}, err //nolint: wrapcheck
	}

	post := Post{Title: string(postData)[7:]}

	return post, nil
}
