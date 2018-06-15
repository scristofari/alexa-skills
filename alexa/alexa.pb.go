// Code generated by protoc-gen-go. DO NOT EDIT.
// source: alexa.proto

package alexa

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type AudioPlayer_PlayerActivity int32

const (
	AudioPlayer_IDLE            AudioPlayer_PlayerActivity = 0
	AudioPlayer_PAUSED          AudioPlayer_PlayerActivity = 1
	AudioPlayer_PLAYING         AudioPlayer_PlayerActivity = 2
	AudioPlayer_BUFFER_UNDERRUN AudioPlayer_PlayerActivity = 3
	AudioPlayer_FINISHED        AudioPlayer_PlayerActivity = 4
	AudioPlayer_STOPPED         AudioPlayer_PlayerActivity = 5
)

var AudioPlayer_PlayerActivity_name = map[int32]string{
	0: "IDLE",
	1: "PAUSED",
	2: "PLAYING",
	3: "BUFFER_UNDERRUN",
	4: "FINISHED",
	5: "STOPPED",
}
var AudioPlayer_PlayerActivity_value = map[string]int32{
	"IDLE":            0,
	"PAUSED":          1,
	"PLAYING":         2,
	"BUFFER_UNDERRUN": 3,
	"FINISHED":        4,
	"STOPPED":         5,
}

func (x AudioPlayer_PlayerActivity) String() string {
	return proto.EnumName(AudioPlayer_PlayerActivity_name, int32(x))
}
func (AudioPlayer_PlayerActivity) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_alexa_5f3d312e1410aa40, []int{12, 0}
}

type AlexaRequest struct {
	Version              string   `protobuf:"bytes,1,opt,name=version" json:"version,omitempty"`
	Session              *Session `protobuf:"bytes,2,opt,name=session" json:"session,omitempty"`
	Context              *Context `protobuf:"bytes,3,opt,name=context" json:"context,omitempty"`
	Request              *Request `protobuf:"bytes,4,opt,name=request" json:"request,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *AlexaRequest) Reset()         { *m = AlexaRequest{} }
func (m *AlexaRequest) String() string { return proto.CompactTextString(m) }
func (*AlexaRequest) ProtoMessage()    {}
func (*AlexaRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_alexa_5f3d312e1410aa40, []int{0}
}
func (m *AlexaRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AlexaRequest.Unmarshal(m, b)
}
func (m *AlexaRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AlexaRequest.Marshal(b, m, deterministic)
}
func (dst *AlexaRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AlexaRequest.Merge(dst, src)
}
func (m *AlexaRequest) XXX_Size() int {
	return xxx_messageInfo_AlexaRequest.Size(m)
}
func (m *AlexaRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_AlexaRequest.DiscardUnknown(m)
}

var xxx_messageInfo_AlexaRequest proto.InternalMessageInfo

func (m *AlexaRequest) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *AlexaRequest) GetSession() *Session {
	if m != nil {
		return m.Session
	}
	return nil
}

func (m *AlexaRequest) GetContext() *Context {
	if m != nil {
		return m.Context
	}
	return nil
}

func (m *AlexaRequest) GetRequest() *Request {
	if m != nil {
		return m.Request
	}
	return nil
}

type AlexaResponse struct {
	Version              string             `protobuf:"bytes,1,opt,name=version" json:"version,omitempty"`
	SessionAttributes    *SessionAttributes `protobuf:"bytes,2,opt,name=sessionAttributes" json:"sessionAttributes,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *AlexaResponse) Reset()         { *m = AlexaResponse{} }
func (m *AlexaResponse) String() string { return proto.CompactTextString(m) }
func (*AlexaResponse) ProtoMessage()    {}
func (*AlexaResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_alexa_5f3d312e1410aa40, []int{1}
}
func (m *AlexaResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AlexaResponse.Unmarshal(m, b)
}
func (m *AlexaResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AlexaResponse.Marshal(b, m, deterministic)
}
func (dst *AlexaResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AlexaResponse.Merge(dst, src)
}
func (m *AlexaResponse) XXX_Size() int {
	return xxx_messageInfo_AlexaResponse.Size(m)
}
func (m *AlexaResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_AlexaResponse.DiscardUnknown(m)
}

var xxx_messageInfo_AlexaResponse proto.InternalMessageInfo

func (m *AlexaResponse) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *AlexaResponse) GetSessionAttributes() *SessionAttributes {
	if m != nil {
		return m.SessionAttributes
	}
	return nil
}

type Session struct {
	New                  bool         `protobuf:"varint,1,opt,name=new" json:"new,omitempty"`
	SessionId            string       `protobuf:"bytes,2,opt,name=sessionId" json:"sessionId,omitempty"`
	Application          *Application `protobuf:"bytes,3,opt,name=Application" json:"Application,omitempty"`
	Attributes           *Attributes  `protobuf:"bytes,4,opt,name=attributes" json:"attributes,omitempty"`
	User                 *User        `protobuf:"bytes,5,opt,name=user" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Session) Reset()         { *m = Session{} }
func (m *Session) String() string { return proto.CompactTextString(m) }
func (*Session) ProtoMessage()    {}
func (*Session) Descriptor() ([]byte, []int) {
	return fileDescriptor_alexa_5f3d312e1410aa40, []int{2}
}
func (m *Session) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Session.Unmarshal(m, b)
}
func (m *Session) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Session.Marshal(b, m, deterministic)
}
func (dst *Session) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Session.Merge(dst, src)
}
func (m *Session) XXX_Size() int {
	return xxx_messageInfo_Session.Size(m)
}
func (m *Session) XXX_DiscardUnknown() {
	xxx_messageInfo_Session.DiscardUnknown(m)
}

var xxx_messageInfo_Session proto.InternalMessageInfo

func (m *Session) GetNew() bool {
	if m != nil {
		return m.New
	}
	return false
}

func (m *Session) GetSessionId() string {
	if m != nil {
		return m.SessionId
	}
	return ""
}

func (m *Session) GetApplication() *Application {
	if m != nil {
		return m.Application
	}
	return nil
}

func (m *Session) GetAttributes() *Attributes {
	if m != nil {
		return m.Attributes
	}
	return nil
}

func (m *Session) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

type SessionAttributes struct {
	Key                  []string `protobuf:"bytes,1,rep,name=key" json:"key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SessionAttributes) Reset()         { *m = SessionAttributes{} }
func (m *SessionAttributes) String() string { return proto.CompactTextString(m) }
func (*SessionAttributes) ProtoMessage()    {}
func (*SessionAttributes) Descriptor() ([]byte, []int) {
	return fileDescriptor_alexa_5f3d312e1410aa40, []int{3}
}
func (m *SessionAttributes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SessionAttributes.Unmarshal(m, b)
}
func (m *SessionAttributes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SessionAttributes.Marshal(b, m, deterministic)
}
func (dst *SessionAttributes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SessionAttributes.Merge(dst, src)
}
func (m *SessionAttributes) XXX_Size() int {
	return xxx_messageInfo_SessionAttributes.Size(m)
}
func (m *SessionAttributes) XXX_DiscardUnknown() {
	xxx_messageInfo_SessionAttributes.DiscardUnknown(m)
}

var xxx_messageInfo_SessionAttributes proto.InternalMessageInfo

func (m *SessionAttributes) GetKey() []string {
	if m != nil {
		return m.Key
	}
	return nil
}

type Context struct {
	System               *System      `protobuf:"bytes,1,opt,name=System" json:"System,omitempty"`
	AudioPlayer          *AudioPlayer `protobuf:"bytes,2,opt,name=AudioPlayer" json:"AudioPlayer,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Context) Reset()         { *m = Context{} }
func (m *Context) String() string { return proto.CompactTextString(m) }
func (*Context) ProtoMessage()    {}
func (*Context) Descriptor() ([]byte, []int) {
	return fileDescriptor_alexa_5f3d312e1410aa40, []int{4}
}
func (m *Context) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Context.Unmarshal(m, b)
}
func (m *Context) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Context.Marshal(b, m, deterministic)
}
func (dst *Context) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Context.Merge(dst, src)
}
func (m *Context) XXX_Size() int {
	return xxx_messageInfo_Context.Size(m)
}
func (m *Context) XXX_DiscardUnknown() {
	xxx_messageInfo_Context.DiscardUnknown(m)
}

var xxx_messageInfo_Context proto.InternalMessageInfo

func (m *Context) GetSystem() *System {
	if m != nil {
		return m.System
	}
	return nil
}

func (m *Context) GetAudioPlayer() *AudioPlayer {
	if m != nil {
		return m.AudioPlayer
	}
	return nil
}

type System struct {
	Device               *Device      `protobuf:"bytes,1,opt,name=device" json:"device,omitempty"`
	Application          *Application `protobuf:"bytes,2,opt,name=application" json:"application,omitempty"`
	User                 *User        `protobuf:"bytes,3,opt,name=user" json:"user,omitempty"`
	ApiEndpoint          string       `protobuf:"bytes,4,opt,name=apiEndpoint" json:"apiEndpoint,omitempty"`
	ApiAccessToken       string       `protobuf:"bytes,5,opt,name=apiAccessToken" json:"apiAccessToken,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *System) Reset()         { *m = System{} }
func (m *System) String() string { return proto.CompactTextString(m) }
func (*System) ProtoMessage()    {}
func (*System) Descriptor() ([]byte, []int) {
	return fileDescriptor_alexa_5f3d312e1410aa40, []int{5}
}
func (m *System) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_System.Unmarshal(m, b)
}
func (m *System) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_System.Marshal(b, m, deterministic)
}
func (dst *System) XXX_Merge(src proto.Message) {
	xxx_messageInfo_System.Merge(dst, src)
}
func (m *System) XXX_Size() int {
	return xxx_messageInfo_System.Size(m)
}
func (m *System) XXX_DiscardUnknown() {
	xxx_messageInfo_System.DiscardUnknown(m)
}

var xxx_messageInfo_System proto.InternalMessageInfo

func (m *System) GetDevice() *Device {
	if m != nil {
		return m.Device
	}
	return nil
}

func (m *System) GetApplication() *Application {
	if m != nil {
		return m.Application
	}
	return nil
}

func (m *System) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *System) GetApiEndpoint() string {
	if m != nil {
		return m.ApiEndpoint
	}
	return ""
}

func (m *System) GetApiAccessToken() string {
	if m != nil {
		return m.ApiAccessToken
	}
	return ""
}

type Device struct {
	DeviceId             string               `protobuf:"bytes,1,opt,name=deviceId" json:"deviceId,omitempty"`
	SupportedInterfaces  *SupportedInterfaces `protobuf:"bytes,2,opt,name=supportedInterfaces" json:"supportedInterfaces,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Device) Reset()         { *m = Device{} }
func (m *Device) String() string { return proto.CompactTextString(m) }
func (*Device) ProtoMessage()    {}
func (*Device) Descriptor() ([]byte, []int) {
	return fileDescriptor_alexa_5f3d312e1410aa40, []int{6}
}
func (m *Device) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Device.Unmarshal(m, b)
}
func (m *Device) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Device.Marshal(b, m, deterministic)
}
func (dst *Device) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Device.Merge(dst, src)
}
func (m *Device) XXX_Size() int {
	return xxx_messageInfo_Device.Size(m)
}
func (m *Device) XXX_DiscardUnknown() {
	xxx_messageInfo_Device.DiscardUnknown(m)
}

var xxx_messageInfo_Device proto.InternalMessageInfo

func (m *Device) GetDeviceId() string {
	if m != nil {
		return m.DeviceId
	}
	return ""
}

func (m *Device) GetSupportedInterfaces() *SupportedInterfaces {
	if m != nil {
		return m.SupportedInterfaces
	}
	return nil
}

type SupportedInterfaces struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SupportedInterfaces) Reset()         { *m = SupportedInterfaces{} }
func (m *SupportedInterfaces) String() string { return proto.CompactTextString(m) }
func (*SupportedInterfaces) ProtoMessage()    {}
func (*SupportedInterfaces) Descriptor() ([]byte, []int) {
	return fileDescriptor_alexa_5f3d312e1410aa40, []int{7}
}
func (m *SupportedInterfaces) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SupportedInterfaces.Unmarshal(m, b)
}
func (m *SupportedInterfaces) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SupportedInterfaces.Marshal(b, m, deterministic)
}
func (dst *SupportedInterfaces) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SupportedInterfaces.Merge(dst, src)
}
func (m *SupportedInterfaces) XXX_Size() int {
	return xxx_messageInfo_SupportedInterfaces.Size(m)
}
func (m *SupportedInterfaces) XXX_DiscardUnknown() {
	xxx_messageInfo_SupportedInterfaces.DiscardUnknown(m)
}

var xxx_messageInfo_SupportedInterfaces proto.InternalMessageInfo

type Application struct {
	ApplicationId        string   `protobuf:"bytes,1,opt,name=applicationId" json:"applicationId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Application) Reset()         { *m = Application{} }
func (m *Application) String() string { return proto.CompactTextString(m) }
func (*Application) ProtoMessage()    {}
func (*Application) Descriptor() ([]byte, []int) {
	return fileDescriptor_alexa_5f3d312e1410aa40, []int{8}
}
func (m *Application) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Application.Unmarshal(m, b)
}
func (m *Application) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Application.Marshal(b, m, deterministic)
}
func (dst *Application) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Application.Merge(dst, src)
}
func (m *Application) XXX_Size() int {
	return xxx_messageInfo_Application.Size(m)
}
func (m *Application) XXX_DiscardUnknown() {
	xxx_messageInfo_Application.DiscardUnknown(m)
}

var xxx_messageInfo_Application proto.InternalMessageInfo

func (m *Application) GetApplicationId() string {
	if m != nil {
		return m.ApplicationId
	}
	return ""
}

type Attributes struct {
	Key                  []string `protobuf:"bytes,1,rep,name=key" json:"key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Attributes) Reset()         { *m = Attributes{} }
func (m *Attributes) String() string { return proto.CompactTextString(m) }
func (*Attributes) ProtoMessage()    {}
func (*Attributes) Descriptor() ([]byte, []int) {
	return fileDescriptor_alexa_5f3d312e1410aa40, []int{9}
}
func (m *Attributes) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Attributes.Unmarshal(m, b)
}
func (m *Attributes) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Attributes.Marshal(b, m, deterministic)
}
func (dst *Attributes) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Attributes.Merge(dst, src)
}
func (m *Attributes) XXX_Size() int {
	return xxx_messageInfo_Attributes.Size(m)
}
func (m *Attributes) XXX_DiscardUnknown() {
	xxx_messageInfo_Attributes.DiscardUnknown(m)
}

var xxx_messageInfo_Attributes proto.InternalMessageInfo

func (m *Attributes) GetKey() []string {
	if m != nil {
		return m.Key
	}
	return nil
}

type User struct {
	UserId               string       `protobuf:"bytes,1,opt,name=userId" json:"userId,omitempty"`
	AccessToken          string       `protobuf:"bytes,2,opt,name=accessToken" json:"accessToken,omitempty"`
	Permissions          *Permissions `protobuf:"bytes,3,opt,name=Permissions" json:"Permissions,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_alexa_5f3d312e1410aa40, []int{10}
}
func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (dst *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(dst, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *User) GetAccessToken() string {
	if m != nil {
		return m.AccessToken
	}
	return ""
}

func (m *User) GetPermissions() *Permissions {
	if m != nil {
		return m.Permissions
	}
	return nil
}

type Permissions struct {
	ConsentToken         string   `protobuf:"bytes,1,opt,name=consentToken" json:"consentToken,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Permissions) Reset()         { *m = Permissions{} }
func (m *Permissions) String() string { return proto.CompactTextString(m) }
func (*Permissions) ProtoMessage()    {}
func (*Permissions) Descriptor() ([]byte, []int) {
	return fileDescriptor_alexa_5f3d312e1410aa40, []int{11}
}
func (m *Permissions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Permissions.Unmarshal(m, b)
}
func (m *Permissions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Permissions.Marshal(b, m, deterministic)
}
func (dst *Permissions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Permissions.Merge(dst, src)
}
func (m *Permissions) XXX_Size() int {
	return xxx_messageInfo_Permissions.Size(m)
}
func (m *Permissions) XXX_DiscardUnknown() {
	xxx_messageInfo_Permissions.DiscardUnknown(m)
}

var xxx_messageInfo_Permissions proto.InternalMessageInfo

func (m *Permissions) GetConsentToken() string {
	if m != nil {
		return m.ConsentToken
	}
	return ""
}

type AudioPlayer struct {
	PlayerActivity       AudioPlayer_PlayerActivity `protobuf:"varint,1,opt,name=playerActivity,enum=alexa.AudioPlayer_PlayerActivity" json:"playerActivity,omitempty"`
	Token                string                     `protobuf:"bytes,2,opt,name=token" json:"token,omitempty"`
	OffsetInMilliseconds int64                      `protobuf:"varint,3,opt,name=offsetInMilliseconds" json:"offsetInMilliseconds,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *AudioPlayer) Reset()         { *m = AudioPlayer{} }
func (m *AudioPlayer) String() string { return proto.CompactTextString(m) }
func (*AudioPlayer) ProtoMessage()    {}
func (*AudioPlayer) Descriptor() ([]byte, []int) {
	return fileDescriptor_alexa_5f3d312e1410aa40, []int{12}
}
func (m *AudioPlayer) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AudioPlayer.Unmarshal(m, b)
}
func (m *AudioPlayer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AudioPlayer.Marshal(b, m, deterministic)
}
func (dst *AudioPlayer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AudioPlayer.Merge(dst, src)
}
func (m *AudioPlayer) XXX_Size() int {
	return xxx_messageInfo_AudioPlayer.Size(m)
}
func (m *AudioPlayer) XXX_DiscardUnknown() {
	xxx_messageInfo_AudioPlayer.DiscardUnknown(m)
}

var xxx_messageInfo_AudioPlayer proto.InternalMessageInfo

func (m *AudioPlayer) GetPlayerActivity() AudioPlayer_PlayerActivity {
	if m != nil {
		return m.PlayerActivity
	}
	return AudioPlayer_IDLE
}

func (m *AudioPlayer) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *AudioPlayer) GetOffsetInMilliseconds() int64 {
	if m != nil {
		return m.OffsetInMilliseconds
	}
	return 0
}

type Request struct {
	Type                 string   `protobuf:"bytes,1,opt,name=type" json:"type,omitempty"`
	RequestId            string   `protobuf:"bytes,2,opt,name=requestId" json:"requestId,omitempty"`
	Timestamp            string   `protobuf:"bytes,3,opt,name=timestamp" json:"timestamp,omitempty"`
	Locale               string   `protobuf:"bytes,4,opt,name=locale" json:"locale,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_alexa_5f3d312e1410aa40, []int{13}
}
func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (dst *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(dst, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Request) GetRequestId() string {
	if m != nil {
		return m.RequestId
	}
	return ""
}

func (m *Request) GetTimestamp() string {
	if m != nil {
		return m.Timestamp
	}
	return ""
}

func (m *Request) GetLocale() string {
	if m != nil {
		return m.Locale
	}
	return ""
}

type Response struct {
	OutputSpeech         *OutputSpeech `protobuf:"bytes,1,opt,name=outputSpeech" json:"outputSpeech,omitempty"`
	Card                 *Card         `protobuf:"bytes,2,opt,name=card" json:"card,omitempty"`
	Reprompt             *OutputSpeech `protobuf:"bytes,3,opt,name=reprompt" json:"reprompt,omitempty"`
	ShouldEndSession     bool          `protobuf:"varint,4,opt,name=shouldEndSession" json:"shouldEndSession,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_alexa_5f3d312e1410aa40, []int{14}
}
func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (dst *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(dst, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetOutputSpeech() *OutputSpeech {
	if m != nil {
		return m.OutputSpeech
	}
	return nil
}

func (m *Response) GetCard() *Card {
	if m != nil {
		return m.Card
	}
	return nil
}

func (m *Response) GetReprompt() *OutputSpeech {
	if m != nil {
		return m.Reprompt
	}
	return nil
}

func (m *Response) GetShouldEndSession() bool {
	if m != nil {
		return m.ShouldEndSession
	}
	return false
}

type OutputSpeech struct {
	Type                 string   `protobuf:"bytes,1,opt,name=type" json:"type,omitempty"`
	Text                 string   `protobuf:"bytes,2,opt,name=text" json:"text,omitempty"`
	Ssml                 string   `protobuf:"bytes,3,opt,name=ssml" json:"ssml,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OutputSpeech) Reset()         { *m = OutputSpeech{} }
func (m *OutputSpeech) String() string { return proto.CompactTextString(m) }
func (*OutputSpeech) ProtoMessage()    {}
func (*OutputSpeech) Descriptor() ([]byte, []int) {
	return fileDescriptor_alexa_5f3d312e1410aa40, []int{15}
}
func (m *OutputSpeech) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OutputSpeech.Unmarshal(m, b)
}
func (m *OutputSpeech) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OutputSpeech.Marshal(b, m, deterministic)
}
func (dst *OutputSpeech) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OutputSpeech.Merge(dst, src)
}
func (m *OutputSpeech) XXX_Size() int {
	return xxx_messageInfo_OutputSpeech.Size(m)
}
func (m *OutputSpeech) XXX_DiscardUnknown() {
	xxx_messageInfo_OutputSpeech.DiscardUnknown(m)
}

var xxx_messageInfo_OutputSpeech proto.InternalMessageInfo

func (m *OutputSpeech) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *OutputSpeech) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

func (m *OutputSpeech) GetSsml() string {
	if m != nil {
		return m.Ssml
	}
	return ""
}

type Card struct {
	Type                 string   `protobuf:"bytes,1,opt,name=type" json:"type,omitempty"`
	Title                string   `protobuf:"bytes,2,opt,name=title" json:"title,omitempty"`
	Content              string   `protobuf:"bytes,3,opt,name=content" json:"content,omitempty"`
	Text                 string   `protobuf:"bytes,4,opt,name=text" json:"text,omitempty"`
	Image                *Image   `protobuf:"bytes,5,opt,name=image" json:"image,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Card) Reset()         { *m = Card{} }
func (m *Card) String() string { return proto.CompactTextString(m) }
func (*Card) ProtoMessage()    {}
func (*Card) Descriptor() ([]byte, []int) {
	return fileDescriptor_alexa_5f3d312e1410aa40, []int{16}
}
func (m *Card) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Card.Unmarshal(m, b)
}
func (m *Card) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Card.Marshal(b, m, deterministic)
}
func (dst *Card) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Card.Merge(dst, src)
}
func (m *Card) XXX_Size() int {
	return xxx_messageInfo_Card.Size(m)
}
func (m *Card) XXX_DiscardUnknown() {
	xxx_messageInfo_Card.DiscardUnknown(m)
}

var xxx_messageInfo_Card proto.InternalMessageInfo

func (m *Card) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Card) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *Card) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func (m *Card) GetText() string {
	if m != nil {
		return m.Text
	}
	return ""
}

func (m *Card) GetImage() *Image {
	if m != nil {
		return m.Image
	}
	return nil
}

type Image struct {
	SmallImageUrl        string   `protobuf:"bytes,1,opt,name=smallImageUrl" json:"smallImageUrl,omitempty"`
	LargeImageUrl        string   `protobuf:"bytes,2,opt,name=largeImageUrl" json:"largeImageUrl,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Image) Reset()         { *m = Image{} }
func (m *Image) String() string { return proto.CompactTextString(m) }
func (*Image) ProtoMessage()    {}
func (*Image) Descriptor() ([]byte, []int) {
	return fileDescriptor_alexa_5f3d312e1410aa40, []int{17}
}
func (m *Image) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Image.Unmarshal(m, b)
}
func (m *Image) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Image.Marshal(b, m, deterministic)
}
func (dst *Image) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Image.Merge(dst, src)
}
func (m *Image) XXX_Size() int {
	return xxx_messageInfo_Image.Size(m)
}
func (m *Image) XXX_DiscardUnknown() {
	xxx_messageInfo_Image.DiscardUnknown(m)
}

var xxx_messageInfo_Image proto.InternalMessageInfo

func (m *Image) GetSmallImageUrl() string {
	if m != nil {
		return m.SmallImageUrl
	}
	return ""
}

func (m *Image) GetLargeImageUrl() string {
	if m != nil {
		return m.LargeImageUrl
	}
	return ""
}

type Directives struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Directives) Reset()         { *m = Directives{} }
func (m *Directives) String() string { return proto.CompactTextString(m) }
func (*Directives) ProtoMessage()    {}
func (*Directives) Descriptor() ([]byte, []int) {
	return fileDescriptor_alexa_5f3d312e1410aa40, []int{18}
}
func (m *Directives) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Directives.Unmarshal(m, b)
}
func (m *Directives) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Directives.Marshal(b, m, deterministic)
}
func (dst *Directives) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Directives.Merge(dst, src)
}
func (m *Directives) XXX_Size() int {
	return xxx_messageInfo_Directives.Size(m)
}
func (m *Directives) XXX_DiscardUnknown() {
	xxx_messageInfo_Directives.DiscardUnknown(m)
}

var xxx_messageInfo_Directives proto.InternalMessageInfo

func init() {
	proto.RegisterType((*AlexaRequest)(nil), "alexa.AlexaRequest")
	proto.RegisterType((*AlexaResponse)(nil), "alexa.AlexaResponse")
	proto.RegisterType((*Session)(nil), "alexa.Session")
	proto.RegisterType((*SessionAttributes)(nil), "alexa.SessionAttributes")
	proto.RegisterType((*Context)(nil), "alexa.Context")
	proto.RegisterType((*System)(nil), "alexa.System")
	proto.RegisterType((*Device)(nil), "alexa.Device")
	proto.RegisterType((*SupportedInterfaces)(nil), "alexa.SupportedInterfaces")
	proto.RegisterType((*Application)(nil), "alexa.Application")
	proto.RegisterType((*Attributes)(nil), "alexa.Attributes")
	proto.RegisterType((*User)(nil), "alexa.User")
	proto.RegisterType((*Permissions)(nil), "alexa.Permissions")
	proto.RegisterType((*AudioPlayer)(nil), "alexa.AudioPlayer")
	proto.RegisterType((*Request)(nil), "alexa.Request")
	proto.RegisterType((*Response)(nil), "alexa.Response")
	proto.RegisterType((*OutputSpeech)(nil), "alexa.OutputSpeech")
	proto.RegisterType((*Card)(nil), "alexa.Card")
	proto.RegisterType((*Image)(nil), "alexa.Image")
	proto.RegisterType((*Directives)(nil), "alexa.Directives")
	proto.RegisterEnum("alexa.AudioPlayer_PlayerActivity", AudioPlayer_PlayerActivity_name, AudioPlayer_PlayerActivity_value)
}

func init() { proto.RegisterFile("alexa.proto", fileDescriptor_alexa_5f3d312e1410aa40) }

var fileDescriptor_alexa_5f3d312e1410aa40 = []byte{
	// 899 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x56, 0xcd, 0x8e, 0x1b, 0x45,
	0x10, 0x66, 0xec, 0xf1, 0x5f, 0xd9, 0x6b, 0xbc, 0xbd, 0x01, 0x59, 0x11, 0x82, 0xa5, 0x45, 0xd0,
	0x8a, 0xc3, 0xa2, 0x6c, 0x90, 0x38, 0x9b, 0xd8, 0x0b, 0x83, 0x96, 0x8d, 0xd5, 0x8e, 0x0f, 0x9c,
	0xd0, 0x64, 0xa6, 0x36, 0x69, 0x65, 0xfe, 0xb6, 0xbb, 0x67, 0x89, 0xaf, 0x5c, 0x78, 0x13, 0xde,
	0x03, 0xf1, 0x02, 0x3c, 0x52, 0xd4, 0x3d, 0x3d, 0x7f, 0xb6, 0x93, 0x5b, 0xd7, 0x57, 0x5f, 0xd7,
	0xcf, 0xe7, 0xea, 0x1a, 0xc3, 0xd8, 0x8f, 0xf0, 0x9d, 0x7f, 0x99, 0x89, 0x54, 0xa5, 0xa4, 0x67,
	0x0c, 0xfa, 0x8f, 0x03, 0x93, 0x85, 0x3e, 0x31, 0xbc, 0xcf, 0x51, 0x2a, 0x32, 0x87, 0xc1, 0x03,
	0x0a, 0xc9, 0xd3, 0x64, 0xee, 0x9c, 0x3b, 0x17, 0x23, 0x56, 0x9a, 0xe4, 0x02, 0x06, 0x12, 0xa5,
	0xf1, 0x74, 0xce, 0x9d, 0x8b, 0xf1, 0xd5, 0xf4, 0xb2, 0x08, 0xb8, 0x29, 0x50, 0x56, 0xba, 0x35,
	0x33, 0x48, 0x13, 0x85, 0xef, 0xd4, 0xbc, 0xdb, 0x62, 0x3e, 0x2f, 0x50, 0x56, 0xba, 0x35, 0x53,
	0x14, 0x89, 0xe7, 0x6e, 0x8b, 0x69, 0xcb, 0x61, 0xa5, 0x9b, 0xde, 0xc3, 0x89, 0xad, 0x53, 0x66,
	0x69, 0x22, 0xf1, 0x23, 0x85, 0x5e, 0xc3, 0xa9, 0xad, 0x64, 0xa1, 0x94, 0xe0, 0xaf, 0x72, 0x85,
	0xd2, 0x96, 0x3c, 0x6f, 0x97, 0x5c, 0xfb, 0xd9, 0xe1, 0x15, 0xfa, 0x9f, 0x03, 0x03, 0x4b, 0x24,
	0x33, 0xe8, 0x26, 0xf8, 0xa7, 0xc9, 0x34, 0x64, 0xfa, 0x48, 0xbe, 0x80, 0x91, 0xbd, 0xe2, 0x85,
	0x26, 0xfa, 0x88, 0xd5, 0x00, 0xf9, 0x01, 0xc6, 0x8b, 0x2c, 0x8b, 0x78, 0xe0, 0x2b, 0x5d, 0x61,
	0x21, 0x03, 0xb1, 0xd9, 0x1b, 0x1e, 0xd6, 0xa4, 0x91, 0xa7, 0x00, 0x7e, 0x5d, 0x72, 0xa1, 0xc8,
	0x69, 0x79, 0xa9, 0xae, 0xb5, 0x41, 0x22, 0x5f, 0x81, 0x9b, 0x4b, 0x14, 0xf3, 0x9e, 0x21, 0x8f,
	0x2d, 0x79, 0x2b, 0x51, 0x30, 0xe3, 0xa0, 0x4f, 0xe0, 0xf4, 0xa0, 0x5b, 0xdd, 0xce, 0x5b, 0xdc,
	0xcd, 0x9d, 0xf3, 0xee, 0xc5, 0x88, 0xe9, 0x23, 0xbd, 0x83, 0x81, 0xfd, 0x75, 0xc8, 0x13, 0xe8,
	0x6f, 0x76, 0x52, 0x61, 0x6c, 0xda, 0x1d, 0x5f, 0x9d, 0x94, 0xa2, 0x19, 0x90, 0x59, 0xa7, 0x69,
	0x31, 0x0f, 0x79, 0xba, 0x8e, 0xfc, 0x1d, 0x0a, 0x2b, 0x70, 0xd5, 0x62, 0xed, 0x61, 0x4d, 0x1a,
	0xfd, 0xdf, 0x29, 0xa3, 0xeb, 0x3c, 0x21, 0x3e, 0xf0, 0x00, 0xf7, 0xf2, 0x2c, 0x0d, 0xc8, 0xac,
	0x53, 0xe7, 0xf1, 0x1b, 0x52, 0x76, 0x3e, 0x2c, 0x65, 0x83, 0x56, 0xe9, 0xd2, 0xfd, 0x80, 0x2e,
	0xe4, 0x5c, 0x87, 0xe5, 0xab, 0x24, 0xcc, 0x52, 0x9e, 0x14, 0xe3, 0x37, 0x62, 0x4d, 0x88, 0x7c,
	0x0b, 0x53, 0x3f, 0xe3, 0x8b, 0x20, 0x40, 0x29, 0x5f, 0xa6, 0x6f, 0x31, 0x31, 0x22, 0x8f, 0xd8,
	0x1e, 0x4a, 0x05, 0xf4, 0x8b, 0x92, 0xc9, 0x63, 0x18, 0x16, 0x45, 0x7b, 0xa1, 0x1d, 0xca, 0xca,
	0x26, 0x37, 0x70, 0x26, 0xf3, 0x2c, 0x4b, 0x85, 0xc2, 0xd0, 0x4b, 0x14, 0x8a, 0x3b, 0x3f, 0xa8,
	0xe6, 0xf2, 0x71, 0x29, 0xf1, 0x21, 0x83, 0x1d, 0xbb, 0x46, 0x3f, 0x83, 0xb3, 0x23, 0x5c, 0xfa,
	0xac, 0x35, 0x76, 0xe4, 0x1b, 0x38, 0x69, 0x68, 0x52, 0x15, 0xd5, 0x06, 0xe9, 0x97, 0x00, 0x1f,
	0x1d, 0x8d, 0x07, 0x70, 0xb5, 0x6e, 0xe4, 0x73, 0xe8, 0x6b, 0xe5, 0xaa, 0x30, 0xd6, 0x32, 0x4a,
	0x36, 0x44, 0xea, 0x58, 0x25, 0x6b, 0x48, 0xff, 0x84, 0x6b, 0x14, 0x31, 0x37, 0x63, 0x28, 0xf7,
	0x5e, 0x43, 0xc3, 0xc3, 0x9a, 0x34, 0xfa, 0xb4, 0x75, 0x8b, 0x50, 0x98, 0x04, 0xfa, 0xe5, 0x27,
	0xaa, 0xc8, 0x53, 0x14, 0xd1, 0xc2, 0xe8, 0xdf, 0x9d, 0xd6, 0x50, 0x12, 0x0f, 0xa6, 0x99, 0x39,
	0x2d, 0x02, 0xc5, 0x1f, 0xb8, 0xda, 0x99, 0x5b, 0xd3, 0xab, 0xaf, 0x0f, 0xc7, 0xf4, 0x72, 0xdd,
	0x22, 0xb2, 0xbd, 0x8b, 0xe4, 0x11, 0xf4, 0x54, 0xa3, 0xbf, 0xc2, 0x20, 0x57, 0xf0, 0x28, 0xbd,
	0xbb, 0x93, 0xa8, 0xbc, 0xe4, 0x37, 0x1e, 0x45, 0x5c, 0x62, 0x90, 0x26, 0x61, 0xd1, 0x62, 0x97,
	0x1d, 0xf5, 0xd1, 0x00, 0xa6, 0xed, 0x5c, 0x64, 0x08, 0xae, 0xb7, 0xbc, 0x59, 0xcd, 0x3e, 0x21,
	0x00, 0xfd, 0xf5, 0x62, 0xbb, 0x59, 0x2d, 0x67, 0x0e, 0x19, 0xc3, 0x60, 0x7d, 0xb3, 0xf8, 0xdd,
	0xbb, 0xfd, 0x79, 0xd6, 0x21, 0x67, 0xf0, 0xe9, 0x4f, 0xdb, 0xeb, 0xeb, 0x15, 0xfb, 0x63, 0x7b,
	0xbb, 0x5c, 0x31, 0xb6, 0xbd, 0x9d, 0x75, 0xc9, 0x04, 0x86, 0xd7, 0xde, 0xad, 0xb7, 0xf9, 0x65,
	0xb5, 0x9c, 0xb9, 0x9a, 0xbf, 0x79, 0xf9, 0x62, 0xbd, 0x5e, 0x2d, 0x67, 0x3d, 0x7a, 0x0f, 0x83,
	0x72, 0xa5, 0x13, 0x70, 0xd5, 0x2e, 0x43, 0x2b, 0x98, 0x39, 0xeb, 0xed, 0x65, 0x37, 0x6b, 0xbd,
	0xbd, 0x2a, 0x40, 0x7b, 0x15, 0x8f, 0x51, 0x2a, 0x3f, 0xce, 0x4c, 0x2b, 0x23, 0x56, 0x03, 0x7a,
	0x0e, 0xa2, 0x34, 0xf0, 0x23, 0xb4, 0x8f, 0xc6, 0x5a, 0xf4, 0x5f, 0x07, 0x86, 0xd5, 0x7a, 0xfe,
	0x11, 0x26, 0x69, 0xae, 0xb2, 0x5c, 0x6d, 0x32, 0xc4, 0xe0, 0x8d, 0x7d, 0xe2, 0x67, 0x56, 0xf7,
	0x17, 0x0d, 0x17, 0x6b, 0x11, 0xf5, 0xc3, 0x0d, 0x7c, 0x11, 0xda, 0x87, 0x51, 0x3e, 0xdc, 0xe7,
	0xbe, 0x08, 0x99, 0x71, 0x90, 0xef, 0x61, 0x28, 0x30, 0x13, 0x69, 0x9c, 0x95, 0x9f, 0x97, 0xa3,
	0x51, 0x2b, 0x12, 0xf9, 0x0e, 0x66, 0xf2, 0x4d, 0x9a, 0x47, 0xe1, 0x2a, 0x09, 0xed, 0x2a, 0x34,
	0x95, 0x0f, 0xd9, 0x01, 0x4e, 0x7f, 0x85, 0x49, 0x33, 0xca, 0x51, 0xed, 0x34, 0xa6, 0xbf, 0x6d,
	0x1d, 0x8b, 0xe9, 0x9d, 0x49, 0xc0, 0x95, 0x32, 0x8e, 0xac, 0x58, 0xe6, 0x4c, 0xff, 0x72, 0xc0,
	0xd5, 0x75, 0x1f, 0x0d, 0xa2, 0xc7, 0x89, 0xab, 0x08, 0xab, 0x71, 0xd2, 0x86, 0xfe, 0xa8, 0x99,
	0x4f, 0x63, 0xa2, 0x6c, 0xa4, 0xd2, 0xac, 0x92, 0xba, 0x8d, 0xa4, 0x14, 0x7a, 0x3c, 0xf6, 0x5f,
	0xa3, 0x5d, 0xfe, 0x13, 0x2b, 0x83, 0xa7, 0x31, 0x56, 0xb8, 0xe8, 0x06, 0x7a, 0xc6, 0xd6, 0xbb,
	0x40, 0xc6, 0x7e, 0x14, 0x19, 0x6b, 0x2b, 0xa2, 0x72, 0x17, 0xb4, 0x40, 0xcd, 0x8a, 0x7c, 0xf1,
	0x1a, 0x2b, 0x56, 0x51, 0x5e, 0x1b, 0xa4, 0x13, 0x80, 0x25, 0x17, 0xa8, 0xc7, 0x17, 0xe5, 0xab,
	0xbe, 0xf9, 0x47, 0xf1, 0xec, 0x7d, 0x00, 0x00, 0x00, 0xff, 0xff, 0x65, 0x5a, 0x26, 0x3c, 0x60,
	0x08, 0x00, 0x00,
}