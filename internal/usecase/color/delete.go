package colorUC

import "context"

func (uc *ColorUC) Delete(ctx context.Context, id int) error {
	return uc.repo.Delete(ctx, uc.db.DB(), id)
}
