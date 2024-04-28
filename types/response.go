package types

type Response struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
}

type FullResponse struct {
	Event string `json:"event"`
	Code  string `json:"code"`
	Msg   string `json:"msg"`
}

type FullResponseWithData struct {
	FullResponse
	Data interface{} `json:"data"`
}

const (
	// success 200
	SuccessCode      = "20000"
	LoginSuccessCode = "20001"
	ForgotPassCode   = "20002"
	ResetPassCode    = "20003"
	MatchSucessCode  = "20004"

	// 201
	UserCreatedCode         = "20101"
	EmailVerifiedCode       = "20102"
	MFAVerifiedCode         = "20103"
	PostCreatedCode         = "20104"
	NotificationCreatedCode = "20105"
	BidCreatedCode          = "20106"

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
	RecordNotFoundCode = "40401"
	UserNotFoundCode   = "40402"

	// 409
	VerificationCodeExistsCode = "40901"

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

// func MatchSuccess() Response
func MatchSuccess() Response {
	return Response{
		Code: MatchSucessCode,
		Msg:  "Match success",
	}
}

// func UserCreated() Response
func UserCreated() Response {
	return Response{
		Code: UserCreatedCode,
		Msg:  "User created, please verify your email",
	}
}

// func PostCreated() Response
func PostCreated() Response {
	return Response{
		Code: PostCreatedCode,
		Msg:  "Post created successfully",
	}
}

// func BidCreated() Response
func BidCreated() Response {
	return Response{
		Code: BidCreatedCode,
		Msg:  "Bid placed successfully",
	}
}

// func NotificationCreated() Response
func NotificationCreated() Response {
	return Response{
		Code: NotificationCreatedCode,
		Msg:  "Notification created successfully",
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

// func RecordNotFound() Response
func RecordNotFound() Response {
	return Response{
		Code: RecordNotFoundCode,
		Msg:  "Record not found",
	}
}

// func UserNotFound() Response
func UserNotFound() Response {
	return Response{
		Code: UserNotFoundCode,
		Msg:  "User not found",
	}
}

// func VerificationCodeExists Response
func VerificationCodeExists() Response {
	return Response{
		Code: VerificationCodeExistsCode,
		Msg:  "Verification code exists",
	}
}

// func TooManyRequest() Respons
func TooManyRequest() Response {
	return Response{
		Code: TooManyRequestsCode,
		Msg:  "Too many request",
	}
}

// func InternalServerError() Response
func InternalServerError() Response {
	return Response{
		Code: InternalServerErrorCode,
		Msg:  "Internal server error",
	}
}
