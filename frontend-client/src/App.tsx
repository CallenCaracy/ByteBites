import { BrowserRouter as Router, Routes, Route, Navigate } from 'react-router-dom';
import { useEffect } from 'react';
import ProtectedRoute from "./components/ProtectedRoute";
import LoginPage from './pages/LoginPage';
import Dashboard from './pages/Dashboard';
import RegisterPage from './pages/Register';
import Account from './pages/Account';
import ForgotPassword from './pages/ForgotPassword'
import UpdatePassword from './pages/UpdatePassword';
import MenuItem from './pages/MenuItems';
import YourCart from './pages/YourCart';
import PaymentPage from './pages/PaymentPage';
import { supabase } from './utils/supabaseClient';

function App() {
  useEffect(() => {
    const accessToken  = localStorage.getItem("accessToken");
    const refreshToken = localStorage.getItem("refreshToken");

    if (accessToken && refreshToken) {
      supabase.auth.setSession({
        access_token:  accessToken,
        refresh_token: refreshToken,
      }).then(({ data, error }) => {
        console.log("setSession result:", data, error);
        if (data?.session) {
          localStorage.setItem("expiresAt", (data.session.expires_at ?? 0).toString());
        }
      });
    }

    const { data: authListener } = supabase.auth.onAuthStateChange(
      (_event, session) => {
        if (session) {
          localStorage.setItem("accessToken",  session.access_token);
          localStorage.setItem("refreshToken", session.refresh_token);
          localStorage.setItem("expiresAt",    (session.expires_at ?? 0).toString());
        }
      }
    );

    return () => {
      authListener.subscription.unsubscribe();
    };
  }, []);

  useEffect(() => {
    let timeoutId: ReturnType<typeof setTimeout>;

    const scheduleRefresh = () => {
      const expiresAt = Number(localStorage.getItem("expiresAt") || "0");
      const msUntilExpiry = expiresAt * 1000 - Date.now();
      const refreshIn = Math.max(msUntilExpiry - 60_000, 0);

      timeoutId = setTimeout(async () => {
        const { data, error } = await supabase.auth.refreshSession();
        if (error) {
          console.error("Failed to refresh session:", error);
        } else {
          console.log("💧 Token refreshed", data.session);
        }
      }, refreshIn);
    };

    scheduleRefresh();

    // Also re-schedule when session changes
    const { data: sub } = supabase.auth.onAuthStateChange((_e, session) => {
      if (session) {
        scheduleRefresh();
      }
    });

    return () => {
      clearTimeout(timeoutId);
      sub.subscription.unsubscribe();
    };
  }, []);

  return (
    <Router>
      <Routes>
        {/* Public Routes */}
        <Route path="/" element={<Navigate to="/login"/>} />
        <Route path="/login" element={<LoginPage />} />
        <Route path="/register" element={<RegisterPage/>} />
        <Route path="/forgot" element={<ForgotPassword/>} />
        <Route path="/reset/:token" element={<UpdatePassword/>} />

          {/* Protected Routes */}
          <Route element={<ProtectedRoute />}>
            <Route path="/dashboard" element={<Dashboard/>} />
            <Route path="/account/:userId" element={<Account/>} />
            <Route path="/menu-item/:userId/:menuId" element={<MenuItem/>} />
            <Route path="/cart/:userId" element={<YourCart/>} />
            <Route path="/payment/:orderId" element={<PaymentPage />} />
          </Route>
      </Routes>
    </Router>
  );
}

export default App;
