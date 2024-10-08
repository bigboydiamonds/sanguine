import { createSlice, PayloadAction } from '@reduxjs/toolkit'

import {
  PendingBridgeTransaction,
  addPendingBridgeTransaction,
  removePendingBridgeTransaction,
  updatePendingBridgeTransaction,
  updatePendingBridgeTransactions,
  updateUserHistoricalTransactions,
  updateIsUserHistoricalTransactionsLoading,
  resetTransactionsState,
} from './actions'
import { BridgeTransaction } from '../api/generated'

export interface TransactionsState {
  userHistoricalTransactions: BridgeTransaction[]
  isUserHistoricalTransactionsLoading: boolean
  pendingBridgeTransactions: PendingBridgeTransaction[]
}

const initialState: TransactionsState = {
  userHistoricalTransactions: [],
  isUserHistoricalTransactionsLoading: true,
  pendingBridgeTransactions: [],
}

export const transactionsSlice = createSlice({
  name: 'transactions',
  initialState,
  reducers: {},
  extraReducers: (builder) => {
    builder
      .addCase(
        addPendingBridgeTransaction,
        (state, action: PayloadAction<PendingBridgeTransaction>) => {
          state.pendingBridgeTransactions = [
            action.payload,
            ...state.pendingBridgeTransactions,
          ]
        }
      )
      .addCase(
        updatePendingBridgeTransaction,
        (
          state,
          action: PayloadAction<{
            id: number
            timestamp: number
            transactionHash: string
            isSubmitted: boolean
          }>
        ) => {
          const { id, timestamp, transactionHash, isSubmitted } = action.payload
          const transactionIndex = state.pendingBridgeTransactions.findIndex(
            (transaction) => transaction.id === id
          )

          if (transactionIndex !== -1) {
            state.pendingBridgeTransactions =
              state.pendingBridgeTransactions.map((transaction, index) =>
                index === transactionIndex
                  ? { ...transaction, transactionHash, isSubmitted, timestamp }
                  : transaction
              )
          }
        }
      )
      .addCase(
        removePendingBridgeTransaction,
        (state, action: PayloadAction<number>) => {
          const idTimestampToRemove = action.payload
          state.pendingBridgeTransactions =
            state.pendingBridgeTransactions.filter(
              (transaction) => transaction.id !== idTimestampToRemove
            )
        }
      )
      .addCase(
        updatePendingBridgeTransactions,
        (state, action: PayloadAction<PendingBridgeTransaction[]>) => {
          state.pendingBridgeTransactions = action.payload
        }
      )
      .addCase(
        updateUserHistoricalTransactions,
        (state, action: PayloadAction<BridgeTransaction[]>) => {
          state.userHistoricalTransactions = action.payload
        }
      )
      .addCase(
        updateIsUserHistoricalTransactionsLoading,
        (state, action: PayloadAction<boolean>) => {
          state.isUserHistoricalTransactionsLoading = action.payload
        }
      )
      .addCase(resetTransactionsState, (state) => {
        state.pendingBridgeTransactions = initialState.pendingBridgeTransactions
        state.userHistoricalTransactions =
          initialState.userHistoricalTransactions
        state.isUserHistoricalTransactionsLoading =
          initialState.isUserHistoricalTransactionsLoading
      })
  },
})

export default transactionsSlice.reducer
