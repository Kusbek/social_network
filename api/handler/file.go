package handler

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func uploadFile(pathToImage string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Maximum upload of 10 MB files
		r.ParseMultipartForm(10 << 20)

		// Get handler for filename, size and headers
		file, handler, err := r.FormFile("myFile")
		if err != nil {
			errorResponse(w, http.StatusBadRequest, err)
			return
		}
		defer file.Close()

		// Create file
		dst, err := os.Create(fmt.Sprintf("./dist/img/%v/%v", pathToImage, handler.Filename))
		if err != nil {
			errorResponse(w, http.StatusInternalServerError, err)
			return
		}
		defer dst.Close()

		// Copy the uploaded file to the created file on the filesystem
		if _, err := io.Copy(dst, file); err != nil {
			errorResponse(w, http.StatusInternalServerError, err)
			return
		}

		successResponse(w, http.StatusCreated, map[string]interface{}{
			"path_to_photo": fmt.Sprintf("/img/%v/%v", pathToImage, handler.Filename),
		})
	})
}

func MakeFileHandlers(r *http.ServeMux) {
	r.Handle("/api/upload", uploadFile("avatars"))
}
