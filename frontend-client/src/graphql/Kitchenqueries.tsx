import { gql } from "@apollo/client";

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
