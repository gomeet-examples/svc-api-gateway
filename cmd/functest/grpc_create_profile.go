// Code generated by protoc-gen-gomeet-service. DO NOT EDIT.
// source: pb/api-gateway.proto
package functest

import (
	"fmt"
)

func TestCreateProfile(config FunctionalTestConfig) (failures []TestFailure) {
	client, ctx, err := grpcClient(config)
	if err != nil {
		failures = append(failures, TestFailure{Procedure: "CreateProfile", Message: fmt.Sprintf("gRPC client initialization error (%v)", err)})
		return failures
	}
	defer client.Close()

	var testCaseResults []*TestCaseResult
	reqs, extras, err := testGetCreateProfileRequest(config)
	if err != nil {
		failures = append(failures, TestFailure{Procedure: "CreateProfile", Message: fmt.Sprintf("HTTP testGetCreateProfileRequest error (%v)", err)})
		return failures
	}

	for _, req := range reqs {
		res, err := client.GetGRPCClient().CreateProfile(ctx, req)
		testCaseResults = append(testCaseResults, &TestCaseResult{req, res, err})
	}

	return testCreateProfileResponse(config, FUNCTEST_GRPC, testCaseResults, extras)
}
