// Code generated by protoc-gen-go. DO NOT EDIT.
// source: tsdb_plugin.proto

/*
Package proto is a generated protocol buffer package.

It is generated from these files:
	tsdb_plugin.proto

It has these top-level messages:
	TsdbQuery
	Query
	TimeRange
	Response
	QueryResult
	Table
	TableColumn
	TableRow
	DatasourceInfo
	TimeSeries
	Point
*/
package proto

import proto1 "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/any"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto1.ProtoPackageIsVersion2 // please upgrade the proto package

type TsdbQuery struct {
	TimeRange  *TimeRange      `protobuf:"bytes,1,opt,name=timeRange" json:"timeRange,omitempty"`
	Datasource *DatasourceInfo `protobuf:"bytes,2,opt,name=datasource" json:"datasource,omitempty"`
	Queries    []*Query        `protobuf:"bytes,3,rep,name=queries" json:"queries,omitempty"`
}

func (m *TsdbQuery) Reset()                    { *m = TsdbQuery{} }
func (m *TsdbQuery) String() string            { return proto1.CompactTextString(m) }
func (*TsdbQuery) ProtoMessage()               {}
func (*TsdbQuery) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *TsdbQuery) GetTimeRange() *TimeRange {
	if m != nil {
		return m.TimeRange
	}
	return nil
}

func (m *TsdbQuery) GetDatasource() *DatasourceInfo {
	if m != nil {
		return m.Datasource
	}
	return nil
}

func (m *TsdbQuery) GetQueries() []*Query {
	if m != nil {
		return m.Queries
	}
	return nil
}

type Query struct {
	RefId         string `protobuf:"bytes,1,opt,name=refId" json:"refId,omitempty"`
	MaxDataPoints int64  `protobuf:"varint,2,opt,name=maxDataPoints" json:"maxDataPoints,omitempty"`
	IntervalMs    int64  `protobuf:"varint,3,opt,name=intervalMs" json:"intervalMs,omitempty"`
	ModelJson     string `protobuf:"bytes,4,opt,name=modelJson" json:"modelJson,omitempty"`
}

func (m *Query) Reset()                    { *m = Query{} }
func (m *Query) String() string            { return proto1.CompactTextString(m) }
func (*Query) ProtoMessage()               {}
func (*Query) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Query) GetRefId() string {
	if m != nil {
		return m.RefId
	}
	return ""
}

func (m *Query) GetMaxDataPoints() int64 {
	if m != nil {
		return m.MaxDataPoints
	}
	return 0
}

func (m *Query) GetIntervalMs() int64 {
	if m != nil {
		return m.IntervalMs
	}
	return 0
}

func (m *Query) GetModelJson() string {
	if m != nil {
		return m.ModelJson
	}
	return ""
}

type TimeRange struct {
	FromRaw     string `protobuf:"bytes,1,opt,name=fromRaw" json:"fromRaw,omitempty"`
	ToRaw       string `protobuf:"bytes,2,opt,name=toRaw" json:"toRaw,omitempty"`
	FromEpochMs int64  `protobuf:"varint,3,opt,name=fromEpochMs" json:"fromEpochMs,omitempty"`
	ToEpochMs   int64  `protobuf:"varint,4,opt,name=toEpochMs" json:"toEpochMs,omitempty"`
}

func (m *TimeRange) Reset()                    { *m = TimeRange{} }
func (m *TimeRange) String() string            { return proto1.CompactTextString(m) }
func (*TimeRange) ProtoMessage()               {}
func (*TimeRange) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *TimeRange) GetFromRaw() string {
	if m != nil {
		return m.FromRaw
	}
	return ""
}

func (m *TimeRange) GetToRaw() string {
	if m != nil {
		return m.ToRaw
	}
	return ""
}

func (m *TimeRange) GetFromEpochMs() int64 {
	if m != nil {
		return m.FromEpochMs
	}
	return 0
}

func (m *TimeRange) GetToEpochMs() int64 {
	if m != nil {
		return m.ToEpochMs
	}
	return 0
}

type Response struct {
	Results []*QueryResult `protobuf:"bytes,1,rep,name=results" json:"results,omitempty"`
}

func (m *Response) Reset()                    { *m = Response{} }
func (m *Response) String() string            { return proto1.CompactTextString(m) }
func (*Response) ProtoMessage()               {}
func (*Response) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Response) GetResults() []*QueryResult {
	if m != nil {
		return m.Results
	}
	return nil
}

type QueryResult struct {
	Error    string        `protobuf:"bytes,1,opt,name=error" json:"error,omitempty"`
	RefId    string        `protobuf:"bytes,2,opt,name=refId" json:"refId,omitempty"`
	MetaJson string        `protobuf:"bytes,3,opt,name=metaJson" json:"metaJson,omitempty"`
	Series   []*TimeSeries `protobuf:"bytes,4,rep,name=series" json:"series,omitempty"`
	Tables   []*Table      `protobuf:"bytes,5,rep,name=tables" json:"tables,omitempty"`
}

func (m *QueryResult) Reset()                    { *m = QueryResult{} }
func (m *QueryResult) String() string            { return proto1.CompactTextString(m) }
func (*QueryResult) ProtoMessage()               {}
func (*QueryResult) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *QueryResult) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *QueryResult) GetRefId() string {
	if m != nil {
		return m.RefId
	}
	return ""
}

func (m *QueryResult) GetMetaJson() string {
	if m != nil {
		return m.MetaJson
	}
	return ""
}

func (m *QueryResult) GetSeries() []*TimeSeries {
	if m != nil {
		return m.Series
	}
	return nil
}

func (m *QueryResult) GetTables() []*Table {
	if m != nil {
		return m.Tables
	}
	return nil
}

type Table struct {
	Columns []*TableColumn `protobuf:"bytes,1,rep,name=columns" json:"columns,omitempty"`
	Rows    []*TableRow    `protobuf:"bytes,2,rep,name=rows" json:"rows,omitempty"`
}

func (m *Table) Reset()                    { *m = Table{} }
func (m *Table) String() string            { return proto1.CompactTextString(m) }
func (*Table) ProtoMessage()               {}
func (*Table) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *Table) GetColumns() []*TableColumn {
	if m != nil {
		return m.Columns
	}
	return nil
}

func (m *Table) GetRows() []*TableRow {
	if m != nil {
		return m.Rows
	}
	return nil
}

type TableColumn struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *TableColumn) Reset()                    { *m = TableColumn{} }
func (m *TableColumn) String() string            { return proto1.CompactTextString(m) }
func (*TableColumn) ProtoMessage()               {}
func (*TableColumn) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *TableColumn) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type TableRow struct {
	Values []*google_protobuf.Any `protobuf:"bytes,1,rep,name=values" json:"values,omitempty"`
}

func (m *TableRow) Reset()                    { *m = TableRow{} }
func (m *TableRow) String() string            { return proto1.CompactTextString(m) }
func (*TableRow) ProtoMessage()               {}
func (*TableRow) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *TableRow) GetValues() []*google_protobuf.Any {
	if m != nil {
		return m.Values
	}
	return nil
}

type DatasourceInfo struct {
	Id             int64  `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	OrgId          int64  `protobuf:"varint,2,opt,name=orgId" json:"orgId,omitempty"`
	Name           string `protobuf:"bytes,3,opt,name=name" json:"name,omitempty"`
	Type           string `protobuf:"bytes,4,opt,name=type" json:"type,omitempty"`
	Url            string `protobuf:"bytes,5,opt,name=url" json:"url,omitempty"`
	JsonData       string `protobuf:"bytes,6,opt,name=jsonData" json:"jsonData,omitempty"`
	SecureJsonData string `protobuf:"bytes,7,opt,name=secureJsonData" json:"secureJsonData,omitempty"`
}

func (m *DatasourceInfo) Reset()                    { *m = DatasourceInfo{} }
func (m *DatasourceInfo) String() string            { return proto1.CompactTextString(m) }
func (*DatasourceInfo) ProtoMessage()               {}
func (*DatasourceInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *DatasourceInfo) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *DatasourceInfo) GetOrgId() int64 {
	if m != nil {
		return m.OrgId
	}
	return 0
}

func (m *DatasourceInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *DatasourceInfo) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *DatasourceInfo) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

func (m *DatasourceInfo) GetJsonData() string {
	if m != nil {
		return m.JsonData
	}
	return ""
}

func (m *DatasourceInfo) GetSecureJsonData() string {
	if m != nil {
		return m.SecureJsonData
	}
	return ""
}

type TimeSeries struct {
	Name   string            `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Tags   map[string]string `protobuf:"bytes,2,rep,name=tags" json:"tags,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Points []*Point          `protobuf:"bytes,3,rep,name=points" json:"points,omitempty"`
}

func (m *TimeSeries) Reset()                    { *m = TimeSeries{} }
func (m *TimeSeries) String() string            { return proto1.CompactTextString(m) }
func (*TimeSeries) ProtoMessage()               {}
func (*TimeSeries) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *TimeSeries) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *TimeSeries) GetTags() map[string]string {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *TimeSeries) GetPoints() []*Point {
	if m != nil {
		return m.Points
	}
	return nil
}

type Point struct {
	Timestamp int64   `protobuf:"varint,1,opt,name=timestamp" json:"timestamp,omitempty"`
	Value     float64 `protobuf:"fixed64,2,opt,name=value" json:"value,omitempty"`
}

func (m *Point) Reset()                    { *m = Point{} }
func (m *Point) String() string            { return proto1.CompactTextString(m) }
func (*Point) ProtoMessage()               {}
func (*Point) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *Point) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *Point) GetValue() float64 {
	if m != nil {
		return m.Value
	}
	return 0
}

func init() {
	proto1.RegisterType((*TsdbQuery)(nil), "plugins.TsdbQuery")
	proto1.RegisterType((*Query)(nil), "plugins.Query")
	proto1.RegisterType((*TimeRange)(nil), "plugins.TimeRange")
	proto1.RegisterType((*Response)(nil), "plugins.Response")
	proto1.RegisterType((*QueryResult)(nil), "plugins.QueryResult")
	proto1.RegisterType((*Table)(nil), "plugins.Table")
	proto1.RegisterType((*TableColumn)(nil), "plugins.TableColumn")
	proto1.RegisterType((*TableRow)(nil), "plugins.TableRow")
	proto1.RegisterType((*DatasourceInfo)(nil), "plugins.DatasourceInfo")
	proto1.RegisterType((*TimeSeries)(nil), "plugins.TimeSeries")
	proto1.RegisterType((*Point)(nil), "plugins.Point")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for TsdbPlugin service

type TsdbPluginClient interface {
	Query(ctx context.Context, in *TsdbQuery, opts ...grpc.CallOption) (*Response, error)
}

type tsdbPluginClient struct {
	cc *grpc.ClientConn
}

func NewTsdbPluginClient(cc *grpc.ClientConn) TsdbPluginClient {
	return &tsdbPluginClient{cc}
}

func (c *tsdbPluginClient) Query(ctx context.Context, in *TsdbQuery, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := grpc.Invoke(ctx, "/plugins.TsdbPlugin/Query", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for TsdbPlugin service

type TsdbPluginServer interface {
	Query(context.Context, *TsdbQuery) (*Response, error)
}

func RegisterTsdbPluginServer(s *grpc.Server, srv TsdbPluginServer) {
	s.RegisterService(&_TsdbPlugin_serviceDesc, srv)
}

func _TsdbPlugin_Query_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TsdbQuery)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TsdbPluginServer).Query(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/plugins.TsdbPlugin/Query",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TsdbPluginServer).Query(ctx, req.(*TsdbQuery))
	}
	return interceptor(ctx, in, info, handler)
}

var _TsdbPlugin_serviceDesc = grpc.ServiceDesc{
	ServiceName: "plugins.TsdbPlugin",
	HandlerType: (*TsdbPluginServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Query",
			Handler:    _TsdbPlugin_Query_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "tsdb_plugin.proto",
}

func init() { proto1.RegisterFile("tsdb_plugin.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 665 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x54, 0x4d, 0x6f, 0xd4, 0x30,
	0x10, 0x55, 0x36, 0xfb, 0xd1, 0x9d, 0x15, 0x2b, 0x6a, 0x2a, 0x11, 0x56, 0x80, 0x4a, 0x04, 0x55,
	0x25, 0x50, 0x0a, 0xe5, 0xd0, 0xaa, 0x70, 0xe1, 0xa3, 0x87, 0x56, 0x42, 0x2a, 0x66, 0x4f, 0x1c,
	0x40, 0xde, 0x5d, 0x6f, 0x08, 0x24, 0x76, 0xb0, 0x9d, 0x96, 0x15, 0x27, 0xfe, 0x09, 0x67, 0xce,
	0xfc, 0x40, 0xe4, 0x49, 0x9c, 0x64, 0x97, 0x9e, 0x92, 0x99, 0xf7, 0x6c, 0x3f, 0xbf, 0x99, 0x31,
	0x6c, 0x1b, 0xbd, 0x98, 0x7d, 0xce, 0xd3, 0x22, 0x4e, 0x44, 0x94, 0x2b, 0x69, 0x24, 0x19, 0x94,
	0x91, 0x9e, 0xdc, 0x89, 0xa5, 0x8c, 0x53, 0x7e, 0x80, 0xe9, 0x59, 0xb1, 0x3c, 0x60, 0x62, 0x55,
	0x72, 0xc2, 0xdf, 0x1e, 0x0c, 0xa7, 0x7a, 0x31, 0x7b, 0x5f, 0x70, 0xb5, 0x22, 0x4f, 0x61, 0x68,
	0x92, 0x8c, 0x53, 0x26, 0x62, 0x1e, 0x78, 0xbb, 0xde, 0xfe, 0xe8, 0x90, 0x44, 0xd5, 0x2e, 0xd1,
	0xd4, 0x21, 0xb4, 0x21, 0x91, 0x23, 0x80, 0x05, 0x33, 0x4c, 0xcb, 0x42, 0xcd, 0x79, 0xd0, 0xc1,
	0x25, 0xb7, 0xeb, 0x25, 0x6f, 0x6b, 0xe8, 0x4c, 0x2c, 0x25, 0x6d, 0x51, 0xc9, 0x3e, 0x0c, 0xbe,
	0x17, 0x5c, 0x25, 0x5c, 0x07, 0xfe, 0xae, 0xbf, 0x3f, 0x3a, 0x1c, 0xd7, 0xab, 0x50, 0x0b, 0x75,
	0x70, 0xf8, 0xcb, 0x83, 0x5e, 0x29, 0x6f, 0x07, 0x7a, 0x8a, 0x2f, 0xcf, 0x16, 0x28, 0x6d, 0x48,
	0xcb, 0x80, 0x3c, 0x84, 0x1b, 0x19, 0xfb, 0x61, 0x8f, 0xba, 0x90, 0x89, 0x30, 0x1a, 0x55, 0xf8,
	0x74, 0x3d, 0x49, 0xee, 0x03, 0x24, 0xc2, 0x70, 0x75, 0xc9, 0xd2, 0x77, 0xf6, 0x48, 0x4b, 0x69,
	0x65, 0xc8, 0x5d, 0x18, 0x66, 0x72, 0xc1, 0xd3, 0x73, 0x2d, 0x45, 0xd0, 0xc5, 0xfd, 0x9b, 0x44,
	0xf8, 0x13, 0x86, 0xf5, 0xf5, 0x49, 0x00, 0x83, 0xa5, 0x92, 0x19, 0x65, 0x57, 0x95, 0x10, 0x17,
	0x5a, 0x81, 0x46, 0xda, 0x7c, 0xa7, 0x14, 0x88, 0x01, 0xd9, 0x85, 0x91, 0x25, 0x9c, 0xe6, 0x72,
	0xfe, 0xa5, 0x3e, 0xbb, 0x9d, 0xb2, 0x87, 0x1b, 0xe9, 0xf0, 0x2e, 0xe2, 0x4d, 0x22, 0x3c, 0x81,
	0x2d, 0xca, 0x75, 0x2e, 0x85, 0xe6, 0x24, 0x82, 0x81, 0xe2, 0xba, 0x48, 0x8d, 0x0e, 0x3c, 0xb4,
	0x6d, 0x67, 0xc3, 0x36, 0x04, 0xa9, 0x23, 0x85, 0x7f, 0x3c, 0x18, 0xb5, 0x00, 0xab, 0x90, 0x2b,
	0x25, 0x95, 0xb3, 0x10, 0x83, 0xc6, 0xd8, 0x4e, 0xdb, 0xd8, 0x09, 0x6c, 0x65, 0xdc, 0x30, 0x74,
	0xc4, 0x47, 0xa0, 0x8e, 0xc9, 0x63, 0xe8, 0xeb, 0xb2, 0x7a, 0x5d, 0x94, 0x71, 0x6b, 0xad, 0x4d,
	0x3e, 0x20, 0x44, 0x2b, 0x0a, 0xd9, 0x83, 0xbe, 0x61, 0xb3, 0x94, 0xeb, 0xa0, 0xb7, 0x51, 0xea,
	0xa9, 0x4d, 0xd3, 0x0a, 0x0d, 0x3f, 0x41, 0x0f, 0x13, 0xf6, 0x96, 0x73, 0x99, 0x16, 0x99, 0xf8,
	0xff, 0x96, 0x48, 0x78, 0x83, 0x20, 0x75, 0x24, 0xf2, 0x08, 0xba, 0x4a, 0x5e, 0xd9, 0xca, 0x5b,
	0xf2, 0xf6, 0xc6, 0xf6, 0xf2, 0x8a, 0x22, 0x1c, 0x3e, 0x80, 0x51, 0x6b, 0x39, 0x21, 0xd0, 0x15,
	0x2c, 0xe3, 0x95, 0x15, 0xf8, 0x1f, 0x1e, 0xc3, 0x96, 0x5b, 0x44, 0x9e, 0x40, 0xff, 0x92, 0xa5,
	0x05, 0x6f, 0x44, 0x94, 0x73, 0x14, 0xb9, 0x39, 0x8a, 0x5e, 0x89, 0x15, 0xad, 0x38, 0xe1, 0x5f,
	0x0f, 0xc6, 0xeb, 0xfd, 0x4e, 0xc6, 0xd0, 0x49, 0xca, 0x66, 0xf5, 0x69, 0x27, 0x59, 0x58, 0x9b,
	0xa5, 0x8a, 0x2b, 0x9b, 0x7d, 0x5a, 0x06, 0xb5, 0x0c, 0xbf, 0x91, 0x61, 0x73, 0x66, 0x95, 0xf3,
	0xaa, 0x11, 0xf1, 0x9f, 0xdc, 0x04, 0xbf, 0x50, 0x69, 0xd0, 0xc3, 0x94, 0xfd, 0xb5, 0x05, 0xfa,
	0xaa, 0xa5, 0xb0, 0xa7, 0x06, 0xfd, 0xb2, 0x40, 0x2e, 0x26, 0x7b, 0x30, 0xd6, 0x7c, 0x5e, 0x28,
	0x7e, 0xee, 0x18, 0x03, 0x64, 0x6c, 0x64, 0xad, 0x6c, 0x68, 0x4a, 0x76, 0x9d, 0x27, 0xe4, 0x19,
	0x74, 0x0d, 0x8b, 0x9d, 0xbb, 0xf7, 0xae, 0xa9, 0x74, 0x34, 0x65, 0xb1, 0x3e, 0x15, 0x46, 0xad,
	0x28, 0x52, 0x6d, 0xc5, 0xf3, 0x72, 0x18, 0x37, 0x87, 0x1b, 0xc7, 0x91, 0x56, 0xe8, 0xe4, 0x08,
	0x86, 0xf5, 0x52, 0x7b, 0xc1, 0x6f, 0x7c, 0x55, 0x1d, 0x6d, 0x7f, 0xad, 0x61, 0xe8, 0xae, 0xeb,
	0x4b, 0x0c, 0x4e, 0x3a, 0xc7, 0x5e, 0xf8, 0x02, 0x7a, 0xb8, 0x13, 0x8e, 0x4e, 0x92, 0x71, 0x6d,
	0x58, 0x96, 0x57, 0x56, 0x37, 0x89, 0xf5, 0x0d, 0xbc, 0x6a, 0x83, 0xc3, 0x97, 0x00, 0xf6, 0xcd,
	0xbb, 0x40, 0x49, 0x24, 0x72, 0xcf, 0x4b, 0xeb, 0xa9, 0x73, 0x2f, 0xe2, 0xa4, 0xe9, 0x25, 0x37,
	0x82, 0xaf, 0x07, 0x1f, 0x7b, 0x65, 0x03, 0xf4, 0xf1, 0xf3, 0xfc, 0x5f, 0x00, 0x00, 0x00, 0xff,
	0xff, 0xa8, 0xec, 0x66, 0xef, 0x7b, 0x05, 0x00, 0x00,
}
