package tgUC

import "context"

func (uc *TgUC) Delete(ctx context.Context, id int) error {
	return uc.repo.Delete(ctx, uc.db.DB(), id)
}