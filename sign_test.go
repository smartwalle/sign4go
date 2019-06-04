package sign4go

import (
	"crypto"
	_ "crypto/md5"
	_ "crypto/sha1"
	"net/url"
	"strings"
	"testing"
)

func BenchmarkSign(b *testing.B) {
	var h = NewHashSign(crypto.MD5)
	for i := 0; i < b.N; i++ {
		var form = make(url.Values, 0)
		form.Set("c", "30")
		form.Set("b", "20")
		form.Set("a", "30")
		_, _ = h.Sign(form, func(buf *strings.Builder) {
			buf.WriteString("&" + "key=this_is_key")
		})
	}
}

func TestVerifyWithSha1(t *testing.T) {
	var h = NewHashSign(crypto.SHA1)
	var form = url.Values{}
	form.Add("jsapi_ticket", "sM4AOVdWfPE4DxkXGEs8VMCPGGVi4C3VM0P37wVUCFvkVAy_90u5h9nbSlYy3-Sl-HhTdfl2fzFy1AOcHKP7qg")
	form.Add("noncestr", "Wm3WZYTPz0wzccnW")
	form.Add("timestamp", "1414587457")
	form.Add("url", "http://mp.weixin.qq.com?params=value")

	if h.Verify(form, "0f9de62fce790f9a083d5c99e95740ceb90c27ed") == false {
		t.Log("sha1 签名错误")
	}
}
