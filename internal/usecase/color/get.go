package colorUC

func (uc *ColorUC) GetAllWithPagination(ctx context.Context, limit, offset int) (*entity.ColorsResp, error) {
	colors, err := uc.repo.GetAllWithPagination(ctx, limit, offset)
	if err != nil {
		return nil, err
	}

	count, err := uc.repo.Count(ctx)
	if err != nil {
		return nil, err
	}

	return &entity.ColorsResp{
		Colors: colors,
		Total:  count,
	}, err
}