package server

import (
	"context"

	abci "github.com/cometbft/cometbft/abci/types"

	servertypes "github.com/cosmos/cosmos-sdk/server/types"
)

type cometApplication struct {
	abci servertypes.ABCI
}

func NewCometApplication(abci servertypes.ABCI) abci.Application {
	return cometApplication{abci: abci}
}

func (a cometApplication) Info(_ context.Context, req *abci.RequestInfo) (*abci.ResponseInfo, error) {
	return a.abci.Info(req)
}

func (a cometApplication) Query(ctx context.Context, req *abci.RequestQuery) (*abci.ResponseQuery, error) {
	return a.abci.Query(ctx, req)
}

func (a cometApplication) CheckTx(_ context.Context, req *abci.RequestCheckTx) (*abci.ResponseCheckTx, error) {
	return a.abci.CheckTx(req)
}

func (a cometApplication) InitChain(_ context.Context, req *abci.RequestInitChain) (*abci.ResponseInitChain, error) {
	return a.abci.InitChain(req)
}

func (a cometApplication) PrepareProposal(_ context.Context, req *abci.RequestPrepareProposal) (*abci.ResponsePrepareProposal, error) {
	return a.abci.PrepareProposal(req)
}

func (a cometApplication) ProcessProposal(_ context.Context, req *abci.RequestProcessProposal) (*abci.ResponseProcessProposal, error) {
	return a.abci.ProcessProposal(req)
}

func (a cometApplication) FinalizeBlock(_ context.Context, req *abci.RequestFinalizeBlock) (*abci.ResponseFinalizeBlock, error) {
	return a.abci.FinalizeBlock(req)
}

func (a cometApplication) ExtendVote(ctx context.Context, req *abci.RequestExtendVote) (*abci.ResponseExtendVote, error) {
	return a.abci.ExtendVote(ctx, req)
}

func (a cometApplication) VerifyVoteExtension(_ context.Context, req *abci.RequestVerifyVoteExtension) (*abci.ResponseVerifyVoteExtension, error) {
	return a.abci.VerifyVoteExtension(req)
}

func (a cometApplication) Commit(_ context.Context, _ *abci.RequestCommit) (*abci.ResponseCommit, error) {
	return a.abci.Commit()
}

func (a cometApplication) ListSnapshots(_ context.Context, req *abci.RequestListSnapshots) (*abci.ResponseListSnapshots, error) {
	return a.abci.ListSnapshots(req)
}

func (a cometApplication) OfferSnapshot(_ context.Context, req *abci.RequestOfferSnapshot) (*abci.ResponseOfferSnapshot, error) {
	return a.abci.OfferSnapshot(req)
}

func (a cometApplication) LoadSnapshotChunk(_ context.Context, req *abci.RequestLoadSnapshotChunk) (*abci.ResponseLoadSnapshotChunk, error) {
	return a.abci.LoadSnapshotChunk(req)
}

func (a cometApplication) ApplySnapshotChunk(_ context.Context, req *abci.RequestApplySnapshotChunk) (*abci.ResponseApplySnapshotChunk, error) {
	return a.abci.ApplySnapshotChunk(req)
}
