// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: pb/api-gateway.proto

/*
Package pb is a generated protocol buffer package.

It is generated from these files:
	pb/api-gateway.proto

It has these top-level messages:
	EmptyMessage
	VersionResponse
	ServiceStatus
	ServicesStatusList
	EchoRequest
	EchoResponse
	ProfileInfo
	ProfileRequest
	ProfileResponse
	ProfileCreationRequest
	ProfileListRequest
	ProfileList
*/
package pb

import faker "github.com/dmgk/faker"
import locales "github.com/dmgk/faker/locales"
import rand "math/rand"
import time "time"
import uuid "github.com/google/uuid"
import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"
import _ "github.com/mwitkow/go-proto-validators"
import _ "github.com/gomeet/go-proto-gomeetfaker"
import _ "github.com/gogo/protobuf/gogoproto"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func GomeetFakerRand() *rand.Rand {
	seed := time.Now().UnixNano()
	return rand.New(rand.NewSource(seed))
}
func init() {
	GomeetFakerSetLocale("en")
}

func GomeetFakerSetLocale(l string) {
	switch l {
	case "pt-br":
		faker.Locale = locales.Pt_BR
	case "fa":
		faker.Locale = locales.Fa
	case "de-ch":
		faker.Locale = locales.De_CH
	case "en-bork":
		faker.Locale = locales.En_BORK
	case "nl":
		faker.Locale = locales.Nl
	case "pl":
		faker.Locale = locales.Pl
	case "en-ca":
		faker.Locale = locales.En_CA
	case "es":
		faker.Locale = locales.Es
	case "ja":
		faker.Locale = locales.Ja
	case "zh-tw":
		faker.Locale = locales.Zh_TW
	case "en-gb":
		faker.Locale = locales.En_GB
	case "en-us":
		faker.Locale = locales.En_US
	case "en-nep":
		faker.Locale = locales.En_NEP
	case "fr":
		faker.Locale = locales.Fr
	case "zh-cn":
		faker.Locale = locales.Zh_CN
	case "en-au":
		faker.Locale = locales.En_AU
	case "en-ind":
		faker.Locale = locales.En_IND
	case "ko":
		faker.Locale = locales.Ko
	case "nb-no":
		faker.Locale = locales.Nb_NO
	case "ru":
		faker.Locale = locales.Ru
	case "de":
		faker.Locale = locales.De
	case "it":
		faker.Locale = locales.It
	case "en":
		faker.Locale = locales.En
	case "sk":
		faker.Locale = locales.Sk
	case "vi":
		faker.Locale = locales.Vi
	case "en-au-ocker":
		faker.Locale = locales.En_AU_OCKER
	case "de-at":
		faker.Locale = locales.De_AT
	case "sv":
		faker.Locale = locales.Sv
	default:
		faker.Locale = locales.En
	}
}
func NewEmptyMessageGomeetFaker() *EmptyMessage {
	this := &EmptyMessage{}
	return this
}

func NewVersionResponseGomeetFaker() *VersionResponse {
	this := &VersionResponse{}
	this.Name = faker.App().Name()
	this.Version = faker.App().Version()
	return this
}

func NewServiceStatusGomeetFaker() *ServiceStatus {
	this := &ServiceStatus{}
	this.Name = faker.App().Name()
	this.Version = faker.App().Version()
	// this.Status is a string or bytes without gommetfaker rules so faker.Lorem().Lorem() is used
	this.Status = ServiceStatus_Status([]int32{0, 1}[GomeetFakerRand().Intn(2)])
	// this.EMsg // skipped by skip rules
	return this
}

func NewServicesStatusListGomeetFaker() *ServicesStatusList {
	this := &ServicesStatusList{}
	for i := 0; i < 3; i++ {
		aCurrentServices := NewServiceStatusGomeetFaker()
		this.Services = append(this.Services, aCurrentServices)
	}
	return this
}

func NewEchoRequestGomeetFaker() *EchoRequest {
	this := &EchoRequest{}
	this.Uuid = uuid.New().String()
	this.Content = faker.Lorem().String()
	return this
}

func NewEchoResponseGomeetFaker() *EchoResponse {
	this := &EchoResponse{}
	this.Uuid = uuid.New().String()
	this.Content = faker.Lorem().String()
	return this
}

func NewProfileInfoGomeetFaker() *ProfileInfo {
	this := &ProfileInfo{}
	this.Uuid = uuid.New().String()
	this.Gender = Genders([]int32{1, 2}[GomeetFakerRand().Intn(2)])
	this.Email = faker.Internet().Email()
	this.Name = faker.Internet().UserName()
	aBirthdayTime := faker.Time().Birthday(17, 99)
	this.Birthday = aBirthdayTime.Format("2006-01-02")
	this.CreatedAt = time.Now().Format("2006-01-02T15:04:05Z07:00")
	this.UpdatedAt = time.Now().Format("2006-01-02T15:04:05Z07:00")
	this.DeletedAt = time.Now().Format("2006-01-02T15:04:05Z07:00")
	return this
}

func NewProfileRequestGomeetFaker() *ProfileRequest {
	this := &ProfileRequest{}
	this.Uuid = uuid.New().String()
	return this
}

func NewProfileResponseGomeetFaker() *ProfileResponse {
	this := &ProfileResponse{}
	this.Ok = true
	this.Info = NewProfileInfoGomeetFaker()
	return this
}

func NewProfileCreationRequestGomeetFaker() *ProfileCreationRequest {
	this := &ProfileCreationRequest{}
	this.Gender = Genders([]int32{1, 2}[GomeetFakerRand().Intn(2)])
	this.Email = faker.Internet().Email()
	this.Name = faker.Internet().UserName()
	aBirthdayTime := faker.Time().Birthday(17, 99)
	this.Birthday = aBirthdayTime.Format("2006-01-02")
	return this
}

func NewProfileListRequestGomeetFaker() *ProfileListRequest {
	this := &ProfileListRequest{}
	this.PageNumber = uint32(1)
	this.PageSize = uint32(200)
	this.Order = "created_at asc"
	this.ExcludeSoftDeleted = true
	this.SoftDeletedOnly = false
	this.Gender = Genders([]int32{1, 2}[GomeetFakerRand().Intn(2)])
	return this
}

func NewProfileListGomeetFaker() *ProfileList {
	this := &ProfileList{}
	this.ResultSetSize = uint32(5)
	this.HasMore = false
	for i := 0; i < 5; i++ {
		aCurrentProfiles := NewProfileInfoGomeetFaker()
		this.Profiles = append(this.Profiles, aCurrentProfiles)
	}
	return this
}
