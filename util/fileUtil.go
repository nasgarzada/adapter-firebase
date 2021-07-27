package util

import (
	"encoding/json"
	"github.com/nasgarzada/adapter-firebase/config"
	"github.com/nasgarzada/adapter-firebase/model"
	"github.com/sirupsen/logrus"
	"io/ioutil"
)

const (
	prefix = ""
	indent = " "
)

func CreateFirebaseConfigFile() error {
	logrus.Info("ActionLog.CreateFirebaseConfigFile.start")
	firebaseConfig := &model.FirebaseConfig{
		Type:                    config.Props.Type,
		ProjectId:               config.Props.ProjectId,
		PrivateKeyId:            config.Props.PrivateKeyId,
		PrivateKey:              config.Props.PrivateKey,
		ClientEmail:             config.Props.ClientEmail,
		ClientId:                config.Props.ClientId,
		AuthUri:                 config.Props.AuthUri,
		TokenUri:                config.Props.TokenUri,
		AuthProviderX509CertUrl: config.Props.AuthProviderX509CertUrl,
		ClientX509CertUrl:       config.Props.ClientX509CertUrl,
	}
	file, err := json.MarshalIndent(firebaseConfig, prefix, indent)
	if err != nil {
		logrus.Error("ActionLog.CreateFirebaseConfigFile.error - ", err)
		return err
	}
	err = ioutil.WriteFile(config.Props.FirebaseConfigPath, file, 0644)
	logrus.Info("ActionLog.CreateFirebaseConfigFile.start")
	return err
}
