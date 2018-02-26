// Code generated by protoc-gen-gomeet-service. DO NOT EDIT.
// source: pb/api-gateway.proto
package functest

import (
	"fmt"
)

func TestDeleteProfile(config FunctionalTestConfig) (failures []TestFailure) {
	client, ctx, err := grpcClient(config)
	if err != nil {
		failures = append(failures, TestFailure{Procedure: "DeleteProfile", Message: fmt.Sprintf("gRPC client initialization error (%v)", err)})
		return failures
	}
	defer client.Close()

	var testCaseResults []*TestCaseResult
	reqs, extras, err := testGetDeleteProfileRequest(config)
	if err != nil {
		failures = append(failures, TestFailure{Procedure: "DeleteProfile", Message: fmt.Sprintf("HTTP testGetDeleteProfileRequest error (%v)", err)})
		return failures
	}

	for _, req := range reqs {
		res, err := client.GetGRPCClient().DeleteProfile(ctx, req)
		testCaseResults = append(testCaseResults, &TestCaseResult{req, res, err})
	}

	return testDeleteProfileResponse(config, FUNCTEST_GRPC, testCaseResults, extras)
}