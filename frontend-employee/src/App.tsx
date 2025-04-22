import { BrowserRouter as Router, Routes, Route, Navigate } from 'react-router-dom';
import ProtectedRoute from "./components/ProtectedRoute";
import LoginPage from './pages/LoginPage';
import Dashboard from './pages/Dashboard';
import RegisterPage from './pages/Register';
import Account from './pages/Account';
import ForgotPassword from './pages/ForgotPassword'
import UpdatePassword from './pages/UpdatePassword';
import MenuItem from './pages/MenuItems';
import PaymentService from './pages/Payment';

function App() {
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
            <Route path="/menu-item/:menuId" element={<MenuItem/>} />
            <Route path="/payment/order-id" element={<PaymentService/>} />
          </Route>
      </Routes>
    </Router>
  );
}

export default App;
