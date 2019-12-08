package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	helper "github.com/AliBedaer/go-blog/helpers"
	"github.com/AliBedaer/go-blog/models"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

type PostsController struct{}

var Posts []models.Post
var validate *validator.Validate

func main() {

}

// ListAll : List All Posts
func (controller *PostsController) ListAll(response http.ResponseWriter, request *http.Request) {
	response.Write([]byte("List Posts"))
}

// NewPost : Create new Post
func (controller *PostsController) NewPost(response http.ResponseWriter, request *http.Request) {
	validate = validator.New()
	var post models.Post
	resp := make(map[string]interface{})
	errors := make([]map[string]interface{}, 0)
	defer request.Body.Close()
	// requestBody, err := ioutil.ReadAll(request.Body)
	err := json.NewDecoder(request.Body).Decode(&post)
	helper.ValidateError(err)

	inputs := &models.Post{
		Title:   post.Title,
		Desc:    post.Desc,
		Content: post.Content,
	}
	err = validate.Struct(inputs)

	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			fmt.Println(err)
			return
		} else {
			fmt.Println(err.(validator.ValidationErrors))
			resp["Message"] = "Validation Errors"
			for _, err := range err.(validator.ValidationErrors) {
				errors = append(errors, map[string]interface{}{err.Field(): fmt.Sprintf("Field %s is %s", err.Field(), err.Tag())})
				// resp["errors"] = []map[string]interface{}{err.Field(): fmt.Sprintf("Field %s is %s", err.Field(), err.Tag())}
			}
			resp["errors"] = errors
			response.Header().Set("Content-Type", "application/json")
			json.NewEncoder(response).Encode(resp)
			return
		}
	}

	// post.CreatePost(post)
	resp["Message"] = "Post Created Successfully"
	response.Header().Set("Content-Type", "application/json")
	json.NewEncoder(response).Encode(resp)

}

// UpdatePost : Update post by Id
func (controller *PostsController) UpdatePost(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	json.NewEncoder(response).Encode(vars)
	// response.Write([]byte("update post"))
}

// DeletePost : Delete Post
func (controller *PostsController) DeletePost(response http.ResponseWriter, request *http.Request) {

	response.Write([]byte("Delete Post"))
}
