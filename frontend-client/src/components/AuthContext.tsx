import React, { createContext, useContext, useEffect, useState } from "react";
import { useQuery } from "@apollo/client";
import { GET_AUTHENTICATED_USER } from "../../../frontend-client/src/graphql/Userqueries";

export const AuthContext = createContext<any>(null);

export const AuthProvider: React.FC<{ children: React.ReactNode }> = ({ children }) => {
  const { data, loading, error } = useQuery(GET_AUTHENTICATED_USER);
  const [user, setUser] = useState(null);

  useEffect(() => {
    if (data?.getAuthenticatedUser) {
      setUser(data.getAuthenticatedUser);
    }
  }, [data]);

  return (
    <AuthContext.Provider value={{ user, loading, error }}>
      {children}
    </AuthContext.Provider>
  );
};

export const useAuth = () => useContext(AuthContext);