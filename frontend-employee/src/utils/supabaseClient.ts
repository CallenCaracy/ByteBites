import { createClient } from '@supabase/supabase-js';

const supabaseUrl = import.meta.env.VITE_SUPABASE_URL as string;
const supabaseKey = import.meta.env.VITE_SUPABASE_ANON_KEY as string;

const accessToken = localStorage.getItem("accessToken");

export const supabase = createClient(supabaseUrl, supabaseKey, {
  global: {
    headers: {
      Authorization: accessToken ? `Bearer ${accessToken}` : '',
    },
  },
});
