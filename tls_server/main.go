package main

import (
	"crypto/tls"
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

const serverKey = `-----BEGIN RSA PRIVATE KEY-----
MIIEogIBAAKCAQEAz4XstVd4Oz40UFE47sTyvVLKRfIXwPoBZYtmg4cJ7zkqT/jp
Ht/DTGyW9i5ciqw50+BUiy03Jz4KZE13p1g0QBGp6+/h6yYlqzTRlQtZyKfnFCSV
maxsMDNg53OmBPcubTvAijTjYi6KmKkkNDUfWhNJDpVOCirzr713e88a7Ofgx2Kg
8AWIv7eMNpPVgf/lqr/Xthbc3wdzbnGeH7PkBcRUL3jHtr8CmVK7/S2sZ+BhhMiz
h2q0PnFlPn7QJnYgSWivEwKeQKD5ihBZEKoTvc5p+thCuZlJbh5BHoyncZYEuGRs
awJ7hnt6Wrj3n90Wx4EmczqQI8eouOEBU2IGPwIDAQABAoIBAG7HtmHjzLoLbc/G
XBcSKeapEHOgVc/dZet4LARgjQwYPx8r5/5kQxRED55hh2M7py8U3Ozi+UNImjGL
LSm3K0VcaJIbn0VxOsCXVJ40zfwQrYmSMg4z/hkJBFzp3qFd9nm3YQQmUY/y+mee
O0ZnaVIjY2eWim9hGG8jPGpqV5xJAht4lkim5okM0JcODdI8f4NTRQK7Ly4ojl6e
frEAN37v7whSoV/nqueWUgljSuyegnRaQii/EdZ/accz2Yx1D7y7C/36SbVu0v6d
1xLPp8niYTsGRMRn4H1g4Yfb+ikzmGBxDl4h4dtCI97LMnqL8c4faH07zE9gKkyL
9gmxIAECgYEA6MaZ1dqRgi0+OJyfVWLPBEacL75wdjpg3ELCN1aaCvEpwbD1itGi
uxv23jloP2KRsoOcBTsBAPrzIr1F6m+nXmGsRAgIpIp3m2ppAjswykabGy4gqaq/
ZYLlhL+3iestnVD34ooskzOn+xX3ZZzL/9VVISmvZaSB4C0v5OO5dj8CgYEA5DpW
jjVbGAZKoSFjdnfRNPyXj8amr924o4OL40XxtVGxnXJNZgPUJT/l7yNvz0talqkY
ydG8EXTn6f0dwkJJsi59X43pkJOvcRCytJZVbz3UoljXl1c+t8N+HOscbkdeud7+
qzysuaMqBlOjp51o+COfHKWpGPw8bUFwp6rTcAECgYBCpJFoe5ILLf7RYG3xmT9h
nwGaXbB++MnAZdCK3V1u4oqs7Ifv/YLUhPYRu2fQNd02uZ1Slgj8idxQRtYcf2jw
iZAzTgWDbi0+LqHvxJh8YOiJSJSLXto7KtWhQu4+Kzm8/3BoyFxorHJcXXad/d0d
2qrruq5/sDCKyQUcR+JSKwKBgEgG7UgmKKCVY7okempp9/l7+i0i6Xo5REs+ech5
S4YIyUBCUZDvhmVXHC7zhRbr1iQUK8GN0ofD2GktO9YR64YzS9f7RoodQdIyDeQW
SdI4tkGiPjACFcYTCVsBDFD1bzzD1qWeGXjguKOUg+KjN/yP6Wg2E/7RpBRrAa24
qNABAoGAHNnuuX3RJOZxhx9r2EDp7z+NDbK+uAkDB2M/pP6XCJmNxJxy4K923u6U
zMynFP8/4CYwfGyYQTpr/Zwvh5sxD3bDAgm0RUbfFCzDDUj+p7/Q9dDO31BjL/dm
TRoqZRnq7mUQpcVdlTUoI1j098B9Tn6SmotcmqTDYYznkAJXVIs=
-----END RSA PRIVATE KEY-----
`

const serverCert = `-----BEGIN CERTIFICATE-----
MIIDIzCCAgugAwIBAgIQUyrA2pkGFYwhpC1wZEh7DjANBgkqhkiG9w0BAQsFADAo
MRQwEgYDVQQKEwtMb2cgQ291cmllcjEQMA4GA1UEAxMHOC44LjguODAeFw0xODEw
MTUxMTMyNDRaFw0yODEwMTIxMTMyNDRaMCgxFDASBgNVBAoTC0xvZyBDb3VyaWVy
MRAwDgYDVQQDEwc4LjguOC44MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKC
AQEAz4XstVd4Oz40UFE47sTyvVLKRfIXwPoBZYtmg4cJ7zkqT/jpHt/DTGyW9i5c
iqw50+BUiy03Jz4KZE13p1g0QBGp6+/h6yYlqzTRlQtZyKfnFCSVmaxsMDNg53Om
BPcubTvAijTjYi6KmKkkNDUfWhNJDpVOCirzr713e88a7Ofgx2Kg8AWIv7eMNpPV
gf/lqr/Xthbc3wdzbnGeH7PkBcRUL3jHtr8CmVK7/S2sZ+BhhMizh2q0PnFlPn7Q
JnYgSWivEwKeQKD5ihBZEKoTvc5p+thCuZlJbh5BHoyncZYEuGRsawJ7hnt6Wrj3
n90Wx4EmczqQI8eouOEBU2IGPwIDAQABo0kwRzAOBgNVHQ8BAf8EBAMCAqQwEwYD
VR0lBAwwCgYIKwYBBQUHAwEwDwYDVR0TAQH/BAUwAwEB/zAPBgNVHREECDAGhwR4
G30IMA0GCSqGSIb3DQEBCwUAA4IBAQBS7IM2cYCW94FCxf1F63HAyDrmcUk6gaXN
SRh+fJUkUM9OHrSDH1yeLeTdBSMT0sHQ80RvB1KuOFs/B2phEK5rBmmp23WyTs8C
fuLh3Svgm5vndxyNYQprrMw85vVf7+noxPH0YODWRBaSsLK52SSsJGiTv3oqehuj
Qxy5ygr+M802qcQiy6DHBOdH2q8RiELwUJEyWxQ4G0Kjzg9ANQLxJoocBcJN+Fm0
jDVLrro9fqMp+6zBSPdJkAkVXYpJPHhSlS/Cajp8HDgfMUe5OfQPEfLvMT18GnVp
FA1ATZO6BHrCaVawPga+oNA+DJmLOfibpVaLyQWg+Gg+BX1yBb4i
-----END CERTIFICATE-----
`

func main() {
	cer, err := tls.X509KeyPair([]byte(serverCert), []byte(serverKey))
	if err != nil {
		log.Fatal(err)
	}
	config := &tls.Config{Certificates: []tls.Certificate{cer}}

	l, err := tls.Listen("tcp", "120.27.125.8:443", config)
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()

	for {
		conn, err := l.Accept()
		if err != nil {
			log.Fatal(err)
		}
		go func(c net.Conn) {
			io.Copy(os.Stdout, c)
			fmt.Println()
			c.Close()
		}(conn)
	}
}
