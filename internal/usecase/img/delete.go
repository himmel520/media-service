package imgUC

import (
	"context"
	"fmt"
)

func (uc *ImgUC) Delete(ctx context.Context, id int) error {
	imageType, err := uc.repo.GetImageTypeByID(ctx, uc.db.DB(), id)
	if err != nil {
		return fmt.Errorf("get image type: %w", err)
	}

	if err := uc.repo.Delete(ctx, uc.db.DB(), id); err != nil {
		return fmt.Errorf("delete img: %w", err)
	}

	uc.DeleteImageCache(ctx, imageType)
	
	return nil
}
