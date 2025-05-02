import { gql } from '@apollo/client';

export const CREATE_TRANSACTION = gql`
  mutation CreateTransaction($orderId: UUID!, $userId: UUID!, $amountPaid: Float!, $paymentMethod: PaymentMethod!) {
    createTransaction(orderId: $orderId, userId: $userId, amountPaid: $amountPaid, paymentMethod: $paymentMethod) {
      id
      orderId
      userId
      amountPaid
      paymentMethod
      transactionStatus
      transactionTimestamp
    }
  }
`;

export const UPDATE_TRANSACTION_STATUS = gql`
  mutation UpdateTransactionStatus($id: UUID!, $status: TransactionStatus!) {
    updateTransactionStatus(id: $id, status: $status) {
      id
      transactionStatus
    }
  }
`;

export const CREATE_RECEIPT = gql`
  mutation CreateReceipt($transactionId: UUID!, $userId: UUID!, $amount: Float!, $paymentMethod: String!) {
    createReceipt(transactionId: $transactionId, userId: $userId, amount: $amount, paymentMethod: $paymentMethod) {
      id
      transactionId
      userId
      amount
      paymentMethod
      timestamp
    }
  }
`;

export const GET_TRANSACTION = gql`
  query GetTransaction($id: UUID!) {
    getTransaction(id: $id) {
      id
      orderId
      userId
      amountPaid
      paymentMethod
      transactionStatus
      transactionTimestamp
    }
  }
`;

export const GET_ALL_TRANSACTIONS = gql`
  query GetAllTransactions {
    getAllTransactions {
      id
      orderId
      userId
      amountPaid
      paymentMethod
      transactionStatus
      transactionTimestamp
    }
  }
`;