// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package fsx_test

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/awserr"
	"github.com/aws/aws-sdk-go-v2/aws/external"
	"github.com/aws/aws-sdk-go-v2/service/fsx"
)

var _ time.Duration
var _ strings.Reader
var _ aws.Config

func parseTime(layout, value string) *time.Time {
	t, err := time.Parse(layout, value)
	if err != nil {
		panic(err)
	}
	return &t
}

// To create a new backup
//
// This operation creates a new backup.
func ExampleFSx_CreateBackupRequest_shared00() {
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		panic("failed to load config, " + err.Error())
	}

	svc := fsx.New(cfg)
	input := &fsx.CreateBackupInput{
		FileSystemId: aws.String("fs-0498eed5fe91001ec"),
		Tags: []fsx.Tag{
			{
				Key:   aws.String("Name"),
				Value: aws.String("MyBackup"),
			},
		},
	}

	req := svc.CreateBackupRequest(input)
	result, err := req.Send(context.Background())
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case fsx.ErrCodeBadRequest:
				fmt.Println(fsx.ErrCodeBadRequest, aerr.Error())
			case fsx.ErrCodeFileSystemNotFound:
				fmt.Println(fsx.ErrCodeFileSystemNotFound, aerr.Error())
			case fsx.ErrCodeBackupInProgress:
				fmt.Println(fsx.ErrCodeBackupInProgress, aerr.Error())
			case fsx.ErrCodeIncompatibleParameterError:
				fmt.Println(fsx.ErrCodeIncompatibleParameterError, aerr.Error())
			case fsx.ErrCodeServiceLimitExceeded:
				fmt.Println(fsx.ErrCodeServiceLimitExceeded, aerr.Error())
			case fsx.ErrCodeInternalServerError:
				fmt.Println(fsx.ErrCodeInternalServerError, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// To create a new file system
//
// This operation creates a new file system.
func ExampleFSx_CreateFileSystemRequest_shared00() {
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		panic("failed to load config, " + err.Error())
	}

	svc := fsx.New(cfg)
	input := &fsx.CreateFileSystemInput{
		ClientRequestToken: aws.String("a8ca07e4-61ec-4399-99f4-19853801bcd5"),
		FileSystemType:     fsx.FileSystemTypeWindows,
		KmsKeyId:           aws.String("arn:aws:kms:us-east-1:012345678912:key/0ff3ea8d-130e-4133-877f-93908b6fdbd6"),
		SecurityGroupIds: []string{
			"sg-edcd9784",
		},
		StorageCapacity: aws.Int64(300),
		SubnetIds: []string{
			"subnet-1234abcd",
		},
		Tags: []fsx.Tag{
			{
				Key:   aws.String("Name"),
				Value: aws.String("MyFileSystem"),
			},
		},
		WindowsConfiguration: &fsx.CreateFileSystemWindowsConfiguration{
			ActiveDirectoryId:             aws.String("d-1234abcd12"),
			AutomaticBackupRetentionDays:  aws.Int64(30),
			DailyAutomaticBackupStartTime: aws.String("05:00"),
			ThroughputCapacity:            aws.Int64(8),
			WeeklyMaintenanceStartTime:    aws.String("1:05:00"),
		},
	}

	req := svc.CreateFileSystemRequest(input)
	result, err := req.Send(context.Background())
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case fsx.ErrCodeBadRequest:
				fmt.Println(fsx.ErrCodeBadRequest, aerr.Error())
			case fsx.ErrCodeActiveDirectoryError:
				fmt.Println(fsx.ErrCodeActiveDirectoryError, aerr.Error())
			case fsx.ErrCodeIncompatibleParameterError:
				fmt.Println(fsx.ErrCodeIncompatibleParameterError, aerr.Error())
			case fsx.ErrCodeInvalidImportPath:
				fmt.Println(fsx.ErrCodeInvalidImportPath, aerr.Error())
			case fsx.ErrCodeInvalidExportPath:
				fmt.Println(fsx.ErrCodeInvalidExportPath, aerr.Error())
			case fsx.ErrCodeInvalidNetworkSettings:
				fmt.Println(fsx.ErrCodeInvalidNetworkSettings, aerr.Error())
			case fsx.ErrCodeServiceLimitExceeded:
				fmt.Println(fsx.ErrCodeServiceLimitExceeded, aerr.Error())
			case fsx.ErrCodeInternalServerError:
				fmt.Println(fsx.ErrCodeInternalServerError, aerr.Error())
			case fsx.ErrCodeMissingFileSystemConfiguration:
				fmt.Println(fsx.ErrCodeMissingFileSystemConfiguration, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// To create a new file system from backup
//
// This operation creates a new file system from backup.
func ExampleFSx_CreateFileSystemFromBackupRequest_shared00() {
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		panic("failed to load config, " + err.Error())
	}

	svc := fsx.New(cfg)
	input := &fsx.CreateFileSystemFromBackupInput{
		BackupId:           aws.String("backup-03e3c82e0183b7b6b"),
		ClientRequestToken: aws.String("f4c94ed7-238d-4c46-93db-48cd62ec33b7"),
		SecurityGroupIds: []string{
			"sg-edcd9784",
		},
		SubnetIds: []string{
			"subnet-1234abcd",
		},
		Tags: []fsx.Tag{
			{
				Key:   aws.String("Name"),
				Value: aws.String("MyFileSystem"),
			},
		},
		WindowsConfiguration: &fsx.CreateFileSystemWindowsConfiguration{
			ThroughputCapacity: aws.Int64(8),
		},
	}

	req := svc.CreateFileSystemFromBackupRequest(input)
	result, err := req.Send(context.Background())
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case fsx.ErrCodeBadRequest:
				fmt.Println(fsx.ErrCodeBadRequest, aerr.Error())
			case fsx.ErrCodeActiveDirectoryError:
				fmt.Println(fsx.ErrCodeActiveDirectoryError, aerr.Error())
			case fsx.ErrCodeIncompatibleParameterError:
				fmt.Println(fsx.ErrCodeIncompatibleParameterError, aerr.Error())
			case fsx.ErrCodeInvalidNetworkSettings:
				fmt.Println(fsx.ErrCodeInvalidNetworkSettings, aerr.Error())
			case fsx.ErrCodeServiceLimitExceeded:
				fmt.Println(fsx.ErrCodeServiceLimitExceeded, aerr.Error())
			case fsx.ErrCodeBackupNotFound:
				fmt.Println(fsx.ErrCodeBackupNotFound, aerr.Error())
			case fsx.ErrCodeInternalServerError:
				fmt.Println(fsx.ErrCodeInternalServerError, aerr.Error())
			case fsx.ErrCodeMissingFileSystemConfiguration:
				fmt.Println(fsx.ErrCodeMissingFileSystemConfiguration, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// To delete a backup
//
// This operation deletes an Amazon FSx file system backup.
func ExampleFSx_DeleteBackupRequest_shared00() {
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		panic("failed to load config, " + err.Error())
	}

	svc := fsx.New(cfg)
	input := &fsx.DeleteBackupInput{
		BackupId: aws.String("backup-03e3c82e0183b7b6b"),
	}

	req := svc.DeleteBackupRequest(input)
	result, err := req.Send(context.Background())
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case fsx.ErrCodeBadRequest:
				fmt.Println(fsx.ErrCodeBadRequest, aerr.Error())
			case fsx.ErrCodeBackupInProgress:
				fmt.Println(fsx.ErrCodeBackupInProgress, aerr.Error())
			case fsx.ErrCodeBackupNotFound:
				fmt.Println(fsx.ErrCodeBackupNotFound, aerr.Error())
			case fsx.ErrCodeBackupRestoring:
				fmt.Println(fsx.ErrCodeBackupRestoring, aerr.Error())
			case fsx.ErrCodeIncompatibleParameterError:
				fmt.Println(fsx.ErrCodeIncompatibleParameterError, aerr.Error())
			case fsx.ErrCodeInternalServerError:
				fmt.Println(fsx.ErrCodeInternalServerError, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// To delete a file system
//
// This operation deletes an Amazon FSx file system.
func ExampleFSx_DeleteFileSystemRequest_shared00() {
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		panic("failed to load config, " + err.Error())
	}

	svc := fsx.New(cfg)
	input := &fsx.DeleteFileSystemInput{
		FileSystemId: aws.String("fs-0498eed5fe91001ec"),
	}

	req := svc.DeleteFileSystemRequest(input)
	result, err := req.Send(context.Background())
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case fsx.ErrCodeBadRequest:
				fmt.Println(fsx.ErrCodeBadRequest, aerr.Error())
			case fsx.ErrCodeIncompatibleParameterError:
				fmt.Println(fsx.ErrCodeIncompatibleParameterError, aerr.Error())
			case fsx.ErrCodeFileSystemNotFound:
				fmt.Println(fsx.ErrCodeFileSystemNotFound, aerr.Error())
			case fsx.ErrCodeServiceLimitExceeded:
				fmt.Println(fsx.ErrCodeServiceLimitExceeded, aerr.Error())
			case fsx.ErrCodeInternalServerError:
				fmt.Println(fsx.ErrCodeInternalServerError, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// To describe Amazon FSx backups
//
// This operation describes all of the Amazon FSx backups in an account.
func ExampleFSx_DescribeBackupsRequest_shared00() {
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		panic("failed to load config, " + err.Error())
	}

	svc := fsx.New(cfg)
	input := &fsx.DescribeBackupsInput{}

	req := svc.DescribeBackupsRequest(input)
	result, err := req.Send(context.Background())
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case fsx.ErrCodeBadRequest:
				fmt.Println(fsx.ErrCodeBadRequest, aerr.Error())
			case fsx.ErrCodeFileSystemNotFound:
				fmt.Println(fsx.ErrCodeFileSystemNotFound, aerr.Error())
			case fsx.ErrCodeBackupNotFound:
				fmt.Println(fsx.ErrCodeBackupNotFound, aerr.Error())
			case fsx.ErrCodeInternalServerError:
				fmt.Println(fsx.ErrCodeInternalServerError, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// To describe an Amazon FSx file system
//
// This operation describes all of the Amazon FSx file systems in an account.
func ExampleFSx_DescribeFileSystemsRequest_shared00() {
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		panic("failed to load config, " + err.Error())
	}

	svc := fsx.New(cfg)
	input := &fsx.DescribeFileSystemsInput{}

	req := svc.DescribeFileSystemsRequest(input)
	result, err := req.Send(context.Background())
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case fsx.ErrCodeBadRequest:
				fmt.Println(fsx.ErrCodeBadRequest, aerr.Error())
			case fsx.ErrCodeFileSystemNotFound:
				fmt.Println(fsx.ErrCodeFileSystemNotFound, aerr.Error())
			case fsx.ErrCodeInternalServerError:
				fmt.Println(fsx.ErrCodeInternalServerError, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// To list tags for a resource
//
// This operation lists tags for an Amazon FSx resource.
func ExampleFSx_ListTagsForResourceRequest_shared00() {
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		panic("failed to load config, " + err.Error())
	}

	svc := fsx.New(cfg)
	input := &fsx.ListTagsForResourceInput{
		ResourceARN: aws.String("arn:aws:fsx:us-east-1:012345678912:file-system/fs-0498eed5fe91001ec"),
	}

	req := svc.ListTagsForResourceRequest(input)
	result, err := req.Send(context.Background())
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case fsx.ErrCodeBadRequest:
				fmt.Println(fsx.ErrCodeBadRequest, aerr.Error())
			case fsx.ErrCodeInternalServerError:
				fmt.Println(fsx.ErrCodeInternalServerError, aerr.Error())
			case fsx.ErrCodeResourceNotFound:
				fmt.Println(fsx.ErrCodeResourceNotFound, aerr.Error())
			case fsx.ErrCodeNotServiceResourceError:
				fmt.Println(fsx.ErrCodeNotServiceResourceError, aerr.Error())
			case fsx.ErrCodeResourceDoesNotSupportTagging:
				fmt.Println(fsx.ErrCodeResourceDoesNotSupportTagging, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// To tag a resource
//
// This operation tags an Amazon FSx resource.
func ExampleFSx_TagResourceRequest_shared00() {
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		panic("failed to load config, " + err.Error())
	}

	svc := fsx.New(cfg)
	input := &fsx.TagResourceInput{
		ResourceARN: aws.String("arn:aws:fsx:us-east-1:012345678912:file-system/fs-0498eed5fe91001ec"),
		Tags: []fsx.Tag{
			{
				Key:   aws.String("Name"),
				Value: aws.String("MyFileSystem"),
			},
		},
	}

	req := svc.TagResourceRequest(input)
	result, err := req.Send(context.Background())
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case fsx.ErrCodeBadRequest:
				fmt.Println(fsx.ErrCodeBadRequest, aerr.Error())
			case fsx.ErrCodeInternalServerError:
				fmt.Println(fsx.ErrCodeInternalServerError, aerr.Error())
			case fsx.ErrCodeResourceNotFound:
				fmt.Println(fsx.ErrCodeResourceNotFound, aerr.Error())
			case fsx.ErrCodeNotServiceResourceError:
				fmt.Println(fsx.ErrCodeNotServiceResourceError, aerr.Error())
			case fsx.ErrCodeResourceDoesNotSupportTagging:
				fmt.Println(fsx.ErrCodeResourceDoesNotSupportTagging, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// To untag a resource
//
// This operation untags an Amazon FSx resource.
func ExampleFSx_UntagResourceRequest_shared00() {
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		panic("failed to load config, " + err.Error())
	}

	svc := fsx.New(cfg)
	input := &fsx.UntagResourceInput{
		ResourceARN: aws.String("arn:aws:fsx:us-east-1:012345678912:file-system/fs-0498eed5fe91001ec"),
		TagKeys: []string{
			"Name",
		},
	}

	req := svc.UntagResourceRequest(input)
	result, err := req.Send(context.Background())
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case fsx.ErrCodeBadRequest:
				fmt.Println(fsx.ErrCodeBadRequest, aerr.Error())
			case fsx.ErrCodeInternalServerError:
				fmt.Println(fsx.ErrCodeInternalServerError, aerr.Error())
			case fsx.ErrCodeResourceNotFound:
				fmt.Println(fsx.ErrCodeResourceNotFound, aerr.Error())
			case fsx.ErrCodeNotServiceResourceError:
				fmt.Println(fsx.ErrCodeNotServiceResourceError, aerr.Error())
			case fsx.ErrCodeResourceDoesNotSupportTagging:
				fmt.Println(fsx.ErrCodeResourceDoesNotSupportTagging, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}

// To update an existing file system
//
// This operation updates an existing file system.
func ExampleFSx_UpdateFileSystemRequest_shared00() {
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		panic("failed to load config, " + err.Error())
	}

	svc := fsx.New(cfg)
	input := &fsx.UpdateFileSystemInput{
		FileSystemId: aws.String("fs-0498eed5fe91001ec"),
		WindowsConfiguration: &fsx.UpdateFileSystemWindowsConfiguration{
			AutomaticBackupRetentionDays:  aws.Int64(10),
			DailyAutomaticBackupStartTime: aws.String("06:00"),
			WeeklyMaintenanceStartTime:    aws.String("3:06:00"),
		},
	}

	req := svc.UpdateFileSystemRequest(input)
	result, err := req.Send(context.Background())
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			case fsx.ErrCodeBadRequest:
				fmt.Println(fsx.ErrCodeBadRequest, aerr.Error())
			case fsx.ErrCodeIncompatibleParameterError:
				fmt.Println(fsx.ErrCodeIncompatibleParameterError, aerr.Error())
			case fsx.ErrCodeInternalServerError:
				fmt.Println(fsx.ErrCodeInternalServerError, aerr.Error())
			case fsx.ErrCodeFileSystemNotFound:
				fmt.Println(fsx.ErrCodeFileSystemNotFound, aerr.Error())
			case fsx.ErrCodeMissingFileSystemConfiguration:
				fmt.Println(fsx.ErrCodeMissingFileSystemConfiguration, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	fmt.Println(result)
}