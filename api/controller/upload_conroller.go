package controller

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"sugam-project/api/responses"

	"github.com/google/uuid"
)

func (server *Server) handleFileupload(w http.ResponseWriter, r *http.Request) {
	// parse incomming image file

	file, fileheader, err := r.FormFile("image")

	if err != nil {
		log.Println("image upload error --> ", err)
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	// generate new uuid for image name
	uniqueId := uuid.New()

	// remove "- from imageName"

	filename := strings.Replace(uniqueId.String(), "-", "", -1)

	// extract image extension from original file filename

	fileExt := strings.Split(fileheader.Filename, ".")[1]

	// generate image from filename and extension
	image := fmt.Sprintf("%s.%s", filename, fileExt)

	// save image to ./images dir
	defer file.Close()
	tempFile, err := ioutil.TempFile("./uploads", image)
	if err != nil {
		log.Println("image upload error --> ", err)
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	defer tempFile.Close()
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		log.Println("image upload error --> ", err)
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	tempFile.Write(fileBytes)

	if err != nil {
		log.Println("image save error --> ", err)
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	// generate image url to serve to client using CDN

	imageUrl := fmt.Sprintf("http://localhost:4000/images/%s", image)

	// create meta data and send to client

	data := map[string]interface{}{

		"imageName": image,
		"imageUrl":  imageUrl,
		"header":    fileheader,
		"size":      fileheader.Size,
	}
	responses.JSON(w, http.StatusCreated, map[string]interface{}{
		"data":    data,
		"message": "Image uploaded successfully",
		"status":  http.StatusCreated,
		"success": true,
	})
}
