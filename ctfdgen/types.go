package ctfdgen

import (
	"context"

	gormodel "github.com/Kinirok/ctfd-acc-gen/internal/storage"
)

func (g *Generator) CheckUser(ctx context.Context, username string) (bool, error) {
	var id uint
	err := g.db.Model(&gormodel.Account{}).
		Where("CTFD_User = ?", username).
		Pluck("id", &id).Error
	if err != nil {
		return false, err
	}
	return g.ctfdClient.UserExists(ctx, id)
}

func (g *Generator) CheckTeam(ctx context.Context, teamname string) (bool, error) {
	var id uint
	err := g.db.Model(&gormodel.Team{}).
		Where("Team_Name = ?", teamname).
		Pluck("id", &id).Error
	if err != nil {
		return false, err
	}
	return g.ctfdClient.TeamExists(ctx, id)
}
