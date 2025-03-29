import { ApolloClient, InMemoryCache, createHttpLink } from "@apollo/client";

const client = new ApolloClient({
  link: createHttpLink({
    uri: "http://localhost:8080/query",
    credentials: "include",
  }),
  cache: new InMemoryCache(),
});

export default client;
