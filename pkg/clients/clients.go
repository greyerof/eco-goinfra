package clients

import (
	"fmt"
	"log"
	"os"

	"github.com/golang/glog"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/kubernetes"

	argocdOperatorv1alpha1 "github.com/argoproj-labs/argocd-operator/api/v1alpha1"
	bmhv1alpha1 "github.com/metal3-io/baremetal-operator/apis/metal3.io/v1alpha1"
	performanceV2 "github.com/openshift/cluster-node-tuning-operator/pkg/apis/performanceprofile/v2"

	clientConfigV1 "github.com/openshift/client-go/config/clientset/versioned/typed/config/v1"
	v1security "github.com/openshift/client-go/security/clientset/versioned/typed/security/v1"
	mcv1 "github.com/openshift/machine-config-operator/pkg/apis/machineconfiguration.openshift.io/v1"
	ptpV1 "github.com/openshift/ptp-operator/pkg/client/clientset/versioned/typed/ptp/v1"
	olm2 "github.com/operator-framework/operator-lifecycle-manager/pkg/api/client/clientset/versioned/scheme"

	olmv1 "github.com/operator-framework/operator-lifecycle-manager/pkg/api/client/clientset/versioned/typed/operators/v1"
	olm "github.com/operator-framework/operator-lifecycle-manager/pkg/api/client/clientset/versioned/typed/operators/v1alpha1"

	pkgManifestV1 "github.com/operator-framework/operator-lifecycle-manager/pkg/package-server/apis/operators/v1"
	clientPkgManifestV1 "github.com/operator-framework/operator-lifecycle-manager/pkg/package-server/client/clientset/versioned/typed/operators/v1"

	apiExt "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	"k8s.io/apimachinery/pkg/runtime"
	appsV1Client "k8s.io/client-go/kubernetes/typed/apps/v1"
	networkV1Client "k8s.io/client-go/kubernetes/typed/networking/v1"
	rbacV1Client "k8s.io/client-go/kubernetes/typed/rbac/v1"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	runtimeClient "sigs.k8s.io/controller-runtime/pkg/client"

	netAttDefV1 "github.com/k8snetworkplumbingwg/network-attachment-definition-client/pkg/apis/k8s.cni.cncf.io/v1"
	clientNetAttDefV1 "github.com/k8snetworkplumbingwg/network-attachment-definition-client/pkg/client/clientset/versioned/typed/k8s.cni.cncf.io/v1"
	srIovV1 "github.com/k8snetworkplumbingwg/sriov-network-operator/api/v1"

	clientSrIov "github.com/k8snetworkplumbingwg/sriov-network-operator/pkg/client/clientset/versioned"
	clientSrIovFake "github.com/k8snetworkplumbingwg/sriov-network-operator/pkg/client/clientset/versioned/fake"
	clientSrIovV1 "github.com/k8snetworkplumbingwg/sriov-network-operator/pkg/client/clientset/versioned/typed/sriovnetwork/v1"

	clientMachineConfigV1 "github.com/openshift/machine-config-operator/pkg/generated/clientset/versioned/typed/machineconfiguration.openshift.io/v1"

	nmstatev1 "github.com/nmstate/kubernetes-nmstate/api/v1"
	nmstateV1alpha1 "github.com/nmstate/kubernetes-nmstate/api/v1alpha1"

	lcasgv1alpha1 "github.com/openshift-kni/lifecycle-agent/api/seedgenerator/v1alpha1"
	lcav1alpha1 "github.com/openshift-kni/lifecycle-agent/api/v1alpha1"
	operatorV1 "github.com/openshift/api/operator/v1"
	routev1 "github.com/openshift/api/route/v1"
	hiveextV1Beta1 "github.com/openshift/assisted-service/api/hiveextension/v1beta1"
	agentInstallV1Beta1 "github.com/openshift/assisted-service/api/v1beta1"
	hiveV1 "github.com/openshift/hive/apis/hive/v1"
	moduleV1Beta1 "github.com/rh-ecosystem-edge/kernel-module-management/api/v1beta1"
	"k8s.io/client-go/kubernetes/scheme"
	coreV1Client "k8s.io/client-go/kubernetes/typed/core/v1"
	storageV1Client "k8s.io/client-go/kubernetes/typed/storage/v1"
	policiesv1 "open-cluster-management.io/governance-policy-propagator/api/v1"

	appsv1 "k8s.io/api/apps/v1"
	scalingv1 "k8s.io/api/autoscaling/v1"
	corev1 "k8s.io/api/core/v1"
	policyv1 "k8s.io/api/policy/v1"
	rbacv1 "k8s.io/api/rbac/v1"
	storagev1 "k8s.io/api/storage/v1"
	k8sFakeClient "k8s.io/client-go/kubernetes/fake"
	fakeRuntimeClient "sigs.k8s.io/controller-runtime/pkg/client/fake"

	nvidiagpuv1 "github.com/NVIDIA/gpu-operator/api/v1"
	grafanaV4V1Alpha1 "github.com/grafana-operator/grafana-operator/v4/api/integreatly/v1alpha1"
	multinetpolicyclientv1 "github.com/k8snetworkplumbingwg/multi-networkpolicy/pkg/client/clientset/versioned/typed/k8s.cni.cncf.io/v1beta1"
	cguapiv1alpha1 "github.com/openshift-kni/cluster-group-upgrades-operator/pkg/api/clustergroupupgrades/v1alpha1"
	machinev1beta1client "github.com/openshift/client-go/machine/clientset/versioned/typed/machine/v1beta1"
	operatorv1alpha1 "github.com/openshift/client-go/operator/clientset/versioned/typed/operator/v1alpha1"
	nfdv1 "github.com/openshift/cluster-nfd-operator/api/v1"
	lsoV1alpha1 "github.com/openshift/local-storage-operator/api/v1alpha1"
	mcmV1Beta1 "github.com/rh-ecosystem-edge/kernel-module-management/api-hub/v1beta1"
	velerov1 "github.com/vmware-tanzu/velero/pkg/apis/velero/v1"
	veleroClient "github.com/vmware-tanzu/velero/pkg/generated/clientset/versioned"
	veleroFakeClient "github.com/vmware-tanzu/velero/pkg/generated/clientset/versioned/fake"
	veleroV1Client "github.com/vmware-tanzu/velero/pkg/generated/clientset/versioned/typed/velero/v1"
	policiesv1beta1 "open-cluster-management.io/governance-policy-propagator/api/v1beta1"
	placementrulev1 "open-cluster-management.io/multicloud-operators-subscription/pkg/apis/apps/placementrule/v1"
)

// Settings provides the struct to talk with relevant API.
type Settings struct {
	KubeconfigPath string
	K8sClient      kubernetes.Interface
	coreV1Client.CoreV1Interface
	clientConfigV1.ConfigV1Interface
	clientMachineConfigV1.MachineconfigurationV1Interface
	networkV1Client.NetworkingV1Interface
	appsV1Client.AppsV1Interface
	rbacV1Client.RbacV1Interface
	ClientSrIov clientSrIov.Interface
	clientSrIovV1.SriovnetworkV1Interface
	Config *rest.Config
	runtimeClient.Client
	ptpV1.PtpV1Interface
	v1security.SecurityV1Interface
	olm.OperatorsV1alpha1Interface
	clientNetAttDefV1.K8sCniCncfIoV1Interface
	dynamic.Interface
	olmv1.OperatorsV1Interface
	multinetpolicyclientv1.K8sCniCncfIoV1beta1Interface
	PackageManifestInterface clientPkgManifestV1.OperatorsV1Interface
	operatorv1alpha1.OperatorV1alpha1Interface
	grafanaV4V1Alpha1.Grafana
	LocalVolumeInterface lsoV1alpha1.LocalVolumeSet
	machinev1beta1client.MachineV1beta1Interface
	storageV1Client.StorageV1Interface
	VeleroClient veleroClient.Interface
	veleroV1Client.VeleroV1Interface
}

// New returns a *Settings with the given kubeconfig.
//
//nolint:funlen
func New(kubeconfig string) *Settings {
	var (
		config *rest.Config
		err    error
	)

	if kubeconfig == "" {
		kubeconfig = os.Getenv("KUBECONFIG")
	}

	if kubeconfig != "" {
		log.Printf("Loading kube client config from path %q", kubeconfig)

		config, err = clientcmd.BuildConfigFromFlags("", kubeconfig)
	} else {
		log.Print("Using in-cluster kube client config")

		config, err = rest.InClusterConfig()
	}

	if err != nil {
		return nil
	}

	clientSet := &Settings{}
	clientSet.CoreV1Interface = coreV1Client.NewForConfigOrDie(config)
	clientSet.ConfigV1Interface = clientConfigV1.NewForConfigOrDie(config)
	clientSet.MachineconfigurationV1Interface = clientMachineConfigV1.NewForConfigOrDie(config)
	clientSet.AppsV1Interface = appsV1Client.NewForConfigOrDie(config)
	clientSet.ClientSrIov = clientSrIov.NewForConfigOrDie(config)
	clientSet.SriovnetworkV1Interface = clientSrIovV1.NewForConfigOrDie(config)
	clientSet.NetworkingV1Interface = networkV1Client.NewForConfigOrDie(config)
	clientSet.PtpV1Interface = ptpV1.NewForConfigOrDie(config)
	clientSet.RbacV1Interface = rbacV1Client.NewForConfigOrDie(config)
	clientSet.OperatorsV1alpha1Interface = olm.NewForConfigOrDie(config)
	clientSet.K8sCniCncfIoV1Interface = clientNetAttDefV1.NewForConfigOrDie(config)
	clientSet.Interface = dynamic.NewForConfigOrDie(config)
	clientSet.OperatorsV1Interface = olmv1.NewForConfigOrDie(config)
	clientSet.PackageManifestInterface = clientPkgManifestV1.NewForConfigOrDie(config)
	clientSet.SecurityV1Interface = v1security.NewForConfigOrDie(config)
	clientSet.OperatorV1alpha1Interface = operatorv1alpha1.NewForConfigOrDie(config)
	clientSet.MachineV1beta1Interface = machinev1beta1client.NewForConfigOrDie(config)
	clientSet.K8sCniCncfIoV1beta1Interface = multinetpolicyclientv1.NewForConfigOrDie(config)
	clientSet.StorageV1Interface = storageV1Client.NewForConfigOrDie(config)
	clientSet.K8sClient = kubernetes.NewForConfigOrDie(config)
	clientSet.VeleroClient = veleroClient.NewForConfigOrDie(config)
	clientSet.VeleroV1Interface = veleroV1Client.NewForConfigOrDie(config)
	clientSet.Config = config

	crScheme := runtime.NewScheme()
	err = SetScheme(crScheme)

	if err != nil {
		log.Print("Error to load apiClient scheme")

		return nil
	}

	clientSet.Client, err = runtimeClient.New(config, runtimeClient.Options{
		Scheme: crScheme,
	})

	if err != nil {
		log.Print("Error to create apiClient")

		return nil
	}

	clientSet.KubeconfigPath = kubeconfig

	return clientSet
}

// SetScheme returns mutated apiClient's scheme.
//
//nolint:funlen
func SetScheme(crScheme *runtime.Scheme) error {
	if err := scheme.AddToScheme(crScheme); err != nil {
		return err
	}

	if err := netAttDefV1.SchemeBuilder.AddToScheme(crScheme); err != nil {
		return err
	}

	if err := srIovV1.AddToScheme(crScheme); err != nil {
		return err
	}

	if err := mcv1.AddToScheme(crScheme); err != nil {
		return err
	}

	if err := apiExt.AddToScheme(crScheme); err != nil {
		return err
	}

	if err := performanceV2.AddToScheme(crScheme); err != nil {
		return err
	}

	if err := lcav1alpha1.AddToScheme(crScheme); err != nil {
		return err
	}

	if err := lcasgv1alpha1.AddToScheme(crScheme); err != nil {
		return err
	}

	if err := operatorV1.Install(crScheme); err != nil {
		return err
	}

	if err := olm2.AddToScheme(crScheme); err != nil {
		return err
	}

	if err := bmhv1alpha1.AddToScheme(crScheme); err != nil {
		return err
	}

	if err := lsoV1alpha1.AddToScheme(crScheme); err != nil {
		return err
	}

	if err := hiveextV1Beta1.AddToScheme(crScheme); err != nil {
		return err
	}

	if err := hiveV1.AddToScheme(crScheme); err != nil {
		return err
	}

	if err := agentInstallV1Beta1.AddToScheme(crScheme); err != nil {
		return err
	}

	if err := moduleV1Beta1.AddToScheme(crScheme); err != nil {
		return err
	}

	if err := mcmV1Beta1.AddToScheme(crScheme); err != nil {
		return err
	}

	if err := nvidiagpuv1.AddToScheme(crScheme); err != nil {
		return err
	}

	if err := nfdv1.AddToScheme(crScheme); err != nil {
		return err
	}

	if err := grafanaV4V1Alpha1.AddToScheme(crScheme); err != nil {
		return err
	}

	if err := pkgManifestV1.AddToScheme(crScheme); err != nil {
		return err
	}

	if err := nmstatev1.AddToScheme(crScheme); err != nil {
		return err
	}

	if err := nmstateV1alpha1.AddToScheme(crScheme); err != nil {
		return err
	}

	if err := argocdOperatorv1alpha1.AddToScheme(crScheme); err != nil {
		return err
	}

	if err := policiesv1.AddToScheme(crScheme); err != nil {
		return err
	}

	if err := cguapiv1alpha1.AddToScheme(crScheme); err != nil {
		return err
	}

	if err := policiesv1beta1.AddToScheme(crScheme); err != nil {
		return err
	}

	if err := placementrulev1.AddToScheme(crScheme); err != nil {
		return err
	}

	if err := routev1.AddToScheme(crScheme); err != nil {
		return err
	}

	return nil
}

// GetAPIClient implements the cluster.APIClientGetter interface.
func (settings *Settings) GetAPIClient() (*Settings, error) {
	if settings == nil {
		glog.V(100).Infof("APIClient is nil")

		return nil, fmt.Errorf("APIClient cannot be nil")
	}

	return settings, nil
}

// GetTestClients returns a fake clientset for testing.
//
//nolint:funlen
func GetTestClients(k8sMockObjects []runtime.Object) *Settings {
	clientSet := &Settings{}

	var k8sClientObjects, genericClientObjects, srIovObjects, veleroClientObjects []runtime.Object

	//nolint:varnamelen
	for _, v := range k8sMockObjects {
		// Based on what type of object is, populate certain object slices
		// with what is supported by a certain client.
		// Add more items below if/when needed.
		switch v.(type) {
		// K8s Client Objects
		case *corev1.ServiceAccount:
			k8sClientObjects = append(k8sClientObjects, v)
		case *rbacv1.ClusterRole:
			k8sClientObjects = append(k8sClientObjects, v)
		case *rbacv1.ClusterRoleBinding:
			k8sClientObjects = append(k8sClientObjects, v)
		case *rbacv1.Role:
			k8sClientObjects = append(k8sClientObjects, v)
		case *rbacv1.RoleBinding:
			k8sClientObjects = append(k8sClientObjects, v)
		case *corev1.Pod:
			k8sClientObjects = append(k8sClientObjects, v)
		case *corev1.Service:
			k8sClientObjects = append(k8sClientObjects, v)
		case *corev1.Node:
			k8sClientObjects = append(k8sClientObjects, v)
		case *appsv1.Deployment:
			k8sClientObjects = append(k8sClientObjects, v)
		case *appsv1.StatefulSet:
			k8sClientObjects = append(k8sClientObjects, v)
		case *corev1.ResourceQuota:
			k8sClientObjects = append(k8sClientObjects, v)
		case *corev1.PersistentVolume:
			k8sClientObjects = append(k8sClientObjects, v)
		case *corev1.PersistentVolumeClaim:
			k8sClientObjects = append(k8sClientObjects, v)
		case *policyv1.PodDisruptionBudget:
			k8sClientObjects = append(k8sClientObjects, v)
		case *scalingv1.HorizontalPodAutoscaler:
			k8sClientObjects = append(k8sClientObjects, v)
		case *storagev1.StorageClass:
			k8sClientObjects = append(k8sClientObjects, v)
		// Generic Client Objects
		case *routev1.Route:
			genericClientObjects = append(genericClientObjects, v)
		// Velero Client Objects
		case *velerov1.Backup:
			veleroClientObjects = append(veleroClientObjects, v)
		case *velerov1.Restore:
			veleroClientObjects = append(veleroClientObjects, v)
		// SrIov Client Objects
		case *srIovV1.SriovNetwork:
			srIovObjects = append(srIovObjects, v)
		case *srIovV1.SriovNetworkNodePolicy:
			srIovObjects = append(srIovObjects, v)
		case *srIovV1.SriovOperatorConfig:
			srIovObjects = append(srIovObjects, v)
		case *srIovV1.SriovNetworkNodeState:
			srIovObjects = append(srIovObjects, v)
		}
	}

	// Assign the fake clientset to the clientSet
	clientSet.K8sClient = k8sFakeClient.NewSimpleClientset(k8sClientObjects...)
	clientSet.CoreV1Interface = clientSet.K8sClient.CoreV1()
	clientSet.AppsV1Interface = clientSet.K8sClient.AppsV1()
	clientSet.NetworkingV1Interface = clientSet.K8sClient.NetworkingV1()
	clientSet.RbacV1Interface = clientSet.K8sClient.RbacV1()
	clientSet.ClientSrIov = clientSrIovFake.NewSimpleClientset(srIovObjects...)

	// Assign the fake velero clientset to the clientSet
	clientSet.VeleroClient = veleroFakeClient.NewSimpleClientset(veleroClientObjects...)
	clientSet.VeleroV1Interface = clientSet.VeleroClient.VeleroV1()

	// Update the generic client with schemes of generic resources
	fakeClientScheme := runtime.NewScheme()

	err := SetScheme(fakeClientScheme)
	if err != nil {
		return nil
	}

	// Add fake runtime client to clientSet runtime client
	clientSet.Client = fakeRuntimeClient.NewClientBuilder().WithScheme(fakeClientScheme).
		WithRuntimeObjects(genericClientObjects...).Build()

	return clientSet
}
