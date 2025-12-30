package ctfdgen

import (
	"context"
	"fmt"

	"github.com/Kinirok/ctfd-acc-gen/internal/ctfd"
	"github.com/Kinirok/ctfd-acc-gen/internal/gen"
	gormodel "github.com/Kinirok/ctfd-acc-gen/internal/storage"
)

func (g *Generator) CreateIndividualAccounts(ctx context.Context, emails []string, hasTeam bool, teamName *string, teamID *int) error {
	g.logger.Printf("Making %d accounts", len(emails))
	for _, email := range emails {
		resp, err := g.ctfdClient.CreateUser(ctx, ctfd.CreateUserRequest{Email: email, Name: gen.GenerateLogin(), Password: gen.GeneratePassword()})
		if err != nil {
			if resp.StatusCode == 400 {
				for range 5 {
					resp, err = g.ctfdClient.CreateUser(ctx, ctfd.CreateUserRequest{Email: email, Name: gen.GenerateLogin(), Password: gen.GeneratePassword()})
					if err == nil {
						break
					}
				}

			}
			if err != nil {
				g.logger.Printf("Failed to create account: %s", email)
				return err
			}
		}

		user := gormodel.Account{ID: resp.Data.ID, Email: resp.Data.Email, CTFDUser: resp.Data.CTFDUser, CTFDPass: resp.Data.CTFDPass, TeamName: teamName, TeamID: *teamID}
		if hasTeam {
			g.ctfdClient.AddUserToTeam(ctx, *teamID, int(resp.Data.ID))
		}
		err = gormodel.AddUserToDB(g.db, user)
		if err != nil {
			g.logger.Printf("Failed to add account data to db: %s", err.Error())
			return err
		}
	}
	return nil
}

func (g *Generator) CreateTeamAccounts(ctx context.Context, teamsCount, teamSize int) error {
	for range teamsCount {
		g.logger.Println("Making request to CTFd API")
		teamName := gen.GenerateTeamName()
		res, err := g.ctfdClient.CreateTeam(ctx, ctfd.CreateTeamRequest{TeamName: teamName})
		if err != nil {
			if res.StatusCode == 400 {
				for range 5 {
					teamName = gen.GenerateTeamName()
					res, err = g.ctfdClient.CreateTeam(ctx, ctfd.CreateTeamRequest{TeamName: teamName})
					if err == nil {
						break
					}
				}

			}
			if err != nil {
				g.logger.Println("Failed to create Team")
				return err
			}
		}
		g.logger.Println("Adding new team to DB")
		err = gormodel.AddTeamToDB(g.db, gormodel.Team{ID: res.Data.ID, TeamName: res.Data.Name})
		if err != nil {
			return err
		}
		emails := make([]string, teamSize)
		for i := range teamSize {
			emails[i] = fmt.Sprintf("%s_user%d@example.com", res.Data.Name, i)
		}
		g.logger.Printf("Team %d:", res.Data.ID)
		teamID := int(res.Data.ID)
		err = g.CreateIndividualAccounts(ctx, emails, true, &teamName, &teamID)
		if err != nil {
			return err
		}
	}

	return nil
}

func (g *Generator) GenerateNEmails(n int) []string {
	emails := make([]string, n)
	for i := range n {
		emails[i] = gen.GenerateEmail()
	}
	return emails
}
