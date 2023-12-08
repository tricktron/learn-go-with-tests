package blogposts_test

import (
	"errors"
	"io/fs"
	"testing"
	"testing/fstest"

	blogposts "learn-go-with-tests/readingfiles"
)

func TestNewBlogPosts(t *testing.T) {
	t.Parallel()

	inMemoryFS := fstest.MapFS{
		"hello world.md":  {Data: []byte("hi")},
		"hello world2.md": {Data: []byte("hola")},
	}

	posts, err := blogposts.NewPostsFromFS(inMemoryFS)
	if err != nil {
		t.Fatal(err)
	}

	if len(posts) != len(inMemoryFS) {
		t.Errorf("got %d posts, wanted %d posts", len(posts), len(inMemoryFS))
	}
}

type StubFailingFS struct{}

func (s StubFailingFS) Open(_ string) (fs.File, error) {
	return nil, errors.New("oh no, i always fail") //nolint: goerr113
}
