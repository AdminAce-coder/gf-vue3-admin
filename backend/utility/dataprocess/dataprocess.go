package dataprocess

import (
	"context"
	"strings"

	"github.com/duke-git/lancet/v2/cryptor"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/gogf/gf/v2/util/gconv"
)

// 获取加密密钥
func getEncryptKey(ctx context.Context) string {
	defaultKey := "abcdefghijklmnop" // 默认秘钥

	key, err := g.Cfg().Get(ctx, "encrypted.key")
	if err != nil {
		glog.New().Error(ctx, err)
		return defaultKey
	}

	// 检查配置值是否为空或只包含空格
	if key.String() == "" || len(strings.TrimSpace(key.String())) == 0 {
		glog.New().Warning(ctx, "未配置加密密钥，使用默认密钥")
		return defaultKey
	}

	return key.String()
}

// 加密数据
func AesEcbEncrypt(ctx context.Context, data string) string {
	keyString := getEncryptKey(ctx)
	glog.New().Infof(ctx, "正在使用秘钥,%s加密", keyString)
	encrypted := cryptor.AesEcbEncrypt([]byte(data), []byte(keyString))
	return cryptor.Base64StdEncode(gconv.String(encrypted)) // base64加密
}

// 解密数据
func AesEcbdecrypted(ctx context.Context, encrypted string) string {
	keyString := getEncryptKey(ctx)
	glog.New().Infof(ctx, "正在使用秘钥,%s 解密", keyString)
	// base64解密
	encrypted = cryptor.Base64StdDecode(encrypted)
	decrypted := cryptor.AesEcbDecrypt(gconv.Bytes(encrypted), []byte(keyString))
	return gconv.String(decrypted)
}
