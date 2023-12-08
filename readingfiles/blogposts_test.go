package blogposts_test

import (
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

	posts := blogposts.NewPostsFromFS(inMemoryFS)

	if len(posts) != len(inMemoryFS) {
		t.Errorf("got %d posts, wanted %d posts", len(posts), len(inMemoryFS))
	}
}
