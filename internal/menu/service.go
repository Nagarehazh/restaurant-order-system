package menu

type Service interface {
	CreateMenuItem(req *MenuItem) error
	GetMenuItem(id uint) (*MenuItem, error)
	UpdateMenuItem(req *MenuItem) error
	DeleteMenuItem(id uint) error
	ListMenuItems() ([]MenuItem, error)
}

type MenuService struct {
	repo Repository
}

func NewMenuService(repo Repository) Service {
	return &MenuService{repo: repo}
}

func (s *MenuService) CreateMenuItem(item *MenuItem) error {
	return s.repo.CreateMenuItem(item)
}

func (s *MenuService) GetMenuItem(id uint) (*MenuItem, error) {
	return s.repo.GetMenuItem(id)
}

func (s *MenuService) UpdateMenuItem(item *MenuItem) error {
	return s.repo.UpdateMenuItem(item)
}

func (s *MenuService) DeleteMenuItem(id uint) error {
	return s.repo.DeleteMenuItem(id)
}

func (s *MenuService) ListMenuItems() ([]MenuItem, error) {
	return s.repo.ListMenuItems()
}
