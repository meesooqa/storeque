package locservice

import (
	"context"

	"github.com/meesooqa/storeque/common/lang"
)

type LocService interface {
	GetLoc(ctx context.Context, chatID int64) lang.Localization
}
