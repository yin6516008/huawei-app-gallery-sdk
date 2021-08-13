package huawei

import (
	"fmt"
	"testing"
)

func TestGetAppDetail(t *testing.T) {
	client, err := NewHuaweiClientWithEnv()
	if err != nil {
		t.Log(err)
	}
	params := &GetAppDetailParams{
		AppID: "100582595",
	}
	app, _, err := client.GetAppDetail(params)
	if err != nil {
		t.Log(err)
	}
	fmt.Println(app.Ret.Code)
}

func TestGetAppId(t *testing.T) {
	client, err := NewHuaweiClientWithEnv()
	if err != nil {
		t.Log(err)
	}
	params := &GetAppIdParams{
		PackageName: "com.gengcon.android.jxc",
	}
	appIdsRes, _, err := client.GetAppId(params)
	if err != nil {
		t.Log(err)
	}
	fmt.Println(appIdsRes.AppIds)
}

func TestSubmitWithDownload(t *testing.T) {
	client, err := NewHuaweiClientWithEnv()
	if err != nil {
		t.Log(err)
	}

	params := &SubmitWithDownloadParams{
		DownloadFileName: "com.gengcon.android.bsydy_1.1.0_4.apk",
		DownloadURL:      "https://oss.niimbot.com/pugna/com.gengcon.android.bsydy_1.1.0_4.apk",
		RequestId:        client.clientID,
	}

	query := &AppIdParams{
		AppId: "101650277",
	}
	res, _, err := client.SubmitWithDownload(query, params)
	if err != nil {
		t.Log(err)
	}
	fmt.Println(res)
}

func TestUpdateLangDescription(t *testing.T) {
	client, err := NewHuaweiClientWithEnv()
	if err != nil {
		t.Log(err)
	}

	params := &UpdateLangDescriptionParams{
		Lang:        "zh-CN",
		NewFeatures: "更新一个心的版本",
	}

	query := &AppIdParams{
		AppId: "101650277",
	}

	res, _, err := client.UpdateLangDescription(query, params)
	if err != nil {
		t.Log(err)
	}
	fmt.Println(res)
}
