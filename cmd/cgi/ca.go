package main

import "xie.sh.cn/panabit-ddns-go-manager/v2/pkg/env"

func updateCertificates() (int, any) {
	if err := env.CopyFile(env.BundledCertificates, env.Certificates, 0644); err != nil {
		return 1, err
	}
	return 0, "ok"
}
