package relayer

import (
	"context"
	"scroll-tech/common/types"

	"github.com/scroll-tech/go-ethereum/log"
)

func (r *Layer2Relayer) MockProver() {
	// Get batch to prove from db
	fields := map[string]interface{}{
		"proving_status": types.ProvingTaskUnassigned,
	}
	orderByList := []string{"index ASC"}
	limit := 1
	batches, err := r.batchOrm.GetBatches(r.ctx, fields, orderByList, limit)
	if err != nil {
		log.Error("Failed to get batches from db", "err", err)
		return
	}
	if len(batches) != 1 {
		log.Warn("Unexpeted result for GetBlockBatches", "number of batches", len(batches))
		return
	}

	// Update proving status to verified
	r.batchOrm.UpdateProvingStatus(context.Background(), batches[0].Hash, types.ProvingTaskVerified)
}