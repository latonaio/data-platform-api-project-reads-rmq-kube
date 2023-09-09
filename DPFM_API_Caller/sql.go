package dpfm_api_caller

import (
	"context"
	dpfm_api_input_reader "data-platform-api-project-reads-rmq-kube/DPFM_API_Input_Reader"
	dpfm_api_output_formatter "data-platform-api-project-reads-rmq-kube/DPFM_API_Output_Formatter"
	"fmt"
	"sync"

	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
)

func (c *DPFMAPICaller) readSqlProcess(
	ctx context.Context,
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	accepter []string,
	errs *[]error,
	log *logger.Logger,
) interface{} {
	var project *dpfm_api_output_formatter.Project
	var wBSElement *dpfm_api_output_formatter.WBSElement
	var network *dpfm_api_output_formatter.Network
	for _, fn := range accepter {
		switch fn {
		case "Project":
			func() {
				project = c.Project(mtx, input, output, errs, log)
			}()
		case "Projects":
			func() {
				project = c.Projects(mtx, input, output, errs, log)
			}()
		case "ProjectsByBusinessPartners":
			func() {
				project = c.ProjectsByBusinessPartners(mtx, input, output, errs, log)
			}()
		case "WBSElement":
			func() {
				wBSElement = c.WBSElement(mtx, input, output, errs, log)
			}()
		case "Network":
			func() {
				network = c.Network(mtx, input, output, errs, log)
			}()
		default:
		}
	}

	data := &dpfm_api_output_formatter.Message{
		Project:    project,
		WBSElement: wBSElement,
		Network:	network,
	}

	return data
}

func (c *DPFMAPICaller) Project(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.Project {
	where := "WHERE 1 = 1"

	if input.Project.Project != nil {
		where = fmt.Sprintf("%s\nAND Project = %d", where, *input.Project.Project)
	}

	if input.Project.IsMarkedForDeletion != nil {
		where = fmt.Sprintf("%s\nAND IsMarkedForDeletion = %t", where, *input.Project.IsMarkedForDeletion)
	}

	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_project_project_data
		` + where + `;`,
	)

	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToProject(input, rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return &((*data)[0])
}

func (c *DPFMAPICaller) Projects(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.Project {
	where := "WHERE 1 = 1"

	if input.Project.BusinessPartner != nil {
		where = fmt.Sprintf("%s\nAND BusinessPartner = %v", where, *input.Project.BusinessPartner)
	}

	if input.Project.IsMarkedForDeletion != nil {
		where = fmt.Sprintf("%s\nAND IsMarkedForDeletion = %t", where, *input.Project.IsMarkedForDeletion)
	}

	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_project_project_data
		` + where + `;`,
	)

	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToProject(input, rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return &((*data)[0])
}

func (c *DPFMAPICaller) ProjectsByBusinessPartners(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *[]dpfm_api_output_formatter.Project {
	log.Info("ProjectsByBusinessPartners")
	in := ""

	for iProject, vProject := range input.Projects {
		businessPartner := vProject.BusinessPartner
		if iProject == 0 {
			in = fmt.Sprintf(
				"( '%d' )",
				businessPartner,
			)
			continue
		}
		in = fmt.Sprintf(
			"%s ,( '%d' )",
			in,
			businessPartner,
		)
	}

	where := fmt.Sprintf(" WHERE ( BusinessPartner ) IN ( %s ) ", in)
	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_project_project_data
		` + where + ` ;`,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToProject(rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) WBSElement(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *dpfm_api_output_formatter.WBSElement {
	where := "WHERE 1 = 1"

	if input.Project.Project != nil {
		where = fmt.Sprintf("%s\nAND Project = %d", where, *input.Project.Project)
	}
	if input.Project.WBSElement != nil {
		where = fmt.Sprintf("%s\nAND WBSElement = %d", where, *input.Project.WBSElement.WBSElement)
	}
	if input.Project.WBSElement.IsMarkedForDeletion != nil {
		where = fmt.Sprintf("%s\nAND IsMarkedForDeletion = %t", where, *input.Project.WBSElement.IsMarkedForDeletion)
	}
	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_project_wbs_element_data
		` + where + `;`,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToWBSElement(input, rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}

func (c *DPFMAPICaller) Network(
	mtx *sync.Mutex,
	input *dpfm_api_input_reader.SDC,
	output *dpfm_api_output_formatter.SDC,
	errs *[]error,
	log *logger.Logger,
) *dpfm_api_output_formatter.Network {
	where := "WHERE 1 = 1"

	if input.Project.Project != nil {
		where = fmt.Sprintf("%s\nAND Project = %d", where, *input.Project.Project)
	}
	if input.Project.WBSElement != nil {
		where = fmt.Sprintf("%s\nAND WBSElement = %d", where, *input.Project.WBSElement.WBSElement)
	}
	if input.Project.WBSElement.Network != nil {
		where = fmt.Sprintf("%s\nAND Network = %d", where, *input.Project.WBSElement.Network.Network)
	}
	if input.Project.WBSElement.Network.IsMarkedForDeletion != nil {
		where = fmt.Sprintf("%s\nAND IsMarkedForDeletion = %t", where, *input.Project.WBSElement.Network.IsMarkedForDeletion)
	}
	rows, err := c.db.Query(
		`SELECT *
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_project_network_data
		` + where + `;`,
	)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}
	defer rows.Close()

	data, err := dpfm_api_output_formatter.ConvertToNetwork(input, rows)
	if err != nil {
		*errs = append(*errs, err)
		return nil
	}

	return data
}
