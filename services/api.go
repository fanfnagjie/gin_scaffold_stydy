package services

import (
	"github.com/e421083458/gin_scaffold_study/dao"
	"github.com/e421083458/gin_scaffold_study/dto"
	"github.com/e421083458/gin_scaffold_study/middleware"
	"github.com/e421083458/gin_scaffold_study/golang_common/lib"
	"github.com/gin-gonic/gin"
)

type ApiService struct {
}


func (as *ApiService) AddUser(c *gin.Context, input *dto.AddUserInput) error {
	tx, err := lib.GetGormPool("default")
	if err != nil {
		middleware.ResponseError(c, 2002, err)
		return err
	}
	user := &dao.User{
		Name:  input.Name,
		Sex:   input.Sex,
		Age:   input.Age,
		Birth: input.Birth,
		Addr:  input.Addr,
	}
	if err := user.Save(c, tx); err != nil {
		middleware.ResponseError(c, 2002, err)
		return err
	}
	return nil
}
