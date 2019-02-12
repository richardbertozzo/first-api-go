package animals

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

const (
	API_DOG = "https://dog.ceo/api/breeds/image/random"
	API_CAT = "https://aws.random.cat/meow"
	API_FOX = "https://randomfox.ca/floof/"
)

type Animal struct {
	Name string `json:"name"`
	Img  string `json:"img"`
}

func Routes() *chi.Mux {
	router := chi.NewRouter()
	router.Get("/", getAnimals)

	return router
}

func getAnimals(w http.ResponseWriter, r *http.Request) {
	channelDog := make(chan string)
	go getDog(channelDog)
	dog := Animal{
		Name: "Dog",
		Img:  <-channelDog,
	}

	channelCat := make(chan string)
	go getCat(channelCat)
	cat := Animal{
		Name: "Cat",
		Img:  <-channelCat,
	}

	channelFox := make(chan string)
	go getFox(channelFox)
	fox := Animal{
		Name: "Fox",
		Img:  <-channelFox,
	}

	animals := []Animal{
		dog,
		cat,
		fox,
	}

	render.JSON(w, r, animals) // A chi router helper for serializing and returning json
}

func getDog(chanDog chan string) {
	type DogResponse struct {
		Status string `json:"status"`
		Link   string `json:"message"`
	}

	var dogResponse DogResponse
	response := getResponse(API_DOG)
	json.Unmarshal(response, &dogResponse)

	chanDog <- dogResponse.Link
}

func getCat(chanCat chan string) {
	type CatResponse struct {
		Link string `json:"file"`
	}

	var catResponse CatResponse
	response := getResponse(API_CAT)
	json.Unmarshal(response, &catResponse)

	chanCat <- catResponse.Link
}

func getFox(chanFox chan string) {
	type FoxResponse struct {
		Image string `json:"image"`
		Link  string `json:"link"`
	}

	var foxResponse FoxResponse
	response := getResponse(API_FOX)
	json.Unmarshal(response, &foxResponse)

	chanFox <- foxResponse.Image
}

func getResponse(url string) []byte {
	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Error in get:", err)
	}
	defer response.Body.Close()

	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error in read body", err)
	}

	return contents
}
