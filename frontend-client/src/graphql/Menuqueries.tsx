import { gql } from "@apollo/client";

export const GET_MENU_ITEMS = gql`
  query GetAllMenuItems {
    getAllMenuItems {
      id
      name
      description
      price
      discounted_price
      category
      discount
      availability_status
      image_url
      created_at
      updated_at
    }
  }
`;

export const GET_MENU_ITEM_BY_ID = gql`
  query GetMenuItemById($id: ID!) {
    getMenuItemById(id: $id) {
      id
      name
      description
      price
      discounted_price
      category
      discount
      availability_status
      image_url
      created_at
      updated_at
    }
  }
`;

export const GET_MENU_ITEM_BY_ID_FOR_CART = gql`
  query GetMenuItemById($id: ID!) {
    getMenuItemById(id: $id) {
      id
      name
      price
      category
      discount
      availability_status
      image_url
    }
  }
`;

export const GET_MENU_ITEM_BY_ID_FOR_GETTING_MENU_NAME = gql`
  query GetMenuItemsByIds($ids: [ID!]!) {
    getMenuItemsByIds(ids: $ids) {
      id
      name
    }
  }
`;

export const MENU_ITEM_CREATED = gql`
  subscription {
    menuItemCreated {
      id
      name
      description
      price
      discounted_price
      category
      discount
      availability_status
      image_url
      created_at
      updated_at
    }
  }
`;