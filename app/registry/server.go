package registry

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"
)

const ServerPort = ":3000"
const ServiceURL = "http://localhost" + ServerPort + "/services"

// Registry

type registry struct {
	registrations []Registration
	mutex         *sync.Mutex
}

func (r *registry) add(reg Registration) error {
	r.mutex.Lock()
	/* We're not using a deferred call here. We're going to need to manually manipulate this mutex because we're going to have things that are called from
	the add method that are also going to manipulate that registration slice. So we need to make sure that we've got full control over the mutex. That's why. */
	r.registrations = append(r.registrations, reg)
	r.mutex.Unlock() // TODO: check what happens if we defer this call
	return nil
}

var reg = registry{
	registrations: make([]Registration, 0),
	mutex:         new(sync.Mutex),
}

// Registry Service

type RegistryService struct{}

func (s RegistryService) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println("Request received")
	switch r.Method {
	case http.MethodPost:
		dec := json.NewDecoder(r.Body)
		var reg Registration
		if err := dec.Decode(&reg); err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		log.Printf("Adding service: %v with URL: %v\n", reg.ServiceName, reg.ServiceURL)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
}
