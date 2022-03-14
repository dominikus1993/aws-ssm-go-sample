package main

import (
	"fmt"
	"log"

	awsssm "github.com/PaddleHQ/go-aws-ssm"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/spf13/viper"
)

func printSSM(prefix string, ssmConfig *aws.Config, v *viper.Viper) error {
	pmstore, err := awsssm.NewParameterStore(ssmConfig)
	if err != nil {
		return err
	}
	params, err := pmstore.GetAllParametersByPath(fmt.Sprintf("/%s/", prefix), true)
	if err != nil {
		return err
	}
	v.SetConfigType(`json`)
	//params object implements the io.Reader interface that is required
	err = v.ReadConfig(params)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	v := viper.New()
	err := printSSM("gosamp", aws.NewConfig().WithRegion("eu-west-1").WithCredentials(credentials.NewEnvCredentials()), v)
	if err != nil {
		panic(err)
	}
	for _, key := range v.AllKeys() {
		log.Printf("%s = %s\n", key, v.GetString(key))
	}
}
