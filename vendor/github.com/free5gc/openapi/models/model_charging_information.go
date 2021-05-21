/*
 * Npcf_SMPolicyControl
 *
 * Session Management Policy Control Service
 *
 * API version: 1.0.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package models

type ChargingInformation struct {
	PrimaryChfAddress   string `json:"primaryChfAddress" yaml:"primaryChfAddress" bson:"primaryChfAddress" mapstructure:"PrimaryChfAddress"`
	SecondaryChfAddress string `json:"secondaryChfAddress" yaml:"secondaryChfAddress" bson:"secondaryChfAddress" mapstructure:"SecondaryChfAddress"`
}