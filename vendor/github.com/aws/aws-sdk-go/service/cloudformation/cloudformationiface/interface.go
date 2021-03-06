// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

// Package cloudformationiface provides an interface to enable mocking the AWS CloudFormation service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package cloudformationiface

import (
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/cloudformation"
)

// CloudFormationAPI provides an interface to enable mocking the
// cloudformation.CloudFormation service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // AWS CloudFormation.
//    func myFunc(svc cloudformationiface.CloudFormationAPI) bool {
//        // Make svc.CancelUpdateStack request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := cloudformation.New(sess)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockCloudFormationClient struct {
//        cloudformationiface.CloudFormationAPI
//    }
//    func (m *mockCloudFormationClient) CancelUpdateStack(input *cloudformation.CancelUpdateStackInput) (*cloudformation.CancelUpdateStackOutput, error) {
//        // mock response/functionality
//    }
//
//    TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockCloudFormationClient{}
//
//        myfunc(mockSvc)
//
//        // Verify myFunc's functionality
//    }
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters. Its suggested to use the pattern above for testing, or using
// tooling to generate mocks to satisfy the interfaces.
type CloudFormationAPI interface {
	CancelUpdateStackRequest(*cloudformation.CancelUpdateStackInput) (*request.Request, *cloudformation.CancelUpdateStackOutput)

	CancelUpdateStack(*cloudformation.CancelUpdateStackInput) (*cloudformation.CancelUpdateStackOutput, error)

	ContinueUpdateRollbackRequest(*cloudformation.ContinueUpdateRollbackInput) (*request.Request, *cloudformation.ContinueUpdateRollbackOutput)

	ContinueUpdateRollback(*cloudformation.ContinueUpdateRollbackInput) (*cloudformation.ContinueUpdateRollbackOutput, error)

	CreateChangeSetRequest(*cloudformation.CreateChangeSetInput) (*request.Request, *cloudformation.CreateChangeSetOutput)

	CreateChangeSet(*cloudformation.CreateChangeSetInput) (*cloudformation.CreateChangeSetOutput, error)

	CreateStackRequest(*cloudformation.CreateStackInput) (*request.Request, *cloudformation.CreateStackOutput)

	CreateStack(*cloudformation.CreateStackInput) (*cloudformation.CreateStackOutput, error)

	DeleteChangeSetRequest(*cloudformation.DeleteChangeSetInput) (*request.Request, *cloudformation.DeleteChangeSetOutput)

	DeleteChangeSet(*cloudformation.DeleteChangeSetInput) (*cloudformation.DeleteChangeSetOutput, error)

	DeleteStackRequest(*cloudformation.DeleteStackInput) (*request.Request, *cloudformation.DeleteStackOutput)

	DeleteStack(*cloudformation.DeleteStackInput) (*cloudformation.DeleteStackOutput, error)

	DescribeAccountLimitsRequest(*cloudformation.DescribeAccountLimitsInput) (*request.Request, *cloudformation.DescribeAccountLimitsOutput)

	DescribeAccountLimits(*cloudformation.DescribeAccountLimitsInput) (*cloudformation.DescribeAccountLimitsOutput, error)

	DescribeChangeSetRequest(*cloudformation.DescribeChangeSetInput) (*request.Request, *cloudformation.DescribeChangeSetOutput)

	DescribeChangeSet(*cloudformation.DescribeChangeSetInput) (*cloudformation.DescribeChangeSetOutput, error)

	DescribeStackEventsRequest(*cloudformation.DescribeStackEventsInput) (*request.Request, *cloudformation.DescribeStackEventsOutput)

	DescribeStackEvents(*cloudformation.DescribeStackEventsInput) (*cloudformation.DescribeStackEventsOutput, error)

	DescribeStackEventsPages(*cloudformation.DescribeStackEventsInput, func(*cloudformation.DescribeStackEventsOutput, bool) bool) error

	DescribeStackResourceRequest(*cloudformation.DescribeStackResourceInput) (*request.Request, *cloudformation.DescribeStackResourceOutput)

	DescribeStackResource(*cloudformation.DescribeStackResourceInput) (*cloudformation.DescribeStackResourceOutput, error)

	DescribeStackResourcesRequest(*cloudformation.DescribeStackResourcesInput) (*request.Request, *cloudformation.DescribeStackResourcesOutput)

	DescribeStackResources(*cloudformation.DescribeStackResourcesInput) (*cloudformation.DescribeStackResourcesOutput, error)

	DescribeStacksRequest(*cloudformation.DescribeStacksInput) (*request.Request, *cloudformation.DescribeStacksOutput)

	DescribeStacks(*cloudformation.DescribeStacksInput) (*cloudformation.DescribeStacksOutput, error)

	DescribeStacksPages(*cloudformation.DescribeStacksInput, func(*cloudformation.DescribeStacksOutput, bool) bool) error

	EstimateTemplateCostRequest(*cloudformation.EstimateTemplateCostInput) (*request.Request, *cloudformation.EstimateTemplateCostOutput)

	EstimateTemplateCost(*cloudformation.EstimateTemplateCostInput) (*cloudformation.EstimateTemplateCostOutput, error)

	ExecuteChangeSetRequest(*cloudformation.ExecuteChangeSetInput) (*request.Request, *cloudformation.ExecuteChangeSetOutput)

	ExecuteChangeSet(*cloudformation.ExecuteChangeSetInput) (*cloudformation.ExecuteChangeSetOutput, error)

	GetStackPolicyRequest(*cloudformation.GetStackPolicyInput) (*request.Request, *cloudformation.GetStackPolicyOutput)

	GetStackPolicy(*cloudformation.GetStackPolicyInput) (*cloudformation.GetStackPolicyOutput, error)

	GetTemplateRequest(*cloudformation.GetTemplateInput) (*request.Request, *cloudformation.GetTemplateOutput)

	GetTemplate(*cloudformation.GetTemplateInput) (*cloudformation.GetTemplateOutput, error)

	GetTemplateSummaryRequest(*cloudformation.GetTemplateSummaryInput) (*request.Request, *cloudformation.GetTemplateSummaryOutput)

	GetTemplateSummary(*cloudformation.GetTemplateSummaryInput) (*cloudformation.GetTemplateSummaryOutput, error)

	ListChangeSetsRequest(*cloudformation.ListChangeSetsInput) (*request.Request, *cloudformation.ListChangeSetsOutput)

	ListChangeSets(*cloudformation.ListChangeSetsInput) (*cloudformation.ListChangeSetsOutput, error)

	ListExportsRequest(*cloudformation.ListExportsInput) (*request.Request, *cloudformation.ListExportsOutput)

	ListExports(*cloudformation.ListExportsInput) (*cloudformation.ListExportsOutput, error)

	ListImportsRequest(*cloudformation.ListImportsInput) (*request.Request, *cloudformation.ListImportsOutput)

	ListImports(*cloudformation.ListImportsInput) (*cloudformation.ListImportsOutput, error)

	ListStackResourcesRequest(*cloudformation.ListStackResourcesInput) (*request.Request, *cloudformation.ListStackResourcesOutput)

	ListStackResources(*cloudformation.ListStackResourcesInput) (*cloudformation.ListStackResourcesOutput, error)

	ListStackResourcesPages(*cloudformation.ListStackResourcesInput, func(*cloudformation.ListStackResourcesOutput, bool) bool) error

	ListStacksRequest(*cloudformation.ListStacksInput) (*request.Request, *cloudformation.ListStacksOutput)

	ListStacks(*cloudformation.ListStacksInput) (*cloudformation.ListStacksOutput, error)

	ListStacksPages(*cloudformation.ListStacksInput, func(*cloudformation.ListStacksOutput, bool) bool) error

	SetStackPolicyRequest(*cloudformation.SetStackPolicyInput) (*request.Request, *cloudformation.SetStackPolicyOutput)

	SetStackPolicy(*cloudformation.SetStackPolicyInput) (*cloudformation.SetStackPolicyOutput, error)

	SignalResourceRequest(*cloudformation.SignalResourceInput) (*request.Request, *cloudformation.SignalResourceOutput)

	SignalResource(*cloudformation.SignalResourceInput) (*cloudformation.SignalResourceOutput, error)

	UpdateStackRequest(*cloudformation.UpdateStackInput) (*request.Request, *cloudformation.UpdateStackOutput)

	UpdateStack(*cloudformation.UpdateStackInput) (*cloudformation.UpdateStackOutput, error)

	ValidateTemplateRequest(*cloudformation.ValidateTemplateInput) (*request.Request, *cloudformation.ValidateTemplateOutput)

	ValidateTemplate(*cloudformation.ValidateTemplateInput) (*cloudformation.ValidateTemplateOutput, error)

	WaitUntilStackCreateComplete(*cloudformation.DescribeStacksInput) error

	WaitUntilStackDeleteComplete(*cloudformation.DescribeStacksInput) error

	WaitUntilStackExists(*cloudformation.DescribeStacksInput) error

	WaitUntilStackUpdateComplete(*cloudformation.DescribeStacksInput) error
}

var _ CloudFormationAPI = (*cloudformation.CloudFormation)(nil)
