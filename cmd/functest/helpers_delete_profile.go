package functest

import (
	pb "github.com/gomeet-examples/svc-api-gateway/pb"
)

func testGetDeleteProfileRequest(
	config FunctionalTestConfig,
) (reqs []*pb.ProfileRequest, extras map[string]interface{}, err error) {
	// TODO

	// return an array of pb.ProfileRequest struct pointers,
	// each of them will be passed as an argument to the grpc DeleteProfile method

	reqs = append(reqs, &pb.ProfileRequest{})
	return reqs, extras, err
}

func testDeleteProfileResponse(
	config FunctionalTestConfig,
	testsType string,
	testCaseResults []*TestCaseResult,
	extras map[string]interface{},
) (failures []TestFailure) {
	// TODO

	// Do something useful functional test with
	// testCaseResults[n].Request, testCaseResults[n].Response and testCaseResults[n].Error
	// then return a array of TestFailure struct
	// testsType value is value of FUNCTEST_HTTP (HTTP) and FUNCTEST_GRPC (GRPC) constants cf. types.go
	for _, tr := range testCaseResults {
		var (
			req *pb.ProfileRequest
			res *pb.ProfileResponse
			err error
			ok  bool
		)
		if tr.Request == nil {
			failures = append(failures, TestFailure{Procedure: "DeleteProfile", Message: "expected request message type pb.ProfileRequest - nil given"})
			continue
		}
		req, ok = tr.Request.(*pb.ProfileRequest)
		if !ok {
			failures = append(failures, TestFailure{Procedure: "DeleteProfile", Message: "expected request message type pb.ProfileRequest - cast fail"})
			continue
		}

		if tr.Response != nil {
			res, ok = tr.Response.(*pb.ProfileResponse)
			if !ok {
				failures = append(failures, TestFailure{Procedure: "DeleteProfile", Message: "expected response message type pb.ProfileRequest - cast fail"})
				continue
			}
		}

		// Do something useful functional test with req, res and err
		err = tr.Error
		if err != nil {
			// if no error are expected do something like this
			// failures = append(failures, TestFailure{Procedure: "DeleteProfile", Message: "no error expected"})
			// continue
		}

		if req != nil && res != nil {
			// for example :
			// if res.GetId() != req.GetId() {
			//     failureMsg := fmt.Sprintf("expected ID \"%s\" but got \"%s\" for request: %v", req.GetId(), res.GetId(), req)
			//     failures = append(failures, TestFailure{Procedure: "DeleteProfile", Message: failureMsg})
			// }
		}
	}

	return failures
}
