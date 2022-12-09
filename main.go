package main

import (
	"embed"
	"fmt"
	"gorm.io/gorm"
	"net/http"
	"os"
	"pergipergi/controller/rest"
	"pergipergi/repository"
	"pergipergi/service"
	"pergipergi/utils"
	"sync"
)

//go:embed views/*
var Resources embed.FS

type RestAPIHandler struct {
	UserRestAPIHandler rest.UserRestAPI
}

func main() {
	os.Setenv("DATABASE_URL", "postgres://postgres:Nero_Cld65@localhost:5432/golang_pergi_pergi")

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()

		mux := http.NewServeMux()

		err := utils.ConnectDB()
		if err != nil {
			panic(err)
		}

		db := utils.GetDBConnection()

		mux = RunServer(db, mux)
		mux = RunClient(mux, Resources)

		fmt.Println("Server is running on port 8080")
		err = http.ListenAndServe(":8080", mux)
		if err != nil {
			panic(err)
		}
	}()

	wg.Wait()
}

func RunServer(db *gorm.DB, mux *http.ServeMux) *http.ServeMux {
	userRepo := repository.NewUserRepository(db)
	//roleRepo := repository.NewRoleRepository(db)
	//destinasiRepo := repository.NewDestinasiRepository(db)
	//travelAgensiRepo := repository.NewTravelAgensiRepository(db)

	userService := service.NewUserService(userRepo)
	//roleService := service.NewRoleService(roleRepo)
	//destinasiService := service.NewDestinasiService(destinasiRepo)
	//travelAgensiService := service.NewTravelAgensiService(travelAgensiRepo)

	userRestAPI := rest.NewUserRestAPI(userService)

	apiHandler := RestAPIHandler{
		UserRestAPIHandler: userRestAPI,
	}

	MuxRoute(mux, "POST", "/api/v1/users/add", http.HandlerFunc(apiHandler.UserRestAPIHandler.AddUser))

	return mux
}

func RunClient(mux *http.ServeMux, embed embed.FS) *http.ServeMux {
	return mux
}

func MuxRoute(mux *http.ServeMux, method string, path string, handler http.Handler, opt ...string) {
	if len(opt) > 0 {
		fmt.Printf("[%s]: %s %v \n", method, path, opt)
	} else {
		fmt.Printf("[%s]: %s \n", method, path)
	}

	mux.Handle(path, handler)
}
