package libreoffice

import (
	"os"
	"fmt"
	"os/exec"
	"strconv"
	"math/rand"
	
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	storagePath="/dev/shm/"
)

func HandleConvertFile(
	ogFileType string, 
	newFileType string, 
	fileData []byte,
)([]byte){
	if len(ogFileType) == 0 || len(newFileType) == 0{
		status.Error(codes.NotFound, "Invlaid file types")
	}

	randBytes := strconv.Itoa(rand.Int())

	filePtr, _ := os.OpenFile(
		storagePath + randBytes, 
		os.O_WRONLY|os.O_CREATE, 
		0644,
	)

	bodyBytes := fileData

	_, fileWriteError := filePtr.Write(bodyBytes)
	if fileWriteError != nil {
		status.Error(codes.FailedPrecondition, fmt.Sprint(fileWriteError))
	}
	filePtr.Close()

	libreofficeCmdError := runLibreoffice(ogFileType, newFileType, randBytes)
	if libreofficeCmdError != nil {
		status.Error(codes.Unknown, fmt.Sprint(libreofficeCmdError))
	}

	data, readFileError := os.ReadFile(storagePath + randBytes + "." + newFileType)
	if readFileError != nil {
		status.Error(codes.FailedPrecondition, fmt.Sprint(readFileError))
	}

	cleanDevShm(randBytes, newFileType)
	return data
}

func runLibreoffice(
	ogFileType string,
	newFileType string, 
	randBytes string,
)error{
	var libreofficeOptions string
	if ogFileType == "pdf" {
		libreofficeOptions = "--infilter='writer_pdf_import'"
	}
	libreofficeCmd := exec.Command(
		"bash",
		"-c",
		"libreoffice " + libreofficeOptions +
		" -env:UserInstallation=file://" + storagePath + randBytes + "_lo " + 
		"--headless --convert-to " + newFileType + " " + randBytes,
	)

	libreofficeCmd.Dir="/dev/shm"
	return libreofficeCmd.Run()
}

func cleanDevShm(fileName string, newFileType string){
	os.Remove(storagePath + fileName)
	os.Remove(storagePath + fileName + "." + newFileType)
	os.RemoveAll(storagePath + fileName + "_lo")
}
