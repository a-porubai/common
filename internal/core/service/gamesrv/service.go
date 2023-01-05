package gamesrv

type Service struct {
	gamesRepository ports.GamesRepository
	uidGen          uidgen.UIDGen
}

func New(gamesRepository ports.GamesRepository, uidGen uidgen.UIDGen) *Service {
	return &Service{
		gamesRepository: gamesRepository,
		uidGen:          uidGen,
	}
}

func (srv *Service) Get(id string) (*domain.Game, error) {
	game, err := srv.gamesRepository.Get(id)
	if err != nil {
		if errors.Is(err, apperrors.NotFound) {
			return nil, errors.New("game not found")
		}

		return nil, err
	}

	game.Board = game.Board.HideBombs()

	return &game, nil
}

func (srv *Service) Create(name string, size uint, bombs uint) (*domain.Game, error) {
	return nil, nil
}

func (srv *Service) Reveal(id string, row uint, col uint) (*domain.Game, error) {
	return nil, nil
}
