package controller

import (
	"clean-arch/models"
	"clean-arch/service"
	"encoding/json"
	"net/http"
)

var (
	postService service.PostService = service.NewPostService()
)

func getPosts(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-type", "application/json")
	posts, err := postService.FindAll()
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte(`{"error": "Error getting posts"}`))
		return
	}
	resp.WriteHeader(http.StatusOK)
	json.NewEncoder(resp).Encode(posts)
}

func addPost(resp http.ResponseWriter, req *http.Request) {
	var post models.Post
	resp.Header().Set("Content-type", "application/json")
	err := json.NewDecoder(req.Body).Decode(&post)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte(`{"error": "Error marshelling the posts array"}`))
		return
	}
	err = postService.Validate(&post)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte(`{"error": "Error marshelling data"}`))
		return
	}
	postService.Create(&post)
	resp.WriteHeader(http.StatusOK)
	json.NewEncoder(resp).Encode(post)
}
