package animals

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

type Animal struct {
	Name string `json:"name"`
	Img  string `json:"img"`
}

func Routes() *chi.Mux {
	router := chi.NewRouter()
	router.Get("/", getImagemAnimals)

	return router
}

func getImagemAnimals(w http.ResponseWriter, r *http.Request) {
	animals := []Animal{
		{
			Name: "Dog",
			Img:  "dasdkoadoa",
		},
		{
			Name: "Cat",
			Img:  "dasdkoadoa",
		},
		{
			Name: "Fox",
			Img:  "dasdkoadoa",
		},
	}

	render.JSON(w, r, animals) // A chi router helper for serializing and returning json
}

func get() {
	response, err := http.Get("https://dog.ceo/api/breeds/image/random")
	if err != nil {
		fmt.Println("Error in get:", err)
	}
	defer response.Body.Close()

	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("%s", err)
	}

	fmt.Printf("%s\n", string(contents))
}
