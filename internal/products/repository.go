package products

import (
	"fmt"
)

var Users []User

var lastID int

func (userRepository) LastID() int {
	return lastID
}

type Repository interface {
	GetAll() ([]User, error)
	Save(Request) (User, error)
	Update(id int, r Request) (User, error)
	Delete(id int) error
	UpdateSobrenomeIdade(id int, sobrenome string, idade int) (User, error)
}
type userRepository struct {
}

func NewUserRepository() Repository {
	return &userRepository{}
}

func (*userRepository) GetAll() ([]User, error) {
	if len(Users) > 0 {
		return Users, nil
	}
	return nil, fmt.Errorf("nenhum usuario encontrado")
}

func (*userRepository) Save(r Request) (User, error) {
	lastID++
	newUser := ReqToUser(r)
	newUser.Id = lastID

	Users = append(Users, newUser)
	return newUser, nil
}

func (*userRepository) Update(id int, r Request) (User, error) {
	newUser := ReqToUser(r)
	var updated bool

	for i, user := range Users {
		if id == user.Id {
			newUser.Id = id
			Users[i] = newUser
			updated = true
		}
	}
	if updated {
		return newUser, nil
	}
	return User{}, fmt.Errorf("id não encontrado")
}

func (*userRepository) Delete(id int) error {
	var deleted bool
	var toRemove int
	for i, user := range Users {
		if id == user.Id {
			deleted = true
			toRemove = i
		}
	}
	if !deleted {
		return fmt.Errorf("id não encontrado")
	}
	Users = append(Users[:toRemove], Users[toRemove+1:]...)
	return nil
}

func (*userRepository) UpdateSobrenomeIdade(id int, sobrenome string, idade int) (User, error) {
	var updated bool
	var editedUser User
	for i, user := range Users {
		if id == user.Id {
			Users[i].Sobrenome = sobrenome
			Users[i].Idade = idade
			updated = true
			editedUser = Users[i]
			break
		}
	}
	if !updated {
		return User{}, fmt.Errorf("id não encontrado")
	}
	return editedUser, nil
}
