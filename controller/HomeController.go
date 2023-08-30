package handler

import (
	"fmt"
	"forum/helper"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {

	err := helper.RenderTemplateWithLoyout(w, "pages/comment", "layout", nil)
	if err != nil {

		fmt.Println(err)
		return
	}

}
