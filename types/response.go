package types

import "github.com/gin-gonic/gin"

type Response struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
}

const (
	// success 200
	SuccessCode      = "20000"
	LoginSuccessCode = "20001"
	ForgotPassCode   = "20002"
	ResetPassCode    = "20003"

	// 201
	UserCreatedCode   = "20101"
	EmailVerifiedCode = "20102"
	MFAVerifiedCode   = "20103"

	// 400
	InvalidRequestCode  = "40001"
	AlreadyExistsCode   = "40002"
	InvalidEmailCode    = "40003"
	AlreadyVerifiedCode = "40004"

	// 401
	InvalidCredentialsCode  = "40101"
	InvalidVerificationCode = "40102"
	EmailNotVerifiedCode    = "40103"
	InvalidSessionCode      = "40104"

	// 404
	UserNotFoundCode = "40401"

	// 429
	TooManyRequestsCode = "42901"

	// 500
	InternalServerErrorCode = "50001"
)

// func Success() Response
func Success() Response {
	return Response{
		Code: SuccessCode,
		Msg:  "Success",
	}
}

// func LoginSuccess() Response
func LoginSuccess() Response {
	return Response{
		Code: LoginSuccessCode,
		Msg:  "Login successful",
	}
}

// func ForgotPass() Response
func ForgotPass() Response {
	return Response{
		Code: ForgotPassCode,
		Msg:  "Password reset email sent",
	}
}

// func PassReset() Response
func ResetPass() Response {
	return Response{
		Code: ResetPassCode,
		Msg:  "Password reset",
	}
}

// func UserCreated() Response
func UserCreated() Response {
	return Response{
		Code: UserCreatedCode,
		Msg:  "User created, please verify your email",
	}
}

// func EmailVerified() Response
func EmailVerified() Response {
	return Response{
		Code: EmailVerifiedCode,
		Msg:  "Email verified",
	}
}

// func MFAVerified() Response
func MFAVerified() Response {
	return Response{
		Code: MFAVerifiedCode,
		Msg:  "MFA verified",
	}
}

// func InvalidRequest() Resopnse
func InvalidRequest() Response {
	return Response{
		Code: InvalidRequestCode,
		Msg:  "Invalid request",
	}
}

// func AlreadyExists() Response
func AlreadyExists() Response {
	return Response{
		Code: AlreadyExistsCode,
		Msg:  "Already exists",
	}
}

// func InvalidCredentials() Response
func InvalidCredentials() Response {
	return Response{
		Code: InvalidCredentialsCode,
		Msg:  "Invalid credentials",
	}
}

// func InvalidVerification() Response
func InvalidVerification() Response {
	return Response{
		Code: InvalidVerificationCode,
		Msg:  "Invalid verification code",
	}
}

// func AlreadyVerified() Response
func AlreadyVerified() Response {
	return Response{
		Code: AlreadyVerifiedCode,
		Msg:  "Email already verified",
	}
}

// func EmailNotVerified() Response
func EmailNotVerified() Response {
	return Response{
		Code: EmailNotVerifiedCode,
		Msg:  "Email not verified",
	}
}

// func InvalidSession() Response
func InvalidSession() Response {
	return Response{
		Code: InvalidSessionCode,
		Msg:  "Invalid session",
	}
}

// func InvalidEmail() Response
func InvalidEmail() Response {
	return Response{
		Code: InvalidEmailCode,
		Msg:  "Invalid email",
	}
}

// func UserNotFound() Response
func UserNotFound() Response {
	return Response{
		Code: UserNotFoundCode,
		Msg:  "User not found",
	}
}

// func InternalServerError() Response
func InternalServerError() Response {
	return Response{
		Code: InternalServerErrorCode,
		Msg:  "Internal server error",
	}
}

// func ResponseSuccess()
func ResponseSuccess(c *gin.Context, status int, event string, userid uint, response Response) {
	c.JSON(status, gin.H{
		"event": event,
		"code":  response.Code,
		"id":    userid,
		"msg":   response.Msg,
	})
}

// func ResponseError()
func ResponseError(c *gin.Context, status int, response Response) {
	c.JSON(status, gin.H{
		"code": response.Code,
		"msg":  response.Msg,
	})
}
