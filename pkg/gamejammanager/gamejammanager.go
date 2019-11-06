package gamejammanager

import (
	"github.com/blockchain-abstraction-middleware/game-jam-abstraction/pkg/contracts/GameJamManager"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

// Manager contains the login for the Manager Contract
type Manager struct {
	ec *ethclient.Client

	auth *bind.TransactOpts

	manager *GameJamManager.GameJamManager

	address string
}

// NewManager Creates a game manager contract with a goven address and private key
func NewManager(ec *ethclient.Client, gmAddress string, auth *bind.TransactOpts) (*Manager, error) {
	gameJamManager, err := GameJamManager.NewGameJamManager(common.HexToAddress(gmAddress), ec)

	return &Manager{
		auth:    auth,
		ec:      ec,
		manager: gameJamManager,
		address: gmAddress,
	}, err
}

// GetAllGameJamAddresses gets all deployed game jam addresses
func (m *Manager) GetAllGameJamAddresses() ([]common.Address, error) {
	return m.manager.GetAllGameJamAddresses(&bind.CallOpts{Pending: true})
}
