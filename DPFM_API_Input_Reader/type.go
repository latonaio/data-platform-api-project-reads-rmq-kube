package dpfm_api_input_reader

type EC_MC struct {
	ConnectionKey string `json:"connection_key"`
	Result        bool   `json:"result"`
	RedisKey      string `json:"redis_key"`
	Filepath      string `json:"filepath"`
	Document      struct {
		DocumentNo     string `json:"document_no"`
		DeliverTo      string `json:"deliver_to"`
		Quantity       string `json:"quantity"`
		PickedQuantity string `json:"picked_quantity"`
		Price          string `json:"price"`
		Batch          string `json:"batch"`
	} `json:"document"`
	BusinessPartner struct {
		DocumentNo           string `json:"document_no"`
		Status               string `json:"status"`
		DeliverTo            string `json:"deliver_to"`
		Quantity             string `json:"quantity"`
		CompletedQuantity    string `json:"completed_quantity"`
		PlannedStartDate     string `json:"planned_start_date"`
		PlannedValidatedDate string `json:"planned_validated_date"`
		ActualStartDate      string `json:"actual_start_date"`
		ActualValidatedDate  string `json:"actual_validated_date"`
		Batch                string `json:"batch"`
		Work                 struct {
			WorkNo                   string `json:"work_no"`
			Quantity                 string `json:"quantity"`
			CompletedQuantity        string `json:"completed_quantity"`
			ErroredQuantity          string `json:"errored_quantity"`
			Component                string `json:"component"`
			PlannedComponentQuantity string `json:"planned_component_quantity"`
			PlannedStartDate         string `json:"planned_start_date"`
			PlannedStartTime         string `json:"planned_start_time"`
			PlannedValidatedDate     string `json:"planned_validated_date"`
			PlannedValidatedTime     string `json:"planned_validated_time"`
			ActualStartDate          string `json:"actual_start_date"`
			ActualStartTime          string `json:"actual_start_time"`
			ActualValidatedDate      string `json:"actual_validated_date"`
			ActualValidatedTime      string `json:"actual_validated_time"`
		} `json:"work"`
	} `json:"business_partner"`
	APISchema     string   `json:"api_schema"`
	Accepter      []string `json:"accepter"`
	MaterialCode  string   `json:"material_code"`
	Project       string   `json:"project/supplier"`
	Stock         string   `json:"stock"`
	DocumentType  string   `json:"document_type"`
	DocumentNo    string   `json:"document_no"`
	PlannedDate   string   `json:"planned_date"`
	ValidatedDate string   `json:"validated_date"`
	Deleted       bool     `json:"deleted"`
}

type SDC struct {
	ConnectionKey     string   `json:"connection_key"`
	Result            bool     `json:"result"`
	RedisKey          string   `json:"redis_key"`
	Filepath          string   `json:"filepath"`
	APIStatusCode     int      `json:"api_status_code"`
	RuntimeSessionID  string   `json:"runtime_session_id"`
	BusinessPartnerID *int     `json:"business_partner"`
	ServiceLabel      string   `json:"service_label"`
	APIType           string   `json:"api_type"`
	Project           Project  `json:"Project"`
	Projects          Projects `json:"Projects"`
	APISchema         string   `json:"api_schema"`
	Accepter          []string `json:"accepter"`
	Deleted           bool     `json:"deleted"`
}

type Project struct {
	Project               int          `json:"Project"`
	ProjectCode           *string      `json:"ProjectCode"`
	ProjectDescription    *string      `json:"ProjectDescription"`
	OwnerBusinessPartner  *int         `json:"OwnerBusinessPartner"`
	OwnerPlant            *string      `json:"OwnerPlant"`
	ProjectProfile        *string      `json:"ProjectProfile"`
	ResponsiblePerson     *int         `json:"ResponsiblePerson"`
	ResponsiblePersonName *string      `json:"ResponsiblePersonName"`
	PlannedStartDate      *string      `json:"PlannedStartDate"`
	PlannedEndDate        *string      `json:"PlannedEndDate"`
	ActualStartDate       *string      `json:"ActualStartDate"`
	ActualEndDate         *string      `json:"ActualEndDate"`
	CreationDate          *string      `json:"CreationDate"`
	LastChangeDate        *string      `json:"LastChangeDate"`
	IsMarkedForDeletion   *bool        `json:"IsMarkedForDeletion"`
	WBSElement            []WBSElement `json:"WBSElement"`
}

type Projects []struct {
	Project               int          `json:"Project"`
	ProjectCode           *string      `json:"ProjectCode"`
	ProjectDescription    *string      `json:"ProjectDescription"`
	OwnerBusinessPartner  *int         `json:"OwnerBusinessPartner"`
	OwnerPlant            *string      `json:"OwnerPlant"`
	ProjectProfile        *string      `json:"ProjectProfile"`
	ResponsiblePerson     *int         `json:"ResponsiblePerson"`
	ResponsiblePersonName *string      `json:"ResponsiblePersonName"`
	PlannedStartDate      *string      `json:"PlannedStartDate"`
	PlannedEndDate        *string      `json:"PlannedEndDate"`
	ActualStartDate       *string      `json:"ActualStartDate"`
	ActualEndDate         *string      `json:"ActualEndDate"`
	CreationDate          *string      `json:"CreationDate"`
	LastChangeDate        *string      `json:"LastChangeDate"`
	IsMarkedForDeletion   *bool        `json:"IsMarkedForDeletion"`
	WBSElement            []WBSElement `json:"WBSElement"`
}

type WBSElement struct {
	Project               int       `json:"Project"`
	WBSElement            int       `json:"WBSElement"`
	WBSElementCode        *string   `json:"WBSElementCode"`
	WBSElementDescription *string   `json:"WBSElementDescription"`
	BusinessPartner       *int      `json:"BusinessPartner"`
	Plant                 *string   `json:"Plant"`
	ResponsiblePerson     *int      `json:"ResponsiblePerson"`
	ResponsiblePersonName *string   `json:"ResponsiblePersonName"`
	PlannedStartDate      *string   `json:"PlannedStartDate"`
	PlannedEndDate        *string   `json:"PlannedEndDate"`
	ActualStartDate       *string   `json:"ActualStartDate"`
	ActualEndDate         *string   `json:"ActualEndDate"`
	CreationDate          *string   `json:"CreationDate"`
	LastChangeDate        *string   `json:"LastChangeDate"`
	IsMarkedForDeletion   *bool     `json:"IsMarkedForDeletion"`
	Network               []Network `json:"Network"`
}

type Network struct {
	Project               int     `json:"Project"`
	WBSElement            int     `json:"WBSElement"`
	Network               int     `json:"Network"`
	NetworkDescription    *string `json:"NetworkDescription"`
	BusinessPartner       *int    `json:"BusinessPartner"`
	Plant                 *string `json:"Plant"`
	ResponsiblePerson     *int    `json:"ResponsiblePerson"`
	ResponsiblePersonName *string `json:"ResponsiblePersonName"`
	PlannedStartDate      *string `json:"PlannedStartDate"`
	PlannedEndDate        *string `json:"PlannedEndDate"`
	ActualStartDate       *string `json:"ActualStartDate"`
	ActualEndDate         *string `json:"ActualEndDate"`
	CreationDate          *string `json:"CreationDate"`
	CreationTime          *string `json:"CreationTime"`
	LastChangeDate        *string `json:"LastChangeDate"`
	LastChangeTime        *string `json:"LastChangeTime"`
	IsReleased            *bool   `json:"IsReleased"`
	IsLocked              *bool   `json:"IsLocked"`
	IsCancelled           *bool   `json:"IsCancelled"`
	IsMarkedForDeletion   *bool   `json:"IsMarkedForDeletion"`
	NetworkComponent	  []NetworkComponent `json:"NetworkComponent"`
}

type NetworkComponent struct {
	Project                                        int      `json:"Project"`
	WBSElement                                     int      `json:"WBSElement"`
	Network                                        int      `json:"Network"`
	BillOfMaterial                                 *int     `json:"BillOfMaterial"`
	BillOfMaterialItem                             *int     `json:"BillOfMaterialItem"`
	SupplyChainRelationshipID                      *int     `json:"SupplyChainRelationshipID"`
	SupplyChainRelationshipDeliveryID              *int     `json:"SupplyChainRelationshipDeliveryID"`
	SupplyChainRelationshipDeliveryPlantID         *int     `json:"SupplyChainRelationshipDeliveryPlantID"`
	SupplyChainRelationshipStockConfPlantID        *int     `json:"SupplyChainRelationshipStockConfPlantID"`
	ProductionPlantBusinessPartner                 *int     `json:"ProductionPlantBusinessPartner"`
	ProductionPlant                                *string  `json:"ProductionPlant"`
	MRPArea                                        *string  `json:"MRPArea"`
	MRPController                                  *string  `json:"MRPController"`
	ProductionVersion                              *int     `json:"ProductionVersion"`
	ProductionVersionItem                          *int     `json:"ProductionVersionItem"`
	ComponentProduct                               *string  `json:"ComponentProduct"`
	ComponentProductBuyer                          *int     `json:"ComponentProductBuyer"`
	ComponentProductSeller                         *int     `json:"ComponentProductSeller"`
	ComponentProductDeliverToParty                 *int     `json:"ComponentProductDeliverToParty"`
	ComponentProductDeliverToPlant                 *string  `json:"ComponentProductDeliverToPlant"`
	ComponentProductDeliverFromParty               *int     `json:"ComponentProductDeliverFromParty"`
	ComponentProductDeliverFromPlant               *string  `json:"ComponentProductDeliverFromPlant"`
	ComponentProductBaseUnit                       *string  `json:"ComponentProductBaseUnit"`
	ComponentProductDeliveryUnit                   *string  `json:"ComponentProductDeliveryUnit"`
	ComponentProductRequirementDate                *string  `json:"ComponentProductRequirementDate"`
	ComponentProductRequirementTime                *string  `json:"ComponentProductRequirementTime"`
	ComponentProductRequiredQuantityInBaseUnit     *float32 `json:"ComponentProductRequiredQuantityInBaseUnit"`
	ComponentProductRequiredQuantityInDeliveryUnit *float32 `json:"ComponentProductRequiredQuantityInDeliveryUnit"`
	ComponentProductPlannedScrapInPercent          *float32 `json:"ComponentProductPlannedScrapInPercent"`
	ComponentProductIsMarkedForBackflush           *bool    `json:"ComponentProductIsMarkedForBackflush"`
	BillOfMaterialItemText                         *string  `json:"BillOfMaterialItemText"`
	StockConfirmationBusinessPartner               *int     `json:"StockConfirmationBusinessPartner"`
	StockConfirmationPlant                         *string  `json:"StockConfirmationPlant"`
	StockConfirmationPlantStorageLocation          *string  `json:"StockConfirmationPlantStorageLocation"`
	PlannedOrder                                   *int     `json:"PlannedOrder"`
	PlannedOrderItem                               *int     `json:"PlannedOrderItem"`
	BOMItemDescription                             *string  `json:"BOMItemDescription"`
	ComponentProductBatch                          *string  `json:"ComponentProductBatch"`
	ComponentProductBatchValidityStartDate         *string  `json:"ComponentProductBatchValidityStartDate"`
	ComponentProductBatchValidityStartTime         *string  `json:"ComponentProductBatchValidityStartTime"`
	ComponentProductBatchValidityEndDate           *string  `json:"ComponentProductBatchValidityEndDate"`
	ComponentProductBatchValidityEndTime           *string  `json:"ComponentProductBatchValidityEndTime"`
	ComponentProductCostingPolicy                  *string  `json:"ComponentProductCostingPolicy"`
	ComponentProductPriceUnitQty                   *int     `json:"ComponentProductPriceUnitQty"`
	ComponentProductStandardPrice                  *float32 `json:"ComponentProductStandardPrice"`
	ComponentProductMovingAveragePrice             *float32 `json:"ComponentProductMovingAveragePrice"`
	ComponentProductWithdrawnQuantity              *float32 `json:"ComponentProductWithdrawnQuantity"`
	CreationDate                                   *string  `json:"CreationDate"`
	CreationTime                                   *string  `json:"CreationTime"`
	LastChangeDate                                 *string  `json:"LastChangeDate"`
	LastChangeTime                                 *string  `json:"LastChangeTime"`
	ComponentProductAvailabilityIsNotChecked       *bool    `json:"ComponentProductAvailabilityIsNotChecked"`
	IsReleased                                     *bool    `json:"IsReleased"`
	IsLocked                                       *bool    `json:"IsLocked"`
	IsCancelled                                    *bool    `json:"IsCancelled"`
	IsMarkedForDeletion                            *bool    `json:"IsMarkedForDeletion"`
	NetworkComponentDeliveryScheduleLine		   []NetworkComponentDeliveryScheduleLine `json:"NetworkComponentDeliveryScheduleLine"`
}

type NetworkComponentDeliveryScheduleLine struct {
	Project                                    int      `json:"Project"`
	WBSElement                                 int      `json:"WBSElement"`
	Network                                    int      `json:"Network"`
	BillOfMaterial                             int      `json:"BillOfMaterial"`
	BillOfMaterialItem                         int      `json:"BillOfMaterialItem"`
	ScheduleLine                               int      `json:"ScheduleLine"`
	SupplyChainRelationshipID                  *int     `json:"SupplyChainRelationshipID"`
	SupplyChainRelationshipStockConfPlantID    *int     `json:"SupplyChainRelationshipStockConfPlantID"`
	ComponentProduct                           *string  `json:"ComponentProduct"`
	StockConfirmationBusinessPartner           *int     `json:"StockConfirmationBusinessPartner"`
	StockConfirmationPlant                     *string  `json:"StockConfirmationPlant"`
	StockConfirmationPlantTimeZone             *string  `json:"StockConfirmationPlantTimeZone"`
	ComponentProductBatch                      *string  `json:"ComponentProductBatch"`
	ComponentProductBatchValidityStartDate     *string  `json:"ComponentProductBatchValidityStartDate"`
	ComponentProductBatchValidityStartTime     *string  `json:"ComponentProductBatchValidityStartTime"`
	ComponentProductBatchValidityEndDate       *string  `json:"ComponentProductBatchValidityEndDate"`
	ComponentProductBatchValidityEndTime       *string  `json:"ComponentProductBatchValidityEndTime"`
	RequestedDeliveryDate                      *string  `json:"RequestedDeliveryDate"`
	RequestedDeliveryTime                      *string  `json:"RequestedDeliveryTime"`
	ConfirmedDeliveryDate                      *string  `json:"ConfirmedDeliveryDate"`
	ConfirmedDeliveryTime                      *string  `json:"ConfirmedDeliveryTime"`
	OriginalRequiredQuantityInBaseUnit         *float32 `json:"OriginalRequiredQuantityInBaseUnit"`
	ConfirmedQuantityByPDTAvailCheckInBaseUnit *float32 `json:"ConfirmedQuantityByPDTAvailCheckInBaseUnit"`
	OpenConfirmedQuantityInBaseUnit            *float32 `json:"OpenConfirmedQuantityInBaseUnit"`
	DeliveredQuantityInBaseUnit                *float32 `json:"DeliveredQuantityInBaseUnit"`
	UndeliveredQuantityInBaseUnit              *float32 `json:"UndeliveredQuantityInBaseUnit"`
	StockIsFullyConfirmed                      *bool    `json:"StockIsFullyConfirmed"`
	PlusMinusFlag                              *string  `json:"PlusMinusFlag"`
	ItemScheduleLineDeliveryBlockStatus        *bool    `json:"ItemScheduleLineDeliveryBlockStatus"`
	CreationDate                               *string  `json:"CreationDate"`
	CreationTime                               *string  `json:"CreationTime"`
	LastChangeDate                             *string  `json:"LastChangeDate"`
	LastChangeTime                             *string  `json:"LastChangeTime"`
	IsReleased                                 *bool    `json:"IsReleased"`
	IsLocked                                   *bool    `json:"IsLocked"`
	IsCancelled                                *bool    `json:"IsCancelled"`
	IsMarkedForDeletion                        *bool    `json:"IsMarkedForDeletion"`
}
