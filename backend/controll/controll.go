package controll

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
)


// https://golangfile-2.onrender.com/ //  hosted in this url 


const BASE_PATH = "C:\\root"

func GetFolderStructure(parentFolderPath string) (map[string][]string, error) {
    folders := []string{}
    files := []string{}

    entries, err := os.ReadDir(parentFolderPath)
    if err != nil {
        return nil, err
    }

    for _, entry := range entries {
        if entry.IsDir() {
            folders = append(folders, entry.Name())
        } else {
            files = append(files, entry.Name())
        }
    }

    return map[string][]string{
        "folders": folders,
        "files":   files,
    }, nil
}


type CreateFolderRequest struct {
	FolderName       string `json:"folderName"`
	ParentFolderPath string `json:"parentFolderPath"`
}

type CreateFolderResponse struct {
	Message    string `json:"message"`
	FolderPath string `json:"folderPath"`
}

func CreateFolderHandler(w http.ResponseWriter, r *http.Request) {
	// Set CORS headers
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	// Parse the request body
	var req CreateFolderRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	// Determine the folder path
	var folderPath string
	if req.ParentFolderPath != "" {
		folderPath = filepath.Join(req.ParentFolderPath, req.FolderName)
	} else {
		folderPath = filepath.Join(BASE_PATH, req.FolderName)
	}

	// Create the folder
	if err := os.MkdirAll(folderPath, os.ModePerm); err != nil {
		log.Println("Error creating folder:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	// Send the response
	resp := CreateFolderResponse{
		Message:    "Folder created at: " + folderPath,
		FolderPath: folderPath,
	}
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		log.Println("Error encoding response:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}
 func Home( w http.ResponseWriter ,r *http.Request){
fmt.Fprintf(w,"hello")
 }



 
 func GetFolderStructureHandler(w http.ResponseWriter, r *http.Request) {
	 // Set CORS headers
	 w.Header().Set("Access-Control-Allow-Origin", "*")
	 w.Header().Set("Content-Type", "application/json")
 
	 parentFolderPath := r.URL.Query().Get("parentFolderPath")
	 if parentFolderPath == "" {
		 parentFolderPath = BASE_PATH
	 }
 
	 folderStructure, err := getFolderStructure(parentFolderPath)
	 if err != nil {
		 http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		 log.Println("Error getting folder structure:", err)
		 return
	 }
 
	 json.NewEncoder(w).Encode(folderStructure)
 }
 
 func getFolderStructure(parentFolderPath string) (map[string][]string, error) {
	 folders := []string{}
	 files := []string{}
 
	 entries, err := os.ReadDir(parentFolderPath)
	 if err != nil {
		 return nil, err
	 }
 
	 for _, entry := range entries {
		 if entry.IsDir() {
			 folders = append(folders, entry.Name())
		 } else {
			 files = append(files, entry.Name())
		 }
	 }
 
	 return map[string][]string{
		 "folders": folders,
		 "files":   files,
	 }, nil
 }
