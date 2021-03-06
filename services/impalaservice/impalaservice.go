// Autogenerated by Thrift Compiler (0.9.2)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package impalaservice

import (
	"bytes"
	"fmt"
	"git.apache.org/thrift.git/lib/go/thrift"
	"github.com/koblas/impalathing/services/beeswax"
	"github.com/koblas/impalathing/services/cli_service"
	"github.com/koblas/impalathing/services/status"
)

// (needed to ensure safety because of naive import list construction.)
var _ = thrift.ZERO
var _ = fmt.Printf
var _ = bytes.Equal

var _ = status.GoUnusedProtection__
var _ = beeswax.GoUnusedProtection__
var _ = cli_service.GoUnusedProtection__

type ImpalaService interface {
	beeswax.BeeswaxService

	// Parameters:
	//  - QueryId
	Cancel(query_id *beeswax.QueryHandle) (r *status.TStatus, err error)
	// Parameters:
	//  - Handle
	CloseInsert(handle *beeswax.QueryHandle) (r *TInsertResult_, err error)
	PingImpalaService() (err error)
}

type ImpalaServiceClient struct {
	*beeswax.BeeswaxServiceClient
}

func NewImpalaServiceClientFactory(t thrift.TTransport, f thrift.TProtocolFactory) *ImpalaServiceClient {
	return &ImpalaServiceClient{BeeswaxServiceClient: beeswax.NewBeeswaxServiceClientFactory(t, f)}
}

func NewImpalaServiceClientProtocol(t thrift.TTransport, iprot thrift.TProtocol, oprot thrift.TProtocol) *ImpalaServiceClient {
	return &ImpalaServiceClient{BeeswaxServiceClient: beeswax.NewBeeswaxServiceClientProtocol(t, iprot, oprot)}
}

// Parameters:
//  - QueryId
func (p *ImpalaServiceClient) Cancel(query_id *beeswax.QueryHandle) (r *status.TStatus, err error) {
	if err = p.sendCancel(query_id); err != nil {
		return
	}
	return p.recvCancel()
}

func (p *ImpalaServiceClient) sendCancel(query_id *beeswax.QueryHandle) (err error) {
	oprot := p.OutputProtocol
	if oprot == nil {
		oprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.OutputProtocol = oprot
	}
	p.SeqId++
	if err = oprot.WriteMessageBegin("Cancel", thrift.CALL, p.SeqId); err != nil {
		return
	}
	args := CancelArgs{
		QueryId: query_id,
	}
	if err = args.Write(oprot); err != nil {
		return
	}
	if err = oprot.WriteMessageEnd(); err != nil {
		return
	}
	return oprot.Flush()
}

func (p *ImpalaServiceClient) recvCancel() (value *status.TStatus, err error) {
	iprot := p.InputProtocol
	if iprot == nil {
		iprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.InputProtocol = iprot
	}
	_, mTypeId, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return
	}
	if mTypeId == thrift.EXCEPTION {
		error2 := thrift.NewTApplicationException(thrift.UNKNOWN_APPLICATION_EXCEPTION, "Unknown Exception")
		var error3 error
		error3, err = error2.Read(iprot)
		if err != nil {
			return
		}
		if err = iprot.ReadMessageEnd(); err != nil {
			return
		}
		err = error3
		return
	}
	if p.SeqId != seqId {
		err = thrift.NewTApplicationException(thrift.BAD_SEQUENCE_ID, "Cancel failed: out of sequence response")
		return
	}
	result := CancelResult{}
	if err = result.Read(iprot); err != nil {
		return
	}
	if err = iprot.ReadMessageEnd(); err != nil {
		return
	}
	if result.ErrorA1 != nil {
		err = result.ErrorA1
		return
	}
	value = result.GetSuccess()
	return
}

// Parameters:
//  - Handle
func (p *ImpalaServiceClient) CloseInsert(handle *beeswax.QueryHandle) (r *TInsertResult_, err error) {
	if err = p.sendCloseInsert(handle); err != nil {
		return
	}
	return p.recvCloseInsert()
}

func (p *ImpalaServiceClient) sendCloseInsert(handle *beeswax.QueryHandle) (err error) {
	oprot := p.OutputProtocol
	if oprot == nil {
		oprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.OutputProtocol = oprot
	}
	p.SeqId++
	if err = oprot.WriteMessageBegin("CloseInsert", thrift.CALL, p.SeqId); err != nil {
		return
	}
	args := CloseInsertArgs{
		Handle: handle,
	}
	if err = args.Write(oprot); err != nil {
		return
	}
	if err = oprot.WriteMessageEnd(); err != nil {
		return
	}
	return oprot.Flush()
}

func (p *ImpalaServiceClient) recvCloseInsert() (value *TInsertResult_, err error) {
	iprot := p.InputProtocol
	if iprot == nil {
		iprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.InputProtocol = iprot
	}
	_, mTypeId, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return
	}
	if mTypeId == thrift.EXCEPTION {
		error4 := thrift.NewTApplicationException(thrift.UNKNOWN_APPLICATION_EXCEPTION, "Unknown Exception")
		var error5 error
		error5, err = error4.Read(iprot)
		if err != nil {
			return
		}
		if err = iprot.ReadMessageEnd(); err != nil {
			return
		}
		err = error5
		return
	}
	if p.SeqId != seqId {
		err = thrift.NewTApplicationException(thrift.BAD_SEQUENCE_ID, "CloseInsert failed: out of sequence response")
		return
	}
	result := CloseInsertResult{}
	if err = result.Read(iprot); err != nil {
		return
	}
	if err = iprot.ReadMessageEnd(); err != nil {
		return
	}
	if result.ErrorA1 != nil {
		err = result.ErrorA1
		return
	} else if result.Error2 != nil {
		err = result.Error2
		return
	}
	value = result.GetSuccess()
	return
}

func (p *ImpalaServiceClient) PingImpalaService() (err error) {
	if err = p.sendPingImpalaService(); err != nil {
		return
	}
	return p.recvPingImpalaService()
}

func (p *ImpalaServiceClient) sendPingImpalaService() (err error) {
	oprot := p.OutputProtocol
	if oprot == nil {
		oprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.OutputProtocol = oprot
	}
	p.SeqId++
	if err = oprot.WriteMessageBegin("PingImpalaService", thrift.CALL, p.SeqId); err != nil {
		return
	}
	args := PingImpalaServiceArgs{}
	if err = args.Write(oprot); err != nil {
		return
	}
	if err = oprot.WriteMessageEnd(); err != nil {
		return
	}
	return oprot.Flush()
}

func (p *ImpalaServiceClient) recvPingImpalaService() (err error) {
	iprot := p.InputProtocol
	if iprot == nil {
		iprot = p.ProtocolFactory.GetProtocol(p.Transport)
		p.InputProtocol = iprot
	}
	_, mTypeId, seqId, err := iprot.ReadMessageBegin()
	if err != nil {
		return
	}
	if mTypeId == thrift.EXCEPTION {
		error6 := thrift.NewTApplicationException(thrift.UNKNOWN_APPLICATION_EXCEPTION, "Unknown Exception")
		var error7 error
		error7, err = error6.Read(iprot)
		if err != nil {
			return
		}
		if err = iprot.ReadMessageEnd(); err != nil {
			return
		}
		err = error7
		return
	}
	if p.SeqId != seqId {
		err = thrift.NewTApplicationException(thrift.BAD_SEQUENCE_ID, "PingImpalaService failed: out of sequence response")
		return
	}
	result := PingImpalaServiceResult{}
	if err = result.Read(iprot); err != nil {
		return
	}
	if err = iprot.ReadMessageEnd(); err != nil {
		return
	}
	return
}

type ImpalaServiceProcessor struct {
	*beeswax.BeeswaxServiceProcessor
}

func NewImpalaServiceProcessor(handler ImpalaService) *ImpalaServiceProcessor {
	self8 := &ImpalaServiceProcessor{beeswax.NewBeeswaxServiceProcessor(handler)}
	self8.AddToProcessorMap("Cancel", &impalaServiceProcessorCancel{handler: handler})
	self8.AddToProcessorMap("CloseInsert", &impalaServiceProcessorCloseInsert{handler: handler})
	self8.AddToProcessorMap("PingImpalaService", &impalaServiceProcessorPingImpalaService{handler: handler})
	return self8
}

type impalaServiceProcessorCancel struct {
	handler ImpalaService
}

func (p *impalaServiceProcessorCancel) Process(seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := CancelArgs{}
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("Cancel", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return false, err
	}

	iprot.ReadMessageEnd()
	result := CancelResult{}
	var retval *status.TStatus
	var err2 error
	if retval, err2 = p.handler.Cancel(args.QueryId); err2 != nil {
		switch v := err2.(type) {
		case *beeswax.BeeswaxException:
			result.ErrorA1 = v
		default:
			x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing Cancel: "+err2.Error())
			oprot.WriteMessageBegin("Cancel", thrift.EXCEPTION, seqId)
			x.Write(oprot)
			oprot.WriteMessageEnd()
			oprot.Flush()
			return true, err2
		}
	} else {
		result.Success = retval
	}
	if err2 = oprot.WriteMessageBegin("Cancel", thrift.REPLY, seqId); err2 != nil {
		err = err2
	}
	if err2 = result.Write(oprot); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.Flush(); err == nil && err2 != nil {
		err = err2
	}
	if err != nil {
		return
	}
	return true, err
}

type impalaServiceProcessorCloseInsert struct {
	handler ImpalaService
}

func (p *impalaServiceProcessorCloseInsert) Process(seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := CloseInsertArgs{}
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("CloseInsert", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return false, err
	}

	iprot.ReadMessageEnd()
	result := CloseInsertResult{}
	var retval *TInsertResult_
	var err2 error
	if retval, err2 = p.handler.CloseInsert(args.Handle); err2 != nil {
		switch v := err2.(type) {
		case *beeswax.QueryNotFoundException:
			result.ErrorA1 = v
		case *beeswax.BeeswaxException:
			result.Error2 = v
		default:
			x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing CloseInsert: "+err2.Error())
			oprot.WriteMessageBegin("CloseInsert", thrift.EXCEPTION, seqId)
			x.Write(oprot)
			oprot.WriteMessageEnd()
			oprot.Flush()
			return true, err2
		}
	} else {
		result.Success = retval
	}
	if err2 = oprot.WriteMessageBegin("CloseInsert", thrift.REPLY, seqId); err2 != nil {
		err = err2
	}
	if err2 = result.Write(oprot); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.Flush(); err == nil && err2 != nil {
		err = err2
	}
	if err != nil {
		return
	}
	return true, err
}

type impalaServiceProcessorPingImpalaService struct {
	handler ImpalaService
}

func (p *impalaServiceProcessorPingImpalaService) Process(seqId int32, iprot, oprot thrift.TProtocol) (success bool, err thrift.TException) {
	args := PingImpalaServiceArgs{}
	if err = args.Read(iprot); err != nil {
		iprot.ReadMessageEnd()
		x := thrift.NewTApplicationException(thrift.PROTOCOL_ERROR, err.Error())
		oprot.WriteMessageBegin("PingImpalaService", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return false, err
	}

	iprot.ReadMessageEnd()
	result := PingImpalaServiceResult{}
	var err2 error
	if err2 = p.handler.PingImpalaService(); err2 != nil {
		x := thrift.NewTApplicationException(thrift.INTERNAL_ERROR, "Internal error processing PingImpalaService: "+err2.Error())
		oprot.WriteMessageBegin("PingImpalaService", thrift.EXCEPTION, seqId)
		x.Write(oprot)
		oprot.WriteMessageEnd()
		oprot.Flush()
		return true, err2
	}
	if err2 = oprot.WriteMessageBegin("PingImpalaService", thrift.REPLY, seqId); err2 != nil {
		err = err2
	}
	if err2 = result.Write(oprot); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.WriteMessageEnd(); err == nil && err2 != nil {
		err = err2
	}
	if err2 = oprot.Flush(); err == nil && err2 != nil {
		err = err2
	}
	if err != nil {
		return
	}
	return true, err
}

// HELPER FUNCTIONS AND STRUCTURES

type CancelArgs struct {
	QueryId *beeswax.QueryHandle `thrift:"query_id,1" json:"query_id"`
}

func NewCancelArgs() *CancelArgs {
	return &CancelArgs{}
}

var CancelArgs_QueryId_DEFAULT *beeswax.QueryHandle

func (p *CancelArgs) GetQueryId() *beeswax.QueryHandle {
	if !p.IsSetQueryId() {
		return CancelArgs_QueryId_DEFAULT
	}
	return p.QueryId
}
func (p *CancelArgs) IsSetQueryId() bool {
	return p.QueryId != nil
}

func (p *CancelArgs) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return fmt.Errorf("%T read error: %s", p, err)
	}
	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return fmt.Errorf("%T field %d read error: %s", p, fieldId, err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if err := p.ReadField1(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return fmt.Errorf("%T read struct end error: %s", p, err)
	}
	return nil
}

func (p *CancelArgs) ReadField1(iprot thrift.TProtocol) error {
	p.QueryId = &beeswax.QueryHandle{}
	if err := p.QueryId.Read(iprot); err != nil {
		return fmt.Errorf("%T error reading struct: %s", p.QueryId, err)
	}
	return nil
}

func (p *CancelArgs) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("Cancel_args"); err != nil {
		return fmt.Errorf("%T write struct begin error: %s", p, err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return fmt.Errorf("write field stop error: %s", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return fmt.Errorf("write struct stop error: %s", err)
	}
	return nil
}

func (p *CancelArgs) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("query_id", thrift.STRUCT, 1); err != nil {
		return fmt.Errorf("%T write field begin error 1:query_id: %s", p, err)
	}
	if err := p.QueryId.Write(oprot); err != nil {
		return fmt.Errorf("%T error writing struct: %s", p.QueryId, err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return fmt.Errorf("%T write field end error 1:query_id: %s", p, err)
	}
	return err
}

func (p *CancelArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("CancelArgs(%+v)", *p)
}

type CancelResult struct {
	Success *status.TStatus           `thrift:"success,0" json:"success"`
	ErrorA1 *beeswax.BeeswaxException `thrift:"error,1" json:"error"`
}

func NewCancelResult() *CancelResult {
	return &CancelResult{}
}

var CancelResult_Success_DEFAULT *status.TStatus

func (p *CancelResult) GetSuccess() *status.TStatus {
	if !p.IsSetSuccess() {
		return CancelResult_Success_DEFAULT
	}
	return p.Success
}

var CancelResult_ErrorA1_DEFAULT *beeswax.BeeswaxException

func (p *CancelResult) GetErrorA1() *beeswax.BeeswaxException {
	if !p.IsSetErrorA1() {
		return CancelResult_ErrorA1_DEFAULT
	}
	return p.ErrorA1
}
func (p *CancelResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *CancelResult) IsSetErrorA1() bool {
	return p.ErrorA1 != nil
}

func (p *CancelResult) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return fmt.Errorf("%T read error: %s", p, err)
	}
	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return fmt.Errorf("%T field %d read error: %s", p, fieldId, err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 0:
			if err := p.ReadField0(iprot); err != nil {
				return err
			}
		case 1:
			if err := p.ReadField1(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return fmt.Errorf("%T read struct end error: %s", p, err)
	}
	return nil
}

func (p *CancelResult) ReadField0(iprot thrift.TProtocol) error {
	p.Success = &status.TStatus{}
	if err := p.Success.Read(iprot); err != nil {
		return fmt.Errorf("%T error reading struct: %s", p.Success, err)
	}
	return nil
}

func (p *CancelResult) ReadField1(iprot thrift.TProtocol) error {
	p.ErrorA1 = &beeswax.BeeswaxException{
		SQLState: "     ",
	}
	if err := p.ErrorA1.Read(iprot); err != nil {
		return fmt.Errorf("%T error reading struct: %s", p.ErrorA1, err)
	}
	return nil
}

func (p *CancelResult) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("Cancel_result"); err != nil {
		return fmt.Errorf("%T write struct begin error: %s", p, err)
	}
	if err := p.writeField0(oprot); err != nil {
		return err
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return fmt.Errorf("write field stop error: %s", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return fmt.Errorf("write struct stop error: %s", err)
	}
	return nil
}

func (p *CancelResult) writeField0(oprot thrift.TProtocol) (err error) {
	if p.IsSetSuccess() {
		if err := oprot.WriteFieldBegin("success", thrift.STRUCT, 0); err != nil {
			return fmt.Errorf("%T write field begin error 0:success: %s", p, err)
		}
		if err := p.Success.Write(oprot); err != nil {
			return fmt.Errorf("%T error writing struct: %s", p.Success, err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return fmt.Errorf("%T write field end error 0:success: %s", p, err)
		}
	}
	return err
}

func (p *CancelResult) writeField1(oprot thrift.TProtocol) (err error) {
	if p.IsSetErrorA1() {
		if err := oprot.WriteFieldBegin("error", thrift.STRUCT, 1); err != nil {
			return fmt.Errorf("%T write field begin error 1:error: %s", p, err)
		}
		if err := p.ErrorA1.Write(oprot); err != nil {
			return fmt.Errorf("%T error writing struct: %s", p.ErrorA1, err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return fmt.Errorf("%T write field end error 1:error: %s", p, err)
		}
	}
	return err
}

func (p *CancelResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("CancelResult(%+v)", *p)
}

type CloseInsertArgs struct {
	Handle *beeswax.QueryHandle `thrift:"handle,1" json:"handle"`
}

func NewCloseInsertArgs() *CloseInsertArgs {
	return &CloseInsertArgs{}
}

var CloseInsertArgs_Handle_DEFAULT *beeswax.QueryHandle

func (p *CloseInsertArgs) GetHandle() *beeswax.QueryHandle {
	if !p.IsSetHandle() {
		return CloseInsertArgs_Handle_DEFAULT
	}
	return p.Handle
}
func (p *CloseInsertArgs) IsSetHandle() bool {
	return p.Handle != nil
}

func (p *CloseInsertArgs) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return fmt.Errorf("%T read error: %s", p, err)
	}
	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return fmt.Errorf("%T field %d read error: %s", p, fieldId, err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 1:
			if err := p.ReadField1(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return fmt.Errorf("%T read struct end error: %s", p, err)
	}
	return nil
}

func (p *CloseInsertArgs) ReadField1(iprot thrift.TProtocol) error {
	p.Handle = &beeswax.QueryHandle{}
	if err := p.Handle.Read(iprot); err != nil {
		return fmt.Errorf("%T error reading struct: %s", p.Handle, err)
	}
	return nil
}

func (p *CloseInsertArgs) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("CloseInsert_args"); err != nil {
		return fmt.Errorf("%T write struct begin error: %s", p, err)
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return fmt.Errorf("write field stop error: %s", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return fmt.Errorf("write struct stop error: %s", err)
	}
	return nil
}

func (p *CloseInsertArgs) writeField1(oprot thrift.TProtocol) (err error) {
	if err := oprot.WriteFieldBegin("handle", thrift.STRUCT, 1); err != nil {
		return fmt.Errorf("%T write field begin error 1:handle: %s", p, err)
	}
	if err := p.Handle.Write(oprot); err != nil {
		return fmt.Errorf("%T error writing struct: %s", p.Handle, err)
	}
	if err := oprot.WriteFieldEnd(); err != nil {
		return fmt.Errorf("%T write field end error 1:handle: %s", p, err)
	}
	return err
}

func (p *CloseInsertArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("CloseInsertArgs(%+v)", *p)
}

type CloseInsertResult struct {
	Success *TInsertResult_                 `thrift:"success,0" json:"success"`
	ErrorA1 *beeswax.QueryNotFoundException `thrift:"error,1" json:"error"`
	Error2  *beeswax.BeeswaxException       `thrift:"error2,2" json:"error2"`
}

func NewCloseInsertResult() *CloseInsertResult {
	return &CloseInsertResult{}
}

var CloseInsertResult_Success_DEFAULT *TInsertResult_

func (p *CloseInsertResult) GetSuccess() *TInsertResult_ {
	if !p.IsSetSuccess() {
		return CloseInsertResult_Success_DEFAULT
	}
	return p.Success
}

var CloseInsertResult_ErrorA1_DEFAULT *beeswax.QueryNotFoundException

func (p *CloseInsertResult) GetErrorA1() *beeswax.QueryNotFoundException {
	if !p.IsSetErrorA1() {
		return CloseInsertResult_ErrorA1_DEFAULT
	}
	return p.ErrorA1
}

var CloseInsertResult_Error2_DEFAULT *beeswax.BeeswaxException

func (p *CloseInsertResult) GetError2() *beeswax.BeeswaxException {
	if !p.IsSetError2() {
		return CloseInsertResult_Error2_DEFAULT
	}
	return p.Error2
}
func (p *CloseInsertResult) IsSetSuccess() bool {
	return p.Success != nil
}

func (p *CloseInsertResult) IsSetErrorA1() bool {
	return p.ErrorA1 != nil
}

func (p *CloseInsertResult) IsSetError2() bool {
	return p.Error2 != nil
}

func (p *CloseInsertResult) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return fmt.Errorf("%T read error: %s", p, err)
	}
	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return fmt.Errorf("%T field %d read error: %s", p, fieldId, err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		switch fieldId {
		case 0:
			if err := p.ReadField0(iprot); err != nil {
				return err
			}
		case 1:
			if err := p.ReadField1(iprot); err != nil {
				return err
			}
		case 2:
			if err := p.ReadField2(iprot); err != nil {
				return err
			}
		default:
			if err := iprot.Skip(fieldTypeId); err != nil {
				return err
			}
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return fmt.Errorf("%T read struct end error: %s", p, err)
	}
	return nil
}

func (p *CloseInsertResult) ReadField0(iprot thrift.TProtocol) error {
	p.Success = &TInsertResult_{}
	if err := p.Success.Read(iprot); err != nil {
		return fmt.Errorf("%T error reading struct: %s", p.Success, err)
	}
	return nil
}

func (p *CloseInsertResult) ReadField1(iprot thrift.TProtocol) error {
	p.ErrorA1 = &beeswax.QueryNotFoundException{}
	if err := p.ErrorA1.Read(iprot); err != nil {
		return fmt.Errorf("%T error reading struct: %s", p.ErrorA1, err)
	}
	return nil
}

func (p *CloseInsertResult) ReadField2(iprot thrift.TProtocol) error {
	p.Error2 = &beeswax.BeeswaxException{
		SQLState: "     ",
	}
	if err := p.Error2.Read(iprot); err != nil {
		return fmt.Errorf("%T error reading struct: %s", p.Error2, err)
	}
	return nil
}

func (p *CloseInsertResult) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("CloseInsert_result"); err != nil {
		return fmt.Errorf("%T write struct begin error: %s", p, err)
	}
	if err := p.writeField0(oprot); err != nil {
		return err
	}
	if err := p.writeField1(oprot); err != nil {
		return err
	}
	if err := p.writeField2(oprot); err != nil {
		return err
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return fmt.Errorf("write field stop error: %s", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return fmt.Errorf("write struct stop error: %s", err)
	}
	return nil
}

func (p *CloseInsertResult) writeField0(oprot thrift.TProtocol) (err error) {
	if p.IsSetSuccess() {
		if err := oprot.WriteFieldBegin("success", thrift.STRUCT, 0); err != nil {
			return fmt.Errorf("%T write field begin error 0:success: %s", p, err)
		}
		if err := p.Success.Write(oprot); err != nil {
			return fmt.Errorf("%T error writing struct: %s", p.Success, err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return fmt.Errorf("%T write field end error 0:success: %s", p, err)
		}
	}
	return err
}

func (p *CloseInsertResult) writeField1(oprot thrift.TProtocol) (err error) {
	if p.IsSetErrorA1() {
		if err := oprot.WriteFieldBegin("error", thrift.STRUCT, 1); err != nil {
			return fmt.Errorf("%T write field begin error 1:error: %s", p, err)
		}
		if err := p.ErrorA1.Write(oprot); err != nil {
			return fmt.Errorf("%T error writing struct: %s", p.ErrorA1, err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return fmt.Errorf("%T write field end error 1:error: %s", p, err)
		}
	}
	return err
}

func (p *CloseInsertResult) writeField2(oprot thrift.TProtocol) (err error) {
	if p.IsSetError2() {
		if err := oprot.WriteFieldBegin("error2", thrift.STRUCT, 2); err != nil {
			return fmt.Errorf("%T write field begin error 2:error2: %s", p, err)
		}
		if err := p.Error2.Write(oprot); err != nil {
			return fmt.Errorf("%T error writing struct: %s", p.Error2, err)
		}
		if err := oprot.WriteFieldEnd(); err != nil {
			return fmt.Errorf("%T write field end error 2:error2: %s", p, err)
		}
	}
	return err
}

func (p *CloseInsertResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("CloseInsertResult(%+v)", *p)
}

type PingImpalaServiceArgs struct {
}

func NewPingImpalaServiceArgs() *PingImpalaServiceArgs {
	return &PingImpalaServiceArgs{}
}

func (p *PingImpalaServiceArgs) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return fmt.Errorf("%T read error: %s", p, err)
	}
	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return fmt.Errorf("%T field %d read error: %s", p, fieldId, err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		if err := iprot.Skip(fieldTypeId); err != nil {
			return err
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return fmt.Errorf("%T read struct end error: %s", p, err)
	}
	return nil
}

func (p *PingImpalaServiceArgs) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("PingImpalaService_args"); err != nil {
		return fmt.Errorf("%T write struct begin error: %s", p, err)
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return fmt.Errorf("write field stop error: %s", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return fmt.Errorf("write struct stop error: %s", err)
	}
	return nil
}

func (p *PingImpalaServiceArgs) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("PingImpalaServiceArgs(%+v)", *p)
}

type PingImpalaServiceResult struct {
}

func NewPingImpalaServiceResult() *PingImpalaServiceResult {
	return &PingImpalaServiceResult{}
}

func (p *PingImpalaServiceResult) Read(iprot thrift.TProtocol) error {
	if _, err := iprot.ReadStructBegin(); err != nil {
		return fmt.Errorf("%T read error: %s", p, err)
	}
	for {
		_, fieldTypeId, fieldId, err := iprot.ReadFieldBegin()
		if err != nil {
			return fmt.Errorf("%T field %d read error: %s", p, fieldId, err)
		}
		if fieldTypeId == thrift.STOP {
			break
		}
		if err := iprot.Skip(fieldTypeId); err != nil {
			return err
		}
		if err := iprot.ReadFieldEnd(); err != nil {
			return err
		}
	}
	if err := iprot.ReadStructEnd(); err != nil {
		return fmt.Errorf("%T read struct end error: %s", p, err)
	}
	return nil
}

func (p *PingImpalaServiceResult) Write(oprot thrift.TProtocol) error {
	if err := oprot.WriteStructBegin("PingImpalaService_result"); err != nil {
		return fmt.Errorf("%T write struct begin error: %s", p, err)
	}
	if err := oprot.WriteFieldStop(); err != nil {
		return fmt.Errorf("write field stop error: %s", err)
	}
	if err := oprot.WriteStructEnd(); err != nil {
		return fmt.Errorf("write struct stop error: %s", err)
	}
	return nil
}

func (p *PingImpalaServiceResult) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("PingImpalaServiceResult(%+v)", *p)
}
