package user

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"

	"github.com/giantswarm/graph/api/response"
	"github.com/giantswarm/graph/ent"
	"github.com/giantswarm/graph/ent/githubuser"
)

type ControllerConfig struct {
	GraphClient *ent.Client
}

type Controller struct {
	graphClient *ent.Client
}

func NewController(config ControllerConfig) (*Controller, error) {
	if config.GraphClient == nil {
		return nil, errors.New(fmt.Sprintf("%T.GraphClient must not be empty", config))
	}

	c := Controller{
		graphClient: config.GraphClient,
	}

	return &c, nil
}

func (c *Controller) Create(ctx *gin.Context, githubUser ent.GitHubUser) {
	res, err := c.graphClient.GitHubUser.
		Create().
		SetLogin(githubUser.Login).
		SetName(githubUser.Name).
		SetEmail(githubUser.Email).
		SetGithubID(githubUser.GithubID).
		Save(ctx)

	if err == nil {
		ctx.JSON(200, &res)
	} else {
		ctx.JSON(500, gin.H{
			"error": err.Error(),
		})
	}
}

func (c *Controller) Update(ctx *gin.Context, githubUser ent.GitHubUser) {
	res, err := c.graphClient.GitHubUser.UpdateOne(&githubUser).Save(ctx)
	if err == nil {
		ctx.JSON(200, &res)
	} else {
		ctx.JSON(500, gin.H{
			"error": err.Error(),
		})
	}
}

func (c *Controller) Get(ctx *gin.Context, username string) {
	log.Println("Username ", username)
	githubUser, err := c.graphClient.GitHubUser.
		Query().
		Where(githubuser.Login(username)).
		Only(ctx)

	if ent.IsNotFound(err) {
		ctx.JSON(200, response.Get{
			Message: "GitHub user not found",
		})
	} else if err == nil {
		ctx.JSON(200, response.Get{
			Result: githubUser,
		})
	} else {
		ctx.JSON(500, gin.H{
			"error": err.Error(),
		})
	}
}
