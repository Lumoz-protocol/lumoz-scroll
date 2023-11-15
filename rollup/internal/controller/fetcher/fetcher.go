package fetcher

import (
	"context"
	"fmt"
	"scroll-tech/rollup/internal/config"
	"scroll-tech/rollup/internal/controller/fetcher/smartcontracts"

	"github.com/scroll-tech/go-ethereum/ethclient"
)

// Fetcher fetch contract information from l1/l2 geth
type Fetcher struct {
	client  *ethclient.Client // The client to retrieve on chain data or send transaction.
	
	ScrollChain *smartcontracts.ScrollChain
}

// NewFetcher returns a new instance of contract information fetcher
func NewFetcher(ctx context.Context, cfg *config.RelayerConfig) (*Fetcher, error) {
	ethClient, err := ethclient.Dial(cfg.SenderConfig.Endpoint)
	if err != nil {
		return nil, fmt.Errorf("fail to new ethClient when new layer2 relayer, err: %w", err)
	}
	scrollchain, err := smartcontracts.NewScrollChain(cfg.RollupContractAddress, ethClient)
	if err != nil {
		return nil, fmt.Errorf("fail to new smartcontracts scrollchain when new layer2 relayer, err: %w", err)
	}

	fetcher := &Fetcher{
		client:        ethClient,
		ScrollChain: scrollchain,
	}

	return fetcher, nil
}



