package controllers

import (
	"gochat/app/svc"
	"net/http"

	"github.com/labstack/echo/v4"
)

type users struct {
	uSvc svc.IUsers
}

// NewUsersController will initialize the controllers
func NewUsersController(grp interface{}, uSvc svc.IUsers) {
	uc := &users{
		uSvc: uSvc,
	}

	g := grp.(*echo.Group)

	g.POST("/v1/user/signup", uc.Create)
}

func (ctr *users) Create(c echo.Context) error {
	// foundUser, getErr := GetUserByAppKey(c, ctr.uSvc)
	// if getErr != nil {
	// 	return c.JSON(getErr.Status, getErr)
	// }

	// var user domain.User

	// if err := c.Bind(&user); err != nil {
	// 	restErr := errors.NewBadRequestError("invalid json body")
	// 	return c.JSON(restErr.Status, restErr)
	// }

	// hashedPass, _ := bcrypt.GenerateFromPassword([]byte(*user.Password), 8)
	// *user.Password = string(hashedPass)
	// user.CompanyID = foundUser.CompanyID
	// user.RoleID = consts.RoleIDSales

	// result, saveErr := ctr.uSvc.CreateUser(user)
	// if saveErr != nil {
	// 	return c.JSON(saveErr.Status, saveErr)
	// }
	// var resp serializers.UserResp
	// respErr := methodsutil.StructToStruct(result, &resp)
	// if respErr != nil {
	// 	return respErr
	// }

	// return c.JSON(http.StatusCreated, resp)
	return c.JSON(http.StatusCreated, "user created")
}
