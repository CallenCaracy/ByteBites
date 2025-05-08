import { StrictMode } from 'react'
import { createRoot } from 'react-dom/client'
import { ApolloProvider } from "@apollo/client";
import './index.css'
import { client } from './apolloClient';
import App from './App.tsx'
import { BrowserRouter } from 'react-router-dom';
import { AuthProvider } from './components/AuthContext.tsx';

createRoot(document.getElementById('root')!).render(
  <StrictMode>
    <ApolloProvider client={client}>
      <BrowserRouter>
          <AuthProvider>
            <App />
          </AuthProvider>
        </BrowserRouter>
    </ApolloProvider>
  </StrictMode>,
)
