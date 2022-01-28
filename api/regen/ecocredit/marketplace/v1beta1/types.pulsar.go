package marketplacev1beta1

import (
	fmt "fmt"
	runtime "github.com/cosmos/cosmos-proto/runtime"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoiface "google.golang.org/protobuf/runtime/protoiface"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	io "io"
	reflect "reflect"
	sync "sync"
)

var _ protoreflect.List = (*_Filter_1_list)(nil)

type _Filter_1_list struct {
	list *[]*Filter_Criteria
}

func (x *_Filter_1_list) Len() int {
	if x.list == nil {
		return 0
	}
	return len(*x.list)
}

func (x *_Filter_1_list) Get(i int) protoreflect.Value {
	return protoreflect.ValueOfMessage((*x.list)[i].ProtoReflect())
}

func (x *_Filter_1_list) Set(i int, value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*Filter_Criteria)
	(*x.list)[i] = concreteValue
}

func (x *_Filter_1_list) Append(value protoreflect.Value) {
	valueUnwrapped := value.Message()
	concreteValue := valueUnwrapped.Interface().(*Filter_Criteria)
	*x.list = append(*x.list, concreteValue)
}

func (x *_Filter_1_list) AppendMutable() protoreflect.Value {
	v := new(Filter_Criteria)
	*x.list = append(*x.list, v)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_Filter_1_list) Truncate(n int) {
	for i := n; i < len(*x.list); i++ {
		(*x.list)[i] = nil
	}
	*x.list = (*x.list)[:n]
}

func (x *_Filter_1_list) NewElement() protoreflect.Value {
	v := new(Filter_Criteria)
	return protoreflect.ValueOfMessage(v.ProtoReflect())
}

func (x *_Filter_1_list) IsValid() bool {
	return x.list != nil
}

var (
	md_Filter    protoreflect.MessageDescriptor
	fd_Filter_or protoreflect.FieldDescriptor
)

func init() {
	file_regen_ecocredit_marketplace_v1beta1_types_proto_init()
	md_Filter = File_regen_ecocredit_marketplace_v1beta1_types_proto.Messages().ByName("Filter")
	fd_Filter_or = md_Filter.Fields().ByName("or")
}

var _ protoreflect.Message = (*fastReflection_Filter)(nil)

type fastReflection_Filter Filter

func (x *Filter) ProtoReflect() protoreflect.Message {
	return (*fastReflection_Filter)(x)
}

func (x *Filter) slowProtoReflect() protoreflect.Message {
	mi := &file_regen_ecocredit_marketplace_v1beta1_types_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_Filter_messageType fastReflection_Filter_messageType
var _ protoreflect.MessageType = fastReflection_Filter_messageType{}

type fastReflection_Filter_messageType struct{}

func (x fastReflection_Filter_messageType) Zero() protoreflect.Message {
	return (*fastReflection_Filter)(nil)
}
func (x fastReflection_Filter_messageType) New() protoreflect.Message {
	return new(fastReflection_Filter)
}
func (x fastReflection_Filter_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_Filter
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_Filter) Descriptor() protoreflect.MessageDescriptor {
	return md_Filter
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_Filter) Type() protoreflect.MessageType {
	return _fastReflection_Filter_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_Filter) New() protoreflect.Message {
	return new(fastReflection_Filter)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_Filter) Interface() protoreflect.ProtoMessage {
	return (*Filter)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_Filter) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if len(x.Or) != 0 {
		value := protoreflect.ValueOfList(&_Filter_1_list{list: &x.Or})
		if !f(fd_Filter_or, value) {
			return
		}
	}
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x *fastReflection_Filter) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "regen.ecocredit.marketplace.v1beta1.Filter.or":
		return len(x.Or) != 0
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.marketplace.v1beta1.Filter"))
		}
		panic(fmt.Errorf("message regen.ecocredit.marketplace.v1beta1.Filter does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Filter) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "regen.ecocredit.marketplace.v1beta1.Filter.or":
		x.Or = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.marketplace.v1beta1.Filter"))
		}
		panic(fmt.Errorf("message regen.ecocredit.marketplace.v1beta1.Filter does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_Filter) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "regen.ecocredit.marketplace.v1beta1.Filter.or":
		if len(x.Or) == 0 {
			return protoreflect.ValueOfList(&_Filter_1_list{})
		}
		listValue := &_Filter_1_list{list: &x.Or}
		return protoreflect.ValueOfList(listValue)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.marketplace.v1beta1.Filter"))
		}
		panic(fmt.Errorf("message regen.ecocredit.marketplace.v1beta1.Filter does not contain field %s", descriptor.FullName()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Filter) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "regen.ecocredit.marketplace.v1beta1.Filter.or":
		lv := value.List()
		clv := lv.(*_Filter_1_list)
		x.Or = *clv.list
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.marketplace.v1beta1.Filter"))
		}
		panic(fmt.Errorf("message regen.ecocredit.marketplace.v1beta1.Filter does not contain field %s", fd.FullName()))
	}
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Filter) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.ecocredit.marketplace.v1beta1.Filter.or":
		if x.Or == nil {
			x.Or = []*Filter_Criteria{}
		}
		value := &_Filter_1_list{list: &x.Or}
		return protoreflect.ValueOfList(value)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.marketplace.v1beta1.Filter"))
		}
		panic(fmt.Errorf("message regen.ecocredit.marketplace.v1beta1.Filter does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_Filter) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.ecocredit.marketplace.v1beta1.Filter.or":
		list := []*Filter_Criteria{}
		return protoreflect.ValueOfList(&_Filter_1_list{list: &list})
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.marketplace.v1beta1.Filter"))
		}
		panic(fmt.Errorf("message regen.ecocredit.marketplace.v1beta1.Filter does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_Filter) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in regen.ecocredit.marketplace.v1beta1.Filter", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_Filter) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Filter) SetUnknown(fields protoreflect.RawFields) {
	x.unknownFields = fields
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x *fastReflection_Filter) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_Filter) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*Filter)
		if x == nil {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Size:              0,
			}
		}
		options := runtime.SizeInputToOptions(input)
		_ = options
		var n int
		var l int
		_ = l
		if len(x.Or) > 0 {
			for _, e := range x.Or {
				l = options.Size(e)
				n += 1 + l + runtime.Sov(uint64(l))
			}
		}
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*Filter)
		if x == nil {
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Buf:               input.Buf,
			}, nil
		}
		options := runtime.MarshalInputToOptions(input)
		_ = options
		size := options.Size(x)
		dAtA := make([]byte, size)
		i := len(dAtA)
		_ = i
		var l int
		_ = l
		if x.unknownFields != nil {
			i -= len(x.unknownFields)
			copy(dAtA[i:], x.unknownFields)
		}
		if len(x.Or) > 0 {
			for iNdEx := len(x.Or) - 1; iNdEx >= 0; iNdEx-- {
				encoded, err := options.Marshal(x.Or[iNdEx])
				if err != nil {
					return protoiface.MarshalOutput{
						NoUnkeyedLiterals: input.NoUnkeyedLiterals,
						Buf:               input.Buf,
					}, err
				}
				i -= len(encoded)
				copy(dAtA[i:], encoded)
				i = runtime.EncodeVarint(dAtA, i, uint64(len(encoded)))
				i--
				dAtA[i] = 0xa
			}
		}
		if input.Buf != nil {
			input.Buf = append(input.Buf, dAtA...)
		} else {
			input.Buf = dAtA
		}
		return protoiface.MarshalOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Buf:               input.Buf,
		}, nil
	}
	unmarshal := func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
		x := input.Message.Interface().(*Filter)
		if x == nil {
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Flags:             input.Flags,
			}, nil
		}
		options := runtime.UnmarshalInputToOptions(input)
		_ = options
		dAtA := input.Buf
		l := len(dAtA)
		iNdEx := 0
		for iNdEx < l {
			preIndex := iNdEx
			var wire uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
				}
				if iNdEx >= l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: Filter: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: Filter: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field Or", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.Or = append(x.Or, &Filter_Criteria{})
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.Or[len(x.Or)-1]); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			default:
				iNdEx = preIndex
				skippy, err := runtime.Skip(dAtA[iNdEx:])
				if err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				if (skippy < 0) || (iNdEx+skippy) < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if (iNdEx + skippy) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if !options.DiscardUnknown {
					x.unknownFields = append(x.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
				}
				iNdEx += skippy
			}
		}

		if iNdEx > l {
			return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
		}
		return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, nil
	}
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             protoiface.SupportMarshalDeterministic | protoiface.SupportUnmarshalDiscardUnknown,
		Size:              size,
		Marshal:           marshal,
		Unmarshal:         unmarshal,
		Merge:             nil,
		CheckInitialized:  nil,
	}
}

var (
	md_Filter_Criteria                  protoreflect.MessageDescriptor
	fd_Filter_Criteria_class_selector   protoreflect.FieldDescriptor
	fd_Filter_Criteria_project_selector protoreflect.FieldDescriptor
	fd_Filter_Criteria_batch_selector   protoreflect.FieldDescriptor
)

func init() {
	file_regen_ecocredit_marketplace_v1beta1_types_proto_init()
	md_Filter_Criteria = File_regen_ecocredit_marketplace_v1beta1_types_proto.Messages().ByName("Filter").Messages().ByName("Criteria")
	fd_Filter_Criteria_class_selector = md_Filter_Criteria.Fields().ByName("class_selector")
	fd_Filter_Criteria_project_selector = md_Filter_Criteria.Fields().ByName("project_selector")
	fd_Filter_Criteria_batch_selector = md_Filter_Criteria.Fields().ByName("batch_selector")
}

var _ protoreflect.Message = (*fastReflection_Filter_Criteria)(nil)

type fastReflection_Filter_Criteria Filter_Criteria

func (x *Filter_Criteria) ProtoReflect() protoreflect.Message {
	return (*fastReflection_Filter_Criteria)(x)
}

func (x *Filter_Criteria) slowProtoReflect() protoreflect.Message {
	mi := &file_regen_ecocredit_marketplace_v1beta1_types_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_Filter_Criteria_messageType fastReflection_Filter_Criteria_messageType
var _ protoreflect.MessageType = fastReflection_Filter_Criteria_messageType{}

type fastReflection_Filter_Criteria_messageType struct{}

func (x fastReflection_Filter_Criteria_messageType) Zero() protoreflect.Message {
	return (*fastReflection_Filter_Criteria)(nil)
}
func (x fastReflection_Filter_Criteria_messageType) New() protoreflect.Message {
	return new(fastReflection_Filter_Criteria)
}
func (x fastReflection_Filter_Criteria_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_Filter_Criteria
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_Filter_Criteria) Descriptor() protoreflect.MessageDescriptor {
	return md_Filter_Criteria
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_Filter_Criteria) Type() protoreflect.MessageType {
	return _fastReflection_Filter_Criteria_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_Filter_Criteria) New() protoreflect.Message {
	return new(fastReflection_Filter_Criteria)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_Filter_Criteria) Interface() protoreflect.ProtoMessage {
	return (*Filter_Criteria)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_Filter_Criteria) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.Selector != nil {
		switch o := x.Selector.(type) {
		case *Filter_Criteria_ClassSelector:
			v := o.ClassSelector
			value := protoreflect.ValueOfMessage(v.ProtoReflect())
			if !f(fd_Filter_Criteria_class_selector, value) {
				return
			}
		case *Filter_Criteria_ProjectSelector:
			v := o.ProjectSelector
			value := protoreflect.ValueOfMessage(v.ProtoReflect())
			if !f(fd_Filter_Criteria_project_selector, value) {
				return
			}
		case *Filter_Criteria_BatchSelector:
			v := o.BatchSelector
			value := protoreflect.ValueOfMessage(v.ProtoReflect())
			if !f(fd_Filter_Criteria_batch_selector, value) {
				return
			}
		}
	}
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x *fastReflection_Filter_Criteria) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "regen.ecocredit.marketplace.v1beta1.Filter.Criteria.class_selector":
		if x.Selector == nil {
			return false
		} else if _, ok := x.Selector.(*Filter_Criteria_ClassSelector); ok {
			return true
		} else {
			return false
		}
	case "regen.ecocredit.marketplace.v1beta1.Filter.Criteria.project_selector":
		if x.Selector == nil {
			return false
		} else if _, ok := x.Selector.(*Filter_Criteria_ProjectSelector); ok {
			return true
		} else {
			return false
		}
	case "regen.ecocredit.marketplace.v1beta1.Filter.Criteria.batch_selector":
		if x.Selector == nil {
			return false
		} else if _, ok := x.Selector.(*Filter_Criteria_BatchSelector); ok {
			return true
		} else {
			return false
		}
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.marketplace.v1beta1.Filter.Criteria"))
		}
		panic(fmt.Errorf("message regen.ecocredit.marketplace.v1beta1.Filter.Criteria does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Filter_Criteria) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "regen.ecocredit.marketplace.v1beta1.Filter.Criteria.class_selector":
		x.Selector = nil
	case "regen.ecocredit.marketplace.v1beta1.Filter.Criteria.project_selector":
		x.Selector = nil
	case "regen.ecocredit.marketplace.v1beta1.Filter.Criteria.batch_selector":
		x.Selector = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.marketplace.v1beta1.Filter.Criteria"))
		}
		panic(fmt.Errorf("message regen.ecocredit.marketplace.v1beta1.Filter.Criteria does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_Filter_Criteria) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "regen.ecocredit.marketplace.v1beta1.Filter.Criteria.class_selector":
		if x.Selector == nil {
			return protoreflect.ValueOfMessage((*ClassSelector)(nil).ProtoReflect())
		} else if v, ok := x.Selector.(*Filter_Criteria_ClassSelector); ok {
			return protoreflect.ValueOfMessage(v.ClassSelector.ProtoReflect())
		} else {
			return protoreflect.ValueOfMessage((*ClassSelector)(nil).ProtoReflect())
		}
	case "regen.ecocredit.marketplace.v1beta1.Filter.Criteria.project_selector":
		if x.Selector == nil {
			return protoreflect.ValueOfMessage((*ProjectSelector)(nil).ProtoReflect())
		} else if v, ok := x.Selector.(*Filter_Criteria_ProjectSelector); ok {
			return protoreflect.ValueOfMessage(v.ProjectSelector.ProtoReflect())
		} else {
			return protoreflect.ValueOfMessage((*ProjectSelector)(nil).ProtoReflect())
		}
	case "regen.ecocredit.marketplace.v1beta1.Filter.Criteria.batch_selector":
		if x.Selector == nil {
			return protoreflect.ValueOfMessage((*BatchSelector)(nil).ProtoReflect())
		} else if v, ok := x.Selector.(*Filter_Criteria_BatchSelector); ok {
			return protoreflect.ValueOfMessage(v.BatchSelector.ProtoReflect())
		} else {
			return protoreflect.ValueOfMessage((*BatchSelector)(nil).ProtoReflect())
		}
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.marketplace.v1beta1.Filter.Criteria"))
		}
		panic(fmt.Errorf("message regen.ecocredit.marketplace.v1beta1.Filter.Criteria does not contain field %s", descriptor.FullName()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Filter_Criteria) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "regen.ecocredit.marketplace.v1beta1.Filter.Criteria.class_selector":
		cv := value.Message().Interface().(*ClassSelector)
		x.Selector = &Filter_Criteria_ClassSelector{ClassSelector: cv}
	case "regen.ecocredit.marketplace.v1beta1.Filter.Criteria.project_selector":
		cv := value.Message().Interface().(*ProjectSelector)
		x.Selector = &Filter_Criteria_ProjectSelector{ProjectSelector: cv}
	case "regen.ecocredit.marketplace.v1beta1.Filter.Criteria.batch_selector":
		cv := value.Message().Interface().(*BatchSelector)
		x.Selector = &Filter_Criteria_BatchSelector{BatchSelector: cv}
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.marketplace.v1beta1.Filter.Criteria"))
		}
		panic(fmt.Errorf("message regen.ecocredit.marketplace.v1beta1.Filter.Criteria does not contain field %s", fd.FullName()))
	}
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Filter_Criteria) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.ecocredit.marketplace.v1beta1.Filter.Criteria.class_selector":
		if x.Selector == nil {
			value := &ClassSelector{}
			oneofValue := &Filter_Criteria_ClassSelector{ClassSelector: value}
			x.Selector = oneofValue
			return protoreflect.ValueOfMessage(value.ProtoReflect())
		}
		switch m := x.Selector.(type) {
		case *Filter_Criteria_ClassSelector:
			return protoreflect.ValueOfMessage(m.ClassSelector.ProtoReflect())
		default:
			value := &ClassSelector{}
			oneofValue := &Filter_Criteria_ClassSelector{ClassSelector: value}
			x.Selector = oneofValue
			return protoreflect.ValueOfMessage(value.ProtoReflect())
		}
	case "regen.ecocredit.marketplace.v1beta1.Filter.Criteria.project_selector":
		if x.Selector == nil {
			value := &ProjectSelector{}
			oneofValue := &Filter_Criteria_ProjectSelector{ProjectSelector: value}
			x.Selector = oneofValue
			return protoreflect.ValueOfMessage(value.ProtoReflect())
		}
		switch m := x.Selector.(type) {
		case *Filter_Criteria_ProjectSelector:
			return protoreflect.ValueOfMessage(m.ProjectSelector.ProtoReflect())
		default:
			value := &ProjectSelector{}
			oneofValue := &Filter_Criteria_ProjectSelector{ProjectSelector: value}
			x.Selector = oneofValue
			return protoreflect.ValueOfMessage(value.ProtoReflect())
		}
	case "regen.ecocredit.marketplace.v1beta1.Filter.Criteria.batch_selector":
		if x.Selector == nil {
			value := &BatchSelector{}
			oneofValue := &Filter_Criteria_BatchSelector{BatchSelector: value}
			x.Selector = oneofValue
			return protoreflect.ValueOfMessage(value.ProtoReflect())
		}
		switch m := x.Selector.(type) {
		case *Filter_Criteria_BatchSelector:
			return protoreflect.ValueOfMessage(m.BatchSelector.ProtoReflect())
		default:
			value := &BatchSelector{}
			oneofValue := &Filter_Criteria_BatchSelector{BatchSelector: value}
			x.Selector = oneofValue
			return protoreflect.ValueOfMessage(value.ProtoReflect())
		}
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.marketplace.v1beta1.Filter.Criteria"))
		}
		panic(fmt.Errorf("message regen.ecocredit.marketplace.v1beta1.Filter.Criteria does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_Filter_Criteria) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.ecocredit.marketplace.v1beta1.Filter.Criteria.class_selector":
		value := &ClassSelector{}
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	case "regen.ecocredit.marketplace.v1beta1.Filter.Criteria.project_selector":
		value := &ProjectSelector{}
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	case "regen.ecocredit.marketplace.v1beta1.Filter.Criteria.batch_selector":
		value := &BatchSelector{}
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.marketplace.v1beta1.Filter.Criteria"))
		}
		panic(fmt.Errorf("message regen.ecocredit.marketplace.v1beta1.Filter.Criteria does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_Filter_Criteria) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	case "regen.ecocredit.marketplace.v1beta1.Filter.Criteria.selector":
		if x.Selector == nil {
			return nil
		}
		switch x.Selector.(type) {
		case *Filter_Criteria_ClassSelector:
			return x.Descriptor().Fields().ByName("class_selector")
		case *Filter_Criteria_ProjectSelector:
			return x.Descriptor().Fields().ByName("project_selector")
		case *Filter_Criteria_BatchSelector:
			return x.Descriptor().Fields().ByName("batch_selector")
		}
	default:
		panic(fmt.Errorf("%s is not a oneof field in regen.ecocredit.marketplace.v1beta1.Filter.Criteria", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_Filter_Criteria) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_Filter_Criteria) SetUnknown(fields protoreflect.RawFields) {
	x.unknownFields = fields
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x *fastReflection_Filter_Criteria) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_Filter_Criteria) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*Filter_Criteria)
		if x == nil {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Size:              0,
			}
		}
		options := runtime.SizeInputToOptions(input)
		_ = options
		var n int
		var l int
		_ = l
		switch x := x.Selector.(type) {
		case *Filter_Criteria_ClassSelector:
			if x == nil {
				break
			}
			l = options.Size(x.ClassSelector)
			n += 1 + l + runtime.Sov(uint64(l))
		case *Filter_Criteria_ProjectSelector:
			if x == nil {
				break
			}
			l = options.Size(x.ProjectSelector)
			n += 1 + l + runtime.Sov(uint64(l))
		case *Filter_Criteria_BatchSelector:
			if x == nil {
				break
			}
			l = options.Size(x.BatchSelector)
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*Filter_Criteria)
		if x == nil {
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Buf:               input.Buf,
			}, nil
		}
		options := runtime.MarshalInputToOptions(input)
		_ = options
		size := options.Size(x)
		dAtA := make([]byte, size)
		i := len(dAtA)
		_ = i
		var l int
		_ = l
		if x.unknownFields != nil {
			i -= len(x.unknownFields)
			copy(dAtA[i:], x.unknownFields)
		}
		switch x := x.Selector.(type) {
		case *Filter_Criteria_ClassSelector:
			encoded, err := options.Marshal(x.ClassSelector)
			if err != nil {
				return protoiface.MarshalOutput{
					NoUnkeyedLiterals: input.NoUnkeyedLiterals,
					Buf:               input.Buf,
				}, err
			}
			i -= len(encoded)
			copy(dAtA[i:], encoded)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(encoded)))
			i--
			dAtA[i] = 0xa
		case *Filter_Criteria_ProjectSelector:
			encoded, err := options.Marshal(x.ProjectSelector)
			if err != nil {
				return protoiface.MarshalOutput{
					NoUnkeyedLiterals: input.NoUnkeyedLiterals,
					Buf:               input.Buf,
				}, err
			}
			i -= len(encoded)
			copy(dAtA[i:], encoded)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(encoded)))
			i--
			dAtA[i] = 0x12
		case *Filter_Criteria_BatchSelector:
			encoded, err := options.Marshal(x.BatchSelector)
			if err != nil {
				return protoiface.MarshalOutput{
					NoUnkeyedLiterals: input.NoUnkeyedLiterals,
					Buf:               input.Buf,
				}, err
			}
			i -= len(encoded)
			copy(dAtA[i:], encoded)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(encoded)))
			i--
			dAtA[i] = 0x1a
		}
		if input.Buf != nil {
			input.Buf = append(input.Buf, dAtA...)
		} else {
			input.Buf = dAtA
		}
		return protoiface.MarshalOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Buf:               input.Buf,
		}, nil
	}
	unmarshal := func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
		x := input.Message.Interface().(*Filter_Criteria)
		if x == nil {
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Flags:             input.Flags,
			}, nil
		}
		options := runtime.UnmarshalInputToOptions(input)
		_ = options
		dAtA := input.Buf
		l := len(dAtA)
		iNdEx := 0
		for iNdEx < l {
			preIndex := iNdEx
			var wire uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
				}
				if iNdEx >= l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: Filter_Criteria: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: Filter_Criteria: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field ClassSelector", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				v := &ClassSelector{}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], v); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				x.Selector = &Filter_Criteria_ClassSelector{v}
				iNdEx = postIndex
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field ProjectSelector", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				v := &ProjectSelector{}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], v); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				x.Selector = &Filter_Criteria_ProjectSelector{v}
				iNdEx = postIndex
			case 3:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field BatchSelector", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				v := &BatchSelector{}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], v); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				x.Selector = &Filter_Criteria_BatchSelector{v}
				iNdEx = postIndex
			default:
				iNdEx = preIndex
				skippy, err := runtime.Skip(dAtA[iNdEx:])
				if err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				if (skippy < 0) || (iNdEx+skippy) < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if (iNdEx + skippy) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if !options.DiscardUnknown {
					x.unknownFields = append(x.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
				}
				iNdEx += skippy
			}
		}

		if iNdEx > l {
			return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
		}
		return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, nil
	}
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             protoiface.SupportMarshalDeterministic | protoiface.SupportUnmarshalDiscardUnknown,
		Size:              size,
		Marshal:           marshal,
		Unmarshal:         unmarshal,
		Merge:             nil,
		CheckInitialized:  nil,
	}
}

var (
	md_ClassSelector                  protoreflect.MessageDescriptor
	fd_ClassSelector_class_id         protoreflect.FieldDescriptor
	fd_ClassSelector_project_location protoreflect.FieldDescriptor
	fd_ClassSelector_min_start_date   protoreflect.FieldDescriptor
	fd_ClassSelector_max_end_date     protoreflect.FieldDescriptor
)

func init() {
	file_regen_ecocredit_marketplace_v1beta1_types_proto_init()
	md_ClassSelector = File_regen_ecocredit_marketplace_v1beta1_types_proto.Messages().ByName("ClassSelector")
	fd_ClassSelector_class_id = md_ClassSelector.Fields().ByName("class_id")
	fd_ClassSelector_project_location = md_ClassSelector.Fields().ByName("project_location")
	fd_ClassSelector_min_start_date = md_ClassSelector.Fields().ByName("min_start_date")
	fd_ClassSelector_max_end_date = md_ClassSelector.Fields().ByName("max_end_date")
}

var _ protoreflect.Message = (*fastReflection_ClassSelector)(nil)

type fastReflection_ClassSelector ClassSelector

func (x *ClassSelector) ProtoReflect() protoreflect.Message {
	return (*fastReflection_ClassSelector)(x)
}

func (x *ClassSelector) slowProtoReflect() protoreflect.Message {
	mi := &file_regen_ecocredit_marketplace_v1beta1_types_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_ClassSelector_messageType fastReflection_ClassSelector_messageType
var _ protoreflect.MessageType = fastReflection_ClassSelector_messageType{}

type fastReflection_ClassSelector_messageType struct{}

func (x fastReflection_ClassSelector_messageType) Zero() protoreflect.Message {
	return (*fastReflection_ClassSelector)(nil)
}
func (x fastReflection_ClassSelector_messageType) New() protoreflect.Message {
	return new(fastReflection_ClassSelector)
}
func (x fastReflection_ClassSelector_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_ClassSelector
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_ClassSelector) Descriptor() protoreflect.MessageDescriptor {
	return md_ClassSelector
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_ClassSelector) Type() protoreflect.MessageType {
	return _fastReflection_ClassSelector_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_ClassSelector) New() protoreflect.Message {
	return new(fastReflection_ClassSelector)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_ClassSelector) Interface() protoreflect.ProtoMessage {
	return (*ClassSelector)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_ClassSelector) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.ClassId != uint64(0) {
		value := protoreflect.ValueOfUint64(x.ClassId)
		if !f(fd_ClassSelector_class_id, value) {
			return
		}
	}
	if x.ProjectLocation != "" {
		value := protoreflect.ValueOfString(x.ProjectLocation)
		if !f(fd_ClassSelector_project_location, value) {
			return
		}
	}
	if x.MinStartDate != nil {
		value := protoreflect.ValueOfMessage(x.MinStartDate.ProtoReflect())
		if !f(fd_ClassSelector_min_start_date, value) {
			return
		}
	}
	if x.MaxEndDate != nil {
		value := protoreflect.ValueOfMessage(x.MaxEndDate.ProtoReflect())
		if !f(fd_ClassSelector_max_end_date, value) {
			return
		}
	}
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x *fastReflection_ClassSelector) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "regen.ecocredit.marketplace.v1beta1.ClassSelector.class_id":
		return x.ClassId != uint64(0)
	case "regen.ecocredit.marketplace.v1beta1.ClassSelector.project_location":
		return x.ProjectLocation != ""
	case "regen.ecocredit.marketplace.v1beta1.ClassSelector.min_start_date":
		return x.MinStartDate != nil
	case "regen.ecocredit.marketplace.v1beta1.ClassSelector.max_end_date":
		return x.MaxEndDate != nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.marketplace.v1beta1.ClassSelector"))
		}
		panic(fmt.Errorf("message regen.ecocredit.marketplace.v1beta1.ClassSelector does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_ClassSelector) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "regen.ecocredit.marketplace.v1beta1.ClassSelector.class_id":
		x.ClassId = uint64(0)
	case "regen.ecocredit.marketplace.v1beta1.ClassSelector.project_location":
		x.ProjectLocation = ""
	case "regen.ecocredit.marketplace.v1beta1.ClassSelector.min_start_date":
		x.MinStartDate = nil
	case "regen.ecocredit.marketplace.v1beta1.ClassSelector.max_end_date":
		x.MaxEndDate = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.marketplace.v1beta1.ClassSelector"))
		}
		panic(fmt.Errorf("message regen.ecocredit.marketplace.v1beta1.ClassSelector does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_ClassSelector) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "regen.ecocredit.marketplace.v1beta1.ClassSelector.class_id":
		value := x.ClassId
		return protoreflect.ValueOfUint64(value)
	case "regen.ecocredit.marketplace.v1beta1.ClassSelector.project_location":
		value := x.ProjectLocation
		return protoreflect.ValueOfString(value)
	case "regen.ecocredit.marketplace.v1beta1.ClassSelector.min_start_date":
		value := x.MinStartDate
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	case "regen.ecocredit.marketplace.v1beta1.ClassSelector.max_end_date":
		value := x.MaxEndDate
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.marketplace.v1beta1.ClassSelector"))
		}
		panic(fmt.Errorf("message regen.ecocredit.marketplace.v1beta1.ClassSelector does not contain field %s", descriptor.FullName()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_ClassSelector) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "regen.ecocredit.marketplace.v1beta1.ClassSelector.class_id":
		x.ClassId = value.Uint()
	case "regen.ecocredit.marketplace.v1beta1.ClassSelector.project_location":
		x.ProjectLocation = value.Interface().(string)
	case "regen.ecocredit.marketplace.v1beta1.ClassSelector.min_start_date":
		x.MinStartDate = value.Message().Interface().(*timestamppb.Timestamp)
	case "regen.ecocredit.marketplace.v1beta1.ClassSelector.max_end_date":
		x.MaxEndDate = value.Message().Interface().(*timestamppb.Timestamp)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.marketplace.v1beta1.ClassSelector"))
		}
		panic(fmt.Errorf("message regen.ecocredit.marketplace.v1beta1.ClassSelector does not contain field %s", fd.FullName()))
	}
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_ClassSelector) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.ecocredit.marketplace.v1beta1.ClassSelector.min_start_date":
		if x.MinStartDate == nil {
			x.MinStartDate = new(timestamppb.Timestamp)
		}
		return protoreflect.ValueOfMessage(x.MinStartDate.ProtoReflect())
	case "regen.ecocredit.marketplace.v1beta1.ClassSelector.max_end_date":
		if x.MaxEndDate == nil {
			x.MaxEndDate = new(timestamppb.Timestamp)
		}
		return protoreflect.ValueOfMessage(x.MaxEndDate.ProtoReflect())
	case "regen.ecocredit.marketplace.v1beta1.ClassSelector.class_id":
		panic(fmt.Errorf("field class_id of message regen.ecocredit.marketplace.v1beta1.ClassSelector is not mutable"))
	case "regen.ecocredit.marketplace.v1beta1.ClassSelector.project_location":
		panic(fmt.Errorf("field project_location of message regen.ecocredit.marketplace.v1beta1.ClassSelector is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.marketplace.v1beta1.ClassSelector"))
		}
		panic(fmt.Errorf("message regen.ecocredit.marketplace.v1beta1.ClassSelector does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_ClassSelector) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.ecocredit.marketplace.v1beta1.ClassSelector.class_id":
		return protoreflect.ValueOfUint64(uint64(0))
	case "regen.ecocredit.marketplace.v1beta1.ClassSelector.project_location":
		return protoreflect.ValueOfString("")
	case "regen.ecocredit.marketplace.v1beta1.ClassSelector.min_start_date":
		m := new(timestamppb.Timestamp)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	case "regen.ecocredit.marketplace.v1beta1.ClassSelector.max_end_date":
		m := new(timestamppb.Timestamp)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.marketplace.v1beta1.ClassSelector"))
		}
		panic(fmt.Errorf("message regen.ecocredit.marketplace.v1beta1.ClassSelector does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_ClassSelector) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in regen.ecocredit.marketplace.v1beta1.ClassSelector", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_ClassSelector) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_ClassSelector) SetUnknown(fields protoreflect.RawFields) {
	x.unknownFields = fields
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x *fastReflection_ClassSelector) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_ClassSelector) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*ClassSelector)
		if x == nil {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Size:              0,
			}
		}
		options := runtime.SizeInputToOptions(input)
		_ = options
		var n int
		var l int
		_ = l
		if x.ClassId != 0 {
			n += 1 + runtime.Sov(uint64(x.ClassId))
		}
		l = len(x.ProjectLocation)
		if l > 0 {
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.MinStartDate != nil {
			l = options.Size(x.MinStartDate)
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.MaxEndDate != nil {
			l = options.Size(x.MaxEndDate)
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*ClassSelector)
		if x == nil {
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Buf:               input.Buf,
			}, nil
		}
		options := runtime.MarshalInputToOptions(input)
		_ = options
		size := options.Size(x)
		dAtA := make([]byte, size)
		i := len(dAtA)
		_ = i
		var l int
		_ = l
		if x.unknownFields != nil {
			i -= len(x.unknownFields)
			copy(dAtA[i:], x.unknownFields)
		}
		if x.MaxEndDate != nil {
			encoded, err := options.Marshal(x.MaxEndDate)
			if err != nil {
				return protoiface.MarshalOutput{
					NoUnkeyedLiterals: input.NoUnkeyedLiterals,
					Buf:               input.Buf,
				}, err
			}
			i -= len(encoded)
			copy(dAtA[i:], encoded)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(encoded)))
			i--
			dAtA[i] = 0x22
		}
		if x.MinStartDate != nil {
			encoded, err := options.Marshal(x.MinStartDate)
			if err != nil {
				return protoiface.MarshalOutput{
					NoUnkeyedLiterals: input.NoUnkeyedLiterals,
					Buf:               input.Buf,
				}, err
			}
			i -= len(encoded)
			copy(dAtA[i:], encoded)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(encoded)))
			i--
			dAtA[i] = 0x1a
		}
		if len(x.ProjectLocation) > 0 {
			i -= len(x.ProjectLocation)
			copy(dAtA[i:], x.ProjectLocation)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(x.ProjectLocation)))
			i--
			dAtA[i] = 0x12
		}
		if x.ClassId != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.ClassId))
			i--
			dAtA[i] = 0x8
		}
		if input.Buf != nil {
			input.Buf = append(input.Buf, dAtA...)
		} else {
			input.Buf = dAtA
		}
		return protoiface.MarshalOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Buf:               input.Buf,
		}, nil
	}
	unmarshal := func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
		x := input.Message.Interface().(*ClassSelector)
		if x == nil {
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Flags:             input.Flags,
			}, nil
		}
		options := runtime.UnmarshalInputToOptions(input)
		_ = options
		dAtA := input.Buf
		l := len(dAtA)
		iNdEx := 0
		for iNdEx < l {
			preIndex := iNdEx
			var wire uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
				}
				if iNdEx >= l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: ClassSelector: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: ClassSelector: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field ClassId", wireType)
				}
				x.ClassId = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.ClassId |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field ProjectLocation", wireType)
				}
				var stringLen uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					stringLen |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				intStringLen := int(stringLen)
				if intStringLen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + intStringLen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				x.ProjectLocation = string(dAtA[iNdEx:postIndex])
				iNdEx = postIndex
			case 3:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field MinStartDate", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if x.MinStartDate == nil {
					x.MinStartDate = &timestamppb.Timestamp{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.MinStartDate); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			case 4:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field MaxEndDate", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if x.MaxEndDate == nil {
					x.MaxEndDate = &timestamppb.Timestamp{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.MaxEndDate); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			default:
				iNdEx = preIndex
				skippy, err := runtime.Skip(dAtA[iNdEx:])
				if err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				if (skippy < 0) || (iNdEx+skippy) < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if (iNdEx + skippy) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if !options.DiscardUnknown {
					x.unknownFields = append(x.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
				}
				iNdEx += skippy
			}
		}

		if iNdEx > l {
			return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
		}
		return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, nil
	}
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             protoiface.SupportMarshalDeterministic | protoiface.SupportUnmarshalDiscardUnknown,
		Size:              size,
		Marshal:           marshal,
		Unmarshal:         unmarshal,
		Merge:             nil,
		CheckInitialized:  nil,
	}
}

var (
	md_ProjectSelector                protoreflect.MessageDescriptor
	fd_ProjectSelector_project_id     protoreflect.FieldDescriptor
	fd_ProjectSelector_min_start_date protoreflect.FieldDescriptor
	fd_ProjectSelector_max_end_date   protoreflect.FieldDescriptor
)

func init() {
	file_regen_ecocredit_marketplace_v1beta1_types_proto_init()
	md_ProjectSelector = File_regen_ecocredit_marketplace_v1beta1_types_proto.Messages().ByName("ProjectSelector")
	fd_ProjectSelector_project_id = md_ProjectSelector.Fields().ByName("project_id")
	fd_ProjectSelector_min_start_date = md_ProjectSelector.Fields().ByName("min_start_date")
	fd_ProjectSelector_max_end_date = md_ProjectSelector.Fields().ByName("max_end_date")
}

var _ protoreflect.Message = (*fastReflection_ProjectSelector)(nil)

type fastReflection_ProjectSelector ProjectSelector

func (x *ProjectSelector) ProtoReflect() protoreflect.Message {
	return (*fastReflection_ProjectSelector)(x)
}

func (x *ProjectSelector) slowProtoReflect() protoreflect.Message {
	mi := &file_regen_ecocredit_marketplace_v1beta1_types_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_ProjectSelector_messageType fastReflection_ProjectSelector_messageType
var _ protoreflect.MessageType = fastReflection_ProjectSelector_messageType{}

type fastReflection_ProjectSelector_messageType struct{}

func (x fastReflection_ProjectSelector_messageType) Zero() protoreflect.Message {
	return (*fastReflection_ProjectSelector)(nil)
}
func (x fastReflection_ProjectSelector_messageType) New() protoreflect.Message {
	return new(fastReflection_ProjectSelector)
}
func (x fastReflection_ProjectSelector_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_ProjectSelector
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_ProjectSelector) Descriptor() protoreflect.MessageDescriptor {
	return md_ProjectSelector
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_ProjectSelector) Type() protoreflect.MessageType {
	return _fastReflection_ProjectSelector_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_ProjectSelector) New() protoreflect.Message {
	return new(fastReflection_ProjectSelector)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_ProjectSelector) Interface() protoreflect.ProtoMessage {
	return (*ProjectSelector)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_ProjectSelector) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.ProjectId != uint64(0) {
		value := protoreflect.ValueOfUint64(x.ProjectId)
		if !f(fd_ProjectSelector_project_id, value) {
			return
		}
	}
	if x.MinStartDate != nil {
		value := protoreflect.ValueOfMessage(x.MinStartDate.ProtoReflect())
		if !f(fd_ProjectSelector_min_start_date, value) {
			return
		}
	}
	if x.MaxEndDate != nil {
		value := protoreflect.ValueOfMessage(x.MaxEndDate.ProtoReflect())
		if !f(fd_ProjectSelector_max_end_date, value) {
			return
		}
	}
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x *fastReflection_ProjectSelector) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "regen.ecocredit.marketplace.v1beta1.ProjectSelector.project_id":
		return x.ProjectId != uint64(0)
	case "regen.ecocredit.marketplace.v1beta1.ProjectSelector.min_start_date":
		return x.MinStartDate != nil
	case "regen.ecocredit.marketplace.v1beta1.ProjectSelector.max_end_date":
		return x.MaxEndDate != nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.marketplace.v1beta1.ProjectSelector"))
		}
		panic(fmt.Errorf("message regen.ecocredit.marketplace.v1beta1.ProjectSelector does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_ProjectSelector) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "regen.ecocredit.marketplace.v1beta1.ProjectSelector.project_id":
		x.ProjectId = uint64(0)
	case "regen.ecocredit.marketplace.v1beta1.ProjectSelector.min_start_date":
		x.MinStartDate = nil
	case "regen.ecocredit.marketplace.v1beta1.ProjectSelector.max_end_date":
		x.MaxEndDate = nil
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.marketplace.v1beta1.ProjectSelector"))
		}
		panic(fmt.Errorf("message regen.ecocredit.marketplace.v1beta1.ProjectSelector does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_ProjectSelector) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "regen.ecocredit.marketplace.v1beta1.ProjectSelector.project_id":
		value := x.ProjectId
		return protoreflect.ValueOfUint64(value)
	case "regen.ecocredit.marketplace.v1beta1.ProjectSelector.min_start_date":
		value := x.MinStartDate
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	case "regen.ecocredit.marketplace.v1beta1.ProjectSelector.max_end_date":
		value := x.MaxEndDate
		return protoreflect.ValueOfMessage(value.ProtoReflect())
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.marketplace.v1beta1.ProjectSelector"))
		}
		panic(fmt.Errorf("message regen.ecocredit.marketplace.v1beta1.ProjectSelector does not contain field %s", descriptor.FullName()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_ProjectSelector) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "regen.ecocredit.marketplace.v1beta1.ProjectSelector.project_id":
		x.ProjectId = value.Uint()
	case "regen.ecocredit.marketplace.v1beta1.ProjectSelector.min_start_date":
		x.MinStartDate = value.Message().Interface().(*timestamppb.Timestamp)
	case "regen.ecocredit.marketplace.v1beta1.ProjectSelector.max_end_date":
		x.MaxEndDate = value.Message().Interface().(*timestamppb.Timestamp)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.marketplace.v1beta1.ProjectSelector"))
		}
		panic(fmt.Errorf("message regen.ecocredit.marketplace.v1beta1.ProjectSelector does not contain field %s", fd.FullName()))
	}
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_ProjectSelector) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.ecocredit.marketplace.v1beta1.ProjectSelector.min_start_date":
		if x.MinStartDate == nil {
			x.MinStartDate = new(timestamppb.Timestamp)
		}
		return protoreflect.ValueOfMessage(x.MinStartDate.ProtoReflect())
	case "regen.ecocredit.marketplace.v1beta1.ProjectSelector.max_end_date":
		if x.MaxEndDate == nil {
			x.MaxEndDate = new(timestamppb.Timestamp)
		}
		return protoreflect.ValueOfMessage(x.MaxEndDate.ProtoReflect())
	case "regen.ecocredit.marketplace.v1beta1.ProjectSelector.project_id":
		panic(fmt.Errorf("field project_id of message regen.ecocredit.marketplace.v1beta1.ProjectSelector is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.marketplace.v1beta1.ProjectSelector"))
		}
		panic(fmt.Errorf("message regen.ecocredit.marketplace.v1beta1.ProjectSelector does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_ProjectSelector) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.ecocredit.marketplace.v1beta1.ProjectSelector.project_id":
		return protoreflect.ValueOfUint64(uint64(0))
	case "regen.ecocredit.marketplace.v1beta1.ProjectSelector.min_start_date":
		m := new(timestamppb.Timestamp)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	case "regen.ecocredit.marketplace.v1beta1.ProjectSelector.max_end_date":
		m := new(timestamppb.Timestamp)
		return protoreflect.ValueOfMessage(m.ProtoReflect())
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.marketplace.v1beta1.ProjectSelector"))
		}
		panic(fmt.Errorf("message regen.ecocredit.marketplace.v1beta1.ProjectSelector does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_ProjectSelector) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in regen.ecocredit.marketplace.v1beta1.ProjectSelector", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_ProjectSelector) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_ProjectSelector) SetUnknown(fields protoreflect.RawFields) {
	x.unknownFields = fields
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x *fastReflection_ProjectSelector) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_ProjectSelector) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*ProjectSelector)
		if x == nil {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Size:              0,
			}
		}
		options := runtime.SizeInputToOptions(input)
		_ = options
		var n int
		var l int
		_ = l
		if x.ProjectId != 0 {
			n += 1 + runtime.Sov(uint64(x.ProjectId))
		}
		if x.MinStartDate != nil {
			l = options.Size(x.MinStartDate)
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.MaxEndDate != nil {
			l = options.Size(x.MaxEndDate)
			n += 1 + l + runtime.Sov(uint64(l))
		}
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*ProjectSelector)
		if x == nil {
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Buf:               input.Buf,
			}, nil
		}
		options := runtime.MarshalInputToOptions(input)
		_ = options
		size := options.Size(x)
		dAtA := make([]byte, size)
		i := len(dAtA)
		_ = i
		var l int
		_ = l
		if x.unknownFields != nil {
			i -= len(x.unknownFields)
			copy(dAtA[i:], x.unknownFields)
		}
		if x.MaxEndDate != nil {
			encoded, err := options.Marshal(x.MaxEndDate)
			if err != nil {
				return protoiface.MarshalOutput{
					NoUnkeyedLiterals: input.NoUnkeyedLiterals,
					Buf:               input.Buf,
				}, err
			}
			i -= len(encoded)
			copy(dAtA[i:], encoded)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(encoded)))
			i--
			dAtA[i] = 0x1a
		}
		if x.MinStartDate != nil {
			encoded, err := options.Marshal(x.MinStartDate)
			if err != nil {
				return protoiface.MarshalOutput{
					NoUnkeyedLiterals: input.NoUnkeyedLiterals,
					Buf:               input.Buf,
				}, err
			}
			i -= len(encoded)
			copy(dAtA[i:], encoded)
			i = runtime.EncodeVarint(dAtA, i, uint64(len(encoded)))
			i--
			dAtA[i] = 0x12
		}
		if x.ProjectId != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.ProjectId))
			i--
			dAtA[i] = 0x8
		}
		if input.Buf != nil {
			input.Buf = append(input.Buf, dAtA...)
		} else {
			input.Buf = dAtA
		}
		return protoiface.MarshalOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Buf:               input.Buf,
		}, nil
	}
	unmarshal := func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
		x := input.Message.Interface().(*ProjectSelector)
		if x == nil {
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Flags:             input.Flags,
			}, nil
		}
		options := runtime.UnmarshalInputToOptions(input)
		_ = options
		dAtA := input.Buf
		l := len(dAtA)
		iNdEx := 0
		for iNdEx < l {
			preIndex := iNdEx
			var wire uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
				}
				if iNdEx >= l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: ProjectSelector: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: ProjectSelector: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field ProjectId", wireType)
				}
				x.ProjectId = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.ProjectId |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			case 2:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field MinStartDate", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if x.MinStartDate == nil {
					x.MinStartDate = &timestamppb.Timestamp{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.MinStartDate); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			case 3:
				if wireType != 2 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field MaxEndDate", wireType)
				}
				var msglen int
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					msglen |= int(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				if msglen < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				postIndex := iNdEx + msglen
				if postIndex < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if postIndex > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if x.MaxEndDate == nil {
					x.MaxEndDate = &timestamppb.Timestamp{}
				}
				if err := options.Unmarshal(dAtA[iNdEx:postIndex], x.MaxEndDate); err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				iNdEx = postIndex
			default:
				iNdEx = preIndex
				skippy, err := runtime.Skip(dAtA[iNdEx:])
				if err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				if (skippy < 0) || (iNdEx+skippy) < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if (iNdEx + skippy) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if !options.DiscardUnknown {
					x.unknownFields = append(x.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
				}
				iNdEx += skippy
			}
		}

		if iNdEx > l {
			return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
		}
		return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, nil
	}
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             protoiface.SupportMarshalDeterministic | protoiface.SupportUnmarshalDiscardUnknown,
		Size:              size,
		Marshal:           marshal,
		Unmarshal:         unmarshal,
		Merge:             nil,
		CheckInitialized:  nil,
	}
}

var (
	md_BatchSelector          protoreflect.MessageDescriptor
	fd_BatchSelector_batch_id protoreflect.FieldDescriptor
)

func init() {
	file_regen_ecocredit_marketplace_v1beta1_types_proto_init()
	md_BatchSelector = File_regen_ecocredit_marketplace_v1beta1_types_proto.Messages().ByName("BatchSelector")
	fd_BatchSelector_batch_id = md_BatchSelector.Fields().ByName("batch_id")
}

var _ protoreflect.Message = (*fastReflection_BatchSelector)(nil)

type fastReflection_BatchSelector BatchSelector

func (x *BatchSelector) ProtoReflect() protoreflect.Message {
	return (*fastReflection_BatchSelector)(x)
}

func (x *BatchSelector) slowProtoReflect() protoreflect.Message {
	mi := &file_regen_ecocredit_marketplace_v1beta1_types_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

var _fastReflection_BatchSelector_messageType fastReflection_BatchSelector_messageType
var _ protoreflect.MessageType = fastReflection_BatchSelector_messageType{}

type fastReflection_BatchSelector_messageType struct{}

func (x fastReflection_BatchSelector_messageType) Zero() protoreflect.Message {
	return (*fastReflection_BatchSelector)(nil)
}
func (x fastReflection_BatchSelector_messageType) New() protoreflect.Message {
	return new(fastReflection_BatchSelector)
}
func (x fastReflection_BatchSelector_messageType) Descriptor() protoreflect.MessageDescriptor {
	return md_BatchSelector
}

// Descriptor returns message descriptor, which contains only the protobuf
// type information for the message.
func (x *fastReflection_BatchSelector) Descriptor() protoreflect.MessageDescriptor {
	return md_BatchSelector
}

// Type returns the message type, which encapsulates both Go and protobuf
// type information. If the Go type information is not needed,
// it is recommended that the message descriptor be used instead.
func (x *fastReflection_BatchSelector) Type() protoreflect.MessageType {
	return _fastReflection_BatchSelector_messageType
}

// New returns a newly allocated and mutable empty message.
func (x *fastReflection_BatchSelector) New() protoreflect.Message {
	return new(fastReflection_BatchSelector)
}

// Interface unwraps the message reflection interface and
// returns the underlying ProtoMessage interface.
func (x *fastReflection_BatchSelector) Interface() protoreflect.ProtoMessage {
	return (*BatchSelector)(x)
}

// Range iterates over every populated field in an undefined order,
// calling f for each field descriptor and value encountered.
// Range returns immediately if f returns false.
// While iterating, mutating operations may only be performed
// on the current field descriptor.
func (x *fastReflection_BatchSelector) Range(f func(protoreflect.FieldDescriptor, protoreflect.Value) bool) {
	if x.BatchId != uint64(0) {
		value := protoreflect.ValueOfUint64(x.BatchId)
		if !f(fd_BatchSelector_batch_id, value) {
			return
		}
	}
}

// Has reports whether a field is populated.
//
// Some fields have the property of nullability where it is possible to
// distinguish between the default value of a field and whether the field
// was explicitly populated with the default value. Singular message fields,
// member fields of a oneof, and proto2 scalar fields are nullable. Such
// fields are populated only if explicitly set.
//
// In other cases (aside from the nullable cases above),
// a proto3 scalar field is populated if it contains a non-zero value, and
// a repeated field is populated if it is non-empty.
func (x *fastReflection_BatchSelector) Has(fd protoreflect.FieldDescriptor) bool {
	switch fd.FullName() {
	case "regen.ecocredit.marketplace.v1beta1.BatchSelector.batch_id":
		return x.BatchId != uint64(0)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.marketplace.v1beta1.BatchSelector"))
		}
		panic(fmt.Errorf("message regen.ecocredit.marketplace.v1beta1.BatchSelector does not contain field %s", fd.FullName()))
	}
}

// Clear clears the field such that a subsequent Has call reports false.
//
// Clearing an extension field clears both the extension type and value
// associated with the given field number.
//
// Clear is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_BatchSelector) Clear(fd protoreflect.FieldDescriptor) {
	switch fd.FullName() {
	case "regen.ecocredit.marketplace.v1beta1.BatchSelector.batch_id":
		x.BatchId = uint64(0)
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.marketplace.v1beta1.BatchSelector"))
		}
		panic(fmt.Errorf("message regen.ecocredit.marketplace.v1beta1.BatchSelector does not contain field %s", fd.FullName()))
	}
}

// Get retrieves the value for a field.
//
// For unpopulated scalars, it returns the default value, where
// the default value of a bytes scalar is guaranteed to be a copy.
// For unpopulated composite types, it returns an empty, read-only view
// of the value; to obtain a mutable reference, use Mutable.
func (x *fastReflection_BatchSelector) Get(descriptor protoreflect.FieldDescriptor) protoreflect.Value {
	switch descriptor.FullName() {
	case "regen.ecocredit.marketplace.v1beta1.BatchSelector.batch_id":
		value := x.BatchId
		return protoreflect.ValueOfUint64(value)
	default:
		if descriptor.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.marketplace.v1beta1.BatchSelector"))
		}
		panic(fmt.Errorf("message regen.ecocredit.marketplace.v1beta1.BatchSelector does not contain field %s", descriptor.FullName()))
	}
}

// Set stores the value for a field.
//
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType.
// When setting a composite type, it is unspecified whether the stored value
// aliases the source's memory in any way. If the composite value is an
// empty, read-only value, then it panics.
//
// Set is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_BatchSelector) Set(fd protoreflect.FieldDescriptor, value protoreflect.Value) {
	switch fd.FullName() {
	case "regen.ecocredit.marketplace.v1beta1.BatchSelector.batch_id":
		x.BatchId = value.Uint()
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.marketplace.v1beta1.BatchSelector"))
		}
		panic(fmt.Errorf("message regen.ecocredit.marketplace.v1beta1.BatchSelector does not contain field %s", fd.FullName()))
	}
}

// Mutable returns a mutable reference to a composite type.
//
// If the field is unpopulated, it may allocate a composite value.
// For a field belonging to a oneof, it implicitly clears any other field
// that may be currently set within the same oneof.
// For extension fields, it implicitly stores the provided ExtensionType
// if not already stored.
// It panics if the field does not contain a composite type.
//
// Mutable is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_BatchSelector) Mutable(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.ecocredit.marketplace.v1beta1.BatchSelector.batch_id":
		panic(fmt.Errorf("field batch_id of message regen.ecocredit.marketplace.v1beta1.BatchSelector is not mutable"))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.marketplace.v1beta1.BatchSelector"))
		}
		panic(fmt.Errorf("message regen.ecocredit.marketplace.v1beta1.BatchSelector does not contain field %s", fd.FullName()))
	}
}

// NewField returns a new value that is assignable to the field
// for the given descriptor. For scalars, this returns the default value.
// For lists, maps, and messages, this returns a new, empty, mutable value.
func (x *fastReflection_BatchSelector) NewField(fd protoreflect.FieldDescriptor) protoreflect.Value {
	switch fd.FullName() {
	case "regen.ecocredit.marketplace.v1beta1.BatchSelector.batch_id":
		return protoreflect.ValueOfUint64(uint64(0))
	default:
		if fd.IsExtension() {
			panic(fmt.Errorf("proto3 declared messages do not support extensions: regen.ecocredit.marketplace.v1beta1.BatchSelector"))
		}
		panic(fmt.Errorf("message regen.ecocredit.marketplace.v1beta1.BatchSelector does not contain field %s", fd.FullName()))
	}
}

// WhichOneof reports which field within the oneof is populated,
// returning nil if none are populated.
// It panics if the oneof descriptor does not belong to this message.
func (x *fastReflection_BatchSelector) WhichOneof(d protoreflect.OneofDescriptor) protoreflect.FieldDescriptor {
	switch d.FullName() {
	default:
		panic(fmt.Errorf("%s is not a oneof field in regen.ecocredit.marketplace.v1beta1.BatchSelector", d.FullName()))
	}
	panic("unreachable")
}

// GetUnknown retrieves the entire list of unknown fields.
// The caller may only mutate the contents of the RawFields
// if the mutated bytes are stored back into the message with SetUnknown.
func (x *fastReflection_BatchSelector) GetUnknown() protoreflect.RawFields {
	return x.unknownFields
}

// SetUnknown stores an entire list of unknown fields.
// The raw fields must be syntactically valid according to the wire format.
// An implementation may panic if this is not the case.
// Once stored, the caller must not mutate the content of the RawFields.
// An empty RawFields may be passed to clear the fields.
//
// SetUnknown is a mutating operation and unsafe for concurrent use.
func (x *fastReflection_BatchSelector) SetUnknown(fields protoreflect.RawFields) {
	x.unknownFields = fields
}

// IsValid reports whether the message is valid.
//
// An invalid message is an empty, read-only value.
//
// An invalid message often corresponds to a nil pointer of the concrete
// message type, but the details are implementation dependent.
// Validity is not part of the protobuf data model, and may not
// be preserved in marshaling or other operations.
func (x *fastReflection_BatchSelector) IsValid() bool {
	return x != nil
}

// ProtoMethods returns optional fastReflectionFeature-path implementations of various operations.
// This method may return nil.
//
// The returned methods type is identical to
// "google.golang.org/protobuf/runtime/protoiface".Methods.
// Consult the protoiface package documentation for details.
func (x *fastReflection_BatchSelector) ProtoMethods() *protoiface.Methods {
	size := func(input protoiface.SizeInput) protoiface.SizeOutput {
		x := input.Message.Interface().(*BatchSelector)
		if x == nil {
			return protoiface.SizeOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Size:              0,
			}
		}
		options := runtime.SizeInputToOptions(input)
		_ = options
		var n int
		var l int
		_ = l
		if x.BatchId != 0 {
			n += 1 + runtime.Sov(uint64(x.BatchId))
		}
		if x.unknownFields != nil {
			n += len(x.unknownFields)
		}
		return protoiface.SizeOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Size:              n,
		}
	}

	marshal := func(input protoiface.MarshalInput) (protoiface.MarshalOutput, error) {
		x := input.Message.Interface().(*BatchSelector)
		if x == nil {
			return protoiface.MarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Buf:               input.Buf,
			}, nil
		}
		options := runtime.MarshalInputToOptions(input)
		_ = options
		size := options.Size(x)
		dAtA := make([]byte, size)
		i := len(dAtA)
		_ = i
		var l int
		_ = l
		if x.unknownFields != nil {
			i -= len(x.unknownFields)
			copy(dAtA[i:], x.unknownFields)
		}
		if x.BatchId != 0 {
			i = runtime.EncodeVarint(dAtA, i, uint64(x.BatchId))
			i--
			dAtA[i] = 0x8
		}
		if input.Buf != nil {
			input.Buf = append(input.Buf, dAtA...)
		} else {
			input.Buf = dAtA
		}
		return protoiface.MarshalOutput{
			NoUnkeyedLiterals: input.NoUnkeyedLiterals,
			Buf:               input.Buf,
		}, nil
	}
	unmarshal := func(input protoiface.UnmarshalInput) (protoiface.UnmarshalOutput, error) {
		x := input.Message.Interface().(*BatchSelector)
		if x == nil {
			return protoiface.UnmarshalOutput{
				NoUnkeyedLiterals: input.NoUnkeyedLiterals,
				Flags:             input.Flags,
			}, nil
		}
		options := runtime.UnmarshalInputToOptions(input)
		_ = options
		dAtA := input.Buf
		l := len(dAtA)
		iNdEx := 0
		for iNdEx < l {
			preIndex := iNdEx
			var wire uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
				}
				if iNdEx >= l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
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
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: BatchSelector: wiretype end group for non-group")
			}
			if fieldNum <= 0 {
				return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: BatchSelector: illegal tag %d (wire type %d)", fieldNum, wire)
			}
			switch fieldNum {
			case 1:
				if wireType != 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, fmt.Errorf("proto: wrong wireType = %d for field BatchId", wireType)
				}
				x.BatchId = 0
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrIntOverflow
					}
					if iNdEx >= l {
						return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					x.BatchId |= uint64(b&0x7F) << shift
					if b < 0x80 {
						break
					}
				}
			default:
				iNdEx = preIndex
				skippy, err := runtime.Skip(dAtA[iNdEx:])
				if err != nil {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, err
				}
				if (skippy < 0) || (iNdEx+skippy) < 0 {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, runtime.ErrInvalidLength
				}
				if (iNdEx + skippy) > l {
					return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
				}
				if !options.DiscardUnknown {
					x.unknownFields = append(x.unknownFields, dAtA[iNdEx:iNdEx+skippy]...)
				}
				iNdEx += skippy
			}
		}

		if iNdEx > l {
			return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, io.ErrUnexpectedEOF
		}
		return protoiface.UnmarshalOutput{NoUnkeyedLiterals: input.NoUnkeyedLiterals, Flags: input.Flags}, nil
	}
	return &protoiface.Methods{
		NoUnkeyedLiterals: struct{}{},
		Flags:             protoiface.SupportMarshalDeterministic | protoiface.SupportUnmarshalDiscardUnknown,
		Size:              size,
		Marshal:           marshal,
		Unmarshal:         unmarshal,
		Merge:             nil,
		CheckInitialized:  nil,
	}
}

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.0
// 	protoc        (unknown)
// source: regen/ecocredit/marketplace/v1beta1/types.proto

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Filter is used to create filtered buy orders which match credit batch
// sell orders based on selection criteria rather than matching individual
// sell orders
type Filter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// or is a list of criteria for matching credit batches. A credit which
	// matches this filter must match at least one of these criteria.
	Or []*Filter_Criteria `protobuf:"bytes,1,rep,name=or,proto3" json:"or,omitempty"`
}

func (x *Filter) Reset() {
	*x = Filter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_regen_ecocredit_marketplace_v1beta1_types_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Filter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Filter) ProtoMessage() {}

// Deprecated: Use Filter.ProtoReflect.Descriptor instead.
func (*Filter) Descriptor() ([]byte, []int) {
	return file_regen_ecocredit_marketplace_v1beta1_types_proto_rawDescGZIP(), []int{0}
}

func (x *Filter) GetOr() []*Filter_Criteria {
	if x != nil {
		return x.Or
	}
	return nil
}

type ClassSelector struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ClassId uint64 `protobuf:"varint,1,opt,name=class_id,json=classId,proto3" json:"class_id,omitempty"`
	// project_location can be specified in three levels of granularity:
	// country, sub-national-code, or postal code. If just country is given,
	// for instance "US" then any credits in the "US" will be matched even
	// their project location is more specific, ex. "US-NY 12345". If
	// a country, sub-national-code and postal code are all provided then
	// only projects in that postal code will match.
	ProjectLocation string `protobuf:"bytes,2,opt,name=project_location,json=projectLocation,proto3" json:"project_location,omitempty"`
	// start_date is the beginning of the period during which a credit batch
	// was quantified and verified. If it is empty then there is no start date
	// limit.
	MinStartDate *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=min_start_date,json=minStartDate,proto3" json:"min_start_date,omitempty"`
	// max_end_date is the end of the period during which a credit batch was
	// quantified and verified. If it is empty then there is no end date
	// limit.
	MaxEndDate *timestamppb.Timestamp `protobuf:"bytes,4,opt,name=max_end_date,json=maxEndDate,proto3" json:"max_end_date,omitempty"`
}

func (x *ClassSelector) Reset() {
	*x = ClassSelector{}
	if protoimpl.UnsafeEnabled {
		mi := &file_regen_ecocredit_marketplace_v1beta1_types_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ClassSelector) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClassSelector) ProtoMessage() {}

// Deprecated: Use ClassSelector.ProtoReflect.Descriptor instead.
func (*ClassSelector) Descriptor() ([]byte, []int) {
	return file_regen_ecocredit_marketplace_v1beta1_types_proto_rawDescGZIP(), []int{1}
}

func (x *ClassSelector) GetClassId() uint64 {
	if x != nil {
		return x.ClassId
	}
	return 0
}

func (x *ClassSelector) GetProjectLocation() string {
	if x != nil {
		return x.ProjectLocation
	}
	return ""
}

func (x *ClassSelector) GetMinStartDate() *timestamppb.Timestamp {
	if x != nil {
		return x.MinStartDate
	}
	return nil
}

func (x *ClassSelector) GetMaxEndDate() *timestamppb.Timestamp {
	if x != nil {
		return x.MaxEndDate
	}
	return nil
}

type ProjectSelector struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProjectId uint64 `protobuf:"varint,1,opt,name=project_id,json=projectId,proto3" json:"project_id,omitempty"`
	// start_date is the beginning of the period during which a credit batch
	// was quantified and verified. If it is empty then there is no start date
	// limit.
	MinStartDate *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=min_start_date,json=minStartDate,proto3" json:"min_start_date,omitempty"`
	// max_end_date is the end of the period during which a credit batch was
	// quantified and verified. If it is empty then there is no end date
	// limit.
	MaxEndDate *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=max_end_date,json=maxEndDate,proto3" json:"max_end_date,omitempty"`
}

func (x *ProjectSelector) Reset() {
	*x = ProjectSelector{}
	if protoimpl.UnsafeEnabled {
		mi := &file_regen_ecocredit_marketplace_v1beta1_types_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProjectSelector) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProjectSelector) ProtoMessage() {}

// Deprecated: Use ProjectSelector.ProtoReflect.Descriptor instead.
func (*ProjectSelector) Descriptor() ([]byte, []int) {
	return file_regen_ecocredit_marketplace_v1beta1_types_proto_rawDescGZIP(), []int{2}
}

func (x *ProjectSelector) GetProjectId() uint64 {
	if x != nil {
		return x.ProjectId
	}
	return 0
}

func (x *ProjectSelector) GetMinStartDate() *timestamppb.Timestamp {
	if x != nil {
		return x.MinStartDate
	}
	return nil
}

func (x *ProjectSelector) GetMaxEndDate() *timestamppb.Timestamp {
	if x != nil {
		return x.MaxEndDate
	}
	return nil
}

type BatchSelector struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BatchId uint64 `protobuf:"varint,1,opt,name=batch_id,json=batchId,proto3" json:"batch_id,omitempty"`
}

func (x *BatchSelector) Reset() {
	*x = BatchSelector{}
	if protoimpl.UnsafeEnabled {
		mi := &file_regen_ecocredit_marketplace_v1beta1_types_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BatchSelector) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BatchSelector) ProtoMessage() {}

// Deprecated: Use BatchSelector.ProtoReflect.Descriptor instead.
func (*BatchSelector) Descriptor() ([]byte, []int) {
	return file_regen_ecocredit_marketplace_v1beta1_types_proto_rawDescGZIP(), []int{3}
}

func (x *BatchSelector) GetBatchId() uint64 {
	if x != nil {
		return x.BatchId
	}
	return 0
}

// Criteria is a simple filter criteria for matching a credit batch.
type Filter_Criteria struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Valid selector types are all
	// attributes which are assigned to credit batches by some authority such
	// as the credit issuer or a curator. Requiring some authority-based selector
	// ensures that buy orders cannot just match some randomly issued credit
	// based on location and dates.
	//
	// Types that are assignable to Selector:
	//	*Filter_Criteria_ClassSelector
	//	*Filter_Criteria_ProjectSelector
	//	*Filter_Criteria_BatchSelector
	Selector isFilter_Criteria_Selector `protobuf_oneof:"selector"`
}

func (x *Filter_Criteria) Reset() {
	*x = Filter_Criteria{}
	if protoimpl.UnsafeEnabled {
		mi := &file_regen_ecocredit_marketplace_v1beta1_types_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Filter_Criteria) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Filter_Criteria) ProtoMessage() {}

// Deprecated: Use Filter_Criteria.ProtoReflect.Descriptor instead.
func (*Filter_Criteria) Descriptor() ([]byte, []int) {
	return file_regen_ecocredit_marketplace_v1beta1_types_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Filter_Criteria) GetSelector() isFilter_Criteria_Selector {
	if x != nil {
		return x.Selector
	}
	return nil
}

func (x *Filter_Criteria) GetClassSelector() *ClassSelector {
	if x, ok := x.GetSelector().(*Filter_Criteria_ClassSelector); ok {
		return x.ClassSelector
	}
	return nil
}

func (x *Filter_Criteria) GetProjectSelector() *ProjectSelector {
	if x, ok := x.GetSelector().(*Filter_Criteria_ProjectSelector); ok {
		return x.ProjectSelector
	}
	return nil
}

func (x *Filter_Criteria) GetBatchSelector() *BatchSelector {
	if x, ok := x.GetSelector().(*Filter_Criteria_BatchSelector); ok {
		return x.BatchSelector
	}
	return nil
}

type isFilter_Criteria_Selector interface {
	isFilter_Criteria_Selector()
}

type Filter_Criteria_ClassSelector struct {
	ClassSelector *ClassSelector `protobuf:"bytes,1,opt,name=class_selector,json=classSelector,proto3,oneof"`
}

type Filter_Criteria_ProjectSelector struct {
	ProjectSelector *ProjectSelector `protobuf:"bytes,2,opt,name=project_selector,json=projectSelector,proto3,oneof"`
}

type Filter_Criteria_BatchSelector struct {
	BatchSelector *BatchSelector `protobuf:"bytes,3,opt,name=batch_selector,json=batchSelector,proto3,oneof"`
}

func (*Filter_Criteria_ClassSelector) isFilter_Criteria_Selector() {}

func (*Filter_Criteria_ProjectSelector) isFilter_Criteria_Selector() {}

func (*Filter_Criteria_BatchSelector) isFilter_Criteria_Selector() {}

var File_regen_ecocredit_marketplace_v1beta1_types_proto protoreflect.FileDescriptor

var file_regen_ecocredit_marketplace_v1beta1_types_proto_rawDesc = []byte{
	0x0a, 0x2f, 0x72, 0x65, 0x67, 0x65, 0x6e, 0x2f, 0x65, 0x63, 0x6f, 0x63, 0x72, 0x65, 0x64, 0x69,
	0x74, 0x2f, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x2f, 0x76, 0x31,
	0x62, 0x65, 0x74, 0x61, 0x31, 0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x23, 0x72, 0x65, 0x67, 0x65, 0x6e, 0x2e, 0x65, 0x63, 0x6f, 0x63, 0x72, 0x65, 0x64,
	0x69, 0x74, 0x2e, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x2e, 0x76,
	0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x84, 0x03, 0x0a, 0x06, 0x46, 0x69, 0x6c, 0x74,
	0x65, 0x72, 0x12, 0x44, 0x0a, 0x02, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x34,
	0x2e, 0x72, 0x65, 0x67, 0x65, 0x6e, 0x2e, 0x65, 0x63, 0x6f, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74,
	0x2e, 0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x62,
	0x65, 0x74, 0x61, 0x31, 0x2e, 0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x2e, 0x43, 0x72, 0x69, 0x74,
	0x65, 0x72, 0x69, 0x61, 0x52, 0x02, 0x6f, 0x72, 0x1a, 0xb3, 0x02, 0x0a, 0x08, 0x43, 0x72, 0x69,
	0x74, 0x65, 0x72, 0x69, 0x61, 0x12, 0x5b, 0x0a, 0x0e, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x5f, 0x73,
	0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x32, 0x2e,
	0x72, 0x65, 0x67, 0x65, 0x6e, 0x2e, 0x65, 0x63, 0x6f, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x2e,
	0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x62, 0x65,
	0x74, 0x61, 0x31, 0x2e, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f,
	0x72, 0x48, 0x00, 0x52, 0x0d, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74,
	0x6f, 0x72, 0x12, 0x61, 0x0a, 0x10, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x73, 0x65,
	0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x34, 0x2e, 0x72,
	0x65, 0x67, 0x65, 0x6e, 0x2e, 0x65, 0x63, 0x6f, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x2e, 0x6d,
	0x61, 0x72, 0x6b, 0x65, 0x74, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74,
	0x61, 0x31, 0x2e, 0x50, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74,
	0x6f, 0x72, 0x48, 0x00, 0x52, 0x0f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x53, 0x65, 0x6c,
	0x65, 0x63, 0x74, 0x6f, 0x72, 0x12, 0x5b, 0x0a, 0x0e, 0x62, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x73,
	0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x32, 0x2e,
	0x72, 0x65, 0x67, 0x65, 0x6e, 0x2e, 0x65, 0x63, 0x6f, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x2e,
	0x6d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x62, 0x65,
	0x74, 0x61, 0x31, 0x2e, 0x42, 0x61, 0x74, 0x63, 0x68, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f,
	0x72, 0x48, 0x00, 0x52, 0x0d, 0x62, 0x61, 0x74, 0x63, 0x68, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74,
	0x6f, 0x72, 0x42, 0x0a, 0x0a, 0x08, 0x73, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x22, 0xd5,
	0x01, 0x0a, 0x0d, 0x43, 0x6c, 0x61, 0x73, 0x73, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72,
	0x12, 0x19, 0x0a, 0x08, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x07, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x49, 0x64, 0x12, 0x29, 0x0a, 0x10, 0x70,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x4c, 0x6f,
	0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x40, 0x0a, 0x0e, 0x6d, 0x69, 0x6e, 0x5f, 0x73, 0x74,
	0x61, 0x72, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0c, 0x6d, 0x69, 0x6e, 0x53,
	0x74, 0x61, 0x72, 0x74, 0x44, 0x61, 0x74, 0x65, 0x12, 0x3c, 0x0a, 0x0c, 0x6d, 0x61, 0x78, 0x5f,
	0x65, 0x6e, 0x64, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x6d, 0x61, 0x78, 0x45,
	0x6e, 0x64, 0x44, 0x61, 0x74, 0x65, 0x22, 0xb0, 0x01, 0x0a, 0x0f, 0x50, 0x72, 0x6f, 0x6a, 0x65,
	0x63, 0x74, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72,
	0x6f, 0x6a, 0x65, 0x63, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x09,
	0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x49, 0x64, 0x12, 0x40, 0x0a, 0x0e, 0x6d, 0x69, 0x6e,
	0x5f, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0c, 0x6d,
	0x69, 0x6e, 0x53, 0x74, 0x61, 0x72, 0x74, 0x44, 0x61, 0x74, 0x65, 0x12, 0x3c, 0x0a, 0x0c, 0x6d,
	0x61, 0x78, 0x5f, 0x65, 0x6e, 0x64, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x6d,
	0x61, 0x78, 0x45, 0x6e, 0x64, 0x44, 0x61, 0x74, 0x65, 0x22, 0x2a, 0x0a, 0x0d, 0x42, 0x61, 0x74,
	0x63, 0x68, 0x53, 0x65, 0x6c, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x12, 0x19, 0x0a, 0x08, 0x62, 0x61,
	0x74, 0x63, 0x68, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x07, 0x62, 0x61,
	0x74, 0x63, 0x68, 0x49, 0x64, 0x42, 0xc6, 0x02, 0x0a, 0x27, 0x63, 0x6f, 0x6d, 0x2e, 0x72, 0x65,
	0x67, 0x65, 0x6e, 0x2e, 0x65, 0x63, 0x6f, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x2e, 0x6d, 0x61,
	0x72, 0x6b, 0x65, 0x74, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61,
	0x31, 0x42, 0x0a, 0x54, 0x79, 0x70, 0x65, 0x73, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a,
	0x60, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x65, 0x67, 0x65,
	0x6e, 0x2d, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2f, 0x72, 0x65, 0x67, 0x65, 0x6e, 0x2d, 0x6e,
	0x65, 0x74, 0x77, 0x6f, 0x72, 0x6b, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x72, 0x65, 0x67, 0x65, 0x6e,
	0x2f, 0x65, 0x63, 0x6f, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x2f, 0x6d, 0x61, 0x72, 0x6b, 0x65,
	0x74, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x2f, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x3b, 0x6d,
	0x61, 0x72, 0x6b, 0x65, 0x74, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61,
	0x31, 0xa2, 0x02, 0x03, 0x52, 0x45, 0x4d, 0xaa, 0x02, 0x23, 0x52, 0x65, 0x67, 0x65, 0x6e, 0x2e,
	0x45, 0x63, 0x6f, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x2e, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74,
	0x70, 0x6c, 0x61, 0x63, 0x65, 0x2e, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0xca, 0x02, 0x23,
	0x52, 0x65, 0x67, 0x65, 0x6e, 0x5c, 0x45, 0x63, 0x6f, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x5c,
	0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x5c, 0x56, 0x31, 0x62, 0x65,
	0x74, 0x61, 0x31, 0xe2, 0x02, 0x2f, 0x52, 0x65, 0x67, 0x65, 0x6e, 0x5c, 0x45, 0x63, 0x6f, 0x63,
	0x72, 0x65, 0x64, 0x69, 0x74, 0x5c, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74, 0x70, 0x6c, 0x61, 0x63,
	0x65, 0x5c, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x26, 0x52, 0x65, 0x67, 0x65, 0x6e, 0x3a, 0x3a, 0x45,
	0x63, 0x6f, 0x63, 0x72, 0x65, 0x64, 0x69, 0x74, 0x3a, 0x3a, 0x4d, 0x61, 0x72, 0x6b, 0x65, 0x74,
	0x70, 0x6c, 0x61, 0x63, 0x65, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_regen_ecocredit_marketplace_v1beta1_types_proto_rawDescOnce sync.Once
	file_regen_ecocredit_marketplace_v1beta1_types_proto_rawDescData = file_regen_ecocredit_marketplace_v1beta1_types_proto_rawDesc
)

func file_regen_ecocredit_marketplace_v1beta1_types_proto_rawDescGZIP() []byte {
	file_regen_ecocredit_marketplace_v1beta1_types_proto_rawDescOnce.Do(func() {
		file_regen_ecocredit_marketplace_v1beta1_types_proto_rawDescData = protoimpl.X.CompressGZIP(file_regen_ecocredit_marketplace_v1beta1_types_proto_rawDescData)
	})
	return file_regen_ecocredit_marketplace_v1beta1_types_proto_rawDescData
}

var file_regen_ecocredit_marketplace_v1beta1_types_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_regen_ecocredit_marketplace_v1beta1_types_proto_goTypes = []interface{}{
	(*Filter)(nil),                // 0: regen.ecocredit.marketplace.v1beta1.Filter
	(*ClassSelector)(nil),         // 1: regen.ecocredit.marketplace.v1beta1.ClassSelector
	(*ProjectSelector)(nil),       // 2: regen.ecocredit.marketplace.v1beta1.ProjectSelector
	(*BatchSelector)(nil),         // 3: regen.ecocredit.marketplace.v1beta1.BatchSelector
	(*Filter_Criteria)(nil),       // 4: regen.ecocredit.marketplace.v1beta1.Filter.Criteria
	(*timestamppb.Timestamp)(nil), // 5: google.protobuf.Timestamp
}
var file_regen_ecocredit_marketplace_v1beta1_types_proto_depIdxs = []int32{
	4, // 0: regen.ecocredit.marketplace.v1beta1.Filter.or:type_name -> regen.ecocredit.marketplace.v1beta1.Filter.Criteria
	5, // 1: regen.ecocredit.marketplace.v1beta1.ClassSelector.min_start_date:type_name -> google.protobuf.Timestamp
	5, // 2: regen.ecocredit.marketplace.v1beta1.ClassSelector.max_end_date:type_name -> google.protobuf.Timestamp
	5, // 3: regen.ecocredit.marketplace.v1beta1.ProjectSelector.min_start_date:type_name -> google.protobuf.Timestamp
	5, // 4: regen.ecocredit.marketplace.v1beta1.ProjectSelector.max_end_date:type_name -> google.protobuf.Timestamp
	1, // 5: regen.ecocredit.marketplace.v1beta1.Filter.Criteria.class_selector:type_name -> regen.ecocredit.marketplace.v1beta1.ClassSelector
	2, // 6: regen.ecocredit.marketplace.v1beta1.Filter.Criteria.project_selector:type_name -> regen.ecocredit.marketplace.v1beta1.ProjectSelector
	3, // 7: regen.ecocredit.marketplace.v1beta1.Filter.Criteria.batch_selector:type_name -> regen.ecocredit.marketplace.v1beta1.BatchSelector
	8, // [8:8] is the sub-list for method output_type
	8, // [8:8] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_regen_ecocredit_marketplace_v1beta1_types_proto_init() }
func file_regen_ecocredit_marketplace_v1beta1_types_proto_init() {
	if File_regen_ecocredit_marketplace_v1beta1_types_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_regen_ecocredit_marketplace_v1beta1_types_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Filter); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_regen_ecocredit_marketplace_v1beta1_types_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ClassSelector); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_regen_ecocredit_marketplace_v1beta1_types_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProjectSelector); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_regen_ecocredit_marketplace_v1beta1_types_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BatchSelector); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_regen_ecocredit_marketplace_v1beta1_types_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Filter_Criteria); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_regen_ecocredit_marketplace_v1beta1_types_proto_msgTypes[4].OneofWrappers = []interface{}{
		(*Filter_Criteria_ClassSelector)(nil),
		(*Filter_Criteria_ProjectSelector)(nil),
		(*Filter_Criteria_BatchSelector)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_regen_ecocredit_marketplace_v1beta1_types_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_regen_ecocredit_marketplace_v1beta1_types_proto_goTypes,
		DependencyIndexes: file_regen_ecocredit_marketplace_v1beta1_types_proto_depIdxs,
		MessageInfos:      file_regen_ecocredit_marketplace_v1beta1_types_proto_msgTypes,
	}.Build()
	File_regen_ecocredit_marketplace_v1beta1_types_proto = out.File
	file_regen_ecocredit_marketplace_v1beta1_types_proto_rawDesc = nil
	file_regen_ecocredit_marketplace_v1beta1_types_proto_goTypes = nil
	file_regen_ecocredit_marketplace_v1beta1_types_proto_depIdxs = nil
}
