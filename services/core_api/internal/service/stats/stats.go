package stats

import (
	"core_api/graph/model"
	"core_api/internal/repository"
	"pkg/models"
	"time"

	"gorm.io/datatypes"
)

var d datatypes.Date

type StatsService struct {
	statsRepository *repository.StatsRepository
}

func NewStatsService(statsRepository *repository.StatsRepository) *StatsService {
	return &StatsService{statsRepository: statsRepository}
}

func (ss *StatsService) GetMaches() ([]*model.GameMatch, error) {
	var gameWeeks []models.GameWeek

	if err := ss.statsRepository.GetGameWeeks(&gameWeeks); err != nil {
		return nil, err
	}

	gameweekMatches := make([]*model.GameMatch, 0, len(gameWeeks))

	for _, gameweek := range gameWeeks {
		gameMatch := model.GameMatch{
			Round: int32(gameweek.Round),
		}
		for _, match := range gameweek.Matches {
			gameMatch.Matches = append(gameMatch.Matches, &model.Match{
				Date:          time.Time(match.Date).Format("2006-01-02"),
				Time:          match.Time.String(),
				Completed:     match.Completed,
				HomeTeamScore: int32(match.HomeScore),
				AwayTeamScore: int32(match.AwayScore),

				Gameweek: mapGameWeek(match.GameWeek),
				HomeTeam: mapTeam(match.HomeTeam),
				AwayTeam: mapTeam(match.AwayTeam),
			})
		}
		gameweekMatches = append(gameweekMatches, &gameMatch)
	}
	return gameweekMatches, nil
}

func (ss *StatsService) GetTable() ([]*model.Table, error) {
	var tables []models.Table

	if err := ss.statsRepository.GetTables(&tables); err != nil {
		return nil, err
	}

	leagueTable := make([]*model.Table, 0, len(tables))
	for _, table := range tables {
		table := model.Table{
			Team:               mapTeam(table.Team),
			Position:           int32(table.Position),
			Played:             int32(table.Played),
			Wins:               int32(table.Wins),
			Losses:             int32(table.Losses),
			ScoresFor:          int32(table.ScoresFor),
			ScoresAgainst:      int32(table.ScoresAgainst),
			Draws:              int32(table.Draws),
			Points:             int32(table.Points),
			ScoreDiffFormatted: int32(table.ScoreDiffFormatted),
		}
		leagueTable = append(leagueTable, &table)
	}
	return leagueTable, nil
}

func mapTeam(t models.Team) *model.Team {
	return &model.Team{
		ID:   int32(t.ID),
		Name: t.Name,
	}
}

func mapGameWeek(gw models.GameWeek) *model.GameWeek {
	return &model.GameWeek{
		Round: int32(gw.Round),
	}
}
