/*
Copyright ApeCloud, Inc.
Licensed under the Apache v2(found in the LICENSE file in the root directory).
*/
package cdc

import (
	"flag"
	"fmt"
	"os"
)

type CdcConfig struct {
	TableSchema     string
	SourceTableName string
	TargetTableName string
	FilterStatement string
	WeScaleHost     string
	WeScaleGrpcPort string
}

var DefaultConfig CdcConfig

func init() {
	RegisterStringVar(&DefaultConfig.TableSchema, "TABLE_SCHEMA", "", "The table schema.")
	RegisterStringVar(&DefaultConfig.SourceTableName, "SOURCE_TABLE_NAME", "", "The source table name.")
	RegisterStringVar(&DefaultConfig.TargetTableName, "TARGET_TABLE_NAME", "", "The target table name.")
	RegisterStringVar(&DefaultConfig.FilterStatement, "FILTER_STATEMENT", "", "The filter statement.")
	RegisterStringVar(&DefaultConfig.WeScaleHost, "WESCALE_HOST", "127.0.0.1", "The WeScale host.")
	RegisterStringVar(&DefaultConfig.WeScaleGrpcPort, "WESCALE_GRPC_PORT", "15991", "The WeScale GRPC port.")
	logDir := os.Getenv("LOG_DIR")
	if logDir != "" {
		flag.Set("log_dir", logDir)
	}
}

func checkFlags() error {
	if DefaultConfig.TableSchema == "" {
		return fmt.Errorf("table-schema is required")
	}
	if DefaultConfig.SourceTableName == "" {
		return fmt.Errorf("table-name is required")
	}
	if DefaultConfig.FilterStatement == "" {
		return fmt.Errorf("filter-statement is required")
	}
	if DefaultConfig.WeScaleHost == "" {
		return fmt.Errorf("wescale-host is required")
	}
	if DefaultConfig.WeScaleGrpcPort == "" {
		return fmt.Errorf("wescale-grpc-port is required")
	}
	return nil
}
