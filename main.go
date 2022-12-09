package main

import (
	"embed"
	"fmt"
	"gorm.io/gorm"
	"net/http"
	"os"
	"pergipergi/utils"
	"sync"
)

//go:embed views/*
var Resources embed.FS

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
