package update

import (
	"context"

	"github.com/STT196/ScopeSentry/internal/repositories/apikey"
)

func Update20() {
	repo := apikey.NewRepository()
	_ = repo.EnsureIndexes(context.Background())
}
