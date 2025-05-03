import { gql } from "@apollo/client";

export const GET_MENU_ITEMS = gql`
  query GetAllMenuItems {
    getAllMenuItems {
      id
      name
      description
      price
      category
      availability_status
      image_url
      created_at
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
      category
      availability_status
      image_url
      created_at
    }
  }
`;

// export const ADD_MENU_ITEM = gql`
//   mutation AddMenuItem($input: NewMenuItem!) {
//     addMenuItem(input: $input) {
//       id
//       name
//       description
//       price
//       category
//       availability_status
//       image_url
//       created_at
//     }
//   }
// `;

export const CREATE_MENU_ITEM = gql`
  mutation CreateMenuItem($input: NewMenuItem!) {
    createMenuItem(input: $input) {
      id
      name
      description
      price
      category
      availability_status
      image_url
      createdAt
      updatedAt
    }
  }
`;


export const UPDATE_MENU_ITEM = gql`
  mutation UpdateMenuItem($id: ID!, $input: UpdateMenuItem!) {
    updateMenuItem(id: $id, input: $input) {
      id
      name
      description
      price
      category
      availability_status
      image_url
      updated_at
    }
  }
`;

// export const DELETE_MENU_ITEM = gql`
//   mutation DeleteMenuItem($id: ID!) {
//     deleteMenuItem(id: $id) {
//       id
//       name
//     }
//   }
// `;


export const DELETE_MENU_ITEM = gql`
  mutation DeleteMenuItem($id: ID!) {
    deleteMenuItem(id: $id)
  }
`;
