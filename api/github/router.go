package github

import (
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"

	"github.com/giantswarm/graph/api/github/user"
	"github.com/giantswarm/graph/ent"
)

var (
	userController *user.Controller
)

func SetupRoutes(r *gin.Engine, graphClient *ent.Client) error {
	var err error

	userControllerConfig := user.ControllerConfig{
		GraphClient: graphClient,
	}
	userController, err = user.NewController(userControllerConfig)
	if err != nil {
		return errors.Wrapf(err, "failed to create GitHubUser controller with config %v", userControllerConfig)
	}

	// Get GitHubUser entity
	r.GET("/github/user/:login", func(ctx *gin.Context) {
		type GitHubUser struct {
			Login string `uri:"login" binding:"required"`
		}

		githubUser := GitHubUser{}
		err = ctx.ShouldBindUri(&githubUser)
		if err != nil {
			ctx.JSON(400, map[string]interface{}{
				"message": "Unknown error occurred",
				"error":   err,
			})
			return
		}

		userController.Get(ctx, githubUser.Login)
	})

	// Create GitHubUser entity
	r.POST("/github/user", func(ctx *gin.Context) {
		githubUser := ent.GitHubUser{}
		err := ctx.ShouldBind(&githubUser)
		if err != nil {
			ctx.JSON(400, map[string]interface{}{
				"message": "Unknown error occurred",
				"error":   err,
			})
			return
		}

		userController.Create(ctx, githubUser)
	})

	// Update GitHubUser entity
	r.PUT("/github/user", func(ctx *gin.Context) {
		githubUser := ent.GitHubUser{}
		err := ctx.ShouldBind(&githubUser)
		if err != nil {
			ctx.JSON(400, map[string]interface{}{
				"message": "Unknown error occurred",
				"error":   err,
			})
			return
		}

		userController.Update(ctx, githubUser)
	})

	return nil
}
