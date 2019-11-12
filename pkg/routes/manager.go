package routes

import (
	"encoding/json"
	"net/http"

	"github.com/blockchain-abstraction-middleware/game-jam-abstraction/pkg/config"
	"github.com/blockchain-abstraction-middleware/game-jam-abstraction/pkg/ethereum"
	"github.com/blockchain-abstraction-middleware/game-jam-abstraction/pkg/gamejammanager"
	log "github.com/blockchain-abstraction-middleware/project-template/pkg/logger"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/go-chi/chi"
)

// ManagerResource implements server.Resource interface
type ManagerResource struct {
	path   string
	config *config.Config
}

// NewResource constructor func
func (r *ManagerResource) NewResource(path string, config *config.Config) *ManagerResource {
	return &ManagerResource{
		path:   path,
		config: config,
	}
}

// Path returns the Resource base path
func (r *ManagerResource) Path() string {
	return r.path
}

// Routes bootstraps health routes
func (r *ManagerResource) Routes() http.Handler {
	router := chi.NewRouter()

	privKey, err := crypto.HexToECDSA(r.config.Keys.Admin)
	if err != nil {
		log.WithError(err).Error("Failed to load key")
	}

	ec, err := ethereum.CreateEthClient(r.config.Infura.URL + r.config.Infura.Key)
	if err != nil {
		log.WithError(err).Error("Failed to initialize ethereum client")
	}

	auth := bind.NewKeyedTransactor(privKey)

	gameJamManagerContract, err := gamejammanager.NewManager(ec, r.config.Contracts.GameManagerAddress, auth)
	if err != nil {
		log.WithError(err).Error("Failed to create game jam manager")
	}

	router.Get("/get-all-gamejams", r.getAllAddresses(gameJamManagerContract))

	log.WithFields(log.Fields{"Contract": "Game Manager", "Address": r.config.Contracts.GameManagerAddress}).Info("Created manager abstraction")

	return router
}

func (r *ManagerResource) getAllAddresses(gm *gamejammanager.Manager) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {

		gameJamAddressList, err := gm.GetAllGameJamAddresses()
		if err != nil {
			log.WithError(err).Error("GetAllGameJamAddresses function on Game Manager contract failed")
		}

		payload := struct {
			AddressList string `json:"addressList"`
			StatusCode  int    `json:"statusCode"`
		}{
			gameJamAddressList[0].String(),
			200,
		}

		json, err := json.Marshal(payload)

		res.Header().Set("Content-Type", "application/json")
		res.WriteHeader(200)
		res.Write(json)
	}
}
