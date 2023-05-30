package hdl

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"

	"blow.com/bogoso.backend/app"
	log "github.com/sirupsen/logrus"
)

const MAX_UPLOAD_SIZE = 1024 * 1024 //1MB
var BogosoCvUpload = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	logger := log.WithFields(log.Fields{"module": "hdl.BogosoCvUpload"})
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	//fmt.Println(r)
	err := r.ParseMultipartForm(128)
	//fmt.Println(r)
	if err != nil {
		// handle err
		http.Error(w, "Error decoding form values", http.StatusInternalServerError)
		return
	}

	cv_id, err := strconv.Atoi(r.FormValue("id"))
	if err != nil {
		// handle err
		http.Error(w, "Error reading id", http.StatusInternalServerError)
		return
	}
	cv_name := r.FormValue("name")
	cv_email := r.FormValue("email")
	cv_phone := r.FormValue("phone")
	_, header, _ := r.FormFile("cv_file")
	fmt.Println(header.Filename)
	file_response_name := header.Filename

	r.Body = http.MaxBytesReader(w, r.Body, MAX_UPLOAD_SIZE)
	if err := r.ParseMultipartForm(MAX_UPLOAD_SIZE); err != nil {
		http.Error(w, "The uploaded file is too big. Please choose an file that's less than 1MB in size", http.StatusBadRequest)
		return
	}

	//file, fileHeader, err := r.FormFile("file")
	file, _, err := r.FormFile("cv_file")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	defer file.Close()

	// Create the uploads folder if it doesn't
	// already exist
	err = os.MkdirAll("./uploads", os.ModePerm)
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Println(file_response_name)
	// Create a new file in the uploads directory
	//dst, err := os.Create(fmt.Sprintf("./uploads/%d%s", time.Now().UnixNano(), filepath.Ext(file_response_name)))
	dst, err := os.Create(fmt.Sprintf("./uploads/%s", file_response_name))
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer dst.Close()

	// Copy the uploaded file to the filesystem
	// at the specified destination
	_, err = io.Copy(dst, file)
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	_, err = app.DBBogoso.Exec("insert into bogoso.cv_files(id,applicant_name,email,phone,file_name) values($1,$2,$3,$4,$5)", cv_id, cv_name, cv_email, cv_phone, file_response_name)
	if err != nil {
		logger.Error(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "Upload successful")
})
