package main

import "time"

type Customer struct {
	Base                       `bson:",inline" json:",inline"`
	Email                      string     `bson:"email,omitempty" json:"email,omitempty"`
	Password                   string     `bson:"password,omitempty" json:"password,omitempty"`
	Name                       string     `bson:"name,omitempty" json:"name,omitempty"`
	PhoneVerified              bool       `bson:"phone_verified" json:"phone_verified"`
	IsEnabled                  bool       `bson:"is_enabled" json:"is_enabled"`
	EmailVerified              bool       `bson:"email_verified" json:"email_verified"`
	IsOfficialAccount          bool       `bson:"is_official_account" json:"is_official_account"`
	PhoneNumber                string     `bson:"phone_number,omitempty" json:"phone_number,omitempty"`
	PhoneCode                  string     `bson:"phone_code,omitempty" json:"phone_code,omitempty"`
	ProfilePicture             string     `bson:"profile_picture,omitempty" json:"profile_picture,omitempty"`
	Addr                       string     `bson:"addr,omitempty" json:"addr,omitempty"`
	City                       string     `bson:"city,omitempty" json:"city,omitempty"`
	District                   string     `bson:"district,omitempty" json:"district,omitempty"`
	Ward                       string     `bson:"ward,omitempty" json:"ward,omitempty"`
	Description                string     `bson:"description,omitempty" json:"description,omitempty"`
	InviterId                  string     `bson:"inviter_id,omitempty" json:"inviter_id,omitempty"`
	IsSocialAccount            bool       `bson:"is_social_account" json:"-"`
	GoogleID                   string     `bson:"google_id,omitempty"`
	FacebookID                 string     `bson:"facebook_id,omitempty"`
	AppleID                    string     `bson:"apple_id,omitempty"`
	VerificationCode           string     `bson:"verification_code,omitempty"`
	VerificationCodeExpiredAt  *time.Time `bson:"verification_code_expired_at,omitempty"`
	ResetPasswordCode          string     `bson:"reset_password_code,omitempty"`
	ResetPasswordCodeExpiredAt *time.Time `bson:"reset_password_code_expired_at,omitempty"`
	ResetPasswordOTP           string     `bson:"reset_password_otp"`
	NewEmail                   string     `bson:"new_email,omitempty"`
	NewPhoneNumber             string     `bson:"new_phone_number,omitempty"`

	LastLogin     *int64 `bson:"last_login,omitempty"`
	PartnershipId string `bson:"partnership_id,omitempty" json:"partnership_id,omitempty"`
}
