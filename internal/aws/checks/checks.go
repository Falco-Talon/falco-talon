package checks

import (
	"context"

	lambdaActionner "github.com/Falco-Talon/falco-talon/actionners/aws/lambda"
	"github.com/Falco-Talon/falco-talon/internal/aws/client"
	"github.com/Falco-Talon/falco-talon/internal/events"
	"github.com/Falco-Talon/falco-talon/internal/rules"
	"github.com/Falco-Talon/falco-talon/utils"

	"github.com/aws/aws-sdk-go-v2/service/lambda"
)

func CheckLambdaExist(_ *events.Event, action *rules.Action) error {
	lambdaClient := client.GetAWSClient().GetLambdaClient()
	parameters := action.GetParameters()

	var lambdaConfig lambdaActionner.Config
	err := utils.DecodeParams(parameters, &lambdaConfig)
	if err != nil {
		return err
	}
	_, err = lambdaClient.GetFunction(context.Background(), &lambda.GetFunctionInput{
		FunctionName: &lambdaConfig.AWSLambdaName,
	})
	if err != nil {
		return err
	}
	return nil
}
