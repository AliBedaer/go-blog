package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/AliBedaer/go-blog/controllers"
	"github.com/AliBedaer/go-blog/models"
	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
)

// User contains user information
type User struct {
	FirstName      string     `validate:"required"`
	LastName       string     `validate:"required"`
	Age            uint8      `validate:"gte=0,lte=130"`
	Email          string     `validate:"required,email"`
	FavouriteColor string     `validate:"iscolor"`                // alias for 'hexcolor|rgb|rgba|hsl|hsla'
	Addresses      []*Address `validate:"required,dive,required"` // a person can have a home and cottage...
}

// Address houses a users address information
type Address struct {
	Street string `validate:"required"`
	City   string `validate:"required"`
	Planet string `validate:"required"`
	Phone  string `validate:"required"`
}

func routes() http.Handler {
	var postsController controllers.PostsController
	var usersController controllers.UsersController

	router := mux.NewRouter().StrictSlash(true)

	// posts Routes
	posts := router.PathPrefix("/posts/").Subrouter()
	posts.HandleFunc("/", postsController.ListAll).Methods("GET")
	posts.HandleFunc("/", postsController.NewPost).Methods("POST")
	// posts.HandleFunc("/", validateStruct).Methods("POST")
	posts.HandleFunc("/{id}", postsController.UpdatePost).Methods("PUT")
	//users Routes
	users := router.PathPrefix("/users/").Subrouter()
	users.HandleFunc("/", usersController.NewUser)

	// return
	return router
}

func validateStruct(writer http.ResponseWriter, request *http.Request) {

	var post models.Post

	err := json.NewDecoder(request.Body).Decode(&post)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(post)

	// address := &Address{
	// 	Street: "Eavesdown Docks",
	// 	Planet: "Persphone",
	// 	Phone:  "none",
	// }

	inputs := &models.Post{
		Title:   post.Title,
		Desc:    post.Desc,
		Content: post.Content,
	}

	// returns nil or ValidationErrors ( []FieldError )
	err = validate.Struct(inputs)

	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			fmt.Println(err)
			return
		} else {
			fmt.Println(err.(validator.ValidationErrors))
			return
		}

		for _, err := range err.(validator.ValidationErrors) {

			fmt.Println(err.Namespace())
			fmt.Println(err.Field())
			fmt.Println(err.StructNamespace()) // can differ when a custom TagNameFunc is registered or
			fmt.Println(err.StructField())     // by passing alt name to ReportError like below
			fmt.Println(err.Tag())
			fmt.Println(err.ActualTag())
			fmt.Println(err.Kind())
			fmt.Println(err.Type())
			fmt.Println(err.Value())
			fmt.Println(err.Param())
			fmt.Println()
		}

		// from here you can create your own error messages in whatever language you wish
		return
	}

}
