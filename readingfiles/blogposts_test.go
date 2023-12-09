package blogposts_test

import (
	"errors"
	"io/fs"
	"reflect"
	"testing"
	"testing/fstest"

	blogposts "learn-go-with-tests/readingfiles"
)

func TestNewBlogPosts(t *testing.T) {
	t.Parallel()

	const (
		firstBody = `Title: Post 1
Description: Description 1`
		secondBody = `Title: Post 2
Description: Description 2`
	)

	inMemoryFS := fstest.MapFS{
		"hello world.md":  {Data: []byte(firstBody)},
		"hello world2.md": {Data: []byte(secondBody)},
	}

	posts, err := blogposts.NewPostsFromFS(inMemoryFS)
	if err != nil {
		t.Fatal(err)
	}

	if len(posts) != len(inMemoryFS) {
		t.Errorf("got %d posts, wanted %d posts", len(posts), len(inMemoryFS))
	}

	assertPost(t, posts[0], blogposts.Post{
		Title:       "Post 1",
		Description: "Description 1",
	})
}

type StubFailingFS struct{}

func (s StubFailingFS) Open(_ string) (fs.File, error) {
	return nil, errors.New("oh no, i always fail") //nolint: goerr113
}

func assertPost(t *testing.T, got blogposts.Post, want blogposts.Post) {
	t.Helper()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %+v want %+v", got, want)
	}
}
