package service

import "errors"

var (
	ErrFailedToGetUser        = errors.New("failed to get user by id")
	ErrSquadAlreadyExists     = errors.New("User already has a squad")
	ErrInvalidPlayerCount     = errors.New("Squad must have exactly 15 players")
	ErrDuplicatePlayer        = errors.New("Duplicate player in squad")
	ErrNoCaptain              = errors.New("Squad must have exactly one captain")
	ErrNoViceCaptain          = errors.New("Squad must have exactly one vice captain")
	ErrCaptainIsViceCaptain   = errors.New("Captain and vice captain must be different")
	ErrInvalidFormation       = errors.New("Invalid squad formation")
	ErrBudgetExceeded         = errors.New("Squad exceeds budget limit")
	ErrPlayerNotFound         = errors.New("Player not found")
	ErrInvalidGoalkeeperCount = errors.New("Invalid goalkeeper count")
	ErrInvalidDefenderCount   = errors.New("Invalid defender count")
	ErrInvalidMidfielderCount = errors.New("Invalid midfielder count")
	ErrInvalidForwardCount    = errors.New("Invalid forward count")
	ErrSquadNotFound          = errors.New("Squad not found")
	ErrSquadPlayerNotFound    = errors.New("Squad player not found")
)
