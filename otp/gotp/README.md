### GoogleAuthenticator 工具类

这是一个用于生成和验证 Google Authenticator 二维码和验证码的工具类。

#### 使用方法

首先，你需要创建一个 GoogleAuthenticator 实例：

```golang
GA := NewGoogleAuthenticator()
```

你可以使用以下方法来配置 GoogleAuthenticator：
- WithSecretSize(secretSize int)：设置密钥的大小。
- WithExpireSecond(expireSecond int)：设置验证码的过期时间（以秒为单位）。
- WithDigits(digits int)：设置验证码的位数。
- WithQrApi(qrApi string)：设置生成二维码的 API。
例如：

```golang
GA := NewGoogleAuthenticator(WithSecretSize(16), WithExpireSecond(30), WithDigits(6))
```

然后，你可以使用以下方法：
- GenerateSecret()：生成一个新的密钥。
- GenerateCode(secret string)：使用给定的密钥生成一个新的验证码。
- GenerateQRCode(label string, issuer string, secret string)：生成一个新的二维码（返回二维码的内容）。
- GenerateQRUrl(label string, issuer string, secret string)：生成一个新的二维码（返回二维码的 URL）。
- ValidateCode(secret, code string)：验证给定的验证码是否有效。
例如:

```golang
secret := GA.GenerateSecret()
code, err := GA.GenerateCode(secret)
isValid, err := GA.ValidateCode(secret, code)
```

#### 注意事项
- GenerateSecret() 方法生成的密钥是随机的，每次调用都会返回一个新的密钥。
- GenerateCode(secret string) 方法生成的验证码是基于当前的 Unix 时间戳和给定的密钥。因此，同一个密钥在同一时间段内（由 ExpireSecond 属性决定）生成的验证码是相同的。
- ValidateCode(secret, code string) 方法会根据当前的 Unix 时间戳和给定的密钥生成一个新的验证码，然后将其与给定的验证码进行比较。如果两者相同，那么给定的验证码就是有效的。