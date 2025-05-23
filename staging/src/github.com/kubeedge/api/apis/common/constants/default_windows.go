//go:build windows

package constants

// Resources
const (
	KubeEdgePath       = "C:\\etc\\kubeedge\\"
	KubeEdgeUsrBinPath = "C:\\usr\\local\\bin"

	DefaultConfigDir   = "c:\\etc\\kubeedge\\config\\"
	EdgecoreConfigPath = "c:\\etc\\kubeedge\\config\\edgecore.yaml"

	// Certificates
	// DefaultCertPath is the default certificate path in edge node
	DefaultCertPath  = "c:\\etc\\kubeedge\\certs"
	DefaultCAFile    = "c:\\etc\\kubeedge\\ca\\rootCA.crt"
	DefaultCAKeyFile = "c:\\etc\\kubeedge\\ca\\rootCA.key"
	DefaultCertFile  = "c:\\etc\\kubeedge\\certs\\server.crt"
	DefaultKeyFile   = "c:\\etc\\kubeedge\\certs\\server.key"

	DefaultStreamCAFile   = "c:\\etc\\kubeedge\\ca\\streamCA.crt"
	DefaultStreamCertFile = "c:\\etc\\kubeedge\\certs\\stream.crt"
	DefaultStreamKeyFile  = "c:\\etc\\kubeedge\\certs\\stream.key"

	DefaultMqttCAFile   = "c:\\etc\\kubeedge\\ca\\rootCA.crt"
	DefaultMqttCertFile = "c:\\etc\\kubeedge\\certs\\server.crt"
	DefaultMqttKeyFile  = "c:\\etc\\kubeedge\\certs\\server.key"

	// Bootstrap file, contains token used by edgecore to apply for ca/cert
	BootstrapFile = "c:\\etc\\kubeedge\\bootstrap-edgecore.conf"

	// Edged
	DefaultRootDir               = "c:\\var\\lib\\kubelet"
	DefaultRemoteRuntimeEndpoint = "npipe://./pipe/containerd-containerd"
	DefaultRemoteImageEndpoint   = "npipe://./pipe/containerd-containerd"
	DefaultCNIConfDir            = "c:\\etc\\cni\\net.d"
	DefaultCNIBinDir             = "c:\\opt\\cni\\bin"
	DefaultCNICacheDir           = "c:\\var\\lib\\cni\\cache"
	DefaultVolumePluginDir       = "C:\\usr\\libexec\\kubernetes\\kubelet-plugins\\volume\\exec\\"

	// DefaultManifestsDir edge node default static pod path
	DefaultManifestsDir = "c:\\etc\\kubeedge\\manifests\\"
)
