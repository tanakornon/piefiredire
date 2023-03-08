package inmemory

import "piefiredire/internal/core/ports"

type beefRepository struct {
	text string
}

func NewBeefRepository() ports.BeefRepository {
	return beefRepository{
		text: "Fatback t-bone t-bone, t-bone-bone pastrami  ..   t-bone.  pork, meatloaf jowl enim.  Bresaola t-bone.",
	}
}

func (repo beefRepository) GetText() (string, error) {
	return repo.text, nil
}
