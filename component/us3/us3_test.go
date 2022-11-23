package us3

var (
	us3Client = &Us3{
		PublicKey:  "TOKEN_8e7f7467-4b29-4310-8d4d-796076d02322",
		PrivateKey: "2f2e7896-d798-47a5-afd5-cedece95d777",
		// PublicKey:  "v3ePASSjFw4+P9z0WcFkQxnKCdnJ40YiZRl8BzbI7+LEBVBk6e6CzQ==",
		// PrivateKey: "sBE0A6FuMDJ9N2/o4V6cQvHIWVKnknpfepOfTGXqRGwkbMtPg+QQtfpmXoEPbWPZ",
		Region:     "cn-gd",
		BucketName: "boshang",
		Type:       "public",
	}
)

// func Test_PrefixFileList(t *testing.T) {
// 	v, err := us3Client.PrefixFileList(context.Background(), "", "", 0)
// 	if err != nil {
// 		t.Error(err)
// 		return
// 	}
// 	t.Log(v)
// }

// func Test_PutFile(t *testing.T) {
// 	f, _ := os.Open("C:\\Users\\huanghu\\Pictures\\FluAY907uESYOjg3OAnw9Gu2hTre.jpg")
// 	err := us3Client.PutFile(context.Background(), "upload/image/20220903/1.txt", f)
// 	if err != nil {
// 		t.Error(err)
// 		return
// 	}
// }

// func Test_PostFile(t *testing.T) {
// 	f, _ := os.Open("C:\\Users\\huanghu\\Pictures\\FluAY907uESYOjg3OAnw9Gu2hTre.jpg")
// 	err := us3Client.PostFile(context.Background(), "upload/image/20220903/2.txt", f)
// 	if err != nil {
// 		t.Error(err)
// 		return
// 	}
// }
