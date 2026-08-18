package services

import (
	"context"

	"players/db/generated"
	"players/models"
)

type PlayerService struct {
	Queries *generated.Queries
}

func NewPlayerService(queries *generated.Queries) *PlayerService {
	return &PlayerService{
		Queries: queries,
	}
}

func toPlayer(p generated.Player) models.Player {
	player := models.Player{
		PlayerID: p.Playerid,
	}

	if p.Birthyear.Valid {
		player.BirthYear = &p.Birthyear.Int32
	}
	if p.Birthmonth.Valid {
		player.BirthMonth = &p.Birthmonth.Int32
	}
	if p.Birthday.Valid {
		player.BirthDay = &p.Birthday.Int32
	}
	if p.Birthcountry.Valid {
		player.BirthCountry = &p.Birthcountry.String
	}
	if p.Birthstate.Valid {
		player.BirthState = &p.Birthstate.String
	}
	if p.Birthcity.Valid {
		player.BirthCity = &p.Birthcity.String
	}
	if p.Deathyear.Valid {
		player.DeathYear = &p.Deathyear.Int32
	}
	if p.Deathmonth.Valid {
		player.DeathMonth = &p.Deathmonth.Int32
	}
	if p.Deathday.Valid {
		player.DeathDay = &p.Deathday.Int32
	}
	if p.Deathcountry.Valid {
		player.DeathCountry = &p.Deathcountry.String
	}
	if p.Deathstate.Valid {
		player.DeathState = &p.Deathstate.String
	}
	if p.Deathcity.Valid {
		player.DeathCity = &p.Deathcity.String
	}
	if p.Namefirst.Valid {
		player.NameFirst = &p.Namefirst.String
	}
	if p.Namelast.Valid {
		player.NameLast = &p.Namelast.String
	}
	if p.Namegiven.Valid {
		player.NameGiven = &p.Namegiven.String
	}
	if p.Weight.Valid {
		player.Weight = &p.Weight.Int32
	}
	if p.Height.Valid {
		player.Height = &p.Height.Int32
	}
	if p.Bats.Valid {
		player.Bats = &p.Bats.String
	}
	if p.Throws.Valid {
		player.Throws = &p.Throws.String
	}
	if p.Debut.Valid {
		player.Debut = &p.Debut.Time
	}
	if p.Finalgame.Valid {
		player.FinalGame = &p.Finalgame.Time
	}
	if p.Retroid.Valid {
		player.RetroID = &p.Retroid.String
	}
	if p.Bbrefid.Valid {
		player.BbrefID = &p.Bbrefid.String
	}

	return player
}

func (s *PlayerService) GetPlayers(
	ctx context.Context,
	limit int32,
	offset int32,
) ([]models.Player, error) {
	players, err := s.Queries.GetPlayers(ctx, generated.GetPlayersParams{
		Limit:  limit,
		Offset: offset,
	})
	if err != nil {
		return nil, err
	}

	result := make([]models.Player, 0, len(players))

	for _, p := range players {
		result = append(result, toPlayer(p))
	}

	return result, nil
}

func (s *PlayerService) GetPlayerByID(
	ctx context.Context,
	playerID string,
) (models.Player, error) {
	player, err := s.Queries.GetPlayerByID(ctx, playerID)
	if err != nil {
		return models.Player{}, err
	}

	return toPlayer(player), nil
}

func (s *PlayerService) GetPlayersCount(
	ctx context.Context,
) (int64, error) {
	return s.Queries.GetPlayersCount(ctx)
}