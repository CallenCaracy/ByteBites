import { gql } from "@apollo/client";

export const ADD_CART_ITEM = gql`
  mutation AddCartItem($input: AddCartItemInput!) {
    addCartItem(input: $input) {
      id
      cart_id
      menu_item_id
      quantity
      price
      customizations
      created_at
    }
  }
`;

// Unused
export const GET_CART = gql`
  query GetCart($user_id: ID!) {
    getCart(user_id: $user_id) {
      id
      user_id
      created_at
      updated_at
      items {
        id
        cart_id
        menu_item_id
        quantity
        price
        customizations
        created_at
      }
    }
  }
`;

// Unused
export const GET_CART_ITEMS = gql`
query GetCartItemsByCartID($cart_id: ID!) {
    getCartItemsByCartId(cart_id: $cart_id) {
      id
      cart_id
      menu_item_id
      quantity
      price
      customizations
      created_at
      updated_at
    }
  }
`;

export const GET_CART_AND_ITEMS = gql`
  query GetCartAndMenuItems($user_id: ID!) {
    getCartAndMenuItems(user_id: $user_id) {
      id
      user_id
      created_at
      updated_at
      items {
        id
        cart_id
        menu_item_id
        quantity
        price
        customizations
        created_at
        updated_at
        menuItem {
          id
          name
          description
          price
          image_url
        }
      }
    }
  }
`;

export const CREATE_ORDER_FROM_CART = gql`
  mutation CreateOrderFromCart(
    $cartID: ID!
    $userID: ID!
    $orderType: String!
    $deliveryAddress: String
    $specialRequests: String
  ) {
    createOrderFromCart(
      cartID: $cartID
      userID: $userID
      orderType: $orderType
      deliveryAddress: $deliveryAddress
      specialRequests: $specialRequests
    ) {
      id
      userID
      totalPrice
      orderStatus
      orderType
      deliveryAddress
      specialRequests
      createdAt
      updatedAt
      items {
        id
        orderID
        menuItemID
        quantity
        price
        customizations
        createdAt
      }
    }
  }
`;

export const UPDATE_CART_ITEM = gql`
  mutation UpdateCartItem($input: UpdateCartItemInput!) {
    updateCartItem(input: $input) {
      id
      cart_id
      menu_item_id
      quantity
      price
      customizations
      created_at
      updated_at
    }
  }
`;

export const REMOVE_CART_ITEM = gql`
    mutation RemoveCartItem($id: ID!) {
        removeCartItem(id: $id)
    }
`;