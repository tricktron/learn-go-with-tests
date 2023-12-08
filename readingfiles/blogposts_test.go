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

	inMemoryFS := fstest.MapFS{
		"hello world.md":  {Data: []byte("Title: Post 1")},
		"hello world2.md": {Data: []byte("Title: Post 2")},
	}
	want := blogposts.Post{Title: "Post 1"}

	posts, err := blogposts.NewPostsFromFS(inMemoryFS)
	if err != nil {
		t.Fatal(err)
	}

	if len(posts) != len(inMemoryFS) {
		t.Errorf("got %d posts, wanted %d posts", len(posts), len(inMemoryFS))
	}

	got := posts[0]

	assertPost(t, got, want)
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
