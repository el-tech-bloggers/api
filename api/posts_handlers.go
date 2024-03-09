package api

import (
	"context"
	"el-tech-bloggers/api/pg"
	"el-tech-bloggers/api/types"
	"encoding/json"
	"fmt"
	"net/http"
)

var (
	ctx         = context.Background()
	postgres, _ = pg.NewPG(ctx)
)

func GetPostsHandler(w http.ResponseWriter, r *http.Request) {
	posts, err := postgres.GetPosts(ctx)
	if err != nil {
		http.Error(w, "unable to get posts", http.StatusInternalServerError)
		return
	}
	fmt.Fprint(w, posts)
}

func CreatePostHandler(w http.ResponseWriter, r *http.Request) {
	var post types.Post
	err := json.NewDecoder(r.Body).Decode(&post)
	if err != nil {
		http.Error(w, "unable to parse request body", http.StatusBadRequest)
		return
	}
	err = postgres.CreatePost(ctx, post)
	if err != nil {
		http.Error(w, "unable to create post", http.StatusInternalServerError)
		return
	}
	fmt.Fprint(w, "Post created successfully!")
}

func GetPostHandler(w http.ResponseWriter, r *http.Request) {
	postId := r.PathValue("id")
	post, err := postgres.GetPost(ctx, postId)
	if err != nil {
		fmt.Println("Error: ", err)
		http.Error(w, "unable to get post "+postId, http.StatusInternalServerError)
		return
	}
	fmt.Fprint(w, *post)
}
