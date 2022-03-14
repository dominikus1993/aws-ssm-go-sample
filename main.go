package main

import (
	"fmt"

	awsssm "github.com/PaddleHQ/go-aws-ssm"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/spf13/viper"
)

func printSSM(prefix string, ssmConfig *aws.Config) (*viper.Viper, error) {
	pmstore, err := awsssm.NewParameterStore(ssmConfig)
	if err != nil {
		return nil, err
	}
	params, err := pmstore.GetAllParametersByPath(fmt.Sprintf("/%s/", prefix), true)
	if err != nil {
		return nil, err
	}
	v := viper.New()
	v.SetConfigType(`json`)
	//params object implements the io.Reader interface that is required
	err = v.ReadConfig(params)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func main() {
	v, err := printSSM("gosamp", aws.NewConfig().WithRegion("eu-west-1").WithCredentials(credentials.NewEnvCredentials()))
	if err != nil {
		panic(err)
	}
	for _, key := range v.AllKeys() {
		fmt.Printf("%s = %s\n", key, v.GetString(key))
	}
}
