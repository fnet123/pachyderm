// Code generated by protoc-gen-go.
// source: server/pfs/db/persist/persist.proto
// DO NOT EDIT!

/*
Package persist is a generated protocol buffer package.

It is generated from these files:
	server/pfs/db/persist/persist.proto

It has these top-level messages:
	Clock
	ClockID
	BranchClock
	Repo
	Branch
	BlockRef
	Diff
	Commit
*/
package persist

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "go.pedge.io/pb/go/google/protobuf"
import _ "go.pedge.io/pb/go/google/protobuf"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type FileType int32

const (
	FileType_NONE FileType = 0
	FileType_FILE FileType = 1
	FileType_DIR  FileType = 2
)

var FileType_name = map[int32]string{
	0: "NONE",
	1: "FILE",
	2: "DIR",
}
var FileType_value = map[string]int32{
	"NONE": 0,
	"FILE": 1,
	"DIR":  2,
}

func (x FileType) String() string {
	return proto.EnumName(FileType_name, int32(x))
}
func (FileType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type Clock struct {
	// a document either has these two fields
	Branch string `protobuf:"bytes,1,opt,name=branch" json:"branch,omitempty"`
	Clock  uint64 `protobuf:"varint,2,opt,name=clock" json:"clock,omitempty"`
}

func (m *Clock) Reset()                    { *m = Clock{} }
func (m *Clock) String() string            { return proto.CompactTextString(m) }
func (*Clock) ProtoMessage()               {}
func (*Clock) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type ClockID struct {
	ID     string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Repo   string `protobuf:"bytes,2,opt,name=repo" json:"repo,omitempty"`
	Branch string `protobuf:"bytes,3,opt,name=branch" json:"branch,omitempty"`
	Clock  uint64 `protobuf:"varint,4,opt,name=clock" json:"clock,omitempty"`
}

func (m *ClockID) Reset()                    { *m = ClockID{} }
func (m *ClockID) String() string            { return proto.CompactTextString(m) }
func (*ClockID) ProtoMessage()               {}
func (*ClockID) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type BranchClock struct {
	Clocks []*Clock `protobuf:"bytes,1,rep,name=clocks" json:"clocks,omitempty"`
}

func (m *BranchClock) Reset()                    { *m = BranchClock{} }
func (m *BranchClock) String() string            { return proto.CompactTextString(m) }
func (*BranchClock) ProtoMessage()               {}
func (*BranchClock) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *BranchClock) GetClocks() []*Clock {
	if m != nil {
		return m.Clocks
	}
	return nil
}

type Repo struct {
	Name    string                     `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Created *google_protobuf.Timestamp `protobuf:"bytes,2,opt,name=created" json:"created,omitempty"`
	Size    uint64                     `protobuf:"varint,3,opt,name=size" json:"size,omitempty"`
}

func (m *Repo) Reset()                    { *m = Repo{} }
func (m *Repo) String() string            { return proto.CompactTextString(m) }
func (*Repo) ProtoMessage()               {}
func (*Repo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Repo) GetCreated() *google_protobuf.Timestamp {
	if m != nil {
		return m.Created
	}
	return nil
}

type Branch struct {
	ID   string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Repo string `protobuf:"bytes,2,opt,name=repo" json:"repo,omitempty"`
	Name string `protobuf:"bytes,3,opt,name=name" json:"name,omitempty"`
}

func (m *Branch) Reset()                    { *m = Branch{} }
func (m *Branch) String() string            { return proto.CompactTextString(m) }
func (*Branch) ProtoMessage()               {}
func (*Branch) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

type BlockRef struct {
	Hash  string `protobuf:"bytes,1,opt,name=hash" json:"hash,omitempty"`
	Lower uint64 `protobuf:"varint,2,opt,name=lower" json:"lower,omitempty"`
	Upper uint64 `protobuf:"varint,3,opt,name=upper" json:"upper,omitempty"`
}

func (m *BlockRef) Reset()                    { *m = BlockRef{} }
func (m *BlockRef) String() string            { return proto.CompactTextString(m) }
func (*BlockRef) ProtoMessage()               {}
func (*BlockRef) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

type Diff struct {
	ID       string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Repo     string `protobuf:"bytes,2,opt,name=repo" json:"repo,omitempty"`
	CommitID string `protobuf:"bytes,3,opt,name=commit_id,json=commitId" json:"commit_id,omitempty"`
	Path     string `protobuf:"bytes,4,opt,name=path" json:"path,omitempty"`
	// block_refs and delete cannot both be set
	BlockRefs    []*BlockRef                `protobuf:"bytes,5,rep,name=block_refs,json=blockRefs" json:"block_refs,omitempty"`
	Delete       bool                       `protobuf:"varint,6,opt,name=delete" json:"delete,omitempty"`
	Size         uint64                     `protobuf:"varint,7,opt,name=size" json:"size,omitempty"`
	BranchClocks []*BranchClock             `protobuf:"bytes,8,rep,name=branch_clocks,json=branchClocks" json:"branch_clocks,omitempty"`
	FileType     FileType                   `protobuf:"varint,9,opt,name=file_type,json=fileType,enum=FileType" json:"file_type,omitempty"`
	Modified     *google_protobuf.Timestamp `protobuf:"bytes,10,opt,name=modified" json:"modified,omitempty"`
}

func (m *Diff) Reset()                    { *m = Diff{} }
func (m *Diff) String() string            { return proto.CompactTextString(m) }
func (*Diff) ProtoMessage()               {}
func (*Diff) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *Diff) GetBlockRefs() []*BlockRef {
	if m != nil {
		return m.BlockRefs
	}
	return nil
}

func (m *Diff) GetBranchClocks() []*BranchClock {
	if m != nil {
		return m.BranchClocks
	}
	return nil
}

func (m *Diff) GetModified() *google_protobuf.Timestamp {
	if m != nil {
		return m.Modified
	}
	return nil
}

type Commit struct {
	ID           string                     `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Repo         string                     `protobuf:"bytes,2,opt,name=repo" json:"repo,omitempty"`
	BranchClocks []*BranchClock             `protobuf:"bytes,3,rep,name=branch_clocks,json=branchClocks" json:"branch_clocks,omitempty"`
	Started      *google_protobuf.Timestamp `protobuf:"bytes,4,opt,name=started" json:"started,omitempty"`
	Finished     *google_protobuf.Timestamp `protobuf:"bytes,5,opt,name=finished" json:"finished,omitempty"`
	Cancelled    bool                       `protobuf:"varint,6,opt,name=cancelled" json:"cancelled,omitempty"`
	Provenance   []string                   `protobuf:"bytes,7,rep,name=provenance" json:"provenance,omitempty"`
	Size         uint64                     `protobuf:"varint,8,opt,name=size" json:"size,omitempty"`
}

func (m *Commit) Reset()                    { *m = Commit{} }
func (m *Commit) String() string            { return proto.CompactTextString(m) }
func (*Commit) ProtoMessage()               {}
func (*Commit) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *Commit) GetBranchClocks() []*BranchClock {
	if m != nil {
		return m.BranchClocks
	}
	return nil
}

func (m *Commit) GetStarted() *google_protobuf.Timestamp {
	if m != nil {
		return m.Started
	}
	return nil
}

func (m *Commit) GetFinished() *google_protobuf.Timestamp {
	if m != nil {
		return m.Finished
	}
	return nil
}

func init() {
	proto.RegisterType((*Clock)(nil), "Clock")
	proto.RegisterType((*ClockID)(nil), "ClockID")
	proto.RegisterType((*BranchClock)(nil), "BranchClock")
	proto.RegisterType((*Repo)(nil), "Repo")
	proto.RegisterType((*Branch)(nil), "Branch")
	proto.RegisterType((*BlockRef)(nil), "BlockRef")
	proto.RegisterType((*Diff)(nil), "Diff")
	proto.RegisterType((*Commit)(nil), "Commit")
	proto.RegisterEnum("FileType", FileType_name, FileType_value)
}

func init() { proto.RegisterFile("server/pfs/db/persist/persist.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 549 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x53, 0x5d, 0x8b, 0xd3, 0x40,
	0x14, 0x75, 0x9b, 0x34, 0x4d, 0x6e, 0xd7, 0xa5, 0x0c, 0x22, 0x65, 0x95, 0x55, 0x22, 0x68, 0x11,
	0x4c, 0x71, 0xfd, 0x78, 0x96, 0xdd, 0xee, 0x42, 0x45, 0x56, 0x18, 0xf6, 0xcd, 0x87, 0x92, 0x34,
	0x93, 0xed, 0x60, 0xd2, 0x84, 0x99, 0xec, 0x2e, 0xfa, 0x67, 0xfc, 0x33, 0xfe, 0x30, 0xef, 0xdc,
	0x4c, 0x9a, 0xa2, 0x42, 0xfb, 0xd4, 0x7b, 0x4f, 0xee, 0xc7, 0xb9, 0xe7, 0x4c, 0xe1, 0x85, 0x16,
	0xea, 0x4e, 0xa8, 0x69, 0x95, 0xe9, 0x69, 0x9a, 0x4c, 0x2b, 0xa1, 0xb4, 0xd4, 0x75, 0xfb, 0x1b,
	0x55, 0xaa, 0xac, 0xcb, 0xe3, 0x67, 0x37, 0x65, 0x79, 0x93, 0x8b, 0x29, 0x65, 0xc9, 0x6d, 0x36,
	0xad, 0x65, 0x21, 0x74, 0x1d, 0x17, 0x95, 0x2d, 0x38, 0xf9, 0xbb, 0xe0, 0x5e, 0xc5, 0x95, 0x99,
	0xd1, 0x7c, 0x0f, 0x3f, 0x40, 0xff, 0x3c, 0x2f, 0x97, 0xdf, 0xd9, 0x63, 0xf0, 0x12, 0x15, 0xaf,
	0x97, 0xab, 0xf1, 0xc1, 0xf3, 0x83, 0x49, 0xc0, 0x6d, 0xc6, 0x1e, 0x41, 0x7f, 0x69, 0x0a, 0xc6,
	0x3d, 0x84, 0x5d, 0xde, 0x24, 0xe1, 0x37, 0x18, 0x50, 0xdb, 0x7c, 0xc6, 0x8e, 0xa0, 0x27, 0x53,
	0xdb, 0x84, 0x11, 0x63, 0xe0, 0x2a, 0x51, 0x95, 0x54, 0x1f, 0x70, 0x8a, 0xb7, 0x86, 0x3b, 0xff,
	0x1f, 0xee, 0x6e, 0x0f, 0x7f, 0x03, 0xc3, 0x33, 0xfa, 0xde, 0x30, 0x3b, 0x01, 0x8f, 0x70, 0x8d,
	0x4b, 0x9c, 0xc9, 0xf0, 0xd4, 0x8b, 0x08, 0xe7, 0x16, 0x0d, 0x53, 0x70, 0xb9, 0x59, 0x82, 0x8b,
	0xd7, 0x71, 0x21, 0x2c, 0x15, 0x8a, 0xd9, 0x7b, 0x18, 0x2c, 0x95, 0x88, 0x6b, 0x91, 0x12, 0x9f,
	0xe1, 0xe9, 0x71, 0xd4, 0x08, 0x12, 0xb5, 0x82, 0x44, 0xd7, 0xad, 0x62, 0xbc, 0x2d, 0x35, 0x93,
	0xb4, 0xfc, 0x29, 0x88, 0xac, 0xcb, 0x29, 0x0e, 0x3f, 0x81, 0xd7, 0x90, 0xda, 0xeb, 0xe0, 0x96,
	0x8b, 0xd3, 0x71, 0x09, 0x3f, 0x83, 0x7f, 0x46, 0xc4, 0x45, 0x66, 0xbe, 0xaf, 0x62, 0xdd, 0x6a,
	0x4d, 0xb1, 0x11, 0x23, 0x2f, 0xef, 0x85, 0x6a, 0x95, 0xa6, 0xc4, 0xa0, 0xb7, 0xc6, 0x30, 0x4b,
	0xa6, 0x49, 0xc2, 0xdf, 0x3d, 0x70, 0x67, 0x32, 0xcb, 0xf6, 0x22, 0xf3, 0x04, 0x82, 0x65, 0x59,
	0x14, 0xb2, 0x5e, 0x60, 0x69, 0xc3, 0xc8, 0x6f, 0x80, 0x39, 0x35, 0x54, 0x71, 0xbd, 0x22, 0x07,
	0xb0, 0xc1, 0xc4, 0x6c, 0x02, 0x90, 0x18, 0xa6, 0x0b, 0x25, 0x32, 0x3d, 0xee, 0x93, 0xea, 0x41,
	0xd4, 0x92, 0xe7, 0x41, 0x62, 0x23, 0x6d, 0x8c, 0x4d, 0x45, 0x2e, 0x6a, 0x31, 0xf6, 0xb0, 0xdf,
	0xe7, 0x36, 0xdb, 0x28, 0x38, 0xe8, 0x14, 0x64, 0x6f, 0xe1, 0x61, 0x63, 0xfb, 0xc2, 0xda, 0xe9,
	0xd3, 0xe0, 0xc3, 0x68, 0xcb, 0x6c, 0x7e, 0x98, 0x74, 0x89, 0x66, 0x2f, 0x21, 0xc8, 0x64, 0x2e,
	0x16, 0xf5, 0x8f, 0x4a, 0x8c, 0x03, 0x9c, 0x75, 0x84, 0x3c, 0x2e, 0x11, 0xb9, 0x46, 0x80, 0xfb,
	0x99, 0x8d, 0xd8, 0x47, 0xf0, 0x8b, 0x32, 0x95, 0x99, 0x44, 0x9f, 0x61, 0xa7, 0xcf, 0x9b, 0xda,
	0xf0, 0x57, 0x0f, 0xbc, 0x73, 0x52, 0x62, 0x2f, 0x21, 0xff, 0xb9, 0xc0, 0xd9, 0x79, 0x01, 0x3e,
	0x40, 0x5c, 0xaa, 0xcc, 0x03, 0x74, 0x77, 0x3f, 0x40, 0x5b, 0x6a, 0xee, 0xc9, 0xe4, 0x5a, 0xea,
	0x15, 0xb6, 0xf5, 0x77, 0xdf, 0xd3, 0xd6, 0xb2, 0xa7, 0xe8, 0x34, 0x2e, 0x17, 0x79, 0x8e, 0x8d,
	0x8d, 0x23, 0x1d, 0x80, 0x7f, 0x24, 0xc0, 0xee, 0x3b, 0xb1, 0x36, 0x08, 0x5a, 0xe3, 0xe0, 0x61,
	0x5b, 0xc8, 0xc6, 0x34, 0xbf, 0x33, 0xed, 0xf5, 0x2b, 0xf0, 0x5b, 0xbd, 0x99, 0x0f, 0xee, 0xd5,
	0xd7, 0xab, 0x8b, 0xd1, 0x03, 0x13, 0x5d, 0xce, 0xbf, 0x5c, 0x8c, 0x0e, 0xd8, 0x00, 0x9c, 0xd9,
	0x9c, 0x8f, 0x7a, 0x89, 0x47, 0xc4, 0xde, 0xfd, 0x09, 0x00, 0x00, 0xff, 0xff, 0xcf, 0xb9, 0x8e,
	0x04, 0xb7, 0x04, 0x00, 0x00,
}