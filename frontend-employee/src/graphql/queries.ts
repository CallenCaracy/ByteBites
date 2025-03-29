import { gql } from "@apollo/client";

export const SIGN_IN_MUTATION = gql`
  mutation SignIn($email: String!, $password: String!) {
    signIn(input: { email: $email, password: $password }) {
      accessToken
      refreshToken
      error
    }
  }
`;
