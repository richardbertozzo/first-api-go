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
	apiDog = "https://dog.ceo/api/breeds/image/random"
	apiCat = "https://aws.random.cat/meow"
	apiFox = "https://randomfox.ca/floof/"
)

// Animal represents an animal
type Animal struct {
	Name string `json:"name"`
	Img  string `json:"img"`
}

// Routes make the routes of package
func Routes() *chi.Mux {
	router := chi.NewRouter()
	router.Get("/", getAnimals)

	return router
}

func getAnimals(w http.ResponseWriter, r *http.Request) {
	channelDog := make(chan string)
	channelCat := make(chan string)
	channelFox := make(chan string)
	go getDog(channelDog)
	go getCat(channelCat)
	go getFox(channelFox)

	animals := []Animal{
		{
			Name: "Dog",
			Img:  <-channelDog,
		},
		{
			Name: "Cat",
			Img:  <-channelCat,
		},
		{
			Name: "Fox",
			Img:  <-channelFox,
		},
	}

	render.JSON(w, r, animals) // A chi router helper for serializing and returning json
}

func getDog(chanDog chan string) {
	type DogResponse struct {
		Status string `json:"status"`
		Link   string `json:"message"`
	}

	var dogResponse DogResponse
	response := getResponse(apiDog)
	json.Unmarshal(response, &dogResponse)

	chanDog <- dogResponse.Link
}

func getCat(chanCat chan string) {
	type CatResponse struct {
		Link string `json:"file"`
	}

	var catResponse CatResponse
	response := getResponse(apiCat)
	json.Unmarshal(response, &catResponse)

	chanCat <- catResponse.Link
}

func getFox(chanFox chan string) {
	type FoxResponse struct {
		Image string `json:"image"`
		Link  string `json:"link"`
	}

	var foxResponse FoxResponse
	response := getResponse(apiFox)
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
