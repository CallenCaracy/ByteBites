import { gql } from "@apollo/client";

export const CREATE_INVENTORY = gql`
  mutation CreateInventory($menuId: ID!, $availableServings: Int!, $lowStockThreshold: Int) {
    createInventory(menuId: $menuId, availableServings: $availableServings, lowStockThreshold: $lowStockThreshold) {
        id
        menuId
        availableServings
        lowStockThreshold
        lastUpdated
    }
  }
`;

export const GET_INVENTORY = gql`
  query GetInventory($menuId: ID!) {
    inventory(id: $menuId) {
        id
        menuId
        availableServings
        lowStockThreshold
        lastUpdated
    }
  }
`;
