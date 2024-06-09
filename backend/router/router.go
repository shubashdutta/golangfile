package router

import (
	"github/shubash/filemanger/controll"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/",controll.Home).Methods("GET")
	router.HandleFunc("/get-folder-structure",controll.GetFolderStructureHandler).Methods("GET")
	 router.HandleFunc("/api/create-folder",controll.CreateFolderHandler).Methods("POST")
	// router.HandleFunc("/api/singup", Controller.Signup).Methods("POST")
	// router.HandleFunc("/api/login", Controller.Login).Methods("POST")
	// router.HandleFunc("/api/users", Controller.GetAllUser).Methods("GET")
	// router.HandleFunc("/api/update/{id}", Controller.UPDATEPASSWORD).Methods("PUT")

	// router.HandleFunc("/api/deleteusers", Controller.DeleteAll).Methods("DELETE")

	return router
}