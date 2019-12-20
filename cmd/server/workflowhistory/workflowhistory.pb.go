// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cmd/server/workflowhistory/workflowhistory.proto

package workflowhistory

import (
	context "context"
	fmt "fmt"
	v1alpha1 "github.com/argoproj/argo/pkg/apis/workflow/v1alpha1"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	_ "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

type WorkflowHistoryListRequest struct {
	ListOptions          *v1.ListOptions `protobuf:"bytes,1,opt,name=listOptions,proto3" json:"listOptions,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *WorkflowHistoryListRequest) Reset()         { *m = WorkflowHistoryListRequest{} }
func (m *WorkflowHistoryListRequest) String() string { return proto.CompactTextString(m) }
func (*WorkflowHistoryListRequest) ProtoMessage()    {}
func (*WorkflowHistoryListRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0a546c3f97ba9ddb, []int{0}
}
func (m *WorkflowHistoryListRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *WorkflowHistoryListRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_WorkflowHistoryListRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *WorkflowHistoryListRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_WorkflowHistoryListRequest.Merge(m, src)
}
func (m *WorkflowHistoryListRequest) XXX_Size() int {
	return m.Size()
}
func (m *WorkflowHistoryListRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_WorkflowHistoryListRequest.DiscardUnknown(m)
}

var xxx_messageInfo_WorkflowHistoryListRequest proto.InternalMessageInfo

func (m *WorkflowHistoryListRequest) GetListOptions() *v1.ListOptions {
	if m != nil {
		return m.ListOptions
	}
	return nil
}

func init() {
	proto.RegisterType((*WorkflowHistoryListRequest)(nil), "WorkflowHistoryListRequest")
}

func init() {
	proto.RegisterFile("cmd/server/workflowhistory/workflowhistory.proto", fileDescriptor_0a546c3f97ba9ddb)
}

var fileDescriptor_0a546c3f97ba9ddb = []byte{
	// 345 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x52, 0xbd, 0x4e, 0xf3, 0x30,
	0x14, 0x95, 0xbf, 0xe1, 0x1b, 0xd2, 0x2d, 0xdf, 0x27, 0x04, 0x2d, 0xaa, 0xa0, 0x13, 0x03, 0xb2,
	0x49, 0x61, 0x40, 0x6c, 0xc0, 0xc2, 0x50, 0x09, 0xa9, 0x1d, 0x90, 0xd8, 0xdc, 0xf4, 0xe2, 0x98,
	0xfc, 0xdc, 0xd4, 0x76, 0x53, 0x75, 0xe5, 0x15, 0x58, 0x78, 0x0f, 0x5e, 0x82, 0x11, 0x89, 0x17,
	0x40, 0x15, 0x0f, 0x82, 0xec, 0x26, 0xa4, 0x0a, 0xed, 0x76, 0x72, 0xaf, 0xce, 0x39, 0xf7, 0x1c,
	0xc7, 0x3b, 0x09, 0xd3, 0x09, 0xd3, 0xa0, 0x0a, 0x50, 0x6c, 0x8e, 0x2a, 0x7e, 0x48, 0x70, 0x1e,
	0x49, 0x6d, 0x50, 0x2d, 0x9a, 0xdf, 0x34, 0x57, 0x68, 0xb0, 0xfd, 0x5f, 0xa0, 0x40, 0x07, 0x99,
	0x45, 0xe5, 0x74, 0x5f, 0x20, 0x8a, 0x04, 0x18, 0xcf, 0x25, 0xe3, 0x59, 0x86, 0x86, 0x1b, 0x89,
	0x99, 0x2e, 0xb7, 0x67, 0xf1, 0xb9, 0xa6, 0x12, 0xed, 0x36, 0xe5, 0x61, 0x24, 0x33, 0x50, 0x0b,
	0x96, 0xc7, 0xc2, 0x0e, 0x34, 0x4b, 0xc1, 0x70, 0x56, 0x04, 0x4c, 0x40, 0x06, 0x8a, 0x1b, 0x98,
	0x94, 0xac, 0x6b, 0x21, 0x4d, 0x34, 0x1b, 0xd3, 0x10, 0x53, 0xc6, 0x95, 0x33, 0x7d, 0x74, 0xa0,
	0xa6, 0x56, 0x27, 0xb2, 0x22, 0xe0, 0x49, 0x1e, 0xf1, 0xdf, 0x22, 0xbd, 0xda, 0x9a, 0x85, 0xa8,
	0x60, 0x83, 0x51, 0x6f, 0xea, 0xb5, 0xef, 0x4a, 0xa1, 0x9b, 0x55, 0xd6, 0x81, 0xd4, 0x66, 0x08,
	0xd3, 0x19, 0x68, 0xe3, 0x8f, 0xbc, 0x56, 0x22, 0xb5, 0xb9, 0xcd, 0x5d, 0xa2, 0x5d, 0x72, 0x40,
	0x8e, 0x5a, 0xfd, 0x80, 0xae, 0x74, 0xe9, 0x7a, 0x24, 0x9a, 0xc7, 0xc2, 0x0e, 0x34, 0xb5, 0x91,
	0x68, 0x11, 0xd0, 0x41, 0x4d, 0x1c, 0xae, 0xab, 0xf4, 0x5f, 0x89, 0xb7, 0xd3, 0xf0, 0x1c, 0x81,
	0x2a, 0x64, 0x08, 0xfe, 0x0b, 0xf1, 0xfe, 0x59, 0x5e, 0x63, 0xed, 0x77, 0xe8, 0xf6, 0x23, 0xdb,
	0x97, 0xb4, 0x2e, 0x8b, 0x56, 0x65, 0x39, 0x50, 0x1f, 0x55, 0x95, 0x45, 0xab, 0xb2, 0x7e, 0x04,
	0xad, 0x52, 0xef, 0xf0, 0xe9, 0xe3, 0xeb, 0xf9, 0x4f, 0xc7, 0xdf, 0x73, 0x5d, 0x15, 0x41, 0xe3,
	0xfd, 0x25, 0xe8, 0xab, 0x8b, 0xb7, 0x65, 0x97, 0xbc, 0x2f, 0xbb, 0xe4, 0x73, 0xd9, 0x25, 0xf7,
	0xc7, 0x5b, 0xdf, 0x67, 0xc3, 0x4f, 0x35, 0xfe, 0xeb, 0xba, 0x3e, 0xfd, 0x0e, 0x00, 0x00, 0xff,
	0xff, 0xa5, 0xaf, 0xb2, 0x48, 0x72, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// WorkflowHistoryServiceClient is the client API for WorkflowHistoryService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type WorkflowHistoryServiceClient interface {
	ListWorkflowHistory(ctx context.Context, in *WorkflowHistoryListRequest, opts ...grpc.CallOption) (*v1alpha1.WorkflowList, error)
}

type workflowHistoryServiceClient struct {
	cc *grpc.ClientConn
}

func NewWorkflowHistoryServiceClient(cc *grpc.ClientConn) WorkflowHistoryServiceClient {
	return &workflowHistoryServiceClient{cc}
}

func (c *workflowHistoryServiceClient) ListWorkflowHistory(ctx context.Context, in *WorkflowHistoryListRequest, opts ...grpc.CallOption) (*v1alpha1.WorkflowList, error) {
	out := new(v1alpha1.WorkflowList)
	err := c.cc.Invoke(ctx, "/WorkflowHistoryService/ListWorkflowHistory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// WorkflowHistoryServiceServer is the server API for WorkflowHistoryService service.
type WorkflowHistoryServiceServer interface {
	ListWorkflowHistory(context.Context, *WorkflowHistoryListRequest) (*v1alpha1.WorkflowList, error)
}

// UnimplementedWorkflowHistoryServiceServer can be embedded to have forward compatible implementations.
type UnimplementedWorkflowHistoryServiceServer struct {
}

func (*UnimplementedWorkflowHistoryServiceServer) ListWorkflowHistory(ctx context.Context, req *WorkflowHistoryListRequest) (*v1alpha1.WorkflowList, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListWorkflowHistory not implemented")
}

func RegisterWorkflowHistoryServiceServer(s *grpc.Server, srv WorkflowHistoryServiceServer) {
	s.RegisterService(&_WorkflowHistoryService_serviceDesc, srv)
}

func _WorkflowHistoryService_ListWorkflowHistory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WorkflowHistoryListRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WorkflowHistoryServiceServer).ListWorkflowHistory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/WorkflowHistoryService/ListWorkflowHistory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WorkflowHistoryServiceServer).ListWorkflowHistory(ctx, req.(*WorkflowHistoryListRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _WorkflowHistoryService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "WorkflowHistoryService",
	HandlerType: (*WorkflowHistoryServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListWorkflowHistory",
			Handler:    _WorkflowHistoryService_ListWorkflowHistory_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cmd/server/workflowhistory/workflowhistory.proto",
}

func (m *WorkflowHistoryListRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *WorkflowHistoryListRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *WorkflowHistoryListRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.ListOptions != nil {
		{
			size, err := m.ListOptions.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintWorkflowhistory(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintWorkflowhistory(dAtA []byte, offset int, v uint64) int {
	offset -= sovWorkflowhistory(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *WorkflowHistoryListRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ListOptions != nil {
		l = m.ListOptions.Size()
		n += 1 + l + sovWorkflowhistory(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovWorkflowhistory(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozWorkflowhistory(x uint64) (n int) {
	return sovWorkflowhistory(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *WorkflowHistoryListRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowWorkflowhistory
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: WorkflowHistoryListRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: WorkflowHistoryListRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ListOptions", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowWorkflowhistory
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthWorkflowhistory
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthWorkflowhistory
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.ListOptions == nil {
				m.ListOptions = &v1.ListOptions{}
			}
			if err := m.ListOptions.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipWorkflowhistory(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthWorkflowhistory
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthWorkflowhistory
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipWorkflowhistory(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowWorkflowhistory
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowWorkflowhistory
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowWorkflowhistory
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthWorkflowhistory
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthWorkflowhistory
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowWorkflowhistory
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipWorkflowhistory(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthWorkflowhistory
				}
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthWorkflowhistory = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowWorkflowhistory   = fmt.Errorf("proto: integer overflow")
)