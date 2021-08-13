package huawei

import (
	"net/http"
)

var HuaweiAppIdMap = map[string]string{
	"com.gengcon.android.jxc":            "100582595",
	"com.niimbot.cxprinter":              "103954035",
	"com.gengcon.android.jccloudprinter": "100455843",
	"com.android.jckdcs":                 "101277937",
	"com.android.jcycgj":                 "100895813",
	"com.gengcon.android.fixedassets":    "100408187",
}

// 查询应用包名对应的appid https://developer.huawei.com/consumer/cn/doc/development/AppGallery-connect-References/agcapi-appid-list-0000001111845086
func (c *HuaweiClient) GetAppId(params *GetAppIdParams) (*GetAppIdRes, *http.Response, error) {
	req, err := c.NewRequest(http.MethodGet, "/api/publish/v2/appid-list", params, nil)
	if err != nil {
		return nil, nil, err
	}

	var appIdRes *GetAppIdRes
	resp, err := c.Do(req, &appIdRes)
	if err != nil {
		return nil, resp, err
	}
	return appIdRes, resp, err
}

// 查询应用信息 https://developer.huawei.com/consumer/cn/doc/development/AppGallery-connect-References/agcapi-app-info-query-0000001158365045
func (c *HuaweiClient) GetAppDetail(params *GetAppDetailParams) (*AppDetailRes, *http.Response, error) {
	req, err := c.NewRequest(http.MethodGet, "/api/publish/v2/app-info", params, nil)
	if err != nil {
		return nil, nil, err
	}

	var appDetail *AppDetailRes
	resp, err := c.Do(req, &appDetail)
	if err != nil {
		return nil, resp, err
	}

	return appDetail, resp, err
}

// 通过下载方式提交发布 https://developer.huawei.com/consumer/cn/doc/development/AppGallery-connect-References/agcapi-app-submit-with-file-0000001111845092
func (c *HuaweiClient) SubmitWithDownload(query *AppIdParams, params *SubmitWithDownloadParams) (*SubmitWithDownloadRes, *http.Response, error) {
	req, err := c.NewRequest(http.MethodPost, "/api/publish/v2/app-submit-with-file", query, params)
	if err != nil {
		return nil, nil, err
	}

	var submitWithDownloadRes *SubmitWithDownloadRes
	resp, err := c.Do(req, &submitWithDownloadRes)
	if err != nil {
		return nil, resp, err
	}

	return submitWithDownloadRes, resp, err
}

// 更新语言描述信息 https://developer.huawei.com/consumer/cn/doc/development/AppGallery-connect-References/agcapi-language-info-update-0000001158245057
func (c *HuaweiClient) UpdateLangDescription(quey *AppIdParams, params *UpdateLangDescriptionParams) (*Ret, *http.Response, error) {
	req, err := c.NewRequest(http.MethodPut, "/api/publish/v2/app-language-info", quey, params)
	if err != nil {
		return nil, nil, err
	}

	var ret *Ret
	resp, err := c.Do(req, &ret)
	if err != nil {
		return nil, resp, err
	}

	return ret, resp, err
}
