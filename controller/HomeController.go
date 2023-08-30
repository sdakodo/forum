package controller

import (
	"fmt"
	"forum/helper"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {

	err := helper.RenderTemplateWithLoyout(w, "pages/home", "layout", nil)
	if err != nil {

		fmt.Println(err)
		return
	}

}
func Comment(w http.ResponseWriter, r *http.Request) {

	err := helper.RenderTemplateWithLoyout(w, "pages/comment", "layout", nil)
	if err != nil {

		fmt.Println(err)
		return
	}

}
func Home2(w http.ResponseWriter, r *http.Request) {

	err := helper.RenderTemplateWithLoyout(w, "pages/home2", "layout", nil)
	if err != nil {

		fmt.Println(err)
		return
	}

}

func Login(w http.ResponseWriter, r *http.Request) {

	err := helper.RenderTemplateWithLoyout(w, "pages/loging", "layoutRegisterAndloging", nil)
	if err != nil {

		fmt.Println(err)
		return
	}

}
func Register(w http.ResponseWriter, r *http.Request) {

	err := helper.RenderTemplateWithLoyout(w, "pages/register", "layoutRegisterAndloging", nil)
	if err != nil {

		fmt.Println(err)
		return
	}

}
