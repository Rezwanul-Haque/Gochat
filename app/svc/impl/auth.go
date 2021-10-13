package impl

import (
	"gochat/app/repository"
	"gochat/app/serializers"
	"gochat/app/svc"
)

type auth struct {
	urepo repository.IUsers
}

func NewAuthService(urepo repository.IUsers) svc.IAuth {
	return &auth{
		urepo: urepo,
	}
}

func (as *auth) Login(req *serializers.LoginReq) (*serializers.LoginResp, error) {
	// var user *domain.User
	// var err error

	// if user, err = as.urepo.GetUserByEmail(req.Email); err != nil {
	// 	return nil, errors.ErrInvalidEmail
	// }

	// if req.Admin && !as.urepo.HasRole(user.ID, consts.RoleIDAdmin) {
	// 	return nil, errors.ErrNotAdmin
	// }

	// loginPass := []byte(req.Password)
	// hashedPass := []byte(*user.Password)

	// if err = bcrypt.CompareHashAndPassword(hashedPass, loginPass); err != nil {
	// 	logger.Error(err.Error(), err)
	// 	return nil, errors.ErrInvalidPassword
	// }

	// var token *serializers.JwtToken

	// if token, err = as.tSvc.CreateToken(user.ID, user.CompanyID); err != nil {
	// 	logger.Error(err.Error(), err)
	// 	return nil, errors.ErrCreateJwt
	// }

	// if err = as.tSvc.StoreTokenUuid(user.ID, user.CompanyID, token); err != nil {
	// 	logger.Error(err.Error(), err)
	// 	return nil, errors.ErrStoreTokenUuid
	// }

	// if err = as.urepo.SetLastLoginAt(user); err != nil {
	// 	logger.Error("error occur when trying to set last login", err)
	// 	return nil, errors.ErrUpdateLastLogin
	// }

	// var userResp *serializers.UserWithParamsResp

	// if userResp, err = as.getUserInfoWithParam(user.ID, user.CompanyID, false); err != nil {
	// 	return nil, err
	// }

	// res := &serializers.LoginResp{
	// 	AccessToken:  token.AccessToken,
	// 	RefreshToken: token.RefreshToken,
	// 	User:         userResp,
	// }
	return nil, nil
}
