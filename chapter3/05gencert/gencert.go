package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"math/big"
	"net"
	"os"
	"time"
)

// 自己署名SSL証明書とサーバの秘密鍵の生成
/*
PKI (Public-Key Infrastructure)
- ロール: Aさん, Bさん(証明書を提示する人), 信頼している第三者(PKIではCA(認証局))
- X.509やRFC3280に基づく扱い方がある
*/
func main() {
	// シリアルナンバーの生成 (本番は認証局によって発行される)
	max := new(big.Int).Lsh(big.NewInt(1), 128)
	serialNumber, _ := rand.Int(rand.Reader, max)

	// CAのサブジェクト
	subject := pkix.Name{
		Organization:       []string{"Manning Publications Co."},
		OrganizationalUnit: []string{"Books"},
		CommonName:         "Go Web Programming",
	}

	// X.509方式によるCA証明書用構造体を設定
	template := x509.Certificate{
		SerialNumber: serialNumber, // CAから発行されるシリアル番号
		Subject:      subject,      // 識別名
		NotBefore:    time.Now(),
		NotAfter:     time.Now().Add(365 * 24 * time.Hour),                         // 有効期限 1年
		KeyUsage:     x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature, // サーバ認証に使用されることを示す
		ExtKeyUsage:  []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},               // サーバ認証に使用されることを示す
		IPAddresses:  []net.IP{net.ParseIP("127.0.0.1")},                           // 127.0.0.1でのみ効力をもつと設定
	}

	// RSA(暗号化方式)秘密鍵を生成
	pk, _ := rsa.GenerateKey(rand.Reader, 2048)

	// 証明書の作成
	// 秘密鍵の構造体には公開鍵があってアクセス可能
	// derBytesはDER形式のバイトデータスライスを生成
	derBytes, _ := x509.CreateCertificate(rand.Reader, &template, &template, &pk.PublicKey, pk)

	// 証明書をPEM(Privacy Enhanced Email: プライバシー強化メール)形式で保存
	certOut, _ := os.Create("cert.pem")
	// pemを使って証明書データを符号化してcert.pemファイルにする
	pem.Encode(certOut, &pem.Block{Type: "CERTIFICATE", Bytes: derBytes})
	certOut.Close()

	// 秘密鍵をPEM形式で保存
	keyOut, _ := os.Create("key.pem")
	pem.Encode(keyOut, &pem.Block{Type: "RSA PRIVATE KEY", Bytes: x509.MarshalPKCS1PrivateKey(pk)})
	keyOut.Close()
}
