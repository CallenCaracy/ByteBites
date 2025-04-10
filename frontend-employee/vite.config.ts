import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react'
import tailwindcss from '@tailwindcss/vite'

// https://vite.dev/config/
export default defineConfig({
  // server: {
  //   host: '0.0.0.0',  // Exposes the server to the network
  //   port: 5173,        // You can change this if you need a different port
  // },
  plugins: [
    react(),
    tailwindcss(),
  ],
})
