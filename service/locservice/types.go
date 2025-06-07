package locservice

import (
	"context"

	"github.com/meesooqa/storeque/common/lang"
)

// LocService defines the interface for localization services
type LocService interface {
	GetLoc(ctx context.Context, chatID int64) lang.Localization
}
