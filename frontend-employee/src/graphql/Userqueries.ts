import { gql } from "@apollo/client";

export const SIGN_UP_MUTATION = gql`
  mutation SignUp($input: SignUpInput!) {
    signUp(input: $input) {
      id
      email
      firstName
      lastName
      role
      address
      phone
      isActive
      age
      userType
      pfp
      gender
      createdAt
      updatedAt
    }
  }
`;

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
      age
      userType
      pfp
      gender
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
