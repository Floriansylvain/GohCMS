package api

import (
	"GoCMS/domain/post"
	"github.com/go-chi/chi/v5"
	"html/template"
	"net/http"
	"os"
	"strconv"
)

type PostEditPageAlert struct {
	IsError bool
	Message string
}

func getPostEditPageTemplate(post post.Post, alert PostEditPageAlert) []byte {
	navbarTmpl, _ := Container.GetPageUseCase.GetPage("componentNavbar", nil)
	postTmpl, _ := Container.GetPageUseCase.GetPage("postEdit", map[string]interface{}{
		"Navbar":  template.HTML(navbarTmpl),
		"Head":    headTmpl,
		"Post":    post,
		"Alert":   alert,
		"Secured": os.Getenv("ENVIRONMENT") == "production",
	})
	return postTmpl
}

func PostPostEditPage(w http.ResponseWriter, r *http.Request) {
	postID := chi.URLParam(r, "id")
	postIDint, err := strconv.Atoi(postID)
	if err != nil {
		_, _ = w.Write(getPostEditPageTemplate(post.Post{}, PostEditPageAlert{
			IsError: true,
			Message: "Could not find the requested post.",
		}))
		return
	}

	_ = r.ParseForm()
	postBody := r.FormValue("postBody")

	getPost, _ := Container.GetPostUseCase.GetPost(uint32(postIDint))
	updatedPost, err := Container.UpdatePostUseCase.UpdateBody(getPost.ID, postBody)
	if err != nil {
		_, _ = w.Write(getPostEditPageTemplate(post.Post{Body: postBody}, PostEditPageAlert{
			IsError: true,
			Message: "Could not save the post: " + err.Error(),
		}))
		return
	}

	_, _ = w.Write(getPostEditPageTemplate(updatedPost, PostEditPageAlert{
		IsError: false,
		Message: "Post successfully edited!",
	}))
}

func GetPostEditPage(w http.ResponseWriter, r *http.Request) {
	postID := chi.URLParam(r, "id")
	postIDint, err := strconv.Atoi(postID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	getPost, _ := Container.GetPostUseCase.GetPost(uint32(postIDint))

	_, _ = w.Write(getPostEditPageTemplate(getPost, PostEditPageAlert{
		IsError: false,
		Message: "",
	}))
}
