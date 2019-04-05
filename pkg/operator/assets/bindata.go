// Code generated by go-bindata.
// sources:
// config/hiveadmission/apiservice.yaml
// config/hiveadmission/clusterdeployment-webhook.yaml
// config/hiveadmission/clusterimageset-webhook.yaml
// config/hiveadmission/deployment.yaml
// config/hiveadmission/dnszones-webhook.yaml
// config/hiveadmission/hiveadmission_rbac_role.yaml
// config/hiveadmission/hiveadmission_rbac_role_binding.yaml
// config/hiveadmission/service-account.yaml
// config/hiveadmission/service.yaml
// config/manager/deployment.yaml
// config/clusterimagesets/openshift-4.0-beta3.yaml
// config/clusterimagesets/openshift-4.0-latest.yaml
// config/external-dns/deployment.yaml
// config/external-dns/rbac_role.yaml
// config/external-dns/rbac_role_binding.yaml
// config/external-dns/service_account.yaml
// config/rbac/hive_admin_role.yaml
// config/rbac/hive_admin_role_binding.yaml
// config/rbac/hive_frontend_role.yaml
// config/rbac/hive_frontend_role_binding.yaml
// config/rbac/rbac_role.yaml
// config/rbac/rbac_role_binding.yaml
// DO NOT EDIT!

package assets

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

func (fi bindataFileInfo) Name() string {
	return fi.name
}
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}
func (fi bindataFileInfo) IsDir() bool {
	return false
}
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _configHiveadmissionApiserviceYaml = []byte(`---
# register as aggregated apiserver; this has a number of benefits:
#
# - allows other kubernetes components to talk to the the admission webhook using the ` + "`" + `kubernetes.default.svc` + "`" + ` service
# - allows other kubernetes components to use their in-cluster credentials to communicate with the webhook
# - allows you to test the webhook using kubectl
# - allows you to govern access to the webhook using RBAC
# - prevents other extension API servers from leaking their service account tokens to the webhook
#
# for more information, see: https://kubernetes.io/blog/2018/01/extensible-admission-is-beta
apiVersion: apiregistration.k8s.io/v1
kind: APIService
metadata:
  name: v1alpha1.admission.hive.openshift.io
  annotations:
    service.alpha.openshift.io/inject-cabundle: "true"
spec:
  group: admission.hive.openshift.io
  groupPriorityMinimum: 1000
  versionPriority: 15
  service:
    name: hiveadmission
    namespace: hive
  version: v1alpha1
`)

func configHiveadmissionApiserviceYamlBytes() ([]byte, error) {
	return _configHiveadmissionApiserviceYaml, nil
}

func configHiveadmissionApiserviceYaml() (*asset, error) {
	bytes, err := configHiveadmissionApiserviceYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "config/hiveadmission/apiservice.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _configHiveadmissionClusterdeploymentWebhookYaml = []byte(`---
apiVersion: admissionregistration.k8s.io/v1beta1
kind: ValidatingWebhookConfiguration
metadata:
  name: clusterdeployments.admission.hive.openshift.io
webhooks:
- name: clusterdeployments.admission.hive.openshift.io
  clientConfig:
    service:
      # reach the webhook via the registered aggregated API
      namespace: default
      name: kubernetes
      path: /apis/admission.hive.openshift.io/v1alpha1/clusterdeployments
  rules:
  - operations:
    - CREATE
    - UPDATE
    apiGroups:
    - hive.openshift.io
    apiVersions:
    - v1alpha1
    resources:
    - clusterdeployments
  failurePolicy: Fail
`)

func configHiveadmissionClusterdeploymentWebhookYamlBytes() ([]byte, error) {
	return _configHiveadmissionClusterdeploymentWebhookYaml, nil
}

func configHiveadmissionClusterdeploymentWebhookYaml() (*asset, error) {
	bytes, err := configHiveadmissionClusterdeploymentWebhookYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "config/hiveadmission/clusterdeployment-webhook.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _configHiveadmissionClusterimagesetWebhookYaml = []byte(`---
apiVersion: admissionregistration.k8s.io/v1beta1
kind: ValidatingWebhookConfiguration
metadata:
  name: clusterimagesets.admission.hive.openshift.io
webhooks:
- name: clusterimagesets.admission.hive.openshift.io
  clientConfig:
    service:
      # reach the webhook via the registered aggregated API
      namespace: default
      name: kubernetes
      path: /apis/admission.hive.openshift.io/v1alpha1/clusterimagesets
  rules:
  - operations:
    - CREATE
    - UPDATE
    apiGroups:
    - hive.openshift.io
    apiVersions:
    - v1alpha1
    resources:
    - clusterimagesets
  failurePolicy: Fail
`)

func configHiveadmissionClusterimagesetWebhookYamlBytes() ([]byte, error) {
	return _configHiveadmissionClusterimagesetWebhookYaml, nil
}

func configHiveadmissionClusterimagesetWebhookYaml() (*asset, error) {
	bytes, err := configHiveadmissionClusterimagesetWebhookYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "config/hiveadmission/clusterimageset-webhook.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _configHiveadmissionDeploymentYaml = []byte(`---
# to create the namespace-reservation-server
apiVersion: apps/v1
kind: Deployment
metadata:
  namespace: hive
  name: hiveadmission
  labels:
    app: hiveadmission
    hiveadmission: "true"
spec:
  replicas: 2
  selector:
    matchLabels:
      app: hiveadmission
      hiveadmission: "true"
  updateStrategy:
    rollingUpdate:
      maxUnavailable: 1
    type: RollingUpdate
  template:
    metadata:
      name: hiveadmission
      labels:
        app: hiveadmission
        hiveadmission: "true"
    spec:
      serviceAccountName: hiveadmission
      containers:
      - name: hiveadmission
        image: registry.svc.ci.openshift.org/openshift/hive-v4.0:hive
        imagePullPolicy: Always
        command:
        - "/opt/services/hiveadmission"
        - "--secure-port=9443"
        - "--audit-log-path=-"
        - "--tls-cert-file=/var/serving-cert/tls.crt"
        - "--tls-private-key-file=/var/serving-cert/tls.key"
        - "--v=8"
        ports:
        - containerPort: 9443
        volumeMounts:
        - mountPath: /var/serving-cert
          name: serving-cert
        readinessProbe:
          httpGet:
            path: /healthz
            port: 9443
            scheme: HTTPS
      volumes:
      - name: serving-cert
        secret:
          defaultMode: 420
          secretName: hiveadmission-serving-cert
`)

func configHiveadmissionDeploymentYamlBytes() ([]byte, error) {
	return _configHiveadmissionDeploymentYaml, nil
}

func configHiveadmissionDeploymentYaml() (*asset, error) {
	bytes, err := configHiveadmissionDeploymentYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "config/hiveadmission/deployment.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _configHiveadmissionDnszonesWebhookYaml = []byte(`---
# register to intercept DNSZone object creates and updates
apiVersion: admissionregistration.k8s.io/v1beta1
kind: ValidatingWebhookConfiguration
metadata:
  name: dnszones.admission.hive.openshift.io
webhooks:
- name: dnszones.admission.hive.openshift.io
  clientConfig:
    service:
      # reach the webhook via the registered aggregated API
      namespace: default
      name: kubernetes
      path: /apis/admission.hive.openshift.io/v1alpha1/dnszones
  rules:
  - operations:
    - CREATE
    - UPDATE
    apiGroups:
    - hive.openshift.io
    apiVersions:
    - v1alpha1
    resources:
    - dnszones
  failurePolicy: Fail
`)

func configHiveadmissionDnszonesWebhookYamlBytes() ([]byte, error) {
	return _configHiveadmissionDnszonesWebhookYaml, nil
}

func configHiveadmissionDnszonesWebhookYaml() (*asset, error) {
	bytes, err := configHiveadmissionDnszonesWebhookYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "config/hiveadmission/dnszones-webhook.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _configHiveadmissionHiveadmission_rbac_roleYaml = []byte(`apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  annotations:
  name: system:openshift:hive:hiveadmission
rules:
- apiGroups:
  - admission.hive.openshift.io
  resources:
  - dnszones
  verbs:
  - get
  - list
  - watch
- apiGroups:
  - ""
  resources:
  - configmaps
  verbs:
  - get
- apiGroups:
  - admissionregistration.k8s.io
  resources:
  - validatingwebhookconfigurations
  - mutatingwebhookconfigurations
  verbs:
  - get
  - list
  - watch
- apiGroups:
  - ""
  resources:
  - namespaces
  verbs:
  - get
  - list
  - watch
- apiGroups:
  - authorization.k8s.io
  resources:
  - subjectaccessreviews
  verbs:
  - create

`)

func configHiveadmissionHiveadmission_rbac_roleYamlBytes() ([]byte, error) {
	return _configHiveadmissionHiveadmission_rbac_roleYaml, nil
}

func configHiveadmissionHiveadmission_rbac_roleYaml() (*asset, error) {
	bytes, err := configHiveadmissionHiveadmission_rbac_roleYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "config/hiveadmission/hiveadmission_rbac_role.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _configHiveadmissionHiveadmission_rbac_role_bindingYaml = []byte(`apiVersion: v1
kind: List
items:
# to delegate authentication and authorization
- apiVersion: rbac.authorization.k8s.io/v1
  kind: ClusterRoleBinding
  metadata:
    name: auth-delegator-hiveadmission
  roleRef:
    kind: ClusterRole
    apiGroup: rbac.authorization.k8s.io
    name: system:auth-delegator
  subjects:
  - kind: ServiceAccount
    namespace: hive
    name: hiveadmission


# to let the admission server read the namespace reservations
- apiVersion: rbac.authorization.k8s.io/v1
  kind: ClusterRoleBinding
  metadata:
    name: hiveadmission-hive-hiveadmission
  roleRef:
    kind: ClusterRole
    apiGroup: rbac.authorization.k8s.io
    name: system:openshift:hive:hiveadmission
  subjects:
  - kind: ServiceAccount
    namespace: hive
    name: hiveadmission

# to read the config for terminating authentication
- apiVersion: rbac.authorization.k8s.io/v1
  kind: RoleBinding
  metadata:
    namespace: kube-system
    name: extension-server-authentication-reader-hiveadmission
  roleRef:
    kind: Role
    apiGroup: rbac.authorization.k8s.io
    name: extension-apiserver-authentication-reader
  subjects:
  - kind: ServiceAccount
    namespace: hive
    name: hiveadmission
`)

func configHiveadmissionHiveadmission_rbac_role_bindingYamlBytes() ([]byte, error) {
	return _configHiveadmissionHiveadmission_rbac_role_bindingYaml, nil
}

func configHiveadmissionHiveadmission_rbac_role_bindingYaml() (*asset, error) {
	bytes, err := configHiveadmissionHiveadmission_rbac_role_bindingYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "config/hiveadmission/hiveadmission_rbac_role_binding.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _configHiveadmissionServiceAccountYaml = []byte(`---
# to be able to assign powers to the hiveadmission process
apiVersion: v1
kind: ServiceAccount
metadata:
  namespace: hive
  name: hiveadmission
`)

func configHiveadmissionServiceAccountYamlBytes() ([]byte, error) {
	return _configHiveadmissionServiceAccountYaml, nil
}

func configHiveadmissionServiceAccountYaml() (*asset, error) {
	bytes, err := configHiveadmissionServiceAccountYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "config/hiveadmission/service-account.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _configHiveadmissionServiceYaml = []byte(`---
apiVersion: v1
kind: Service
metadata:
  namespace: hive
  name: hiveadmission
  annotations:
    service.alpha.openshift.io/serving-cert-secret-name: hiveadmission-serving-cert
spec:
  selector:
    app: hiveadmission
  ports:
  - port: 443
    targetPort: 9443
`)

func configHiveadmissionServiceYamlBytes() ([]byte, error) {
	return _configHiveadmissionServiceYaml, nil
}

func configHiveadmissionServiceYaml() (*asset, error) {
	bytes, err := configHiveadmissionServiceYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "config/hiveadmission/service.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _configManagerDeploymentYaml = []byte(`apiVersion: apps/v1
kind: Deployment
metadata:
  name: hive-controllers
  namespace: hive
  labels:
    control-plane: controller-manager
    controller-tools.k8s.io: "1.0"
spec:
  selector:
    matchLabels:
      control-plane: controller-manager
      controller-tools.k8s.io: "1.0"
  replicas: 1
  revisionHistoryLimit: 4
  template:
    metadata:
      labels:
        control-plane: controller-manager
        controller-tools.k8s.io: "1.0"
    spec:
      serviceAccountName: default
      volumes:
      - name: kubectl-cache
        emptyDir: {}
      containers:
      # By default we will use the latest CI images published from hive master:
      - image: registry.svc.ci.openshift.org/openshift/hive-v4.0:hive
        imagePullPolicy: Always
        name: manager
        resources:
          limits:
            cpu: 100m
            memory: 256Mi
          requests:
            cpu: 100m
            memory: 75Mi
        command:
          - /opt/services/manager
          - --log-level
          - debug
        volumeMounts:
        - name: kubectl-cache
          mountPath: /var/cache/kubectl
        env:
        - name: CLI_CACHE_DIR
          value: /var/cache/kubectl
      terminationGracePeriodSeconds: 10
`)

func configManagerDeploymentYamlBytes() ([]byte, error) {
	return _configManagerDeploymentYaml, nil
}

func configManagerDeploymentYaml() (*asset, error) {
	bytes, err := configManagerDeploymentYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "config/manager/deployment.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _configClusterimagesetsOpenshift40Beta3Yaml = []byte(`apiVersion: hive.openshift.io/v1alpha1
kind: ClusterImageSet
metadata:
  labels:
    controller-tools.k8s.io: "1.0"
  name: openshift-v4.0-beta3
spec:
  releaseImage: "quay.io/openshift-release-dev/ocp-release:4.0.0-0.9"
  hiveImage: "quay.io/twiest/hive-controller:20190403"
`)

func configClusterimagesetsOpenshift40Beta3YamlBytes() ([]byte, error) {
	return _configClusterimagesetsOpenshift40Beta3Yaml, nil
}

func configClusterimagesetsOpenshift40Beta3Yaml() (*asset, error) {
	bytes, err := configClusterimagesetsOpenshift40Beta3YamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "config/clusterimagesets/openshift-4.0-beta3.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _configClusterimagesetsOpenshift40LatestYaml = []byte(`apiVersion: hive.openshift.io/v1alpha1
kind: ClusterImageSet
metadata:
  labels:
    controller-tools.k8s.io: "1.0"
  name: openshift-v4.0-latest
spec:
  releaseImage: "registry.svc.ci.openshift.org/openshift/origin-release:v4.0"
  hiveImage: "registry.svc.ci.openshift.org/openshift/hive-v4.0:hive"
`)

func configClusterimagesetsOpenshift40LatestYamlBytes() ([]byte, error) {
	return _configClusterimagesetsOpenshift40LatestYaml, nil
}

func configClusterimagesetsOpenshift40LatestYaml() (*asset, error) {
	bytes, err := configClusterimagesetsOpenshift40LatestYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "config/clusterimagesets/openshift-4.0-latest.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _configExternalDnsDeploymentYaml = []byte(`apiVersion: apps/v1
kind: Deployment
metadata:
  name: external-dns
  namespace: hive
spec:
  replicas: 1
  selector:
    matchLabels:
      name: external-dns
  template:
    metadata:
      labels:
        name: external-dns
    spec:
      strategy:
        type: Recreate
      serviceAccountName: external-dns
      containers:
      - name: external-dns
        image: registry.svc.ci.openshift.org/openshift/hive-v4.0:external-dns
        args:
        - --source=crd
        - --crd-source-apiversion=hive.openshift.io/v1alpha1
        - --crd-source-kind=DNSEndpoint
        - --registry=noop
        - --policy=upsert-only
`)

func configExternalDnsDeploymentYamlBytes() ([]byte, error) {
	return _configExternalDnsDeploymentYaml, nil
}

func configExternalDnsDeploymentYaml() (*asset, error) {
	bytes, err := configExternalDnsDeploymentYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "config/external-dns/deployment.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _configExternalDnsRbac_roleYaml = []byte(`apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  name: external-dns
rules:
- apiGroups:
  - ""
  resources:
  - services
  verbs:
  - get
  - watch
  - list
- apiGroups:
  - hive.openshift.io
  resources:
  - dnsendpoints
  - dnsendpoints/status
  verbs:
  - get
  - list
  - watch
  - create
  - update
  - patch
  - delete
`)

func configExternalDnsRbac_roleYamlBytes() ([]byte, error) {
	return _configExternalDnsRbac_roleYaml, nil
}

func configExternalDnsRbac_roleYaml() (*asset, error) {
	bytes, err := configExternalDnsRbac_roleYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "config/external-dns/rbac_role.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _configExternalDnsRbac_role_bindingYaml = []byte(`apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  name: external-dns
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: external-dns
subjects:
- kind: ServiceAccount
  name: external-dns
  namespace: hive
`)

func configExternalDnsRbac_role_bindingYamlBytes() ([]byte, error) {
	return _configExternalDnsRbac_role_bindingYaml, nil
}

func configExternalDnsRbac_role_bindingYaml() (*asset, error) {
	bytes, err := configExternalDnsRbac_role_bindingYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "config/external-dns/rbac_role_binding.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _configExternalDnsService_accountYaml = []byte(`apiVersion: v1
kind: ServiceAccount
metadata:
  name: external-dns
  namespace: hive
`)

func configExternalDnsService_accountYamlBytes() ([]byte, error) {
	return _configExternalDnsService_accountYaml, nil
}

func configExternalDnsService_accountYaml() (*asset, error) {
	bytes, err := configExternalDnsService_accountYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "config/external-dns/service_account.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _configRbacHive_admin_roleYaml = []byte(`# hive-admin is a role intended for hive administrators who need to be able to debug
# cluster installations, and modify hive configuration.
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  name: hive-admin
rules:
- apiGroups:
  - batch
  resources:
  - jobs
  verbs:
  - get
  - list
  - watch
- apiGroups:
  - ""
  resources:
  - jobs
  - pods
  - pods/log
  verbs:
  - get
  - list
  - watch
- apiGroups:
  - hive.openshift.io
  resources:
  - clusterdeployments
  - dnszones
  - dnsendpoints
  - selectorsyncidentityprovider
  - selectorsyncset
  - syncidentityprovider
  - syncset
  verbs:
  - get
  - list
  - watch
- apiGroups:
  - hive.openshift.io
  resources:
  - clusterimagesets
  - hiveconfigs
  verbs:
  - get
  - list
  - watch
  - create
  - update
  - delete
`)

func configRbacHive_admin_roleYamlBytes() ([]byte, error) {
	return _configRbacHive_admin_roleYaml, nil
}

func configRbacHive_admin_roleYaml() (*asset, error) {
	bytes, err := configRbacHive_admin_roleYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "config/rbac/hive_admin_role.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _configRbacHive_admin_role_bindingYaml = []byte(`apiVersion: authorization.openshift.io/v1
kind: ClusterRoleBinding
metadata:
  name: hive-admin
roleRef:
  name: hive-admin
groupNames:
- hive-admins
subjects:
- kind: Group
  name: hive-admins
`)

func configRbacHive_admin_role_bindingYamlBytes() ([]byte, error) {
	return _configRbacHive_admin_role_bindingYaml, nil
}

func configRbacHive_admin_role_bindingYaml() (*asset, error) {
	bytes, err := configRbacHive_admin_role_bindingYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "config/rbac/hive_admin_role_binding.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _configRbacHive_frontend_roleYaml = []byte(`# hive-frontend is a role intended for integrating applications acting as a frontend
# to Hive. These applications will need quite powerful permissions in the Hive cluster
# to create namespaces to organize clusters, as well as all the required objects in those
# clusters.
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  name: hive-frontend
rules:
- apiGroups:
  - batch
  resources:
  - jobs
  verbs:
  - get
  - list
  - watch
- apiGroups:
  - ""
  resources:
  - pods
  - pods/log
  verbs:
  - get
  - list
  - watch
- apiGroups:
  - ""
  resources:
  - secrets
  - configmaps
  - namespaces
  verbs:
  - get
  - list
  - watch
  - create
  - update
  - patch
  - delete
- apiGroups:
  - hive.openshift.io
  resources:
  - clusterdeployments
  - dnszones
  - selectorsyncidentityprovider
  - syncidentityprovider
  - selectorsyncset
  - syncset
  verbs:
  - get
  - list
  - watch
  - create
  - update
  - patch
  - delete
- apiGroups:
  - hive.openshift.io
  resources:
  - clusterimagesets
  - hiveconfigs
  verbs:
  - get
  - list
  - watch
`)

func configRbacHive_frontend_roleYamlBytes() ([]byte, error) {
	return _configRbacHive_frontend_roleYaml, nil
}

func configRbacHive_frontend_roleYaml() (*asset, error) {
	bytes, err := configRbacHive_frontend_roleYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "config/rbac/hive_frontend_role.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _configRbacHive_frontend_role_bindingYaml = []byte(`apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  creationTimestamp: null
  name: hive-frontend
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: hive-frontend
subjects:
- kind: ServiceAccount
  name: hive-frontend
  namespace: hive
`)

func configRbacHive_frontend_role_bindingYamlBytes() ([]byte, error) {
	return _configRbacHive_frontend_role_bindingYaml, nil
}

func configRbacHive_frontend_role_bindingYaml() (*asset, error) {
	bytes, err := configRbacHive_frontend_role_bindingYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "config/rbac/hive_frontend_role_binding.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _configRbacRbac_roleYaml = []byte(`apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  creationTimestamp: null
  name: manager-role
rules:
- apiGroups:
  - batch
  resources:
  - jobs
  verbs:
  - get
  - list
  - watch
  - create
  - update
  - patch
  - delete
- apiGroups:
  - ""
  resources:
  - serviceaccounts
  - secrets
  - configmaps
  verbs:
  - get
  - list
  - watch
  - create
  - update
  - patch
  - delete
- apiGroups:
  - rbac.authorization.k8s.io
  resources:
  - roles
  - rolebindings
  verbs:
  - get
  - list
  - watch
  - create
  - update
  - patch
  - delete
- apiGroups:
  - hive.openshift.io
  resources:
  - clusterdeployments
  - clusterdeployments/status
  - clusterdeployments/finalizers
  verbs:
  - get
  - list
  - watch
  - create
  - update
  - patch
  - delete
- apiGroups:
  - hive.openshift.io
  resources:
  - clusterimagesets
  verbs:
  - get
  - list
  - watch
  - create
  - update
  - patch
  - delete
- apiGroups:
  - hive.openshift.io
  resources:
  - clusterimagesets/status
  verbs:
  - get
  - update
  - patch
- apiGroups:
  - hive.openshift.io
  resources:
  - dnszones
  - dnszones/status
  - dnszones/finalizers
  - dnsendpoints
  verbs:
  - get
  - list
  - watch
  - create
  - update
  - patch
  - delete
- apiGroups:
  - apiextensions.k8s.io
  resources:
  - customresourcedefinitions
  verbs:
  - get
  - list
  - watch
- apiGroups:
  - clusterregistry.k8s.io
  resources:
  - clusters
  verbs:
  - get
  - list
  - watch
  - create
  - update
  - patch
  - delete
- apiGroups:
  - core.federation.k8s.io
  resources:
  - federatedclusters
  verbs:
  - get
  - list
  - watch
  - create
  - update
  - patch
  - delete
- apiGroups:
  - hive.openshift.io
  resources:
  - clusterdeployments
  verbs:
  - get
  - watch
  - update
- apiGroups:
  - hive.openshift.io
  resources:
  - syncsets
  verbs:
  - get
  - create
  - update
  - delete
  - patch
  - list
  - watch
- apiGroups:
  - hive.openshift.io
  resources:
  - syncidentityproviders
  verbs:
  - get
  - list
  - watch
- apiGroups:
  - hive.openshift.io
  resources:
  - selectorsyncidentityproviders
  verbs:
  - get
  - list
  - watch
- apiGroups:
  - hive.openshift.io
  resources:
  - clusterdeployments
  verbs:
  - get
  - list
  - watch
- apiGroups:
  - hive.openshift.io
  resources:
  - clusterdeployments
  verbs:
  - get
  - list
  - watch
- apiGroups:
  - hive.openshift.io
  resources:
  - syncsets
  verbs:
  - get
  - create
  - update
  - delete
  - patch
  - list
  - watch
- apiGroups:
  - hive.openshift.io
  resources:
  - selectorsyncsets
  verbs:
  - get
  - create
  - update
  - delete
  - patch
  - list
  - watch
`)

func configRbacRbac_roleYamlBytes() ([]byte, error) {
	return _configRbacRbac_roleYaml, nil
}

func configRbacRbac_roleYaml() (*asset, error) {
	bytes, err := configRbacRbac_roleYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "config/rbac/rbac_role.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

var _configRbacRbac_role_bindingYaml = []byte(`apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  creationTimestamp: null
  name: manager-rolebinding
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: manager-role
subjects:
- kind: ServiceAccount
  name: default
  namespace: system
`)

func configRbacRbac_role_bindingYamlBytes() ([]byte, error) {
	return _configRbacRbac_role_bindingYaml, nil
}

func configRbacRbac_role_bindingYaml() (*asset, error) {
	bytes, err := configRbacRbac_role_bindingYamlBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "config/rbac/rbac_role_binding.yaml", size: 0, mode: os.FileMode(0), modTime: time.Unix(0, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"config/hiveadmission/apiservice.yaml":                      configHiveadmissionApiserviceYaml,
	"config/hiveadmission/clusterdeployment-webhook.yaml":       configHiveadmissionClusterdeploymentWebhookYaml,
	"config/hiveadmission/clusterimageset-webhook.yaml":         configHiveadmissionClusterimagesetWebhookYaml,
	"config/hiveadmission/deployment.yaml":                      configHiveadmissionDeploymentYaml,
	"config/hiveadmission/dnszones-webhook.yaml":                configHiveadmissionDnszonesWebhookYaml,
	"config/hiveadmission/hiveadmission_rbac_role.yaml":         configHiveadmissionHiveadmission_rbac_roleYaml,
	"config/hiveadmission/hiveadmission_rbac_role_binding.yaml": configHiveadmissionHiveadmission_rbac_role_bindingYaml,
	"config/hiveadmission/service-account.yaml":                 configHiveadmissionServiceAccountYaml,
	"config/hiveadmission/service.yaml":                         configHiveadmissionServiceYaml,
	"config/manager/deployment.yaml":                            configManagerDeploymentYaml,
	"config/clusterimagesets/openshift-4.0-beta3.yaml":          configClusterimagesetsOpenshift40Beta3Yaml,
	"config/clusterimagesets/openshift-4.0-latest.yaml":         configClusterimagesetsOpenshift40LatestYaml,
	"config/external-dns/deployment.yaml":                       configExternalDnsDeploymentYaml,
	"config/external-dns/rbac_role.yaml":                        configExternalDnsRbac_roleYaml,
	"config/external-dns/rbac_role_binding.yaml":                configExternalDnsRbac_role_bindingYaml,
	"config/external-dns/service_account.yaml":                  configExternalDnsService_accountYaml,
	"config/rbac/hive_admin_role.yaml":                          configRbacHive_admin_roleYaml,
	"config/rbac/hive_admin_role_binding.yaml":                  configRbacHive_admin_role_bindingYaml,
	"config/rbac/hive_frontend_role.yaml":                       configRbacHive_frontend_roleYaml,
	"config/rbac/hive_frontend_role_binding.yaml":               configRbacHive_frontend_role_bindingYaml,
	"config/rbac/rbac_role.yaml":                                configRbacRbac_roleYaml,
	"config/rbac/rbac_role_binding.yaml":                        configRbacRbac_role_bindingYaml,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}

var _bintree = &bintree{nil, map[string]*bintree{
	"config": {nil, map[string]*bintree{
		"clusterimagesets": {nil, map[string]*bintree{
			"openshift-4.0-beta3.yaml":  {configClusterimagesetsOpenshift40Beta3Yaml, map[string]*bintree{}},
			"openshift-4.0-latest.yaml": {configClusterimagesetsOpenshift40LatestYaml, map[string]*bintree{}},
		}},
		"external-dns": {nil, map[string]*bintree{
			"deployment.yaml":        {configExternalDnsDeploymentYaml, map[string]*bintree{}},
			"rbac_role.yaml":         {configExternalDnsRbac_roleYaml, map[string]*bintree{}},
			"rbac_role_binding.yaml": {configExternalDnsRbac_role_bindingYaml, map[string]*bintree{}},
			"service_account.yaml":   {configExternalDnsService_accountYaml, map[string]*bintree{}},
		}},
		"hiveadmission": {nil, map[string]*bintree{
			"apiservice.yaml":                      {configHiveadmissionApiserviceYaml, map[string]*bintree{}},
			"clusterdeployment-webhook.yaml":       {configHiveadmissionClusterdeploymentWebhookYaml, map[string]*bintree{}},
			"clusterimageset-webhook.yaml":         {configHiveadmissionClusterimagesetWebhookYaml, map[string]*bintree{}},
			"deployment.yaml":                      {configHiveadmissionDeploymentYaml, map[string]*bintree{}},
			"dnszones-webhook.yaml":                {configHiveadmissionDnszonesWebhookYaml, map[string]*bintree{}},
			"hiveadmission_rbac_role.yaml":         {configHiveadmissionHiveadmission_rbac_roleYaml, map[string]*bintree{}},
			"hiveadmission_rbac_role_binding.yaml": {configHiveadmissionHiveadmission_rbac_role_bindingYaml, map[string]*bintree{}},
			"service-account.yaml":                 {configHiveadmissionServiceAccountYaml, map[string]*bintree{}},
			"service.yaml":                         {configHiveadmissionServiceYaml, map[string]*bintree{}},
		}},
		"manager": {nil, map[string]*bintree{
			"deployment.yaml": {configManagerDeploymentYaml, map[string]*bintree{}},
		}},
		"rbac": {nil, map[string]*bintree{
			"hive_admin_role.yaml":            {configRbacHive_admin_roleYaml, map[string]*bintree{}},
			"hive_admin_role_binding.yaml":    {configRbacHive_admin_role_bindingYaml, map[string]*bintree{}},
			"hive_frontend_role.yaml":         {configRbacHive_frontend_roleYaml, map[string]*bintree{}},
			"hive_frontend_role_binding.yaml": {configRbacHive_frontend_role_bindingYaml, map[string]*bintree{}},
			"rbac_role.yaml":                  {configRbacRbac_roleYaml, map[string]*bintree{}},
			"rbac_role_binding.yaml":          {configRbacRbac_role_bindingYaml, map[string]*bintree{}},
		}},
	}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}
