import { gql } from "@apollo/client";

export const GET_ORDER_QUEUES = gql`
  query {
    orderQueues {
      id
      orderId
      status
      createdAt
    }
  }
`;

export const ORDER_QUEUE_CREATED = gql`
  subscription {
    orderQueueCreated {
      id
      orderId
      status
      createdAt
    }
  }
`;