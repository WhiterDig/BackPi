package repository

import (
	"BackPi/model"
	"context"
	"encoding/json"
	"fmt"
	"github.com/juju/loggo"
	"github.com/pkg/errors"
	"io/ioutil"
	"os"
)

type Credential interface {
	WriteToFile(newCred model.Credentials) (*string, error)
}

type credential struct {
	ctx context.Context
	log loggo.Logger
}

func NewCredentialRepo(ctx context.Context, log loggo.Logger) *credential {
	return &credential{
		ctx: ctx,
		log: log,
	}
}

//WriteToFile is a repository method that writes a newCred to json file "creds.json"
func (cd *credential) WriteToFile(newCred model.Credentials) (*string, error) {
	cd.log.Infof("Repository Layer - Checking if creds.json has already been made")
	err := createFileIfNotExist(cd.log)
	if err != nil {
		return nil, err
	}
	cd.log.Infof("Repository Layer - creds.json already exists or has been successfully created")
	cd.log.Infof("Repository Layer - Now Opening creds.json to begin update")
	//Open the json file so that we may edit it
	file, err := os.Open("./assets/json/creds.json")
	if err != nil {
		return nil, errors.Wrapf(err, "error opening file")
	}
	//Gotta make sure to close our opened files!
	defer file.Close()
	//Reads our file and retrieves a byte array.
	byteValue, _ := ioutil.ReadAll(file)
	var creds []model.Credentials
	//Since the json file will hold an array of model.Credentials, it understands how to unmarshal the json into the right properties of an array of model.Credentials
	json.Unmarshal(byteValue, &creds)
	//At this point, it's just an array.  We can append a new value to it.  In this case, the newCred from the frontend.
	creds = append(creds, newCred)
	//Now we just marshal the array with the appended newCred and get another byte array.
	b, err := json.Marshal(creds)
	if err != nil {
		return nil, errors.Wrapf(err, "error marshaling json")
	}
	//Write the byte array back to the file, and set the file permission to 0755.
	err = ioutil.WriteFile("./assets/json/creds.json", b, 0755)
	if err != nil {
		return nil, errors.Wrapf(err, "error writing file")
	}
	cd.log.Infof("Repository Layer - Update to creds.json has completed successfully")
	resp := fmt.Sprint("Saved Cred Object: ", newCred.String())
	return &resp, nil
}

func createFileIfNotExist(log loggo.Logger) error {
	_, err := os.Stat("./assets/json/creds.json")
	//Create file if not exists
	if os.IsNotExist(err) {
		log.Infof("Repository Layer - create creds.json - creds.json doesn't exist, attempting to create")
		file, err := os.Create("./assets/json/creds.json")
		if err != nil {
			defer file.Close()
			return errors.Wrapf(err, "create file failed")
		}
		log.Infof("Repository Layer - create creds.json - creds.json now exists")
	}
	return nil
}
