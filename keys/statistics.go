package keys

import "fmt"

func GenSiteUv(host string) string {
	return fmt.Sprintf("siteuv:%s", host)
}

func GenSitePv(host string) string {
	return fmt.Sprintf("sitepv:%s", host)
}

func GenPagePv(host string, path string) string {
	return fmt.Sprintf("pagepv:%s:%s", host, path)
}

func GenSiteUvArchive(host string) string {
	return fmt.Sprintf("siteuv:archive:%s", host)
}
