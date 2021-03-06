// Code generated by protoc-gen-go.
// source: file.proto
// DO NOT EDIT!

/*
Package proto is a generated protocol buffer package.

It is generated from these files:
	file.proto
	stats.proto

It has these top-level messages:
	RequestHeader
	PathError
	SyscallError
	Error
	FileInfo
	WriteRequest
	WriteReply
	ReadRequest
	ReadReply
	RenameRequest
	RenameReply
	ReadDirRequest
	ReadDirReply
	RemoveRequest
	RemoveReply
	MkdirRequest
	MkdirReply
	ReconstructSrc
	ReconstructDst
	ReconstructRequest
	ReconstructReply
*/
package proto

import proto1 "github.com/golang/protobuf/proto"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto1.Marshal

type RequestHeader struct {
	ClientID int64 `protobuf:"varint,1,opt,name=clientID" json:"clientID,omitempty"`
}

func (m *RequestHeader) Reset()         { *m = RequestHeader{} }
func (m *RequestHeader) String() string { return proto1.CompactTextString(m) }
func (*RequestHeader) ProtoMessage()    {}

// PathError records an error and the operation and file path that caused it.
type PathError struct {
	Op    string `protobuf:"bytes,1,opt,name=op" json:"op,omitempty"`
	Path  string `protobuf:"bytes,2,opt,name=path" json:"path,omitempty"`
	Error string `protobuf:"bytes,3,opt,name=error" json:"error,omitempty"`
}

func (m *PathError) Reset()         { *m = PathError{} }
func (m *PathError) String() string { return proto1.CompactTextString(m) }
func (*PathError) ProtoMessage()    {}

// SyscallError records an error from a specific system call.
type SyscallError struct {
	Syscall string `protobuf:"bytes,1,opt,name=syscall" json:"syscall,omitempty"`
	Error   string `protobuf:"bytes,2,opt,name=error" json:"error,omitempty"`
}

func (m *SyscallError) Reset()         { *m = SyscallError{} }
func (m *SyscallError) String() string { return proto1.CompactTextString(m) }
func (*SyscallError) ProtoMessage()    {}

type Error struct {
	PathErr *PathError    `protobuf:"bytes,1,opt,name=pathErr" json:"pathErr,omitempty"`
	SysErr  *SyscallError `protobuf:"bytes,2,opt,name=sysErr" json:"sysErr,omitempty"`
}

func (m *Error) Reset()         { *m = Error{} }
func (m *Error) String() string { return proto1.CompactTextString(m) }
func (*Error) ProtoMessage()    {}

func (m *Error) GetPathErr() *PathError {
	if m != nil {
		return m.PathErr
	}
	return nil
}

func (m *Error) GetSysErr() *SyscallError {
	if m != nil {
		return m.SysErr
	}
	return nil
}

type FileInfo struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Size int64  `protobuf:"varint,2,opt,name=size" json:"size,omitempty"`
	// including block header(CRC), padding zero bytes
	TotalSize int64 `protobuf:"varint,3,opt,name=total_size" json:"total_size,omitempty"`
	ModTime   int64 `protobuf:"varint,4,opt,name=mod_time" json:"mod_time,omitempty"`
	IsDir     bool  `protobuf:"varint,5,opt,name=is_dir" json:"is_dir,omitempty"`
}

func (m *FileInfo) Reset()         { *m = FileInfo{} }
func (m *FileInfo) String() string { return proto1.CompactTextString(m) }
func (*FileInfo) ProtoMessage()    {}

// Write writes len(b) bytes from the given offset. It returns the number
// of bytes written and an error, if any.
// Write returns an error when n != len(b).
type WriteRequest struct {
	Header *RequestHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	Name   string         `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Offset int64          `protobuf:"varint,3,opt,name=offset" json:"offset,omitempty"`
	Data   []byte         `protobuf:"bytes,4,opt,name=data,proto3" json:"data,omitempty"`
	Append bool           `protobuf:"varint,5,opt,name=append" json:"append,omitempty"`
}

func (m *WriteRequest) Reset()         { *m = WriteRequest{} }
func (m *WriteRequest) String() string { return proto1.CompactTextString(m) }
func (*WriteRequest) ProtoMessage()    {}

func (m *WriteRequest) GetHeader() *RequestHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

type WriteReply struct {
	Error        *Error `protobuf:"bytes,1,opt,name=error" json:"error,omitempty"`
	BytesWritten int64  `protobuf:"varint,2,opt,name=bytes_written" json:"bytes_written,omitempty"`
}

func (m *WriteReply) Reset()         { *m = WriteReply{} }
func (m *WriteReply) String() string { return proto1.CompactTextString(m) }
func (*WriteReply) ProtoMessage()    {}

func (m *WriteReply) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

// Read reads up to length bytes. The checksum of the data must match the exp_checksum if given, or an error is returned.
type ReadRequest struct {
	Header      *RequestHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	Name        string         `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Offset      int64          `protobuf:"varint,3,opt,name=offset" json:"offset,omitempty"`
	Length      int64          `protobuf:"varint,4,opt,name=length" json:"length,omitempty"`
	ExpChecksum uint32         `protobuf:"fixed32,5,opt,name=exp_checksum" json:"exp_checksum,omitempty"`
}

func (m *ReadRequest) Reset()         { *m = ReadRequest{} }
func (m *ReadRequest) String() string { return proto1.CompactTextString(m) }
func (*ReadRequest) ProtoMessage()    {}

func (m *ReadRequest) GetHeader() *RequestHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

type ReadReply struct {
	Error     *Error `protobuf:"bytes,1,opt,name=error" json:"error,omitempty"`
	BytesRead int64  `protobuf:"varint,2,opt,name=bytes_read" json:"bytes_read,omitempty"`
	Data      []byte `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
	Checksum  uint32 `protobuf:"fixed32,4,opt,name=checksum" json:"checksum,omitempty"`
}

func (m *ReadReply) Reset()         { *m = ReadReply{} }
func (m *ReadReply) String() string { return proto1.CompactTextString(m) }
func (*ReadReply) ProtoMessage()    {}

func (m *ReadReply) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

type RenameRequest struct {
	Header  *RequestHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	Oldname string         `protobuf:"bytes,2,opt,name=oldname" json:"oldname,omitempty"`
	Newname string         `protobuf:"bytes,3,opt,name=newname" json:"newname,omitempty"`
}

func (m *RenameRequest) Reset()         { *m = RenameRequest{} }
func (m *RenameRequest) String() string { return proto1.CompactTextString(m) }
func (*RenameRequest) ProtoMessage()    {}

func (m *RenameRequest) GetHeader() *RequestHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

type RenameReply struct {
	Error *Error `protobuf:"bytes,1,opt,name=error" json:"error,omitempty"`
}

func (m *RenameReply) Reset()         { *m = RenameReply{} }
func (m *RenameReply) String() string { return proto1.CompactTextString(m) }
func (*RenameReply) ProtoMessage()    {}

func (m *RenameReply) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

type ReadDirRequest struct {
	Header *RequestHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	Name   string         `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
}

func (m *ReadDirRequest) Reset()         { *m = ReadDirRequest{} }
func (m *ReadDirRequest) String() string { return proto1.CompactTextString(m) }
func (*ReadDirRequest) ProtoMessage()    {}

func (m *ReadDirRequest) GetHeader() *RequestHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

type ReadDirReply struct {
	Error     *Error      `protobuf:"bytes,1,opt,name=error" json:"error,omitempty"`
	FileInfos []*FileInfo `protobuf:"bytes,2,rep,name=fileInfos" json:"fileInfos,omitempty"`
}

func (m *ReadDirReply) Reset()         { *m = ReadDirReply{} }
func (m *ReadDirReply) String() string { return proto1.CompactTextString(m) }
func (*ReadDirReply) ProtoMessage()    {}

func (m *ReadDirReply) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

func (m *ReadDirReply) GetFileInfos() []*FileInfo {
	if m != nil {
		return m.FileInfos
	}
	return nil
}

// Remove removes the named file or directory. If there is an error, it will be of type *PathError.
type RemoveRequest struct {
	Header *RequestHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	Name   string         `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	// All removes path and any children it contains. It removes everything it can but returns the first error it
	// encounters. If the path does not exist, RemoveAll returns nil (no error).
	All bool `protobuf:"varint,3,opt,name=all" json:"all,omitempty"`
}

func (m *RemoveRequest) Reset()         { *m = RemoveRequest{} }
func (m *RemoveRequest) String() string { return proto1.CompactTextString(m) }
func (*RemoveRequest) ProtoMessage()    {}

func (m *RemoveRequest) GetHeader() *RequestHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

type RemoveReply struct {
	Error *Error `protobuf:"bytes,1,opt,name=error" json:"error,omitempty"`
}

func (m *RemoveReply) Reset()         { *m = RemoveReply{} }
func (m *RemoveReply) String() string { return proto1.CompactTextString(m) }
func (*RemoveReply) ProtoMessage()    {}

func (m *RemoveReply) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

// Mkdir creates a new directory with the specified name. If all is set, Mkdir creates a directory named path,
// along with any necessary parents. If path is already a directory, Mkdir does nothing.
type MkdirRequest struct {
	Header *RequestHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	Name   string         `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	All    bool           `protobuf:"varint,3,opt,name=all" json:"all,omitempty"`
}

func (m *MkdirRequest) Reset()         { *m = MkdirRequest{} }
func (m *MkdirRequest) String() string { return proto1.CompactTextString(m) }
func (*MkdirRequest) ProtoMessage()    {}

func (m *MkdirRequest) GetHeader() *RequestHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

type MkdirReply struct {
	Error *Error `protobuf:"bytes,1,opt,name=error" json:"error,omitempty"`
}

func (m *MkdirReply) Reset()         { *m = MkdirReply{} }
func (m *MkdirReply) String() string { return proto1.CompactTextString(m) }
func (*MkdirReply) ProtoMessage()    {}

func (m *MkdirReply) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

type ReconstructSrc struct {
	Header *RequestHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	Remote string         `protobuf:"bytes,2,opt,name=remote" json:"remote,omitempty"`
	Name   string         `protobuf:"bytes,3,opt,name=name" json:"name,omitempty"`
}

func (m *ReconstructSrc) Reset()         { *m = ReconstructSrc{} }
func (m *ReconstructSrc) String() string { return proto1.CompactTextString(m) }
func (*ReconstructSrc) ProtoMessage()    {}

func (m *ReconstructSrc) GetHeader() *RequestHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

type ReconstructDst struct {
	Header *RequestHeader `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	// The destination should always be local server.
	Name string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
}

func (m *ReconstructDst) Reset()         { *m = ReconstructDst{} }
func (m *ReconstructDst) String() string { return proto1.CompactTextString(m) }
func (*ReconstructDst) ProtoMessage()    {}

func (m *ReconstructDst) GetHeader() *RequestHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

// http://web.eecs.utk.edu/~plank/plank/papers/2013-02-11-FAST-Tutorial.pdf
// https://www.usenix.org/legacy/events/fast09/tech/full_papers/plank/plank_html/
// Optimized for Cauchy Reed-Solomon (CRS) Codes, but should also be applied to
// RAID5 and RAID6
type ReconstructRequest struct {
	Header *RequestHeader    `protobuf:"bytes,1,opt,name=header" json:"header,omitempty"`
	Srcs   []*ReconstructSrc `protobuf:"bytes,2,rep,name=srcs" json:"srcs,omitempty"`
	Dsts   []*ReconstructDst `protobuf:"bytes,3,rep,name=dsts" json:"dsts,omitempty"`
	// each src has multiple strips. the length of src must be
	// a multiply of stripe_size or it should be zero filled.
	//
	// a strip (also called block) is partitioned into w packets
	// Invariant: strip_size = packet_size * w
	// w MUST be in the range [1, 32]
	//
	// https://www.usenix.org/legacy/events/fast09/tech/full_papers/plank/plank_html Section 2.2
	StripSize  int32 `protobuf:"varint,4,opt,name=strip_size" json:"strip_size,omitempty"`
	PacketSize int32 `protobuf:"varint,5,opt,name=packet_size" json:"packet_size,omitempty"`
	W          int32 `protobuf:"varint,6,opt,name=w" json:"w,omitempty"`
	// wk * wn matrix of bits
	// k is the number of sources, n is the number of dests.
	// bit_matrix[i][j] = i * k * w + j
	// TODO: make this a dense bytes array and each bytes contains
	// 8 bits.
	BitMatrix []int32 `protobuf:"varint,7,rep,name=bit_matrix" json:"bit_matrix,omitempty"`
}

func (m *ReconstructRequest) Reset()         { *m = ReconstructRequest{} }
func (m *ReconstructRequest) String() string { return proto1.CompactTextString(m) }
func (*ReconstructRequest) ProtoMessage()    {}

func (m *ReconstructRequest) GetHeader() *RequestHeader {
	if m != nil {
		return m.Header
	}
	return nil
}

func (m *ReconstructRequest) GetSrcs() []*ReconstructSrc {
	if m != nil {
		return m.Srcs
	}
	return nil
}

func (m *ReconstructRequest) GetDsts() []*ReconstructDst {
	if m != nil {
		return m.Dsts
	}
	return nil
}

type ReconstructReply struct {
	Error *Error `protobuf:"bytes,1,opt,name=error" json:"error,omitempty"`
}

func (m *ReconstructReply) Reset()         { *m = ReconstructReply{} }
func (m *ReconstructReply) String() string { return proto1.CompactTextString(m) }
func (*ReconstructReply) ProtoMessage()    {}

func (m *ReconstructReply) GetError() *Error {
	if m != nil {
		return m.Error
	}
	return nil
}

func init() {
}

// Client API for Cfs service

type CfsClient interface {
	Write(ctx context.Context, in *WriteRequest, opts ...grpc.CallOption) (*WriteReply, error)
	Read(ctx context.Context, in *ReadRequest, opts ...grpc.CallOption) (*ReadReply, error)
	Rename(ctx context.Context, in *RenameRequest, opts ...grpc.CallOption) (*RenameReply, error)
	Remove(ctx context.Context, in *RemoveRequest, opts ...grpc.CallOption) (*RemoveReply, error)
	ReadDir(ctx context.Context, in *ReadDirRequest, opts ...grpc.CallOption) (*ReadDirReply, error)
	Mkdir(ctx context.Context, in *MkdirRequest, opts ...grpc.CallOption) (*MkdirReply, error)
}

type cfsClient struct {
	cc *grpc.ClientConn
}

func NewCfsClient(cc *grpc.ClientConn) CfsClient {
	return &cfsClient{cc}
}

func (c *cfsClient) Write(ctx context.Context, in *WriteRequest, opts ...grpc.CallOption) (*WriteReply, error) {
	out := new(WriteReply)
	err := grpc.Invoke(ctx, "/proto.cfs/Write", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cfsClient) Read(ctx context.Context, in *ReadRequest, opts ...grpc.CallOption) (*ReadReply, error) {
	out := new(ReadReply)
	err := grpc.Invoke(ctx, "/proto.cfs/Read", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cfsClient) Rename(ctx context.Context, in *RenameRequest, opts ...grpc.CallOption) (*RenameReply, error) {
	out := new(RenameReply)
	err := grpc.Invoke(ctx, "/proto.cfs/Rename", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cfsClient) Remove(ctx context.Context, in *RemoveRequest, opts ...grpc.CallOption) (*RemoveReply, error) {
	out := new(RemoveReply)
	err := grpc.Invoke(ctx, "/proto.cfs/Remove", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cfsClient) ReadDir(ctx context.Context, in *ReadDirRequest, opts ...grpc.CallOption) (*ReadDirReply, error) {
	out := new(ReadDirReply)
	err := grpc.Invoke(ctx, "/proto.cfs/ReadDir", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cfsClient) Mkdir(ctx context.Context, in *MkdirRequest, opts ...grpc.CallOption) (*MkdirReply, error) {
	out := new(MkdirReply)
	err := grpc.Invoke(ctx, "/proto.cfs/Mkdir", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Cfs service

type CfsServer interface {
	Write(context.Context, *WriteRequest) (*WriteReply, error)
	Read(context.Context, *ReadRequest) (*ReadReply, error)
	Rename(context.Context, *RenameRequest) (*RenameReply, error)
	Remove(context.Context, *RemoveRequest) (*RemoveReply, error)
	ReadDir(context.Context, *ReadDirRequest) (*ReadDirReply, error)
	Mkdir(context.Context, *MkdirRequest) (*MkdirReply, error)
}

func RegisterCfsServer(s *grpc.Server, srv CfsServer) {
	s.RegisterService(&_Cfs_serviceDesc, srv)
}

func _Cfs_Write_Handler(srv interface{}, ctx context.Context, codec grpc.Codec, buf []byte) (interface{}, error) {
	in := new(WriteRequest)
	if err := codec.Unmarshal(buf, in); err != nil {
		return nil, err
	}
	out, err := srv.(CfsServer).Write(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _Cfs_Read_Handler(srv interface{}, ctx context.Context, codec grpc.Codec, buf []byte) (interface{}, error) {
	in := new(ReadRequest)
	if err := codec.Unmarshal(buf, in); err != nil {
		return nil, err
	}
	out, err := srv.(CfsServer).Read(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _Cfs_Rename_Handler(srv interface{}, ctx context.Context, codec grpc.Codec, buf []byte) (interface{}, error) {
	in := new(RenameRequest)
	if err := codec.Unmarshal(buf, in); err != nil {
		return nil, err
	}
	out, err := srv.(CfsServer).Rename(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _Cfs_Remove_Handler(srv interface{}, ctx context.Context, codec grpc.Codec, buf []byte) (interface{}, error) {
	in := new(RemoveRequest)
	if err := codec.Unmarshal(buf, in); err != nil {
		return nil, err
	}
	out, err := srv.(CfsServer).Remove(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _Cfs_ReadDir_Handler(srv interface{}, ctx context.Context, codec grpc.Codec, buf []byte) (interface{}, error) {
	in := new(ReadDirRequest)
	if err := codec.Unmarshal(buf, in); err != nil {
		return nil, err
	}
	out, err := srv.(CfsServer).ReadDir(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _Cfs_Mkdir_Handler(srv interface{}, ctx context.Context, codec grpc.Codec, buf []byte) (interface{}, error) {
	in := new(MkdirRequest)
	if err := codec.Unmarshal(buf, in); err != nil {
		return nil, err
	}
	out, err := srv.(CfsServer).Mkdir(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var _Cfs_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.cfs",
	HandlerType: (*CfsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Write",
			Handler:    _Cfs_Write_Handler,
		},
		{
			MethodName: "Read",
			Handler:    _Cfs_Read_Handler,
		},
		{
			MethodName: "Rename",
			Handler:    _Cfs_Rename_Handler,
		},
		{
			MethodName: "Remove",
			Handler:    _Cfs_Remove_Handler,
		},
		{
			MethodName: "ReadDir",
			Handler:    _Cfs_ReadDir_Handler,
		},
		{
			MethodName: "Mkdir",
			Handler:    _Cfs_Mkdir_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}
