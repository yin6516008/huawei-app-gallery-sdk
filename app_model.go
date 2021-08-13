package huawei

type AppIdParams struct {
	AppId string `json:"appId"`
}

type GetAppDetailParams struct {
	AppID       string `url:"appId" json:"appId"`
	Lang        string `url:"lang,omitempty" json:"lang,omitempty"`
	ReleaseType int    `url:"releaseType,omitempty" json:"releaseType,omitempty"`
}

type GetAppIdParams struct {
	PackageName string `url:"packageName" json:"packageName"`
}

// 获取AppId
type GetAppIdRes struct {
	Ret    Ret      `json:"ret"`
	AppIds []AppIds `json:"appids"`
}

type AppIds struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type AppInfo struct {
	ReleaseState      int64        `json:"releaseState"`
	DefaultLang       string       `json:"defaultLang"`
	ParentType        int64        `json:"parentType"`
	ChildType         int64        `json:"childType"`
	GrandChildType    int64        `json:"grandChildType"`
	PrivacyPolicy     string       `json:"privacyPolicy"`
	AppAdapters       string       `json:"appAdapters"`
	IsFree            int64        `json:"isFree"`
	Price             string       `json:"price"`
	PriceDetail       string       `json:"priceDetail"`
	PublishCountry    string       `json:"publishCountry"`
	ContentRate       string       `json:"contentRate"`
	HispaceAutoDown   int64        `json:"hispaceAutoDown"`
	AppTariffType     string       `json:"appTariffType"`
	DeveloperNameCN   string       `json:"developerNameCn"`
	DeveloperNameEn   string       `json:"developerNameEn"`
	DeveloperAddr     string       `json:"developerAddr"`
	DeveloperEmail    string       `json:"developerEmail"`
	DeveloperPhone    string       `json:"developerPhone"`
	DeveloperWebsite  string       `json:"developerWebsite"`
	CertificateURLs   string       `json:"certificateURLs"`
	PublicationURLs   string       `json:"publicationURLs"`
	CultureRecordURLs string       `json:"cultureRecordURLs"`
	UpdateTime        string       `json:"updateTime"`
	VersionNumber     string       `json:"versionNumber"`
	FamilyShareTag    int64        `json:"familyShareTag"`
	DeviceTypes       []DeviceType `json:"deviceTypes"`
}

type Language struct {
	Lang            string           `json:"lang"`
	AppName         string           `json:"appName"`
	AppDesc         string           `json:"appDesc"`
	BriefInfo       string           `json:"briefInfo"`
	NewFeatures     string           `json:"newFeatures"`
	Icon            string           `json:"icon"`
	ShowType        int64            `json:"showType"`
	VideoShowType   int64            `json:"videoShowType"`
	IntroPic        string           `json:"introPic"`
	DeviceMaterials []DeviceMaterial `json:"deviceMaterials"`
	RcmdPic         *string          `json:"rcmdPic,omitempty"`
}

type DeviceMaterial struct {
	DeviceType          int64         `json:"deviceType"`
	AppIcon             string        `json:"appIcon"`
	ScreenShots         []string      `json:"screenShots"`
	ShowType            int64         `json:"showType"`
	VRCoverLayeredImage []interface{} `json:"vrCoverLayeredImage"`
	VRRecomGraphic4To3  []interface{} `json:"vrRecomGraphic4to3"`
	VRRecomGraphic1To1  []interface{} `json:"vrRecomGraphic1to1"`
	PromoGraphics       []interface{} `json:"promoGraphics"`
	VideoShowType       int64         `json:"videoShowType"`
	RcmdPics            *string       `json:"rcmdPics,omitempty"`
}

type DeviceType struct {
	DeviceType  int64  `json:"deviceType"`
	AppAdapters string `json:"appAdapters"`
}

type AppAuditInfo struct {
	AuditOpinion string `json:"auditOpinion"`
}

type AppDetailRes struct {
	Ret       Ret          `json:"ret"`
	AppInfo   AppInfo      `json:"appInfo"`
	AuditInfo AppAuditInfo `json:"auditInfo"`
	Languages []Language   `json:"languages"`
}

type SubmitWithDownloadParams struct {
	DownloadURL      string `json:"downloadUrl"`
	DownloadFileName string `json:"downloadFileName"`
	RequestId        string `json:"requestId"`
}

type SubmitWithDownloadRes struct {
	Ret Ret `json:"ret"`
}

type UpdateLangDescriptionParams struct {
	Lang        string `json:"lang"`
	NewFeatures string `json:"newFeatures"`
}
