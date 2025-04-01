import { gql } from "@apollo/client";

export const SIGN_IN_MUTATION = gql`
  mutation SignInOnlyEmployee($email: String!, $password: String!) {
    signInOnlyEmployee(input: { email: $email, password: $password }) {
      accessToken
      refreshToken
      error
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
