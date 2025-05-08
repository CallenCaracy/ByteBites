import { Routes, Route, Navigate, useLocation } from 'react-router-dom';
import { useEffect, useRef } from 'react';
import ProtectedRoute from "./components/ProtectedRoute";
import LoginPage from './pages/LoginPage';
import Dashboard from './pages/Dashboard';
import RegisterPage from './pages/Register';
import Account from './pages/Account';
import ForgotPassword from './pages/ForgotPassword'
import UpdatePassword from './pages/UpdatePassword';
import MenuItem from './pages/MenuItems';
import AddMenu from './pages/AddMenu';
import MakeInventory from './pages/MakeInventory';
import OrderQueues from './pages/OrderQueues'
import { supabase } from './utils/supabaseClient';

function App() {
  const location = useLocation();
  const timeoutRef = useRef<ReturnType<typeof setTimeout> | null>(null);
  const protectedPaths = ["/dashboard", "/account", "/menu-item", "/order-queues", "/add-inventory", "/add-menu"];
  
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
    if (!protectedPaths.some(path => location.pathname.startsWith(path))) return;

    console.log("🔐 You're in a protected route");

    const scheduleRefresh = () => {
      if (timeoutRef.current) clearTimeout(timeoutRef.current);

      const expiresAt = Number(localStorage.getItem("expiresAt") || "0");
      const msUntilExpiry = expiresAt * 1000 - Date.now();
      const refreshIn = Math.max(msUntilExpiry - 60_000, 0); // refresh 1 min before

      timeoutRef.current = setTimeout(async () => {
        const { data, error } = await supabase.auth.refreshSession();
        if (error) {
          console.error("❌ Failed to refresh session:", error);
        } else {
          console.log("💧 Token refreshed", data.session);
          // Optional: Update expiresAt in localStorage here if needed
        }
      }, refreshIn);
    };

    // Initial setup
    scheduleRefresh();

    const { data: sub } = supabase.auth.onAuthStateChange((_event, session) => {
      if (session) {
        scheduleRefresh();
      }
    });

    return () => {
      if (timeoutRef.current) clearTimeout(timeoutRef.current);
      sub.subscription.unsubscribe();
    };
  }, [location.pathname]);

  return (
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
            <Route path="/menu-item/:menuId" element={<MenuItem/>} />
            <Route path="/add-menu/:userId" element={<AddMenu/>} />
            <Route path="/add-inventory/:menuId" element={<MakeInventory/>} />
            <Route path="/order-queues" element={<OrderQueues/>} />
          </Route>
      </Routes>
  );
}

export default App;
