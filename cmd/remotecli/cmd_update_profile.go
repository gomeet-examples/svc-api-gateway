package remotecli

import (
	"errors"
	"fmt"
	"strings"

	pb "github.com/gomeet-examples/svc-api-gateway/pb"
)

func (c *remoteCli) cmdUpdateProfile(args []string) (string, error) {
	if len(args) < 8 {
		return "", errors.New("Bad arguments : update_profile <uuid [string]> <gender [UNKNOW|MALE|FEMALE]> <email [string]> <name [string]> <birthday [string]> <created_at [string]> <updated_at [string]> <deleted_at [string]>")
	}

	// request message
	var req *pb.ProfileInfo

	// decl req for no nil panic
	req = &pb.ProfileInfo{}

	// cast args[0] in req.Uuid - type TYPE_STRING to go type string
	req.Uuid = args[0]

	// cast args[1] in req.Gender - type TYPE_ENUM to go type *grpc.Genders
	reqGender, ok := pb.Genders_value[strings.ToUpper(args[1])]
	if !ok {
		return "", fmt.Errorf("Bad arguments : unknown gender \"%s\"", args[1])
	}
	req.Gender = pb.Genders(reqGender)

	// cast args[2] in req.Email - type TYPE_STRING to go type string
	req.Email = args[2]

	// cast args[3] in req.Name - type TYPE_STRING to go type string
	req.Name = args[3]

	// cast args[4] in req.Birthday - type TYPE_STRING to go type string
	req.Birthday = args[4]

	// cast args[5] in req.CreatedAt - type TYPE_STRING to go type string
	req.CreatedAt = args[5]

	// cast args[6] in req.UpdatedAt - type TYPE_STRING to go type string
	req.UpdatedAt = args[6]

	// cast args[7] in req.DeletedAt - type TYPE_STRING to go type string
	req.DeletedAt = args[7]

	// message validation - github.com/mwitkow/go-proto-validators
	if reqValidator, ok := interface{}(*req).(interface {
		Validate() error
	}); ok {
		if err := reqValidator.Validate(); err != nil {
			return "", err
		}
	}

	// sending message to server
	r, err := c.c.UpdateProfile(c.ctx, req)
	if err != nil {
		return "", fmt.Errorf("UpdateProfile service call fail - %v", err)
	}

	return fmt.Sprintf("UpdateProfile: %v", r), nil
}
