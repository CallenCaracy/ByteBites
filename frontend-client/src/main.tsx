import { StrictMode } from 'react'
import { createRoot } from 'react-dom/client'
import { ApolloProvider } from "@apollo/client";
// import './index.css'
import { client } from './apolloClient';
import { AuthProvider } from "./components/AuthContext";
import App from './App.tsx'

createRoot(document.getElementById('root')!).render(
  <StrictMode>
    <ApolloProvider client={client}>
      <AuthProvider>
        <App />
      </AuthProvider>
    </ApolloProvider>
  </StrictMode>,
)