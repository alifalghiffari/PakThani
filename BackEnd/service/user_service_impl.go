package service

import (
	"context"
	"crypto/rand"
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/big"

	// "log"
	// "math/rand"
	"net/http"
	"net/smtp"
	"net/url"
	"strings"
	"time"
	"unicode"

	"project-workshop/go-api-ecom/helper"
	"project-workshop/go-api-ecom/model/domain"
	"project-workshop/go-api-ecom/model/web"
	"project-workshop/go-api-ecom/repository"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type UserServiceImpl struct {
	UserRepository repository.UserRepository
	DB             *sql.DB
	Validate       *validator.Validate
	GoogleConfig   *oauth2.Config
	Error          error
	SMTPAuth       smtp.Auth
	SMTPHost       string
	SMTPPort       string
}

func NewUserService(userRepository repository.UserRepository, DB *sql.DB, validate *validator.Validate, googleConfig *oauth2.Config, smtpAuth smtp.Auth, smtpHost, smtpPort string) UserService {
	return &UserServiceImpl{
		UserRepository: userRepository,
		DB:             DB,
		Validate:       validate,
		GoogleConfig:   googleConfig,
		SMTPAuth:       smtpAuth,
		SMTPHost:       smtpHost,
		SMTPPort:       smtpPort,
	}
}

type Claims struct {
	Email    string
	UserID   int
	Role     bool
	jwt.RegisteredClaims
}

type Config struct {
	GoogleLoginConfig oauth2.Config
}

func (service *UserServiceImpl) Register(ctx context.Context, request web.UserCreateRequest) web.UserResponse {
	// Validasi request
	err := service.Validate.Struct(request)
	if err != nil {
		fmt.Println("Validation error:", err)
		return web.UserResponse{Error: "Validation error"}
	}

	// Validasi password
	if !IsValidPassword(request.Password) {
		fmt.Println("Password does not meet security criteria")
		return web.UserResponse{Error: "Password must be at least 8 characters long and include an uppercase letter, a lowercase letter, a number, and a special character"}
	}

	// Mulai transaksi database
	tx, err := service.DB.Begin()
	if err != nil {
		fmt.Println("DB Begin error:", err)
		return web.UserResponse{Error: "Database error"}
	}
	defer helper.CommitOrRollback(tx)

	// Hash password
	hashedPassword, err := HashPassword(request.Password)
	if err != nil {
		fmt.Println("HashPassword error:", err)
		return web.UserResponse{Error: "Internal server error"}
	}

	// Generate OTP untuk verifikasi email
	otp := GenerateOTP()

	// Buat objek User
	user := domain.User{
		Username:          request.Username,
		Password:          hashedPassword,
		IsVerified:        false,
		VerificationToken: otp,
		Email:             request.Email,
		Role:              request.Role,
		NoTelepon:         request.NoTelepon,
	}

	// Simpan user ke database
	user = service.UserRepository.Register(ctx, tx, user)

	// Kirim email verifikasi
	err = SendOTPEmailVerification(user.Email, otp, service.SMTPHost, service.SMTPPort, service.SMTPAuth)
	if err != nil {
		fmt.Println("SendOTPEmailVerification error:", err)
		return web.UserResponse{Error: "Failed to send verification email"}
	}

	// Kembalikan respons sukses
	return helper.ToUserResponse(user)
}

func (service *UserServiceImpl) VerifyEmail(ctx context.Context, request web.UserVerifyRequest) web.UserResponse {
	// Validasi request
	err := service.Validate.Struct(request)
	if err != nil {
		return web.UserResponse{Error: "Invalid request"}
	}

	// Mulai transaksi database
	tx, err := service.DB.Begin()
	if err != nil {
		return web.UserResponse{Error: "Database error"}
	}
	defer helper.CommitOrRollback(tx)

	// Verifikasi token
	user, err := service.UserRepository.VerifyEmail(ctx, tx, request.VerificationToken)
	if err != nil {
		return web.UserResponse{Error: "Invalid verification token"}
	}

	// Update status verifikasi
	user, err = service.UserRepository.UpdateVerificationStatus(ctx, tx, user.Email, request.VerificationToken)
	if err != nil {
		return web.UserResponse{Error: "Failed to update verification status"}
	}

	// Kembalikan respons sukses
	return helper.ToUserResponse(user)
}

func (service *UserServiceImpl) Login(ctx context.Context, request web.UserLoginRequest) (web.UserResponse, error) {
	err := service.Validate.Struct(request)
	if err != nil {
		return web.UserResponse{}, err
	}

	tx, err := service.DB.Begin()
	if err != nil {
		return web.UserResponse{}, err
	}
	defer helper.CommitOrRollback(tx)

	user, err := service.UserRepository.FindByEmail(ctx, tx, request.Email)
	if err != nil || user.IsVerified != true {
		return web.UserResponse{}, fmt.Errorf("user not found")
	}

	err = ComparePassword(user.Password, request.Password)
	if err != nil {
		return web.UserResponse{}, err
	}

	token, err := GenerateToken(user.Email, user.Role, user.Id, "yourSecretKey")
	if err != nil {
		return web.UserResponse{}, err
	}

	userResponse := helper.ToUserResponse(user)
	userResponse.Token = token

	return userResponse, nil
}

func (service *UserServiceImpl) GetGoogleAuthURL() string {
	return service.GoogleConfig.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
}

func (service *UserServiceImpl) LoginGoogle(ctx context.Context, code string) (web.UserResponse, error) {
	tokenURL := "https://oauth2.googleapis.com/token"
	data := url.Values{}
	data.Set("code", code)
	data.Set("client_id", "1083611601525-sm55jlhpioajjvbp811ikq5bbn2m04vh.apps.googleusercontent.com")
	data.Set("client_secret", "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")
	data.Set("redirect_uri", "postmessage")
	data.Set("grant_type", "authorization_code")

	req, err := http.NewRequest("POST", tokenURL, strings.NewReader(data.Encode()))
	if err != nil {
		return web.UserResponse{}, fmt.Errorf("failed to create request: %v", err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return web.UserResponse{}, fmt.Errorf("failed to perform request: %v", err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return web.UserResponse{}, fmt.Errorf("failed to read response body: %v", err)
	}

	var tokenResp struct {
		AccessToken string `json:"access_token"`
	}
	if err := json.Unmarshal(body, &tokenResp); err != nil {
		return web.UserResponse{}, fmt.Errorf("failed to unmarshal token response: %v", err)
	}

	resp, err := http.Get("https://www.googleapis.com/oauth2/v3/userinfo?access_token=" + tokenResp.AccessToken)
	if err != nil {
		return web.UserResponse{}, fmt.Errorf("failed to get user info: %v", err)
	}
	defer resp.Body.Close()

	var userInfo struct {
		Email string `json:"email"`
		Name  string `json:"name"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&userInfo); err != nil {
		return web.UserResponse{}, fmt.Errorf("failed to decode user info: %v", err)
	}

	tx, err := service.DB.Begin()
	if err != nil {
		return web.UserResponse{}, fmt.Errorf("failed to begin transaction: %v", err)
	}
	defer helper.CommitOrRollback(tx)

	user, err := service.UserRepository.FindByEmailGoogle(ctx, tx, userInfo.Email)
	if err != nil {
		if err.Error() == "user not found" {
			user = domain.User{
				Username:  userInfo.Name,
				Email:     userInfo.Email,
				Password:  "-",
				Role:      true,
				NoTelepon: "0",
			}
			user = service.UserRepository.Register(ctx, tx, user)
		} else {
			return web.UserResponse{}, fmt.Errorf("failed to find user: %v", err)
		}
	}

	jwtToken, err := GenerateToken(user.Email, user.Role, user.Id, "yourSecretKey")
	if err != nil {
		return web.UserResponse{}, fmt.Errorf("failed to generate token: %v", err)
	}

	userResponse := helper.ToUserResponse(user)
	userResponse.Token = jwtToken

	return userResponse, nil
}

func GoogleConfig() *oauth2.Config {
	return &oauth2.Config{
		ClientID:     "1083611601525-sm55jlhpioajjvbp811ikq5bbn2m04vh.apps.googleusercontent.com",
		ClientSecret: "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",
		RedirectURL:  "http://localhost:3000/auth/callback",
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile",
		},
		Endpoint: google.Endpoint,
	}
}

func (service *UserServiceImpl) SendResetPassword(ctx context.Context, request web.ForgotPasswordRequest) web.UserResponse {
	err := service.Validate.Struct(request)
	if err != nil {
		return web.UserResponse{Error: "Invalid request"}
	}

	tx, err := service.DB.Begin()
	if err != nil {
		return web.UserResponse{Error: "Database error"}
	}
	defer helper.CommitOrRollback(tx)

	user, err := service.UserRepository.FindByEmail(ctx, tx, request.Email)
	if err != nil {
		return web.UserResponse{Error: "User not found"}
	}

	otp := GenerateOTP()

	resetPassword := domain.ResetPassword{
		UserId:     user.Id,
		Token:      otp,
		Expired_at: time.Now().Add(2 * time.Minute),
	}

	resetPassword, err = service.UserRepository.InsertResetPassword(ctx, tx, resetPassword)
	if err != nil {
		return web.UserResponse{Error: "Failed to insert reset password"}
	}

	err = SendOTPResetPassword(user.Email, otp, service.SMTPHost, service.SMTPPort, service.SMTPAuth)
	if err != nil {
		return web.UserResponse{Error: "Failed to send OTP"}
	}

	return helper.ToUserResponse(user)
}

func (service *UserServiceImpl) VerifyResetPassword(ctx context.Context, request web.ResetPasswordRequest) web.UserResponse {
    err := service.Validate.Struct(request)
    if err != nil {
        return web.UserResponse{Error: "Invalid request"}
    }

    tx, err := service.DB.Begin()
    if err != nil {
        return web.UserResponse{Error: "Database error"}
    }
    defer helper.CommitOrRollback(tx)

    resetPassword, err := service.UserRepository.FindByToken(ctx, tx, request.Token)
    if err != nil {
        // Log the error for debugging purposes
        fmt.Printf("Error finding reset password by token: %v\n", err)
        return web.UserResponse{Error: "Invalid token"}
    }

    if time.Now().After(resetPassword.Expired_at) {
        return web.UserResponse{Error: "Token expired"}
    }

    user, err := service.UserRepository.FindById(ctx, tx, resetPassword.UserId)
    if err != nil {
        // Log the error for debugging purposes
        fmt.Printf("Error finding user by ID: %v\n", err)
        return web.UserResponse{Error: "User not found"}
    }

    // Here we assume GenerateToken takes only user ID and role for simplicity
    token, err := GenerateToken(user.Email, user.Role, user.Id, "yourSecretKey")
    if err != nil {
        // Log the error for debugging purposes
        fmt.Printf("Error generating token: %v\n", err)
        return web.UserResponse{Error: "Failed to generate token"}
    }

    userResponse := helper.ToUserResponse(user)
    userResponse.Token = token

    return userResponse
}

func (service *UserServiceImpl) ResetPassword(ctx context.Context, request web.ChangePasswordRequest, userId int) web.UserResponse {
    err := service.Validate.Struct(request)
    if err != nil {
        return web.UserResponse{Error: "Invalid request"}
    }

    tx, err := service.DB.Begin()
    if err != nil {
        return web.UserResponse{Error: "Database error"}
    }
    defer helper.CommitOrRollback(tx)

    user, err := service.UserRepository.FindById(ctx, tx, userId)
    if err != nil {
        return web.UserResponse{Error: "User not found"}
    }

    if !IsValidPassword(request.Password) {
        return web.UserResponse{Error: "Password must be at least 8 characters long and include an uppercase letter, a lowercase letter, a number, and a special character"}
    }

    hashedPassword, err := HashPassword(request.Password)
    if err != nil {
        return web.UserResponse{Error: "Failed to hash password"}
    }

    user.Password = hashedPassword

    user, err = service.UserRepository.UpdatePassword(ctx, tx, user)
    if err != nil {
        return web.UserResponse{Error: "Failed to update password"}
    }

	user, err = service.UserRepository.DeletedByUserId(ctx, tx, userId)
	if err != nil {
		return web.UserResponse{Error: "Failed to reset password"}
	}

    return helper.ToUserResponse(user)
}

// Utilities function >>>

func GenerateToken(email string, role bool, userId int, secretKey string) (string, error) {
	claims := &Claims{
		Email:  email,
		UserID: userId,
		Role:   role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte("secretKey"))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func IsValidPassword(password string) bool {
	var (
		hasMinLen  = false
		hasUpper   = false
		hasLower   = false
		hasNumber  = false
		hasSpecial = false
	)

	// Panjang minimal 8 karakter
	if len(password) >= 8 {
		hasMinLen = true
	}

	for _, char := range password {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsDigit(char):
			hasNumber = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			hasSpecial = true
		}
	}

	// Setidaknya harus ada satu karakter dari setiap kriteria
	return hasMinLen && hasUpper && hasLower && hasNumber && hasSpecial
}

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func ComparePassword(hashedPassword string, password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return err
	}
	return nil
}

func GenerateOTP() string {
	const number = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	const otpLength = 6

	otp := make([]byte, otpLength)
	for i := range otp {
		randomIndex, err := rand.Int(rand.Reader, big.NewInt(int64(len(number))))
		if err != nil {
			// Handle error
			return ""
		}
		otp[i] = number[randomIndex.Int64()]
	}
	return string(otp)
}

func SendOTPResetPassword(email, otp, host, port string, auth smtp.Auth) error {
	from := ""
	msg := fmt.Sprintf("To: %s\r\nSubject: Reset Password\r\n\r\nYour Code for reset password is: %s\r\n", email, otp)
	return smtp.SendMail(host+":"+port, auth, from, []string{email}, []byte(msg))
}

func SendOTPEmailVerification(email, otp, host, port string, auth smtp.Auth) error {
	from := ""
	msg := fmt.Sprintf("To: %s\r\nSubject: Email Verification\r\n\r\nYour Code for email verification is: %s\r\n", email, otp)
	return smtp.SendMail(host+":"+port, auth, from, []string{email}, []byte(msg))
}
