import { useState, useEffect } from "react";
import { useNavigate } from 'react-router-dom';
import { supabase } from "../utils/supabaseClient";
import bg from "../assets/forgotpasswordbg.jpg";
import logo from "../assets/ByteBitesLogo/logo.png";

const UpdatePassword = () => {
  const [newPassword, setNewPassword] = useState("");
  const [loading, setLoading] = useState(false);
  const [success, setSuccess] = useState(false);
  const [error, setError] = useState<string | null>(null);
  const [accessToken, setAccessToken] = useState<string | null>(null);
  const [parsedToken, setParsedToken] = useState<{ access_token: string; refresh_token: string } | null>(null);

  const navigate = useNavigate();

  useEffect(() => {
    const storedToken = localStorage.getItem('sb-hzjjmfwrtvqjwxunfcue-auth-token');
    if (storedToken) {
      const parsed = JSON.parse(storedToken);
      setParsedToken(parsed);
      if (parsed) {
        setAccessToken(parsed.access_token);
      }
    } else {
      setError("Missing access token in localStorage.");
    }
  }, []);  

  console.log("Token", accessToken);

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
  
    if (!accessToken || !parsedToken) {
      setError("Access token or parsed token is not available.");
      return;
    }
  
    setLoading(true);
    setError(null);
  
    try {
      const { error: sessionError } = await supabase.auth.setSession({
        access_token: accessToken,
        refresh_token: parsedToken.refresh_token || "",
      });
  
      if (sessionError) {
        setError(sessionError.message);
        setLoading(false);
        return;
      }
  
      const { error } = await supabase.auth.updateUser({ password: newPassword });
  
      if (error) {
        setError(error.message);
      } else {
        setSuccess(true);
        setTimeout(() => {
          navigate("/login");
        }, 2000);
      }
    } catch (err) {
      setError("An error occurred while setting the session.");
    } finally {
      setLoading(false);
    }
  };

  return (
    <div className="py-16 flex justify-center items-center min-h-screen bg-gray-100">
      <div className="flex flex-col lg:flex-row bg-white rounded-lg shadow-lg overflow-hidden w-full max-w-4xl">
        <div className="hidden lg:block w-1/2">
          <img src={bg} alt="Background" className="w-full h-full object-cover" />
        </div>
        <form onSubmit={handleSubmit} className="flex flex-col justify-center p-8 w-full lg:w-1/2 space-y-6">
          <div className="flex items-center justify-center mb-6 w-full">
            <img src={logo} alt="Logo" className="h-20 w-20 object-contain" />
          </div>
          <h2 className="text-3xl font-semibold text-gray-700 text-center">Reset Password</h2>
          <p className="text-lg text-gray-600 text-center">Enter your new password below.</p>
  
          {success && (
            <p className="text-green-600 mb-4 text-center">
              Password updated successfully! Redirecting to login...
            </p>
          )}
          {error && (
            <p className="text-red-600 mb-4 text-center">{error}</p>
          )}
  
          <div className="mt-4">
            <label className="block text-gray-700 text-sm font-bold mb-2">New Password</label>
            <input
              type="password"
              placeholder="Enter your new password"
              value={newPassword}
              onChange={(e) => setNewPassword(e.target.value)}
              className="bg-gray-200 text-gray-700 focus:outline-none focus:shadow-outline border border-gray-300 rounded py-2 px-4 block w-full appearance-none"
              disabled={!accessToken || loading}
              autoComplete="new-password"
            />
          </div>
  
          <div className="mt-8">
            <button
              type="submit"
              className="bg-gray-700 text-white font-bold py-2 px-4 w-full rounded hover:bg-gray-600"
              disabled={!accessToken || loading}
            >
              {loading ? "Updating..." : "Update Password"}
            </button>
          </div>
  
          <p className="text-center text-gray-600 text-sm mt-4">
            Remembered your password?{" "}
            <span
              className="text-blue-500 cursor-pointer"
              onClick={() => navigate("/login")}
            >
              Login here
            </span>
          </p>
        </form>
      </div>
    </div>
  );
};

export default UpdatePassword;
