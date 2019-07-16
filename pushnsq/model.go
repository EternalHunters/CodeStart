package main

import "github.com/golang/protobuf/proto"

type Condition struct {
	Filter          string   `json:"filter"`
	Relation        string   `json:"relation"`
	ProductIds      []int    `json:"product_ids,omitempty"`
	Regions         []Region `json:"regions,omitempty"`
	EnterpriseIds   []int    `json:"enterprise_ids,omitempty"`
	HomeTags        []int    `json:"home_tags,omitempty"`
	UserTags        []int    `json:"user_tags,omitempty"`
	Platforms       []string `json:"platforms,omitempty"`
	OrderTime       []string `json:"order_time,omitempty"`
	OrderCommoditys []int    `json:"order_commodity,omitempty"`
	AppVersions     App      `json:"app_versions,omitempty"`
	DevelopIdents   []string `json:"develop_idents,omitempty"`
	ClientTypes     []string `json:"client_types,omitempty"`
	Phones          []string `json:"phones,omitempty"`
	UserIds         []int64  `json:"user_ids,omitempty"`
}
type Region struct {
	ProvinceId int64 `json:"province_id"`
	CityId     int64 `json:"city_id,omitempty"`
	CountyId   int64 `json:"county_id,omitempty"`
}
type App struct {
	Version      []string `json:"version"`
	ClientType   string   `json:"client_type,omitempty"`
	Platform     string   `json:"platform"`
	DevelopIdent string   `json:"develop_ident,omitempty"`
}
type Header struct {
	Id               *int32  `protobuf:"varint,1,req,name=id" json:"id,omitempty"`
	Type             *int32  `protobuf:"varint,2,req,name=type,enum=gpb.Type,def=0" json:"type,omitempty"`
	ServiceName      *string `protobuf:"bytes,3,req,name=service_name" json:"service_name,omitempty"`
	From             *string `protobuf:"bytes,4,req,name=from" json:"from,omitempty"`
	FromType         *int32  `protobuf:"varint,5,req,name=from_type,enum=gpb.ObjectType" json:"from_type,omitempty"`
	To               *string `protobuf:"bytes,6,req,name=to" json:"to,omitempty"`
	ToType           *int32  `protobuf:"varint,7,req,name=to_type,enum=gpb.ObjectType" json:"to_type,omitempty"`
	SessionId        *string `protobuf:"bytes,8,opt,name=session_id" json:"session_id,omitempty"`
	ContentType      *int32  `protobuf:"varint,9,opt,name=content_type,enum=gpb.ContentType,def=0" json:"content_type,omitempty"`
	Version          *string `protobuf:"bytes,10,opt,name=version" json:"version,omitempty"`
	CreateTime       *int64  `protobuf:"varint,11,req,name=create_time" json:"create_time,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}
type Message struct {
	Header           *Header `protobuf:"bytes,1,req,name=header" json:"header,omitempty"`
	Body             []byte  `protobuf:"bytes,2,opt,name=body" json:"body,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Header) Reset()          { *m = Header{} }
func (m *Header) String() string  { return proto.CompactTextString(m) }
func (*Header) ProtoMessage()     {}
func (m *Message) Reset()         { *m = Message{} }
func (m *Message) String() string { return proto.CompactTextString(m) }
func (*Message) ProtoMessage()    {}

type TaskObject struct {
	Phone        string `json:"phone,omitempty"`
	UserId       int64  `json:"user_id,omitempty"`
	ClientType   string `json:"client_type,omitempty"`
	Platform     string `json:"platform,omitempty"`
	DevelopIdent string `json:"develop_ident,omitempty"`
	ClientId     string `json:"-"` //不解析，客户端不传此值
}
type TaskModel struct {
	Id          int                    `json:"message_id, omitempty"`  //来自oss的消息id
	Title       string                 `json:"title, omitempty"`       //消息标题
	Body        string                 `json:"body,omitempty"`         //消息体
	Extend      map[string]interface{} `json:"extend,omitempty"`       //notice的扩展字段
	Type        string                 `json:"type,omitempty"`         //条件类型，所有、指定、导入
	Condition   []Condition            `json:"condition,omitempty"`    //条件
	Objects     []TaskObject           `json:"objects,omitempty"`      //推送对象列表
	Go          string                 `json:"go,omitempty"`           //点击行为
	PushType    string                 `json:"push_type,omitempty"`    //推送类型，实时或者定时
	Timing      string                 `json:"timing,omitempty"`       //定时
	MessageType string                 `json:"message_type,omitempty"` //短信或者消息  notice 、sms
	SmsType     int                    `json:"sms_type,omitempty"`     //短信的类型
	Channel     string                 `json:"channel,omitempty"`      //短信通道
	ExpireTime  int                    `json:"expire_time,omitempty"`  //过期时间
	FilePath    string                 `json:"file_path,omitempty"`    //文件路径
	Callback    string                 `json:"callback,omitempty"`     //回调地址
	Action      string                 `json:"action,omitempty"`
}

type Address struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}
type EmailModel struct {
	Sender    string    `json:"sender,omitempty"`
	Subject   string    `json:"subject,omitempty"`
	Body      string    `json:"body,omitempty"`
	IsHtml    int       `json:"is_html,omitempty"`
	Addresses []Address `json:"addresses,omitempty"`
	Replys    []Address `json:"replys,omitempty"`
}
type AppFaultModel struct {
	UserId     int    `json:"user_id,omitempty"`
	AppId      string `json:"app_id,omitempty"`
	HomeId     int    `json:"home_id,omitempty"`
	DeviceId   string `json:"device_id,omitempty"`
	Code       string `json:"code,omitempty"`
	Type       string `json:"type,omitempty"`
	Remark     string `json:"remark,omitempty"`
	File       string `json:"file,omitempty"`
	ReportTime string `json:"report_time,omitempty"`
}

func NewSmsBody() TaskModel {
	task := TaskModel{}
	task.Id = 309
	task.Title = "test_sms"
	task.Body = "test sms send"
	task.Type = "condition"
	task.PushType = "realtime"
	//task.Timing = "2018-09-12 13:40:15"
	task.MessageType = "sms"
	task.SmsType = 0
	task.Channel = "smart_lock"
	task.ExpireTime = 0
	task.Callback = "/v1/message/message/callback"
	task.Action = ""

	condition := Condition{Filter: "client", Relation: "contain",
		AppVersions: App{Platform: ""}, ClientTypes: []string{"phone"}}
	condition2 := Condition{Filter: "develop_ident", Relation: "contain",
		DevelopIdents: []string{"spotmau"}}
	task.Condition = []Condition{condition, condition2}
	task.Extend = make(map[string]interface{})
	task.Extend["msg_type"] = "sys"
	task.Go = "none"

	return task
}

func NewAppBody() TaskModel {
	task := TaskModel{}
	task.Id = 310
	task.Title = "test_app"
	task.Body = "test app send"
	task.Type = "condition"
	//task.PushType = "timing"
	task.Timing = "2018-09-18 13:40:15"
	task.MessageType = "notice"
	//task.SmsType = 0
	//task.Channel = "smart_lock"
	task.ExpireTime = 0
	task.Callback = "/v1/message/message/callback"
	task.Action = ""

	//	condition := Condition{Filter: "client", Relation: "contain",
	//		AppVersions: App{Platform: ""}, ClientTypes: []string{"phone"}}
	condition2 := Condition{Filter: "app_version", Relation: "contain",
		AppVersions: App{Platform: "ios",
			DevelopIdent: "spotmau",
			ClientType:   "phone",
			Version:      []string{"2.01.04"}}}
	task.Condition = []Condition{condition2}
	task.Extend = make(map[string]interface{})
	task.Extend["msg_type"] = "sys"
	task.Extend["develop_ident"] = "spotmau"
	task.Go = "none"

	return task
}

func NewEnterpriseAppBody() TaskModel {
	task := TaskModel{}
	task.Id = 309
	task.Title = "test_app"
	task.Body = "test app send"
	task.Type = "condition"
	task.PushType = "realtime"
	//task.Timing = "2018-09-12 13:40:15"
	task.MessageType = "notice"
	//task.SmsType = 0
	//task.Channel = "smart_lock"
	task.ExpireTime = 0
	task.Callback = "/v1/message/message/callback"
	task.Action = ""

	condition := Condition{Filter: "client", Relation: "contain",
		AppVersions: App{Platform: ""}, ClientTypes: []string{"phone"}}
	condition2 := Condition{Filter: "develop_ident", Relation: "contain",
		DevelopIdents: []string{"spotmau"}}
	task.Condition = []Condition{condition, condition2}
	task.Extend = make(map[string]interface{})
	task.Extend["msg_type"] = "sys"
	task.Extend["develop_ident"] = "spotmau"
	//task.Extend["enterprise_id"] = 76
	task.Go = "none"

	return task
}

func NewCustomBody() TaskModel {
	task := TaskModel{}
	task.Id = 309
	task.Title = "test_app"
	task.Body = "test app send"
	task.Type = "condition"
	task.PushType = "realtime"
	//task.Timing = "2018-09-12 13:40:15"
	task.MessageType = "custom"
	//task.SmsType = 0
	//task.Channel = "smart_lock"
	task.ExpireTime = 0
	task.Callback = "/v1/message/message/callback"
	task.Action = "AD_POPUP"

	condition := Condition{Filter: "region", Relation: "contain",
		Regions: []Region{Region{ProvinceId: 1019, CityId: 1019002, CountyId: 1019002002}}}
	condition2 := Condition{Filter: "app_version", Relation: "exclusive",
		AppVersions: App{Platform: "ios",
			DevelopIdent: "smart_lock",
			ClientType:   "phone",
			Version:      []string{"1.01.01"}}}
	task.Condition = []Condition{condition, condition2}
	task.Extend = make(map[string]interface{})
	task.Extend["msg_type"] = "sys"
	task.Extend["develop_ident"] = "spotmau"
	task.Go = "none"

	return task
}

func NewEmailBody() EmailModel {
	task := EmailModel{}
	task.Sender = "小松"
	task.Subject = "测试"
	task.Body = "Hello"
	addr := Address{Email: "xiaoss@wondershare.cn", Name: "xiaoss"}
	task.Addresses = []Address{addr}
	return task
}

func NewAppFaultBody() AppFaultModel {
	task := AppFaultModel{}
	task.UserId = 103
	task.AppId = "8ec1743d0a5073830e4212dd72308165"
	task.HomeId = 103
	//task.DeviceId = "3119948908c73de357317fe567628cd8"
	task.Code = "denglu-chaoshi"
	task.Type = "passport"
	task.Remark = "remark"
	task.File = "/aa/aaaa/aaa"
	task.ReportTime = "2018-10-18 11:11:11"
	return task
}
