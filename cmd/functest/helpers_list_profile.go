package functest

import (
	pb "github.com/gomeet-examples/svc-api-gateway/pb"
)

func testGetListProfileRequest(
	config FunctionalTestConfig,
) (reqs []*pb.ProfileListRequest, extras map[string]interface{}, err error) {
	// return an array of pb.ProfileListRequest struct pointers,
	// each of them will be passed as an argument to the grpc ListProfile method

	reqs = append(reqs, &pb.ProfileListRequest{})
	return reqs, extras, err
}

func testListProfileResponse(
	config FunctionalTestConfig,
	testsType string,
	testCaseResults []*TestCaseResult,
	extras map[string]interface{},
) (failures []TestFailure) {
	// Do something useful functional test with
	// testCaseResults[n].Request, testCaseResults[n].Response and testCaseResults[n].Error
	// then return a array of TestFailure struct
	// testsType value is value of FUNCTEST_HTTP (HTTP) and FUNCTEST_GRPC (GRPC) constants cf. types.go
	for _, tr := range testCaseResults {
		var (
			req *pb.ProfileListRequest
			res *pb.ProfileList
			err error
			ok  bool
		)
		if tr.Request == nil {
			failures = append(failures, TestFailure{Procedure: "ListProfile", Message: "expected request message type pb.ProfileListRequest - nil given"})
			continue
		}
		req, ok = tr.Request.(*pb.ProfileListRequest)
		if !ok {
			failures = append(failures, TestFailure{Procedure: "ListProfile", Message: "expected request message type pb.ProfileListRequest - cast fail"})
			continue
		}

		if tr.Response != nil {
			res, ok = tr.Response.(*pb.ProfileList)
			if !ok {
				failures = append(failures, TestFailure{Procedure: "ListProfile", Message: "expected response message type pb.ProfileListRequest - cast fail"})
				continue
			}
		}

		// Do something useful functional test with req, res and err
		err = tr.Error
		if err != nil {
			// if no error are expected do something like this
			// failures = append(failures, TestFailure{Procedure: "ListProfile", Message: "no error expected"})
			// continue
		}

		if req != nil && res != nil {
			// for example :
			// if res.GetId() != req.GetId() {
			//     failureMsg := fmt.Sprintf("expected ID \"%s\" but got \"%s\" for request: %v", req.GetId(), res.GetId(), req)
			//     failures = append(failures, TestFailure{Procedure: "ListProfile", Message: failureMsg})
			// }
		}
	}

	return failures
}
