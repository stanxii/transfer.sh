// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/cloud/talent/v4beta1/resume_service.proto

package talent // import "google.golang.org/genproto/googleapis/cloud/talent/v4beta1"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Parse resume request.
type ParseResumeRequest struct {
	// Required.
	//
	// The resource name of the project.
	//
	// The format is "projects/{project_id}", for example,
	// "projects/api-test-project".
	Parent string `protobuf:"bytes,1,opt,name=parent,proto3" json:"parent,omitempty"`
	// Required.
	//
	// The bytes of the resume file in common format, for example, PDF, TXT.
	// UTF-8 encoding is required if the resume is text-based, otherwise an error
	// is thrown.
	Resume []byte `protobuf:"bytes,2,opt,name=resume,proto3" json:"resume,omitempty"`
	// Optional.
	//
	// The region code indicating where the resume is from. Values
	// are as per the ISO-3166-2 format. For example, US, FR, DE.
	//
	// This value is optional, but providing this value improves the resume
	// parsing quality and performance.
	//
	// An error is thrown if the regionCode is invalid.
	RegionCode string `protobuf:"bytes,3,opt,name=region_code,json=regionCode,proto3" json:"region_code,omitempty"`
	// Optional.
	//
	// The language code of contents in the resume.
	//
	// Language codes must be in BCP-47 format, such as "en-US" or "sr-Latn".
	// For more information, see
	// [Tags for Identifying Languages](https://tools.ietf.org/html/bcp47){:
	// class="external" target="_blank" }.
	LanguageCode string `protobuf:"bytes,4,opt,name=language_code,json=languageCode,proto3" json:"language_code,omitempty"`
	// Optional.
	//
	// Options that change how the resume parse is performed.
	Options              *ParseResumeOptions `protobuf:"bytes,5,opt,name=options,proto3" json:"options,omitempty"`
	XXX_NoUnkeyedLiteral struct{}            `json:"-"`
	XXX_unrecognized     []byte              `json:"-"`
	XXX_sizecache        int32               `json:"-"`
}

func (m *ParseResumeRequest) Reset()         { *m = ParseResumeRequest{} }
func (m *ParseResumeRequest) String() string { return proto.CompactTextString(m) }
func (*ParseResumeRequest) ProtoMessage()    {}
func (*ParseResumeRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_resume_service_d325ab94b73b8b68, []int{0}
}
func (m *ParseResumeRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ParseResumeRequest.Unmarshal(m, b)
}
func (m *ParseResumeRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ParseResumeRequest.Marshal(b, m, deterministic)
}
func (dst *ParseResumeRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ParseResumeRequest.Merge(dst, src)
}
func (m *ParseResumeRequest) XXX_Size() int {
	return xxx_messageInfo_ParseResumeRequest.Size(m)
}
func (m *ParseResumeRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ParseResumeRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ParseResumeRequest proto.InternalMessageInfo

func (m *ParseResumeRequest) GetParent() string {
	if m != nil {
		return m.Parent
	}
	return ""
}

func (m *ParseResumeRequest) GetResume() []byte {
	if m != nil {
		return m.Resume
	}
	return nil
}

func (m *ParseResumeRequest) GetRegionCode() string {
	if m != nil {
		return m.RegionCode
	}
	return ""
}

func (m *ParseResumeRequest) GetLanguageCode() string {
	if m != nil {
		return m.LanguageCode
	}
	return ""
}

func (m *ParseResumeRequest) GetOptions() *ParseResumeOptions {
	if m != nil {
		return m.Options
	}
	return nil
}

// Options that change how the resume parse is performed.
type ParseResumeOptions struct {
	// Optional.
	//
	// Controls whether Optical Character Recognition (OCR) is enabled.
	//
	// OCR is used to decipher pictorial resumes, or resumes that have some
	// element of pictorial detail (for example, contact information placed within
	// an image in a pdf). Note that the API call has a higher latency if OCR is
	// enabled.
	EnableOcr bool `protobuf:"varint,1,opt,name=enable_ocr,json=enableOcr,proto3" json:"enable_ocr,omitempty"`
	// Optional.
	//
	// Controls whether detected skills are included in the parsed profile from
	// sections of the resume other than just skills sections.
	//
	// Normally, returned skills are limited to those taken from a resume section
	// intended to list skills. When enabled, this feature causes detected
	// skills in other sections to also be included in the returned profile.
	EnableFullSkillDetection bool     `protobuf:"varint,2,opt,name=enable_full_skill_detection,json=enableFullSkillDetection,proto3" json:"enable_full_skill_detection,omitempty"`
	XXX_NoUnkeyedLiteral     struct{} `json:"-"`
	XXX_unrecognized         []byte   `json:"-"`
	XXX_sizecache            int32    `json:"-"`
}

func (m *ParseResumeOptions) Reset()         { *m = ParseResumeOptions{} }
func (m *ParseResumeOptions) String() string { return proto.CompactTextString(m) }
func (*ParseResumeOptions) ProtoMessage()    {}
func (*ParseResumeOptions) Descriptor() ([]byte, []int) {
	return fileDescriptor_resume_service_d325ab94b73b8b68, []int{1}
}
func (m *ParseResumeOptions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ParseResumeOptions.Unmarshal(m, b)
}
func (m *ParseResumeOptions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ParseResumeOptions.Marshal(b, m, deterministic)
}
func (dst *ParseResumeOptions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ParseResumeOptions.Merge(dst, src)
}
func (m *ParseResumeOptions) XXX_Size() int {
	return xxx_messageInfo_ParseResumeOptions.Size(m)
}
func (m *ParseResumeOptions) XXX_DiscardUnknown() {
	xxx_messageInfo_ParseResumeOptions.DiscardUnknown(m)
}

var xxx_messageInfo_ParseResumeOptions proto.InternalMessageInfo

func (m *ParseResumeOptions) GetEnableOcr() bool {
	if m != nil {
		return m.EnableOcr
	}
	return false
}

func (m *ParseResumeOptions) GetEnableFullSkillDetection() bool {
	if m != nil {
		return m.EnableFullSkillDetection
	}
	return false
}

// Parse resume response.
type ParseResumeResponse struct {
	// The profile parsed from resume.
	Profile *Profile `protobuf:"bytes,1,opt,name=profile,proto3" json:"profile,omitempty"`
	// Raw text from resume.
	RawText              string   `protobuf:"bytes,2,opt,name=raw_text,json=rawText,proto3" json:"raw_text,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ParseResumeResponse) Reset()         { *m = ParseResumeResponse{} }
func (m *ParseResumeResponse) String() string { return proto.CompactTextString(m) }
func (*ParseResumeResponse) ProtoMessage()    {}
func (*ParseResumeResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_resume_service_d325ab94b73b8b68, []int{2}
}
func (m *ParseResumeResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ParseResumeResponse.Unmarshal(m, b)
}
func (m *ParseResumeResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ParseResumeResponse.Marshal(b, m, deterministic)
}
func (dst *ParseResumeResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ParseResumeResponse.Merge(dst, src)
}
func (m *ParseResumeResponse) XXX_Size() int {
	return xxx_messageInfo_ParseResumeResponse.Size(m)
}
func (m *ParseResumeResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ParseResumeResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ParseResumeResponse proto.InternalMessageInfo

func (m *ParseResumeResponse) GetProfile() *Profile {
	if m != nil {
		return m.Profile
	}
	return nil
}

func (m *ParseResumeResponse) GetRawText() string {
	if m != nil {
		return m.RawText
	}
	return ""
}

func init() {
	proto.RegisterType((*ParseResumeRequest)(nil), "google.cloud.talent.v4beta1.ParseResumeRequest")
	proto.RegisterType((*ParseResumeOptions)(nil), "google.cloud.talent.v4beta1.ParseResumeOptions")
	proto.RegisterType((*ParseResumeResponse)(nil), "google.cloud.talent.v4beta1.ParseResumeResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ResumeServiceClient is the client API for ResumeService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ResumeServiceClient interface {
	// Parses a resume into a [Profile][google.cloud.talent.v4beta1.Profile]. The
	// API attempts to fill out the following profile fields if present within the
	// resume:
	//
	// * personNames
	// * addresses
	// * emailAddress
	// * phoneNumbers
	// * personalUris
	// * employmentRecords
	// * educationRecords
	// * skills
	//
	// Note that some attributes in these fields may not be populated if they're
	// not present within the resume or unrecognizable by the resume parser.
	//
	// This API does not save the resume or profile. To create a profile from this
	// resume, clients need to call the CreateProfile method again with the
	// profile returned.
	//
	// The following list of formats are supported:
	//
	// * PDF
	// * TXT
	// * DOC
	// * RTF
	// * DOCX
	// * PNG (only when [ParseResumeRequest.enable_ocr][] is set to `true`,
	// otherwise an error is thrown)
	ParseResume(ctx context.Context, in *ParseResumeRequest, opts ...grpc.CallOption) (*ParseResumeResponse, error)
}

type resumeServiceClient struct {
	cc *grpc.ClientConn
}

func NewResumeServiceClient(cc *grpc.ClientConn) ResumeServiceClient {
	return &resumeServiceClient{cc}
}

func (c *resumeServiceClient) ParseResume(ctx context.Context, in *ParseResumeRequest, opts ...grpc.CallOption) (*ParseResumeResponse, error) {
	out := new(ParseResumeResponse)
	err := c.cc.Invoke(ctx, "/google.cloud.talent.v4beta1.ResumeService/ParseResume", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ResumeServiceServer is the server API for ResumeService service.
type ResumeServiceServer interface {
	// Parses a resume into a [Profile][google.cloud.talent.v4beta1.Profile]. The
	// API attempts to fill out the following profile fields if present within the
	// resume:
	//
	// * personNames
	// * addresses
	// * emailAddress
	// * phoneNumbers
	// * personalUris
	// * employmentRecords
	// * educationRecords
	// * skills
	//
	// Note that some attributes in these fields may not be populated if they're
	// not present within the resume or unrecognizable by the resume parser.
	//
	// This API does not save the resume or profile. To create a profile from this
	// resume, clients need to call the CreateProfile method again with the
	// profile returned.
	//
	// The following list of formats are supported:
	//
	// * PDF
	// * TXT
	// * DOC
	// * RTF
	// * DOCX
	// * PNG (only when [ParseResumeRequest.enable_ocr][] is set to `true`,
	// otherwise an error is thrown)
	ParseResume(context.Context, *ParseResumeRequest) (*ParseResumeResponse, error)
}

func RegisterResumeServiceServer(s *grpc.Server, srv ResumeServiceServer) {
	s.RegisterService(&_ResumeService_serviceDesc, srv)
}

func _ResumeService_ParseResume_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ParseResumeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ResumeServiceServer).ParseResume(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/google.cloud.talent.v4beta1.ResumeService/ParseResume",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ResumeServiceServer).ParseResume(ctx, req.(*ParseResumeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ResumeService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "google.cloud.talent.v4beta1.ResumeService",
	HandlerType: (*ResumeServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ParseResume",
			Handler:    _ResumeService_ParseResume_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "google/cloud/talent/v4beta1/resume_service.proto",
}

func init() {
	proto.RegisterFile("google/cloud/talent/v4beta1/resume_service.proto", fileDescriptor_resume_service_d325ab94b73b8b68)
}

var fileDescriptor_resume_service_d325ab94b73b8b68 = []byte{
	// 476 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xc1, 0x6e, 0x13, 0x31,
	0x10, 0x86, 0xe5, 0x16, 0x9a, 0xc4, 0x69, 0x2f, 0x46, 0x42, 0x4b, 0x0a, 0x6a, 0xb4, 0x70, 0x08,
	0x39, 0xac, 0x4b, 0x80, 0x4b, 0x51, 0x91, 0x68, 0x11, 0x12, 0xa7, 0x46, 0x9b, 0x9e, 0xb8, 0xac,
	0x9c, 0xcd, 0x74, 0xb5, 0xe0, 0x7a, 0x8c, 0xed, 0x6d, 0x2b, 0x21, 0x24, 0xc4, 0x2b, 0xf0, 0x06,
	0x5c, 0x78, 0x06, 0x9e, 0x03, 0xf1, 0x06, 0x3c, 0x08, 0x5a, 0xdb, 0x41, 0x8d, 0x84, 0xa2, 0x1e,
	0x3d, 0xf3, 0x7f, 0x3b, 0xf3, 0xcf, 0xfe, 0x74, 0xbf, 0x42, 0xac, 0x24, 0xf0, 0x52, 0x62, 0xb3,
	0xe0, 0x4e, 0x48, 0x50, 0x8e, 0x5f, 0x3c, 0x9b, 0x83, 0x13, 0x4f, 0xb8, 0x01, 0xdb, 0x9c, 0x43,
	0x61, 0xc1, 0x5c, 0xd4, 0x25, 0x64, 0xda, 0xa0, 0x43, 0xb6, 0x1b, 0x88, 0xcc, 0x13, 0x59, 0x20,
	0xb2, 0x48, 0x0c, 0xee, 0xc7, 0xcf, 0x09, 0x5d, 0x73, 0xa1, 0x14, 0x3a, 0xe1, 0x6a, 0x54, 0x36,
	0xa0, 0x83, 0xc7, 0xeb, 0x86, 0x69, 0x83, 0x67, 0xb5, 0x8c, 0x53, 0xd2, 0xdf, 0x84, 0xb2, 0xa9,
	0x30, 0x16, 0x72, 0xbf, 0x43, 0x0e, 0x1f, 0x1b, 0xb0, 0x8e, 0xdd, 0xa5, 0x5b, 0x5a, 0x18, 0x50,
	0x2e, 0x21, 0x43, 0x32, 0xea, 0xe5, 0xf1, 0xd5, 0xd6, 0xc3, 0xb2, 0xc9, 0xc6, 0x90, 0x8c, 0xb6,
	0xf3, 0xf8, 0x62, 0x7b, 0xb4, 0x6f, 0xa0, 0xaa, 0x51, 0x15, 0x25, 0x2e, 0x20, 0xd9, 0xf4, 0x10,
	0x0d, 0xa5, 0x63, 0x5c, 0x00, 0x7b, 0x48, 0x77, 0xa4, 0x50, 0x55, 0x23, 0x2a, 0x08, 0x92, 0x5b,
	0x5e, 0xb2, 0xbd, 0x2c, 0x7a, 0xd1, 0x5b, 0xda, 0x41, 0xed, 0x8d, 0x24, 0xb7, 0x87, 0x64, 0xd4,
	0x9f, 0xf0, 0x6c, 0xcd, 0x11, 0xb2, 0x6b, 0x7b, 0x9f, 0x04, 0x2c, 0x5f, 0xf2, 0xa9, 0x59, 0xb1,
	0x15, 0xdb, 0xec, 0x01, 0xa5, 0xa0, 0xc4, 0x5c, 0x42, 0x81, 0xa5, 0xf1, 0xd6, 0xba, 0x79, 0x2f,
	0x54, 0x4e, 0x4a, 0xc3, 0x0e, 0xe9, 0x6e, 0x6c, 0x9f, 0x35, 0x52, 0x16, 0xf6, 0x43, 0x2d, 0x65,
	0xb1, 0x00, 0x07, 0x65, 0x8b, 0x7b, 0xcb, 0xdd, 0x3c, 0x09, 0x92, 0x37, 0x8d, 0x94, 0xb3, 0x56,
	0xf0, 0x7a, 0xd9, 0x4f, 0x35, 0xbd, 0xb3, 0x72, 0x4a, 0xab, 0x51, 0x59, 0x60, 0x2f, 0x69, 0x27,
	0xde, 0xdc, 0x4f, 0xec, 0x4f, 0x1e, 0xad, 0x77, 0x15, 0xb4, 0xf9, 0x12, 0x62, 0xf7, 0x68, 0xd7,
	0x88, 0xcb, 0xc2, 0xc1, 0x95, 0xf3, 0x2b, 0xf4, 0xf2, 0x8e, 0x11, 0x97, 0xa7, 0x70, 0xe5, 0x26,
	0x3f, 0x09, 0xdd, 0x09, 0xd3, 0x66, 0x21, 0x3b, 0xec, 0x07, 0xa1, 0xfd, 0x6b, 0x4b, 0xb0, 0x1b,
	0x5f, 0x30, 0xfe, 0xf9, 0xc1, 0xfe, 0xcd, 0x81, 0xe0, 0x2f, 0x7d, 0xfe, 0xf5, 0xd7, 0x9f, 0x6f,
	0x1b, 0x3c, 0x1d, 0xff, 0x8b, 0xd8, 0xa7, 0x90, 0x96, 0x43, 0x6d, 0xf0, 0x3d, 0x94, 0xce, 0xf2,
	0xf1, 0xe7, 0x98, 0x71, 0x7b, 0xa0, 0xdb, 0x2f, 0x1c, 0x90, 0xf1, 0xd1, 0x17, 0x42, 0xf7, 0x4a,
	0x3c, 0x5f, 0x37, 0xee, 0x88, 0xad, 0x98, 0x9b, 0xb6, 0x89, 0x9d, 0x92, 0x77, 0xaf, 0x22, 0x52,
	0x61, 0x9b, 0x9e, 0x0c, 0x4d, 0xc5, 0x2b, 0x50, 0x3e, 0xcf, 0x3c, 0xb4, 0x84, 0xae, 0xed, 0x7f,
	0xd3, 0xff, 0x22, 0x3c, 0xbf, 0x6f, 0x6c, 0x1e, 0x9f, 0xce, 0xe6, 0x5b, 0x9e, 0x79, 0xfa, 0x37,
	0x00, 0x00, 0xff, 0xff, 0xf9, 0x33, 0x2d, 0x87, 0x9d, 0x03, 0x00, 0x00,
}
