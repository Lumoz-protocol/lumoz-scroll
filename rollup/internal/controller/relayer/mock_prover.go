package relayer

import (
	"context"
	"scroll-tech/common/types"
)

func (r *Layer2Relayer) MockProver() {
	r.batchOrm.UpdateProvingStatus(context.Background(), batchHash2, types.ProvingTaskVerified)
}