package akp

import (
	"context"
	"fmt"
	"os"
	"testing"

	hashitype "github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"

	"time"

	"github.com/akuity/api-client-go/pkg/api/gateway/accesscontrol"
	gwoption "github.com/akuity/api-client-go/pkg/api/gateway/option"
	argocdv1 "github.com/akuity/api-client-go/pkg/api/gen/argocd/v1"
	kargov1 "github.com/akuity/api-client-go/pkg/api/gen/kargo/v1"
	orgcv1 "github.com/akuity/api-client-go/pkg/api/gen/organization/v1"
	idv1 "github.com/akuity/api-client-go/pkg/api/gen/types/id/v1"
	httpctx "github.com/akuity/grpc-gateway-client/pkg/http/context"
	"github.com/akuity/terraform-provider-akp/akp/types"
)

var (
	instanceId        string
	createdInstanceId string
	testAkpCli        *AkpCli
)

func getInstanceId() string {
	if instanceId == "" {
		if v := os.Getenv("AKUITY_INSTANCE_ID"); v == "" {
			// Create a new instance for testing
			instanceId = createTestInstance()
		} else {
			instanceId = v
		}
	}

	return instanceId
}

func getTestAkpCli() *AkpCli {
	if testAkpCli != nil {
		return testAkpCli
	}

	ctx := context.Background()

	serverUrl := os.Getenv("AKUITY_SERVER_URL")
	if serverUrl == "" {
		serverUrl = "https://akuity.cloud"
	}

	apiKeyID := os.Getenv("AKUITY_API_KEY_ID")
	apiKeySecret := os.Getenv("AKUITY_API_KEY_SECRET")

	if apiKeyID == "" || apiKeySecret == "" {
		panic("API key credentials are required")
	}

	// Create client following the same logic as the provider
	cred := accesscontrol.NewAPIKeyCredential(apiKeyID, apiKeySecret)
	ctx = httpctx.SetAuthorizationHeader(ctx, cred.Scheme(), cred.Credential())
	gwc := gwoption.NewClient(serverUrl, false)
	orgc := orgcv1.NewOrganizationServiceGatewayClient(gwc)

	// Get Organization ID by name
	res, err := orgc.GetOrganization(ctx, &orgcv1.GetOrganizationRequest{
		Id:     orgName,
		IdType: idv1.Type_NAME,
	})
	if err != nil {
		panic(fmt.Sprintf("Failed to get organization: %v", err))
	}

	orgID := res.Organization.Id

	// Create service clients
	argoc := argocdv1.NewArgoCDServiceGatewayClient(gwc)
	kargoc := kargov1.NewKargoServiceGatewayClient(gwc)

	testAkpCli = &AkpCli{
		Cli:      argoc,
		KargoCli: kargoc,
		Cred:     cred,
		OrgId:    orgID,
		OrgCli:   orgc,
	}
	return testAkpCli
}

func createTestInstance() string {
	akpCli := getTestAkpCli()
	ctx := context.Background()
	ctx = httpctx.SetAuthorizationHeader(ctx, akpCli.Cred.Scheme(), akpCli.Cred.Credential())
	instanceName := fmt.Sprintf("test-cluster-provider-%s", acctest.RandString(8))

	createReq := &argocdv1.CreateInstanceRequest{
		OrganizationId: akpCli.OrgId,
		Name:           instanceName,
		Version:        "v3.0.0",
	}
	instance, err := akpCli.Cli.CreateInstance(ctx, createReq)
	if err != nil {
		panic(fmt.Sprintf("Failed to create instance: %v", err))
	}

	// Wait for instance to become available and get its ID
	for i := 0; i < 30; i++ { // Wait up to 5 minutes
		resp, err := akpCli.Cli.GetInstance(ctx, &argocdv1.GetInstanceRequest{
			OrganizationId: akpCli.OrgId,
			Id:             instance.GetInstance().Id,
			IdType:         idv1.Type_NAME,
		})
		if err == nil && resp.Instance != nil && resp.Instance.Id != "" {
			return resp.Instance.Id
		}
		time.Sleep(10 * time.Second)
	}

	panic("Test instance did not become available within timeout")
}

func cleanupTestInstance() {
	if createdInstanceId == "" || testAkpCli == nil {
		return
	}

	ctx := context.Background()
	ctx = httpctx.SetAuthorizationHeader(ctx, testAkpCli.Cred.Scheme(), testAkpCli.Cred.Credential())

	// Delete the instance
	_, _ = testAkpCli.Cli.DeleteInstance(ctx, &argocdv1.DeleteInstanceRequest{
		Id:             createdInstanceId,
		OrganizationId: testAkpCli.OrgId,
	})
}

func TestAccClusterResource(t *testing.T) {
	name := fmt.Sprintf("cluster-%s", acctest.RandString(10))
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			// Create and Read testing
			{
				Config: providerConfig + testAccClusterResourceConfig("small", name, "test one", getInstanceId()),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("akp_cluster.test", "id"),
					resource.TestCheckResourceAttr("akp_cluster.test", "namespace", "test"),
					resource.TestCheckResourceAttr("akp_cluster.test", "labels.test-label", "true"),
					resource.TestCheckResourceAttr("akp_cluster.test", "annotations.test-annotation", "false"),
					// spec
					resource.TestCheckResourceAttr("akp_cluster.test", "spec.description", "test one"),
					resource.TestCheckResourceAttr("akp_cluster.test", "spec.namespace_scoped", "true"),
					// spec.data
					resource.TestCheckResourceAttr("akp_cluster.test", "spec.data.size", "small"),
					resource.TestCheckResourceAttr("akp_cluster.test", "spec.data.auto_upgrade_disabled", "true"),
					resource.TestCheckResourceAttr("akp_cluster.test", "spec.data.kustomization", `  apiVersion: kustomize.config.k8s.io/v1beta1
  kind: Kustomization
  patches:
    - patch: |-
        apiVersion: apps/v1
        kind: Deployment
        metadata:
          name: argocd-repo-server
        spec:
          template:
            spec:
              containers:
              - name: argocd-repo-server
                resources:
                  limits:
                    memory: 2Gi
                  requests:
                    cpu: 750m
                    memory: 1Gi
      target:
        kind: Deployment
        name: argocd-repo-server
`),
					resource.TestCheckResourceAttr("akp_cluster.test", "spec.data.app_replication", "false"),
					resource.TestCheckResourceAttr("akp_cluster.test", "spec.data.redis_tunneling", "false"),
					resource.TestCheckResourceAttr("akp_cluster.test", "remove_agent_resources_on_destroy", "true"),
				),
			},
			// Update and Read testing
			{
				Config: providerConfig + testAccClusterResourceConfig("medium", name, "test two", getInstanceId()),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("akp_cluster.test", "spec.description", "test two"),
					resource.TestCheckResourceAttr("akp_cluster.test", "spec.data.size", "medium"),
				),
			},
			// Delete testing automatically occurs in TestCase
		},
	})
}

func testAccClusterResourceConfig(size, name, description, instanceId string) string {
	return fmt.Sprintf(`
resource "akp_cluster" "test" {
  instance_id = %q
  name      = %q
  namespace = "test"
  labels = {
    test-label = "true"
  }
  annotations = {
    test-annotation = "false"
  }
  spec = {
    namespace_scoped = true
    description      = %q
    data = {
      size                  = %q
      auto_upgrade_disabled = true
      kustomization         = <<EOF
  apiVersion: kustomize.config.k8s.io/v1beta1
  kind: Kustomization
  patches:
    - patch: |-
        apiVersion: apps/v1
        kind: Deployment
        metadata:
          name: argocd-repo-server
        spec:
          template:
            spec:
              containers:
              - name: argocd-repo-server
                resources:
                  limits:
                    memory: 2Gi
                  requests:
                    cpu: 750m
                    memory: 1Gi
      target:
        kind: Deployment
        name: argocd-repo-server
EOF
    }
  }
}
`, instanceId, name, description, size)
}

func TestAkpClusterResource_applyInstance(t *testing.T) {
	type args struct {
		plan             *types.Cluster
		apiReq           *argocdv1.ApplyInstanceRequest
		isCreate         bool
		applyInstance    func(context.Context, *argocdv1.ApplyInstanceRequest) (*argocdv1.ApplyInstanceResponse, error)
		upsertKubeConfig func(ctx context.Context, plan *types.Cluster) error
	}
	tests := []struct {
		name  string
		args  args
		want  *types.Cluster
		error error
	}{
		{
			name: "happy path, no kubeconfig",
			args: args{
				plan:     &types.Cluster{},
				isCreate: true,
				applyInstance: func(ctx context.Context, request *argocdv1.ApplyInstanceRequest) (*argocdv1.ApplyInstanceResponse, error) {
					return &argocdv1.ApplyInstanceResponse{}, nil
				},
				upsertKubeConfig: func(ctx context.Context, plan *types.Cluster) error {
					return errors.New("this should not be called")
				},
			},
			want:  &types.Cluster{},
			error: nil,
		},
		{
			name: "error path, no kubeconfig",
			args: args{
				plan:     &types.Cluster{},
				isCreate: true,
				applyInstance: func(ctx context.Context, request *argocdv1.ApplyInstanceRequest) (*argocdv1.ApplyInstanceResponse, error) {
					return &argocdv1.ApplyInstanceResponse{}, errors.New("some error")
				},
				upsertKubeConfig: func(ctx context.Context, plan *types.Cluster) error {
					return errors.New("this should not be called")
				},
			},
			want:  nil,
			error: fmt.Errorf("unable to create Argo CD instance: some error"),
		},
		{
			name: "happy path, with kubeconfig",
			args: args{
				plan: &types.Cluster{
					Kubeconfig: &types.Kubeconfig{
						Host: hashitype.StringValue("some-host"),
					},
				},
				applyInstance: func(ctx context.Context, request *argocdv1.ApplyInstanceRequest) (*argocdv1.ApplyInstanceResponse, error) {
					return &argocdv1.ApplyInstanceResponse{}, nil
				},
				isCreate: true,
				upsertKubeConfig: func(ctx context.Context, plan *types.Cluster) error {
					assert.Equal(t, &types.Cluster{
						Kubeconfig: &types.Kubeconfig{
							Host: hashitype.StringValue("some-host"),
						},
					}, plan)
					return nil
				},
			},
			want: &types.Cluster{
				Kubeconfig: &types.Kubeconfig{
					Host: hashitype.StringValue("some-host"),
				},
			},
			error: nil,
		},
		{
			name: "error path, with kubeconfig",
			args: args{
				plan: &types.Cluster{
					Kubeconfig: &types.Kubeconfig{
						Host: hashitype.StringValue("some-host"),
					},
				},
				isCreate: true,
				applyInstance: func(ctx context.Context, request *argocdv1.ApplyInstanceRequest) (*argocdv1.ApplyInstanceResponse, error) {
					return &argocdv1.ApplyInstanceResponse{}, nil
				},
				upsertKubeConfig: func(ctx context.Context, plan *types.Cluster) error {
					return errors.New("some kube apply error")
				},
			},
			want:  &types.Cluster{},
			error: fmt.Errorf("unable to apply manifests: some kube apply error"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &AkpClusterResource{}
			ctx := context.Background()
			got, err := r.applyInstance(ctx, tt.args.plan, tt.args.apiReq, tt.args.isCreate, tt.args.applyInstance, tt.args.upsertKubeConfig)
			assert.Equal(t, tt.error, err)
			assert.Equalf(t, tt.want, got, "applyInstance(%v, %v, %v)", tt.args.plan, tt.args.apiReq, tt.args.isCreate)
		})
	}
}

func TestAkpClusterResource_reApplyManifests(t *testing.T) {
	type args struct {
		plan             *types.Cluster
		apiReq           *argocdv1.ApplyInstanceRequest
		applyInstance    func(context.Context, *argocdv1.ApplyInstanceRequest) (*argocdv1.ApplyInstanceResponse, error)
		upsertKubeConfig func(ctx context.Context, plan *types.Cluster) error
	}
	tests := []struct {
		name  string
		args  args
		want  *types.Cluster
		error error
	}{
		{
			name: "error path, with kubeconfig",
			args: args{
				plan: &types.Cluster{
					Kubeconfig: &types.Kubeconfig{
						Host: hashitype.StringValue("some-host"),
					},
					ReapplyManifestsOnUpdate: hashitype.BoolValue(true),
				},
				applyInstance: func(ctx context.Context, request *argocdv1.ApplyInstanceRequest) (*argocdv1.ApplyInstanceResponse, error) {
					return &argocdv1.ApplyInstanceResponse{}, nil
				},
				upsertKubeConfig: func(ctx context.Context, plan *types.Cluster) error {
					return errors.New("some kube apply error")
				},
			},
			want:  &types.Cluster{},
			error: fmt.Errorf("unable to apply manifests: some kube apply error"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &AkpClusterResource{}
			ctx := context.Background()
			_, err := r.applyInstance(ctx, tt.args.plan, tt.args.apiReq, false, tt.args.applyInstance, tt.args.upsertKubeConfig)
			assert.Equal(t, tt.error, err)
		})
	}
}
