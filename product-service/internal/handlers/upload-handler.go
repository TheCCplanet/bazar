package handlers

import (
	"bazar/config"
	"bazar/internal/model"
	"bazar/internal/service"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

var uploadDir string = "./uploads"

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:8080")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Access-Control-Allow-Credentials", "true")

	if r.Method == http.MethodOptions {
		log.Println("options is started")

		w.WriteHeader(http.StatusNoContent)
		return
	}
	log.Println("upload +")
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusLocked)
		return
	}
	userID, err := service.CheckTokenValidity(r)
	if err != nil {
		log.Println("auth error", err)
		http.Error(w, "Invalid token", http.StatusUnauthorized)
		return
	}
	r.Body = http.MaxBytesReader(w, r.Body, config.MaxUploadSize)
	if err := r.ParseMultipartForm(config.MaxUploadSize); err != nil {
		// request to large
		log.Println("Error: parsing form goes wrong:", err)
		http.Error(w, "Faild to parse multipart form", http.StatusBadRequest)
		return
	}
	product := &model.Product{
		SellerID:    userID,
		Title:       r.FormValue("title"),
		Description: r.FormValue("description"),
		Category:    r.FormValue("category"),
	}
	fmt.Println("product:", product)
	if product.Price, err = strconv.ParseFloat(r.FormValue("price"), 64); err != nil {
		log.Println("faild to format uploaded price to float64", err)
		return
	}
	quantityStr := r.FormValue("quantity")
	if quantityStr == "" {
		// Handle empty quantity case
		log.Println("quantity field is empty")
		return
	}

	// Correct parsing with proper error handling
	q, err := strconv.ParseInt(quantityStr, 10, 64)
	if err != nil {
		log.Println("failed to convert quantity to int64:", err)
		return
	}
	product.Quantity = q
	primaryImageID := r.FormValue("primaryImage")
	files := r.MultipartForm.File["images"]
	fmt.Println("image Count:", len(files))
	for _, fileHeader := range files {
		file, err := fileHeader.Open()
		if err != nil {
			log.Println(err)
			http.Error(w, "Error retrieving the file", http.StatusInternalServerError)
			return
		}
		defer file.Close()

		//ceck file type
		ext := filepath.Ext(fileHeader.Filename)
		newFileName := fmt.Sprintf("%d-%d%s", time.Now().UnixNano(), rand.Intn(10000), ext)
		filePath := filepath.Join(uploadDir, newFileName)
		// uploads/27374328478.png
		dst, err := os.Create(filePath)
		if err != nil {
			log.Println("Error create file ---", err)
			return
		}
		defer dst.Close()
		if _, err := io.Copy(dst, file); err != nil {
			log.Println("Faild to copy file in to the path:", err)
			return
		}
		product.Images = append(product.Images, model.ProductImage{
			URL:       newFileName,
			IsPrimary: fileHeader.Filename == primaryImageID,
		})
		fmt.Println("success ++++++++++++")

	}
	return
}
