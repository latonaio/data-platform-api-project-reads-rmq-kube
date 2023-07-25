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
			Project:				data.Project,
			ProjectCode:			data.ProjectCode,
			ProjectDescription:		data.ProjectDescription,
			OwnerBusinessPartner:	data.OwnerBusinessPartner,
			OwnerPlant:				data.OwnerPlant,
			ProjectProfile:			data.ProjectProfile,
			ResponsiblePerson:		data.ResponsiblePerson,
			ResponsiblePersonName:	data.ResponsiblePersonName,
			PlannedStartDate:		data.PlannedStartDate,
			PlannedEndDate:			data.PlannedEndDate,
			ActualStartDate:		data.ActualStartDate,
			ActualEndDate:			data.ActualEndDate,
			CreationDate:			data.CreationDate,
			LastChangeDate:			data.LastChangeDate,
			IsMarkedForDeletion:	data.IsMarkedForDeletion,
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
			Project:				data.Project,
			WBSElement:				data.WBSElement,
			WBSElementCode:			data.WBSElementCode,
			WBSElementDescription:	data.WBSElementDescription,
			BusinessPartner:		data.BusinessPartner,
			Plant:					data.Plant,
			ResponsiblePerson:		data.ResponsiblePerson,
			ResponsiblePersonName:	data.ResponsiblePersonName,
			PlannedStartDate:		data.PlannedStartDate,
			PlannedEndDate:			data.PlannedEndDate,
			ActualStartDate:		data.ActualStartDate,
			ActualEndDate:			data.ActualEndDate,
			CreationDate:			data.CreationDate,
			LastChangeDate:			data.LastChangeDate,
			IsMarkedForDeletion:	data.IsMarkedForDeletion,
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
			Project:				data.Project,
			WBSElement:				data.WBSElement,
			Network:				data.Network,
			NetworkDescription:		data.NetworkDescription,
			BusinessPartner:		data.BusinessPartner,
			Plant:					data.Plant,
			ResponsiblePerson:		data.ResponsiblePerson,
			ResponsiblePersonName:	data.ResponsiblePersonName,
			PlannedStartDate:		data.PlannedStartDate,
			PlannedEndDate:			data.PlannedEndDate,
			ActualStartDate:		data.ActualStartDate,
			ActualEndDate:			data.ActualEndDate,
			CreationDate:			data.CreationDate,
			CreationTime:			data.CreationTime,
			LastChangeDate:			data.LastChangeDate,
			LastChangeTime:			data.LastChangeTime,
			IsMarkedForDeletion:	data.IsMarkedForDeletion,
	}
	return network, nil
}
