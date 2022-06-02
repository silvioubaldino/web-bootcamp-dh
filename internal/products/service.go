package products

type Service interface {
	GetAll() ([]User, error)
	Save(r Request) (User, error)
	Update(id int, r Request) (User, error)
	Delete(id int) error
	UpdateSobrenomeIdade(id int, sobrenome string, idade int) (User, error)
}

type userService struct {
	rep Repository
}

func NewUserService(rep Repository) Service {
	return &userService{rep}
}

func (s *userService) GetAll() ([]User, error) {
	users, err := s.rep.GetAll()
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (s *userService) Save(r Request) (User, error) {
	user, err := s.rep.Save(r)
	if err != nil {
		return user, err
	}
	return user, nil
}

func (s *userService) Update(id int, r Request) (User, error) {
	user, err := s.rep.Update(id, r)
	if err != nil {
		return user, err
	}
	return user, nil
}

func (s *userService) Delete(id int) error {
	err := s.rep.Delete(id)
	return err
}

func (s *userService) UpdateSobrenomeIdade(id int, sobrenome string, idade int) (User, error) {
	user, err := s.rep.UpdateSobrenomeIdade(id, sobrenome, idade)
	return user, err
}
