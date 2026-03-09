package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

type DigestAuth struct {
	Username string
	Password string
	Client   *http.Client
}

func (d *DigestAuth) Do(req *http.Request) (*http.Response, error) {
	authRequest, _ := http.NewRequest(req.Method, req.URL.String(), nil)
	resp, err := d.Client.Do(authRequest)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusUnauthorized {
		return nil, fmt.Errorf("unexpected status: %d", resp.StatusCode)
	}

	authInfo := resp.Header.Get("WWW-Authenticate")
	fmt.Println(authInfo)

	// 解析认证信息
	authMap, err := parseDigestAuth(authInfo)
	if err != nil {
		return nil, err
	}
	fmt.Println(authMap)

	authorizationHeader := GenerateAuthorizationHeader(req.Method, req.URL.Path, d.Username, authMap["realm"], d.Password, authMap["nonce"], authMap["opaque"])
	fmt.Printf("Authorization: %s\n", authorizationHeader)

	req.Header.Set("Authorization", authorizationHeader)
	resp, err = d.Client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

// parseDigestAuth 解析 401 响应中的 Authorization 头并返回认证数据。
func parseDigestAuth(info string) (map[string]string, error) {
	if info == "" {
		return nil, fmt.Errorf("no WWW-Authenticate header")
	}
	if len(info) < 6 || info[:7] != "Digest " {
		return nil, fmt.Errorf("invalid WWW-Authenticate header: %s", info)
	}

	auth := info[7:]

	// 将认证数据拆分为其键值对。
	data := make(map[string]string)
	for _, pair := range strings.Split(auth, ",") {
		pair = strings.Trim(pair, " ")
		kv := strings.Split(pair, "=")
		if len(kv) != 2 {
			return nil, fmt.Errorf("无效的键值对：%s", pair)
		}
		data[kv[0]] = strings.Trim(kv[1], "\"")
	}

	// 检查必需的字段。
	if _, ok := data["realm"]; !ok {
		return nil, fmt.Errorf("Authorization 头中缺少 realm")
	}
	if _, ok := data["nonce"]; !ok {
		return nil, fmt.Errorf("Authorization 头中缺少 nonce")
	}
	if _, ok := data["qop"]; !ok {
		return nil, fmt.Errorf("Authorization 头中缺少 qop")
	}
	if _, ok := data["opaque"]; !ok {
		return nil, fmt.Errorf("Authorization 头中缺少 opaque")
	}

	// 返回认证数据。
	return data, nil
}

// generateNonce generates a random nonce using the current time as a seed.
//
// No parameters.
// Returns a string.
func generateNonce() string {
	rd := rand.New(rand.NewSource(time.Now().UnixNano()))
	return fmt.Sprintf("%X", rd.Int31())
}

// calculateHA1 calculates the HA1 value for Digest Access Authentication.
//
// Parameters:
// - username: the username
// - realm: the realm
// - password: the password
// Return type: string
func calculateHA1(username, realm, password string) string {
	h := md5.New()
	h.Write([]byte(username + ":" + realm + ":" + password))
	return hex.EncodeToString(h.Sum(nil))
}

// calculateHA2 calculates the hash of the method and URI using MD5 and returns the hexadecimal encoded string.
//
// method, uri string. string
// string
func calculateHA2(method, uri string) string {
	h := md5.New()
	h.Write([]byte(method + ":" + uri))
	return hex.EncodeToString(h.Sum(nil))
}

// GenerateAuthorizationHeader generates an authorization header for a given set of parameters.
//
// Parameters: method, uri, username, realm, password, nonce, opaque string.
// Returns a string.
func GenerateAuthorizationHeader(method, uri, username, realm, password, nonce, opaque string) string {
	nc := "00000000" // 固定值
	cnonce := generateNonce()
	qop := "auth"
	ha1 := calculateHA1(username, realm, password)
	ha2 := calculateHA2(method, uri)
	response := md5.Sum([]byte(ha1 + ":" + nonce + ":" + nc + ":" + cnonce + ":" + qop + ":" + ha2))

	return fmt.Sprintf(`Digest username="%s", realm="%s", nonce="%s", uri="%s", algorithm="MD5", qop=auth, nc=%s, cnonce="%s", response="%s", opaque="%s"`, username, realm, nonce, uri, nc, cnonce, hex.EncodeToString(response[:]), opaque)
}
