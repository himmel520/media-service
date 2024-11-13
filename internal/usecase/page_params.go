package usecase

type PageParams struct {
	// Номер страницы для пагинации.
    Page uint64
    // Количество категорий на странице.
    PerPage uint64
}