// Code generated by github.com/Khan/genqlient, DO NOT EDIT.

package graphql

import (
	"context"

	"github.com/Khan/genqlient/graphql"
)

// Search after specified block number and alertId
type AlertEndCursorInput struct {
	BlockNumber uint   `json:"blockNumber"`
	AlertId     string `json:"alertId"`
}

// GetBlockNumber returns AlertEndCursorInput.BlockNumber, and is useful for accessing the field via an interface.
func (v *AlertEndCursorInput) GetBlockNumber() uint { return v.BlockNumber }

// GetAlertId returns AlertEndCursorInput.AlertId, and is useful for accessing the field via an interface.
func (v *AlertEndCursorInput) GetAlertId() string { return v.AlertId }

// Alert list input
type AlertsInput struct {
	// Indicate a list of addresses.
	// Alerts returned will have those addresses involved.
	Addresses []string `json:"addresses"`
	// Search results after the specified cursor
	After *AlertEndCursorInput `json:"after,omitempty"`
	// Filter alerts by alert-id.
	AlertId string `json:"alertId"`
	// Get a specific alert by alert hash.
	AlertHash string `json:"alertHash"`
	// Filter alerts by alert name.
	AlertName string `json:"alertName"`
	// Block Date range
	// Alerts returned will be between the specified start and end block timestamp dates when the threats were detected.
	// By default, start and end date will be set to local query execution date.
	BlockDateRange *DateRange `json:"blockDateRange,omitempty"`
	// Block number range
	// Alerts for the block number range will be returned.
	BlockNumberRange *BlockRange `json:"blockNumberRange,omitempty"`
	// Block Timestamp range
	// Alerts returned will be between the specified start and end block timestamp when the threats were detected.
	BlockTimestampRange *TimestampRange `json:"blockTimestampRange,omitempty"`
	// Indicate sorting order by block number,
	// 'desc' or 'asc'.
	// Default is 'desc'.
	BlockSortDirection Sort `json:"blockSortDirection"`
	// Indicate a list of bot hashes.
	// Alerts returned will only be from any of those bots.
	Bots []string `json:"bots"`
	// Indicate a chain Id: EIP155 identifier of the chain
	// Alerts returned will only be from the specific chain Id
	// Default is 1 = Ethereum Mainnet.
	ChainId uint `json:"chainId"`
	// Indicate number of milliseconds
	// Alerts returned will be alerts created since the number of milliseconds indicated ago.
	CreatedSince uint `json:"createdSince"`
	// Indicate number of milliseconds
	// Alerts returned will be alerts created before the number of milliseconds indicated ago.
	CreatedBefore uint `json:"createdBefore"`
	// Indicate max number of results.
	First uint `json:"first"`
	// Filter alerts by number of scan nodes confirming the alert.
	ScanNodeConfirmations *ScanNodeFilters `json:"scanNodeConfirmations,omitempty"`
	// Filter alerts by severity levels.
	Severities []string `json:"severities"`
	// Indicate a transaction hash
	// Alerts returned will only be from that transaction.
	TransactionHash string `json:"transactionHash"`
	// Filter alerts by multiple alert-ids.
	AlertIds []string `json:"alertIds"`
}

// GetAddresses returns AlertsInput.Addresses, and is useful for accessing the field via an interface.
func (v *AlertsInput) GetAddresses() []string { return v.Addresses }

// GetAfter returns AlertsInput.After, and is useful for accessing the field via an interface.
func (v *AlertsInput) GetAfter() *AlertEndCursorInput { return v.After }

// GetAlertId returns AlertsInput.AlertId, and is useful for accessing the field via an interface.
func (v *AlertsInput) GetAlertId() string { return v.AlertId }

// GetAlertHash returns AlertsInput.AlertHash, and is useful for accessing the field via an interface.
func (v *AlertsInput) GetAlertHash() string { return v.AlertHash }

// GetAlertName returns AlertsInput.AlertName, and is useful for accessing the field via an interface.
func (v *AlertsInput) GetAlertName() string { return v.AlertName }

// GetBlockDateRange returns AlertsInput.BlockDateRange, and is useful for accessing the field via an interface.
func (v *AlertsInput) GetBlockDateRange() *DateRange { return v.BlockDateRange }

// GetBlockNumberRange returns AlertsInput.BlockNumberRange, and is useful for accessing the field via an interface.
func (v *AlertsInput) GetBlockNumberRange() *BlockRange { return v.BlockNumberRange }

// GetBlockTimestampRange returns AlertsInput.BlockTimestampRange, and is useful for accessing the field via an interface.
func (v *AlertsInput) GetBlockTimestampRange() *TimestampRange { return v.BlockTimestampRange }

// GetBlockSortDirection returns AlertsInput.BlockSortDirection, and is useful for accessing the field via an interface.
func (v *AlertsInput) GetBlockSortDirection() Sort { return v.BlockSortDirection }

// GetBots returns AlertsInput.Bots, and is useful for accessing the field via an interface.
func (v *AlertsInput) GetBots() []string { return v.Bots }

// GetChainId returns AlertsInput.ChainId, and is useful for accessing the field via an interface.
func (v *AlertsInput) GetChainId() uint { return v.ChainId }

// GetCreatedSince returns AlertsInput.CreatedSince, and is useful for accessing the field via an interface.
func (v *AlertsInput) GetCreatedSince() uint { return v.CreatedSince }

// GetCreatedBefore returns AlertsInput.CreatedBefore, and is useful for accessing the field via an interface.
func (v *AlertsInput) GetCreatedBefore() uint { return v.CreatedBefore }

// GetFirst returns AlertsInput.First, and is useful for accessing the field via an interface.
func (v *AlertsInput) GetFirst() uint { return v.First }

// GetScanNodeConfirmations returns AlertsInput.ScanNodeConfirmations, and is useful for accessing the field via an interface.
func (v *AlertsInput) GetScanNodeConfirmations() *ScanNodeFilters { return v.ScanNodeConfirmations }

// GetSeverities returns AlertsInput.Severities, and is useful for accessing the field via an interface.
func (v *AlertsInput) GetSeverities() []string { return v.Severities }

// GetTransactionHash returns AlertsInput.TransactionHash, and is useful for accessing the field via an interface.
func (v *AlertsInput) GetTransactionHash() string { return v.TransactionHash }

// GetAlertIds returns AlertsInput.AlertIds, and is useful for accessing the field via an interface.
func (v *AlertsInput) GetAlertIds() []string { return v.AlertIds }

// Block range
type BlockRange struct {
	StartBlockNumber uint `json:"startBlockNumber"`
	EndBlockNumber   uint `json:"endBlockNumber"`
}

// GetStartBlockNumber returns BlockRange.StartBlockNumber, and is useful for accessing the field via an interface.
func (v *BlockRange) GetStartBlockNumber() uint { return v.StartBlockNumber }

// GetEndBlockNumber returns BlockRange.EndBlockNumber, and is useful for accessing the field via an interface.
func (v *BlockRange) GetEndBlockNumber() uint { return v.EndBlockNumber }

// Date range
// Date format: YYYY-MM-DD
type DateRange struct {
	StartDate string `json:"startDate"`
	EndDate   string `json:"endDate"`
}

// GetStartDate returns DateRange.StartDate, and is useful for accessing the field via an interface.
func (v *DateRange) GetStartDate() string { return v.StartDate }

// GetEndDate returns DateRange.EndDate, and is useful for accessing the field via an interface.
func (v *DateRange) GetEndDate() string { return v.EndDate }

// Filter by number of scan nodes confirming the alert.
type ScanNodeFilters struct {
	Gte uint `json:"gte"`
	Lte uint `json:"lte"`
}

// GetGte returns ScanNodeFilters.Gte, and is useful for accessing the field via an interface.
func (v *ScanNodeFilters) GetGte() uint { return v.Gte }

// GetLte returns ScanNodeFilters.Lte, and is useful for accessing the field via an interface.
func (v *ScanNodeFilters) GetLte() uint { return v.Lte }

type Sort string

const (
	SortAsc  Sort = "asc"
	SortDesc Sort = "desc"
)

// Timestamp range
// Timestamp format: number of milliseconds from start of UNIX epoch
type TimestampRange struct {
	StartTimestamp uint `json:"startTimestamp"`
	EndTimestamp   uint `json:"endTimestamp"`
}

// GetStartTimestamp returns TimestampRange.StartTimestamp, and is useful for accessing the field via an interface.
func (v *TimestampRange) GetStartTimestamp() uint { return v.StartTimestamp }

// GetEndTimestamp returns TimestampRange.EndTimestamp, and is useful for accessing the field via an interface.
func (v *TimestampRange) GetEndTimestamp() uint { return v.EndTimestamp }

// __getAlertsInput is used internally by genqlient
type __getAlertsInput struct {
	Input *AlertsInput `json:"input,omitempty"`
}

// GetInput returns __getAlertsInput.Input, and is useful for accessing the field via an interface.
func (v *__getAlertsInput) GetInput() *AlertsInput { return v.Input }

// getAlertsAlertsAlertsResponse includes the requested fields of the GraphQL type AlertsResponse.
type getAlertsAlertsAlertsResponse struct {
	PageInfo *getAlertsAlertsAlertsResponsePageInfoAlertPageInfo `json:"pageInfo"`
	Alerts   []*getAlertsAlertsAlertsResponseAlertsAlert         `json:"alerts"`
}

// GetPageInfo returns getAlertsAlertsAlertsResponse.PageInfo, and is useful for accessing the field via an interface.
func (v *getAlertsAlertsAlertsResponse) GetPageInfo() *getAlertsAlertsAlertsResponsePageInfoAlertPageInfo {
	return v.PageInfo
}

// GetAlerts returns getAlertsAlertsAlertsResponse.Alerts, and is useful for accessing the field via an interface.
func (v *getAlertsAlertsAlertsResponse) GetAlerts() []*getAlertsAlertsAlertsResponseAlertsAlert {
	return v.Alerts
}

// getAlertsAlertsAlertsResponseAlertsAlert includes the requested fields of the GraphQL type Alert.
// The GraphQL type's documentation follows.
//
// Alert information
type getAlertsAlertsAlertsResponseAlertsAlert struct {
	// Unique string to identify this class of finding,
	// primarily used to group similar findings for the end user
	AlertId string `json:"alertId"`
	// List of addresses involved in the alert
	Addresses []string `json:"addresses"`
	// List of contracts related to the alert
	Contracts []*getAlertsAlertsAlertsResponseAlertsAlertContractsContract `json:"contracts"`
	// Timestamp when the alert was published
	CreatedAt string `json:"createdAt"`
	// Text description of the alert
	Description string `json:"description"`
	// Alert hash identifier
	Hash string `json:"hash"`
	// Extra alert information
	Metadata map[string]string `json:"metadata"`
	// Alert name
	Name string `json:"name"`
	// List of Web3 projects related to the alert
	Projects []*getAlertsAlertsAlertsResponseAlertsAlertProjectsProject `json:"projects"`
	// Name of protocol being reported on
	Protocol string `json:"protocol"`
	// Number of scanners that found the alert
	ScanNodeCount int `json:"scanNodeCount"`
	// Impact level of finding
	//
	// CRITICAL - Exploitable vulnerabilities, massive impact on users/funds
	//
	// HIGH - Exploitable under more specific conditions, significant impact on users/funds
	//
	// MEDIUM - Notable unexpected behaviours, moderate to low impact on users/funds
	//
	// LOW - Minor oversights, negligible impact on users/funds
	//
	// INFO - Miscellaneous behaviours worth describing
	Severity string `json:"severity"`
	// Source where the alert was detected
	Source *getAlertsAlertsAlertsResponseAlertsAlertSource `json:"source"`
	// Type of alert
	// UNKNOWN
	// PRIVATE
	// BLOCK
	// TRANSACTION
	// COMBINATION
	AlertDocumentType string `json:"alertDocumentType"`
	// Type of finding
	//
	// UNKNOWN_TYPE
	//
	// EXPLOIT
	//
	// SUSPICIOUS
	//
	// DEGRADED
	//
	// INFO
	FindingType string `json:"findingType"`
	// List of alerts involved in the alert
	RelatedAlerts []string `json:"relatedAlerts"`
	// Alert chain id
	ChainId uint `json:"chainId"`
}

// GetAlertId returns getAlertsAlertsAlertsResponseAlertsAlert.AlertId, and is useful for accessing the field via an interface.
func (v *getAlertsAlertsAlertsResponseAlertsAlert) GetAlertId() string { return v.AlertId }

// GetAddresses returns getAlertsAlertsAlertsResponseAlertsAlert.Addresses, and is useful for accessing the field via an interface.
func (v *getAlertsAlertsAlertsResponseAlertsAlert) GetAddresses() []string { return v.Addresses }

// GetContracts returns getAlertsAlertsAlertsResponseAlertsAlert.Contracts, and is useful for accessing the field via an interface.
func (v *getAlertsAlertsAlertsResponseAlertsAlert) GetContracts() []*getAlertsAlertsAlertsResponseAlertsAlertContractsContract {
	return v.Contracts
}

// GetCreatedAt returns getAlertsAlertsAlertsResponseAlertsAlert.CreatedAt, and is useful for accessing the field via an interface.
func (v *getAlertsAlertsAlertsResponseAlertsAlert) GetCreatedAt() string { return v.CreatedAt }

// GetDescription returns getAlertsAlertsAlertsResponseAlertsAlert.Description, and is useful for accessing the field via an interface.
func (v *getAlertsAlertsAlertsResponseAlertsAlert) GetDescription() string { return v.Description }

// GetHash returns getAlertsAlertsAlertsResponseAlertsAlert.Hash, and is useful for accessing the field via an interface.
func (v *getAlertsAlertsAlertsResponseAlertsAlert) GetHash() string { return v.Hash }

// GetMetadata returns getAlertsAlertsAlertsResponseAlertsAlert.Metadata, and is useful for accessing the field via an interface.
func (v *getAlertsAlertsAlertsResponseAlertsAlert) GetMetadata() map[string]string { return v.Metadata }

// GetName returns getAlertsAlertsAlertsResponseAlertsAlert.Name, and is useful for accessing the field via an interface.
func (v *getAlertsAlertsAlertsResponseAlertsAlert) GetName() string { return v.Name }

// GetProjects returns getAlertsAlertsAlertsResponseAlertsAlert.Projects, and is useful for accessing the field via an interface.
func (v *getAlertsAlertsAlertsResponseAlertsAlert) GetProjects() []*getAlertsAlertsAlertsResponseAlertsAlertProjectsProject {
	return v.Projects
}

// GetProtocol returns getAlertsAlertsAlertsResponseAlertsAlert.Protocol, and is useful for accessing the field via an interface.
func (v *getAlertsAlertsAlertsResponseAlertsAlert) GetProtocol() string { return v.Protocol }

// GetScanNodeCount returns getAlertsAlertsAlertsResponseAlertsAlert.ScanNodeCount, and is useful for accessing the field via an interface.
func (v *getAlertsAlertsAlertsResponseAlertsAlert) GetScanNodeCount() int { return v.ScanNodeCount }

// GetSeverity returns getAlertsAlertsAlertsResponseAlertsAlert.Severity, and is useful for accessing the field via an interface.
func (v *getAlertsAlertsAlertsResponseAlertsAlert) GetSeverity() string { return v.Severity }

// GetSource returns getAlertsAlertsAlertsResponseAlertsAlert.Source, and is useful for accessing the field via an interface.
func (v *getAlertsAlertsAlertsResponseAlertsAlert) GetSource() *getAlertsAlertsAlertsResponseAlertsAlertSource {
	return v.Source
}

// GetAlertDocumentType returns getAlertsAlertsAlertsResponseAlertsAlert.AlertDocumentType, and is useful for accessing the field via an interface.
func (v *getAlertsAlertsAlertsResponseAlertsAlert) GetAlertDocumentType() string {
	return v.AlertDocumentType
}

// GetFindingType returns getAlertsAlertsAlertsResponseAlertsAlert.FindingType, and is useful for accessing the field via an interface.
func (v *getAlertsAlertsAlertsResponseAlertsAlert) GetFindingType() string { return v.FindingType }

// GetRelatedAlerts returns getAlertsAlertsAlertsResponseAlertsAlert.RelatedAlerts, and is useful for accessing the field via an interface.
func (v *getAlertsAlertsAlertsResponseAlertsAlert) GetRelatedAlerts() []string {
	return v.RelatedAlerts
}

// GetChainId returns getAlertsAlertsAlertsResponseAlertsAlert.ChainId, and is useful for accessing the field via an interface.
func (v *getAlertsAlertsAlertsResponseAlertsAlert) GetChainId() uint { return v.ChainId }

// getAlertsAlertsAlertsResponseAlertsAlertContractsContract includes the requested fields of the GraphQL type Contract.
// The GraphQL type's documentation follows.
//
// Smart Contract Information
type getAlertsAlertsAlertsResponseAlertsAlertContractsContract struct {
	// User-friendly name of the contract
	Name string `json:"name"`
	// Related web3 project
	ProjectId string `json:"projectId"`
}

// GetName returns getAlertsAlertsAlertsResponseAlertsAlertContractsContract.Name, and is useful for accessing the field via an interface.
func (v *getAlertsAlertsAlertsResponseAlertsAlertContractsContract) GetName() string { return v.Name }

// GetProjectId returns getAlertsAlertsAlertsResponseAlertsAlertContractsContract.ProjectId, and is useful for accessing the field via an interface.
func (v *getAlertsAlertsAlertsResponseAlertsAlertContractsContract) GetProjectId() string {
	return v.ProjectId
}

// getAlertsAlertsAlertsResponseAlertsAlertProjectsProject includes the requested fields of the GraphQL type Project.
// The GraphQL type's documentation follows.
//
// Web3 Project Information
type getAlertsAlertsAlertsResponseAlertsAlertProjectsProject struct {
	// Project identifier
	Id string `json:"id"`
}

// GetId returns getAlertsAlertsAlertsResponseAlertsAlertProjectsProject.Id, and is useful for accessing the field via an interface.
func (v *getAlertsAlertsAlertsResponseAlertsAlertProjectsProject) GetId() string { return v.Id }

// getAlertsAlertsAlertsResponseAlertsAlertSource includes the requested fields of the GraphQL type AlertSource.
// The GraphQL type's documentation follows.
//
// Source where the threat was detected
type getAlertsAlertsAlertsResponseAlertsAlertSource struct {
	// Transaction where the threat was detected
	TransactionHash string `json:"transactionHash"`
	// Bot that triggered the alert
	Bot *getAlertsAlertsAlertsResponseAlertsAlertSourceBot `json:"bot"`
	// Block where the threat was detected
	Block *getAlertsAlertsAlertsResponseAlertsAlertSourceBlock `json:"block"`
	// Source alert event the threat was detected
	SourceAlert *getAlertsAlertsAlertsResponseAlertsAlertSourceSourceAlertSourceAlertEvent `json:"sourceAlert"`
}

// GetTransactionHash returns getAlertsAlertsAlertsResponseAlertsAlertSource.TransactionHash, and is useful for accessing the field via an interface.
func (v *getAlertsAlertsAlertsResponseAlertsAlertSource) GetTransactionHash() string {
	return v.TransactionHash
}

// GetBot returns getAlertsAlertsAlertsResponseAlertsAlertSource.Bot, and is useful for accessing the field via an interface.
func (v *getAlertsAlertsAlertsResponseAlertsAlertSource) GetBot() *getAlertsAlertsAlertsResponseAlertsAlertSourceBot {
	return v.Bot
}

// GetBlock returns getAlertsAlertsAlertsResponseAlertsAlertSource.Block, and is useful for accessing the field via an interface.
func (v *getAlertsAlertsAlertsResponseAlertsAlertSource) GetBlock() *getAlertsAlertsAlertsResponseAlertsAlertSourceBlock {
	return v.Block
}

// GetSourceAlert returns getAlertsAlertsAlertsResponseAlertsAlertSource.SourceAlert, and is useful for accessing the field via an interface.
func (v *getAlertsAlertsAlertsResponseAlertsAlertSource) GetSourceAlert() *getAlertsAlertsAlertsResponseAlertsAlertSourceSourceAlertSourceAlertEvent {
	return v.SourceAlert
}

// getAlertsAlertsAlertsResponseAlertsAlertSourceBlock includes the requested fields of the GraphQL type Block.
// The GraphQL type's documentation follows.
//
// Block information
type getAlertsAlertsAlertsResponseAlertsAlertSourceBlock struct {
	// Block number
	Number uint `json:"number"`
	// Block hash
	Hash string `json:"hash"`
	// Block's timestamp
	Timestamp string `json:"timestamp"`
	// Block's chain id
	ChainId uint `json:"chainId"`
}

// GetNumber returns getAlertsAlertsAlertsResponseAlertsAlertSourceBlock.Number, and is useful for accessing the field via an interface.
func (v *getAlertsAlertsAlertsResponseAlertsAlertSourceBlock) GetNumber() uint { return v.Number }

// GetHash returns getAlertsAlertsAlertsResponseAlertsAlertSourceBlock.Hash, and is useful for accessing the field via an interface.
func (v *getAlertsAlertsAlertsResponseAlertsAlertSourceBlock) GetHash() string { return v.Hash }

// GetTimestamp returns getAlertsAlertsAlertsResponseAlertsAlertSourceBlock.Timestamp, and is useful for accessing the field via an interface.
func (v *getAlertsAlertsAlertsResponseAlertsAlertSourceBlock) GetTimestamp() string {
	return v.Timestamp
}

// GetChainId returns getAlertsAlertsAlertsResponseAlertsAlertSourceBlock.ChainId, and is useful for accessing the field via an interface.
func (v *getAlertsAlertsAlertsResponseAlertsAlertSourceBlock) GetChainId() uint { return v.ChainId }

// getAlertsAlertsAlertsResponseAlertsAlertSourceBot includes the requested fields of the GraphQL type Bot.
// The GraphQL type's documentation follows.
//
// Bot information
type getAlertsAlertsAlertsResponseAlertsAlertSourceBot struct {
	ChainIds    []string `json:"chainIds"`
	CreatedAt   string   `json:"createdAt"`
	Description string   `json:"description"`
	// Bot Developer address
	Developer string `json:"developer"`
	// Bot Documentation IPFS reference
	DocReference string `json:"docReference"`
	Enabled      bool   `json:"enabled"`
	// Bot id
	Id string `json:"id"`
	// Bot docker image hash
	Image string `json:"image"`
	Name  string `json:"name"`
	// Bot manifest IPFS reference
	Reference string `json:"reference"`
	// Bot code repository url
	Repository string `json:"repository"`
	// Projects the bot monitors
	Projects []string `json:"projects"`
	// List of scanNodes running the bot
	ScanNodes []string `json:"scanNodes"`
	Version   string   `json:"version"`
}

// GetChainIds returns getAlertsAlertsAlertsResponseAlertsAlertSourceBot.ChainIds, and is useful for accessing the field via an interface.
func (v *getAlertsAlertsAlertsResponseAlertsAlertSourceBot) GetChainIds() []string { return v.ChainIds }

// GetCreatedAt returns getAlertsAlertsAlertsResponseAlertsAlertSourceBot.CreatedAt, and is useful for accessing the field via an interface.
func (v *getAlertsAlertsAlertsResponseAlertsAlertSourceBot) GetCreatedAt() string { return v.CreatedAt }

// GetDescription returns getAlertsAlertsAlertsResponseAlertsAlertSourceBot.Description, and is useful for accessing the field via an interface.
func (v *getAlertsAlertsAlertsResponseAlertsAlertSourceBot) GetDescription() string {
	return v.Description
}

// GetDeveloper returns getAlertsAlertsAlertsResponseAlertsAlertSourceBot.Developer, and is useful for accessing the field via an interface.
func (v *getAlertsAlertsAlertsResponseAlertsAlertSourceBot) GetDeveloper() string { return v.Developer }

// GetDocReference returns getAlertsAlertsAlertsResponseAlertsAlertSourceBot.DocReference, and is useful for accessing the field via an interface.
func (v *getAlertsAlertsAlertsResponseAlertsAlertSourceBot) GetDocReference() string {
	return v.DocReference
}

// GetEnabled returns getAlertsAlertsAlertsResponseAlertsAlertSourceBot.Enabled, and is useful for accessing the field via an interface.
func (v *getAlertsAlertsAlertsResponseAlertsAlertSourceBot) GetEnabled() bool { return v.Enabled }

// GetId returns getAlertsAlertsAlertsResponseAlertsAlertSourceBot.Id, and is useful for accessing the field via an interface.
func (v *getAlertsAlertsAlertsResponseAlertsAlertSourceBot) GetId() string { return v.Id }

// GetImage returns getAlertsAlertsAlertsResponseAlertsAlertSourceBot.Image, and is useful for accessing the field via an interface.
func (v *getAlertsAlertsAlertsResponseAlertsAlertSourceBot) GetImage() string { return v.Image }

// GetName returns getAlertsAlertsAlertsResponseAlertsAlertSourceBot.Name, and is useful for accessing the field via an interface.
func (v *getAlertsAlertsAlertsResponseAlertsAlertSourceBot) GetName() string { return v.Name }

// GetReference returns getAlertsAlertsAlertsResponseAlertsAlertSourceBot.Reference, and is useful for accessing the field via an interface.
func (v *getAlertsAlertsAlertsResponseAlertsAlertSourceBot) GetReference() string { return v.Reference }

// GetRepository returns getAlertsAlertsAlertsResponseAlertsAlertSourceBot.Repository, and is useful for accessing the field via an interface.
func (v *getAlertsAlertsAlertsResponseAlertsAlertSourceBot) GetRepository() string {
	return v.Repository
}

// GetProjects returns getAlertsAlertsAlertsResponseAlertsAlertSourceBot.Projects, and is useful for accessing the field via an interface.
func (v *getAlertsAlertsAlertsResponseAlertsAlertSourceBot) GetProjects() []string { return v.Projects }

// GetScanNodes returns getAlertsAlertsAlertsResponseAlertsAlertSourceBot.ScanNodes, and is useful for accessing the field via an interface.
func (v *getAlertsAlertsAlertsResponseAlertsAlertSourceBot) GetScanNodes() []string {
	return v.ScanNodes
}

// GetVersion returns getAlertsAlertsAlertsResponseAlertsAlertSourceBot.Version, and is useful for accessing the field via an interface.
func (v *getAlertsAlertsAlertsResponseAlertsAlertSourceBot) GetVersion() string { return v.Version }

// getAlertsAlertsAlertsResponseAlertsAlertSourceSourceAlertSourceAlertEvent includes the requested fields of the GraphQL type SourceAlertEvent.
// The GraphQL type's documentation follows.
//
// Source alert information
type getAlertsAlertsAlertsResponseAlertsAlertSourceSourceAlertSourceAlertEvent struct {
	// Alert hash
	Hash string `json:"hash"`
	// Bot id
	BotId string `json:"botId"`
	// Alert timestamp
	Timestamp string `json:"timestamp"`
	// Source chain id
	ChainId string `json:"chainId"`
}

// GetHash returns getAlertsAlertsAlertsResponseAlertsAlertSourceSourceAlertSourceAlertEvent.Hash, and is useful for accessing the field via an interface.
func (v *getAlertsAlertsAlertsResponseAlertsAlertSourceSourceAlertSourceAlertEvent) GetHash() string {
	return v.Hash
}

// GetBotId returns getAlertsAlertsAlertsResponseAlertsAlertSourceSourceAlertSourceAlertEvent.BotId, and is useful for accessing the field via an interface.
func (v *getAlertsAlertsAlertsResponseAlertsAlertSourceSourceAlertSourceAlertEvent) GetBotId() string {
	return v.BotId
}

// GetTimestamp returns getAlertsAlertsAlertsResponseAlertsAlertSourceSourceAlertSourceAlertEvent.Timestamp, and is useful for accessing the field via an interface.
func (v *getAlertsAlertsAlertsResponseAlertsAlertSourceSourceAlertSourceAlertEvent) GetTimestamp() string {
	return v.Timestamp
}

// GetChainId returns getAlertsAlertsAlertsResponseAlertsAlertSourceSourceAlertSourceAlertEvent.ChainId, and is useful for accessing the field via an interface.
func (v *getAlertsAlertsAlertsResponseAlertsAlertSourceSourceAlertSourceAlertEvent) GetChainId() string {
	return v.ChainId
}

// getAlertsAlertsAlertsResponsePageInfoAlertPageInfo includes the requested fields of the GraphQL type AlertPageInfo.
type getAlertsAlertsAlertsResponsePageInfoAlertPageInfo struct {
	HasNextPage bool                                                                       `json:"hasNextPage"`
	EndCursor   *getAlertsAlertsAlertsResponsePageInfoAlertPageInfoEndCursorAlertEndCursor `json:"endCursor"`
}

// GetHasNextPage returns getAlertsAlertsAlertsResponsePageInfoAlertPageInfo.HasNextPage, and is useful for accessing the field via an interface.
func (v *getAlertsAlertsAlertsResponsePageInfoAlertPageInfo) GetHasNextPage() bool {
	return v.HasNextPage
}

// GetEndCursor returns getAlertsAlertsAlertsResponsePageInfoAlertPageInfo.EndCursor, and is useful for accessing the field via an interface.
func (v *getAlertsAlertsAlertsResponsePageInfoAlertPageInfo) GetEndCursor() *getAlertsAlertsAlertsResponsePageInfoAlertPageInfoEndCursorAlertEndCursor {
	return v.EndCursor
}

// getAlertsAlertsAlertsResponsePageInfoAlertPageInfoEndCursorAlertEndCursor includes the requested fields of the GraphQL type AlertEndCursor.
type getAlertsAlertsAlertsResponsePageInfoAlertPageInfoEndCursorAlertEndCursor struct {
	AlertId     string `json:"alertId"`
	BlockNumber uint   `json:"blockNumber"`
}

// GetAlertId returns getAlertsAlertsAlertsResponsePageInfoAlertPageInfoEndCursorAlertEndCursor.AlertId, and is useful for accessing the field via an interface.
func (v *getAlertsAlertsAlertsResponsePageInfoAlertPageInfoEndCursorAlertEndCursor) GetAlertId() string {
	return v.AlertId
}

// GetBlockNumber returns getAlertsAlertsAlertsResponsePageInfoAlertPageInfoEndCursorAlertEndCursor.BlockNumber, and is useful for accessing the field via an interface.
func (v *getAlertsAlertsAlertsResponsePageInfoAlertPageInfoEndCursorAlertEndCursor) GetBlockNumber() uint {
	return v.BlockNumber
}

// getAlertsResponse is returned by getAlerts on success.
type getAlertsResponse struct {
	// Fetches alerts
	Alerts *getAlertsAlertsAlertsResponse `json:"alerts"`
}

// GetAlerts returns getAlertsResponse.Alerts, and is useful for accessing the field via an interface.
func (v *getAlertsResponse) GetAlerts() *getAlertsAlertsAlertsResponse { return v.Alerts }

// The query or mutation executed by getAlerts.
const getAlerts_Operation = `
query getAlerts ($input: AlertsInput) {
	alerts(input: $input) {
		pageInfo {
			hasNextPage
			endCursor {
				alertId
				blockNumber
			}
		}
		alerts {
			alertId
			addresses
			contracts {
				name
				projectId
			}
			createdAt
			description
			hash
			metadata
			name
			projects {
				id
			}
			protocol
			scanNodeCount
			severity
			source {
				transactionHash
				bot {
					chainIds
					createdAt
					description
					developer
					docReference
					enabled
					id
					image
					name
					reference
					repository
					projects
					scanNodes
					version
				}
				block {
					number
					hash
					timestamp
					chainId
				}
				sourceAlert {
					hash
					botId
					timestamp
					chainId
				}
			}
			alertDocumentType
			findingType
			relatedAlerts
			chainId
		}
	}
}
`

func getAlerts(
	ctx context.Context,
	client graphql.Client,
	input *AlertsInput,
) (*getAlertsResponse, error) {
	req := &graphql.Request{
		OpName: "getAlerts",
		Query:  getAlerts_Operation,
		Variables: &__getAlertsInput{
			Input: input,
		},
	}
	var err error

	var data getAlertsResponse
	resp := &graphql.Response{Data: &data}

	err = client.MakeRequest(
		ctx,
		req,
		resp,
	)

	return &data, err
}
