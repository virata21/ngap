package ngap

import "github.com/lvdund/ngap/ies"

func createMessage(present uint8, procedureCode ies.ProcedureCode) MessageUnmarshaller {
	switch present {
	case ies.NgapPduInitiatingMessage:
		switch int64(procedureCode.Value) {
		case ies.ProcedureCode_AMFConfigurationUpdate:
			return new(ies.AMFConfigurationUpdate)
		case ies.ProcedureCode_AMFStatusIndication:
			return new(ies.AMFStatusIndication)
		case ies.ProcedureCode_CellTrafficTrace:
			return new(ies.CellTrafficTrace)
		case ies.ProcedureCode_DeactivateTrace:
			return new(ies.DeactivateTrace)
		case ies.ProcedureCode_DownlinkNASTransport:
			return new(ies.DownlinkNASTransport)
		case ies.ProcedureCode_DownlinkNonUEAssociatedNRPPaTransport:
			return new(ies.DownlinkNonUEAssociatedNRPPaTransport)
		case ies.ProcedureCode_DownlinkRANConfigurationTransfer:
			return new(ies.DownlinkRANConfigurationTransfer)
		case ies.ProcedureCode_DownlinkRANStatusTransfer:
			return new(ies.DownlinkRANStatusTransfer)
		case ies.ProcedureCode_DownlinkUEAssociatedNRPPaTransport:
			return new(ies.DownlinkUEAssociatedNRPPaTransport)
		case ies.ProcedureCode_ErrorIndication:
			return new(ies.ErrorIndication)
		case ies.ProcedureCode_HandoverCancel:
			return new(ies.HandoverCancel)
		case ies.ProcedureCode_HandoverNotification:
			return new(ies.HandoverNotify)
		case ies.ProcedureCode_HandoverPreparation:
			return new(ies.HandoverRequired)
		case ies.ProcedureCode_HandoverResourceAllocation:
			return new(ies.HandoverRequest)
		case ies.ProcedureCode_InitialContextSetup:
			return new(ies.InitialContextSetupRequest)
		case ies.ProcedureCode_InitialUEMessage:
			return new(ies.InitialUEMessage)
		case ies.ProcedureCode_LocationReport:
			return new(ies.LocationReport)
		case ies.ProcedureCode_LocationReportingControl:
			return new(ies.LocationReportingControl)
		case ies.ProcedureCode_LocationReportingFailureIndication:
			return new(ies.LocationReportingFailureIndication)
		case ies.ProcedureCode_NASNonDeliveryIndication:
			return new(ies.NASNonDeliveryIndication)
		case ies.ProcedureCode_NGReset:
			return new(ies.NGReset)
		case ies.ProcedureCode_NGSetup:
			return new(ies.NGSetupRequest)
		case ies.ProcedureCode_OverloadStart:
			return new(ies.OverloadStart)
		// case ies.ProcedureCode_OverloadStop:
		// 	return new(ies.OverloadStop)
		case ies.ProcedureCode_Paging:
			return new(ies.Paging)
		case ies.ProcedureCode_PathSwitchRequest:
			return new(ies.PathSwitchRequest)
		case ies.ProcedureCode_PDUSessionResourceModify:
			return new(ies.PDUSessionResourceModifyRequest)
		case ies.ProcedureCode_PDUSessionResourceModifyIndication:
			return new(ies.PDUSessionResourceModifyIndication)
		case ies.ProcedureCode_PDUSessionResourceNotify:
			return new(ies.PDUSessionResourceNotify)
		case ies.ProcedureCode_PDUSessionResourceRelease:
			return new(ies.PDUSessionResourceReleaseCommand)
		case ies.ProcedureCode_PDUSessionResourceSetup:
			return new(ies.PDUSessionResourceSetupRequest)
		// case ies.ProcedureCode_PrivateMessage:
		// 	return new(ies.PrivateMessage)
		case ies.ProcedureCode_PWSCancel:
			return new(ies.PWSCancelRequest)
		case ies.ProcedureCode_PWSFailureIndication:
			return new(ies.PWSFailureIndication)
		case ies.ProcedureCode_PWSRestartIndication:
			return new(ies.PWSRestartIndication)
		case ies.ProcedureCode_RANConfigurationUpdate:
			return new(ies.RANConfigurationUpdate)
		case ies.ProcedureCode_RerouteNASRequest:
			return new(ies.RerouteNASRequest)
		case ies.ProcedureCode_RRCInactiveTransitionReport:
			return new(ies.RRCInactiveTransitionReport)
		case ies.ProcedureCode_SecondaryRATDataUsageReport:
			return new(ies.SecondaryRATDataUsageReport)
		case ies.ProcedureCode_TraceFailureIndication:
			return new(ies.TraceFailureIndication)
		case ies.ProcedureCode_TraceStart:
			return new(ies.TraceStart)
		case ies.ProcedureCode_UEContextModification:
			return new(ies.UEContextModificationRequest)
		case ies.ProcedureCode_UEContextRelease:
			return new(ies.UEContextReleaseCommand)
		case ies.ProcedureCode_UEContextReleaseRequest:
			return new(ies.UEContextReleaseRequest)
		case ies.ProcedureCode_UERadioCapabilityCheck:
			return new(ies.UERadioCapabilityCheckRequest)
		case ies.ProcedureCode_UERadioCapabilityInfoIndication:
			return new(ies.UERadioCapabilityInfoIndication)
		case ies.ProcedureCode_UETNLABindingRelease:
			return new(ies.UETNLABindingReleaseRequest)
		case ies.ProcedureCode_UplinkNASTransport:
			return new(ies.UplinkNASTransport)
		case ies.ProcedureCode_UplinkNonUEAssociatedNRPPaTransport:
			return new(ies.UplinkNonUEAssociatedNRPPaTransport)
		case ies.ProcedureCode_UplinkRANConfigurationTransfer:
			return new(ies.UplinkRANConfigurationTransfer)
		case ies.ProcedureCode_UplinkRANStatusTransfer:
			return new(ies.UplinkRANStatusTransfer)
		case ies.ProcedureCode_UplinkUEAssociatedNRPPaTransport:
			return new(ies.UplinkUEAssociatedNRPPaTransport)
		case ies.ProcedureCode_WriteReplaceWarning:
			return new(ies.WriteReplaceWarningRequest)
		}
	case ies.NgapPduSuccessfulOutcome:
		switch int64(procedureCode.Value) {
		case ies.ProcedureCode_AMFConfigurationUpdate:
			return new(ies.AMFConfigurationUpdateAcknowledge)
		case ies.ProcedureCode_HandoverCancel:
			return new(ies.HandoverCancelAcknowledge)
		case ies.ProcedureCode_HandoverPreparation:
			return new(ies.HandoverCommand)
		case ies.ProcedureCode_HandoverResourceAllocation:
			return new(ies.HandoverRequestAcknowledge)
		case ies.ProcedureCode_InitialContextSetup:
			return new(ies.InitialContextSetupResponse)
		case ies.ProcedureCode_NGReset:
			return new(ies.NGResetAcknowledge)
		case ies.ProcedureCode_NGSetup:
			return new(ies.NGSetupResponse)
		case ies.ProcedureCode_PathSwitchRequest:
			return new(ies.PathSwitchRequestAcknowledge)
		case ies.ProcedureCode_PDUSessionResourceModify:
			return new(ies.PDUSessionResourceModifyResponse)
		case ies.ProcedureCode_PDUSessionResourceModifyIndication:
			return new(ies.PDUSessionResourceModifyConfirm)
		case ies.ProcedureCode_PDUSessionResourceRelease:
			return new(ies.PDUSessionResourceReleaseResponse)
		case ies.ProcedureCode_PDUSessionResourceSetup:
			return new(ies.PDUSessionResourceSetupResponse)
		case ies.ProcedureCode_PWSCancel:
			return new(ies.PWSCancelResponse)
		case ies.ProcedureCode_RANConfigurationUpdate:
			return new(ies.RANConfigurationUpdateAcknowledge)
		case ies.ProcedureCode_UEContextModification:
			return new(ies.UEContextModificationResponse)
		case ies.ProcedureCode_UEContextRelease:
			return new(ies.UEContextReleaseComplete)
		case ies.ProcedureCode_UERadioCapabilityCheck:
			return new(ies.UERadioCapabilityCheckResponse)
		case ies.ProcedureCode_WriteReplaceWarning:
			return new(ies.WriteReplaceWarningResponse)
		}
	case ies.NgapPduUnsuccessfulOutcome:
		switch int64(procedureCode.Value) {
		case ies.ProcedureCode_AMFConfigurationUpdate:
			return new(ies.AMFConfigurationUpdateFailure)
		case ies.ProcedureCode_HandoverPreparation:
			return new(ies.HandoverPreparationFailure)
		case ies.ProcedureCode_HandoverResourceAllocation:
			return new(ies.HandoverFailure)
		case ies.ProcedureCode_InitialContextSetup:
			return new(ies.InitialContextSetupFailure)
		case ies.ProcedureCode_NGSetup:
			return new(ies.NGSetupFailure)
		case ies.ProcedureCode_PathSwitchRequest:
			return new(ies.PathSwitchRequestFailure)
		case ies.ProcedureCode_RANConfigurationUpdate:
			return new(ies.RANConfigurationUpdateFailure)
		case ies.ProcedureCode_UEContextModification:
			return new(ies.UEContextModificationFailure)
		}
	default:
		return nil
	}
	return nil
}
