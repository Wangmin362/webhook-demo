package webhook

import (
	"k8s.io/api/admission/v1beta1"
	corev1 "k8s.io/api/core/v1"
	"net/http"
)

type WebHookServerInt interface {
	mutating(ar *v1beta1.AdmissionReview) *v1beta1.AdmissionResponse
	validating(ar *v1beta1.AdmissionReview) *v1beta1.AdmissionResponse
	Start()
	Stop()
}

// WebHookServerParameters webhook参数，port, certFile, keyFile是常规参数，应该是所有的webhook都有，而sidecarCfgFile未此
// webhook业务相关参数
type WebHookServerParameters struct {
	Port     int    // webhook server port
	CertFile string // path to the x509 certificate for https
	KeyFile  string // path to the x509 private key matching `CertFile`

	// 要注入的容器是啥？ 要注入的卷是啥?
	SidecarCfgFile string
}

type webHookServer struct {
	server        *http.Server
	sidecarConfig *Config
}

type Config struct {
	Containers []corev1.Container `yaml:"containers"`
	Volumes    []corev1.Volume    `yaml:"volumes"`
}

type patchOperation struct {
	Op    string      `json:"op"`
	Path  string      `json:"path"`
	Value interface{} `json:"value,omitempty"`
}
