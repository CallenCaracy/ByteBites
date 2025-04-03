import { gql } from "@apollo/client";

export const SIGN_UP_MUTATION = gql`
  mutation SignUp(
  $email: String!
  $password: String!
  $firstName: String!
  $lastName: String!
  $role: String!
  $address: String
  $phone: String
) {
  signUp(
    input: {
      email: $email
      password: $password
      firstName: $firstName
      lastName: $lastName
      role: $role
      address: $address
      phone: $phone
    }
  ) {
    id
    email
    firstName
    lastName
    role
    address
    phone
    isActive
    createdAt
  }
}
`

export const SIGN_IN_MUTATION = gql`
  mutation SignInOnlyEmployee($email: String!, $password: String!) {
    signInOnlyEmployee(input: { email: $email, password: $password }) {
      accessToken
      refreshToken
    }
  }
`;

export const GET_AUTHENTICATED_USER = gql`
  query GetAuthenticatedUser {
    getAuthenticatedUser {
      id
      email
      firstName
      lastName
      role
      address
      phone
      isActive
      createdAt
      updatedAt
    }
  }
`;

export const SIGN_OUT_USER = gql`
  mutation SignOut {
    signOut
  }
`
