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

export const GET_USER_BY_ID = gql`
  query GetUserByID($id: ID!) {
    getUserById(id: $id) {
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
`

export const UPDATE_USER = gql`
  mutation UpdateUser($id: ID!, $input: UpdateUserInput!) {
    updateUser(id: $id, input: $input) {
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
`

export const CHECK_TOKEN = gql`
  query CheckToken {
    checkToken {
      id
      email
    }
  }
`
export const FORGOT_PASSWORD = gql `
  mutation ForgotPassword($email: String!) {
    forgotPassword(input: { email: $email }) {
      success
      message
    }
  }
`