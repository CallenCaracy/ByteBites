import React, { createContext, useContext, useEffect, useState } from "react";
import { useQuery } from "@apollo/client";
import { GET_AUTHENTICATED_USER } from "../graphql/Userqueries";

export const AuthContext = createContext<any>(null);

export const AuthProvider: React.FC<{ children: React.ReactNode }> = ({ children }) => {
  const [tokenReady, setTokenReady] = useState(false);
  const { data, loading, error } = useQuery(GET_AUTHENTICATED_USER, {
    skip: !localStorage.getItem("accessToken"),
    fetchPolicy: "network-only",
  });
  const [user, setUser] = useState(null);

  useEffect(() => {
    if (data?.getAuthenticatedUser) {
      setUser(data.getAuthenticatedUser);
    }
    setTokenReady(true);
  }, [data]);

  if (!tokenReady) return <div>Loading Auth...</div>;

  return (
    <AuthContext.Provider value={{ user, loading, error }}>
      {children}
    </AuthContext.Provider>
  );
};

export const useAuth = () => useContext(AuthContext);