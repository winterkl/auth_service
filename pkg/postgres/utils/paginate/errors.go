package paginate

type InvalidPage struct{}

func (p *InvalidPage) Error() string {
	return "Введен некорректный номер страницы"
}
