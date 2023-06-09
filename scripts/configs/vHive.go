package configs

type VHiveConfigStruct struct {
	FirecrackerKernelImgDownloadUrl string
	StargzVersion                   string
	VHiveRepoPath                   string
	VHiveRepoBranch                 string
	VHiveRepoUrl                    string
	VHiveSetupConfigPath            string
}

var VHive = VHiveConfigStruct{
	FirecrackerKernelImgDownloadUrl: "https://s3.amazonaws.com/spec.ccfc.min/img/hello/kernel/hello-vmlinux.bin",
	StargzVersion:                   "0.13.0",
	VHiveRepoPath:                   "",
	VHiveRepoBranch:                 "main",
	VHiveRepoUrl:                    "https://github.com/vhive-serverless/vHive.git",
	VHiveSetupConfigPath:            "",
}
