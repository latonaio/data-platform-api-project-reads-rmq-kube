package dpfm_api_output_formatter

import (
	"data-platform-api-project-reads-rmq-kube/DPFM_API_Caller/requests"
	api_input_reader "data-platform-api-project-reads-rmq-kube/DPFM_API_Input_Reader"
	"database/sql"
	"fmt"
)

func ConvertToProject(sdc *api_input_reader.SDC, rows *sql.Rows) (*[]Project, error) {
	defer rows.Close()
	projects := make([]Project, 0, len(sdc.Projects))

	i := 0
	for rows.Next() {
		i++
		pm := &requests.Project{}

		err := rows.Scan(
			&pm.Project,
			&pm.ProjectCode,
			&pm.ProjectDescription,
			&pm.OwnerBusinessPartner,
			&pm.OwnerPlant,
			&pm.ProjectProfile,
			&pm.ResponsiblePerson,
			&pm.ResponsiblePersonName,
			&pm.PlannedStartDate,
			&pm.PlannedEndDate,
			&pm.ActualStartDate,
			&pm.ActualEndDate,
			&pm.CreationDate,
			&pm.LastChangeDate,
			&pm.IsMarkedForDeletion,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return &projects, err
		}

		data := pm

		projects = append(projects, Project{
			Project:               data.Project,
			ProjectCode:           data.ProjectCode,
			ProjectDescription:    data.ProjectDescription,
			OwnerBusinessPartner:  data.OwnerBusinessPartner,
			OwnerPlant:            data.OwnerPlant,
			ProjectProfile:        data.ProjectProfile,
			ResponsiblePerson:     data.ResponsiblePerson,
			ResponsiblePersonName: data.ResponsiblePersonName,
			PlannedStartDate:      data.PlannedStartDate,
			PlannedEndDate:        data.PlannedEndDate,
			ActualStartDate:       data.ActualStartDate,
			ActualEndDate:         data.ActualEndDate,
			CreationDate:          data.CreationDate,
			LastChangeDate:        data.LastChangeDate,
			IsMarkedForDeletion:   data.IsMarkedForDeletion,
		})
	}
	if i == 0 {
		fmt.Printf("DBに対象のレコードが存在しません。")
		return &projects, nil
	}

	return &projects, nil
}

func ConvertToWBSElement(sdc *api_input_reader.SDC, rows *sql.Rows) (*[]WBSElement, error) {
	defer rows.Close()
	pm := &requests.WBSElement{}

	i := 0
	for rows.Next() {
		i++

		err := rows.Scan(
			&pm.Project,
			&pm.WBSElement,
			&pm.WBSElementCode,
			&pm.WBSElementDescription,
			&pm.BusinessPartner,
			&pm.Plant,
			&pm.ResponsiblePerson,
			&pm.ResponsiblePersonName,
			&pm.PlannedStartDate,
			&pm.PlannedEndDate,
			&pm.ActualStartDate,
			&pm.ActualEndDate,
			&pm.CreationDate,
			&pm.LastChangeDate,
			&pm.IsMarkedForDeletion,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return &wBSElement{}, err
		}
	}
	if i == 0 {
		fmt.Printf("DBに対象のレコードが存在しません。")
		return &wBSElement{}, nil
	}

	data := pm
	wBSElement := &wBSElement{
		Project:               data.Project,
		WBSElement:            data.WBSElement,
		WBSElementCode:        data.WBSElementCode,
		WBSElementDescription: data.WBSElementDescription,
		BusinessPartner:       data.BusinessPartner,
		Plant:                 data.Plant,
		ResponsiblePerson:     data.ResponsiblePerson,
		ResponsiblePersonName: data.ResponsiblePersonName,
		PlannedStartDate:      data.PlannedStartDate,
		PlannedEndDate:        data.PlannedEndDate,
		ActualStartDate:       data.ActualStartDate,
		ActualEndDate:         data.ActualEndDate,
		CreationDate:          data.CreationDate,
		LastChangeDate:        data.LastChangeDate,
		IsMarkedForDeletion:   data.IsMarkedForDeletion,
	}
	return wBSElement, nil
}

func ConvertToNetwork(sdc *api_input_reader.SDC, rows *sql.Rows) (*[]Network, error) {
	defer rows.Close()
	pm := &requests.Network{}

	i := 0
	for rows.Next() {
		i++

		err := rows.Scan(
			&pm.Project,
			&pm.WBSElement,
			&pm.Network,
			&pm.NetworkDescription,
			&pm.BusinessPartner,
			&pm.Plant,
			&pm.ResponsiblePerson,
			&pm.ResponsiblePersonName,
			&pm.PlannedStartDate,
			&pm.PlannedEndDate,
			&pm.ActualStartDate,
			&pm.ActualEndDate,
			&pm.CreationDate,
			&pm.CreationTime,
			&pm.LastChangeDate,
			&pm.LastChangeTime,
			&pm.IsMarkedForDeletion,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return &network{}, err
		}
	}
	if i == 0 {
		fmt.Printf("DBに対象のレコードが存在しません。")
		return &network{}, nil
	}

	data := pm
	network := &network{
		Project:               data.Project,
		WBSElement:            data.WBSElement,
		Network:               data.Network,
		NetworkDescription:    data.NetworkDescription,
		BusinessPartner:       data.BusinessPartner,
		Plant:                 data.Plant,
		ResponsiblePerson:     data.ResponsiblePerson,
		ResponsiblePersonName: data.ResponsiblePersonName,
		PlannedStartDate:      data.PlannedStartDate,
		PlannedEndDate:        data.PlannedEndDate,
		ActualStartDate:       data.ActualStartDate,
		ActualEndDate:         data.ActualEndDate,
		CreationDate:          data.CreationDate,
		CreationTime:          data.CreationTime,
		LastChangeDate:        data.LastChangeDate,
		LastChangeTime:        data.LastChangeTime,
		IsMarkedForDeletion:   data.IsMarkedForDeletion,
	}
	return network, nil
}

func ConvertToNetworkComponent(sdc *api_input_reader.SDC, rows *sql.Rows) (*[]NetworkComponent, error) {
	defer rows.Close()
	pm := &requests.NetworkComponent{}

	i := 0
	for rows.Next() {
		i++

		err := rows.Scan(
			&pm.Project,
			&pm.WBSElement,
			&pm.Network,
			&pm.BillOfMaterial,
			&pm.BillOfMaterialItem,
			&pm.SupplyChainRelationshipID,
			&pm.SupplyChainRelationshipDeliveryID,
			&pm.SupplyChainRelationshipDeliveryPlantID,
			&pm.SupplyChainRelationshipStockConfPlantID,
			&pm.ProductionPlantBusinessPartner,
			&pm.ProductionPlant,
			&pm.MRPArea,
			&pm.MRPController,
			&pm.ProductionVersion,
			&pm.ProductionVersionItem,
			&pm.ComponentProduct,
			&pm.ComponentProductBuyer,
			&pm.ComponentProductSeller,
			&pm.ComponentProductDeliverToParty,
			&pm.ComponentProductDeliverToPlant,
			&pm.ComponentProductDeliverFromParty,
			&pm.ComponentProductDeliverFromPlant,
			&pm.ComponentProductBaseUnit,
			&pm.ComponentProductDeliveryUnit,
			&pm.ComponentProductRequirementDate,
			&pm.ComponentProductRequirementTime,
			&pm.ComponentProductRequiredQuantityInBaseUnit,
			&pm.ComponentProductRequiredQuantityInDeliveryUnit,
			&pm.ComponentProductPlannedScrapInPercent,
			&pm.ComponentProductIsMarkedForBackflush,
			&pm.BillOfMaterialItemText,
			&pm.StockConfirmationBusinessPartner,
			&pm.StockConfirmationPlant,
			&pm.StockConfirmationPlantStorageLocation,
			&pm.PlannedOrder,
			&pm.PlannedOrderItem,
			&pm.BOMItemDescription,
			&pm.ComponentProductBatch,
			&pm.ComponentProductBatchValidityStartDate,
			&pm.ComponentProductBatchValidityStartTime,
			&pm.ComponentProductBatchValidityEndDate,
			&pm.ComponentProductBatchValidityEndTime,
			&pm.ComponentProductCostingPolicy,
			&pm.ComponentProductPriceUnitQty,
			&pm.ComponentProductStandardPrice,
			&pm.ComponentProductMovingAveragePrice,
			&pm.ComponentProductWithdrawnQuantity,
			&pm.CreationDate,
			&pm.CreationTime,
			&pm.LastChangeDate,
			&pm.LastChangeTime,
			&pm.ComponentProductAvailabilityIsNotChecked,
			&pm.IsReleased,
			&pm.IsLocked,
			&pm.IsCancelled,
			&pm.IsMarkedForDeletion,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return &[]networkComponent{}, err
		}
	}
	if i == 0 {
		fmt.Printf("DBに対象のレコードが存在しません。")
		return &networkComponent{}, nil
	}

	data := pm
	networkComponent := &networkComponent{
		Project:                                        data.Project,
		WBSElement:                                     data.WBSElement,
		Network:                                        data.Network,
		BillOfMaterial:                                 data.BillOfMaterial,
		BillOfMaterialItem:                             data.BillOfMaterialItem,
		SupplyChainRelationshipID:                      data.SupplyChainRelationshipID,
		SupplyChainRelationshipDeliveryID:              data.SupplyChainRelationshipDeliveryID,
		SupplyChainRelationshipDeliveryPlantID:         data.SupplyChainRelationshipDeliveryPlantID,
		SupplyChainRelationshipStockConfPlantID:        data.SupplyChainRelationshipStockConfPlantID,
		ProductionPlantBusinessPartner:                 data.ProductionPlantBusinessPartner,
		ProductionPlant:                                data.ProductionPlant,
		MRPArea:                                        data.MRPArea,
		MRPController:                                  data.MRPController,
		ProductionVersion:                              data.ProductionVersion,
		ProductionVersionItem:                          data.ProductionVersionItem,
		ComponentProduct:                               data.ComponentProduct,
		ComponentProductBuyer:                          data.ComponentProductBuyer,
		ComponentProductSeller:                         data.ComponentProductSeller,
		ComponentProductDeliverToParty:                 data.ComponentProductDeliverToParty,
		ComponentProductDeliverToPlant:                 data.ComponentProductDeliverToPlant,
		ComponentProductDeliverFromParty:               data.ComponentProductDeliverFromParty,
		ComponentProductDeliverFromPlant:               data.ComponentProductDeliverFromPlant,
		ComponentProductBaseUnit:                       data.ComponentProductBaseUnit,
		ComponentProductDeliveryUnit:                   data.ComponentProductDeliveryUnit,
		ComponentProductRequirementDate:                data.ComponentProductRequirementDate,
		ComponentProductRequirementTime:                data.ComponentProductRequirementTime,
		ComponentProductRequiredQuantityInBaseUnit:     data.ComponentProductRequiredQuantityInBaseUnit,
		ComponentProductRequiredQuantityInDeliveryUnit: data.ComponentProductRequiredQuantityInDeliveryUnit,
		ComponentProductPlannedScrapInPercent:          data.ComponentProductPlannedScrapInPercent,
		ComponentProductIsMarkedForBackflush:           data.ComponentProductIsMarkedForBackflush,
		BillOfMaterialItemText:                         data.BillOfMaterialItemText,
		StockConfirmationBusinessPartner:               data.StockConfirmationBusinessPartner,
		StockConfirmationPlant:                         data.StockConfirmationPlant,
		StockConfirmationPlantStorageLocation:          data.StockConfirmationPlantStorageLocation,
		PlannedOrder:                                   data.PlannedOrder,
		PlannedOrderItem:                               data.PlannedOrderItem,
		BOMItemDescription:                             data.BOMItemDescription,
		ComponentProductBatch:                          data.ComponentProductBatch,
		ComponentProductBatchValidityStartDate:         data.ComponentProductBatchValidityStartDate,
		ComponentProductBatchValidityStartTime:         data.ComponentProductBatchValidityStartTime,
		ComponentProductBatchValidityEndDate:           data.ComponentProductBatchValidityEndDate,
		ComponentProductBatchValidityEndTime:           data.ComponentProductBatchValidityEndTime,
		ComponentProductCostingPolicy:                  data.ComponentProductCostingPolicy,
		ComponentProductPriceUnitQty:                   data.ComponentProductPriceUnitQty,
		ComponentProductStandardPrice:                  data.ComponentProductStandardPrice,
		ComponentProductMovingAveragePrice:             data.ComponentProductMovingAveragePrice,
		ComponentProductWithdrawnQuantity:              data.ComponentProductWithdrawnQuantity,
		CreationDate:                                   data.CreationDate,
		CreationTime:                                   data.CreationTime,
		LastChangeDate:                                 data.LastChangeDate,
		LastChangeTime:                                 data.LastChangeTime,
		ComponentProductAvailabilityIsNotChecked:       data.ComponentProductAvailabilityIsNotChecked,
		IsReleased:                                     data.IsReleased,
		IsLocked:                                       data.IsLocked,
		IsCancelled:                                    data.IsCancelled,
		IsMarkedForDeletion:                            data.IsMarkedForDeletion,
	}
	return networkComponent, nil
}

func ConvertToNetworkComponentDeliveryScheduleLine(sdc *api_input_reader.SDC, rows *sql.Rows) (*[]NetworkComponentDeliveryScheduleLine, error) {
	defer rows.Close()
	pm := &requests.NetworkComponentDeliveryScheduleLine{}

	i := 0
	for rows.Next() {
		i++

		err := rows.Scan(
			&pm.Project,
			&pm.WBSElement,
			&pm.Network,
			&pm.BillOfMaterial,
			&pm.BillOfMaterialItem,
			&pm.ScheduleLine,
			&pm.SupplyChainRelationshipID,
			&pm.SupplyChainRelationshipStockConfPlantID,
			&pm.ComponentProduct,
			&pm.StockConfirmationBusinessPartner,
			&pm.StockConfirmationPlant,
			&pm.StockConfirmationPlantTimeZone,
			&pm.ComponentProductBatch,
			&pm.ComponentProductBatchValidityStartDate,
			&pm.ComponentProductBatchValidityStartTime,
			&pm.ComponentProductBatchValidityEndDate,
			&pm.ComponentProductBatchValidityEndTime,
			&pm.RequestedDeliveryDate,
			&pm.RequestedDeliveryTime,
			&pm.ConfirmedDeliveryDate,
			&pm.ConfirmedDeliveryTime,
			&pm.OriginalRequiredQuantityInBaseUnit,
			&pm.ConfirmedQuantityByPDTAvailCheckInBaseUnit,
			&pm.OpenConfirmedQuantityInBaseUnit,
			&pm.DeliveredQuantityInBaseUnit,
			&pm.UndeliveredQuantityInBaseUnit,
			&pm.StockIsFullyConfirmed,
			&pm.PlusMinusFlag,
			&pm.ItemScheduleLineDeliveryBlockStatus,
			&pm.CreationDate,
			&pm.CreationTime,
			&pm.LastChangeDate,
			&pm.LastChangeTime,
			&pm.IsReleased,
			&pm.IsLocked,
			&pm.IsCancelled,
			&pm.IsMarkedForDeletion,
		)
		if err != nil {
			fmt.Printf("err = %+v \n", err)
			return &networkComponentDeliveryScheduleLine{}, err
		}
	}
	if i == 0 {
		fmt.Printf("DBに対象のレコードが存在しません。")
		return &networkComponentDeliveryScheduleLine{}, nil
	}

	data := pm
	networkComponentDeliveryScheduleLine := &networkComponentDeliveryScheduleLine{
		Project:                                    data.Project,
		WBSElement:                                 data.WBSElement,
		Network:                                    data.Network,
		BillOfMaterial:                             data.BillOfMaterial,
		BillOfMaterialItem:                         data.BillOfMaterialItem,
		ScheduleLine:                               data.ScheduleLine,
		SupplyChainRelationshipID:                  data.SupplyChainRelationshipID,
		SupplyChainRelationshipStockConfPlantID:    data.SupplyChainRelationshipStockConfPlantID,
		ComponentProduct:                           data.ComponentProduct,
		StockConfirmationBusinessPartner:           data.StockConfirmationBusinessPartner,
		StockConfirmationPlant:                     data.StockConfirmationPlant,
		StockConfirmationPlantTimeZone:             data.StockConfirmationPlantTimeZone,
		ComponentProductBatch:                      data.ComponentProductBatch,
		ComponentProductBatchValidityStartDate:     data.ComponentProductBatchValidityStartDate,
		ComponentProductBatchValidityStartTime:     data.ComponentProductBatchValidityStartTime,
		ComponentProductBatchValidityEndDate:       data.ComponentProductBatchValidityEndDate,
		ComponentProductBatchValidityEndTime:       data.ComponentProductBatchValidityEndTime,
		RequestedDeliveryDate:                      data.RequestedDeliveryDate,
		RequestedDeliveryTime:                      data.RequestedDeliveryTime,
		ConfirmedDeliveryDate:                      data.ConfirmedDeliveryDate,
		ConfirmedDeliveryTime:                      data.ConfirmedDeliveryTime,
		OriginalRequiredQuantityInBaseUnit:         data.OriginalRequiredQuantityInBaseUnit,
		ConfirmedQuantityByPDTAvailCheckInBaseUnit: data.ConfirmedQuantityByPDTAvailCheckInBaseUnit,
		OpenConfirmedQuantityInBaseUnit:            data.OpenConfirmedQuantityInBaseUnit,
		DeliveredQuantityInBaseUnit:                data.DeliveredQuantityInBaseUnit,
		UndeliveredQuantityInBaseUnit:              data.UndeliveredQuantityInBaseUnit,
		StockIsFullyConfirmed:                      data.StockIsFullyConfirmed,
		PlusMinusFlag:                              data.PlusMinusFlag,
		ItemScheduleLineDeliveryBlockStatus:        data.ItemScheduleLineDeliveryBlockStatus,
		CreationDate:                               data.CreationDate,
		CreationTime:                               data.CreationTime,
		LastChangeDate:                             data.LastChangeDate,
		LastChangeTime:                             data.LastChangeTime,
		IsReleased:                                 data.IsReleased,
		IsLocked:                                   data.IsLocked,
		IsCancelled:                                data.IsCancelled,
		IsMarkedForDeletion:                        data.IsMarkedForDeletion,
	}
	return networkComponentDeliveryScheduleLine, nil
}
