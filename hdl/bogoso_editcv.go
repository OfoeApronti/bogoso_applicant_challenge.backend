package hdl

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"blow.com/bogoso.backend/app"
	log "github.com/sirupsen/logrus"
)

var BogosoEditApplication = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	logger := log.WithFields(log.Fields{"module": "hdl.BogosoEditApplication"})
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

	cv_id := r.FormValue("id")

	cv_name := r.FormValue("name")
	cv_email := r.FormValue("email")
	cv_phone := r.FormValue("phone")
	if cv_id == "" || cv_name == "" || cv_email == "" || cv_phone == "" {
		http.Error(w, "Some required fields are empty", http.StatusInternalServerError)
		return
	}
	var file_response_name string
	_, header, err := r.FormFile("cv_file")
	if err != nil {
		file_response_name = ""
	} else {
		fmt.Println(header.Filename)
		file_response_name = header.Filename
	}

	if file_response_name != "" {
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
		_, err = app.DBBogoso.Exec(`UPDATE bogoso.cv_files
	SET applicant_name=$2, email=$3, phone=$4, file_name=$5, created=now()
	WHERE id=$1`, cv_id, cv_name, cv_email, cv_phone, file_response_name)
		if err != nil {
			logger.Error(err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	} else {
		_, err = app.DBBogoso.Exec(`UPDATE bogoso.cv_files
	SET applicant_name=$2, email=$3, phone=$4, created=now()
	WHERE id=$1`, cv_id, cv_name, cv_email, cv_phone)
		if err != nil {

			logger.Error(err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
	fmt.Fprintf(w, "Upload successful")
})
