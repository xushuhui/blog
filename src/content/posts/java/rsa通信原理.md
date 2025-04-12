---
title: "Rsa 通信原理"
published: 2020-07-24T10:23:14+08:00
toc: true
---

## 原理

- A 和 B 通信，双方生成各自的公钥和私钥，把各自的公钥给对方。
- A 向 B 发送消息，A 使用 B 的公钥加密消息，发给 B。
- B 收到消息，用 B 的私钥解密消息。
- B 回复消息，用 A 的公钥加密消息，发送给 A。
- A 收到消息，用 A 的私钥解密消息

```sh
curl 'https://www.douyin.com/aweme/v1/web/channel/feed/?device_platform=webapp&aid=6383&channel=channel_pc_web&tag_id=&count=10&version_code=160100&version_name=16.1.0&_signature=_02B4Z6wo00f01JqNG5QAAIDBG-N4jKVOoLCajB8AAEZ3aa' 
-H 'User-Agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:89.0) Gecko/20100101 Firefox/89.0' 
-H 'Accept: application/json, text/plain, */*' 
-H 'Accept-Language: zh-CN,zh;q=0.8,zh-TW;q=0.7,zh-HK;q=0.5,en-US;q=0.3,en;q=0.2' 
--compressed -H 'withCredentials: true' 
-H 'Connection: keep-alive' 
-H 'Referer: https://www.douyin.com/' 
-H 'Cookie: ttwid=1%7CBiSfQKa4mBHj8ILfoXh4y0G_4pM1GRNRBOrMieRZr9c%7C1624347048%7Cdd6a2e488596fa1dc7a4de1691f5a58f9e9d8dc1e9c4303ed844235f11196ff9; douyin.com; MONITOR_WEB_ID=222cf879-7019-4458-bcc2-6a45c00364cf; s_v_web_id=verify_kq7qant4_Z8G4Jhqk_708o_4g5m_9Qjt_3HmGI4BX2QjZ; passport_csrf_token_default=2f4a83eea195e8eb919354b1f10e4946; passport_csrf_token=2f4a83eea195e8eb919354b1f10e4946' 
-H 'TE: Trailers'
```
```sh
curl 'https://www.douyin.com/aweme/v1/web/channel/feed/?device_platform=webapp&aid=6383&channel=channel_pc_web&tag_id=&count=10&version_code=160100&version_name=16.1.0&_signature=_02B4Z6wo00f01Nppj2wAAIDBWwfsdtWK5WzaaIvAAFZRd0' 
-H 'User-Agent: Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:89.0) Gecko/20100101 Firefox/89.0' 
-H 'Accept: application/json, text/plain, */*' 
-H 'Accept-Language: zh-CN,zh;q=0.8,zh-TW;q=0.7,zh-HK;q=0.5,en-US;q=0.3,en;q=0.2' 
--compressed -H 'withCredentials: true' 
-H 'Connection: keep-alive' 
-H 'Referer: https://www.douyin.com/' -H 'Cookie: ttwid=1%7CBiSfQKa4mBHj8ILfoXh4y0G_4pM1GRNRBOrMieRZr9c%7C1624347048%7Cdd6a2e488596fa1dc7a4de1691f5a58f9e9d8dc1e9c4303ed844235f11196ff9; douyin.com; MONITOR_WEB_ID=222cf879-7019-4458-bcc2-6a45c00364cf; s_v_web_id=verify_kq7qant4_Z8G4Jhqk_708o_4g5m_9Qjt_3HmGI4BX2QjZ' 
-H 'TE: Trailers'
```