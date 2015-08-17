// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

package machinelearning_test

import (
	"bytes"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/awsutil"
	"github.com/aws/aws-sdk-go/service/machinelearning"
)

var _ time.Duration
var _ bytes.Buffer

func ExampleMachineLearning_CreateBatchPrediction() {
	svc := machinelearning.New(nil)

	params := &machinelearning.CreateBatchPredictionInput{
		BatchPredictionDataSourceId: aws.String("EntityId"), // Required
		BatchPredictionId:           aws.String("EntityId"), // Required
		MLModelId:                   aws.String("EntityId"), // Required
		OutputUri:                   aws.String("S3Url"),    // Required
		BatchPredictionName:         aws.String("EntityName"),
	}
	resp, err := svc.CreateBatchPrediction(params)

	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			// Generic AWS error with Code, Message, and original error (if any)
			fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
			if reqErr, ok := err.(awserr.RequestFailure); ok {
				// A service error occurred
				fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
			}
		} else {
			// This case should never be hit, the SDK should always return an
			// error which satisfies the awserr.Error interface.
			fmt.Println(err.Error())
		}
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.Prettify(resp))
}

func ExampleMachineLearning_CreateDataSourceFromRDS() {
	svc := machinelearning.New(nil)

	params := &machinelearning.CreateDataSourceFromRDSInput{
		DataSourceId: aws.String("EntityId"), // Required
		RDSData: &machinelearning.RDSDataSpec{ // Required
			DatabaseCredentials: &machinelearning.RDSDatabaseCredentials{ // Required
				Password: aws.String("RDSDatabasePassword"), // Required
				Username: aws.String("RDSDatabaseUsername"), // Required
			},
			DatabaseInformation: &machinelearning.RDSDatabase{ // Required
				DatabaseName:       aws.String("RDSDatabaseName"),       // Required
				InstanceIdentifier: aws.String("RDSInstanceIdentifier"), // Required
			},
			ResourceRole:      aws.String("EDPResourceRole"), // Required
			S3StagingLocation: aws.String("S3Url"),           // Required
			SecurityGroupIds: []*string{ // Required
				aws.String("EDPSecurityGroupId"), // Required
				// More values...
			},
			SelectSqlQuery:    aws.String("RDSSelectSqlQuery"), // Required
			ServiceRole:       aws.String("EDPServiceRole"),    // Required
			SubnetId:          aws.String("EDPSubnetId"),       // Required
			DataRearrangement: aws.String("DataRearrangement"),
			DataSchema:        aws.String("DataSchema"),
			DataSchemaUri:     aws.String("S3Url"),
		},
		RoleARN:           aws.String("RoleARN"), // Required
		ComputeStatistics: aws.Bool(true),
		DataSourceName:    aws.String("EntityName"),
	}
	resp, err := svc.CreateDataSourceFromRDS(params)

	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			// Generic AWS error with Code, Message, and original error (if any)
			fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
			if reqErr, ok := err.(awserr.RequestFailure); ok {
				// A service error occurred
				fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
			}
		} else {
			// This case should never be hit, the SDK should always return an
			// error which satisfies the awserr.Error interface.
			fmt.Println(err.Error())
		}
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.Prettify(resp))
}

func ExampleMachineLearning_CreateDataSourceFromRedshift() {
	svc := machinelearning.New(nil)

	params := &machinelearning.CreateDataSourceFromRedshiftInput{
		DataSourceId: aws.String("EntityId"), // Required
		DataSpec: &machinelearning.RedshiftDataSpec{ // Required
			DatabaseCredentials: &machinelearning.RedshiftDatabaseCredentials{ // Required
				Password: aws.String("RedshiftDatabasePassword"), // Required
				Username: aws.String("RedshiftDatabaseUsername"), // Required
			},
			DatabaseInformation: &machinelearning.RedshiftDatabase{ // Required
				ClusterIdentifier: aws.String("RedshiftClusterIdentifier"), // Required
				DatabaseName:      aws.String("RedshiftDatabaseName"),      // Required
			},
			S3StagingLocation: aws.String("S3Url"),                  // Required
			SelectSqlQuery:    aws.String("RedshiftSelectSqlQuery"), // Required
			DataRearrangement: aws.String("DataRearrangement"),
			DataSchema:        aws.String("DataSchema"),
			DataSchemaUri:     aws.String("S3Url"),
		},
		RoleARN:           aws.String("RoleARN"), // Required
		ComputeStatistics: aws.Bool(true),
		DataSourceName:    aws.String("EntityName"),
	}
	resp, err := svc.CreateDataSourceFromRedshift(params)

	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			// Generic AWS error with Code, Message, and original error (if any)
			fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
			if reqErr, ok := err.(awserr.RequestFailure); ok {
				// A service error occurred
				fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
			}
		} else {
			// This case should never be hit, the SDK should always return an
			// error which satisfies the awserr.Error interface.
			fmt.Println(err.Error())
		}
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.Prettify(resp))
}

func ExampleMachineLearning_CreateDataSourceFromS3() {
	svc := machinelearning.New(nil)

	params := &machinelearning.CreateDataSourceFromS3Input{
		DataSourceId: aws.String("EntityId"), // Required
		DataSpec: &machinelearning.S3DataSpec{ // Required
			DataLocationS3:       aws.String("S3Url"), // Required
			DataRearrangement:    aws.String("DataRearrangement"),
			DataSchema:           aws.String("DataSchema"),
			DataSchemaLocationS3: aws.String("S3Url"),
		},
		ComputeStatistics: aws.Bool(true),
		DataSourceName:    aws.String("EntityName"),
	}
	resp, err := svc.CreateDataSourceFromS3(params)

	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			// Generic AWS error with Code, Message, and original error (if any)
			fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
			if reqErr, ok := err.(awserr.RequestFailure); ok {
				// A service error occurred
				fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
			}
		} else {
			// This case should never be hit, the SDK should always return an
			// error which satisfies the awserr.Error interface.
			fmt.Println(err.Error())
		}
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.Prettify(resp))
}

func ExampleMachineLearning_CreateEvaluation() {
	svc := machinelearning.New(nil)

	params := &machinelearning.CreateEvaluationInput{
		EvaluationDataSourceId: aws.String("EntityId"), // Required
		EvaluationId:           aws.String("EntityId"), // Required
		MLModelId:              aws.String("EntityId"), // Required
		EvaluationName:         aws.String("EntityName"),
	}
	resp, err := svc.CreateEvaluation(params)

	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			// Generic AWS error with Code, Message, and original error (if any)
			fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
			if reqErr, ok := err.(awserr.RequestFailure); ok {
				// A service error occurred
				fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
			}
		} else {
			// This case should never be hit, the SDK should always return an
			// error which satisfies the awserr.Error interface.
			fmt.Println(err.Error())
		}
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.Prettify(resp))
}

func ExampleMachineLearning_CreateMLModel() {
	svc := machinelearning.New(nil)

	params := &machinelearning.CreateMLModelInput{
		MLModelId:            aws.String("EntityId"),    // Required
		MLModelType:          aws.String("MLModelType"), // Required
		TrainingDataSourceId: aws.String("EntityId"),    // Required
		MLModelName:          aws.String("EntityName"),
		Parameters: map[string]*string{
			"Key": aws.String("StringType"), // Required
			// More values...
		},
		Recipe:    aws.String("Recipe"),
		RecipeUri: aws.String("S3Url"),
	}
	resp, err := svc.CreateMLModel(params)

	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			// Generic AWS error with Code, Message, and original error (if any)
			fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
			if reqErr, ok := err.(awserr.RequestFailure); ok {
				// A service error occurred
				fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
			}
		} else {
			// This case should never be hit, the SDK should always return an
			// error which satisfies the awserr.Error interface.
			fmt.Println(err.Error())
		}
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.Prettify(resp))
}

func ExampleMachineLearning_CreateRealtimeEndpoint() {
	svc := machinelearning.New(nil)

	params := &machinelearning.CreateRealtimeEndpointInput{
		MLModelId: aws.String("EntityId"), // Required
	}
	resp, err := svc.CreateRealtimeEndpoint(params)

	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			// Generic AWS error with Code, Message, and original error (if any)
			fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
			if reqErr, ok := err.(awserr.RequestFailure); ok {
				// A service error occurred
				fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
			}
		} else {
			// This case should never be hit, the SDK should always return an
			// error which satisfies the awserr.Error interface.
			fmt.Println(err.Error())
		}
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.Prettify(resp))
}

func ExampleMachineLearning_DeleteBatchPrediction() {
	svc := machinelearning.New(nil)

	params := &machinelearning.DeleteBatchPredictionInput{
		BatchPredictionId: aws.String("EntityId"), // Required
	}
	resp, err := svc.DeleteBatchPrediction(params)

	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			// Generic AWS error with Code, Message, and original error (if any)
			fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
			if reqErr, ok := err.(awserr.RequestFailure); ok {
				// A service error occurred
				fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
			}
		} else {
			// This case should never be hit, the SDK should always return an
			// error which satisfies the awserr.Error interface.
			fmt.Println(err.Error())
		}
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.Prettify(resp))
}

func ExampleMachineLearning_DeleteDataSource() {
	svc := machinelearning.New(nil)

	params := &machinelearning.DeleteDataSourceInput{
		DataSourceId: aws.String("EntityId"), // Required
	}
	resp, err := svc.DeleteDataSource(params)

	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			// Generic AWS error with Code, Message, and original error (if any)
			fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
			if reqErr, ok := err.(awserr.RequestFailure); ok {
				// A service error occurred
				fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
			}
		} else {
			// This case should never be hit, the SDK should always return an
			// error which satisfies the awserr.Error interface.
			fmt.Println(err.Error())
		}
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.Prettify(resp))
}

func ExampleMachineLearning_DeleteEvaluation() {
	svc := machinelearning.New(nil)

	params := &machinelearning.DeleteEvaluationInput{
		EvaluationId: aws.String("EntityId"), // Required
	}
	resp, err := svc.DeleteEvaluation(params)

	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			// Generic AWS error with Code, Message, and original error (if any)
			fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
			if reqErr, ok := err.(awserr.RequestFailure); ok {
				// A service error occurred
				fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
			}
		} else {
			// This case should never be hit, the SDK should always return an
			// error which satisfies the awserr.Error interface.
			fmt.Println(err.Error())
		}
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.Prettify(resp))
}

func ExampleMachineLearning_DeleteMLModel() {
	svc := machinelearning.New(nil)

	params := &machinelearning.DeleteMLModelInput{
		MLModelId: aws.String("EntityId"), // Required
	}
	resp, err := svc.DeleteMLModel(params)

	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			// Generic AWS error with Code, Message, and original error (if any)
			fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
			if reqErr, ok := err.(awserr.RequestFailure); ok {
				// A service error occurred
				fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
			}
		} else {
			// This case should never be hit, the SDK should always return an
			// error which satisfies the awserr.Error interface.
			fmt.Println(err.Error())
		}
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.Prettify(resp))
}

func ExampleMachineLearning_DeleteRealtimeEndpoint() {
	svc := machinelearning.New(nil)

	params := &machinelearning.DeleteRealtimeEndpointInput{
		MLModelId: aws.String("EntityId"), // Required
	}
	resp, err := svc.DeleteRealtimeEndpoint(params)

	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			// Generic AWS error with Code, Message, and original error (if any)
			fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
			if reqErr, ok := err.(awserr.RequestFailure); ok {
				// A service error occurred
				fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
			}
		} else {
			// This case should never be hit, the SDK should always return an
			// error which satisfies the awserr.Error interface.
			fmt.Println(err.Error())
		}
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.Prettify(resp))
}

func ExampleMachineLearning_DescribeBatchPredictions() {
	svc := machinelearning.New(nil)

	params := &machinelearning.DescribeBatchPredictionsInput{
		EQ:             aws.String("ComparatorValue"),
		FilterVariable: aws.String("BatchPredictionFilterVariable"),
		GE:             aws.String("ComparatorValue"),
		GT:             aws.String("ComparatorValue"),
		LE:             aws.String("ComparatorValue"),
		LT:             aws.String("ComparatorValue"),
		Limit:          aws.Int64(1),
		NE:             aws.String("ComparatorValue"),
		NextToken:      aws.String("StringType"),
		Prefix:         aws.String("ComparatorValue"),
		SortOrder:      aws.String("SortOrder"),
	}
	resp, err := svc.DescribeBatchPredictions(params)

	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			// Generic AWS error with Code, Message, and original error (if any)
			fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
			if reqErr, ok := err.(awserr.RequestFailure); ok {
				// A service error occurred
				fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
			}
		} else {
			// This case should never be hit, the SDK should always return an
			// error which satisfies the awserr.Error interface.
			fmt.Println(err.Error())
		}
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.Prettify(resp))
}

func ExampleMachineLearning_DescribeDataSources() {
	svc := machinelearning.New(nil)

	params := &machinelearning.DescribeDataSourcesInput{
		EQ:             aws.String("ComparatorValue"),
		FilterVariable: aws.String("DataSourceFilterVariable"),
		GE:             aws.String("ComparatorValue"),
		GT:             aws.String("ComparatorValue"),
		LE:             aws.String("ComparatorValue"),
		LT:             aws.String("ComparatorValue"),
		Limit:          aws.Int64(1),
		NE:             aws.String("ComparatorValue"),
		NextToken:      aws.String("StringType"),
		Prefix:         aws.String("ComparatorValue"),
		SortOrder:      aws.String("SortOrder"),
	}
	resp, err := svc.DescribeDataSources(params)

	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			// Generic AWS error with Code, Message, and original error (if any)
			fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
			if reqErr, ok := err.(awserr.RequestFailure); ok {
				// A service error occurred
				fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
			}
		} else {
			// This case should never be hit, the SDK should always return an
			// error which satisfies the awserr.Error interface.
			fmt.Println(err.Error())
		}
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.Prettify(resp))
}

func ExampleMachineLearning_DescribeEvaluations() {
	svc := machinelearning.New(nil)

	params := &machinelearning.DescribeEvaluationsInput{
		EQ:             aws.String("ComparatorValue"),
		FilterVariable: aws.String("EvaluationFilterVariable"),
		GE:             aws.String("ComparatorValue"),
		GT:             aws.String("ComparatorValue"),
		LE:             aws.String("ComparatorValue"),
		LT:             aws.String("ComparatorValue"),
		Limit:          aws.Int64(1),
		NE:             aws.String("ComparatorValue"),
		NextToken:      aws.String("StringType"),
		Prefix:         aws.String("ComparatorValue"),
		SortOrder:      aws.String("SortOrder"),
	}
	resp, err := svc.DescribeEvaluations(params)

	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			// Generic AWS error with Code, Message, and original error (if any)
			fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
			if reqErr, ok := err.(awserr.RequestFailure); ok {
				// A service error occurred
				fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
			}
		} else {
			// This case should never be hit, the SDK should always return an
			// error which satisfies the awserr.Error interface.
			fmt.Println(err.Error())
		}
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.Prettify(resp))
}

func ExampleMachineLearning_DescribeMLModels() {
	svc := machinelearning.New(nil)

	params := &machinelearning.DescribeMLModelsInput{
		EQ:             aws.String("ComparatorValue"),
		FilterVariable: aws.String("MLModelFilterVariable"),
		GE:             aws.String("ComparatorValue"),
		GT:             aws.String("ComparatorValue"),
		LE:             aws.String("ComparatorValue"),
		LT:             aws.String("ComparatorValue"),
		Limit:          aws.Int64(1),
		NE:             aws.String("ComparatorValue"),
		NextToken:      aws.String("StringType"),
		Prefix:         aws.String("ComparatorValue"),
		SortOrder:      aws.String("SortOrder"),
	}
	resp, err := svc.DescribeMLModels(params)

	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			// Generic AWS error with Code, Message, and original error (if any)
			fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
			if reqErr, ok := err.(awserr.RequestFailure); ok {
				// A service error occurred
				fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
			}
		} else {
			// This case should never be hit, the SDK should always return an
			// error which satisfies the awserr.Error interface.
			fmt.Println(err.Error())
		}
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.Prettify(resp))
}

func ExampleMachineLearning_GetBatchPrediction() {
	svc := machinelearning.New(nil)

	params := &machinelearning.GetBatchPredictionInput{
		BatchPredictionId: aws.String("EntityId"), // Required
	}
	resp, err := svc.GetBatchPrediction(params)

	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			// Generic AWS error with Code, Message, and original error (if any)
			fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
			if reqErr, ok := err.(awserr.RequestFailure); ok {
				// A service error occurred
				fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
			}
		} else {
			// This case should never be hit, the SDK should always return an
			// error which satisfies the awserr.Error interface.
			fmt.Println(err.Error())
		}
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.Prettify(resp))
}

func ExampleMachineLearning_GetDataSource() {
	svc := machinelearning.New(nil)

	params := &machinelearning.GetDataSourceInput{
		DataSourceId: aws.String("EntityId"), // Required
		Verbose:      aws.Bool(true),
	}
	resp, err := svc.GetDataSource(params)

	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			// Generic AWS error with Code, Message, and original error (if any)
			fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
			if reqErr, ok := err.(awserr.RequestFailure); ok {
				// A service error occurred
				fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
			}
		} else {
			// This case should never be hit, the SDK should always return an
			// error which satisfies the awserr.Error interface.
			fmt.Println(err.Error())
		}
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.Prettify(resp))
}

func ExampleMachineLearning_GetEvaluation() {
	svc := machinelearning.New(nil)

	params := &machinelearning.GetEvaluationInput{
		EvaluationId: aws.String("EntityId"), // Required
	}
	resp, err := svc.GetEvaluation(params)

	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			// Generic AWS error with Code, Message, and original error (if any)
			fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
			if reqErr, ok := err.(awserr.RequestFailure); ok {
				// A service error occurred
				fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
			}
		} else {
			// This case should never be hit, the SDK should always return an
			// error which satisfies the awserr.Error interface.
			fmt.Println(err.Error())
		}
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.Prettify(resp))
}

func ExampleMachineLearning_GetMLModel() {
	svc := machinelearning.New(nil)

	params := &machinelearning.GetMLModelInput{
		MLModelId: aws.String("EntityId"), // Required
		Verbose:   aws.Bool(true),
	}
	resp, err := svc.GetMLModel(params)

	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			// Generic AWS error with Code, Message, and original error (if any)
			fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
			if reqErr, ok := err.(awserr.RequestFailure); ok {
				// A service error occurred
				fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
			}
		} else {
			// This case should never be hit, the SDK should always return an
			// error which satisfies the awserr.Error interface.
			fmt.Println(err.Error())
		}
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.Prettify(resp))
}

func ExampleMachineLearning_Predict() {
	svc := machinelearning.New(nil)

	params := &machinelearning.PredictInput{
		MLModelId:       aws.String("EntityId"), // Required
		PredictEndpoint: aws.String("VipURL"),   // Required
		Record: map[string]*string{ // Required
			"Key": aws.String("VariableValue"), // Required
			// More values...
		},
	}
	resp, err := svc.Predict(params)

	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			// Generic AWS error with Code, Message, and original error (if any)
			fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
			if reqErr, ok := err.(awserr.RequestFailure); ok {
				// A service error occurred
				fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
			}
		} else {
			// This case should never be hit, the SDK should always return an
			// error which satisfies the awserr.Error interface.
			fmt.Println(err.Error())
		}
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.Prettify(resp))
}

func ExampleMachineLearning_UpdateBatchPrediction() {
	svc := machinelearning.New(nil)

	params := &machinelearning.UpdateBatchPredictionInput{
		BatchPredictionId:   aws.String("EntityId"),   // Required
		BatchPredictionName: aws.String("EntityName"), // Required
	}
	resp, err := svc.UpdateBatchPrediction(params)

	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			// Generic AWS error with Code, Message, and original error (if any)
			fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
			if reqErr, ok := err.(awserr.RequestFailure); ok {
				// A service error occurred
				fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
			}
		} else {
			// This case should never be hit, the SDK should always return an
			// error which satisfies the awserr.Error interface.
			fmt.Println(err.Error())
		}
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.Prettify(resp))
}

func ExampleMachineLearning_UpdateDataSource() {
	svc := machinelearning.New(nil)

	params := &machinelearning.UpdateDataSourceInput{
		DataSourceId:   aws.String("EntityId"),   // Required
		DataSourceName: aws.String("EntityName"), // Required
	}
	resp, err := svc.UpdateDataSource(params)

	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			// Generic AWS error with Code, Message, and original error (if any)
			fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
			if reqErr, ok := err.(awserr.RequestFailure); ok {
				// A service error occurred
				fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
			}
		} else {
			// This case should never be hit, the SDK should always return an
			// error which satisfies the awserr.Error interface.
			fmt.Println(err.Error())
		}
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.Prettify(resp))
}

func ExampleMachineLearning_UpdateEvaluation() {
	svc := machinelearning.New(nil)

	params := &machinelearning.UpdateEvaluationInput{
		EvaluationId:   aws.String("EntityId"),   // Required
		EvaluationName: aws.String("EntityName"), // Required
	}
	resp, err := svc.UpdateEvaluation(params)

	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			// Generic AWS error with Code, Message, and original error (if any)
			fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
			if reqErr, ok := err.(awserr.RequestFailure); ok {
				// A service error occurred
				fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
			}
		} else {
			// This case should never be hit, the SDK should always return an
			// error which satisfies the awserr.Error interface.
			fmt.Println(err.Error())
		}
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.Prettify(resp))
}

func ExampleMachineLearning_UpdateMLModel() {
	svc := machinelearning.New(nil)

	params := &machinelearning.UpdateMLModelInput{
		MLModelId:      aws.String("EntityId"), // Required
		MLModelName:    aws.String("EntityName"),
		ScoreThreshold: aws.Float64(1.0),
	}
	resp, err := svc.UpdateMLModel(params)

	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			// Generic AWS error with Code, Message, and original error (if any)
			fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
			if reqErr, ok := err.(awserr.RequestFailure); ok {
				// A service error occurred
				fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
			}
		} else {
			// This case should never be hit, the SDK should always return an
			// error which satisfies the awserr.Error interface.
			fmt.Println(err.Error())
		}
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.Prettify(resp))
}
