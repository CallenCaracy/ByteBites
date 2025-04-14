import { ApolloClient, InMemoryCache, createHttpLink } from "@apollo/client";
import { setContext } from "@apollo/client/link/context";

const httpLink = createHttpLink({
  uri: "http://localhost:8080/query",
  credentials: "include",
});

const authLink = setContext((_, { headers }) => {
  const token = localStorage.getItem("accessToken");
  return {
      headers: {
          ...headers,
          Authorization: token ? `Bearer ${token}` : "",
      }
  };
});

export const client = new ApolloClient({
  link: authLink.concat(httpLink),
  cache: new InMemoryCache(),
});

export default client;
